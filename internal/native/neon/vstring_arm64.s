// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__vstring_entry__(SB), NOSPLIT, $32
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

_vstring:
	WORD $0xa9bd4ff4  // stp	x20, x19, [sp, #-48]!
	WORD $0xa9017bfd  // stp	fp, lr, [sp, #16]
	WORD $0xa93ffbfd  // stp	fp, lr, [sp, #-8]
	WORD $0xd10023fd  // sub	fp, sp, #8
	WORD $0xf9400028  // ldr	x8, [x1]
	WORD $0xf9400409  // ldr	x9, [x0, #8]
	WORD $0x37280c43  // tbnz	w3, #5, LBB0_11 $392(%rip)
	WORD $0xeb08012b  // subs	x11, x9, x8
	WORD $0x54003d00  // b.eq	LBB0_77 $1952(%rip)
	WORD $0xf940000a  // ldr	x10, [x0]
	WORD $0xf101017f  // cmp	x11, #64
	WORD $0x54001e03  // b.lo	LBB0_27 $960(%rip)
	WORD $0xd280000c  // mov	x12, #0
	WORD $0x92800009  // mov	x9, #-1
	WORD $0x4f01e440  // movi.16b	v0, #34
	WORD $0x4f02e781  // movi.16b	v1, #92
Lloh0:
	WORD $0x10fffd0d  // adr	x13, lCPI0_0 $-96(%rip)
Lloh1:
	WORD $0x3dc001a2  // ldr	q2, [x13, lCPI0_0@PAGEOFF] $0(%rip)
Lloh2:
	WORD $0x10fffd4d  // adr	x13, lCPI0_1 $-88(%rip)
Lloh3:
	WORD $0x3dc001a3  // ldr	q3, [x13, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0xaa0803ed  // mov	x13, x8
LBB0_4:
	WORD $0x8b0d014e  // add	x14, x10, x13
	WORD $0xad4015c4  // ldp	q4, q5, [x14]
	WORD $0xad411dc6  // ldp	q6, q7, [x14, #32]
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
	WORD $0x1e26020e  // fmov	w14, s16
	WORD $0x4e221e30  // and.16b	v16, v17, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e26020f  // fmov	w15, s16
	WORD $0x4e221e50  // and.16b	v16, v18, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260210  // fmov	w16, s16
	WORD $0x4e221e70  // and.16b	v16, v19, v2
	WORD $0x4e030210  // tbl.16b	v16, { v16 }, v3
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260211  // fmov	w17, s16
	WORD $0x4e221c84  // and.16b	v4, v4, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260083  // fmov	w3, s4
	WORD $0x4e221ca4  // and.16b	v4, v5, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260084  // fmov	w4, s4
	WORD $0x4e221cc4  // and.16b	v4, v6, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260085  // fmov	w5, s4
	WORD $0x4e221ce4  // and.16b	v4, v7, v2
	WORD $0x4e030084  // tbl.16b	v4, { v4 }, v3
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260086  // fmov	w6, s4
	WORD $0xd3607e10  // lsl	x16, x16, #32
	WORD $0xaa11c210  // orr	x16, x16, x17, lsl #48
	WORD $0x53103def  // lsl	w15, w15, #16
	WORD $0xaa0f020f  // orr	x15, x16, x15
	WORD $0xaa0e01ee  // orr	x14, x15, x14
	WORD $0xd3607caf  // lsl	x15, x5, #32
	WORD $0xaa06c1ef  // orr	x15, x15, x6, lsl #48
	WORD $0x53103c90  // lsl	w16, w4, #16
	WORD $0xaa1001ef  // orr	x15, x15, x16
	WORD $0xaa0301ef  // orr	x15, x15, x3
	WORD $0xb500010f  // cbnz	x15, LBB0_8 $32(%rip)
	WORD $0xb500018c  // cbnz	x12, LBB0_9 $48(%rip)
	WORD $0xb50002ce  // cbnz	x14, LBB0_10 $88(%rip)
LBB0_7:
	WORD $0xd101016b  // sub	x11, x11, #64
	WORD $0x910101ad  // add	x13, x13, #64
	WORD $0xf100fd7f  // cmp	x11, #63
	WORD $0x54fff8a8  // b.hi	LBB0_4 $-236(%rip)
	WORD $0x1400009f  // b	LBB0_24 $636(%rip)
LBB0_8:
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0xdac001f0  // rbit	x16, x15
	WORD $0xdac01210  // clz	x16, x16
	WORD $0x8b0d0210  // add	x16, x16, x13
	WORD $0x9a901129  // csel	x9, x9, x16, ne
LBB0_9:
	WORD $0x8a2c01f0  // bic	x16, x15, x12
	WORD $0xaa100591  // orr	x17, x12, x16, lsl #1
	WORD $0x8a3101ec  // bic	x12, x15, x17
	WORD $0x9201f18c  // and	x12, x12, #0xaaaaaaaaaaaaaaaa
	WORD $0xab10018f  // adds	x15, x12, x16
	WORD $0x1a9f37ec  // cset	w12, hs
	WORD $0xd37ff9ef  // lsl	x15, x15, #1
	WORD $0xd200f1ef  // eor	x15, x15, #0x5555555555555555
	WORD $0x8a1101ef  // and	x15, x15, x17
	WORD $0x8a2f01ce  // bic	x14, x14, x15
	WORD $0xb4fffd8e  // cbz	x14, LBB0_7 $-80(%rip)
LBB0_10:
	WORD $0xdac001ca  // rbit	x10, x14
	WORD $0xdac0114a  // clz	x10, x10
	WORD $0x8b0d014a  // add	x10, x10, x13
	WORD $0x9100054e  // add	x14, x10, #1
	WORD $0xb6f8102e  // tbz	x14, #63, LBB0_23 $516(%rip)
	WORD $0x1400016b  // b	LBB0_71 $1452(%rip)
LBB0_11:
	WORD $0xeb08012c  // subs	x12, x9, x8
	WORD $0x540030e0  // b.eq	LBB0_77 $1564(%rip)
	WORD $0xf940000a  // ldr	x10, [x0]
	WORD $0x10fff1af  // adr	x15, lCPI0_0 $-460(%rip)
	WORD $0x10fff20e  // adr	x14, lCPI0_1 $-448(%rip)
	WORD $0xf101019f  // cmp	x12, #64
	WORD $0x540016a3  // b.lo	LBB0_33 $724(%rip)
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
	WORD $0xaa04c063  // orr	x3, x3, x4, lsl #48
	WORD $0x53103e31  // lsl	w17, w17, #16
	WORD $0xaa110071  // orr	x17, x3, x17
	WORD $0xaa100230  // orr	x16, x17, x16
	WORD $0xd3607cf1  // lsl	x17, x7, #32
	WORD $0xaa13c231  // orr	x17, x17, x19, lsl #48
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
	WORD $0xaa05c084  // orr	x4, x4, x5, lsl #48
	WORD $0x53103c63  // lsl	w3, w3, #16
	WORD $0xaa030083  // orr	x3, x4, x3
	WORD $0xaa110071  // orr	x17, x3, x17
	WORD $0xb50002f0  // cbnz	x16, LBB0_21 $92(%rip)
	WORD $0xb5000551  // cbnz	x17, LBB0_25 $168(%rip)
	WORD $0xd101018c  // sub	x12, x12, #64
	WORD $0x910101ad  // add	x13, x13, #64
	WORD $0xf100fd9f  // cmp	x12, #63
	WORD $0x54fff568  // b.hi	LBB0_14 $-340(%rip)
	WORD $0x1400002c  // b	LBB0_26 $176(%rip)
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
	WORD $0x54001ea3  // b.lo	LBB0_72 $980(%rip)
	WORD $0x8b0d014a  // add	x10, x10, x13
	WORD $0x9100054e  // add	x14, x10, #1
	WORD $0xb7f81d8e  // tbnz	x14, #63, LBB0_71 $944(%rip)
LBB0_23:
	WORD $0xf900002e  // str	x14, [x1]
	WORD $0x528000ea  // mov	w10, #7
	WORD $0xf900004a  // str	x10, [x2]
	WORD $0xeb0e013f  // cmp	x9, x14
	WORD $0xda9fb129  // csinv	x9, x9, xzr, lt
	WORD $0xa9012448  // stp	x8, x9, [x2, #16]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_24:
	WORD $0x8b0d014d  // add	x13, x10, x13
	WORD $0x1400000d  // b	LBB0_28 $52(%rip)
LBB0_25:
	WORD $0x9280002e  // mov	x14, #-2
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004e  // str	x14, [x2]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_26:
	WORD $0x8b0d014d  // add	x13, x10, x13
	WORD $0x1400002c  // b	LBB0_34 $176(%rip)
LBB0_27:
	WORD $0xd280000c  // mov	x12, #0
	WORD $0x8b08014d  // add	x13, x10, x8
	WORD $0x92800009  // mov	x9, #-1
LBB0_28:
	WORD $0xf100816e  // subs	x14, x11, #32
	WORD $0x54000e83  // b.lo	LBB0_43 $464(%rip)
	WORD $0xad4005a0  // ldp	q0, q1, [x13]
	WORD $0x4f01e442  // movi.16b	v2, #34
	WORD $0x6e228c03  // cmeq.16b	v3, v0, v2
	WORD $0x6e228c22  // cmeq.16b	v2, v1, v2
	WORD $0x4f02e784  // movi.16b	v4, #92
	WORD $0x6e248c00  // cmeq.16b	v0, v0, v4
	WORD $0x6e248c21  // cmeq.16b	v1, v1, v4
Lloh4:
	WORD $0x10ffde2b  // adr	x11, lCPI0_0 $-1084(%rip)
Lloh5:
	WORD $0x3dc00164  // ldr	q4, [x11, lCPI0_0@PAGEOFF] $0(%rip)
	WORD $0x4e241c63  // and.16b	v3, v3, v4
Lloh6:
	WORD $0x10ffde4b  // adr	x11, lCPI0_1 $-1080(%rip)
Lloh7:
	WORD $0x3dc00165  // ldr	q5, [x11, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x4e050063  // tbl.16b	v3, { v3 }, v5
	WORD $0x4e71b863  // addv.8h	h3, v3
	WORD $0x1e26006b  // fmov	w11, s3
	WORD $0x4e241c42  // and.16b	v2, v2, v4
	WORD $0x4e050042  // tbl.16b	v2, { v2 }, v5
	WORD $0x4e71b842  // addv.8h	h2, v2
	WORD $0x1e260050  // fmov	w16, s2
	WORD $0x4e241c00  // and.16b	v0, v0, v4
	WORD $0x4e050000  // tbl.16b	v0, { v0 }, v5
	WORD $0x4e71b800  // addv.8h	h0, v0
	WORD $0x1e26000f  // fmov	w15, s0
	WORD $0x4e241c20  // and.16b	v0, v1, v4
	WORD $0x4e050000  // tbl.16b	v0, { v0 }, v5
	WORD $0x4e71b800  // addv.8h	h0, v0
	WORD $0x1e260011  // fmov	w17, s0
	WORD $0x33103e0b  // bfi	w11, w16, #16, #16
	WORD $0x33103e2f  // bfi	w15, w17, #16, #16
	WORD $0x3500080f  // cbnz	w15, LBB0_40 $256(%rip)
	WORD $0xb50008ac  // cbnz	x12, LBB0_41 $276(%rip)
	WORD $0xb4000a4b  // cbz	x11, LBB0_42 $328(%rip)
LBB0_32:
	WORD $0xdac0016b  // rbit	x11, x11
	WORD $0xdac0116b  // clz	x11, x11
	WORD $0x14000036  // b	LBB0_39 $216(%rip)
LBB0_33:
	WORD $0xd280000b  // mov	x11, #0
	WORD $0x8b08014d  // add	x13, x10, x8
	WORD $0x92800009  // mov	x9, #-1
LBB0_34:
	WORD $0xf1008190  // subs	x16, x12, #32
	WORD $0x54001103  // b.lo	LBB0_60 $544(%rip)
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
	WORD $0x3500092f  // cbnz	w15, LBB0_56 $292(%rip)
	WORD $0xb50009cb  // cbnz	x11, LBB0_57 $312(%rip)
	WORD $0xb4000b6c  // cbz	x12, LBB0_58 $364(%rip)
LBB0_38:
	WORD $0xdac0018b  // rbit	x11, x12
	WORD $0xdac0116b  // clz	x11, x11
	WORD $0xdac001cc  // rbit	x12, x14
	WORD $0xdac0118c  // clz	x12, x12
	WORD $0xeb0b019f  // cmp	x12, x11
	WORD $0x54001023  // b.lo	LBB0_72 $516(%rip)
LBB0_39:
	WORD $0xcb0a01aa  // sub	x10, x13, x10
	WORD $0x8b0b014a  // add	x10, x10, x11
	WORD $0x9100054e  // add	x14, x10, #1
	WORD $0xb6fff18e  // tbz	x14, #63, LBB0_23 $-464(%rip)
	WORD $0x14000076  // b	LBB0_71 $472(%rip)
LBB0_40:
	WORD $0xdac001f0  // rbit	x16, x15
	WORD $0xdac01210  // clz	x16, x16
	WORD $0xcb0a01b1  // sub	x17, x13, x10
	WORD $0x8b100230  // add	x16, x17, x16
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a901129  // csel	x9, x9, x16, ne
LBB0_41:
	WORD $0x0a2c01f0  // bic	w16, w15, w12
	WORD $0x531f7a11  // lsl	w17, w16, #1
	WORD $0x331f7a0c  // bfi	w12, w16, #1, #31
	WORD $0x0a3101ef  // bic	w15, w15, w17
	WORD $0x1201f1ef  // and	w15, w15, #0xaaaaaaaa
	WORD $0x2b1001ef  // adds	w15, w15, w16
	WORD $0x3200f3f0  // mov	w16, #1431655765
	WORD $0x4a0f060f  // eor	w15, w16, w15, lsl #1
	WORD $0x0a0c01ec  // and	w12, w15, w12
	WORD $0x1a9f37ef  // cset	w15, hs
	WORD $0x2a2c03ec  // mvn	w12, w12
	WORD $0x8a0b018b  // and	x11, x12, x11
	WORD $0xaa0f03ec  // mov	x12, x15
	WORD $0xb5fff60b  // cbnz	x11, LBB0_32 $-320(%rip)
LBB0_42:
	WORD $0x910081ad  // add	x13, x13, #32
	WORD $0xaa0e03eb  // mov	x11, x14
LBB0_43:
	WORD $0xb5000d8c  // cbnz	x12, LBB0_73 $432(%rip)
	WORD $0xb40003ab  // cbz	x11, LBB0_53 $116(%rip)
LBB0_45:
	WORD $0xcb0a03ec  // neg	x12, x10
LBB0_46:
	WORD $0xd280000f  // mov	x15, #0
LBB0_47:
	WORD $0x386f69ae  // ldrb	w14, [x13, x15]
	WORD $0x710089df  // cmp	w14, #34
	WORD $0x540002c0  // b.eq	LBB0_52 $88(%rip)
	WORD $0x710171df  // cmp	w14, #92
	WORD $0x540000a0  // b.eq	LBB0_50 $20(%rip)
	WORD $0x910005ef  // add	x15, x15, #1
	WORD $0xeb0f017f  // cmp	x11, x15
	WORD $0x54ffff21  // b.ne	LBB0_47 $-28(%rip)
	WORD $0x14000015  // b	LBB0_54 $84(%rip)
LBB0_50:
	WORD $0xd100056e  // sub	x14, x11, #1
	WORD $0xeb0f01df  // cmp	x14, x15
	WORD $0x54000dc0  // b.eq	LBB0_77 $440(%rip)
	WORD $0x8b0f01ad  // add	x13, x13, x15
	WORD $0x8b0c01ae  // add	x14, x13, x12
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a8901c9  // csel	x9, x14, x9, eq
	WORD $0x910009ad  // add	x13, x13, #2
	WORD $0xcb0f016e  // sub	x14, x11, x15
	WORD $0xd1000970  // sub	x16, x11, #2
	WORD $0xd10009cb  // sub	x11, x14, #2
	WORD $0x9280000e  // mov	x14, #-1
	WORD $0xeb0f021f  // cmp	x16, x15
	WORD $0x54fffd21  // b.ne	LBB0_46 $-92(%rip)
	WORD $0x14000044  // b	LBB0_71 $272(%rip)
LBB0_52:
	WORD $0x8b0f01ab  // add	x11, x13, x15
	WORD $0x9100056d  // add	x13, x11, #1
LBB0_53:
	WORD $0xcb0a01ae  // sub	x14, x13, x10
	WORD $0xb6ffeaae  // tbz	x14, #63, LBB0_23 $-684(%rip)
	WORD $0x1400003f  // b	LBB0_71 $252(%rip)
LBB0_54:
	WORD $0x710089df  // cmp	w14, #34
	WORD $0x54000b61  // b.ne	LBB0_77 $364(%rip)
	WORD $0x8b0f01ad  // add	x13, x13, x15
	WORD $0x17fffffa  // b	LBB0_53 $-24(%rip)
LBB0_56:
	WORD $0xdac001f1  // rbit	x17, x15
	WORD $0xdac01231  // clz	x17, x17
	WORD $0xcb0a01a3  // sub	x3, x13, x10
	WORD $0x8b110071  // add	x17, x3, x17
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a911129  // csel	x9, x9, x17, ne
LBB0_57:
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
	WORD $0xb5fff4ec  // cbnz	x12, LBB0_38 $-356(%rip)
LBB0_58:
	WORD $0x3500058e  // cbnz	w14, LBB0_72 $176(%rip)
	WORD $0x910081ad  // add	x13, x13, #32
	WORD $0xaa1003ec  // mov	x12, x16
LBB0_60:
	WORD $0xb500072b  // cbnz	x11, LBB0_75 $228(%rip)
	WORD $0xb400080c  // cbz	x12, LBB0_77 $256(%rip)
LBB0_62:
	WORD $0xcb0a03eb  // neg	x11, x10
LBB0_63:
	WORD $0xd280000f  // mov	x15, #0
LBB0_64:
	WORD $0x386f69ae  // ldrb	w14, [x13, x15]
	WORD $0x710089df  // cmp	w14, #34
	WORD $0x54000320  // b.eq	LBB0_70 $100(%rip)
	WORD $0x710171df  // cmp	w14, #92
	WORD $0x54000100  // b.eq	LBB0_68 $32(%rip)
	WORD $0x710081df  // cmp	w14, #32
	WORD $0x540003e3  // b.lo	LBB0_72 $124(%rip)
	WORD $0x910005ef  // add	x15, x15, #1
	WORD $0x9280000e  // mov	x14, #-1
	WORD $0xeb0f019f  // cmp	x12, x15
	WORD $0x54fffec1  // b.ne	LBB0_64 $-40(%rip)
	WORD $0x14000014  // b	LBB0_71 $80(%rip)
LBB0_68:
	WORD $0xd100058e  // sub	x14, x12, #1
	WORD $0xeb0f01df  // cmp	x14, x15
	WORD $0x540005e0  // b.eq	LBB0_77 $188(%rip)
	WORD $0x8b0f01ad  // add	x13, x13, x15
	WORD $0x8b0b01ae  // add	x14, x13, x11
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a8901c9  // csel	x9, x14, x9, eq
	WORD $0x910009ad  // add	x13, x13, #2
	WORD $0xd1000990  // sub	x16, x12, #2
	WORD $0xcb0f018c  // sub	x12, x12, x15
	WORD $0xd100098c  // sub	x12, x12, #2
	WORD $0x9280000e  // mov	x14, #-1
	WORD $0xeb0f021f  // cmp	x16, x15
	WORD $0x54fffcc1  // b.ne	LBB0_63 $-104(%rip)
	WORD $0x14000005  // b	LBB0_71 $20(%rip)
LBB0_70:
	WORD $0xcb0a01aa  // sub	x10, x13, x10
	WORD $0x8b0f014a  // add	x10, x10, x15
	WORD $0x9100054e  // add	x14, x10, #1
	WORD $0xb6ffe2ce  // tbz	x14, #63, LBB0_23 $-936(%rip)
LBB0_71:
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004e  // str	x14, [x2]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_72:
	WORD $0x9280002e  // mov	x14, #-2
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004e  // str	x14, [x2]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
LBB0_73:
	WORD $0xb400022b  // cbz	x11, LBB0_77 $68(%rip)
	WORD $0xaa2a03ec  // mvn	x12, x10
	WORD $0x8b0c01ac  // add	x12, x13, x12
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a890189  // csel	x9, x12, x9, eq
	WORD $0x910005ad  // add	x13, x13, #1
	WORD $0xd100056b  // sub	x11, x11, #1
	WORD $0xb5fff1eb  // cbnz	x11, LBB0_45 $-452(%rip)
	WORD $0x17ffffaa  // b	LBB0_53 $-344(%rip)
LBB0_75:
	WORD $0xb400010c  // cbz	x12, LBB0_77 $32(%rip)
	WORD $0xaa2a03eb  // mvn	x11, x10
	WORD $0x8b0b01ab  // add	x11, x13, x11
	WORD $0xb100053f  // cmn	x9, #1
	WORD $0x9a890169  // csel	x9, x11, x9, eq
	WORD $0x910005ad  // add	x13, x13, #1
	WORD $0xd100058c  // sub	x12, x12, #1
	WORD $0xb5fff84c  // cbnz	x12, LBB0_62 $-248(%rip)
LBB0_77:
	WORD $0x9280000e  // mov	x14, #-1
	WORD $0xf9400408  // ldr	x8, [x0, #8]
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0xf900004e  // str	x14, [x2]
	WORD $0xa9417bfd  // ldp	fp, lr, [sp, #16]
	WORD $0xa8c34ff4  // ldp	x20, x19, [sp], #48
	WORD $0xd65f03c0  // ret
	  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	WORD $0x00000002  // .long 2

TEXT ·__vstring(SB), NOSPLIT, $0-32
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $112, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_vstring:
	MOVD s+0(FP), R0
	MOVD p+8(FP), R1
	MOVD v+16(FP), R2
	MOVD flags+24(FP), R3
	MOVD ·_subr__vstring(SB), R11
	WORD $0x1000005e // adr x30, .+8
	JMP (R11)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
