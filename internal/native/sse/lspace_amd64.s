// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__lspace_entry(SB), NOSPLIT, $8
	NO_LOCAL_POINTERS
	LONG $0xf20d8d4c; WORD $0xffff; BYTE $0xff  // leaq         $-14(%rip), %r9
	LONG $0x244c894c; BYTE $0x18  // movq         %r9, $24(%rsp)
	BYTE $0xc3  // retq         
	BYTE $0x00
	BYTE $0x00
	BYTE $0x00
	BYTE $0x00
	BYTE $0x00
	  // .p2align 4, 0x90
_lspace:
	BYTE $0x55  // pushq        %rbp
	WORD $0x8948; BYTE $0xe5  // movq         %rsp, %rbp
	WORD $0x3948; BYTE $0xd6  // cmpq         %rdx, %rsi
	LONG $0x004e840f; WORD $0x0000  // je           LBB0_1, $78(%rip)
	LONG $0x37048d4c  // leaq         (%rdi,%rsi), %r8
	LONG $0x3a448d48; BYTE $0x01  // leaq         $1(%rdx,%rdi), %rax
	WORD $0x2948; BYTE $0xf2  // subq         %rsi, %rdx
	QUAD $0x000100002600be48; WORD $0x0000  // movabsq      $4294977024, %rsi
	QUAD $0x9090909090909090; LONG $0x90909090; BYTE $0x90  // .p2align 4, 0x90
LBB0_3:
	LONG $0xff48be0f  // movsbl       $-1(%rax), %ecx
	WORD $0xf983; BYTE $0x20  // cmpl         $32, %ecx
	LONG $0x002c870f; WORD $0x0000  // ja           LBB0_5, $44(%rip)
	LONG $0xcea30f48  // btq          %rcx, %rsi
	LONG $0x0022830f; WORD $0x0000  // jae          LBB0_5, $34(%rip)
	WORD $0xff48; BYTE $0xc0  // incq         %rax
	WORD $0xff48; BYTE $0xc2  // incq         %rdx
	LONG $0xffdd850f; WORD $0xffff  // jne          LBB0_3, $-35(%rip)
	WORD $0x2949; BYTE $0xf8  // subq         %rdi, %r8
	WORD $0x894c; BYTE $0xc0  // movq         %r8, %rax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_1:
	WORD $0x0148; BYTE $0xfa  // addq         %rdi, %rdx
	WORD $0x8949; BYTE $0xd0  // movq         %rdx, %r8
	WORD $0x2949; BYTE $0xf8  // subq         %rdi, %r8
	WORD $0x894c; BYTE $0xc0  // movq         %r8, %rax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_5:
	WORD $0xf748; BYTE $0xd7  // notq         %rdi
	WORD $0x0148; BYTE $0xf8  // addq         %rdi, %rax
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         

TEXT ·__lspace(SB), NOSPLIT | NOFRAME, $0 - 32
	NO_LOCAL_POINTERS

_entry:
	MOVQ (TLS), R14
	LEAQ -8(SP), R12
	CMPQ R12, 16(R14)
	JBE  _stack_grow

_lspace:
	MOVQ sp+0(FP), DI
	MOVQ nb+8(FP), SI
	MOVQ off+16(FP), DX
	CALL ·__lspace_entry+32(SB)  // _lspace
	MOVQ AX, ret+24(FP)
	RET

_stack_grow:
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
