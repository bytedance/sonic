// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__skip_one_fast_entry__(SB), NOSPLIT, $176
	NO_LOCAL_POINTERS
	WORD $0x10000000  // adr x0, . $0(%rip)
	WORD $0xd65f03c0  // ret
	WORD $0x00000000; WORD $0x00000000  // .p2align 4, 0x00
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

_skip_one_fast:
	WORD $0xd10303ff  // sub	sp, sp, #192
	WORD $0xa9056ffc  // stp	x28, x27, [sp, #80]
	WORD $0xa90667fa  // stp	x26, x25, [sp, #96]
	WORD $0xa9075ff8  // stp	x24, x23, [sp, #112]
	WORD $0xa90857f6  // stp	x22, x21, [sp, #128]
	WORD $0xa9094ff4  // stp	x20, x19, [sp, #144]
	WORD $0xa90a7bfd  // stp	fp, lr, [sp, #160]
	WORD $0xa93ffbfd  // stp	fp, lr, [sp, #-8]
	WORD $0xd10023fd  // sub	fp, sp, #8
	WORD $0xf940002b  // ldr	x11, [x1]
	WORD $0xa9402809  // ldp	x9, x10, [x0]
	WORD $0xeb0a017f  // cmp	x11, x10
	WORD $0x54000142  // b.hs	LBB0_4 $40(%rip)
	WORD $0x386b6928  // ldrb	w8, [x9, x11]
	WORD $0x7100351f  // cmp	w8, #13
	WORD $0x540000e0  // b.eq	LBB0_4 $28(%rip)
	WORD $0x7100811f  // cmp	w8, #32
	WORD $0x540000a0  // b.eq	LBB0_4 $20(%rip)
	WORD $0x51002d0c  // sub	w12, w8, #11
	WORD $0xaa0b03e8  // mov	x8, x11
	WORD $0x3100099f  // cmn	w12, #2
	WORD $0x54000683  // b.lo	LBB0_21 $208(%rip)
LBB0_4:
	WORD $0x91000568  // add	x8, x11, #1
	WORD $0xeb0a011f  // cmp	x8, x10
	WORD $0x54000122  // b.hs	LBB0_8 $36(%rip)
	WORD $0x3868692c  // ldrb	w12, [x9, x8]
	WORD $0x7100359f  // cmp	w12, #13
	WORD $0x540000c0  // b.eq	LBB0_8 $24(%rip)
	WORD $0x7100819f  // cmp	w12, #32
	WORD $0x54000080  // b.eq	LBB0_8 $16(%rip)
	WORD $0x51002d8c  // sub	w12, w12, #11
	WORD $0x3100099f  // cmn	w12, #2
	WORD $0x54000523  // b.lo	LBB0_21 $164(%rip)
LBB0_8:
	WORD $0x91000968  // add	x8, x11, #2
	WORD $0xeb0a011f  // cmp	x8, x10
	WORD $0x54000122  // b.hs	LBB0_12 $36(%rip)
	WORD $0x3868692c  // ldrb	w12, [x9, x8]
	WORD $0x7100359f  // cmp	w12, #13
	WORD $0x540000c0  // b.eq	LBB0_12 $24(%rip)
	WORD $0x7100819f  // cmp	w12, #32
	WORD $0x54000080  // b.eq	LBB0_12 $16(%rip)
	WORD $0x51002d8c  // sub	w12, w12, #11
	WORD $0x3100099f  // cmn	w12, #2
	WORD $0x540003c3  // b.lo	LBB0_21 $120(%rip)
LBB0_12:
	WORD $0x91000d68  // add	x8, x11, #3
	WORD $0xeb0a011f  // cmp	x8, x10
	WORD $0x54000122  // b.hs	LBB0_16 $36(%rip)
	WORD $0x3868692c  // ldrb	w12, [x9, x8]
	WORD $0x7100359f  // cmp	w12, #13
	WORD $0x540000c0  // b.eq	LBB0_16 $24(%rip)
	WORD $0x7100819f  // cmp	w12, #32
	WORD $0x54000080  // b.eq	LBB0_16 $16(%rip)
	WORD $0x51002d8c  // sub	w12, w12, #11
	WORD $0x3100099f  // cmn	w12, #2
	WORD $0x54000263  // b.lo	LBB0_21 $76(%rip)
LBB0_16:
	WORD $0x91001168  // add	x8, x11, #4
	WORD $0xeb0a011f  // cmp	x8, x10
	WORD $0x54005ae2  // b.hs	LBB0_115 $2908(%rip)
	WORD $0x5280002b  // mov	w11, #1
	WORD $0xd284c00c  // mov	x12, #9728
	WORD $0xf2c0002c  // movk	x12, #1, lsl #32
LBB0_18:
	WORD $0x3868692d  // ldrb	w13, [x9, x8]
	WORD $0x710081bf  // cmp	w13, #32
	WORD $0x9acd216d  // lsl	x13, x11, x13
	WORD $0x8a0c01ad  // and	x13, x13, x12
	WORD $0xfa4099a4  // ccmp	x13, #0, #4, ls
	WORD $0x540000a0  // b.eq	LBB0_20 $20(%rip)
	WORD $0x91000508  // add	x8, x8, #1
	WORD $0xeb08015f  // cmp	x10, x8
	WORD $0x54ffff01  // b.ne	LBB0_18 $-32(%rip)
	WORD $0x140002cb  // b	LBB0_116 $2860(%rip)
LBB0_20:
	WORD $0xeb0a011f  // cmp	x8, x10
	WORD $0x54005922  // b.hs	LBB0_116 $2852(%rip)
LBB0_21:
	WORD $0x91000510  // add	x16, x8, #1
	WORD $0xf9000030  // str	x16, [x1]
	WORD $0x3868692a  // ldrb	w10, [x9, x8]
	WORD $0x7101695f  // cmp	w10, #90
	WORD $0x540006ec  // b.gt	LBB0_39 $220(%rip)
	WORD $0x7100bd5f  // cmp	w10, #47
	WORD $0x54000b4d  // b.le	LBB0_44 $360(%rip)
	WORD $0x5100c14a  // sub	w10, w10, #48
	WORD $0x7100295f  // cmp	w10, #10
	WORD $0x540055a2  // b.hs	LBB0_110 $2740(%rip)
LBB0_24:
	WORD $0xf940040a  // ldr	x10, [x0, #8]
	WORD $0xcb10014a  // sub	x10, x10, x16
	WORD $0xf100415f  // cmp	x10, #16
	WORD $0x540002c3  // b.lo	LBB0_28 $88(%rip)
	WORD $0x4f01e580  // movi.16b	v0, #44
	WORD $0x4f06e7e1  // movi.16b	v1, #223
	WORD $0x4f02e7a2  // movi.16b	v2, #93
Lloh0:
	WORD $0x10fff3cb  // adr	x11, lCPI0_0 $-392(%rip)
Lloh1:
	WORD $0x3dc00163  // ldr	q3, [x11, lCPI0_0@PAGEOFF] $0(%rip)
Lloh2:
	WORD $0x10fff40b  // adr	x11, lCPI0_1 $-384(%rip)
Lloh3:
	WORD $0x3dc00164  // ldr	q4, [x11, lCPI0_1@PAGEOFF] $0(%rip)
LBB0_26:
	WORD $0x3cf06925  // ldr	q5, [x9, x16]
	WORD $0x6e208ca6  // cmeq.16b	v6, v5, v0
	WORD $0x4e211ca5  // and.16b	v5, v5, v1
	WORD $0x6e228ca5  // cmeq.16b	v5, v5, v2
	WORD $0x4ea61ca5  // orr.16b	v5, v5, v6
	WORD $0x4e231ca5  // and.16b	v5, v5, v3
	WORD $0x4e0400a5  // tbl.16b	v5, { v5 }, v4
	WORD $0x4e71b8a5  // addv.8h	h5, v5
	WORD $0x1e2600ab  // fmov	w11, s5
	WORD $0x350002eb  // cbnz	w11, LBB0_36 $92(%rip)
	WORD $0xd100414a  // sub	x10, x10, #16
	WORD $0x91004210  // add	x16, x16, #16
	WORD $0xf1003d5f  // cmp	x10, #15
	WORD $0x54fffe68  // b.hi	LBB0_26 $-52(%rip)
LBB0_28:
	WORD $0x8b10012b  // add	x11, x9, x16
	WORD $0xb40001ea  // cbz	x10, LBB0_35 $60(%rip)
	WORD $0x8b0a016c  // add	x12, x11, x10
	WORD $0xcb09016d  // sub	x13, x11, x9
LBB0_30:
	WORD $0x3940016e  // ldrb	w14, [x11]
	WORD $0x7100b1df  // cmp	w14, #44
	WORD $0x540052c0  // b.eq	LBB0_112 $2648(%rip)
	WORD $0x7101f5df  // cmp	w14, #125
	WORD $0x54005280  // b.eq	LBB0_112 $2640(%rip)
	WORD $0x710175df  // cmp	w14, #93
	WORD $0x54005240  // b.eq	LBB0_112 $2632(%rip)
	WORD $0x9100056b  // add	x11, x11, #1
	WORD $0x910005ad  // add	x13, x13, #1
	WORD $0xf100054a  // subs	x10, x10, #1
	WORD $0x54fffec1  // b.ne	LBB0_30 $-40(%rip)
	WORD $0xaa0c03eb  // mov	x11, x12
LBB0_35:
	WORD $0xcb090169  // sub	x9, x11, x9
	WORD $0x14000004  // b	LBB0_37 $16(%rip)
LBB0_36:
	WORD $0x5ac00169  // rbit	w9, w11
	WORD $0x5ac01129  // clz	w9, w9
	WORD $0x8b100129  // add	x9, x9, x16
LBB0_37:
	WORD $0xf9000029  // str	x9, [x1]
LBB0_38:
	WORD $0xaa0803e0  // mov	x0, x8
	WORD $0x1400028f  // b	LBB0_117 $2620(%rip)
LBB0_39:
	WORD $0x7101b55f  // cmp	w10, #109
	WORD $0x5400054d  // b.le	LBB0_47 $168(%rip)
	WORD $0x7101b95f  // cmp	w10, #110
	WORD $0x54002320  // b.eq	LBB0_74 $1124(%rip)
	WORD $0x7101d15f  // cmp	w10, #116
	WORD $0x540022e0  // b.eq	LBB0_74 $1116(%rip)
	WORD $0x7101ed5f  // cmp	w10, #123
	WORD $0x54004e81  // b.ne	LBB0_110 $2512(%rip)
	WORD $0xd2800007  // mov	x7, #0
	WORD $0xd280000f  // mov	x15, #0
	WORD $0xd280000a  // mov	x10, #0
	WORD $0xd280000b  // mov	x11, #0
	WORD $0xb201e3ec  // mov	x12, #-8608480567731124088
	WORD $0xf2e1110c  // movk	x12, #2184, lsl #48
	WORD $0xb202e3ed  // mov	x13, #4919131752989213764
	WORD $0xf2e0888d  // movk	x13, #1092, lsl #48
	WORD $0xb203e3ee  // mov	x14, #2459565876494606882
	WORD $0xf2e0444e  // movk	x14, #546, lsl #48
	WORD $0xf9400411  // ldr	x17, [x0, #8]
	WORD $0xcb100225  // sub	x5, x17, x16
	WORD $0x8b100130  // add	x16, x9, x16
	WORD $0x910043e9  // add	x9, sp, #16
	WORD $0x91008129  // add	x9, x9, #32
	WORD $0x4f01e440  // movi.16b	v0, #34
Lloh4:
	WORD $0x10ffeb91  // adr	x17, lCPI0_0 $-656(%rip)
Lloh5:
	WORD $0x3dc00221  // ldr	q1, [x17, lCPI0_0@PAGEOFF] $0(%rip)
Lloh6:
	WORD $0x10ffebd1  // adr	x17, lCPI0_1 $-648(%rip)
Lloh7:
	WORD $0x3dc00222  // ldr	q2, [x17, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x4f02e783  // movi.16b	v3, #92
	WORD $0xb200e3f1  // mov	x17, #1229782938247303441
	WORD $0xb203e3e2  // mov	x2, #2459565876494606882
	WORD $0xb202e3e3  // mov	x3, #4919131752989213764
	WORD $0xb201e3e4  // mov	x4, #-8608480567731124088
	WORD $0x4f03e764  // movi.16b	v4, #123
	WORD $0x4f03e7a5  // movi.16b	v5, #125
	WORD $0x6f00e406  // movi.2d	v6, #0000000000000000
	WORD $0x14000018  // b	LBB0_51 $96(%rip)
LBB0_44:
	WORD $0x34004d0a  // cbz	w10, LBB0_116 $2464(%rip)
	WORD $0x7100895f  // cmp	w10, #34
	WORD $0x54001f40  // b.eq	LBB0_75 $1000(%rip)
	WORD $0x7100b55f  // cmp	w10, #45
	WORD $0x54fff4c0  // b.eq	LBB0_24 $-360(%rip)
	WORD $0x14000251  // b	LBB0_110 $2372(%rip)
LBB0_47:
	WORD $0x71016d5f  // cmp	w10, #91
	WORD $0x540027a0  // b.eq	LBB0_84 $1268(%rip)
	WORD $0x7101995f  // cmp	w10, #102
	WORD $0x540049a1  // b.ne	LBB0_110 $2356(%rip)
	WORD $0x91001509  // add	x9, x8, #5
	WORD $0xf940040a  // ldr	x10, [x0, #8]
	WORD $0xeb0a013f  // cmp	x9, x10
	WORD $0x54fff969  // b.ls	LBB0_37 $-212(%rip)
	WORD $0x1400025a  // b	LBB0_116 $2408(%rip)
LBB0_50:
	WORD $0x937ffce7  // asr	x7, x7, #63
	WORD $0x9e670267  // fmov	d7, x19
	WORD $0x0e2058e7  // cnt.8b	v7, v7
	WORD $0x2e3038e7  // uaddlv.8b	h7, v7
	WORD $0x1e2600e5  // fmov	w5, s7
	WORD $0x8b0a00aa  // add	x10, x5, x10
	WORD $0x91010210  // add	x16, x16, #64
	WORD $0xaa0603e5  // mov	x5, x6
LBB0_51:
	WORD $0xf10100a6  // subs	x6, x5, #64
	WORD $0x540015cb  // b.lt	LBB0_58 $696(%rip)
LBB0_52:
	WORD $0xad404612  // ldp	q18, q17, [x16]
	WORD $0xad411e10  // ldp	q16, q7, [x16, #32]
	WORD $0x6e238e53  // cmeq.16b	v19, v18, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260273  // fmov	w19, s19
	WORD $0x6e238e33  // cmeq.16b	v19, v17, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260274  // fmov	w20, s19
	WORD $0x6e238e13  // cmeq.16b	v19, v16, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260275  // fmov	w21, s19
	WORD $0x6e238cf3  // cmeq.16b	v19, v7, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260276  // fmov	w22, s19
	WORD $0xd3607eb5  // lsl	x21, x21, #32
	WORD $0xaa16c2b5  // orr	x21, x21, x22, lsl #48
	WORD $0x53103e94  // lsl	w20, w20, #16
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0xaa130293  // orr	x19, x20, x19
	WORD $0xaa0f0274  // orr	x20, x19, x15
	WORD $0xb5000094  // cbnz	x20, LBB0_54 $16(%rip)
	WORD $0xd280000f  // mov	x15, #0
	WORD $0xd2800013  // mov	x19, #0
	WORD $0x1400000a  // b	LBB0_55 $40(%rip)
LBB0_54:
	WORD $0x8a2f0274  // bic	x20, x19, x15
	WORD $0xaa1405f5  // orr	x21, x15, x20, lsl #1
	WORD $0x8a35026f  // bic	x15, x19, x21
	WORD $0x9201f1ef  // and	x15, x15, #0xaaaaaaaaaaaaaaaa
	WORD $0xab1401f3  // adds	x19, x15, x20
	WORD $0x1a9f37ef  // cset	w15, hs
	WORD $0xd37ffa73  // lsl	x19, x19, #1
	WORD $0xd200f273  // eor	x19, x19, #0x5555555555555555
	WORD $0x8a150273  // and	x19, x19, x21
LBB0_55:
	WORD $0x6e208e53  // cmeq.16b	v19, v18, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260274  // fmov	w20, s19
	WORD $0x6e208e33  // cmeq.16b	v19, v17, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260275  // fmov	w21, s19
	WORD $0x6e208e13  // cmeq.16b	v19, v16, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260276  // fmov	w22, s19
	WORD $0x6e208cf3  // cmeq.16b	v19, v7, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260277  // fmov	w23, s19
	WORD $0xd3607ed6  // lsl	x22, x22, #32
	WORD $0xaa17c2d6  // orr	x22, x22, x23, lsl #48
	WORD $0x53103eb5  // lsl	w21, w21, #16
	WORD $0xaa1502d5  // orr	x21, x22, x21
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0x8a330293  // bic	x19, x20, x19
	WORD $0x9200e274  // and	x20, x19, #0x1111111111111111
	WORD $0x9203e275  // and	x21, x19, #0x2222222222222222
	WORD $0x9202e276  // and	x22, x19, #0x4444444444444444
	WORD $0x9201e273  // and	x19, x19, #0x8888888888888888
	WORD $0x9b117e97  // mul	x23, x20, x17
	WORD $0x9b0c7eb8  // mul	x24, x21, x12
	WORD $0xca1802f7  // eor	x23, x23, x24
	WORD $0x9b0d7ed8  // mul	x24, x22, x13
	WORD $0x9b0e7e79  // mul	x25, x19, x14
	WORD $0xca190318  // eor	x24, x24, x25
	WORD $0xca1802f7  // eor	x23, x23, x24
	WORD $0x9b027e98  // mul	x24, x20, x2
	WORD $0x9b117eb9  // mul	x25, x21, x17
	WORD $0xca190318  // eor	x24, x24, x25
	WORD $0x9b0c7ed9  // mul	x25, x22, x12
	WORD $0x9b0d7e7a  // mul	x26, x19, x13
	WORD $0xca1a0339  // eor	x25, x25, x26
	WORD $0xca190318  // eor	x24, x24, x25
	WORD $0x9b037e99  // mul	x25, x20, x3
	WORD $0x9b027eba  // mul	x26, x21, x2
	WORD $0xca1a0339  // eor	x25, x25, x26
	WORD $0x9b117eda  // mul	x26, x22, x17
	WORD $0x9b0c7e7b  // mul	x27, x19, x12
	WORD $0xca1b035a  // eor	x26, x26, x27
	WORD $0xca1a0339  // eor	x25, x25, x26
	WORD $0x9b047e94  // mul	x20, x20, x4
	WORD $0x9b037eb5  // mul	x21, x21, x3
	WORD $0xca150294  // eor	x20, x20, x21
	WORD $0x9b027ed5  // mul	x21, x22, x2
	WORD $0x9b117e73  // mul	x19, x19, x17
	WORD $0xca1302b3  // eor	x19, x21, x19
	WORD $0xca130293  // eor	x19, x20, x19
	WORD $0x9200e2f4  // and	x20, x23, #0x1111111111111111
	WORD $0x9203e315  // and	x21, x24, #0x2222222222222222
	WORD $0x9202e336  // and	x22, x25, #0x4444444444444444
	WORD $0x9201e273  // and	x19, x19, #0x8888888888888888
	WORD $0xaa150294  // orr	x20, x20, x21
	WORD $0xaa1302d3  // orr	x19, x22, x19
	WORD $0xaa130293  // orr	x19, x20, x19
	WORD $0xca070267  // eor	x7, x19, x7
	WORD $0x6e248e53  // cmeq.16b	v19, v18, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260273  // fmov	w19, s19
	WORD $0x6e248e33  // cmeq.16b	v19, v17, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260274  // fmov	w20, s19
	WORD $0x6e248e13  // cmeq.16b	v19, v16, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260275  // fmov	w21, s19
	WORD $0x6e248cf3  // cmeq.16b	v19, v7, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260276  // fmov	w22, s19
	WORD $0xd3607eb5  // lsl	x21, x21, #32
	WORD $0xaa16c2b5  // orr	x21, x21, x22, lsl #48
	WORD $0x53103e94  // lsl	w20, w20, #16
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0xaa130293  // orr	x19, x20, x19
	WORD $0x8a270273  // bic	x19, x19, x7
	WORD $0x6e258e52  // cmeq.16b	v18, v18, v5
	WORD $0x4e211e52  // and.16b	v18, v18, v1
	WORD $0x4e020252  // tbl.16b	v18, { v18 }, v2
	WORD $0x4e71ba52  // addv.8h	h18, v18
	WORD $0x1e260254  // fmov	w20, s18
	WORD $0x6e258e31  // cmeq.16b	v17, v17, v5
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260235  // fmov	w21, s17
	WORD $0x6e258e10  // cmeq.16b	v16, v16, v5
	WORD $0x4e211e10  // and.16b	v16, v16, v1
	WORD $0x4e020210  // tbl.16b	v16, { v16 }, v2
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260216  // fmov	w22, s16
	WORD $0x6e258ce7  // cmeq.16b	v7, v7, v5
	WORD $0x4e211ce7  // and.16b	v7, v7, v1
	WORD $0x4e0200e7  // tbl.16b	v7, { v7 }, v2
	WORD $0x4e71b8e7  // addv.8h	h7, v7
	WORD $0x1e2600f7  // fmov	w23, s7
	WORD $0xd3607ed6  // lsl	x22, x22, #32
	WORD $0xaa17c2d6  // orr	x22, x22, x23, lsl #48
	WORD $0x53103eb5  // lsl	w21, w21, #16
	WORD $0xaa1502d5  // orr	x21, x22, x21
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0xea270294  // bics	x20, x20, x7
	WORD $0x54ffeae0  // b.eq	LBB0_50 $-676(%rip)
LBB0_56:
	WORD $0xd1000695  // sub	x21, x20, #1
	WORD $0x8a1302b6  // and	x22, x21, x19
	WORD $0x9e6702c7  // fmov	d7, x22
	WORD $0x0e2058e7  // cnt.8b	v7, v7
	WORD $0x2e3038e7  // uaddlv.8b	h7, v7
	WORD $0x1e2600f6  // fmov	w22, s7
	WORD $0x8b0a02d6  // add	x22, x22, x10
	WORD $0xeb0b02df  // cmp	x22, x11
	WORD $0x54003109  // b.ls	LBB0_109 $1568(%rip)
	WORD $0x9100056b  // add	x11, x11, #1
	WORD $0xea1402b4  // ands	x20, x21, x20
	WORD $0x54fffea1  // b.ne	LBB0_56 $-44(%rip)
	WORD $0x17ffff4a  // b	LBB0_50 $-728(%rip)
LBB0_58:
	WORD $0xf10000bf  // cmp	x5, #0
	WORD $0x540033ed  // b.le	LBB0_114 $1660(%rip)
	WORD $0xad019be6  // stp	q6, q6, [sp, #48]
	WORD $0xad009be6  // stp	q6, q6, [sp, #16]
	WORD $0x92402e13  // and	x19, x16, #0xfff
	WORD $0xf13f067f  // cmp	x19, #4033
	WORD $0x54ffe9a3  // b.lo	LBB0_52 $-716(%rip)
	WORD $0xf10080b4  // subs	x20, x5, #32
	WORD $0x540000a3  // b.lo	LBB0_62 $20(%rip)
	WORD $0xacc14207  // ldp	q7, q16, [x16], #32
	WORD $0xad00c3e7  // stp	q7, q16, [sp, #16]
	WORD $0xaa0903f3  // mov	x19, x9
	WORD $0x14000003  // b	LBB0_63 $12(%rip)
LBB0_62:
	WORD $0x910043f3  // add	x19, sp, #16
	WORD $0xaa0503f4  // mov	x20, x5
LBB0_63:
	WORD $0xf1004295  // subs	x21, x20, #16
	WORD $0x54000243  // b.lo	LBB0_69 $72(%rip)
	WORD $0x3cc10607  // ldr	q7, [x16], #16
	WORD $0x3c810667  // str	q7, [x19], #16
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xf10022b5  // subs	x21, x21, #8
	WORD $0x540001e2  // b.hs	LBB0_70 $60(%rip)
LBB0_65:
	WORD $0xf1001295  // subs	x21, x20, #4
	WORD $0x54000243  // b.lo	LBB0_71 $72(%rip)
LBB0_66:
	WORD $0xb8404614  // ldr	w20, [x16], #4
	WORD $0xb8004674  // str	w20, [x19], #4
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xf1000ab5  // subs	x21, x21, #2
	WORD $0x540001e2  // b.hs	LBB0_72 $60(%rip)
LBB0_67:
	WORD $0xb4000254  // cbz	x20, LBB0_73 $72(%rip)
LBB0_68:
	WORD $0x39400210  // ldrb	w16, [x16]
	WORD $0x39000270  // strb	w16, [x19]
	WORD $0x910043f0  // add	x16, sp, #16
	WORD $0x17ffff32  // b	LBB0_52 $-824(%rip)
LBB0_69:
	WORD $0xf1002295  // subs	x21, x20, #8
	WORD $0x54fffe63  // b.lo	LBB0_65 $-52(%rip)
LBB0_70:
	WORD $0xf8408614  // ldr	x20, [x16], #8
	WORD $0xf8008674  // str	x20, [x19], #8
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xf10012b5  // subs	x21, x21, #4
	WORD $0x54fffe02  // b.hs	LBB0_66 $-64(%rip)
LBB0_71:
	WORD $0xf1000a95  // subs	x21, x20, #2
	WORD $0x54fffe63  // b.lo	LBB0_67 $-52(%rip)
LBB0_72:
	WORD $0x78402614  // ldrh	w20, [x16], #2
	WORD $0x78002674  // strh	w20, [x19], #2
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xb5fffe15  // cbnz	x21, LBB0_68 $-64(%rip)
LBB0_73:
	WORD $0x910043f0  // add	x16, sp, #16
	WORD $0x17ffff23  // b	LBB0_52 $-884(%rip)
LBB0_74:
	WORD $0x91001109  // add	x9, x8, #4
	WORD $0xf940040a  // ldr	x10, [x0, #8]
	WORD $0xeb0a013f  // cmp	x9, x10
	WORD $0x54ffdbc9  // b.ls	LBB0_37 $-1160(%rip)
	WORD $0x1400016d  // b	LBB0_116 $1460(%rip)
LBB0_75:
	WORD $0xf940040b  // ldr	x11, [x0, #8]
	WORD $0xcb10016a  // sub	x10, x11, x16
	WORD $0xf100815f  // cmp	x10, #32
	WORD $0x54002c4b  // b.lt	LBB0_113 $1416(%rip)
	WORD $0xd280000a  // mov	x10, #0
	WORD $0xd280000d  // mov	x13, #0
	WORD $0x8b08012c  // add	x12, x9, x8
	WORD $0x4f01e440  // movi.16b	v0, #34
Lloh8:
	WORD $0x10ffc96e  // adr	x14, lCPI0_0 $-1748(%rip)
Lloh9:
	WORD $0x3dc001c1  // ldr	q1, [x14, lCPI0_0@PAGEOFF] $0(%rip)
	WORD $0xcb08016b  // sub	x11, x11, x8
Lloh10:
	WORD $0x10ffc98e  // adr	x14, lCPI0_1 $-1744(%rip)
Lloh11:
	WORD $0x3dc001c2  // ldr	q2, [x14, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x528003ee  // mov	w14, #31
	WORD $0x4f02e783  // movi.16b	v3, #92
LBB0_77:
	WORD $0x8b0a018f  // add	x15, x12, x10
	WORD $0x3cc011e4  // ldur	q4, [x15, #1]
	WORD $0x3cc111e5  // ldur	q5, [x15, #17]
	WORD $0x6e208c86  // cmeq.16b	v6, v4, v0
	WORD $0x4e211cc6  // and.16b	v6, v6, v1
	WORD $0x4e0200c6  // tbl.16b	v6, { v6 }, v2
	WORD $0x4e71b8c6  // addv.8h	h6, v6
	WORD $0x1e2600cf  // fmov	w15, s6
	WORD $0x6e208ca6  // cmeq.16b	v6, v5, v0
	WORD $0x4e211cc6  // and.16b	v6, v6, v1
	WORD $0x4e0200c6  // tbl.16b	v6, { v6 }, v2
	WORD $0x4e71b8c6  // addv.8h	h6, v6
	WORD $0x1e2600d0  // fmov	w16, s6
	WORD $0x33103e0f  // bfi	w15, w16, #16, #16
	WORD $0x6e238c84  // cmeq.16b	v4, v4, v3
	WORD $0x4e211c84  // and.16b	v4, v4, v1
	WORD $0x4e020084  // tbl.16b	v4, { v4 }, v2
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260090  // fmov	w16, s4
	WORD $0x6e238ca4  // cmeq.16b	v4, v5, v3
	WORD $0x4e211c84  // and.16b	v4, v4, v1
	WORD $0x4e020084  // tbl.16b	v4, { v4 }, v2
	WORD $0x4e71b884  // addv.8h	h4, v4
	WORD $0x1e260091  // fmov	w17, s4
	WORD $0x33103e30  // bfi	w16, w17, #16, #16
	WORD $0x7100021f  // cmp	w16, #0
	WORD $0xfa4009a0  // ccmp	x13, #0, #0, eq
	WORD $0x540001a0  // b.eq	LBB0_79 $52(%rip)
	WORD $0x0a2d0211  // bic	w17, w16, w13
	WORD $0x2a1105a0  // orr	w0, w13, w17, lsl #1
	WORD $0x0a20020d  // bic	w13, w16, w0
	WORD $0x1201f1ad  // and	w13, w13, #0xaaaaaaaa
	WORD $0x2b1101b0  // adds	w16, w13, w17
	WORD $0x1a9f37ed  // cset	w13, hs
	WORD $0x531f7a10  // lsl	w16, w16, #1
	WORD $0x5200f210  // eor	w16, w16, #0x55555555
	WORD $0x0a000210  // and	w16, w16, w0
	WORD $0x2a3003f0  // mvn	w16, w16
	WORD $0x8a0f020f  // and	x15, x16, x15
	WORD $0x14000002  // b	LBB0_80 $8(%rip)
LBB0_79:
	WORD $0xd280000d  // mov	x13, #0
LBB0_80:
	WORD $0xb50024af  // cbnz	x15, LBB0_111 $1172(%rip)
	WORD $0x9100814a  // add	x10, x10, #32
	WORD $0xd10081ce  // sub	x14, x14, #32
	WORD $0x8b0e016f  // add	x15, x11, x14
	WORD $0xf100fdff  // cmp	x15, #63
	WORD $0x54fffa4c  // b.gt	LBB0_77 $-184(%rip)
	WORD $0xb50026ed  // cbnz	x13, LBB0_118 $1244(%rip)
	WORD $0x8b08012c  // add	x12, x9, x8
	WORD $0x8b0a018c  // add	x12, x12, x10
	WORD $0x9100058c  // add	x12, x12, #1
	WORD $0xaa2a03ea  // mvn	x10, x10
	WORD $0x8b0b014a  // add	x10, x10, x11
	WORD $0x92800000  // mov	x0, #-1
	WORD $0xf100055f  // cmp	x10, #1
	WORD $0x5400280a  // b.ge	LBB0_121 $1280(%rip)
	WORD $0x14000126  // b	LBB0_117 $1176(%rip)
LBB0_84:
	WORD $0xd2800007  // mov	x7, #0
	WORD $0xd280000f  // mov	x15, #0
	WORD $0xd280000a  // mov	x10, #0
	WORD $0xd280000b  // mov	x11, #0
	WORD $0xb201e3ec  // mov	x12, #-8608480567731124088
	WORD $0xf2e1110c  // movk	x12, #2184, lsl #48
	WORD $0xb202e3ed  // mov	x13, #4919131752989213764
	WORD $0xf2e0888d  // movk	x13, #1092, lsl #48
	WORD $0xb203e3ee  // mov	x14, #2459565876494606882
	WORD $0xf2e0444e  // movk	x14, #546, lsl #48
	WORD $0xf9400411  // ldr	x17, [x0, #8]
	WORD $0xcb100225  // sub	x5, x17, x16
	WORD $0x8b100130  // add	x16, x9, x16
	WORD $0x910043e9  // add	x9, sp, #16
	WORD $0x91008129  // add	x9, x9, #32
	WORD $0x4f01e440  // movi.16b	v0, #34
Lloh12:
	WORD $0x10ffbf71  // adr	x17, lCPI0_0 $-2068(%rip)
Lloh13:
	WORD $0x3dc00221  // ldr	q1, [x17, lCPI0_0@PAGEOFF] $0(%rip)
Lloh14:
	WORD $0x10ffbfb1  // adr	x17, lCPI0_1 $-2060(%rip)
Lloh15:
	WORD $0x3dc00222  // ldr	q2, [x17, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0x4f02e783  // movi.16b	v3, #92
	WORD $0xb200e3f1  // mov	x17, #1229782938247303441
	WORD $0xb203e3e2  // mov	x2, #2459565876494606882
	WORD $0xb202e3e3  // mov	x3, #4919131752989213764
	WORD $0xb201e3e4  // mov	x4, #-8608480567731124088
	WORD $0x4f02e764  // movi.16b	v4, #91
	WORD $0x4f02e7a5  // movi.16b	v5, #93
	WORD $0x6f00e406  // movi.2d	v6, #0000000000000000
	WORD $0x14000009  // b	LBB0_86 $36(%rip)
LBB0_85:
	WORD $0x937ffce7  // asr	x7, x7, #63
	WORD $0x9e670267  // fmov	d7, x19
	WORD $0x0e2058e7  // cnt.8b	v7, v7
	WORD $0x2e3038e7  // uaddlv.8b	h7, v7
	WORD $0x1e2600e5  // fmov	w5, s7
	WORD $0x8b0a00aa  // add	x10, x5, x10
	WORD $0x91010210  // add	x16, x16, #64
	WORD $0xaa0603e5  // mov	x5, x6
LBB0_86:
	WORD $0xf10100a6  // subs	x6, x5, #64
	WORD $0x540015cb  // b.lt	LBB0_93 $696(%rip)
LBB0_87:
	WORD $0xad404612  // ldp	q18, q17, [x16]
	WORD $0xad411e10  // ldp	q16, q7, [x16, #32]
	WORD $0x6e238e53  // cmeq.16b	v19, v18, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260273  // fmov	w19, s19
	WORD $0x6e238e33  // cmeq.16b	v19, v17, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260274  // fmov	w20, s19
	WORD $0x6e238e13  // cmeq.16b	v19, v16, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260275  // fmov	w21, s19
	WORD $0x6e238cf3  // cmeq.16b	v19, v7, v3
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260276  // fmov	w22, s19
	WORD $0xd3607eb5  // lsl	x21, x21, #32
	WORD $0xaa16c2b5  // orr	x21, x21, x22, lsl #48
	WORD $0x53103e94  // lsl	w20, w20, #16
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0xaa130293  // orr	x19, x20, x19
	WORD $0xaa0f0274  // orr	x20, x19, x15
	WORD $0xb5000094  // cbnz	x20, LBB0_89 $16(%rip)
	WORD $0xd280000f  // mov	x15, #0
	WORD $0xd2800013  // mov	x19, #0
	WORD $0x1400000a  // b	LBB0_90 $40(%rip)
LBB0_89:
	WORD $0x8a2f0274  // bic	x20, x19, x15
	WORD $0xaa1405f5  // orr	x21, x15, x20, lsl #1
	WORD $0x8a35026f  // bic	x15, x19, x21
	WORD $0x9201f1ef  // and	x15, x15, #0xaaaaaaaaaaaaaaaa
	WORD $0xab1401f3  // adds	x19, x15, x20
	WORD $0x1a9f37ef  // cset	w15, hs
	WORD $0xd37ffa73  // lsl	x19, x19, #1
	WORD $0xd200f273  // eor	x19, x19, #0x5555555555555555
	WORD $0x8a150273  // and	x19, x19, x21
LBB0_90:
	WORD $0x6e208e53  // cmeq.16b	v19, v18, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260274  // fmov	w20, s19
	WORD $0x6e208e33  // cmeq.16b	v19, v17, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260275  // fmov	w21, s19
	WORD $0x6e208e13  // cmeq.16b	v19, v16, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260276  // fmov	w22, s19
	WORD $0x6e208cf3  // cmeq.16b	v19, v7, v0
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260277  // fmov	w23, s19
	WORD $0xd3607ed6  // lsl	x22, x22, #32
	WORD $0xaa17c2d6  // orr	x22, x22, x23, lsl #48
	WORD $0x53103eb5  // lsl	w21, w21, #16
	WORD $0xaa1502d5  // orr	x21, x22, x21
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0x8a330293  // bic	x19, x20, x19
	WORD $0x9200e274  // and	x20, x19, #0x1111111111111111
	WORD $0x9203e275  // and	x21, x19, #0x2222222222222222
	WORD $0x9202e276  // and	x22, x19, #0x4444444444444444
	WORD $0x9201e273  // and	x19, x19, #0x8888888888888888
	WORD $0x9b117e97  // mul	x23, x20, x17
	WORD $0x9b0c7eb8  // mul	x24, x21, x12
	WORD $0xca1802f7  // eor	x23, x23, x24
	WORD $0x9b0d7ed8  // mul	x24, x22, x13
	WORD $0x9b0e7e79  // mul	x25, x19, x14
	WORD $0xca190318  // eor	x24, x24, x25
	WORD $0xca1802f7  // eor	x23, x23, x24
	WORD $0x9b027e98  // mul	x24, x20, x2
	WORD $0x9b117eb9  // mul	x25, x21, x17
	WORD $0xca190318  // eor	x24, x24, x25
	WORD $0x9b0c7ed9  // mul	x25, x22, x12
	WORD $0x9b0d7e7a  // mul	x26, x19, x13
	WORD $0xca1a0339  // eor	x25, x25, x26
	WORD $0xca190318  // eor	x24, x24, x25
	WORD $0x9b037e99  // mul	x25, x20, x3
	WORD $0x9b027eba  // mul	x26, x21, x2
	WORD $0xca1a0339  // eor	x25, x25, x26
	WORD $0x9b117eda  // mul	x26, x22, x17
	WORD $0x9b0c7e7b  // mul	x27, x19, x12
	WORD $0xca1b035a  // eor	x26, x26, x27
	WORD $0xca1a0339  // eor	x25, x25, x26
	WORD $0x9b047e94  // mul	x20, x20, x4
	WORD $0x9b037eb5  // mul	x21, x21, x3
	WORD $0xca150294  // eor	x20, x20, x21
	WORD $0x9b027ed5  // mul	x21, x22, x2
	WORD $0x9b117e73  // mul	x19, x19, x17
	WORD $0xca1302b3  // eor	x19, x21, x19
	WORD $0xca130293  // eor	x19, x20, x19
	WORD $0x9200e2f4  // and	x20, x23, #0x1111111111111111
	WORD $0x9203e315  // and	x21, x24, #0x2222222222222222
	WORD $0x9202e336  // and	x22, x25, #0x4444444444444444
	WORD $0x9201e273  // and	x19, x19, #0x8888888888888888
	WORD $0xaa150294  // orr	x20, x20, x21
	WORD $0xaa1302d3  // orr	x19, x22, x19
	WORD $0xaa130293  // orr	x19, x20, x19
	WORD $0xca070267  // eor	x7, x19, x7
	WORD $0x6e248e53  // cmeq.16b	v19, v18, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260273  // fmov	w19, s19
	WORD $0x6e248e33  // cmeq.16b	v19, v17, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260274  // fmov	w20, s19
	WORD $0x6e248e13  // cmeq.16b	v19, v16, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260275  // fmov	w21, s19
	WORD $0x6e248cf3  // cmeq.16b	v19, v7, v4
	WORD $0x4e211e73  // and.16b	v19, v19, v1
	WORD $0x4e020273  // tbl.16b	v19, { v19 }, v2
	WORD $0x4e71ba73  // addv.8h	h19, v19
	WORD $0x1e260276  // fmov	w22, s19
	WORD $0xd3607eb5  // lsl	x21, x21, #32
	WORD $0xaa16c2b5  // orr	x21, x21, x22, lsl #48
	WORD $0x53103e94  // lsl	w20, w20, #16
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0xaa130293  // orr	x19, x20, x19
	WORD $0x8a270273  // bic	x19, x19, x7
	WORD $0x6e258e52  // cmeq.16b	v18, v18, v5
	WORD $0x4e211e52  // and.16b	v18, v18, v1
	WORD $0x4e020252  // tbl.16b	v18, { v18 }, v2
	WORD $0x4e71ba52  // addv.8h	h18, v18
	WORD $0x1e260254  // fmov	w20, s18
	WORD $0x6e258e31  // cmeq.16b	v17, v17, v5
	WORD $0x4e211e31  // and.16b	v17, v17, v1
	WORD $0x4e020231  // tbl.16b	v17, { v17 }, v2
	WORD $0x4e71ba31  // addv.8h	h17, v17
	WORD $0x1e260235  // fmov	w21, s17
	WORD $0x6e258e10  // cmeq.16b	v16, v16, v5
	WORD $0x4e211e10  // and.16b	v16, v16, v1
	WORD $0x4e020210  // tbl.16b	v16, { v16 }, v2
	WORD $0x4e71ba10  // addv.8h	h16, v16
	WORD $0x1e260216  // fmov	w22, s16
	WORD $0x6e258ce7  // cmeq.16b	v7, v7, v5
	WORD $0x4e211ce7  // and.16b	v7, v7, v1
	WORD $0x4e0200e7  // tbl.16b	v7, { v7 }, v2
	WORD $0x4e71b8e7  // addv.8h	h7, v7
	WORD $0x1e2600f7  // fmov	w23, s7
	WORD $0xd3607ed6  // lsl	x22, x22, #32
	WORD $0xaa17c2d6  // orr	x22, x22, x23, lsl #48
	WORD $0x53103eb5  // lsl	w21, w21, #16
	WORD $0xaa1502d5  // orr	x21, x22, x21
	WORD $0xaa1402b4  // orr	x20, x21, x20
	WORD $0xea270294  // bics	x20, x20, x7
	WORD $0x54ffeae0  // b.eq	LBB0_85 $-676(%rip)
LBB0_91:
	WORD $0xd1000695  // sub	x21, x20, #1
	WORD $0x8a1302b6  // and	x22, x21, x19
	WORD $0x9e6702c7  // fmov	d7, x22
	WORD $0x0e2058e7  // cnt.8b	v7, v7
	WORD $0x2e3038e7  // uaddlv.8b	h7, v7
	WORD $0x1e2600f6  // fmov	w22, s7
	WORD $0x8b0a02d6  // add	x22, x22, x10
	WORD $0xeb0b02df  // cmp	x22, x11
	WORD $0x540006c9  // b.ls	LBB0_109 $216(%rip)
	WORD $0x9100056b  // add	x11, x11, #1
	WORD $0xea1402b4  // ands	x20, x21, x20
	WORD $0x54fffea1  // b.ne	LBB0_91 $-44(%rip)
	WORD $0x17ffff4a  // b	LBB0_85 $-728(%rip)
LBB0_93:
	WORD $0xf10000bf  // cmp	x5, #0
	WORD $0x540009ad  // b.le	LBB0_114 $308(%rip)
	WORD $0xad019be6  // stp	q6, q6, [sp, #48]
	WORD $0xad009be6  // stp	q6, q6, [sp, #16]
	WORD $0x92402e13  // and	x19, x16, #0xfff
	WORD $0xf13f067f  // cmp	x19, #4033
	WORD $0x54ffe9a3  // b.lo	LBB0_87 $-716(%rip)
	WORD $0xf10080b4  // subs	x20, x5, #32
	WORD $0x540000a3  // b.lo	LBB0_97 $20(%rip)
	WORD $0xacc14207  // ldp	q7, q16, [x16], #32
	WORD $0xad00c3e7  // stp	q7, q16, [sp, #16]
	WORD $0xaa0903f3  // mov	x19, x9
	WORD $0x14000003  // b	LBB0_98 $12(%rip)
LBB0_97:
	WORD $0x910043f3  // add	x19, sp, #16
	WORD $0xaa0503f4  // mov	x20, x5
LBB0_98:
	WORD $0xf1004295  // subs	x21, x20, #16
	WORD $0x54000243  // b.lo	LBB0_104 $72(%rip)
	WORD $0x3cc10607  // ldr	q7, [x16], #16
	WORD $0x3c810667  // str	q7, [x19], #16
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xf10022b5  // subs	x21, x21, #8
	WORD $0x540001e2  // b.hs	LBB0_105 $60(%rip)
LBB0_100:
	WORD $0xf1001295  // subs	x21, x20, #4
	WORD $0x54000243  // b.lo	LBB0_106 $72(%rip)
LBB0_101:
	WORD $0xb8404614  // ldr	w20, [x16], #4
	WORD $0xb8004674  // str	w20, [x19], #4
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xf1000ab5  // subs	x21, x21, #2
	WORD $0x540001e2  // b.hs	LBB0_107 $60(%rip)
LBB0_102:
	WORD $0xb4000254  // cbz	x20, LBB0_108 $72(%rip)
LBB0_103:
	WORD $0x39400210  // ldrb	w16, [x16]
	WORD $0x39000270  // strb	w16, [x19]
	WORD $0x910043f0  // add	x16, sp, #16
	WORD $0x17ffff32  // b	LBB0_87 $-824(%rip)
LBB0_104:
	WORD $0xf1002295  // subs	x21, x20, #8
	WORD $0x54fffe63  // b.lo	LBB0_100 $-52(%rip)
LBB0_105:
	WORD $0xf8408614  // ldr	x20, [x16], #8
	WORD $0xf8008674  // str	x20, [x19], #8
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xf10012b5  // subs	x21, x21, #4
	WORD $0x54fffe02  // b.hs	LBB0_101 $-64(%rip)
LBB0_106:
	WORD $0xf1000a95  // subs	x21, x20, #2
	WORD $0x54fffe63  // b.lo	LBB0_102 $-52(%rip)
LBB0_107:
	WORD $0x78402614  // ldrh	w20, [x16], #2
	WORD $0x78002674  // strh	w20, [x19], #2
	WORD $0xaa1503f4  // mov	x20, x21
	WORD $0xb5fffe15  // cbnz	x21, LBB0_103 $-64(%rip)
LBB0_108:
	WORD $0x910043f0  // add	x16, sp, #16
	WORD $0x17ffff23  // b	LBB0_87 $-884(%rip)
LBB0_109:
	WORD $0xf9400409  // ldr	x9, [x0, #8]
	WORD $0xdac0028a  // rbit	x10, x20
	WORD $0xdac0114a  // clz	x10, x10
	WORD $0xcb05014a  // sub	x10, x10, x5
	WORD $0x8b090149  // add	x9, x10, x9
	WORD $0x9100052a  // add	x10, x9, #1
	WORD $0xf900002a  // str	x10, [x1]
	WORD $0xf940040b  // ldr	x11, [x0, #8]
	WORD $0xeb0b015f  // cmp	x10, x11
	WORD $0x9a892569  // csinc	x9, x11, x9, hs
	WORD $0xf9000029  // str	x9, [x1]
	WORD $0xda9f9100  // csinv	x0, x8, xzr, ls
	WORD $0x14000014  // b	LBB0_117 $80(%rip)
LBB0_110:
	WORD $0xf9000028  // str	x8, [x1]
	WORD $0x92800020  // mov	x0, #-2
	WORD $0x14000011  // b	LBB0_117 $68(%rip)
LBB0_111:
	WORD $0xdac001e9  // rbit	x9, x15
	WORD $0xdac01129  // clz	x9, x9
	WORD $0x8b0a010a  // add	x10, x8, x10
	WORD $0x8b0a0129  // add	x9, x9, x10
	WORD $0x91000929  // add	x9, x9, #2
	WORD $0x17fffd7a  // b	LBB0_37 $-2584(%rip)
LBB0_112:
	WORD $0xf900002d  // str	x13, [x1]
	WORD $0x17fffd79  // b	LBB0_38 $-2588(%rip)
LBB0_113:
	WORD $0x8b10012c  // add	x12, x9, x16
	WORD $0x92800000  // mov	x0, #-1
	WORD $0xf100055f  // cmp	x10, #1
	WORD $0x540003ca  // b.ge	LBB0_121 $120(%rip)
	WORD $0x14000004  // b	LBB0_117 $16(%rip)
LBB0_114:
	WORD $0xf9400408  // ldr	x8, [x0, #8]
LBB0_115:
	WORD $0xf9000028  // str	x8, [x1]
LBB0_116:
	WORD $0x92800000  // mov	x0, #-1
LBB0_117:
	WORD $0xa94a7bfd  // ldp	fp, lr, [sp, #160]
	WORD $0xa9494ff4  // ldp	x20, x19, [sp, #144]
	WORD $0xa94857f6  // ldp	x22, x21, [sp, #128]
	WORD $0xa9475ff8  // ldp	x24, x23, [sp, #112]
	WORD $0xa94667fa  // ldp	x26, x25, [sp, #96]
	WORD $0xa9456ffc  // ldp	x28, x27, [sp, #80]
	WORD $0x910303ff  // add	sp, sp, #192
	WORD $0xd65f03c0  // ret
LBB0_118:
	WORD $0xd100056c  // sub	x12, x11, #1
	WORD $0xeb0a019f  // cmp	x12, x10
	WORD $0x54fffea0  // b.eq	LBB0_116 $-44(%rip)
	WORD $0x8b08012c  // add	x12, x9, x8
	WORD $0x8b0a018c  // add	x12, x12, x10
	WORD $0x9100098c  // add	x12, x12, #2
	WORD $0xcb0a016a  // sub	x10, x11, x10
	WORD $0xd100094a  // sub	x10, x10, #2
	WORD $0x92800000  // mov	x0, #-1
	WORD $0xf100055f  // cmp	x10, #1
	WORD $0x540000ea  // b.ge	LBB0_121 $28(%rip)
	WORD $0x17ffffed  // b	LBB0_117 $-76(%rip)
LBB0_120:
	WORD $0x9280002b  // mov	x11, #-2
	WORD $0x5280004d  // mov	w13, #2
	WORD $0x8b0d018c  // add	x12, x12, x13
	WORD $0xab0b014a  // adds	x10, x10, x11
	WORD $0x54fffd0d  // b.le	LBB0_117 $-96(%rip)
LBB0_121:
	WORD $0x3940018b  // ldrb	w11, [x12]
	WORD $0x7101717f  // cmp	w11, #92
	WORD $0x54ffff20  // b.eq	LBB0_120 $-28(%rip)
	WORD $0x7100897f  // cmp	w11, #34
	WORD $0x540000e0  // b.eq	LBB0_124 $28(%rip)
	WORD $0x9280000b  // mov	x11, #-1
	WORD $0x5280002d  // mov	w13, #1
	WORD $0x8b0d018c  // add	x12, x12, x13
	WORD $0xab0b014a  // adds	x10, x10, x11
	WORD $0x54fffeec  // b.gt	LBB0_121 $-36(%rip)
	WORD $0x17ffffdd  // b	LBB0_117 $-140(%rip)
LBB0_124:
	WORD $0xcb090189  // sub	x9, x12, x9
	WORD $0x91000529  // add	x9, x9, #1
	WORD $0x17fffd49  // b	LBB0_37 $-2780(%rip)
	  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	WORD $0x00000002  // .long 2

TEXT ·__skip_one_fast(SB), $0-24
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $256, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_skip_one_fast:
	MOVD s+0(FP), R0
	MOVD p+8(FP), R1
	CALL ·__skip_one_fast_entry__+48(SB)  // _skip_one_fast
	MOVD R0, ret+16(FP)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
