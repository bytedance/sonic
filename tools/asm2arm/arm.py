#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import string
import argparse
import itertools
import functools

from typing import Any
from typing import Dict
from typing import List
from typing import Type
from typing import Tuple
from typing import Union
from typing import Callable
from typing import Iterable
from typing import Optional

import mcasm

class InstrStreamer(mcasm.Streamer):
    data:   bytes
    instr:  Optional[mcasm.mc.Instruction]
    fixups: List[mcasm.mc.Fixup]

    def __init__(self):
        self.data = b''
        self.instr = None
        self.fixups = []
        super().__init__()

    def unhandled_event(self, name: str, base_impl, *args, **kwargs):
        if name == 'emit_instruction':
            self.instr = args[1]
            self.data = args[2]
            self.fixups = args[3]
        return super().unhandled_event(name, base_impl, *args, **kwargs)

### Instruction Parser (GAS Syntax) ###
class Token:
    tag: int
    val: Union[int, str]

    def __init__(self, tag: int, val: Union[int, str]):
        self.val = val
        self.tag = tag

    @classmethod
    def end(cls):
        return cls(TOKEN_END, '')

    @classmethod
    def reg(cls, reg: str):
        return cls(TOKEN_REG, reg)

    @classmethod
    def imm(cls, imm: int):
        return cls(TOKEN_IMM, imm)

    @classmethod
    def num(cls, num: int):
        return cls(TOKEN_NUM, num)

    @classmethod
    def name(cls, name: str):
        return cls(TOKEN_NAME, name)

    @classmethod
    def punc(cls, punc: str):
        return cls(TOKEN_PUNC, punc)

    def __repr__(self):
        if self.tag == TOKEN_END:
            return '<END>'
        elif self.tag == TOKEN_REG:
            return '<REG %s>' % self.val
        elif self.tag == TOKEN_IMM:
            return '<IMM %d>' % self.val
        elif self.tag == TOKEN_NUM:
            return '<NUM %d>' % self.val
        elif self.tag == TOKEN_NAME:
            return '<NAME %s>' % repr(self.val)
        elif self.tag == TOKEN_PUNC:
            return '<PUNC %s>' % repr(self.val)
        else:
            return '<UNK:%d %r>' % (self.tag, self.val)

class Label:
    name: str
    offs: Optional[int]

    def __init__(self, name: str):
        self.name = name
        self.offs = None

    def __str__(self):
        return self.name

    def __repr__(self):
        if self.offs is None:
            return '{LABEL %s (unresolved)}' % self.name
        else:
            return '{LABEL %s (offset: %d)}' % (self.name, self.offs)

    def resolve(self, offs: int):
        self.offs = offs

class Index:
    base  : 'Register'
    scale : int

    def __init__(self, base: 'Register', scale: int = 1):
        self.base  = base
        self.scale = scale

    def __str__(self):
        if self.scale == 1:
            return ',%s' % self.base
        elif self.scale >= 2:
            return ',%s,%d' % (self.base, self.scale)
        else:
            raise RuntimeError('invalid parser state: invalid scale')

    def __repr__(self):
        if self.scale == 1:
            return repr(self.base)
        elif self.scale >= 2:
            return '%d * %r' % (self.scale, self.base)
        else:
            raise RuntimeError('invalid parser state: invalid scale')

class Memory:
    base  : Optional['Register']
    disp  : Optional['Displacement']
    index : Optional[Index]

    def __init__(self, base: Optional['Register'], disp: Optional['Displacement'], index: Optional[Index]):
        self.base  = base
        self.disp  = disp
        self.index = index
        self._validate()

    def __str__(self):
        return '%s(%s%s)' % (
            '' if self.disp  is None else self.disp,
            '' if self.base  is None else self.base,
            '' if self.index is None else self.index
        )

    def __repr__(self):
        return '{MEM %r%s%s}' % (
            '' if self.base  is None else self.base,
            '' if self.index is None else ' + ' + repr(self.index),
            '' if self.disp  is None else ' + ' + repr(self.disp)
        )

    def _validate(self):
        if self.base is None and self.index is None:
            raise SyntaxError('either base or index must be specified')

class Register:
    reg: str

    def __init__(self, reg: str):
        self.reg = reg.lower()

    def __str__(self):
        return '%' + self.reg

    def __repr__(self):
        return '{REG %s}' % self.reg

class Immediate:
    val: int
    ref: str

    def __init__(self, val: int):
        self.ref = ''
        self.val = val

    def __str__(self):
        return '$%d' % self.val

    def __repr__(self):
        return '{IMM bin:%s, oct:%s, dec:%d, hex:%s}' % (
            bin(self.val)[2:],
            oct(self.val)[2:],
            self.val,
            hex(self.val)[2:],
        )

class Reference:
    ref: str
    disp: int
    off: Optional[int]

    def __init__(self, ref: str, disp: int = 0):
        self.ref = ref
        self.disp = disp
        self.off = None

    def __str__(self):
        if self.off is None:
            return self.ref
        else:
            return '$' + str(self.off)

    def __repr__(self):
        if self.off is None:
            return '{REF %s + %d (unresolved)}' % (self.ref, self.disp)
        else:
            return '{REF %s + %d (offset: %d)}' % (self.ref, self.disp, self.off)

    @property
    def offset(self) -> int:
        if self.off is None:
            raise SyntaxError('unresolved reference to ' + repr(self.ref))
        else:
            return self.off

    def resolve(self, off: int):
        self.off = self.disp + off

Operand = Union[
    Label,
    Memory,
    Register,
    Immediate,
    Reference,
]

Displacement = Union[
    Immediate,
    Reference,
]

TOKEN_END  = 0
TOKEN_REG  = 1
TOKEN_IMM  = 2
TOKEN_NUM  = 3
TOKEN_NAME = 4
TOKEN_PUNC = 5

ARM_ADRP_IMM_BIT_SIZE = 21
ARM_ADR_WIDTH = 1024 * 1024

class Instruction:
    comments:   str
    mnemonic:   str
    asm_code:   str
    data:       bytes
    instr:      Optional[mcasm.mc.Instruction]
    fixups:     List[mcasm.mc.Fixup]
    offs_:      Optional[int]
    ADRP_label: Optional[str]
    text_label: Optional[str]
    back_label: Optional[str]
    ADR_instr:  Optional[str]
    adrp_asm:   Optional[str]
    is_adrp:    bool

    def __init__(self, line: str, adrp_count=0):
        self.comments = ''
        self.offs_ = None
        self.is_adrp = False
        self.asm = mcasm.Assembler('aarch64-apple-macos11')

        self.parse(line, adrp_count)

    def __str__(self):
        return self.asm_code

    def __repr__(self):
        return '{INSTR %s}' % ( self.asm_code )
    
    @property
    def jmptab(self) -> Optional[str]:
        if self.is_adrp and self.label_name.find(CLANG_JUMPTABLE_LABLE) != -1:
            return self.label_name

    @property
    def size(self) -> int:
        return len(self.data)

    @functools.cached_property
    def label_name(self) -> Optional[str]:
        if len(self.fixups) > 1:
            raise RuntimeError('has more than 1 fixup: ' + self.asm_code)
        if self.need_reloc:
            if self.mnemonic == 'adr':
                return self.fixups[0].value.sub_expr.symbol.name
            else:
                return self.fixups[0].value.symbol.name
        else:
            return None

    @functools.cached_property
    def is_branch(self) -> bool:
        return self.instr.desc.is_branch or self.is_invoke

    @functools.cached_property
    def is_return(self) -> bool:
        return self.instr.desc.is_return

    @functools.cached_property
    def is_jmpq(self) -> bool:
        # return self.mnemonic == 'br'
        return False

    @functools.cached_property
    def is_jmp(self) -> bool:
        return self.mnemonic == 'b'

    @functools.cached_property
    def is_invoke(self) -> bool:
        return self.instr.desc.is_call

    @property
    def is_branch_label(self) -> bool:
        return self.is_branch and (len(self.fixups) != 0)

    @property
    def need_reloc(self) -> bool:
        return (len(self.fixups) != 0)

    def set_label_offset(self, off):
        # arm64
        self.offs_ = off + 4

    # def _encode_normal_instr(self) -> str:
    #     return self.encode(self.data, self.asm_code)

    @functools.cached_property
    def encoded(self) -> str:
        if self.need_reloc:
            return self._encode_reloc_instr()
        else:
            return self._encode_normal_instr()

    def _check_offs_is_valid(self, bit_size: int):
        if abs(self.offs_) > (1 << bit_size):
            raise RuntimeError('offset is too larger, [assembly]: %s, [offset]: %d, [valid off size]: %d'
                % (self.asm_code, self.offs_, self.fixups[0].kind_info.bit_size))

    def _encode_adr(self):
        buf = int.from_bytes(self.data, byteorder='little')
        bit_size = ARM_ADRP_IMM_BIT_SIZE

        self._check_offs_is_valid(bit_size)

        # adrp op: | op | immlo | 1 0 0 0 0 | immhi | Rd |
        #          |31  |30   29|28       24|23    5|4  0|
        imm_lo = (self.offs_ << 29) & 0x60000000
        imm_hi = (self.offs_ << 3) & 0x00FFFFE0
        encode_data = (buf + imm_lo + imm_hi).to_bytes(4, byteorder='little')
        self.data = encode_data
        # return self.encode(encode_data, '%s $%s(%%rip)' % (str(self), self.offs_))

    def _encode_rel32(self):
        if self.mnemonic == 'adrp' or self.mnemonic == 'adr':
            return self._encode_adr()
        buf = int.from_bytes(self.data, byteorder='little')

        imm = self.offs_
        imm_size = self.fixups[0].kind_info.bit_size
        imm_offset = self.fixups[0].kind_info.bit_offset
        if self.fixups[0].kind_info.is_pc_rel == 1:
            # except adr and adrp, other PC-releative instructions need times 4
            imm = imm >> 2
            # immediate bit size has 1-bit for sign
            self._check_offs_is_valid(imm_size - 1 + 2)
        else:
            self._check_offs_is_valid(imm_size)

        imm = imm << imm_offset
        mask = (0x1 << (imm_size + imm_offset)) - 1
        buf = buf | (imm & mask)
        buf = buf.to_bytes(4, byteorder='little')
        self.data = buf
        # return self.encode(buf, '%s $%s(%%rip)' % (str(self), self.offs_))

    def _encode_page(self):
        if self.mnemonic != 'adrp':
            raise RuntimeError("not adrp instruction: %s" % self.asm_code)
        self.offs_ = self.offs_ >> 12
        return self._encode_rel32()

    def _encode_pageoff(self):
        self.offs_ = 0
        return self._encode_rel32()

    def _fixup_rel32(self):
        if self.offs_ is None:
            raise RuntimeError('unresolved label %s' % self.label_name)

        if self.mnemonic == 'adr':
            self._encode_adr()
        elif self.fixups[0].value.variant_kind == mcasm.mc.SymbolRefExpr.VariantKind.PAGEOFF:
            self._encode_pageoff()
        elif self.fixups[0].value.variant_kind == mcasm.mc.SymbolRefExpr.VariantKind.PAGE:
            self._encode_page()
        else:
            self._encode_rel32()

    def _encode_reloc_instr(self) -> str:
        self._fixup_rel32()
        return self.encode(self.data, '%s $%s(%%rip)' % (str(self), self.offs_))

    def _encode_normal_instr(self) -> str:
        return self.encode(self.data, str(self))

    def _raw_instr(self) -> bytes:
        if self.need_reloc:
            self._fixup_rel32()
        return self.data

    def _fixup_adrp(self, line: str, adrp_count: int) -> str:
        reg = line.split()[1].split(',')[0]
        self.text_label = line.split()[2].split('@')[0]
        self.ADRP_label = self.text_label + '_' + reg + '_' + str(adrp_count)
        self.back_label = '_back_adrp_' + str(adrp_count)
        self.ADR_instr = 'adr ' + reg + ', ' + self.text_label
        self.adrp_asm = line
        self.is_adrp = True
        line = 'b ' + self.ADRP_label
        self.asm_code = line + ' // ' + self.adrp_asm

        return line

    def _parse_by_mcasm(self, line: str):
        streamer = InstrStreamer()
        # self.asm.assemble(streamer, line, MCPU="", features_str="")
        self.asm.assemble(streamer, line)
        if streamer.instr is None:
            raise RuntimeError('cannot parse assembly: %s' % line)
        self.instr = streamer.instr

        # instead of short jump instruction
        self.data = streamer.data

        self.fixups = streamer.fixups
        self.mnemonic = line.split()[0]

    def convert_to_adr(self):
        self.is_adrp = True
        adr_asm = self.adrp_asm.replace('adrp', 'adr')
        self.asm_code = adr_asm + ' // ' + self.adrp_asm
        # self._parse_by_mcasm(adr_asm)
        return self.asm_code

    def parse(self, line: str, adrp_count: int):
        # machine code
        menmonic = line.split()[0]

        self.ADRP_label = None
        self.text_label = None
        # turn adrp to jmp
        if (menmonic == 'adrp'):
            line = self.convert_to_adr()
        else:
            self.asm_code = line

        self._parse_by_mcasm(line)

    @staticmethod
    def encode(buf: bytes, comments: str = '') -> str:
        i = 0
        r = []
        n = len(buf)

        # @debug
        # while i < n - 3:
        #     r.append('%08x' % int.from_bytes(buf[i:i + 4], 'little'))
        #     i += 4
        # return '\n\t'.join(r)

        if (n % 4 != 0):
            raise RuntimeError("Unkown instruction which not encoding 4 bytes: %s " % comments, buf)

        while i < n - 3:
            r.append('WORD $0x%08x' % int.from_bytes(buf[i:i + 4], 'little'))
            i += 4

        # join them together, and attach the comment if any
        if not comments:
            return '; '.join(r)
        else:
            return '%s  // %s' % ('; '.join(r), comments)

    Reg  = Optional[Register]
    Disp = Optional[Displacement]

### Prototype Parser ###

ARGS_ORDER_C = [
    Register('x0'),
    Register('x1'),
    Register('x2'),
    Register('x3'),
    Register('x4'),
    Register('x5'),
    Register('x6'),
    Register('x7'),
]

ARGS_ORDER_GO = [
    Register('R0'),
    Register('R1'),
    Register('R2'),
    Register('R3'),
    Register('R4'),
    Register('R5'),
    Register('R6'),
    Register('R7'),
]

FPARGS_ORDER = [
    Register('D0'),
    Register('D1'),
    Register('D2'),
    Register('D3'),
    Register('D4'),
    Register('D5'),
    Register('D6'),
    Register('D7'),
]

class Parameter:
    name : str
    size : int
    creg : Register
    goreg: Register

    def __init__(self, name: str, size: int, reg: Register, goreg: Register):
        self.creg  = reg
        self.goreg = reg
        self.name = name
        self.size = size

    def __repr__(self):
        return '<ARG %s(%d): %s>' % (self.name, self.size, self.creg)

class Pcsp:
    entry: int
    maxpc: int
    out  : List[Tuple[int, int]]
    pc   : int
    sp   : int

    def __init__(self, entry: int):
        self.out = []
        self.maxpc = entry
        self.entry = entry
        self.pc = entry
        self.sp = 0

    def __str__(self) -> str:
        ret = '[][2]uint32{\n'
        for pc, sp in self.out:
            ret += '        {%d, %d},\n' % (pc, sp)
        return ret + '    }'

    def optimize(self):
        # push the last record
        self.out.append((self.pc - self.entry, self.sp))
        # sort by pc
        self.out.sort(key=lambda x: x[0])
        # NOTICE: first pair {1, 0} to be compitable with golang
        tmp = [(1, 0)]
        lpc, lsp = 0, -1
        for pc, sp in self.out:
            # sp changed, push new record
            if pc != lpc and sp != lsp:
                    tmp.append((pc, sp))
            # sp unchanged, replace with the higher pc
            if pc != lpc and sp == lsp:
                if len(tmp) > 0:
                    tmp.pop(-1)
                tmp.append((pc, sp))

            lpc, lsp = pc, sp
        self.out = tmp

    def update(self, dpc: int, dsp: int):
        self.out.append((self.pc - self.entry, self.sp))
        self.pc += dpc
        self.sp += dsp
        if self.pc > self.maxpc:
            self.maxpc = self.pc

class Prototype:
    args: List[Parameter]
    retv: Optional[Parameter]

    def __init__(self, retv: Optional[Parameter], args: List[Parameter]):
        self.retv = retv
        self.args = args

    def __repr__(self):
        if self.retv is None:
            return '<PROTO (%s)>' % repr(self.args)
        else:
            return '<PROTO (%r) -> %r>' % (self.args, self.retv)

    @property
    def argspace(self) -> int:
        return sum(
            [v.size for v in self.args],
            (0 if self.retv is None else self.retv.size)
        )
        
    @property
    def inputspace(self) -> int:
        return sum([v.size for v in self.args])

class PrototypeMap(Dict[str, Prototype]):
    @staticmethod
    def _dv(c: str) -> int:
        if c == '(':
            return 1
        elif c == ')':
            return -1
        else:
            return 0

    @staticmethod
    def _tk(s: str, p: str) -> bool:
        return s.startswith(p) and (s == p or s[len(p)].isspace())

    @classmethod
    def _punc(cls, s: str) -> bool:
        return s in cls.__puncs_

    @staticmethod
    def _err(msg: str) -> SyntaxError:
        return SyntaxError(
            msg + ', ' +
            'the parser integrated in this tool is just a text-based parser, ' +
            'so please keep the companion .go file as simple as possible and do not use defined types'
        )

    @staticmethod
    def _align(nb: int) -> int:
        return (((nb - 1) >> 3) + 1) << 3

    @classmethod
    def _retv(cls, ret: str) -> Tuple[str, int, Register, Register]:
        name, size, xmm = cls._args(ret)
        reg = Register('d0') if xmm else Register('x0')
        return name, size, reg, reg

    @classmethod
    def _args(cls, arg: str, sv: str = '') -> Tuple[str, int, bool]:
        while True:
            if not arg:
                raise SyntaxError('missing type for parameter: ' + sv)
            elif arg[0] != '_' and not arg[0].isalnum():
                return (sv,) + cls._size(arg.strip())
            elif not sv and arg[0].isdigit():
                raise SyntaxError('invalid character: ' + repr(arg[0]))
            else:
                sv += arg[0]
                arg = arg[1:]

    @classmethod
    def _size(cls, name: str) -> Tuple[int, bool]:
        if name[0] == '*':
            return cls._align(8), False
        elif name in ('int8', 'uint8', 'byte', 'bool'):
            return cls._align(1), False
        elif name in ('int16', 'uint16'):
            return cls._align(2), False
        elif name == 'float32':
            return cls._align(4), True
        elif name in ('int32', 'uint32', 'rune'):
            return cls._align(4), False
        elif name == 'float64':
            return cls._align(8), True
        elif name in ('int64', 'uint64', 'uintptr', 'int', 'Pointer', 'unsafe.Pointer'):
            return cls._align(8), False
        else:
            raise cls._err('unrecognized type "%s"' % name)

    @classmethod
    def _func(cls, src: List[str], idx: int, depth: int = 0) -> Tuple[str, int]:
        for i in range(idx, len(src)):
            for x in map(cls._dv, src[i]):
                if depth + x >= 0:
                    depth += x
                else:
                    raise cls._err('encountered ")" more than "(" on line %d' % (i + 1))
            else:
                if depth == 0:
                    return ' '.join(src[idx:i + 1]), i + 1
        else:
            raise cls._err('unexpected EOF when parsing function signatures')

    @classmethod
    def parse(cls, src: str) -> Tuple[str, 'PrototypeMap']:
        idx = 0
        pkg = ''
        ret = PrototypeMap()
        buf = src.splitlines()

        # scan through all the lines
        while idx < len(buf):
            line = buf[idx]
            line = line.strip()

            # skip empty lines
            if not line:
                idx += 1
                continue

            # check for package name
            if cls._tk(line, 'package'):
                idx, pkg = idx + 1, line[7:].strip().split()[0]
                continue

            if OUTPUT_RAW:

                # extract funcname like "[var ]{funcname} = func(..."
                end = line.find('func(')
                if end == -1:
                    idx += 1
                    continue
                name = line[:end].strip()
                if name.startswith('var '):
                    name = name[4:].strip()

                # function names must be identifiers
                if not name.isidentifier():
                    raise cls._err('invalid function prototype: ' + name)

                # register a empty prototype
                ret[name] = Prototype(None, [])
                idx += 1

            else:

                # only cares about those functions that does not have bodies
                if line[-1] == '{' or not cls._tk(line, 'func'):
                    idx += 1
                    continue

                # prevent type-aliasing primitive types into other names
                if cls._tk(line, 'type'):
                    raise cls._err('please do not declare any type with in the companion .go file')

                # find the next function declaration
                decl, pos = cls._func(buf, idx)
                func, idx = decl[4:].strip(), pos

                # find the beginning '('
                nd = 1
                pos = func.find('(')

                # must have a '('
                if pos == -1:
                    raise cls._err('invalid function prototype: ' + decl)

                # extract the name and signature
                args = ''
                name = func[:pos].strip()
                func = func[pos + 1:].strip()

                # skip the method declaration
                if not name:
                    continue

                # function names must be identifiers
                if not name.isidentifier():
                    raise cls._err('invalid function prototype: ' + decl)

                # extract the argument list
                while nd and func:
                    nch  = func[0]
                    func = func[1:]

                    # adjust the nesting level
                    nd   += cls._dv(nch)
                    args += nch

                # check for EOF
                if not nd:
                    func = func.strip()
                else:
                    raise cls._err('unexpected EOF when parsing function prototype: ' + decl)

                # check for multiple returns
                if ',' in func:
                    raise cls._err('can only return a single value (detected by looking for "," within the return list)')

                # check for return signature
                if not func:
                    retv = None
                elif func[0] == '(' and func[-1] == ')':
                    retv = Parameter(*cls._retv(func[1:-1]))
                else:
                    raise SyntaxError('badly formatted return argument (please use parenthesis and proper arguments naming): ' + func)

                # extract the argument list
                if not args[:-1]:
                    args, alens, axmm = [], [], []
                else:
                    args, alens, axmm = list(zip(*[cls._args(v.strip()) for v in args[:-1].split(',')]))

                # check for the result
                cregs = []
                goregs = []
                idxs = [0, 0]

                # split the integer & floating point registers
                for xmm in axmm:
                    key = 0 if xmm else 1
                    seq = FPARGS_ORDER if xmm else ARGS_ORDER_C
                    goseq = FPARGS_ORDER if xmm else ARGS_ORDER_GO

                    # check the argument count
                    if idxs[key] >= len(seq):
                        raise cls._err("too many arguments, consider pack some into a pointer")

                    # add the register
                    cregs.append(seq[idxs[key]])
                    goregs.append(goseq[idxs[key]])
                    idxs[key] += 1

                # register the prototype
                ret[name] = Prototype(retv, [
                    Parameter(arg, size, creg, goreg)
                    for arg, size, creg, goreg in zip(args, alens, cregs, goregs)
                ])

        # all done
        return pkg, ret

### Assembly Source Parser ###

ESC_IDLE = 0    # escape parser is idleing
ESC_ISTR = 1    # currently inside a string
ESC_BKSL = 2    # encountered backslash, prepare for escape sequences
ESC_HEX0 = 3    # expect the first hexadecimal character of a "\x" escape
ESC_HEX1 = 4    # expect the second hexadecimal character of a "\x" escape
ESC_OCT1 = 5    # expect the second octal character of a "\000" escape
ESC_OCT2 = 6    # expect the third octal character of a "\000" escape

class Command:
    cmd  : str
    args : List[Union[str, bytes]]

    def __init__(self, cmd: str, args: List[Union[str, bytes]]):
        self.cmd  = cmd
        self.args = args

    def __repr__(self):
        return '<CMD %s %s>' % (self.cmd, ', '.join(map(repr, self.args)))

    @classmethod
    def parse(cls, src: str) -> 'Command':
        val = src.split(None, 1)
        cmd = val[0]

        # no parameters
        if len(val) == 1:
            return cls(cmd, [])

        # extract the argument string
        idx = 0
        esc = 0
        pos = None
        args = []
        vstr = val[1]

        # scan through the whole string
        while idx < len(vstr):
            nch = vstr[idx]
            idx += 1

            # mark the start of the argument
            if pos is None:
                pos = idx - 1

            # encountered the delimiter outside of a string
            if nch == ',' and esc == ESC_IDLE:
                pos, p = None, pos
                args.append(vstr[p:idx - 1].strip())

            # start of a string
            elif nch == '"' and esc == ESC_IDLE:
                esc = ESC_ISTR

            # end of string
            elif nch == '"' and esc == ESC_ISTR:
                esc = ESC_IDLE
                pos, p = None, pos
                args.append(vstr[p:idx].strip()[1:-1].encode('utf-8').decode('unicode_escape'))

            # escape characters
            elif nch == '\\' and esc == ESC_ISTR:
                esc = ESC_BKSL

            # hexadecimal escape characters (3 chars)
            elif esc == ESC_BKSL and nch == 'x':
                esc = ESC_HEX0

            # octal escape characters (3 chars)
            elif esc == ESC_BKSL and nch in string.octdigits:
                esc = ESC_OCT1

            # generic escape characters (single char)
            elif esc == ESC_BKSL and nch in ('a', 'b', 'f', 'r', 'n', 't', 'v', '"', '\\'):
                esc = ESC_ISTR

            # invalid escape sequence
            elif esc == ESC_BKSL:
                raise SyntaxError('invalid escape character: ' + repr(nch))

            # normal characters, simply advance to the next character
            elif esc in (ESC_IDLE, ESC_ISTR):
                pass

            # hexadecimal escape characters
            elif esc in (ESC_HEX0, ESC_HEX1) and nch.lower() in string.hexdigits:
                esc = ESC_HEX1 if esc == ESC_HEX0 else ESC_ISTR

            # invalid hexadecimal character
            elif esc in (ESC_HEX0, ESC_HEX1):
                raise SyntaxError('invalid hexdecimal character: ' + repr(nch))

            # octal escape characters
            elif esc in (ESC_OCT1, ESC_OCT2) and nch.lower() in string.octdigits:
                esc = ESC_OCT2 if esc == ESC_OCT1 else ESC_ISTR

            # at most 3 octal digits
            elif esc in (ESC_OCT1, ESC_OCT2):
                esc = ESC_ISTR

            # illegal state, should not happen
            else:
                raise RuntimeError('illegal state: %d' % esc)

        # check for the last argument
        if pos is None:
            return cls(cmd, args)

        # add the last argument and build the command
        args.append(vstr[pos:].strip())
        return cls(cmd, args)

class Expression:
    pos: int
    src: str

    def __init__(self, src: str):
        self.pos = 0
        self.src = src

    @property
    def _ch(self) -> str:
        return self.src[self.pos]

    @property
    def _eof(self) -> bool:
        return self.pos >= len(self.src)

    def _rch(self) -> str:
        pos, self.pos = self.pos, self.pos + 1
        return self.src[pos]

    def _hex(self, ch: str) -> bool:
        if len(ch) == 1 and ch[0] == '0':
            return self._ch.lower() == 'x'
        elif len(ch) <= 1 or ch[1].lower() != 'x':
            return self._ch.isdigit()
        else:
            return self._ch in string.hexdigits

    def _int(self, ch: str) -> Token:
        while not self._eof and self._hex(ch):
            ch += self._rch()
        else:
            if ch.lower().startswith('0x'):
                return Token.num(int(ch, 16))
            elif ch[0] == '0':
                return Token.num(int(ch, 8))
            else:
                return Token.num(int(ch))

    def _name(self, ch: str) -> Token:
        while not self._eof and (self._ch == '_' or self._ch.isalnum()):
            ch += self._rch()
        else:
            return Token.name(ch)

    def _read(self, ch: str) -> Token:
        if ch.isdigit():
            return self._int(ch)
        elif ch.isidentifier():
            return self._name(ch)
        elif ch in ('*', '<', '>') and not self._eof and self._ch == ch:
            return Token.punc(self._rch() * 2)
        elif ch in ('+', '-', '*', '/', '%', '&', '|', '^', '~', '(', ')'):
            return Token.punc(ch)
        else:
            raise SyntaxError('invalid character: ' + repr(ch))

    def _peek(self) -> Optional[Token]:
        pos = self.pos
        ret = self._next()
        self.pos = pos
        return ret

    def _next(self) -> Optional[Token]:
        while not self._eof and self._ch.isspace():
            self.pos += 1
        else:
            return Token.end() if self._eof else self._read(self._rch())

    def _grab(self, tk: Token, getvalue: Callable[[str], int]) -> int:
        if tk.tag == TOKEN_NUM:
            return tk.val
        elif tk.tag == TOKEN_NAME:
            return getvalue(tk.val)
        else:
            raise SyntaxError('integer or identifier expected, got ' + repr(tk))

    __pred__ = [
        {'<<', '>>'},
        {'|'},
        {'^'},
        {'&'},
        {'+', '-'},
        {'*', '/', '%'},
        {'**'},
    ]

    __binary__ = {
        '+'  : lambda a, b: a + b,
        '-'  : lambda a, b: a - b,
        '*'  : lambda a, b: a * b,
        '/'  : lambda a, b: a / b,
        '%'  : lambda a, b: a % b,
        '&'  : lambda a, b: a & b,
        '^'  : lambda a, b: a ^ b,
        '|'  : lambda a, b: a | b,
        '<<' : lambda a, b: a << b,
        '>>' : lambda a, b: a >> b,
        '**' : lambda a, b: a ** b,
    }

    def _eval(self, op: str, v1: int, v2: int) -> int:
        return self.__binary__[op](v1, v2)

    def _nest(self, nest: int, getvalue: Callable[[str], int]) -> int:
        ret = self._expr(0, nest + 1, getvalue)
        ntk = self._next()

        # it must follows with a ')' operator
        if ntk.tag != TOKEN_PUNC or ntk.val != ')':
            raise SyntaxError('")" expected, got ' + repr(ntk))
        else:
            return ret

    def _unit(self, nest: int, getvalue: Callable[[str], int]) -> int:
        tk = self._next()
        tt, tv = tk.tag, tk.val

        # check for unary operators
        if tt == TOKEN_NUM:
            return tv
        elif tt == TOKEN_NAME:
            return getvalue(tv)
        elif tt == TOKEN_PUNC and tv == '(':
            return self._nest(nest, getvalue)
        elif tt == TOKEN_PUNC and tv == '+':
            return self._unit(nest, getvalue)
        elif tt == TOKEN_PUNC and tv == '-':
            return -self._unit(nest, getvalue)
        elif tt == TOKEN_PUNC and tv == '~':
            return ~self._unit(nest, getvalue)
        else:
            raise SyntaxError('integer, unary operator or nested expression expected, got ' + repr(tk))

    def _term(self, pred: int, nest: int, getvalue: Callable[[str], int]) -> int:
        lv = self._expr(pred + 1, nest, getvalue)
        tk = self._peek()

        # scan to the end
        while True:
            tt = tk.tag
            tv = tk.val

            # encountered EOF
            if tt == TOKEN_END:
                return lv

            # must be an operator here
            if tt != TOKEN_PUNC:
                raise SyntaxError('operator expected, got ' + repr(tk))

            # check for the operator precedence
            if tv not in self.__pred__[pred]:
                return lv

            # apply the operator
            op = self._next().val
            rv = self._expr(pred + 1, nest, getvalue)
            lv = self._eval(op, lv, rv)
            tk = self._peek()

    def _expr(self, pred: int, nest: int, getvalue: Callable[[str], int]) -> int:
        if pred >= len(self.__pred__):
            return self._unit(nest, getvalue)
        else:
            return self._term(pred, nest, getvalue)

    def eval(self, getvalue: Callable[[str], int]) -> int:
        return self._expr(0, 0, getvalue)


class Instr:
    ALIGN_WIDTH = 48
    len   : int                     = NotImplemented
    instr : Union[str, Instruction] = NotImplemented

    def size(self, _: int) -> int:
        return self.len

    def formatted(self, pc: int) -> str:
        raise NotImplementedError

    @staticmethod
    def raw_formatted(bs: bytes, comm: str, pc: int) -> str:
        t = '\t'
        if bs:
            for b in bs:
                t +='0x%02x, ' % b
            # if len(bs)<Instr.ALIGN_WIDTH:
            #     t += '\b' * (Instr.ALIGN_WIDTH - len(bs))
        return '%s//%s%s' % (t, ('0x%08x ' % pc) if pc else ' ', comm)

class RawInstr(Instr):
    bs: bytes
    def __init__(self, size: int, instr: str, bs: bytes):
        self.len = size
        self.instr = instr
        self.bs = bs

    def formatted(self, _: int) -> str:
        return '\t' + self.instr

    def raw_formatted(self, pc: int) -> str:
        return Instr.raw_formatted(self.bs, self.instr, pc)

class IntInstr(Instr):
    comm: str
    func: Callable[[], int]

    def __init__(self, size: int, func: Callable[[], int], comments: str = ''):
        self.len = size
        self.func = func
        self.comm = comments

    @property
    def raw_bytes(self):
        return self.func().to_bytes(self.len, 'little')

    @property
    def instr(self) -> str:
        return Instruction.encode(self.func().to_bytes(self.len, 'little'), self.comm)

    def formatted(self, _: int) -> str:
        return '\t' + self.instr

    def raw_formatted(self, pc: int) -> str:
        return Instr.raw_formatted(self.func().to_bytes(self.len, 'little'), self.comm, pc)

class X86Instr(Instr):
    def __init__(self, instr: Instruction):
        self.len = instr.size
        self.instr = instr

    def resize(self, size: int) -> int:
        self.len = size
        return size

    def formatted(self, _: int) -> str:
        return '\t' + self.instr.encoded

    def raw_formatted(self, pc: int) -> str:
        return Instr.raw_formatted(self.instr._raw_instr(), str(self.instr), pc)

class LabelInstr(Instr):
    def __init__(self, name: str):
        self.len = 0
        self.instr = name

    def formatted(self, _: int) -> str:
        if self.instr.isidentifier():
            return self.instr + ':'
        else:
            return '_LB_%08x: // %s' % (hash(self.instr) & 0xffffffff, self.instr)

    def raw_formatted(self, pc: int) -> str:
        return Instr.raw_formatted(None, str(self.instr), pc)

class BranchInstr(Instr):
    def __init__(self, instr: Instruction):
        self.len = instr.size
        self.instr = instr

    def formatted(self, _: int) -> str:
        return '\t' + self.instr.encoded

    def raw_formatted(self, pc: int) -> str:
        return Instr.raw_formatted(self.instr._raw_instr(), str(self.instr), pc)

class CommentInstr(Instr):
    def __init__(self, text: str):
        self.len = 0
        self.instr = '// ' + text

    def formatted(self, _: int) -> str:
        return '\t' + self.instr

    def raw_formatted(self, pc: int) -> str:
        return  Instr.raw_formatted(None, str(self.instr), None)

class AlignmentInstr(Instr):
    bits: int
    fill: int

    def __init__(self, bits: int, fill: int = 0):
        self.bits = bits
        self.fill = fill

    def size(self, pc: int) -> int:
        mask = (1 << self.bits) - 1
        return (mask - (pc & mask) + 1) & mask

    def formatted(self, pc: int) -> str:
        buf = bytes([self.fill]) * self.size(pc)
        return '\t' + Instruction.encode(buf, '.p2align %d, 0x%02x' % (self.bits, self.fill))

    def raw_formatted(self, pc: int) -> str:
        buf = bytes([self.fill]) * self.size(pc)
        return Instr.raw_formatted(buf, '.p2align %d, 0x%02x' % (self.bits, self.fill), pc)

REG_MAP = {
    'x0'  : ('MOVD'  , 'R0'),
    'x1'  : ('MOVD'  , 'R1'),
    'x2'  : ('MOVD'  , 'R2'),
    'x3'  : ('MOVD'  , 'R3'),
    'x4'  : ('MOVD'  , 'R4'),
    'x5'  : ('MOVD'  , 'R5'),
    'x6'  : ('MOVD'  , 'R6'),
    'x7'  : ('MOVD'  , 'R7'),
    'd0'  : ('FMOVD' , 'F0'),
    'd1'  : ('FMOVD' , 'F1'),
    'd2'  : ('FMOVD' , 'F2'),
    'd3'  : ('FMOVD' , 'F3'),
    'd4'  : ('FMOVD' , 'F4'),
    'd5'  : ('FMOVD' , 'F5'),
    'd6'  : ('FMOVD' , 'F6'),
    'd7'  : ('FMOVD' , 'F7'),
}

class Counter:
    value: int = 0

    @classmethod
    def next(cls) -> int:
        val, cls.value = cls.value, cls.value + 1
        return val

class BasicBlock:
    maxsp: int
    name: str
    weak: bool
    jmptab: bool
    func: bool
    body: List[Instr]
    prevs: List['BasicBlock']
    next: Optional['BasicBlock']
    jump: Optional['BasicBlock']

    def __init__(self, name: str, weak: bool = True, jmptab: bool = False, func: bool = False):
        self.maxsp = -1
        self.body = []
        self.prevs = []
        self.name = name
        self.weak = weak
        self.next = None
        self.jump = None
        self.jmptab = jmptab
        self.func = func

    def __repr__(self):
        return '{BasicBlock %s}' % repr(self.name)

    @property
    def last(self) -> Optional[Instr]:
        return next((v for v in reversed(self.body) if not isinstance(v, CommentInstr)), None)

    def if_all_IntInstr_then_2_RawInstr(self):
        is_table = False
        instr_size = 0
        for instr in self.body:
            if isinstance(instr, IntInstr):
               if not is_table:
                   instr_size = instr.len
               is_table = True
               if instr_size != instr.len:
                   instr_size = 0
               continue
            if isinstance(instr, AlignmentInstr):
               continue
            if isinstance(instr, LabelInstr):
               continue
            # others
            return

        if not is_table:
            return

        # .long or .quad
        if instr_size == 8 or instr_size == 4:
            return

        # All instrs are IntInstr, golang asm only suuport WORD and DWORD for arm. We need
        # combine them as 4-bytes RawInstr and align block
        nb = [] # new body
        raw_buf = [];
        comment = ''

        # first element is LabelInstr
        for i in range(1, len(self.body)):
            if isinstance(self.body[i], AlignmentInstr):
                if i != len(self.body) -1:
                    raise RuntimeError("Not support p2algin in : %s" % self.name)
                continue

            raw_buf += self.body[i].raw_bytes
            comment += '// ' + self.body[i].comm + '\n'

        align_size = len(raw_buf) % 4
        if align_size != 0:
            raw_buf += int(0).to_bytes(4 - align_size, 'little')

        if isinstance(self.body[0], LabelInstr):
            nb.append(self.body[0])

        for i in range(0, len(raw_buf), 4):
            buf = raw_buf[i: i + 4]
            nb.append(RawInstr(len(buf), Instruction.encode(buf), buf))

        nb.append(CommentInstr(comment))

        if isinstance(self.body[-1:-1], AlignmentInstr):
            nb.append(self.body[-1:-1])
        self.body = nb

    def size_of(self, pc: int) -> int:
        return functools.reduce(lambda p, v: p + v.size(pc + p), self.body, 0)

    def link_to(self, block: 'BasicBlock'):
        self.next = block
        block.prevs.append(self)

    def jump_to(self, block: 'BasicBlock'):
        self.jump = block
        block.prevs.append(self)

    @classmethod
    def annonymous(cls) -> 'BasicBlock':
        return cls('// bb.%d' % Counter.next(), weak = False)

CLANG_JUMPTABLE_LABLE = 'LJTI'

class CodeSection:
    dead   : bool
    export : bool
    blocks : List[BasicBlock]
    labels : Dict[str, BasicBlock]
    jmptabs: Dict[str, List[BasicBlock]]
    funcs  : Dict[str, Pcsp]
    bsmap_ : Dict[str, int]

    def __init__(self):
        self.dead   = False
        self.labels = {}
        self.export = False
        self.blocks = [BasicBlock.annonymous()]
        self.jmptabs = {}
        self.funcs = {}
        self.bsmap_ = {}

    @classmethod
    def _dfs_jump_first(cls, bb: BasicBlock, visited: Dict[BasicBlock, bool], hook: Callable[[BasicBlock], bool]) -> bool:
        if bb not in visited or not visited[bb]:
            visited[bb] = True
            if bb.jump and not cls._dfs_jump_first(bb.jump, visited, hook):
                return False
            if bb.next and not cls._dfs_jump_first(bb.next, visited, hook):
                return False
            return hook(bb)
        else:
            return True

    def get_jmptab(self, name: str) -> List[BasicBlock]:
        return self.jmptabs.setdefault(name, [])

    def get_block(self, name: str) -> BasicBlock:
        for block in self.blocks:
            if block.name == name:
                return block

    @property
    def block(self) -> BasicBlock:
        return self.blocks[-1]

    @property
    def instrs(self) -> Iterable[Instr]:
        for block in self.blocks:
            yield from block.body

    def _make(self, name: str, jmptab: bool = False, func: bool = False):
        if func:
        #NOTICE: if it is a function, always set func to be True
            if (old := self.labels.get(name)) and (old.func != func):
                old.func = True
        return self.labels.setdefault(name, BasicBlock(name, jmptab = jmptab, func = func))

    def _next(self, link: BasicBlock):
        if self.dead:
            self.dead = False
        else:
            self.block.link_to(link)

    def _decl(self, name: str, block: BasicBlock):
        block.weak = False
        block.body.append(LabelInstr(name))
        self._next(block)
        self.blocks.append(block)

    def _kill(self, name: str):
        self.dead = True
        self.block.link_to(self._make(name))

    def _split(self, jmp: BasicBlock):
        self.jump = True
        link = BasicBlock.annonymous()
        self.labels[link.name] = link
        self.block.link_to(link)
        self.block.jump_to(jmp)
        self.blocks.append(link)

    @staticmethod
    def _mk_align(v: int) -> int:
        if v & 15 == 0:
            return v
        else:
            print('* warning: SP is not aligned with 16 bytes.', file = sys.stderr)
            return (v + 15) & -16

    @staticmethod
    def _is_spadj(ins: Instruction) -> bool:
        return len(ins.instr.operands) == 3                         and \
               isinstance(ins.instr.operands[1], mcasm.mc.Register) and \
               isinstance(ins.instr.operands[2], int)               and \
               ins.instr.operands[1].name == 'RSP'

    @staticmethod
    def _is_spmove(ins: Instruction, i: int) -> bool:
        return len(ins.operands) == 2                and \
               isinstance(ins.operands[0], Register) and \
               isinstance(ins.operands[1], Register) and \
               ins.operands[i].reg == 'rsp'

    @staticmethod
    def _is_rjump(ins: Optional[Instr]) -> bool:
        return isinstance(ins, X86Instr) and ins.instr.is_branch_label

    def _find_label(self, name: str, adjs: Iterable[int], size: int = 0) -> int:
        for adj, block in zip(adjs, self.blocks):
            if block.name == name:
                return size
            else:
                # find block size from cache
                v = self.bsmap_.get(block.name)
                if v is not None:
                    size += v + adj
                else:
                    block_size = block.size_of(size)
                    size += block_size + adj
                    self.bsmap_[block.name] = block_size
        else:
            raise SyntaxError('unresolved reference to name: ' + name)

    def _alloc_instr(self, instr: Instruction):
        if not instr.is_branch_label:
            self.block.body.append(X86Instr(instr))
        else:
            self.block.body.append(BranchInstr(instr))

    # it seems to not be able to specify stack aligment inside the Go ASM so we
    # need to replace the aligned instructions with unaligned one if either of it's
    # operand is an RBP relative addressing memory operand

    def _check_align(self, instr: Instruction) -> bool:
        # TODO: check
        return False

    def _check_split(self, instr: Instruction):
        if instr.is_return:
            self.dead = True

        elif instr.is_jmpq: # jmpq
            # backtrace jump table from current block (BFS)
            prevs = [self.block]
            visited = set()
            while len(prevs) > 0:
                curb = prevs.pop()
                if curb in visited:
                    continue
                else:
                    visited.add(curb)

                # backtrace instructions
                for ins in reversed(curb.body):
                    if isinstance(ins, X86Instr) and ins.instr.jmptab:
                        self._split(self._make(ins.instr.jmptab, jmptab = True))
                        return

                if curb.prevs:
                    prevs.extend(curb.prevs)

        elif instr.is_branch_label:
            if instr.is_jmp:
                self._kill(instr.label_name)
            
            elif instr.is_invoke: # call
                fname = instr.label_name
                self._split(self._make(fname, func = True))

            else: # jeq, ja, jae ...
                self._split(self._make(instr.label_name))

    def _trace_block(self, bb: BasicBlock, pcsp: Optional[Pcsp]) -> int:
        if (pcsp is not None):
            if bb.name in self.funcs:
                # already traced
                pcsp = None
            else:
                # continue tracing, update the pcsp
                # NOTICE: must mark pcsp at block entry because go only calculate delta value
                pcsp.pc = self.get(bb.name)
                if bb.func or pcsp.pc < pcsp.entry:
                    # new func
                    pcsp = Pcsp(pcsp.pc)
                    self.funcs[bb.name] = pcsp

        if bb.maxsp == -1:
            ret = self._trace_nocache(bb, pcsp)
            return ret
        elif bb.maxsp >= 0:
            return bb.maxsp
        else:
            return 0

    def _trace_nocache(self, bb: BasicBlock, pcsp: Optional[Pcsp]) -> int:
        bb.maxsp = -2

        # ## FIXME:
        # if pcsp is None:
        #     pcsp = Pcsp(0)

        # make a fake object just for reducing redundant checking
        if pcsp:
            pc0, sp0 = pcsp.pc, pcsp.sp

        maxsp, term = self._trace_instructions(bb, pcsp)

        # this is a terminating block
        if term:
            return maxsp

        # don't trace it's next block if it's an unconditional jump
        a, b = 0, 0
        if pcsp:
            pc, sp = pcsp.pc, pcsp.sp

        if bb.jump:
            if bb.jump.jmptab:
                cases = self.get_jmptab(bb.jump.name)
                for case in cases:
                    nsp = self._trace_block(case, pcsp)
                    if pcsp:
                        pcsp.pc, pcsp.sp = pc, sp
                    if nsp > a:
                        a = nsp
            else:
                a = self._trace_block(bb.jump, pcsp)
                if pcsp:
                    pcsp.pc, pcsp.sp = pc, sp

        if bb.next:
            b = self._trace_block(bb.next, pcsp)

        if pcsp:
            pcsp.pc, pcsp.sp = pc0, sp0

        # select the maximum stack depth
        bb.maxsp = maxsp + max(a, b)
        return bb.maxsp

    def _trace_instructions(self, bb: BasicBlock, pcsp: Pcsp) -> Tuple[int, bool]:
        cursp = 0
        maxsp = 0
        close = False

        # scan every instruction
        for ins in bb.body:
            diff = 0

            if isinstance(ins, X86Instr):
                name = ins.instr.mnemonic
                operands = ins.instr.instr.operands

                # check for instructions
                if name == 'ret':
                    close = True
                elif isinstance(operands[0], mcasm.mc.Register) and operands[0].name == 'SP':
                    # print(ins.instr.asm_code)
                    if name == 'add':
                        diff = -self._mk_align(operands[2])
                    elif name == 'sub':
                        diff = self._mk_align(operands[2])
                    elif name == 'stp':
                        diff = -self._mk_align(operands[4] * 8)
                    elif name == 'ldp':
                        diff = -self._mk_align(operands[4] * 8)
                    elif name == 'str':
                        diff = -self._mk_align(operands[3])
                    else:
                        raise RuntimeError("An instruction adjsut sp but bot processed: %s" % ins.instr.asm_code)

                cursp += diff

                # update the max stack depth
                if cursp > maxsp:
                    maxsp = cursp

            # update pcsp
            if pcsp:
                pcsp.update(ins.size(pcsp.pc), diff)

        # trace successful
        return maxsp, close

    def get(self, key: str) -> Optional[int]:
        if key not in self.labels:
            raise SyntaxError('unresolved reference to name: %s' % key)
        else:
            return self._find_label(key, itertools.repeat(0, len(self.blocks)))

    def has(self, key: str) -> bool:
        return key in self.labels

    def emit(self, buf: bytes, comments: str = ''):
        if not self.dead:
            self.block.body.append(RawInstr(len(buf), Instruction.encode(buf, comments or buf.hex()), buf))

    def lazy(self, size: int, func: Callable[[], int], comments: str = ''):
        if not self.dead:
            self.block.body.append(IntInstr(size, func, comments))

    def label(self, name: str):
        if name not in self.labels or self.labels[name].weak:
            self._decl(name, self._make(name))
        else:
            raise SyntaxError('duplicated label: ' + name)

    def instr(self, instr: Instruction):
        if not self.dead:
            if self._check_align(instr):
                return
            self._alloc_instr(instr)
            self._check_split(instr)

    # @functools.cache
    def stacksize(self, name: str) -> int:
        if name not in self.labels:
            raise SyntaxError('undefined function: ' + name)
        else:
            return self._trace_block(self.labels[name], None)

    # @functools.cache
    def pcsp(self, name: str, entry: int) -> int:
        if name not in self.labels:
            raise SyntaxError('undefined function: ' + name)
        else:
            pcsp = Pcsp(entry)
            self.labels[name].func = True
            return self._trace_block(self.labels[name], pcsp)

    def debug(self, pos: int, inss: List[Instruction]):
        def inject(bb: BasicBlock) -> bool:
            if (not bb.func) and (bb.name not in self.funcs):
                return True
            nonlocal pos
            if pos >= len(bb.body):
                return
            for ins in inss:
                bb.body.insert(pos, ins)
                pos += 1
        visited = {}
        for _, bb in self.labels.items():
            CodeSection._dfs_jump_first(bb, visited, inject)
    def debug(self):
        for label, bb in self.labels.items():
            print(label)
            for v in bb.body:
                if isinstance(v, (X86Instr, BranchInstr)):
                    print(v.instr.asm_code)

STUB_NAME = '__native_entry__'
STUB_SIZE = 67
WITH_OFFS = os.getenv('ASM2ASM_DEBUG_OFFSET', '').lower() in ('1', 'yes', 'true')

class Assembler:
    out  : List[str]
    subr : Dict[str, int]
    code : CodeSection
    vals : Dict[str, Union[str, int]]

    def __init__(self):
        self.out  = []
        self.subr = {}
        self.vals = {}
        self.code = CodeSection()

    def _get(self, v: str) -> int:
        if v not in self.vals:
            return self.code.get(v)
        elif isinstance(self.vals[v], int):
            return self.vals[v]
        else:
            ret = self.vals[v] = self._eval(self.vals[v])
            return ret

    def _eval(self, v: str) -> int:
        return Expression(v).eval(self._get)

    def _emit(self, v: bytes, cmd: str):
        align_size = len(v) % 4
        if align_size != 0:
            v += int(0).to_bytes(4 - align_size, 'little')

        for i in range(0, len(v), 4):
            self.code.emit(v[i:i + 4], '%s %d, %s' % (cmd, len(v[i:i + 4]), repr(v[i:i + 16])[1:]))

    def _limit(self, v: int, a: int, b: int) -> int:
        if not (a <= v <= b):
            raise SyntaxError('integer constant out of bound [%d, %d): %d' % (a, b, v))
        else:
            return v

    def _vfill(self, cmd: str, args: List[str]) -> Tuple[int, int]:
        if len(args) == 1:
            return self._limit(self._eval(args[0]), 1, 1 << 64), 0
        elif len(args) == 2:
            return self._limit(self._eval(args[0]), 1, 1 << 64), self._limit(self._eval(args[1]), 0, 255)
        else:
            raise SyntaxError(cmd + ' takes 1 ~ 2 arguments')

    def _bytes(self, cmd: str, args: List[str], low: int, high: int, size: int):
        if len(args) != 1:
            raise SyntaxError(cmd + ' takes exact 1 argument')
        else:
            self.code.lazy(size, lambda: self._limit(self._eval(args[0]), low, high) & high, '%s %s' % (cmd, args[0]))

    def _comment(self, msg: str):
        self.code.blocks[-1].body.append(CommentInstr(msg))

    def _cmd_nop(self, _: List[str]):
        pass

    def _cmd_set(self, args: List[str]):
        if len(args) != 2:
            raise SyntaxError(".set takes exact 2 argument")
        elif not args[0].isidentifier():
            raise SyntaxError(repr(args[0]) + " is not a valid identifier")
        else:
            key = args[0]
            val = args[1]
            self.vals[key] = val
            self._comment('.set ' + ', '.join(args))
            # special case: clang-generated jump tables are always like '{block}_{table}'
            jt = val.find(CLANG_JUMPTABLE_LABLE)
            if jt > 0:
                tab = self.code.get_jmptab(val[jt:])
                tab.append(self.code.get_block(val[:jt-1]))

    def _cmd_byte(self, args: List[str]):
        self._bytes('.byte', args, -0x80, 0xff, 1)

    def _cmd_word(self, args: List[str]):
        self._bytes('.word', args, -0x8000, 0xffff, 2)

    def _cmd_long(self, args: List[str]):
        self._bytes('.long', args, -0x80000000, 0xffffffff, 4)

    def _cmd_quad(self, args: List[str]):
        self._bytes('.quad', args, -0x8000000000000000, 0xffffffffffffffff, 8)

    def _cmd_ascii(self, args: List[str]):
        if len(args) != 1:
            raise SyntaxError('.ascii takes exact 1 argument')
        else:
            self._emit(args[0].encode('latin-1'), '.ascii')

    def _cmd_asciz(self, args: List[str]):
        if len(args) != 1:
            raise SyntaxError('.asciz takes exact 1 argument')
        else:
            self._emit(args[0].encode('latin-1') + b'\0', '.asciz')

    def _cmd_space(self, args: List[str]):
        nb, fv = self._vfill('.space', args)
        self._emit(bytes([fv] * nb), '.space')

    def _cmd_p2align(self, args: List[str]):
        if len(args) == 1:
            self.code.block.body.append(AlignmentInstr(self._eval(args[0])))
        elif len(args) == 2:
            self.code.block.body.append(AlignmentInstr(self._eval(args[0]), self._eval(args[1])))
        else:
            raise SyntaxError('.p2align takes 1 ~ 2 arguments')

    @functools.cached_property
    def _commands(self) -> dict:
        return {
            '.set'                     : self._cmd_set,
            '.int'                     : self._cmd_long,
            '.long'                    : self._cmd_long,
            '.byte'                    : self._cmd_byte,
            '.quad'                    : self._cmd_quad,
            '.word'                    : self._cmd_word,
            '.hword'                   : self._cmd_word,
            '.short'                   : self._cmd_word,
            '.ascii'                   : self._cmd_ascii,
            '.asciz'                   : self._cmd_asciz,
            '.space'                   : self._cmd_space,
            '.globl'                   : self._cmd_nop,
            '.text'                    : self._cmd_nop,
            '.file'                    : self._cmd_nop,
            '.type'                    : self._cmd_nop,
            '.p2align'                 : self._cmd_p2align,
            '.align'                   : self._cmd_nop,
            '.size'                    : self._cmd_nop,
            '.section'                 : self._cmd_nop,
            '.loh'                     : self._cmd_nop,
            '.data_region'             : self._cmd_nop,
            '.build_version'           : self._cmd_nop,
            '.end_data_region'         : self._cmd_nop,
            '.subsections_via_symbols' : self._cmd_nop,
            # linux-gnu
            '.xword'                   :self._cmd_nop,
        }

    @staticmethod
    def _is_rip_relative(op: Operand) -> bool:
        return isinstance(op, Memory) and \
               op.base is not None    and \
               op.base.reg == 'rip'   and \
               op.index is None       and \
               isinstance(op.disp, Reference)

    @staticmethod
    def _remove_comments(line: str, *, st: str = 'normal') -> str:
        for i, ch in enumerate(line):
            if   st == 'normal' and ch == '/'        : st = 'slcomm'
            elif st == 'normal' and ch == '\"'       : st = 'string'
            # elif st == 'normal' and ch in ('#', ';') : return line[:i]
            elif st == 'normal' and ch in (';')      : return line[:i]
            elif st == 'slcomm' and ch == '/'        : return line[:i - 1]
            elif st == 'slcomm'                      : st = 'normal'
            elif st == 'string' and ch == '\"'       : st = 'normal'
            elif st == 'string' and ch == '\\'       : st = 'escape'
            elif st == 'escape'                      : st = 'string'
        else:
            return line

    @staticmethod
    def _replace_adrp_line(line: str) -> str:
        if 'adrp' in line:
            line = line.replace('adrp', 'adr').replace('@PAGE', '')
        return line

    @staticmethod
    def _replace_adrp(src: List[str]) -> List[str]:
        back_label_count = 0
        adrp_label_map = {}
        new_src = []
        for line in src:
            line = Assembler._remove_comments(line)
            line = line.strip()

            if not line:
                continue
            # is instructions
            if line[-1] != ':' and line[0] != '.':
                instr = Instruction(line, back_label_count)
                if instr.ADRP_label:
                    back_label_count += 1
                    new_src.append(instr.asm_code)
                    new_src.append(instr.back_label + ':')
                    if instr.text_label in adrp_label_map:
                        adrp_label_map[instr.text_label] += [(instr.ADRP_label, instr.ADR_instr, instr.back_label)]
                    else:
                        adrp_label_map[instr.text_label] = [(instr.ADRP_label, instr.ADR_instr, instr.back_label)]
                else:
                    new_src.append(line)
            else:
                new_src.append(line)

        nn_src = []

        for line in new_src:
            if line[-1] == ':': # is label
                if line[:-1] in adrp_label_map:
                    for item in adrp_label_map[line[:-1]]:
                        nn_src.append(item[0] + ':')       # label that adrp will jump to
                        nn_src.append(item[1])             # adr to get really symbol address
                        nn_src.append('b ' + item[2])      # jump back to adrp next instruction
            nn_src.append(line)

        return nn_src

    def _parse(self, src: List[str]):
        # src = self._replace_adrp(o_src)

        for line in src:
            line = Assembler._remove_comments(line)
            line = line.strip()

            # skip empty lines
            if not line:
                continue

            # labels, resolve the offset
            if line[-1] == ':':
                self.code.label(line[:-1])
                continue

            # instructions
            if line[0] != '.':
                line = self._replace_adrp_line(line)
                self.code.instr(Instruction(line, 0))
                continue

            # parse the command
            cmd = Command.parse(line)
            func = self._commands.get(cmd.cmd)

            # handle the command
            if func is not None:
                func(cmd.args)
            else:
                raise SyntaxError('invalid assembly command: ' + cmd.cmd)

    def _reloc(self, rip: int = 0):
        for block in self.code.blocks:
            for instr in block.body:
                rip += self._reloc_one(instr, rip)

    def _reloc_one(self, instr: Instr, rip: int) -> int:
        if not isinstance(instr, (X86Instr, BranchInstr)):
            return instr.size(rip)
        elif instr.instr.need_reloc:
            return self._reloc_branch(instr.instr, rip)
        else:
            return instr.resize(self._reloc_normal(instr.instr, rip))

    def _reloc_branch(self, instr: Instruction, rip: int) -> int:
        label = instr.label_name
        if label is None:
            raise RuntimeError('cannnot found label name: %s' % instr.asm_code)
        if instr.mnemonic == 'adr' and label == 'Ltmp0':
            instr.set_label_offset(-4)
        else:
            instr.set_label_offset(self.code.get(label)- rip - instr.size)

        return instr.size

    def _reloc_normal(self, instr: Instruction, rip: int) -> int:
        if instr.need_reloc:
            raise SyntaxError('unresolved instruction when relocation: ' + instr.asm_code)
        return instr.size

    def _LE_4bytes_IntIntr_2_RawIntr(self):
        for block in self.code.blocks:
            block.if_all_IntInstr_then_2_RawInstr()

    def _declare(self, protos: PrototypeMap):
        if OUTPUT_RAW:
            self._declare_body_raw()
        else:
            name = next(iter(protos))
            self._declare_body(name[1:])
        self._declare_functions(protos)

    def _declare_body(self, name: str):
        size = self.code.stacksize(name)
        gosize = 0 if size < 16 else size-16
        self.out.append('TEXT ·_%s_entry__(SB), NOSPLIT, $%d' % (name, gosize))
        self.out.append('\tNO_LOCAL_POINTERS')
        # get current PC
        self.out.append('\tWORD $0x100000a0 // adr x0, .+20')
        # self.out.append('\t'+Instruction('add sp, sp, #%d' % size).encoded)
        self.out.append('\tMOVD R0, ret(FP)')
        self.out.append('\tRET')
        self._LE_4bytes_IntIntr_2_RawIntr()
        self._reloc()

        # instruction buffer
        pc = 0
        ins = self.code.instrs

        for v in ins:
            self.out.append(('// +%d\n' % pc if WITH_OFFS else '') + v.formatted(pc))
            pc += v.size(pc)

    def _declare_body_raw(self):
        self._reloc()

        # instruction buffer
        pc = 0
        ins = self.code.instrs

        # dump every instruction
        for v in ins:
            self.out.append(v.raw_formatted(pc))
            pc += v.size(pc)

    def _declare_function(self, name: str, proto: Prototype):
        offs = 0
        subr = name[1:]
        addr = self.code.get(subr)
        self.subr[subr] = addr
        size = self.code.stacksize(subr)

        m_size = size + 64
        # rsp_sub_size = size + 16

        if OUTPUT_RAW:
            return

        # function header and stack checking
        self.out.append('')
        # frame size is 16 to store x29 and x30
        # self.out.append('TEXT ·%s(SB), NOSPLIT | NOFRAME, $0-%d' % (name, proto.argspace))
        self.out.append('TEXT ·%s(SB), NOSPLIT, $%d-%d' % (name, 0, proto.argspace))
        self.out.append('\tNO_LOCAL_POINTERS')

        # add stack check if needed
        if m_size != 0:
            self.out.append('')
            self.out.append('_entry:')
            self.out.append('\tMOVD 16(g), R16')
            if size > 0:
             if size < (0x1 << 12) - 1:
                 self.out.append('\tSUB $%d, RSP, R17' % (m_size))
             elif size < (0x1 << 16) - 1:
                 self.out.append('\tMOVD $%d, R17' % (m_size))
                 self.out.append('\tSUB R17, RSP, R17')
             else:
                 raise RuntimeError('too large stack size: %d' % (m_size))
             self.out.append('\tCMP  R16, R17')
            else:
             self.out.append('\tCMP R16, RSP')
            self.out.append('\tBLS  _stack_grow')

        # function name
        self.out.append('')
        self.out.append('%s:' % subr)

        # self.out.append('\tMOVD.W R30, -16(RSP)')
        # self.out.append('\tMOVD R29, -8(RSP)')
        # self.out.append('\tSUB $8, RSP, R29')

        # intialize all the arguments
        for arg in proto.args:
            offs += arg.size
            op, reg = REG_MAP[arg.creg.reg]
            self.out.append('\t%s %s+%d(FP), %s' % (op, arg.name, offs - arg.size, reg))


        # Go ASM completely ignores the offset of the JMP instruction,
        # so we need to use indirect jumps instead for tail-call elimination
        
        # LEA and JUMP
        self.out.append('\tMOVD ·_subr_%s(SB), R11' % (subr))
        self.out.append('\tWORD $0x1000005e // adr x30, .+8')
        self.out.append('\tJMP (R11)')
        # self.out.append('\tCALL ·_%s_entry__(SB)  // %s' % (subr, subr))
        
        # normal functions, call the real function, and return the result
        if proto.retv is not None:
            self.out.append('\t%s, %s+%d(FP)' % (' '.join(REG_MAP[proto.retv.creg.reg]), proto.retv.name, offs))
        # Restore LR and Frame Pointer
        # self.out.append('\tLDP -8(RSP), (R29, R30)')
        # self.out.append('\tADD $16, RSP')
        
        self.out.append('\tRET')

        # add stack growing if needed
        if m_size != 0:
            self.out.append('')
            self.out.append('_stack_grow:')
            self.out.append('\tMOVD R30, R3')
            self.out.append('\tCALL runtime·morestack_noctxt<>(SB)')
            self.out.append('\tJMP  _entry')

    def _declare_functions(self, protos: PrototypeMap):
        for name, proto in sorted(protos.items()):
            if name[0] == '_':
                self._declare_function(name, proto)
            else:
                raise SyntaxError('function prototype must have a "_" prefix: ' + repr(name))

    def parse(self, src: List[str], proto: PrototypeMap):
        # self.code.instr(Instruction('adr x0, .'))
        # self.code.instr(Instruction('add sp, sp, #%d'%self.code.stacksize(name)))
        # self.code.instr(Instruction('ret'))
        # cmd = Command.parse(".p2align 4")
        # func = self._commands.get(cmd.cmd)
        # func(cmd.args)

        self._parse(src)
        self._declare(proto)

GOOS = {
    'aix',
    'android',
    'darwin',
    'dragonfly',
    'freebsd',
    'hurd',
    'illumos',
    'js',
    'linux',
    'nacl',
    'netbsd',
    'openbsd',
    'plan9',
    'solaris',
    'windows',
    'zos',
}

GOARCH = {
    '386',
    'amd64',
    'amd64p32',
    'arm',
    'armbe',
    'arm64',
    'arm64be',
    'ppc64',
    'ppc64le',
    'mips',
    'mipsle',
    'mips64',
    'mips64le',
    'mips64p32',
    'mips64p32le',
    'ppc',
    'riscv',
    'riscv64',
    's390',
    's390x',
    'sparc',
    'sparc64',
    'wasm',
}

def make_subr_filename(name: str) -> str:
    name = os.path.basename(name)
    base = os.path.splitext(name)[0].rsplit('_', 2)

    # construct the new name
    if base[-1] in GOOS:
        return '%s_subr_%s.go' % ('_'.join(base[:-1]), base[-1])
    elif base[-1] not in GOARCH:
        return '%s_subr.go' % '_'.join(base)
    elif len(base) > 2 and base[-2] in GOOS:
        return '%s_subr_%s_%s.go' % ('_'.join(base[:-2]), base[-2], base[-1])
    else:
        return '%s_subr_%s.go' % ('_'.join(base[:-1]), base[-1])

def parse_args():
    parser = argparse.ArgumentParser(description='Convert llvm asm to golang asm.')
    parser.add_argument('proto_file', type=str, help = 'The go file that declares go functions')
    parser.add_argument('asm_file', type=str, nargs='+', help = 'The llvm assembly file')
    parser.add_argument('-r', default=False, action='store_true', help = 'Ture: output as raw; default is False')
    return parser.parse_args()

def main():
    src = []
    args = parse_args()

    # check if optional flag is enabled
    global OUTPUT_RAW
    OUTPUT_RAW = False
    if args.r:
        OUTPUT_RAW = True

    proto_name = os.path.splitext(args.proto_file)[0]

    # parse the prototype
    with open(proto_name + '.go', 'r', newline = None) as fp:
        pkg, proto = PrototypeMap.parse(fp.read())

    # read all the sources, and combine them together
    for fn in args.asm_file:
        with open(fn, 'r', newline = None) as fp:
            src.extend(fp.read().splitlines())

    asm = Assembler()

    # convert the original sources
    if OUTPUT_RAW:
        asm.out.append('// +build arm64')
        asm.out.append('// Code generated by asm2asm, DO NOT EDIT.')
        asm.out.append('')
        asm.out.append('package %s' % pkg)
        asm.out.append('')
        ## native text
        asm.out.append('var Text%s = []byte{' % STUB_NAME)
    else:
        asm.out.append('// +build !noasm !appengine')
        asm.out.append('// Code generated by asm2asm, DO NOT EDIT.')
        asm.out.append('')
        asm.out.append('#include "go_asm.h"')
        asm.out.append('#include "funcdata.h"')
        asm.out.append('#include "textflag.h"')
        asm.out.append('')

    asm.parse(src, proto)

    if OUTPUT_RAW:
        asrc = proto_name[:proto_name.rfind('_')] + '_text_arm.go'
    else:
        asrc = proto_name + '.s'

    # save the converted result
    with open(asrc, 'w')  as fp:
        for line in asm.out:
            print(line, file = fp)
        if OUTPUT_RAW:
            print('}', file = fp)

    # calculate the subroutine stub file name
    subr = make_subr_filename(args.proto_file)
    subr = os.path.join(os.path.dirname(args.proto_file), subr)

    # save the compiled code stub
    with open(subr, 'w') as fp:
        print('// +build !noasm !appengine', file = fp)
        print('// Code generated by asm2asm, DO NOT EDIT.', file = fp)
        print(file = fp)
        print('package %s' % pkg, file = fp)

        # also save the actual function addresses if any
        if not asm.subr:
            return

        if OUTPUT_RAW:
            print(file = fp)
            print('import (\n\t`github.com/bytedance/sonic/loader`\n)', file = fp)

            # dump every entry for all functions
            print(file = fp)
            print('const (', file = fp)
            for name in asm.code.funcs.keys():
                addr = asm.code.get(name)
                if addr is not None:
                    print(f'    _entry_{name} = %d' % addr, file = fp)
            print(')', file = fp)

            # dump max stack depth for all functions
            print(file = fp)
            print('const (', file = fp)
            for name in asm.code.funcs.keys():
                print('    _stack_%s = %d' % (name, asm.code.stacksize(name)), file = fp)
            print(')', file = fp)

            # dump every text size for all functions
            print(file = fp)
            print('const (', file = fp)
            for name, pcsp in asm.code.funcs.items():
                if pcsp is not None:
                    # print(f'before {name} optimize {pcsp}')
                    pcsp.optimize()
                    # print(f'after {name} optimize {pcsp}')
                    print(f'    _size_{name} = %d' % (pcsp.maxpc - pcsp.entry), file = fp)
            print(')', file = fp)

            # dump every pcsp for all functions
            print(file = fp)
            print('var (', file = fp)
            for name, pcsp in asm.code.funcs.items():
                if pcsp is not None:
                    print(f'    _pcsp_{name} = %s' % pcsp, file = fp)
            print(')', file = fp)

            # insert native entry info
            print(file = fp)
            print('var Funcs = []loader.CFunc{', file = fp)
            print('    {"%s", 0, %d, 0, nil},' % (STUB_NAME, STUB_SIZE), file = fp)
            # dump every native function info for all functions
            for name in asm.code.funcs.keys():
                print('    {"%s", _entry_%s, _size_%s, _stack_%s, _pcsp_%s},' % (name, name, name, name, name), file = fp)
            print('}', file = fp)

        else:
            # native entry for entry function
            print(file = fp)
            print('//go:nosplit', file = fp)
            print('//go:noescape', file = fp)
            print('//goland:noinspection ALL', file = fp)
            for name, entry in asm.subr.items():
                print('func _%s_entry__() uintptr' % name, file = fp)
            
            # dump exported function entry for exported functions
            print(file = fp)
            print('var (', file = fp)
            mlen = max(len(s) for s in asm.subr)
            for name, entry in asm.subr.items():
                print('    _subr_%s uintptr = _%s_entry__() + %d' % (name.ljust(mlen, ' '), name, entry), file = fp)
                # print('    _subr_%s uintptr = %d' % (name.ljust(mlen, ' '), entry), file = fp)
            print(')', file = fp)

            # dump max stack depth for exported functions
            print(file = fp)
            print('const (', file = fp)
            for name in asm.subr.keys():
                print('    _stack_%s = %d' % (name, asm.code.stacksize(name)), file = fp)
            print(')', file = fp)

            # assign subroutine offsets to '_' to mute the "unused" warnings
            print(file = fp)
            print('var (', file = fp)
            for name in asm.subr:
                print('    _ = _subr_%s' % name, file = fp)
            print(')', file = fp)

            # dump every constant
            print(file = fp)
            print('const (', file = fp)
            for name in asm.subr:
                print('    _ = _stack_%s' % name, file = fp)
            else:
                print(')', file = fp)

if __name__ == '__main__':
    main()
