#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import string
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

from peachpy import x86_64
from peachpy.x86_64 import generic
from peachpy.x86_64 import XMMRegister
from peachpy.x86_64.operand import is_rel32
from peachpy.x86_64.operand import MemoryOperand
from peachpy.x86_64.operand import MemoryAddress
from peachpy.x86_64.operand import RIPRelativeOffset
from peachpy.x86_64.instructions import Instruction as PInstr
from peachpy.x86_64.instructions import BranchInstruction

### Instruction Parser (GAS Syntax) ###

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

    @functools.cached_property
    def native(self) -> x86_64.registers.Register:
        if self.reg == 'rip':
            raise SyntaxError('%rip is not directly accessible')
        else:
            return getattr(x86_64.registers, self.reg)

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

REGISTERS = {
    'rax'   , 'eax'   , 'ax'    , 'al'    , 'ah'   ,
    'rbx'   , 'ebx'   , 'bx'    , 'bl'    , 'bh'   ,
    'rcx'   , 'ecx'   , 'cx'    , 'cl'    , 'ch'   ,
    'rdx'   , 'edx'   , 'dx'    , 'dl'    , 'dh'   ,
    'rsi'   , 'esi'   , 'si'    , 'sil'   ,
    'rdi'   , 'edi'   , 'di'    , 'dil'   ,
    'rbp'   , 'ebp'   , 'bp'    , 'bpl'   ,
    'rsp'   , 'esp'   , 'sp'    , 'spl'   ,
    'r8'    , 'r8d'   , 'r8w'   , 'r8b'   ,
    'r9'    , 'r9d'   , 'r9w'   , 'r9b'   ,
    'r10'   , 'r10d'  , 'r10w'  , 'r10b'  ,
    'r11'   , 'r11d'  , 'r11w'  , 'r11b'  ,
    'r12'   , 'r12d'  , 'r12w'  , 'r12b'  ,
    'r13'   , 'r13d'  , 'r13w'  , 'r13b'  ,
    'r14'   , 'r14d'  , 'r14w'  , 'r14b'  ,
    'r15'   , 'r15d'  , 'r15w'  , 'r15b'  ,
    'mm0'   , 'mm1'   , 'mm2'   , 'mm3'   , 'mm4'   , 'mm5'   , 'mm6'   , 'mm7'   ,
    'xmm0'  , 'xmm1'  , 'xmm2'  , 'xmm3'  , 'xmm4'  , 'xmm5'  , 'xmm6'  , 'xmm7'  ,
    'xmm8'  , 'xmm9'  , 'xmm10' , 'xmm11' , 'xmm12' , 'xmm13' , 'xmm14' , 'xmm15' ,
    'xmm16' , 'xmm17' , 'xmm18' , 'xmm19' , 'xmm20' , 'xmm21' , 'xmm22' , 'xmm23' ,
    'xmm24' , 'xmm25' , 'xmm26' , 'xmm27' , 'xmm28' , 'xmm29' , 'xmm30' , 'xmm31' ,
    'ymm0'  , 'ymm1'  , 'ymm2'  , 'ymm3'  , 'ymm4'  , 'ymm5'  , 'ymm6'  , 'ymm7'  ,
    'ymm8'  , 'ymm9'  , 'ymm10' , 'ymm11' , 'ymm12' , 'ymm13' , 'ymm14' , 'ymm15' ,
    'ymm16' , 'ymm17' , 'ymm18' , 'ymm19' , 'ymm20' , 'ymm21' , 'ymm22' , 'ymm23' ,
    'ymm24' , 'ymm25' , 'ymm26' , 'ymm27' , 'ymm28' , 'ymm29' , 'ymm30' , 'ymm31' ,
    'zmm0'  , 'zmm1'  , 'zmm2'  , 'zmm3'  , 'zmm4'  , 'zmm5'  , 'zmm6'  , 'zmm7'  ,
    'zmm8'  , 'zmm9'  , 'zmm10' , 'zmm11' , 'zmm12' , 'zmm13' , 'zmm14' , 'zmm15' ,
    'zmm16' , 'zmm17' , 'zmm18' , 'zmm19' , 'zmm20' , 'zmm21' , 'zmm22' , 'zmm23' ,
    'zmm24' , 'zmm25' , 'zmm26' , 'zmm27' , 'zmm28' , 'zmm29' , 'zmm30' , 'zmm31' ,
    'rip'   ,
}

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

class Tokenizer:
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
        ret, self.pos = self.src[self.pos], self.pos + 1
        return ret

    def _rid(self, s: str, allow_dot: bool) -> str:
        while not self._eof and (self._ch == '_' or self._ch.isalnum() or (allow_dot and self._ch == '.')):
            s += self._rch()
        else:
            return s

    def _reg(self) -> Token:
        if self._eof:
            raise SyntaxError('unexpected EOF when parsing register names')
        else:
            return self._regx()

    def _imm(self) -> Token:
        if self._eof:
            raise SyntaxError('unexpected EOF when parsing immediate values')
        else:
            return self._immx(self._rch())

    def _regx(self) -> Token:
        nch = self._rch()
        reg = self._rid(nch, allow_dot = False).lower()

        # check for register names
        if reg not in REGISTERS:
            raise SyntaxError('invalid register: ' + reg)
        else:
            return Token.reg(reg)

    def _immv(self, ch: str) -> int:
        while not self._eof and self._ch in string.digits:
            ch += self._rch()
        else:
            return int(ch)

    def _immx(self, ch: str) -> Token:
        if ch.isdigit():
            return Token.imm(self._immv(ch))
        elif ch == '-':
            return Token.imm(-self._immv(self._rch()))
        else:
            raise SyntaxError('unexpected character when parsing immediate value: ' + ch)

    def _name(self, ch: str) -> Token:
        return Token.name(self._rid(ch, allow_dot = True))

    def _read(self, ch: str) -> Token:
        if ch == '%':
            return self._reg()
        elif ch == '$':
            return self._imm()
        elif ch == '-':
            return Token.num(-self._immv(self._rch()))
        elif ch == '+':
            return Token.num(self._immv(self._rch()))
        elif ch.isdigit():
            return Token.num(self._immv(ch))
        elif ch.isidentifier():
            return self._name(ch)
        elif ch in ('(', ')', ',', '*'):
            return Token.punc(ch)
        else:
            raise SyntaxError('invalid character: ' + repr(ch))

    def next(self) -> Token:
        while not self._eof and self._ch.isspace():
            self.pos += 1
        else:
            return Token.end() if self._eof else self._read(self._rch())

class Instruction:
    comments: str
    mnemonic: str
    operands: List[Operand]

    def __init__(self, mnemonic: str, operands: List[Operand]):
        self.comments = ''
        self.operands = operands
        self.mnemonic = mnemonic.lower()

    def __str__(self):
        ops = ', '.join(map(str, self.operands))
        com = self.comments and '  /* %s */' % self.comments

        # ordinal instructions
        if not self.is_branch:
            return '%-12s %s%s' % (self.mnemonic, ops, com)
        elif len(self.operands) != 1:
            raise SyntaxError('invalid branch instruction: ' + self.mnemonic)
        elif isinstance(self.operands[0], Label):
            return '%-12s %s%s' % (self.mnemonic, ops, com)
        else:
            return '%-12s *%s%s' % (self.mnemonic, ops, com)

    def __repr__(self):
        return '{INSTR %s: %s%s}' % (
            self.mnemonic,
            ', '.join(map(repr, self.operands)),
            self.comments and ' (%s)' % self.comments
        )

    class Basic:
        @staticmethod
        def INT3(*args, **kwargs):
            return x86_64.INT(3, *args, **kwargs)

        @staticmethod
        def MOVQ(*args, **kwargs):
            if not any(isinstance(v, XMMRegister) for v in args):
                return x86_64.MOV(*args, **kwargs)
            else:
                return x86_64.MOVQ(*args, **kwargs)

    class BitShift:
        op: Type[PInstr]

        def __init__(self, op: Type[PInstr]):
            self.op = op

        def __call__(self, *args, **kwargs):
            if len(args) != 1:
                return self.op(*args, **kwargs)
            else:
                return self.op(*args, 1, **kwargs)

    class VectorCompare:
        fn: int
        op: Type[PInstr]

        def __init__(self, op: Type[PInstr], fn: int):
            self.fn = fn
            self.op = op

        def __call__(self, *args, **kwargs):
            return self.op(*args, self.fn, **kwargs)

    __instr_map__ = {
        'INT3'       : Basic.INT3,
        'SALB'       : BitShift(x86_64.SAL),
        'SALW'       : BitShift(x86_64.SAL),
        'SALL'       : BitShift(x86_64.SAL),
        'SALQ'       : BitShift(x86_64.SAL),
        'SARB'       : BitShift(x86_64.SAR),
        'SARW'       : BitShift(x86_64.SAR),
        'SARL'       : BitShift(x86_64.SAR),
        'SARQ'       : BitShift(x86_64.SAR),
        'SHLB'       : BitShift(x86_64.SHL),
        'SHLW'       : BitShift(x86_64.SHL),
        'SHLL'       : BitShift(x86_64.SHL),
        'SHLQ'       : BitShift(x86_64.SHL),
        'SHRB'       : BitShift(x86_64.SHR),
        'SHRW'       : BitShift(x86_64.SHR),
        'SHRL'       : BitShift(x86_64.SHR),
        'SHRQ'       : BitShift(x86_64.SHR),
        'MOVQ'       : Basic.MOVQ,
        'CBTW'       : x86_64.CBW,
        'CWTL'       : x86_64.CWDE,
        'CLTQ'       : x86_64.CDQE,
        'MOVZBW'     : x86_64.MOVZX,
        'MOVZBL'     : x86_64.MOVZX,
        'MOVZWL'     : x86_64.MOVZX,
        'MOVZBQ'     : x86_64.MOVZX,
        'MOVZWQ'     : x86_64.MOVZX,
        'MOVSBW'     : x86_64.MOVSX,
        'MOVSBL'     : x86_64.MOVSX,
        'MOVSWL'     : x86_64.MOVSX,
        'MOVSBQ'     : x86_64.MOVSX,
        'MOVSWQ'     : x86_64.MOVSX,
        'MOVSLQ'     : x86_64.MOVSXD,
        'MOVABSQ'    : x86_64.MOV,
        'VCMPEQPS'   : VectorCompare(x86_64.VCMPPS, 0x00),
        'VCMPTRUEPS' : VectorCompare(x86_64.VCMPPS, 0x0f),
    }

    @functools.cached_property
    def _instr(self) -> Union[Type[PInstr], Callable[..., PInstr]]:
        name = self.mnemonic.upper()
        func = self.__instr_map__.get(name)

        # not found, resolve as x86_64 instruction
        if func is None:
            func = getattr(x86_64, name, None)

        # try with size suffix removed (only for generic instructions)
        if func is None and name[-1] in 'BWLQ':
            func = getattr(generic, name[:-1], func)

        # still not found, it should be an error
        if func is None:
            raise SyntaxError('unknown instruction: ' + self.mnemonic)
        else:
            return func
    
    @property
    def jmptab(self) -> Optional[str]:
        if self.mnemonic == 'leaq' and isinstance(self.operands[0], Memory) and self.operands[0].base.reg == 'rip':
            dis = self.operands[0].disp
            if dis and dis.ref.find(CLANG_JUMPTABLE_LABLE) != -1:
                return dis.ref

    @property
    def _instr_size(self) -> Optional[int]:
        ops = self.operands
        key = self.mnemonic.upper()

        # special case of sign/zero extension instructions
        if key in self.__instr_size__:
            return self.__instr_size__[key]

        # check for register operands
        for op in ops:
            if isinstance(op, Register):
                return None

        # check for size suffix, and this only applies to generic instructions
        if key[-1] not in self.__size_map__ or not hasattr(generic, key[:-1]):
            raise SyntaxError('ambiguous operand sizes')
        else:
            return self.__size_map__[key[-1]]

    __size_map__ = {
        'B': 1,
        'W': 2,
        'L': 4,
        'Q': 8,
    }

    __instr_size__ = {
        'MOVZBW' : 1,
        'MOVZBL' : 1,
        'MOVZWL' : 2,
        'MOVZBQ' : 1,
        'MOVZWQ' : 2,
        'MOVSBW' : 1,
        'MOVSBL' : 1,
        'MOVSWL' : 2,
        'MOVSBQ' : 1,
        'MOVSWQ' : 2,
        'MOVSLQ' : 4,
    }

    @staticmethod
    def _encode_r32(ins: PInstr) -> bytes:
        ret = [fn(ins.operands) for _, fn in ins.encodings]
        ret.sort(key = len)
        return ret[-1]

    @classmethod
    def _encode_ins(cls, ins: PInstr, force_rel32: bool = False) -> bytes:
        if not isinstance(ins, BranchInstruction):
            return ins.encode()
        elif not is_rel32(ins.operands[0]) or not force_rel32:
            return ins.encode()
        else:
            return cls._encode_r32(ins)

    def _encode_rel(self, rel: Label, sizing: bool, offset: int) -> RIPRelativeOffset:
        if rel.offs is not None:
            return RIPRelativeOffset(rel.offs)
        elif sizing:
            return RIPRelativeOffset(offset)
        else:
            raise SyntaxError('unresolved reference to name: ' + rel.name)

    def _encode_mem(self, mem: Memory, sizing: bool, offset: int) -> MemoryOperand:
        if mem.base is not None and mem.base.reg == 'rip':
            return self._encode_mem_rip(mem, sizing, offset)
        else:
            return self._encode_mem_std(mem, sizing, offset)

    def _encode_mem_rip(self, mem: Memory, sizing: bool, offset: int) -> MemoryOperand:
        if mem.disp is None:
            return MemoryOperand(RIPRelativeOffset(0))
        elif mem.index is not None:
            raise SyntaxError('%rip relative addresing does not support indexing')
        elif isinstance(mem.disp, Immediate):
            return MemoryOperand(RIPRelativeOffset(mem.disp.val))
        elif isinstance(mem.disp, Reference):
            return MemoryOperand(RIPRelativeOffset(offset if sizing else mem.disp.offset))
        else:
            raise RuntimeError('illegal memory displacement')

    def _encode_mem_std(self, mem: Memory, sizing: bool, offset: int) -> MemoryOperand:
        disp  = 0
        base  = None
        index = None
        scale = None

        # add optional base
        if mem.base is not None:
            base = mem.base.native

        # add optional indexing
        if mem.index is not None:
            scale = mem.index.scale
            index = mem.index.base.native

        # add optional displacement
        if mem.disp is not None:
            if isinstance(mem.disp, Immediate):
                disp = mem.disp.val
            elif isinstance(mem.disp, Reference):
                disp = offset if sizing else mem.disp.offset
            else:
                raise RuntimeError('illegal memory displacement')

        # construct the memory address
        return MemoryOperand(
            size    = self._instr_size,
            address = MemoryAddress(base, index, scale, disp),
        )

    def _encode_operands(self, sizing: bool, offset: int) -> Iterable[Any]:
        for op in self.operands:
            if isinstance(op, Label):
                yield self._encode_rel(op, sizing, offset)
            elif isinstance(op, Memory):
                yield self._encode_mem(op, sizing, offset)
            elif isinstance(op, Register):
                yield op.native
            elif isinstance(op, Immediate):
                yield op.val
            else:
                raise SyntaxError('cannot encode %s as operand' % repr(op))

    def _encode_branch_rel(self, rel: Label) -> str:
        if rel.offs is not None:
            return self._encode_normal_instr()
        else:
            raise RuntimeError('invalid relative branching instruction')
        
    def _raw_branch_rel(self, rel: Label) -> bytes:
        if rel.offs is not None:
            return self._raw_normal_instr()
        else:
            raise RuntimeError('invalid relative branching instruction')

    def _encode_branch_mem(self, mem: Memory) -> str:
        raise NotImplementedError('not implemented: memory indirect jump')
    
    def _raw_branch_mem(self, mem: Memory) -> bytes:
        raise NotImplementedError('not implemented: memory indirect jump')

    def _encode_branch_reg(self, reg: Register) -> str:
        if reg.reg == 'rip':
            raise SyntaxError('%rip cannot be used as a jump target')
        elif self.mnemonic != 'jmpq':
            raise SyntaxError('invalid indirect jump for instruction: ' + self.mnemonic)
        else:
            return x86_64.JMP(reg.native).format('go')
        
    def _raw_branch_reg(self, reg: Register) -> bytes:
        if reg.reg == 'rip':
            raise SyntaxError('%rip cannot be used as a jump target')
        elif self.mnemonic != 'jmpq':
            raise SyntaxError('invalid indirect jump for instruction: ' + self.mnemonic)
        else:
            return x86_64.JMP(reg.native).encode()

    def _encode_branch_instr(self) -> str:
        if len(self.operands) != 1:
            raise RuntimeError('illegal branch instruction')
        elif isinstance(self.operands[0], Label):
            return self._encode_branch_rel(self.operands[0])
        elif isinstance(self.operands[0], Memory):
            return self._encode_branch_mem(self.operands[0])
        elif isinstance(self.operands[0], Register):
            return self._encode_branch_reg(self.operands[0])
        else:
            raise RuntimeError('invalid operand type ' + repr(self.operands[0]))
        
    def _raw_branch_instr(self) -> str:
        if len(self.operands) != 1:
            raise RuntimeError('illegal branch instruction')
        elif isinstance(self.operands[0], Label):
            return self._raw_branch_rel(self.operands[0])
        elif isinstance(self.operands[0], Memory):
            return self._raw_branch_mem(self.operands[0])
        elif isinstance(self.operands[0], Register):
            return self._raw_branch_reg(self.operands[0])
        else:
            raise RuntimeError('invalid operand type ' + repr(self.operands[0]))

    def _encode_normal_instr(self) -> str:
        ops = self._encode_operands(False, 0)
        ret = self._instr(*list(ops)[::-1])

        # encode all instructions as raw bytes
        if not self.is_branch_label:
            return self.encode(self._encode_ins(ret), str(self))
        else:
            return self.encode(self._encode_ins(ret, force_rel32 = True), '%s, $%s(%%rip)' % (self, self.operands[0].offs))
        
    def _raw_normal_instr(self) -> str:
        ops = self._encode_operands(False, 0)
        ret = self._instr(*list(ops)[::-1])

        # encode all instructions as raw bytes
        if not self.is_branch_label:
            return self._encode_ins(ret)
        else:
            return self._encode_ins(ret, force_rel32 = True)


    @property
    def size(self) -> int:
        return self.encoded_size(0)

    @functools.cached_property
    def encoded(self) -> str:
        if self.is_branch:
            return self._encode_branch_instr()
        else:
            return self._encode_normal_instr()
        
    def raw(self) -> bytes:
        if self.is_branch:
            return self._raw_branch_instr()
        else:
            return self._raw_normal_instr()

    @functools.cached_property
    def is_return(self) -> bool:
        return self._instr is x86_64.RET

    @functools.cached_property
    def is_invoke(self) -> bool:
        return self._instr is x86_64.CALL

    @functools.cached_property
    def is_branch(self) -> bool:
        try:
            return self.is_invoke or issubclass(self._instr, BranchInstruction)
        except TypeError:
            return False

    @functools.cached_property
    def is_jmp(self) -> bool:
        return self._instr is x86_64.JMP
    
    @functools.cached_property
    def is_jmpq(self) -> bool:
        return self.mnemonic == 'jmpq'

    @property
    def is_branch_label(self) -> bool:
        return self.is_branch and isinstance(self.operands[0], Label)

    def encoded_size(self, offset: int) -> int:
        op = self._encode_operands(True, offset)
        return len(self._encode_ins(self._instr(*list(op)[::-1]), force_rel32 = True))

    @classmethod
    def parse(cls, line: str) -> 'Instruction':
        lex = Tokenizer(line)
        ntk = lex.next()

        # the first token must be a name
        if ntk.tag != TOKEN_NAME:
            raise SyntaxError('mnemonic expected, got ' + repr(ntk))
        else:
            return cls(ntk.val, cls._parse_operands(lex))

    @staticmethod
    def encode(buf: bytes, comments: str = '') -> str:
        i = 0
        r = []
        n = len(buf)

        # try "QUAD" first
        while i < n - 7:
            r.append('QUAD $0x%016x' % int.from_bytes(buf[i:i + 8], 'little'))
            i += 8

        # then "LONG"
        while i < n - 3:
            r.append('LONG $0x%08x' % int.from_bytes(buf[i:i + 4], 'little'))
            i += 4

        # then "SHORT"
        while i < n - 1:
            r.append('WORD $0x%04x' % int.from_bytes(buf[i:i + 2], 'little'))
            i += 2

        # then "BYTE"
        while i < n:
            r.append('BYTE $0x%02x' % buf[i])
            i += 1

        # join them together, and attach the comment if any
        if not comments:
            return '; '.join(r)
        else:
            return '%s  // %s' % ('; '.join(r), comments)

    Reg  = Optional[Register]
    Disp = Optional[Displacement]

    @classmethod
    def _parse_mend(cls, ntk: Token, base: Reg, index: Register, scale: int, disp: Disp) -> Operand:
        if ntk.tag != TOKEN_PUNC or ntk.val != ')':
            raise SyntaxError('")" expected, got ' + repr(ntk))
        else:
            return Memory(base, disp, Index(index, scale))

    @classmethod
    def _parse_base(cls, lex: Tokenizer, ntk: Token, disp: Disp) -> Operand:
        if ntk.tag == TOKEN_REG:
            return cls._parse_idelim(lex, lex.next(), Register(ntk.val), disp)
        elif ntk.tag == TOKEN_PUNC and ntk.val == ',':
            return cls._parse_ibase(lex, lex.next(), None, disp)
        else:
            raise SyntaxError('register expected, got ' + repr(ntk))

    @classmethod
    def _parse_ibase(cls, lex: Tokenizer, ntk: Token, base: Reg, disp: Disp) -> Operand:
        if ntk.tag != TOKEN_REG:
            raise SyntaxError('register expected, got ' + repr(ntk))
        else:
            return cls._parse_sdelim(lex, lex.next(), base, Register(ntk.val), disp)

    @classmethod
    def _parse_idelim(cls, lex: Tokenizer, ntk: Token, base: Reg, disp: Disp) -> Operand:
        if ntk.tag == TOKEN_END:
            raise SyntaxError('unexpected EOF when parsing memory operands')
        elif ntk.tag == TOKEN_PUNC and ntk.val == ')':
            return Memory(base, disp, None)
        elif ntk.tag == TOKEN_PUNC and ntk.val == ',':
            return cls._parse_ibase(lex, lex.next(), base, disp)
        else:
            raise SyntaxError('"," or ")" expected, got ' + repr(ntk))

    @classmethod
    def _parse_iscale(cls, lex: Tokenizer, ntk: Token, base: Reg, index: Register, disp: Disp) -> Operand:
        if ntk.tag != TOKEN_NUM:
            raise SyntaxError('integer expected, got ' + repr(ntk))
        elif ntk.val not in (1, 2, 4, 8):
            raise SyntaxError('indexing scale can only be 1, 2, 4 or 8')
        else:
            return cls._parse_mend(lex.next(), base, index, ntk.val, disp)

    @classmethod
    def _parse_sdelim(cls, lex: Tokenizer, ntk: Token, base: Reg, index: Register, disp: Disp) -> Operand:
        if ntk.tag == TOKEN_END:
            raise SyntaxError('unexpected EOF when parsing memory operands')
        elif ntk.tag == TOKEN_PUNC and ntk.val == ')':
            return Memory(base, disp, Index(index))
        elif ntk.tag == TOKEN_PUNC and ntk.val == ',':
            return cls._parse_iscale(lex, lex.next(), base, index, disp)
        else:
            raise SyntaxError('"," or ")" expected, got ' + repr(ntk))

    @classmethod
    def _parse_refmem(cls, lex: Tokenizer, ntk: Token, ref: str) -> Operand:
        if ntk.tag == TOKEN_END:
            return Label(ref)
        elif ntk.tag == TOKEN_PUNC and ntk.val == '(' :
            return cls._parse_memory(lex, ntk, Reference(ref, 0))
        elif ntk.tag == TOKEN_NUM:
            ntk = lex.next()
            if ntk.tag == TOKEN_PUNC and ntk.val == '(':
                return cls._parse_refmem(lex, ntk, Reference(ref, ntk.val))
        
        raise SyntaxError(f'identifier "{ref}" must either be a label or a displacement reference')

    @classmethod
    def _parse_memory(cls, lex: Tokenizer, ntk: Token, disp: Optional[Displacement]) -> Operand:
        if ntk.tag != TOKEN_PUNC or ntk.val != '(':
            raise SyntaxError('"(" expected, got ' + repr(ntk))
        else:
            return cls._parse_base(lex, lex.next(), disp)

    @classmethod
    def _parse_operand(cls, lex: Tokenizer, ntk: Token, can_indir: bool = True) -> Operand:
        if ntk.tag == TOKEN_REG:
            return Register(ntk.val)
        elif ntk.tag == TOKEN_IMM:
            return Immediate(ntk.val)
        elif ntk.tag == TOKEN_NUM:
            return cls._parse_memory(lex, lex.next(), Immediate(ntk.val))
        elif ntk.tag == TOKEN_NAME:
            return cls._parse_refmem(lex, lex.next(), ntk.val)
        elif ntk.tag == TOKEN_PUNC and ntk.val == '(':
            return cls._parse_memory(lex, ntk, None)
        elif ntk.tag == TOKEN_PUNC and ntk.val == '*' and can_indir:
            return cls._parse_operand(lex, lex.next(), False)
        else:
            raise SyntaxError('invalid token: ' + repr(ntk))

    @classmethod
    def _parse_operands(cls, lex: Tokenizer) -> List[Operand]:
        ret = []
        ntk = lex.next()

        # check for empty operand
        if ntk.tag == TOKEN_END:
            return []

        # parse every operand
        while True:
            ret.append(cls._parse_operand(lex, ntk))
            ntk = lex.next()

            # check for the ',' delimiter or the end of input
            if ntk.tag == TOKEN_PUNC and ntk.val == ',':
                ntk = lex.next()
            elif ntk.tag != TOKEN_END:
                raise SyntaxError('"," expected, got ' + repr(ntk))
            else:
                return ret

### Prototype Parser ###

ARGS_ORDER_C = [
    Register('rdi'),
    Register('rsi'),
    Register('rdx'),
    Register('rcx'),
    Register('r8'),
    Register('r9'),
]

ARGS_ORDER_GO = [
    Register('rax'),
    Register('rbx'),
    Register('rcx'),
    Register('rdi'),
    Register('rsi'),
    Register('r8'),
]

FPARGS_ORDER = [
    Register('xmm0'),
    Register('xmm1'),
    Register('xmm2'),
    Register('xmm3'),
    Register('xmm4'),
    Register('xmm5'),
    Register('xmm6'),
    Register('xmm7'),
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
        reg = Register('xmm0') if xmm else Register('rax')
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

    def size(self, pc: int) -> int:
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
        return '\t' + str(self.instr.encoded)
    
    def raw_formatted(self, pc: int) -> str:
        return Instr.raw_formatted(self.instr._raw_normal_instr(), str(self.instr), pc)

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
        return Instr.raw_formatted(self.instr._raw_branch_instr(), str(self.instr), pc)
    
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
    'rax'  : ('MOVQ'  , 'AX'),
    'rdi'  : ('MOVQ'  , 'DI'),
    'rsi'  : ('MOVQ'  , 'SI'),
    'rdx'  : ('MOVQ'  , 'DX'),
    'rcx'  : ('MOVQ'  , 'CX'),
    'r8'   : ('MOVQ'  , 'R8'),
    'r9'   : ('MOVQ'  , 'R9'),
    'xmm0' : ('MOVSD' , 'X0'),
    'xmm1' : ('MOVSD' , 'X1'),
    'xmm2' : ('MOVSD' , 'X2'),
    'xmm3' : ('MOVSD' , 'X3'),
    'xmm4' : ('MOVSD' , 'X4'),
    'xmm5' : ('MOVSD' , 'X5'),
    'xmm6' : ('MOVSD' , 'X6'),
    'xmm7' : ('MOVSD' , 'X7'),
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

    def __init__(self):
        self.dead   = False
        self.labels = {}
        self.export = False
        self.blocks = [BasicBlock.annonymous()]
        self.jmptabs = {}
        self.funcs = {}
    
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
        if v & 7 == 0:
            return v
        else:
            print('* warning: SP is not aligned with 8 bytes.', file = sys.stderr)
            return (v + 7) & -8

    @staticmethod
    def _is_spadj(ins: Instruction) -> bool:
        return len(ins.operands) == 2                 and \
               isinstance(ins.operands[0], Immediate) and \
               isinstance(ins.operands[1], Register)  and \
               ins.operands[1].reg == 'rsp'

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
                size += block.size_of(size) + adj
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

    __instr_repl__ = {
        'movdqa'  : 'movdqu',
        'movaps'  : 'movups',
        'vmovdqa' : 'vmovdqu',
        'vmovaps' : 'vmovups',
        'vmovapd' : 'vmovupd',
    }

    def _check_align(self, instr: Instruction) -> bool:
        if instr.mnemonic in self.__instr_repl__:
            # NOTICE: since we need use unaligned instruction, thus SP can be fixed according to PC
            for op in instr.operands:
                if isinstance(op, Memory):
                    if op.base is not None and (op.base.reg == 'rbp' or op.base.reg == 'rsp'):
                        instr.mnemonic = self.__instr_repl__[instr.mnemonic]
                        return False
        elif instr.mnemonic == 'andq' and self._is_spadj(instr):
            # NOTICE: since we always use unaligned instruction above, we don't need align SP
            return True

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
            if instr.is_jmp: # jmp
                self._kill(instr.operands[0].name)
                
            elif instr.is_invoke: # call
                fname = instr.operands[0].name
                self._split(self._make(fname, func = True))
                
            else: # jeq, ja, jae ...
                self._split(self._make(instr.operands[0].name)) 

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
                args = ins.instr.operands

                # check for instructions
                if name == 'retq':
                    close = True
                elif name == 'popq':
                    diff = -8
                elif name == 'pushq':
                    diff = 8
                elif name == 'addq' and self._is_spadj(ins.instr):
                    diff = -self._mk_align(args[0].val)
                elif name == 'subq' and self._is_spadj(ins.instr):
                    diff = self._mk_align(args[0].val)
                    
                # FIXME: andq is usually used for aligment of memory address, we can't handle it correctly now
                # elif name == 'andq' and self._is_spadj(ins.instr): 
                #     diff = self._mk_align(max(-args[0].val - 8, 0))
                
                cursp += diff
                
                #NOTICE: pcsp no need to update here
                if name == 'callq':
                    cursp += 8
                        
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
            raise SyntaxError('unresolved reference to name: ' + key)
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

    def stacksize(self, name: str) -> int:
        if name not in self.labels:
            raise SyntaxError('undefined function: ' + name)
        else:
            return self._trace_block(self.labels[name], None)
        
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

_STUB_NAME = '__native_entry__'
STUB_SIZE = 67
WITH_OFFS = os.getenv('ASM2ASM_DEBUG_OFFSET', '').lower() in ('1', 'yes', 'true')

def stub_name(name :str) -> str:
    return name+ '_entry'

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
        for i in range(0, len(v), 16):
            self.code.emit(v[i:i + 16], '%s %d, %s' % (cmd, len(v[i:i + 16]), repr(v[i:i + 16])[1:]))

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
            '.p2align'                 : self._cmd_p2align,
            '.section'                 : self._cmd_nop,
            '.data_region'             : self._cmd_nop,
            '.build_version'           : self._cmd_nop,
            '.end_data_region'         : self._cmd_nop,
            '.subsections_via_symbols' : self._cmd_nop,
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
            elif st == 'normal' and ch in ('#', ';') : return line[:i]
            elif st == 'slcomm' and ch == '/'        : return line[:i - 1]
            elif st == 'slcomm'                      : st = 'normal'
            elif st == 'string' and ch == '\"'       : st = 'normal'
            elif st == 'string' and ch == '\\'       : st = 'escape'
            elif st == 'escape'                      : st = 'string'
        else:
            return line

    def _parse(self, src: List[str]):
        for line in src:
            line = self._remove_comments(line)
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
                self.code.instr(Instruction.parse(line))
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
        elif instr.instr.is_branch_label and isinstance(instr.instr.operands[0], Label):
            return self._reloc_branch(instr.instr, rip)
        else:
            return instr.resize(self._reloc_normal(instr.instr, rip))

    def _reloc_branch(self, instr: Instruction, rip: int) -> int:
        instr.operands[0].resolve(self.code.get(instr.operands[0].name) - rip - instr.size)
        return instr.size

    def _reloc_normal(self, instr: Instruction, rip: int) -> int:
        msg = []
        ops = instr.operands

        # relocate RIP relative operands
        for i, op in enumerate(ops):
            if self._is_rip_relative(op):
                if self.code.has(str(op.disp.ref)):
                    self._reloc_static(op.disp, msg, rip + instr.size)
                else:
                    raise SyntaxError('unresolved reference to name ' + str(op.disp.ref))

        # attach comments if any
        instr.comments = ', '.join(msg) or instr.comments
        return instr.size

    def _reloc_static(self, ref: Reference, msg: List[str], rip: int):
        msg.append('%s+%d(%%rip)' % (ref.ref, ref.disp))
        ref.resolve(self.code.get(str(ref.ref)) - rip)

    def _declare(self, protos: PrototypeMap):
        if OUTPUT_RAW:
            self._declare_body_raw()
        self._declare_functions(protos)

    def _declare_body(self, name :str):
        size = self.code.stacksize(name[1:])
        gosize = 0 if size < 8 else size - 8
        self.out.append('TEXT ·%s(SB), NOSPLIT, $%d' % (stub_name(name), gosize))
        self.out.append('\tNO_LOCAL_POINTERS')
        self._reloc()

        # instruction buffer
        pc = 0
        ins = self.code.instrs

        # dump every instruction
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
        size = self.code.pcsp(subr, addr)   
        # 14 for reserve go frame instructions
        addr += 14     
        self.subr[subr] = addr

        if OUTPUT_RAW:
            return
        
        # function header and stack checking
        self.out.append('')
        self.out.append('TEXT ·%s(SB), NOSPLIT | NOFRAME, $0 - %d' % (name, proto.argspace))
        self.out.append('\tNO_LOCAL_POINTERS')
        
        # add stack check if needed
        if size != 0:
            self.out.append('')
            self.out.append('_entry:')
            self.out.append('\tMOVQ (TLS), R14')
            self.out.append('\tLEAQ -%d(SP), R12' % size)
            self.out.append('\tCMPQ R12, 16(R14)')
            self.out.append('\tJBE  _stack_grow')

        # function name
        self.out.append('')
        self.out.append('%s:' % subr)

        # intialize all the arguments
        for arg in proto.args:
            offs += arg.size
            op, reg = REG_MAP[arg.creg.reg]
            self.out.append('\t%s %s+%d(FP), %s' % (op, arg.name, offs - arg.size, reg))

        # the function starts at zero
        if addr == 0 and proto.retv is None:
            self.out.append('\tJMP ·%s(SB)  // %s' % (stub_name(name), subr))

        # Go ASM completely ignores the offset of the JMP instruction,
        # so we need to use indirect jumps instead for tail-call elimination
        elif proto.retv is None:
            self.out.append('\tLEAQ ·%s+%d(SB), AX  // %s' % (stub_name(name), addr, subr))
            self.out.append('\tJMP AX')

        # normal functions, call the real function, and return the result
        else:
            self.out.append('\tCALL ·%s+%d(SB)  // %s' % (stub_name(name), addr, subr))
            self.out.append('\t%s, %s+%d(FP)' % (' '.join(REG_MAP[proto.retv.creg.reg]), proto.retv.name, offs))
            self.out.append('\tRET')

        # add stack growing if needed
        if size != 0:
            self.out.append('')
            self.out.append('_stack_grow:')
            self.out.append('\tCALL runtime·morestack_noctxt<>(SB)')
            self.out.append('\tJMP  _entry')

    def _declare_functions(self, protos: PrototypeMap):
        for name, proto in sorted(protos.items()):
            if name[0] == '_':
                self._declare_body(name)
                self._declare_function(name, proto)
            else:
                raise SyntaxError('function prototype must have a "_" prefix: ' + repr(name))

    def parse(self, src: List[str], proto: PrototypeMap):
        self.code.instr(Instruction('leaq', [Memory(Register('rip'), Immediate(-7), None), Register('rax')]))
        self.code.instr(Instruction('movq', [Register('rax'), Memory(Register('rsp'), Immediate(8), None)]))
        self.code.instr(Instruction('retq', []))
        self._parse(src)
        # print("DEBUG...")
        # self.code.debug(0, [
        #     X86Instr(Instruction('int3', []))
        #     # X86Instr(Instruction('xorq', [Register('rax'), Register('rax')])),
        #     # X86Instr(Instruction('movq', [Memory(Register('rax'), Immediate(0), None), Register('rax')]))
        # ])
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

def main():
    src = []
    asm = Assembler()
    
    
    # check for arguments
    if len(sys.argv) < 3:
        print('* usage: %s [-r|-d] <output-file> <clang-asm> ...' % sys.argv[0], file = sys.stderr)
        sys.exit(1)

    # check if optional flag is enabled
    global OUTPUT_RAW
    OUTPUT_RAW = False
    if len(sys.argv) >= 4:
        i = 0
        while i<len(sys.argv):
            flag = sys.argv[i]
            if flag == '-r':
                OUTPUT_RAW = True
                for j in range(i, len(sys.argv)-1):
                    sys.argv[j] = sys.argv[j + 1]  
                sys.argv.pop()
                continue
            i += 1
            
    # parse the prototype
    with open(os.path.splitext(sys.argv[1])[0] + '.go', 'r', newline = None) as fp:
        pkg, proto = PrototypeMap.parse(fp.read())

    # read all the sources, and combine them together
    for fn in sys.argv[2:]:
        with open(fn, 'r', newline = None) as fp:
            src.extend(fp.read().splitlines())

    # convert the original sources
    if OUTPUT_RAW:
        asm.out.append('// +build amd64')
        asm.out.append('// Code generated by asm2asm, DO NOT EDIT.')
        asm.out.append('')
        asm.out.append('package %s' % pkg)
        asm.out.append('')
        ## native text
        asm.out.append('var Text%s = []byte{' % _STUB_NAME)
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
        asrc = os.path.splitext(sys.argv[1])[0]
        asrc = asrc[:asrc.rfind('_')] + '_text_amd64.go'
    else:
        asrc = os.path.splitext(sys.argv[1])[0] + '.s'
      
    # save the converted result  
    with open(asrc, 'w')  as fp:
        for line in asm.out:
            print(line, file = fp)
        if OUTPUT_RAW:
            print('}', file = fp)

    # calculate the subroutine stub file name
    subr = make_subr_filename(sys.argv[1])
    subr = os.path.join(os.path.dirname(sys.argv[1]), subr)

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
            print('    {"%s", 0, %d, 0, nil},' % (_STUB_NAME, STUB_SIZE), file = fp)
            # dump every native function info for all functions
            for name in asm.code.funcs.keys():
                print('    {"%s", _entry_%s, _size_%s, _stack_%s, _pcsp_%s},' % (name, name, name, name, name), file = fp)
            print('}', file = fp)

        else:
            print(file = fp)
            print('import (\n`github.com/bytedance/sonic/internal/rt`\n)\n', file = fp)
            print('//go:nosplit', file = fp)
            print('//go:noescape', file = fp)
            print('//goland:noinspection ALL', file = fp)
            
            # native entry for entry function
            for (name, _) in asm.subr.items():
                print('func _%s() uintptr' % stub_name(name), file = fp)
            
            # dump exported function entry for exported functions
            print(file = fp)
            print('var (', file = fp)
            mlen = max(len(s) for s in asm.subr)
            for name, entry in asm.subr.items():
                print('    _subr_%s uintptr = rt.GetFuncPC(_%s) + %d' % (name.ljust(mlen, ' '), stub_name(name), entry), file = fp)
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
