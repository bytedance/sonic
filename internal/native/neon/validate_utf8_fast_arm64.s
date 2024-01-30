// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__validate_utf8_fast_entry__(SB), NOSPLIT, $0
	NO_LOCAL_POINTERS
	WORD $0x10000000  // adr x0, . $0(%rip)
	WORD $0xd65f03c0  // ret
	WORD $0x00000000; WORD $0x00000000  // .p2align 4, 0x00
	  // .p2align 2, 0x00
_validate_utf8_fast:
	WORD $0xd10043ff  // sub	sp, sp, #16
	WORD $0xa9402408  // ldp	x8, x9, [x0]
	WORD $0x8b090109  // add	x9, x8, x9
	WORD $0xd1000d2b  // sub	x11, x9, #3
	WORD $0xeb0b011f  // cmp	x8, x11
	WORD $0x54000622  // b.hs	LBB0_13 $196(%rip)
	WORD $0x52981e0c  // mov	w12, #49392
	WORD $0x72a0180c  // movk	w12, #192, lsl #16
	WORD $0x52901c0d  // mov	w13, #32992
	WORD $0x72a0100d  // movk	w13, #128, lsl #16
	WORD $0x528401ee  // mov	w14, #8207
	WORD $0x528401af  // mov	w15, #8205
	WORD $0x52981c10  // mov	w16, #49376
	WORD $0x52901811  // mov	w17, #32960
	WORD $0x52981f00  // mov	w0, #49400
	WORD $0x72b81800  // movk	w0, #49344, lsl #16
	WORD $0x528600e1  // mov	w1, #12295
	WORD $0x52901e02  // mov	w2, #33008
	WORD $0x72b01002  // movk	w2, #32896, lsl #16
	WORD $0x52860063  // mov	w3, #12291
	WORD $0xaa0803ea  // mov	x10, x8
	WORD $0x14000005  // b	LBB0_4 $20(%rip)
LBB0_2:
	WORD $0x52800025  // mov	w5, #1
LBB0_3:
	WORD $0x8b05014a  // add	x10, x10, x5
	WORD $0xeb0b015f  // cmp	x10, x11
	WORD $0x540003c2  // b.hs	LBB0_14 $120(%rip)
LBB0_4:
	WORD $0x39c00144  // ldrsb	w4, [x10]
	WORD $0x36ffff64  // tbz	w4, #31, LBB0_2 $-20(%rip)
	WORD $0xb9400144  // ldr	w4, [x10]
	WORD $0x0a0c0085  // and	w5, w4, w12
	WORD $0x6b0d00bf  // cmp	w5, w13
	WORD $0x0a0e0085  // and	w5, w4, w14
	WORD $0x7a4f00a4  // ccmp	w5, w15, #4, eq
	WORD $0x7a4018a4  // ccmp	w5, #0, #4, ne
	WORD $0x54000241  // b.ne	LBB0_12 $72(%rip)
	WORD $0x0a100085  // and	w5, w4, w16
	WORD $0x121f0c86  // and	w6, w4, #0x1e
	WORD $0x6b1100bf  // cmp	w5, w17
	WORD $0x7a4008c4  // ccmp	w6, #0, #4, eq
	WORD $0x54000161  // b.ne	LBB0_11 $44(%rip)
	WORD $0x0a000085  // and	w5, w4, w0
	WORD $0x6b0200bf  // cmp	w5, w2
	WORD $0x54000961  // b.ne	LBB0_30 $300(%rip)
	WORD $0x0a010085  // and	w5, w4, w1
	WORD $0x34000925  // cbz	w5, LBB0_30 $292(%rip)
	WORD $0x52800085  // mov	w5, #4
	WORD $0x3617fd24  // tbz	w4, #2, LBB0_3 $-92(%rip)
	WORD $0x0a030084  // and	w4, w4, w3
	WORD $0x34fffce4  // cbz	w4, LBB0_3 $-100(%rip)
	WORD $0x14000044  // b	LBB0_30 $272(%rip)
LBB0_11:
	WORD $0x52800045  // mov	w5, #2
	WORD $0x17ffffe4  // b	LBB0_3 $-112(%rip)
LBB0_12:
	WORD $0x52800065  // mov	w5, #3
	WORD $0x17ffffe2  // b	LBB0_3 $-120(%rip)
LBB0_13:
	WORD $0xaa0803ea  // mov	x10, x8
LBB0_14:
	WORD $0xeb09015f  // cmp	x10, x9
	WORD $0x54000742  // b.hs	LBB0_29 $232(%rip)
	WORD $0x52981e0b  // mov	w11, #49392
	WORD $0x72a0180b  // movk	w11, #192, lsl #16
	WORD $0x52901c0c  // mov	w12, #32992
	WORD $0x72a0100c  // movk	w12, #128, lsl #16
	WORD $0x528401ed  // mov	w13, #8207
	WORD $0x528401ae  // mov	w14, #8205
	WORD $0x52981c0f  // mov	w15, #49376
	WORD $0x52901810  // mov	w16, #32960
	WORD $0x14000004  // b	LBB0_18 $16(%rip)
LBB0_16:
	WORD $0x9100054a  // add	x10, x10, #1
LBB0_17:
	WORD $0xeb09015f  // cmp	x10, x9
	WORD $0x540005c2  // b.hs	LBB0_29 $184(%rip)
LBB0_18:
	WORD $0x39c00151  // ldrsb	w17, [x10]
	WORD $0x36ffff91  // tbz	w17, #31, LBB0_16 $-16(%rip)
	WORD $0x390033ff  // strb	wzr, [sp, #12]
	WORD $0x39002bff  // strb	wzr, [sp, #10]
	WORD $0xcb0a0122  // sub	x2, x9, x10
	WORD $0xf1000844  // subs	x4, x2, #2
	WORD $0x540001a3  // b.lo	LBB0_22 $52(%rip)
	WORD $0x39400151  // ldrb	w17, [x10]
	WORD $0x39400540  // ldrb	w0, [x10, #1]
	WORD $0x390033f1  // strb	w17, [sp, #12]
	WORD $0x91000943  // add	x3, x10, #2
	WORD $0x91002be1  // add	x1, sp, #10
	WORD $0xaa0403e2  // mov	x2, x4
	WORD $0xb4000164  // cbz	x4, LBB0_23 $44(%rip)
LBB0_21:
	WORD $0x39400071  // ldrb	w17, [x3]
	WORD $0x39000031  // strb	w17, [x1]
	WORD $0x394033f1  // ldrb	w17, [sp, #12]
	WORD $0x39402be1  // ldrb	w1, [sp, #10]
	WORD $0x14000007  // b	LBB0_24 $28(%rip)
LBB0_22:
	WORD $0x52800011  // mov	w17, #0
	WORD $0x52800000  // mov	w0, #0
	WORD $0x910033e1  // add	x1, sp, #12
	WORD $0xaa0a03e3  // mov	x3, x10
	WORD $0xb5fffee2  // cbnz	x2, LBB0_21 $-36(%rip)
LBB0_23:
	WORD $0x52800001  // mov	w1, #0
LBB0_24:
	WORD $0x53185c00  // lsl	w0, w0, #8
	WORD $0x2a014000  // orr	w0, w0, w1, lsl #16
	WORD $0x2a110000  // orr	w0, w0, w17
	WORD $0x0a0b0001  // and	w1, w0, w11
	WORD $0x6b0c003f  // cmp	w1, w12
	WORD $0x0a0d0001  // and	w1, w0, w13
	WORD $0x7a4e0024  // ccmp	w1, w14, #4, eq
	WORD $0x7a401824  // ccmp	w1, #0, #4, ne
	WORD $0x54000121  // b.ne	LBB0_28 $36(%rip)
	WORD $0x721f0e3f  // tst	w17, #0x1e
	WORD $0x540001a0  // b.eq	LBB0_30 $52(%rip)
	WORD $0x0a0f0011  // and	w17, w0, w15
	WORD $0x6b10023f  // cmp	w17, w16
	WORD $0x54000141  // b.ne	LBB0_30 $40(%rip)
	WORD $0x52800051  // mov	w17, #2
	WORD $0x8b11014a  // add	x10, x10, x17
	WORD $0x17ffffd5  // b	LBB0_17 $-172(%rip)
LBB0_28:
	WORD $0x52800071  // mov	w17, #3
	WORD $0x8b11014a  // add	x10, x10, x17
	WORD $0x17ffffd2  // b	LBB0_17 $-184(%rip)
LBB0_29:
	WORD $0xd2800000  // mov	x0, #0
	WORD $0x910043ff  // add	sp, sp, #16
	WORD $0xd65f03c0  // ret
LBB0_30:
	WORD $0xaa2a03e9  // mvn	x9, x10
	WORD $0x8b080120  // add	x0, x9, x8
	WORD $0x910043ff  // add	sp, sp, #16
	WORD $0xd65f03c0  // ret
	  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	WORD $0x00000002  // .long 2

TEXT ·__validate_utf8_fast(SB), $0-16
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $80, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_validate_utf8_fast:
	MOVD s+0(FP), R0
	CALL ·__validate_utf8_fast_entry__+16(SB)  // _validate_utf8_fast
	MOVD R0, ret+8(FP)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
