// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__skip_number_entry__(SB), NOSPLIT, $32
	NO_LOCAL_POINTERS
	WORD $0x100000a0 // adr x0, .+20
	MOVD R0, ret(FP)
	RET
	  // .p2align 4, 0x00
lCPI0_0:
	WORD $0x08040201
	WORD $0x80402010
	WORD $0x08040201
	WORD $0x80402010
	// // .byte 1
// .byte 2
// .byte 4
// .byte 8
// .byte 16
// .byte 32
// .byte 64
// .byte 128
// .byte 1
// .byte 2
// .byte 4
// .byte 8
// .byte 16
// .byte 32
// .byte 64
// .byte 128

lCPI0_1:
	WORD $0x09010800
	WORD $0x0b030a02
	WORD $0x0d050c04
	WORD $0x0f070e06
	// // .byte 0
// .byte 8
// .byte 1
// .byte 9
// .byte 2
// .byte 10
// .byte 3
// .byte 11
// .byte 4
// .byte 12
// .byte 5
// .byte 13
// .byte 6
// .byte 14
// .byte 7
// .byte 15

_skip_number:
	WORD $0xa9bd4ff4  // stp	x20, x19, [sp, #-48]!
	WORD $0xa9017bfd  // stp	fp, lr, [sp, #16]
	WORD $0xa93ffbfd  // stp	fp, lr, [sp, #-8]
	WORD $0xd10023fd  // sub	fp, sp, #8
	WORD $0xaa0003e8  // mov	x8, x0
	WORD $0xf9400020  // ldr	x0, [x1]
	WORD $0xa9402909  // ldp	x9, x10, [x8]
	WORD $0x8b000128  // add	x8, x9, x0
	WORD $0xaa0803eb  // mov	x11, x8
	WORD $0x3840156c  // ldrb	w12, [x11], #1
	WORD $0x7100b59f  // cmp	w12, #45
	WORD $0x1a9f17ed  // cset	w13, eq
	WORD $0x9a8b1108  // csel	x8, x8, x11, ne
	WORD $0xcb00014a  // sub	x10, x10, x0
	WORD $0xeb0d014e  // subs	x14, x10, x13
	WORD $0x54001ca0  // b.eq	LBB0_59 $916(%rip)
	WORD $0x3940010a  // ldrb	w10, [x8]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x3100297f  // cmn	w11, #10
	WORD $0x54001823  // b.lo	LBB0_52 $772(%rip)
	WORD $0x7100c15f  // cmp	w10, #48
	WORD $0x54000281  // b.ne	LBB0_6 $80(%rip)
	WORD $0xf10005df  // cmp	x14, #1
	WORD $0x54000101  // b.ne	LBB0_5 $32(%rip)
LBB0_4:
	WORD $0x5280002d  // mov	w13, #1
	WORD $0x8b0d0108  // add	x8, x8, x13
	WORD $0xcb090108  // sub	x8, x8, x9
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_5:
	WORD $0x3940050a  // ldrb	w10, [x8, #1]
	WORD $0x5100b94a  // sub	w10, w10, #46
	WORD $0x7100dd5f  // cmp	w10, #55
	WORD $0x5280002b  // mov	w11, #1
	WORD $0x9aca216a  // lsl	x10, x11, x10
	WORD $0xb20903eb  // mov	x11, #36028797027352576
	WORD $0xf280002b  // movk	x11, #1
	WORD $0x8a0b014a  // and	x10, x10, x11
	WORD $0xfa409944  // ccmp	x10, #0, #4, ls
	WORD $0x54fffe00  // b.eq	LBB0_4 $-64(%rip)
LBB0_6:
	WORD $0xf10041df  // cmp	x14, #16
	WORD $0x54001a03  // b.lo	LBB0_60 $832(%rip)
	WORD $0xd2800010  // mov	x16, #0
	WORD $0xd280000f  // mov	x15, #0
	WORD $0x9280000a  // mov	x10, #-1
	WORD $0x4f01e5c0  // movi.16b	v0, #46
	WORD $0x4f01e561  // movi.16b	v1, #43
	WORD $0x4f01e5a2  // movi.16b	v2, #45
	WORD $0x4f06e603  // movi.16b	v3, #208
	WORD $0x4f00e544  // movi.16b	v4, #10
Lloh0:
	WORD $0x10fff8ab  // adr	x11, lCPI0_0 $-236(%rip)
Lloh1:
	WORD $0x3dc00165  // ldr	q5, [x11, lCPI0_0@PAGEOFF] $0(%rip)
	WORD $0x4f06e7e6  // movi.16b	v6, #223
	WORD $0x4f02e4a7  // movi.16b	v7, #69
Lloh2:
	WORD $0x10fff8ab  // adr	x11, lCPI0_1 $-236(%rip)
Lloh3:
	WORD $0x3dc00170  // ldr	q16, [x11, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x12800011  // mov	w17, #-1
	WORD $0x9280000c  // mov	x12, #-1
	WORD $0x9280000b  // mov	x11, #-1
LBB0_8:
	WORD $0x3cef6911  // ldr	q17, [x8, x15]
	WORD $0x6e208e32  // cmeq.16b	v18, v17, v0
	WORD $0x6e218e33  // cmeq.16b	v19, v17, v1
	WORD $0x6e228e34  // cmeq.16b	v20, v17, v2
	WORD $0x4e238635  // add.16b	v21, v17, v3
	WORD $0x6e353495  // cmhi.16b	v21, v4, v21
	WORD $0x4e261e31  // and.16b	v17, v17, v6
	WORD $0x6e278e31  // cmeq.16b	v17, v17, v7
	WORD $0x4eb41e73  // orr.16b	v19, v19, v20
	WORD $0x4eb21eb4  // orr.16b	v20, v21, v18
	WORD $0x4eb31e35  // orr.16b	v21, v17, v19
	WORD $0x4eb51e94  // orr.16b	v20, v20, v21
	WORD $0x4e251e52  // and.16b	v18, v18, v5
	WORD $0x4e100252  // tbl.16b	v18, { v18 }, v16
	WORD $0x4e71ba52  // addv.8h	h18, v18
	WORD $0x1e260243  // fmov	w3, s18
	WORD $0x4e251e31  // and.16b	v17, v17, v5
	WORD $0x4e100231  // tbl.16b	v17, { v17 }, v16
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260224  // fmov	w4, s17
	WORD $0x4e251e71  // and.16b	v17, v19, v5
	WORD $0x4e100231  // tbl.16b	v17, { v17 }, v16
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260226  // fmov	w6, s17
	WORD $0x4e251e91  // and.16b	v17, v20, v5
	WORD $0x4e100231  // tbl.16b	v17, { v17 }, v16
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260222  // fmov	w2, s17
	WORD $0x2a2203e2  // mvn	w2, w2
	WORD $0x32103c42  // orr	w2, w2, #0xffff0000
	WORD $0x5ac00042  // rbit	w2, w2
	WORD $0x5ac01042  // clz	w2, w2
	WORD $0x1ac22225  // lsl	w5, w17, w2
	WORD $0x0a250067  // bic	w7, w3, w5
	WORD $0x0a250093  // bic	w19, w4, w5
	WORD $0x0a2500d4  // bic	w20, w6, w5
	WORD $0x7100405f  // cmp	w2, #16
	WORD $0x1a870065  // csel	w5, w3, w7, eq
	WORD $0x1a930084  // csel	w4, w4, w19, eq
	WORD $0x1a9400c3  // csel	w3, w6, w20, eq
	WORD $0x510004a6  // sub	w6, w5, #1
	WORD $0x6a0500c6  // ands	w6, w6, w5
	WORD $0x54001001  // b.ne	LBB0_55 $512(%rip)
	WORD $0x51000486  // sub	w6, w4, #1
	WORD $0x6a0400c6  // ands	w6, w6, w4
	WORD $0x54000fa1  // b.ne	LBB0_55 $500(%rip)
	WORD $0x51000466  // sub	w6, w3, #1
	WORD $0x6a0300c6  // ands	w6, w6, w3
	WORD $0x54000f41  // b.ne	LBB0_55 $488(%rip)
	WORD $0x340000c5  // cbz	w5, LBB0_14 $24(%rip)
	WORD $0x5ac000a5  // rbit	w5, w5
	WORD $0x5ac010a5  // clz	w5, w5
	WORD $0xb100057f  // cmn	x11, #1
	WORD $0x54000f41  // b.ne	LBB0_56 $488(%rip)
	WORD $0x8b0501eb  // add	x11, x15, x5
LBB0_14:
	WORD $0x340000c4  // cbz	w4, LBB0_17 $24(%rip)
	WORD $0x5ac00084  // rbit	w4, w4
	WORD $0x5ac01084  // clz	w4, w4
	WORD $0xb100059f  // cmn	x12, #1
	WORD $0x54000ee1  // b.ne	LBB0_57 $476(%rip)
	WORD $0x8b0401ec  // add	x12, x15, x4
LBB0_17:
	WORD $0x340000c3  // cbz	w3, LBB0_20 $24(%rip)
	WORD $0x5ac00063  // rbit	w3, w3
	WORD $0x5ac01063  // clz	w3, w3
	WORD $0xb100055f  // cmn	x10, #1
	WORD $0x54000e81  // b.ne	LBB0_58 $464(%rip)
	WORD $0x8b0301ea  // add	x10, x15, x3
LBB0_20:
	WORD $0x7100405f  // cmp	w2, #16
	WORD $0x54000621  // b.ne	LBB0_35 $196(%rip)
	WORD $0x910041ef  // add	x15, x15, #16
	WORD $0xd1004210  // sub	x16, x16, #16
	WORD $0x8b1001c2  // add	x2, x14, x16
	WORD $0xf1003c5f  // cmp	x2, #15
	WORD $0x54fff6e8  // b.hi	LBB0_8 $-292(%rip)
	WORD $0x8b0f0110  // add	x16, x8, x15
	WORD $0xeb0f01df  // cmp	x14, x15
	WORD $0x54000560  // b.eq	LBB0_36 $172(%rip)
LBB0_23:
	WORD $0x8b02020e  // add	x14, x16, x2
	WORD $0xaa3003ef  // mvn	x15, x16
	WORD $0x8b090011  // add	x17, x0, x9
	WORD $0x8b1101ef  // add	x15, x15, x17
	WORD $0x8b0d01ed  // add	x13, x15, x13
	WORD $0xcb08020f  // sub	x15, x16, x8
	WORD $0xaa1003f1  // mov	x17, x16
	WORD $0x14000009  // b	LBB0_26 $36(%rip)
LBB0_24:
	WORD $0xb100059f  // cmn	x12, #1
	WORD $0xaa0f03ec  // mov	x12, x15
	WORD $0x54000661  // b.ne	LBB0_46 $204(%rip)
LBB0_25:
	WORD $0xd10005ad  // sub	x13, x13, #1
	WORD $0x910005ef  // add	x15, x15, #1
	WORD $0xaa1103f0  // mov	x16, x17
	WORD $0xd1000442  // sub	x2, x2, #1
	WORD $0xb4000842  // cbz	x2, LBB0_53 $264(%rip)
LBB0_26:
	WORD $0x38401623  // ldrb	w3, [x17], #1
	WORD $0x5100c064  // sub	w4, w3, #48
	WORD $0x7100289f  // cmp	w4, #10
	WORD $0x54ffff03  // b.lo	LBB0_25 $-32(%rip)
	WORD $0x7100b47f  // cmp	w3, #45
	WORD $0x5400016d  // b.le	LBB0_32 $44(%rip)
	WORD $0x7101947f  // cmp	w3, #101
	WORD $0x54fffe20  // b.eq	LBB0_24 $-60(%rip)
	WORD $0x7101147f  // cmp	w3, #69
	WORD $0x54fffde0  // b.eq	LBB0_24 $-68(%rip)
	WORD $0x7100b87f  // cmp	w3, #46
	WORD $0x540001e1  // b.ne	LBB0_36 $60(%rip)
	WORD $0xb100057f  // cmn	x11, #1
	WORD $0xaa0f03eb  // mov	x11, x15
	WORD $0x54fffda0  // b.eq	LBB0_25 $-76(%rip)
	WORD $0x1400001e  // b	LBB0_46 $120(%rip)
LBB0_32:
	WORD $0x7100ac7f  // cmp	w3, #43
	WORD $0x54000060  // b.eq	LBB0_34 $12(%rip)
	WORD $0x7100b47f  // cmp	w3, #45
	WORD $0x540000e1  // b.ne	LBB0_36 $28(%rip)
LBB0_34:
	WORD $0xb100055f  // cmn	x10, #1
	WORD $0xaa0f03ea  // mov	x10, x15
	WORD $0x54fffca0  // b.eq	LBB0_25 $-108(%rip)
	WORD $0x14000016  // b	LBB0_46 $88(%rip)
LBB0_35:
	WORD $0x8b22410d  // add	x13, x8, w2, uxtw
	WORD $0x8b0f01b0  // add	x16, x13, x15
LBB0_36:
	WORD $0x9280000d  // mov	x13, #-1
	WORD $0xb40003cb  // cbz	x11, LBB0_51 $120(%rip)
LBB0_37:
	WORD $0xb40003aa  // cbz	x10, LBB0_51 $116(%rip)
	WORD $0xb400038c  // cbz	x12, LBB0_51 $112(%rip)
	WORD $0xcb08020d  // sub	x13, x16, x8
	WORD $0xd10005ae  // sub	x14, x13, #1
	WORD $0xeb0e017f  // cmp	x11, x14
	WORD $0x54000160  // b.eq	LBB0_45 $44(%rip)
	WORD $0xeb0e015f  // cmp	x10, x14
	WORD $0x54000120  // b.eq	LBB0_45 $36(%rip)
	WORD $0xeb0e019f  // cmp	x12, x14
	WORD $0x540000e0  // b.eq	LBB0_45 $28(%rip)
	WORD $0xf100054e  // subs	x14, x10, #1
	WORD $0x540001ab  // b.lt	LBB0_48 $52(%rip)
	WORD $0xeb0e019f  // cmp	x12, x14
	WORD $0x54000160  // b.eq	LBB0_48 $44(%rip)
	WORD $0xaa2a03ed  // mvn	x13, x10
	WORD $0x1400000e  // b	LBB0_51 $56(%rip)
LBB0_45:
	WORD $0xcb0d03ed  // neg	x13, x13
LBB0_46:
	WORD $0xb7f8018d  // tbnz	x13, #63, LBB0_51 $48(%rip)
	WORD $0x8b0d0108  // add	x8, x8, x13
	WORD $0xcb090108  // sub	x8, x8, x9
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_48:
	WORD $0xaa0c016a  // orr	x10, x11, x12
	WORD $0xb7f8020a  // tbnz	x10, #63, LBB0_54 $64(%rip)
	WORD $0xeb0c017f  // cmp	x11, x12
	WORD $0x540001cb  // b.lt	LBB0_54 $56(%rip)
	WORD $0xaa2b03ed  // mvn	x13, x11
LBB0_51:
	WORD $0xaa2d03ea  // mvn	x10, x13
	WORD $0x8b0a0108  // add	x8, x8, x10
LBB0_52:
	WORD $0x92800020  // mov	x0, #-2
	WORD $0xcb090108  // sub	x8, x8, x9
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_53:
	WORD $0xaa0e03f0  // mov	x16, x14
	WORD $0x9280000d  // mov	x13, #-1
	WORD $0xb5fffb2b  // cbnz	x11, LBB0_37 $-156(%rip)
	WORD $0x17fffff5  // b	LBB0_51 $-44(%rip)
LBB0_54:
	WORD $0xd37ffd4a  // lsr	x10, x10, #63
	WORD $0x5200014a  // eor	w10, w10, #0x1
	WORD $0xd100058e  // sub	x14, x12, #1
	WORD $0xeb0e017f  // cmp	x11, x14
	WORD $0x1a9f17eb  // cset	w11, eq
	WORD $0x6a0b015f  // tst	w10, w11
	WORD $0xda8c01ad  // csinv	x13, x13, x12, eq
	WORD $0x17ffffe1  // b	LBB0_46 $-124(%rip)
LBB0_55:
	WORD $0x5ac000ca  // rbit	w10, w6
	WORD $0x5ac0114a  // clz	w10, w10
	WORD $0xaa2f03eb  // mvn	x11, x15
	WORD $0xcb0a016d  // sub	x13, x11, x10
	WORD $0x17ffffdc  // b	LBB0_46 $-144(%rip)
LBB0_56:
	WORD $0xaa2f03ea  // mvn	x10, x15
	WORD $0xcb25414d  // sub	x13, x10, w5, uxtw
	WORD $0x17ffffd9  // b	LBB0_46 $-156(%rip)
LBB0_57:
	WORD $0xaa2f03ea  // mvn	x10, x15
	WORD $0xcb24414d  // sub	x13, x10, w4, uxtw
	WORD $0x17ffffd6  // b	LBB0_46 $-168(%rip)
LBB0_58:
	WORD $0xaa2f03ea  // mvn	x10, x15
	WORD $0xcb23414d  // sub	x13, x10, w3, uxtw
	WORD $0x17ffffd3  // b	LBB0_46 $-180(%rip)
LBB0_59:
	WORD $0x92800000  // mov	x0, #-1
	WORD $0xcb090108  // sub	x8, x8, x9
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_60:
	WORD $0x9280000b  // mov	x11, #-1
	WORD $0xaa0803f0  // mov	x16, x8
	WORD $0xaa0e03e2  // mov	x2, x14
	WORD $0x9280000c  // mov	x12, #-1
	WORD $0x9280000a  // mov	x10, #-1
	WORD $0x17ffff8a  // b	LBB0_23 $-472(%rip)
	  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	WORD $0x00000002  // .long 2

TEXT ·__skip_number(SB), NOSPLIT, $0-24
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $112, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_skip_number:
	MOVD s+0(FP), R0
	MOVD p+8(FP), R1
	MOVD ·_subr__skip_number(SB), R11
	WORD $0x1000005e // adr x30, .+8
	JMP (R11)
	MOVD R0, ret+16(FP)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
