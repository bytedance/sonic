// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__vunsigned_entry(SB), NOSPLIT, $8
	NO_LOCAL_POINTERS
	LONG $0xeb0d8d4c; WORD $0xffff; BYTE $0xff  // leaq         $-21(%rip), %r9
	LONG $0x244c894c; BYTE $0x18  // movq         %r9, $24(%rsp)
	LONG $0x10c48348  // addq         $16, %rsp
	BYTE $0xc3  // retq         
	BYTE $0x00
	  // .p2align 4, 0x90
_vunsigned:
	BYTE $0x55  // pushq        %rbp
	WORD $0x8948; BYTE $0xe5  // movq         %rsp, %rbp
	WORD $0x8949; BYTE $0xd0  // movq         %rdx, %r8
	WORD $0x8b48; BYTE $0x0e  // movq         (%rsi), %rcx
	WORD $0x8b4c; BYTE $0x0f  // movq         (%rdi), %r9
	LONG $0x085f8b4c  // movq         $8(%rdi), %r11
	LONG $0x0902c748; WORD $0x0000; BYTE $0x00  // movq         $9, (%rdx)
	QUAD $0x000000000842c748  // movq         $0, $8(%rdx)
	QUAD $0x000000001042c748  // movq         $0, $16(%rdx)
	WORD $0x8b48; BYTE $0x06  // movq         (%rsi), %rax
	LONG $0x18428948  // movq         %rax, $24(%rdx)
	WORD $0x394c; BYTE $0xd9  // cmpq         %r11, %rcx
	LONG $0x0018830f; WORD $0x0000  // jae          LBB0_1, $24(%rip)
	LONG $0x09048a41  // movb         (%r9,%rcx), %al
	WORD $0x2d3c  // cmpb         $45, %al
	LONG $0x0018850f; WORD $0x0000  // jne          LBB0_4, $24(%rip)
LBB0_3:
	WORD $0x8948; BYTE $0x0e  // movq         %rcx, (%rsi)
	LONG $0xfa00c749; WORD $0xffff; BYTE $0xff  // movq         $-6, (%r8)
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_1:
	WORD $0x894c; BYTE $0x1e  // movq         %r11, (%rsi)
	LONG $0xff00c749; WORD $0xffff; BYTE $0xff  // movq         $-1, (%r8)
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_4:
	WORD $0x508d; BYTE $0xd0  // leal         $-48(%rax), %edx
	WORD $0xfa80; BYTE $0x0a  // cmpb         $10, %dl
	LONG $0x000c820f; WORD $0x0000  // jb           LBB0_6, $12(%rip)
	WORD $0x8948; BYTE $0x0e  // movq         %rcx, (%rsi)
	LONG $0xfe00c749; WORD $0xffff; BYTE $0xff  // movq         $-2, (%r8)
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_6:
	WORD $0x303c  // cmpb         $48, %al
	LONG $0x0026850f; WORD $0x0000  // jne          LBB0_7, $38(%rip)
	LONG $0x09448a41; BYTE $0x01  // movb         $1(%r9,%rcx), %al
	WORD $0xd204  // addb         $-46, %al
	WORD $0x373c  // cmpb         $55, %al
	LONG $0x00af870f; WORD $0x0000  // ja           LBB0_16, $175(%rip)
	WORD $0xb60f; BYTE $0xc0  // movzbl       %al, %eax
	QUAD $0x000000800001ba48; WORD $0x0080  // movabsq      $36028797027352577, %rdx
	LONG $0xc2a30f48  // btq          %rax, %rdx
	LONG $0x0098830f; WORD $0x0000  // jae          LBB0_16, $152(%rip)
LBB0_7:
	WORD $0xc031  // xorl         %eax, %eax
	LONG $0x000aba41; WORD $0x0000  // movl         $10, %r10d
	LONG $0x90909090; WORD $0x9090  // .p2align 4, 0x90
LBB0_8:
	WORD $0x394c; BYTE $0xd9  // cmpq         %r11, %rcx
	LONG $0x0078830f; WORD $0x0000  // jae          LBB0_20, $120(%rip)
	LONG $0x3cbe0f41; BYTE $0x09  // movsbl       (%r9,%rcx), %edi
	WORD $0x578d; BYTE $0xd0  // leal         $-48(%rdi), %edx
	WORD $0xfa80; BYTE $0x09  // cmpb         $9, %dl
	LONG $0x0049870f; WORD $0x0000  // ja           LBB0_17, $73(%rip)
	WORD $0xf749; BYTE $0xe2  // mulq         %r10
	LONG $0x0031800f; WORD $0x0000  // jo           LBB0_13, $49(%rip)
	WORD $0xff48; BYTE $0xc1  // incq         %rcx
	WORD $0xc783; BYTE $0xd0  // addl         $-48, %edi
	WORD $0x6348; BYTE $0xd7  // movslq       %edi, %rdx
	WORD $0x8948; BYTE $0xd7  // movq         %rdx, %rdi
	LONG $0x3fffc148  // sarq         $63, %rdi
	WORD $0x0148; BYTE $0xd0  // addq         %rdx, %rax
	LONG $0x00d78348  // adcq         $0, %rdi
	WORD $0xfa89  // movl         %edi, %edx
	WORD $0xe283; BYTE $0x01  // andl         $1, %edx
	WORD $0xf748; BYTE $0xda  // negq         %rdx
	WORD $0x3148; BYTE $0xd7  // xorq         %rdx, %rdi
	LONG $0x0009850f; WORD $0x0000  // jne          LBB0_13, $9(%rip)
	WORD $0x8548; BYTE $0xd2  // testq        %rdx, %rdx
	LONG $0xffac890f; WORD $0xffff  // jns          LBB0_8, $-84(%rip)
LBB0_13:
	WORD $0xff48; BYTE $0xc9  // decq         %rcx
	WORD $0x8948; BYTE $0x0e  // movq         %rcx, (%rsi)
	LONG $0xfb00c749; WORD $0xffff; BYTE $0xff  // movq         $-5, (%r8)
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_17:
	LONG $0x65ff8040  // cmpb         $101, %dil
	LONG $0xff27840f; WORD $0xffff  // je           LBB0_3, $-217(%rip)
	LONG $0x45ff8040  // cmpb         $69, %dil
	LONG $0xff1d840f; WORD $0xffff  // je           LBB0_3, $-227(%rip)
	LONG $0x2eff8040  // cmpb         $46, %dil
	LONG $0xff13840f; WORD $0xffff  // je           LBB0_3, $-237(%rip)
LBB0_20:
	WORD $0x8948; BYTE $0x0e  // movq         %rcx, (%rsi)
	LONG $0x10408949  // movq         %rax, $16(%r8)
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_16:
	WORD $0xff48; BYTE $0xc1  // incq         %rcx
	WORD $0x8948; BYTE $0x0e  // movq         %rcx, (%rsi)
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
	WORD $0x0000  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	LONG $0x00000002  // .long 2

TEXT ·__vunsigned(SB), NOSPLIT | NOFRAME, $0 - 24
	NO_LOCAL_POINTERS

_entry:
	MOVQ (TLS), R14
	LEAQ -8(SP), R12
	CMPQ R12, 16(R14)
	JBE  _stack_grow

_vunsigned:
	MOVQ s+0(FP), DI
	MOVQ p+8(FP), SI
	MOVQ v+16(FP), DX
	MOVQ ·_subr__vunsigned(SB), R9
	LONG $0x05158d4c; WORD $0x0000; BYTE $0x00  // leaq         $5(%rip), %r10
	WORD $0x5241  // pushq        %r10
	JMP R9
	RET

_stack_grow:
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
