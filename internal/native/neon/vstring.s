// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__vstring_entry__(SB), NOSPLIT, $32
	NO_LOCAL_POINTERS
	WORD $0x10000000  // adr x0, . $0(%rip)
	WORD $0x9100c3ff  // add sp, sp, #48
	WORD $0xd65f03c0  // ret
	WORD $0x00000000  // .p2align 4, 0x00
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

_vstring:
	MOVD.W R30, -48(RSP)	// 	WORD $0xf81d0ffe  // str	x30, [sp, #-48]!
	WORD $0xa9024ff4  // stp	x20, x19, [sp, #32]
	WORD $0xf9400028  // ldr	x8, [x1]
	WORD $0x37280ca3  // tbnz	w3, #5, LBB0_11 $404(%rip)
	WORD $0xf9400409  // ldr	x9, [x0, #8]
	WORD $0xeb08012c  // subs	x12, x9, x8
	WORD $0x54003dc0  // b.eq	LBB0_75 $1976(%rip)
	WORD $0xf940000a  // ldr	x10, [x0]
	WORD $0xf101019f  // cmp	x12, #64
	WORD $0x54002923  // b.lo	LBB0_37 $1316(%rip)
	WORD $0xd280000d  // mov	x13, #0
	WORD $0x92800009  // mov	x9, #-1
	WORD $0x4f01e440  // movi.16b	v0, #34
	WORD $0x4f02e781  // movi.16b	v1, #92
Lloh0:
	WORD $0x10fffd4b  // adr	x11, lCPI0_0 $-88(%rip)
Lloh1:
	WORD $0x3dc00162  // ldr	q2, [x11, lCPI0_0@PAGEOFF] $0(%rip)
Lloh2:
	WORD $0x10fffd8b  // adr	x11, lCPI0_1 $-80(%rip)
Lloh3:
	WORD $0x3dc00163  // ldr	q3, [x11, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0xaa0803ee  // mov	x14, x8
	WORD $0x9280000b  // mov	x11, #-1
LBB0_4:
	WORD $0x8b0e014f  // add	x15, x10, x14
	WORD $0xad4015e4  // ldp	q4, q5, [x15]
	WORD $0xad411de6  // ldp	q6, q7, [x15, #32]
	WORD $0x6e208c90  // cmeq.16b	v16, v4, v0
	WORD $0x6e208cb1  // cmeq.16b	v17, v5, v0
	WORD $0x6e208cd2  // cmeq.16b	v18, v6, v0
	WORD $0x6e208cf3  // cmeq.16b	v19, v7, v0
	WORD $0x6e218c84  // cmeq.16b	v4, v4, v1
	WORD $0x6e218ca5  // cmeq.16b	v5, v5, v1
	WORD $0x6e218cc6  // cmeq.16b	v6, v6, v1
	WORD $0x6e218ce7  // cmeq.16b	v7, v7, v1
	WORD $0x4e221e10  // and.16b	v16, v16, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e26020f  // fmov	w15, s16
	WORD $0x4e221e30  // and.16b	v16, v17, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260210  // fmov	w16, s16
	WORD $0x4e221e50  // and.16b	v16, v18, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260211  // fmov	w17, s16
	WORD $0x4e221e70  // and.16b	v16, v19, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260203  // fmov	w3, s16
	WORD $0x4e221c84  // and.16b	v4, v4, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260084  // fmov	w4, s4
	WORD $0x4e221ca4  // and.16b	v4, v5, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260085  // fmov	w5, s4
	WORD $0x4e221cc4  // and.16b	v4, v6, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260086  // fmov	w6, s4
	WORD $0x4e221ce4  // and.16b	v4, v7, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260087  // fmov	w7, s4
	WORD $0xd3607e31  // lsl	x17, x17, #32
	WORD $0xb3503c71  // bfi	x17, x3, #48, #16
	WORD $0x53103e10  // lsl	w16, w16, #16
	WORD $0xaa100230  // orr	x16, x17, x16
	WORD $0xaa0f020f  // orr	x15, x16, x15
	WORD $0xd3607cd0  // lsl	x16, x6, #32
	WORD $0xb3503cf0  // bfi	x16, x7, #48, #16
	WORD $0x53103cb1  // lsl	w17, w5, #16
	WORD $0xaa110210  // orr	x16, x16, x17
	WORD $0xaa040210  // orr	x16, x16, x4
	WORD $0xb5000110  // cbnz	x16, LBB0_8 $32(%rip)
	WORD $0xb50001ad  // cbnz	x13, LBB0_9 $52(%rip)
	WORD $0xb50002ef  // cbnz	x15, LBB0_10 $92(%rip)
LBB0_7:
	WORD $0xd101018c  // sub	x12, x12, #64
	WORD $0x910101ce  // add	x14, x14, #64
	WORD $0xf100fd9f  // cmp	x12, #63
	WORD $0x54fff8a8  // b.hi	LBB0_4 $-236(%rip)
	WORD $0x140000a1  // b	LBB0_24 $644(%rip)
LBB0_8:
	WORD $0xb100057f  // cmn	x11, #1
	WORD $0xdac00211  // rbit	x17, x16
	WORD $0xdac01231  // clz	x17, x17
	WORD $0x8b0e0231  // add	x17, x17, x14
	WORD $0x9a911129  // csel	x9, x9, x17, ne
	WORD $0x9a91116b  // csel	x11, x11, x17, ne
LBB0_9:
	WORD $0x8a2d0211  // bic	x17, x16, x13
	WORD $0xaa1105a3  // orr	x3, x13, x17, lsl #1
	WORD $0x8a23020d  // bic	x13, x16, x3
	WORD $0x9201f1ad  // and	x13, x13, #0xaaaaaaaaaaaaaaaa
	WORD $0xab1101b0  // adds	x16, x13, x17
	WORD $0x1a9f37ed  // cset	w13, hs
	WORD $0xd37ffa10  // lsl	x16, x16, #1
	WORD $0xd200f210  // eor	x16, x16, #0x5555555555555555
	WORD $0x8a030210  // and	x16, x16, x3
	WORD $0x8a3001ef  // bic	x15, x15, x16
	WORD $0xb4fffd6f  // cbz	x15, LBB0_7 $-84(%rip)
LBB0_10:
	WORD $0xdac001ea  // rbit	x10, x15
	WORD $0xdac0114a  // clz	x10, x10
	WORD $0x8b0e014a  // add	x10, x10, x14
	WORD $0x9100054f  // add	x15, x10, #1
	WORD $0xb6f8104f  // tbz	x15, #63, LBB0_23 $520(%rip)
	WORD $0x1400016e  // b	LBB0_69 $1464(%rip)
LBB0_11:
	WORD $0xf9400409  // ldr	x9, [x0, #8]
	WORD $0xeb08012c  // subs	x12, x9, x8
	WORD $0x54003140  // b.eq	LBB0_75 $1576(%rip)
	WORD $0xf940000a  // ldr	x10, [x0]
	WORD $0x10fff18f  // adr	x15, lCPI0_0 $-464(%rip)
	WORD $0x10fff1ee  // adr	x14, lCPI0_1 $-452(%rip)
	WORD $0xf101019f  // cmp	x12, #64
	WORD $0x54001d43  // b.lo	LBB0_38 $936(%rip)
	WORD $0xd280000b  // mov	x11, #0
	WORD $0x92800009  // mov	x9, #-1
	WORD $0x4f01e440  // movi.16b	v0, #34
	WORD $0x3dc001e1  // ldr	q1, [x15, lCPI0_0@PAGEOFF] $0(%rip)
	WORD $0x3dc001c2  // ldr	q2, [x14, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x4f02e783  // movi.16b	v3, #92
	WORD $0x4f01e404  // movi.16b	v4, #32
	WORD $0xaa0803ed  // mov	x13, x8
LBB0_14:
	WORD $0x8b0d0150  // add	x16, x10, x13
	WORD $0xad401e10  // ldp	q16, q7, [x16]
	WORD $0xad411606  // ldp	q6, q5, [x16, #32]
	WORD $0x6e208e11  // cmeq.16b	v17, v16, v0
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260230  // fmov	w16, s17
	WORD $0x6e208cf1  // cmeq.16b	v17, v7, v0
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260231  // fmov	w17, s17
	WORD $0x6e208cd1  // cmeq.16b	v17, v6, v0
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260223  // fmov	w3, s17
	WORD $0x6e208cb1  // cmeq.16b	v17, v5, v0
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260224  // fmov	w4, s17
	WORD $0x6e238e11  // cmeq.16b	v17, v16, v3
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260225  // fmov	w5, s17
	WORD $0x6e238cf1  // cmeq.16b	v17, v7, v3
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260226  // fmov	w6, s17
	WORD $0x6e238cd1  // cmeq.16b	v17, v6, v3
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260227  // fmov	w7, s17
	WORD $0x6e238cb1  // cmeq.16b	v17, v5, v3
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260233  // fmov	w19, s17
	WORD $0xd3607c63  // lsl	x3, x3, #32
	WORD $0xb3503c83  // bfi	x3, x4, #48, #16
	WORD $0x53103e31  // lsl	w17, w17, #16
	WORD $0xaa110071  // orr	x17, x3, x17
	WORD $0xaa100230  // orr	x16, x17, x16
	WORD $0xd3607cf1  // lsl	x17, x7, #32
	WORD $0xb3503e71  // bfi	x17, x19, #48, #16
	WORD $0x53103cc3  // lsl	w3, w6, #16
	WORD $0xaa030231  // orr	x17, x17, x3
	WORD $0xaa050231  // orr	x17, x17, x5
	WORD $0xb5000451  // cbnz	x17, LBB0_19 $136(%rip)
	WORD $0xb50004cb  // cbnz	x11, LBB0_20 $152(%rip)
LBB0_16:
	WORD $0x6e303490  // cmhi.16b	v16, v4, v16
	WORD $0x4e211e10  // and.16b	v16, v16, v1
	WORD $0x4e020210  // tbl.16b	v16, { v16 }, v2
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260211  // fmov	w17, s16
	WORD $0x6e273487  // cmhi.16b	v7, v4, v7
	WORD $0x4e211ce7  // and.16b	v7, v7, v1
	WORD $0x4e0200e7  // tbl.16b	v7, { v7 }, v2
	WORD $0x4e71b8e7  // addv.8h	h7, v7
	WORD $0x1e2600e3  // fmov	w3, s7
	WORD $0x6e263486  // cmhi.16b	v6, v4, v6
	WORD $0x4e211cc6  // and.16b	v6, v6, v1
	WORD $0x4e0200c6  // tbl.16b	v6, { v6 }, v2
	WORD $0x4e71b8c6  // addv.8h	h6, v6
	WORD $0x1e2600c4  // fmov	w4, s6
	WORD $0x6e253485  // cmhi.16b	v5, v4, v5
	WORD $0x4e211ca5  // and.16b	v5, v5, v1
	WORD $0x4e0200a5  // tbl.16b	v5, { v5 }, v2
	WORD $0x4e71b8a5  // addv.8h	h5, v5
	WORD $0x1e2600a5  // fmov	w5, s5
	WORD $0xd3607c84  // lsl	x4, x4, #32
	WORD $0xb3503ca4  // bfi	x4, x5, #48, #16
	WORD $0x53103c63  // lsl	w3, w3, #16
	WORD $0xaa030083  // orr	x3, x4, x3
	WORD $0xaa110071  // orr	x17, x3, x17
	WORD $0xb50002f0  // cbnz	x16, LBB0_21 $92(%rip)
	WORD $0xb50009f1  // cbnz	x17, LBB0_29 $316(%rip)
	WORD $0xd101018c  // sub	x12, x12, #64
	WORD $0x910101ad  // add	x13, x13, #64
	WORD $0xf100fd9f  // cmp	x12, #63
	WORD $0x54fff568  // b.hi	LBB0_14 $-340(%rip)
	WORD $0x1400004c  // b	LBB0_30 $304(%rip)
LBB0_19:
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0xdac00223  // rbit	x3, x17
	WORD $0xdac01063  // clz	x3, x3
	WORD $0x8b0d0063  // add	x3, x3, x13
	WORD $0x9a831129  // csel	x9, x9, x3, ne
LBB0_20:
	WORD $0x8a2b0223  // bic	x3, x17, x11
	WORD $0xaa030564  // orr	x4, x11, x3, lsl #1
	WORD $0x8a24022b  // bic	x11, x17, x4
	WORD $0x9201f16b  // and	x11, x11, #0xaaaaaaaaaaaaaaaa
	WORD $0xab030171  // adds	x17, x11, x3
	WORD $0x1a9f37eb  // cset	w11, hs
	WORD $0xd37ffa31  // lsl	x17, x17, #1
	WORD $0xd200f231  // eor	x17, x17, #0x5555555555555555
	WORD $0x8a040231  // and	x17, x17, x4
	WORD $0x8a310210  // bic	x16, x16, x17
	WORD $0x17ffffd1  // b	LBB0_16 $-188(%rip)
LBB0_21:
	WORD $0xdac0020a  // rbit	x10, x16
	WORD $0xdac0114a  // clz	x10, x10
	WORD $0xdac0022b  // rbit	x11, x17
	WORD $0xdac0116b  // clz	x11, x11
	WORD $0xeb0a017f  // cmp	x11, x10
	WORD $0x54001ee3  // b.lo	LBB0_70 $988(%rip)
	WORD $0x8b0d014a  // add	x10, x10, x13
	WORD $0x9100054f  // add	x15, x10, #1
	WORD $0xb7f81dcf  // tbnz	x15, #63, LBB0_69 $952(%rip)
LBB0_23:
	WORD $0xf900002f  // str	x15, [x1]
	WORD $0x528000ea  // mov	w10, #7
	WORD $0xf900004a  // str	x10, [x2]
	WORD $0xeb0f013f  // cmp	x9, x15
	WORD $0xda9fb129  // csinv	x9, x9, xzr, lt
	WORD $0xa9012448  // stp	x8, x9, [x2, #16]
	WORD $0xa9424ff4  // ldp	x20, x19, [sp, #32]
	WORD $0x9100c3ff  // add	sp, sp, #48
	WORD $0xd65f03c0  // ret
LBB0_24:
	WORD $0x8b0e014e  // add	x14, x10, x14
	WORD $0xf100818f  // subs	x15, x12, #32
	WORD $0x54001063  // b.lo	LBB0_42 $524(%rip)
LBB0_25:
	WORD $0xad4005c0  // ldp	q0, q1, [x14]
	WORD $0x4f01e442  // movi.16b	v2, #34
	WORD $0x6e228c03  // cmeq.16b	v3, v0, v2
	WORD $0x6e228c22  // cmeq.16b	v2, v1, v2
	WORD $0x4f02e784  // movi.16b	v4, #92
	WORD $0x6e248c00  // cmeq.16b	v0, v0, v4
	WORD $0x6e248c21  // cmeq.16b	v1, v1, v4
Lloh4:
	WORD $0x10ffdfac  // adr	x12, lCPI0_0 $-1036(%rip)
Lloh5:
	WORD $0x3dc00184  // ldr	q4, [x12, lCPI0_0@PAGEOFF] $0(%rip)
	WORD $0x4e241c63  // and.16b	v3, v3, v4
Lloh6:
	WORD $0x10ffdfcc  // adr	x12, lCPI0_1 $-1032(%rip)
Lloh7:
	WORD $0x3dc00185  // ldr	q5, [x12, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x4e050063  // tbl.16b	v3, { v3 }, v5
	WORD $0x4e71b863  // addv.8h	h3, v3
	WORD $0x1e26006c  // fmov	w12, s3
	WORD $0x4e241c42  // and.16b	v2, v2, v4
	WORD $0x4e050042  // tbl.16b	v2, { v2 }, v5
	WORD $0x4e71b842  // addv.8h	h2, v2
	WORD $0x1e260051  // fmov	w17, s2
	WORD $0x4e241c00  // and.16b	v0, v0, v4
	WORD $0x4e050000  // tbl.16b	v0, { v0 }, v5
	WORD $0x4e71b800  // addv.8h	h0, v0
	WORD $0x1e260010  // fmov	w16, s0
	WORD $0x4e241c20  // and.16b	v0, v1, v4
	WORD $0x4e050000  // tbl.16b	v0, { v0 }, v5
	WORD $0x4e71b800  // addv.8h	h0, v0
	WORD $0x1e260003  // fmov	w3, s0
	WORD $0x33103e2c  // bfi	w12, w17, #16, #16
	WORD $0x33103c70  // bfi	w16, w3, #16, #16
	WORD $0x350009d0  // cbnz	w16, LBB0_39 $312(%rip)
	WORD $0xb5000a8d  // cbnz	x13, LBB0_40 $336(%rip)
	WORD $0xb4000c2c  // cbz	x12, LBB0_41 $388(%rip)
LBB0_28:
	WORD $0xdac0018b  // rbit	x11, x12
	WORD $0xdac0116b  // clz	x11, x11
	WORD $0xcb0a01ca  // sub	x10, x14, x10
	WORD $0x14000037  // b	LBB0_36 $220(%rip)
LBB0_29:
	WORD $0x9280002f  // mov	x15, #-2
	WORD $0x140000bc  // b	LBB0_69 $752(%rip)
LBB0_30:
	WORD $0x8b0d014d  // add	x13, x10, x13
	WORD $0xf1008190  // subs	x16, x12, #32
	WORD $0x54001303  // b.lo	LBB0_59 $608(%rip)
LBB0_31:
	WORD $0xad4005a0  // ldp	q0, q1, [x13]
	WORD $0x4f01e442  // movi.16b	v2, #34
	WORD $0x6e228c03  // cmeq.16b	v3, v0, v2
	WORD $0x3dc001e4  // ldr	q4, [x15, lCPI0_0@PAGEOFF] $0(%rip)
	WORD $0x4e241c63  // and.16b	v3, v3, v4
	WORD $0x3dc001c5  // ldr	q5, [x14, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x4e050063  // tbl.16b	v3, { v3 }, v5
	WORD $0x4e71b863  // addv.8h	h3, v3
	WORD $0x1e26006c  // fmov	w12, s3
	WORD $0x6e228c22  // cmeq.16b	v2, v1, v2
	WORD $0x4e241c42  // and.16b	v2, v2, v4
	WORD $0x4e050042  // tbl.16b	v2, { v2 }, v5
	WORD $0x4e71b842  // addv.8h	h2, v2
	WORD $0x1e260051  // fmov	w17, s2
	WORD $0x4f02e782  // movi.16b	v2, #92
	WORD $0x6e228c03  // cmeq.16b	v3, v0, v2
	WORD $0x4e241c63  // and.16b	v3, v3, v4
	WORD $0x4e050063  // tbl.16b	v3, { v3 }, v5
	WORD $0x4e71b863  // addv.8h	h3, v3
	WORD $0x1e26006f  // fmov	w15, s3
	WORD $0x6e228c22  // cmeq.16b	v2, v1, v2
	WORD $0x4e241c42  // and.16b	v2, v2, v4
	WORD $0x4e050042  // tbl.16b	v2, { v2 }, v5
	WORD $0x4e71b842  // addv.8h	h2, v2
	WORD $0x1e260043  // fmov	w3, s2
	WORD $0x4f01e402  // movi.16b	v2, #32
	WORD $0x6e203440  // cmhi.16b	v0, v2, v0
	WORD $0x4e241c00  // and.16b	v0, v0, v4
	WORD $0x4e050000  // tbl.16b	v0, { v0 }, v5
	WORD $0x4e71b800  // addv.8h	h0, v0
	WORD $0x1e26000e  // fmov	w14, s0
	WORD $0x6e213440  // cmhi.16b	v0, v2, v1
	WORD $0x4e241c00  // and.16b	v0, v0, v4
	WORD $0x4e050000  // tbl.16b	v0, { v0 }, v5
	WORD $0x4e71b800  // addv.8h	h0, v0
	WORD $0x1e260004  // fmov	w4, s0
	WORD $0x33103e2c  // bfi	w12, w17, #16, #16
	WORD $0x33103c6f  // bfi	w15, w3, #16, #16
	WORD $0x33103c8e  // bfi	w14, w4, #16, #16
	WORD $0x35000b2f  // cbnz	w15, LBB0_55 $356(%rip)
	WORD $0xb5000bcb  // cbnz	x11, LBB0_56 $376(%rip)
	WORD $0xb4000d6c  // cbz	x12, LBB0_57 $428(%rip)
LBB0_34:
	WORD $0xdac0018b  // rbit	x11, x12
	WORD $0xdac0116b  // clz	x11, x11
	WORD $0xdac001cc  // rbit	x12, x14
	WORD $0xdac0118c  // clz	x12, x12
	WORD $0xeb0b019f  // cmp	x12, x11
	WORD $0x540011e3  // b.lo	LBB0_70 $572(%rip)
	WORD $0xcb0a01aa  // sub	x10, x13, x10
LBB0_36:
	WORD $0x8b0b014a  // add	x10, x10, x11
	WORD $0x9100054f  // add	x15, x10, #1
	WORD $0xb6fff30f  // tbz	x15, #63, LBB0_23 $-416(%rip)
	WORD $0x14000084  // b	LBB0_69 $528(%rip)
LBB0_37:
	WORD $0xd280000d  // mov	x13, #0
	WORD $0x8b08014e  // add	x14, x10, x8
	WORD $0x92800009  // mov	x9, #-1
	WORD $0x9280000b  // mov	x11, #-1
	WORD $0xf100818f  // subs	x15, x12, #32
	WORD $0x54fff3a2  // b.hs	LBB0_25 $-396(%rip)
	WORD $0x1400001e  // b	LBB0_42 $120(%rip)
LBB0_38:
	WORD $0xd280000b  // mov	x11, #0
	WORD $0x8b08014d  // add	x13, x10, x8
	WORD $0x92800009  // mov	x9, #-1
	WORD $0xf1008190  // subs	x16, x12, #32
	WORD $0x54fff802  // b.hs	LBB0_31 $-256(%rip)
	WORD $0x14000056  // b	LBB0_59 $344(%rip)
LBB0_39:
	WORD $0xdac00211  // rbit	x17, x16
	WORD $0xdac01231  // clz	x17, x17
	WORD $0xcb0a01c3  // sub	x3, x14, x10
	WORD $0x8b030231  // add	x17, x17, x3
	WORD $0xb100057f  // cmn	x11, #1
	WORD $0x9a911129  // csel	x9, x9, x17, ne
	WORD $0x9a91116b  // csel	x11, x11, x17, ne
LBB0_40:
	WORD $0x0a2d0211  // bic	w17, w16, w13
	WORD $0x531f7a23  // lsl	w3, w17, #1
	WORD $0x331f7a2d  // bfi	w13, w17, #1, #31
	WORD $0x0a230210  // bic	w16, w16, w3
	WORD $0x1201f210  // and	w16, w16, #0xaaaaaaaa
	WORD $0x2b110210  // adds	w16, w16, w17
	WORD $0x3200f3f1  // mov	w17, #1431655765
	WORD $0x4a100630  // eor	w16, w17, w16, lsl #1
	WORD $0x0a0d020d  // and	w13, w16, w13
	WORD $0x1a9f37f0  // cset	w16, hs
	WORD $0x2a2d03ed  // mvn	w13, w13
	WORD $0x8a0c01ac  // and	x12, x13, x12
	WORD $0xaa1003ed  // mov	x13, x16
	WORD $0xb5fff42c  // cbnz	x12, LBB0_28 $-380(%rip)
LBB0_41:
	WORD $0x910081ce  // add	x14, x14, #32
	WORD $0xaa0f03ec  // mov	x12, x15
LBB0_42:
	WORD $0xb5000d8d  // cbnz	x13, LBB0_71 $432(%rip)
	WORD $0xb40003ec  // cbz	x12, LBB0_52 $124(%rip)
LBB0_44:
	WORD $0xcb0a03ed  // neg	x13, x10
LBB0_45:
	WORD $0xd2800010  // mov	x16, #0
LBB0_46:
	WORD $0x387069cf  // ldrb	w15, [x14, x16]
	WORD $0x710089ff  // cmp	w15, #34
	WORD $0x54000300  // b.eq	LBB0_51 $96(%rip)
	WORD $0x710171ff  // cmp	w15, #92
	WORD $0x540000a0  // b.eq	LBB0_49 $20(%rip)
	WORD $0x91000610  // add	x16, x16, #1
	WORD $0xeb10019f  // cmp	x12, x16
	WORD $0x54ffff21  // b.ne	LBB0_46 $-28(%rip)
	WORD $0x14000017  // b	LBB0_53 $92(%rip)
LBB0_49:
	WORD $0xd100058f  // sub	x15, x12, #1
	WORD $0xeb1001ff  // cmp	x15, x16
	WORD $0x54000de0  // b.eq	LBB0_75 $444(%rip)
	WORD $0x8b0e01af  // add	x15, x13, x14
	WORD $0x8b1001ef  // add	x15, x15, x16
	WORD $0xb100057f  // cmn	x11, #1
	WORD $0x9a8901e9  // csel	x9, x15, x9, eq
	WORD $0x9a8b01eb  // csel	x11, x15, x11, eq
	WORD $0x8b1001ce  // add	x14, x14, x16
	WORD $0x910009ce  // add	x14, x14, #2
	WORD $0xcb10018f  // sub	x15, x12, x16
	WORD $0xd1000991  // sub	x17, x12, #2
	WORD $0xd10009ec  // sub	x12, x15, #2
	WORD $0x9280000f  // mov	x15, #-1
	WORD $0xeb10023f  // cmp	x17, x16
	WORD $0x54fffce1  // b.ne	LBB0_45 $-100(%rip)
	WORD $0x14000042  // b	LBB0_69 $264(%rip)
LBB0_51:
	WORD $0x8b1001cb  // add	x11, x14, x16
	WORD $0x9100056e  // add	x14, x11, #1
LBB0_52:
	WORD $0xcb0a01cf  // sub	x15, x14, x10
	WORD $0xb6ffea2f  // tbz	x15, #63, LBB0_23 $-700(%rip)
	WORD $0x1400003d  // b	LBB0_69 $244(%rip)
LBB0_53:
	WORD $0x710089ff  // cmp	w15, #34
	WORD $0x54000b41  // b.ne	LBB0_75 $360(%rip)
	WORD $0x8b0c01ce  // add	x14, x14, x12
	WORD $0x17fffffa  // b	LBB0_52 $-24(%rip)
LBB0_55:
	WORD $0xdac001f1  // rbit	x17, x15
	WORD $0xdac01231  // clz	x17, x17
	WORD $0xcb0a01a3  // sub	x3, x13, x10
	WORD $0x8b110071  // add	x17, x3, x17
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a911129  // csel	x9, x9, x17, ne
LBB0_56:
	WORD $0x0a2b01f1  // bic	w17, w15, w11
	WORD $0x531f7a23  // lsl	w3, w17, #1
	WORD $0x331f7a2b  // bfi	w11, w17, #1, #31
	WORD $0x0a2301ef  // bic	w15, w15, w3
	WORD $0x1201f1ef  // and	w15, w15, #0xaaaaaaaa
	WORD $0x2b1101ef  // adds	w15, w15, w17
	WORD $0x3200f3f1  // mov	w17, #1431655765
	WORD $0x4a0f062f  // eor	w15, w17, w15, lsl #1
	WORD $0x0a0b01eb  // and	w11, w15, w11
	WORD $0x1a9f37ef  // cset	w15, hs
	WORD $0x2a2b03eb  // mvn	w11, w11
	WORD $0x8a0c016c  // and	x12, x11, x12
	WORD $0xaa0f03eb  // mov	x11, x15
	WORD $0xb5fff2ec  // cbnz	x12, LBB0_34 $-420(%rip)
LBB0_57:
	WORD $0x3500054e  // cbnz	w14, LBB0_70 $168(%rip)
	WORD $0x910081ad  // add	x13, x13, #32
	WORD $0xaa1003ec  // mov	x12, x16
LBB0_59:
	WORD $0xb500070b  // cbnz	x11, LBB0_73 $224(%rip)
	WORD $0xb40007ec  // cbz	x12, LBB0_75 $252(%rip)
LBB0_61:
	WORD $0x9280002b  // mov	x11, #-2
	WORD $0x5280004e  // mov	w14, #2
LBB0_62:
	WORD $0x394001af  // ldrb	w15, [x13]
	WORD $0x710089ff  // cmp	w15, #34
	WORD $0x54000300  // b.eq	LBB0_68 $96(%rip)
	WORD $0x710171ff  // cmp	w15, #92
	WORD $0x54000140  // b.eq	LBB0_66 $40(%rip)
	WORD $0x710081ff  // cmp	w15, #32
	WORD $0x540003a3  // b.lo	LBB0_70 $116(%rip)
	WORD $0x92800010  // mov	x16, #-1
	WORD $0x5280002f  // mov	w15, #1
	WORD $0x8b0f01ad  // add	x13, x13, x15
	WORD $0x9280000f  // mov	x15, #-1
	WORD $0xab0c020c  // adds	x12, x16, x12
	WORD $0x54fffe81  // b.ne	LBB0_62 $-48(%rip)
	WORD $0x14000010  // b	LBB0_69 $64(%rip)
LBB0_66:
	WORD $0xf100059f  // cmp	x12, #1
	WORD $0x540005a0  // b.eq	LBB0_75 $180(%rip)
	WORD $0xcb0a01af  // sub	x15, x13, x10
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a8f1129  // csel	x9, x9, x15, ne
	WORD $0x9a8e11cf  // csel	x15, x14, x14, ne
	WORD $0x9a8b1170  // csel	x16, x11, x11, ne
	WORD $0x8b0f01ad  // add	x13, x13, x15
	WORD $0x9280000f  // mov	x15, #-1
	WORD $0xab0c020c  // adds	x12, x16, x12
	WORD $0x54fffd01  // b.ne	LBB0_62 $-96(%rip)
	WORD $0x14000004  // b	LBB0_69 $16(%rip)
LBB0_68:
	WORD $0xcb0a01aa  // sub	x10, x13, x10
	WORD $0x9100054f  // add	x15, x10, #1
	WORD $0xb6ffe28f  // tbz	x15, #63, LBB0_23 $-944(%rip)
LBB0_69:
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004f  // str	x15, [x2]
	WORD $0xa9424ff4  // ldp	x20, x19, [sp, #32]
	WORD $0x9100c3ff  // add	sp, sp, #48
	WORD $0xd65f03c0  // ret
LBB0_70:
	WORD $0x9280002f  // mov	x15, #-2
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004f  // str	x15, [x2]
	WORD $0xa9424ff4  // ldp	x20, x19, [sp, #32]
	WORD $0x9100c3ff  // add	sp, sp, #48
	WORD $0xd65f03c0  // ret
LBB0_71:
	WORD $0xb400024c  // cbz	x12, LBB0_75 $72(%rip)
	WORD $0xaa2a03ed  // mvn	x13, x10
	WORD $0x8b0d01cd  // add	x13, x14, x13
	WORD $0xb100057f  // cmn	x11, #1
	WORD $0x9a8901a9  // csel	x9, x13, x9, eq
	WORD $0x9a8b01ab  // csel	x11, x13, x11, eq
	WORD $0x910005ce  // add	x14, x14, #1
	WORD $0xd100058c  // sub	x12, x12, #1
	WORD $0xb5fff1cc  // cbnz	x12, LBB0_44 $-456(%rip)
	WORD $0x17ffffab  // b	LBB0_52 $-340(%rip)
LBB0_73:
	WORD $0xb400010c  // cbz	x12, LBB0_75 $32(%rip)
	WORD $0xaa2a03eb  // mvn	x11, x10
	WORD $0x8b0b01ab  // add	x11, x13, x11
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a890169  // csel	x9, x11, x9, eq
	WORD $0x910005ad  // add	x13, x13, #1
	WORD $0xd100058c  // sub	x12, x12, #1
	WORD $0xb5fff86c  // cbnz	x12, LBB0_61 $-244(%rip)
LBB0_75:
	WORD $0x9280000f  // mov	x15, #-1
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004f  // str	x15, [x2]
	WORD $0xa9424ff4  // ldp	x20, x19, [sp, #32]
	WORD $0x9100c3ff  // add	sp, sp, #48
	WORD $0xd65f03c0  // ret
	  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	WORD $0x00000002  // .long 2

TEXT ·__vstring(SB), $0-32
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $80, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_vstring:
	MOVD s+0(FP), R0
	MOVD p+8(FP), R1
	MOVD v+16(FP), R2
	MOVD flags+24(FP), R3
	CALL ·__vstring_entry__+60(SB)  // _vstring
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
