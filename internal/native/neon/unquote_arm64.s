// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__unquote_entry__(SB), NOSPLIT, $80
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

_unquote:
	WORD $0xa9ba6ffc  // stp	x28, x27, [sp, #-96]!
	WORD $0xa90167fa  // stp	x26, x25, [sp, #16]
	WORD $0xa9025ff8  // stp	x24, x23, [sp, #32]
	WORD $0xa90357f6  // stp	x22, x21, [sp, #48]
	WORD $0xa9044ff4  // stp	x20, x19, [sp, #64]
	WORD $0xa9057bfd  // stp	fp, lr, [sp, #80]
	WORD $0xb4002201  // cbz	x1, LBB0_62 $1088(%rip)
	WORD $0x3203cbe9  // mov	w9, #-522133280
	WORD $0x5299fa0a  // mov	w10, #53200
	WORD $0x72b9f9ea  // movk	w10, #53199, lsl #16
	WORD $0x3202c7eb  // mov	w11, #-1061109568
	WORD $0x3201c3ec  // mov	w12, #-2139062144
	WORD $0x3200dbed  // mov	w13, #2139062143
	WORD $0x5288c8ce  // mov	w14, #17990
	WORD $0x72a8c8ce  // movk	w14, #17990, lsl #16
	WORD $0x5287272f  // mov	w15, #14649
	WORD $0x72a7272f  // movk	w15, #14649, lsl #16
	WORD $0x52832330  // mov	w16, #6425
	WORD $0x72a32330  // movk	w16, #6425, lsl #16
	WORD $0x3200cff1  // mov	w17, #252645135
	WORD $0x3200c3e6  // mov	w6, #16843009
	WORD $0x92400085  // and	x5, x4, #0x1
	WORD $0x528017b3  // mov	w19, #189
	WORD $0x52848014  // mov	w20, #9216
	WORD $0x72bf9414  // movk	w20, #64672, lsl #16
	WORD $0x4f02e780  // movi.16b	v0, #92
Lloh0:
	WORD $0x10fffbc8  // adr	x8, lCPI0_0 $-136(%rip)
Lloh1:
	WORD $0x3dc00101  // ldr	q1, [x8, lCPI0_0@PAGEOFF] $0(%rip)
Lloh2:
	WORD $0x10fffc08  // adr	x8, lCPI0_1 $-128(%rip)
Lloh3:
	WORD $0x3dc00102  // ldr	q2, [x8, lCPI0_1@PAGEOFF] $0(%rip)
	WORD $0xaa0003e8  // mov	x8, x0
	WORD $0xaa0103f5  // mov	x21, x1
	WORD $0xaa0203f6  // mov	x22, x2
Lloh4:
	WORD $0x10002ed7  // adr	x23, __UnquoteTab $1496(%rip)
Lloh5:
	WORD $0x910002f7  // add	x23, x23, __UnquoteTab@PAGEOFF $0(%rip)
LBB0_2:
	WORD $0x39400118  // ldrb	w24, [x8]
	WORD $0x7101731f  // cmp	w24, #92
	WORD $0x54000061  // b.ne	LBB0_4 $12(%rip)
	WORD $0xd2800018  // mov	x24, #0
	WORD $0x1400002b  // b	LBB0_15 $172(%rip)
LBB0_4:
	WORD $0xaa1503f9  // mov	x25, x21
	WORD $0xaa1603fb  // mov	x27, x22
	WORD $0xaa0803fa  // mov	x26, x8
	WORD $0xf10042bf  // cmp	x21, #16
	WORD $0x540002cb  // b.lt	LBB0_9 $88(%rip)
	WORD $0xd2800018  // mov	x24, #0
	WORD $0xaa1503fa  // mov	x26, x21
LBB0_6:
	WORD $0x8b180119  // add	x25, x8, x24
	WORD $0x8b1802db  // add	x27, x22, x24
	WORD $0xa9407339  // ldp	x25, x28, [x25]
	WORD $0x9e670323  // fmov	d3, x25
	WORD $0x4e181f83  // mov.d	v3[1], x28
	WORD $0xa9007379  // stp	x25, x28, [x27]
	WORD $0x6e208c63  // cmeq.16b	v3, v3, v0
	WORD $0x4e211c63  // and.16b	v3, v3, v1
	WORD $0x4e020063  // tbl.16b	v3, { v3 }, v2
	WORD $0x4e71b863  // addv.8h	h3, v3
	WORD $0x1e260079  // fmov	w25, s3
	WORD $0x35000279  // cbnz	w25, LBB0_13 $76(%rip)
	WORD $0xd1004359  // sub	x25, x26, #16
	WORD $0x91004318  // add	x24, x24, #16
	WORD $0xf1007f5f  // cmp	x26, #31
	WORD $0xaa1903fa  // mov	x26, x25
	WORD $0x54fffe08  // b.hi	LBB0_6 $-64(%rip)
	WORD $0x8b18011a  // add	x26, x8, x24
	WORD $0x8b1802db  // add	x27, x22, x24
LBB0_9:
	WORD $0xb4001ad9  // cbz	x25, LBB0_63 $856(%rip)
	WORD $0xcb080358  // sub	x24, x26, x8
LBB0_11:
	WORD $0x3940035c  // ldrb	w28, [x26]
	WORD $0x7101739f  // cmp	w28, #92
	WORD $0x54000140  // b.eq	LBB0_14 $40(%rip)
	WORD $0x9100075a  // add	x26, x26, #1
	WORD $0x3800177c  // strb	w28, [x27], #1
	WORD $0x91000718  // add	x24, x24, #1
	WORD $0xf1000739  // subs	x25, x25, #1
	WORD $0x54ffff21  // b.ne	LBB0_11 $-28(%rip)
	WORD $0x140000cc  // b	LBB0_63 $816(%rip)
LBB0_13:
	WORD $0x5ac00339  // rbit	w25, w25
	WORD $0x5ac01339  // clz	w25, w25
	WORD $0x8b180338  // add	x24, x25, x24
LBB0_14:
	WORD $0xb100071f  // cmn	x24, #1
	WORD $0x540018e0  // b.eq	LBB0_63 $796(%rip)
LBB0_15:
	WORD $0x91000b19  // add	x25, x24, #2
	WORD $0xeb1902b5  // subs	x21, x21, x25
	WORD $0x54002804  // b.mi	LBB0_93 $1280(%rip)
	WORD $0x8b190108  // add	x8, x8, x25
	WORD $0xb5000145  // cbnz	x5, LBB0_20 $40(%rip)
	WORD $0x8b1802da  // add	x26, x22, x24
	WORD $0x385ff119  // ldurb	w25, [x8, #-1]
	WORD $0x38796af9  // ldrb	w25, [x23, x25]
	WORD $0x7103ff3f  // cmp	w25, #255
	WORD $0x540003a0  // b.eq	LBB0_28 $116(%rip)
LBB0_18:
	WORD $0x34001df9  // cbz	w25, LBB0_74 $956(%rip)
	WORD $0x38001759  // strb	w25, [x26], #1
	WORD $0xaa1a03f6  // mov	x22, x26
	WORD $0x14000093  // b	LBB0_54 $588(%rip)
LBB0_20:
	WORD $0x34002695  // cbz	w21, LBB0_93 $1232(%rip)
	WORD $0x385ff119  // ldurb	w25, [x8, #-1]
	WORD $0x7101733f  // cmp	w25, #92
	WORD $0x54002381  // b.ne	LBB0_86 $1136(%rip)
	WORD $0x39400119  // ldrb	w25, [x8]
	WORD $0x7101733f  // cmp	w25, #92
	WORD $0x54000161  // b.ne	LBB0_27 $44(%rip)
	WORD $0x710006bf  // cmp	w21, #1
	WORD $0x5400258d  // b.le	LBB0_93 $1200(%rip)
	WORD $0xaa0803f9  // mov	x25, x8
	WORD $0x38401f3a  // ldrb	w26, [x25, #1]!
	WORD $0x71008b5f  // cmp	w26, #34
	WORD $0x54000060  // b.eq	LBB0_26 $12(%rip)
	WORD $0x7101735f  // cmp	w26, #92
	WORD $0x54002281  // b.ne	LBB0_87 $1104(%rip)
LBB0_26:
	WORD $0xd10006b5  // sub	x21, x21, #1
	WORD $0xaa1903e8  // mov	x8, x25
LBB0_27:
	WORD $0x91000508  // add	x8, x8, #1
	WORD $0xd10006b5  // sub	x21, x21, #1
	WORD $0x8b1802da  // add	x26, x22, x24
	WORD $0x385ff119  // ldurb	w25, [x8, #-1]
	WORD $0x38796af9  // ldrb	w25, [x23, x25]
	WORD $0x7103ff3f  // cmp	w25, #255
	WORD $0x54fffca1  // b.ne	LBB0_18 $-108(%rip)
LBB0_28:
	WORD $0xf1000ebf  // cmp	x21, #3
	WORD $0x54002369  // b.ls	LBB0_93 $1132(%rip)
	WORD $0xb9400119  // ldr	w25, [x8]
	WORD $0x0a39019b  // bic	w27, w12, w25
	WORD $0x0b0a033c  // add	w28, w25, w10
	WORD $0x0a1c037c  // and	w28, w27, w28
	WORD $0x7100039f  // cmp	w28, #0
	WORD $0x0b10033c  // add	w28, w25, w16
	WORD $0x2a19039c  // orr	w28, w28, w25
	WORD $0x0a0c039c  // and	w28, w28, w12
	WORD $0x7a400b80  // ccmp	w28, #0, #0, eq
	WORD $0x540013c1  // b.ne	LBB0_65 $632(%rip)
	WORD $0x0a0d033c  // and	w28, w25, w13
	WORD $0x4b1c017e  // sub	w30, w11, w28
	WORD $0x0b0e0387  // add	w7, w28, w14
	WORD $0x0a1e00e7  // and	w7, w7, w30
	WORD $0x6a1b00ff  // tst	w7, w27
	WORD $0x54001301  // b.ne	LBB0_65 $608(%rip)
	WORD $0x4b1c0127  // sub	w7, w9, w28
	WORD $0x0b0f039c  // add	w28, w28, w15
	WORD $0x0a070387  // and	w7, w28, w7
	WORD $0x6a1b00ff  // tst	w7, w27
	WORD $0x54001261  // b.ne	LBB0_65 $588(%rip)
	WORD $0x5ac00b27  // rev	w7, w25
	WORD $0x0a6710d9  // bic	w25, w6, w7, lsr #4
	WORD $0x0b190f39  // add	w25, w25, w25, lsl #3
	WORD $0x0a1100e7  // and	w7, w7, w17
	WORD $0x0b070327  // add	w7, w25, w7
	WORD $0x2a4710e7  // orr	w7, w7, w7, lsr #4
	WORD $0x53105cf9  // ubfx	w25, w7, #16, #8
	WORD $0x12001ce7  // and	w7, w7, #0xff
	WORD $0x2a1920f9  // orr	w25, w7, w25, lsl #8
	WORD $0x91001108  // add	x8, x8, #4
	WORD $0xd10012b5  // sub	x21, x21, #4
	WORD $0x7102033f  // cmp	w25, #128
	WORD $0x54000b23  // b.lo	LBB0_55 $356(%rip)
	WORD $0x8b1802c7  // add	x7, x22, x24
	WORD $0x910008f6  // add	x22, x7, #2
LBB0_34:
	WORD $0x711fff3f  // cmp	w25, #2047
	WORD $0x54000b29  // b.ls	LBB0_57 $356(%rip)
	WORD $0x51403b27  // sub	w7, w25, #14, lsl #12
	WORD $0x312004ff  // cmn	w7, #2049
	WORD $0x540008e9  // b.ls	LBB0_53 $284(%rip)
	WORD $0xb50006c5  // cbnz	x5, LBB0_48 $216(%rip)
	WORD $0xaa1503f8  // mov	x24, x21
	WORD $0x530a7f27  // lsr	w7, w25, #10
	WORD $0x7100d8ff  // cmp	w7, #54
	WORD $0x54000788  // b.hi	LBB0_51 $240(%rip)
LBB0_38:
	WORD $0xf1001b15  // subs	x21, x24, #6
	WORD $0x5400074b  // b.lt	LBB0_51 $232(%rip)
	WORD $0x39400107  // ldrb	w7, [x8]
	WORD $0x710170ff  // cmp	w7, #92
	WORD $0x540006e1  // b.ne	LBB0_51 $220(%rip)
	WORD $0x39400507  // ldrb	w7, [x8, #1]
	WORD $0x7101d4ff  // cmp	w7, #117
	WORD $0x54000681  // b.ne	LBB0_51 $208(%rip)
	WORD $0xb8402118  // ldur	w24, [x8, #2]
	WORD $0x0b0a0307  // add	w7, w24, w10
	WORD $0x0a38019a  // bic	w26, w12, w24
	WORD $0x6a07035f  // tst	w26, w7
	WORD $0x54001401  // b.ne	LBB0_77 $640(%rip)
	WORD $0x0b100307  // add	w7, w24, w16
	WORD $0x2a1800e7  // orr	w7, w7, w24
	WORD $0x6a0c00ff  // tst	w7, w12
	WORD $0x54001381  // b.ne	LBB0_77 $624(%rip)
	WORD $0x0a0d031b  // and	w27, w24, w13
	WORD $0x4b1b0167  // sub	w7, w11, w27
	WORD $0x0b0e037c  // add	w28, w27, w14
	WORD $0x0a070387  // and	w7, w28, w7
	WORD $0x6a1a00ff  // tst	w7, w26
	WORD $0x540012c1  // b.ne	LBB0_77 $600(%rip)
	WORD $0x4b1b0127  // sub	w7, w9, w27
	WORD $0x0b0f037b  // add	w27, w27, w15
	WORD $0x0a070367  // and	w7, w27, w7
	WORD $0x6a1a00ff  // tst	w7, w26
	WORD $0x54001221  // b.ne	LBB0_77 $580(%rip)
	WORD $0x5ac00b07  // rev	w7, w24
	WORD $0x0a6710d8  // bic	w24, w6, w7, lsr #4
	WORD $0x0b180f18  // add	w24, w24, w24, lsl #3
	WORD $0x0a1100e7  // and	w7, w7, w17
	WORD $0x0b070307  // add	w7, w24, w7
	WORD $0x2a4710fa  // orr	w26, w7, w7, lsr #4
	WORD $0x53087f47  // lsr	w7, w26, #8
	WORD $0x12181cf8  // and	w24, w7, #0xff00
	WORD $0x91001908  // add	x8, x8, #6
	WORD $0x51403b07  // sub	w7, w24, #14, lsl #12
	WORD $0x33001f58  // bfxil	w24, w26, #0, #8
	WORD $0x311004ff  // cmn	w7, #1025
	WORD $0x540005e8  // b.hi	LBB0_58 $188(%rip)
	WORD $0x36081684  // tbz	w4, #1, LBB0_88 $720(%rip)
	WORD $0x5297fde7  // mov	w7, #49135
	WORD $0x781fe2c7  // sturh	w7, [x22, #-2]
	WORD $0x380036d3  // strb	w19, [x22], #3
	WORD $0xaa1803f9  // mov	x25, x24
	WORD $0x7102031f  // cmp	w24, #128
	WORD $0x54fff8e2  // b.hs	LBB0_34 $-228(%rip)
	WORD $0x14000037  // b	LBB0_59 $220(%rip)
LBB0_48:
	WORD $0xf10002bf  // cmp	x21, #0
	WORD $0x5400166d  // b.le	LBB0_91 $716(%rip)
	WORD $0x39400107  // ldrb	w7, [x8]
	WORD $0x710170ff  // cmp	w7, #92
	WORD $0x54000681  // b.ne	LBB0_60 $208(%rip)
	WORD $0xd10006b8  // sub	x24, x21, #1
	WORD $0x91000508  // add	x8, x8, #1
	WORD $0x530a7f27  // lsr	w7, w25, #10
	WORD $0x7100d8ff  // cmp	w7, #54
	WORD $0x54fff8c9  // b.ls	LBB0_38 $-232(%rip)
LBB0_51:
	WORD $0x360814e4  // tbz	w4, #1, LBB0_90 $668(%rip)
	WORD $0x5297fde7  // mov	w7, #49135
	WORD $0x781fe2c7  // sturh	w7, [x22, #-2]
	WORD $0x380016d3  // strb	w19, [x22], #1
	WORD $0xaa1803f5  // mov	x21, x24
	WORD $0x1400000a  // b	LBB0_54 $40(%rip)
LBB0_53:
	WORD $0x530c7f27  // lsr	w7, w25, #12
	WORD $0x321b08e7  // orr	w7, w7, #0xe0
	WORD $0x381fe2c7  // sturb	w7, [x22, #-2]
	WORD $0x52801007  // mov	w7, #128
	WORD $0x33062f27  // bfxil	w7, w25, #6, #6
	WORD $0x381ff2c7  // sturb	w7, [x22, #-1]
	WORD $0x52801007  // mov	w7, #128
	WORD $0x33001727  // bfxil	w7, w25, #0, #6
	WORD $0x380016c7  // strb	w7, [x22], #1
LBB0_54:
	WORD $0xb5ffe635  // cbnz	x21, LBB0_2 $-828(%rip)
	WORD $0x14000025  // b	LBB0_63 $148(%rip)
LBB0_55:
	WORD $0xaa1903f8  // mov	x24, x25
LBB0_56:
	WORD $0x38001758  // strb	w24, [x26], #1
	WORD $0xaa1a03f6  // mov	x22, x26
	WORD $0x17fffffb  // b	LBB0_54 $-20(%rip)
LBB0_57:
	WORD $0x53067f27  // lsr	w7, w25, #6
	WORD $0x321a04e7  // orr	w7, w7, #0xc0
	WORD $0x381fe2c7  // sturb	w7, [x22, #-2]
	WORD $0x52801007  // mov	w7, #128
	WORD $0x33001727  // bfxil	w7, w25, #0, #6
	WORD $0x381ff2c7  // sturb	w7, [x22, #-1]
	WORD $0x17fffff4  // b	LBB0_54 $-48(%rip)
LBB0_58:
	WORD $0x0b192b07  // add	w7, w24, w25, lsl #10
	WORD $0x0b1400e7  // add	w7, w7, w20
	WORD $0x53127cf8  // lsr	w24, w7, #18
	WORD $0x321c0f18  // orr	w24, w24, #0xf0
	WORD $0x381fe2d8  // sturb	w24, [x22, #-2]
	WORD $0x52801018  // mov	w24, #128
	WORD $0x330c44f8  // bfxil	w24, w7, #12, #6
	WORD $0x381ff2d8  // sturb	w24, [x22, #-1]
	WORD $0x52801018  // mov	w24, #128
	WORD $0x33062cf8  // bfxil	w24, w7, #6, #6
	WORD $0x390002d8  // strb	w24, [x22]
	WORD $0x52801007  // mov	w7, #128
	WORD $0x33001747  // bfxil	w7, w26, #0, #6
	WORD $0x390006c7  // strb	w7, [x22, #1]
	WORD $0x91000ad6  // add	x22, x22, #2
	WORD $0x17ffffe4  // b	LBB0_54 $-112(%rip)
LBB0_59:
	WORD $0xd1000ada  // sub	x26, x22, #2
	WORD $0x17ffffe5  // b	LBB0_56 $-108(%rip)
LBB0_60:
	WORD $0x36080e84  // tbz	w4, #1, LBB0_88 $464(%rip)
	WORD $0x5297fde7  // mov	w7, #49135
	WORD $0x781fe2c7  // sturh	w7, [x22, #-2]
	WORD $0x380016d3  // strb	w19, [x22], #1
	WORD $0x17ffffdd  // b	LBB0_54 $-140(%rip)
LBB0_62:
	WORD $0xd2800015  // mov	x21, #0
	WORD $0xaa0203f6  // mov	x22, x2
LBB0_63:
	WORD $0x8b1502c8  // add	x8, x22, x21
	WORD $0xcb020100  // sub	x0, x8, x2
LBB0_64:
	WORD $0xa9457bfd  // ldp	fp, lr, [sp, #80]
	WORD $0xa9444ff4  // ldp	x20, x19, [sp, #64]
	WORD $0xa94357f6  // ldp	x22, x21, [sp, #48]
	WORD $0xa9425ff8  // ldp	x24, x23, [sp, #32]
	WORD $0xa94167fa  // ldp	x26, x25, [sp, #16]
	WORD $0xa8c66ffc  // ldp	x28, x27, [sp], #96
	WORD $0xd65f03c0  // ret
LBB0_65:
	WORD $0xcb000109  // sub	x9, x8, x0
	WORD $0xf9000069  // str	x9, [x3]
	WORD $0x3940010a  // ldrb	w10, [x8]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x31002d7f  // cmn	w11, #11
	WORD $0x540000a8  // b.hi	LBB0_67 $20(%rip)
	WORD $0x121a794a  // and	w10, w10, #0xffffffdf
	WORD $0x51011d4a  // sub	w10, w10, #71
	WORD $0x3100195f  // cmn	w10, #6
	WORD $0x540003e3  // b.lo	LBB0_73 $124(%rip)
LBB0_67:
	WORD $0x9100052a  // add	x10, x9, #1
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x3940050a  // ldrb	w10, [x8, #1]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x31002d7f  // cmn	w11, #11
	WORD $0x540000a8  // b.hi	LBB0_69 $20(%rip)
	WORD $0x121a794a  // and	w10, w10, #0xffffffdf
	WORD $0x51011d4a  // sub	w10, w10, #71
	WORD $0x3100195f  // cmn	w10, #6
	WORD $0x540002a3  // b.lo	LBB0_73 $84(%rip)
LBB0_69:
	WORD $0x9100092a  // add	x10, x9, #2
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x3940090a  // ldrb	w10, [x8, #2]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x31002d7f  // cmn	w11, #11
	WORD $0x540000a8  // b.hi	LBB0_71 $20(%rip)
	WORD $0x121a794a  // and	w10, w10, #0xffffffdf
	WORD $0x51011d4a  // sub	w10, w10, #71
	WORD $0x3100195f  // cmn	w10, #6
	WORD $0x54000163  // b.lo	LBB0_73 $44(%rip)
LBB0_71:
	WORD $0x91000d2a  // add	x10, x9, #3
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x39400d08  // ldrb	w8, [x8, #3]
	WORD $0x5100e90a  // sub	w10, w8, #58
	WORD $0x31002d5f  // cmn	w10, #11
	WORD $0x54000188  // b.hi	LBB0_75 $48(%rip)
	WORD $0x121a7908  // and	w8, w8, #0xffffffdf
	WORD $0x51011d08  // sub	w8, w8, #71
	WORD $0x3100191f  // cmn	w8, #6
	WORD $0x54000102  // b.hs	LBB0_75 $32(%rip)
LBB0_73:
	WORD $0x92800020  // mov	x0, #-2
	WORD $0x17ffffd0  // b	LBB0_64 $-192(%rip)
LBB0_74:
	WORD $0xaa2003e9  // mvn	x9, x0
	WORD $0x8b090108  // add	x8, x8, x9
	WORD $0xf9000068  // str	x8, [x3]
	WORD $0x92800040  // mov	x0, #-3
	WORD $0x17ffffcb  // b	LBB0_64 $-212(%rip)
LBB0_75:
	WORD $0x91001128  // add	x8, x9, #4
LBB0_76:
	WORD $0xf9000068  // str	x8, [x3]
	WORD $0x92800020  // mov	x0, #-2
	WORD $0x17ffffc7  // b	LBB0_64 $-228(%rip)
LBB0_77:
	WORD $0xcb000109  // sub	x9, x8, x0
	WORD $0x9100092a  // add	x10, x9, #2
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x3940090a  // ldrb	w10, [x8, #2]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x31002d7f  // cmn	w11, #11
	WORD $0x540000a8  // b.hi	LBB0_79 $20(%rip)
	WORD $0x121a794a  // and	w10, w10, #0xffffffdf
	WORD $0x51011d4a  // sub	w10, w10, #71
	WORD $0x3100195f  // cmn	w10, #6
	WORD $0x54fffd63  // b.lo	LBB0_73 $-84(%rip)
LBB0_79:
	WORD $0x91000d2a  // add	x10, x9, #3
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x39400d0a  // ldrb	w10, [x8, #3]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x31002d7f  // cmn	w11, #11
	WORD $0x540000a8  // b.hi	LBB0_81 $20(%rip)
	WORD $0x121a794a  // and	w10, w10, #0xffffffdf
	WORD $0x51011d4a  // sub	w10, w10, #71
	WORD $0x3100195f  // cmn	w10, #6
	WORD $0x54fffc23  // b.lo	LBB0_73 $-124(%rip)
LBB0_81:
	WORD $0x9100112a  // add	x10, x9, #4
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x3940110a  // ldrb	w10, [x8, #4]
	WORD $0x5100e94b  // sub	w11, w10, #58
	WORD $0x31002d7f  // cmn	w11, #11
	WORD $0x540000a8  // b.hi	LBB0_83 $20(%rip)
	WORD $0x121a794a  // and	w10, w10, #0xffffffdf
	WORD $0x51011d4a  // sub	w10, w10, #71
	WORD $0x3100195f  // cmn	w10, #6
	WORD $0x54fffae3  // b.lo	LBB0_73 $-164(%rip)
LBB0_83:
	WORD $0x9100152a  // add	x10, x9, #5
	WORD $0xf900006a  // str	x10, [x3]
	WORD $0x39401508  // ldrb	w8, [x8, #5]
	WORD $0x5100e90a  // sub	w10, w8, #58
	WORD $0x31002d5f  // cmn	w10, #11
	WORD $0x540000a8  // b.hi	LBB0_85 $20(%rip)
	WORD $0x121a7908  // and	w8, w8, #0xffffffdf
	WORD $0x51011d08  // sub	w8, w8, #71
	WORD $0x3100191f  // cmn	w8, #6
	WORD $0x54fff9a3  // b.lo	LBB0_73 $-204(%rip)
LBB0_85:
	WORD $0x91001928  // add	x8, x9, #6
	WORD $0x17ffffd3  // b	LBB0_76 $-180(%rip)
LBB0_86:
	WORD $0xaa2003e9  // mvn	x9, x0
	WORD $0x8b090108  // add	x8, x8, x9
	WORD $0x17ffffd0  // b	LBB0_76 $-192(%rip)
LBB0_87:
	WORD $0xcb000108  // sub	x8, x8, x0
	WORD $0x91000508  // add	x8, x8, #1
	WORD $0x17ffffcd  // b	LBB0_76 $-204(%rip)
LBB0_88:
	WORD $0xcb000108  // sub	x8, x8, x0
LBB0_89:
	WORD $0xd1001108  // sub	x8, x8, #4
	WORD $0xf9000068  // str	x8, [x3]
	WORD $0x92800060  // mov	x0, #-4
	WORD $0x17ffff91  // b	LBB0_64 $-444(%rip)
LBB0_90:
	WORD $0x8b0000a9  // add	x9, x5, x0
	WORD $0xcb090108  // sub	x8, x8, x9
	WORD $0x17fffffa  // b	LBB0_89 $-24(%rip)
LBB0_91:
	WORD $0x360800e4  // tbz	w4, #1, LBB0_93 $28(%rip)
	WORD $0xd2800015  // mov	x21, #0
	WORD $0x5297fde8  // mov	w8, #49135
	WORD $0x781fe2c8  // sturh	w8, [x22, #-2]
	WORD $0x528017a8  // mov	w8, #189
	WORD $0x380016c8  // strb	w8, [x22], #1
	WORD $0x17ffff85  // b	LBB0_63 $-492(%rip)
LBB0_93:
	WORD $0xf9000061  // str	x1, [x3]
	WORD $0x92800000  // mov	x0, #-1
	WORD $0x17ffff84  // b	LBB0_64 $-496(%rip)
__UnquoteTab:
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00"\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00220000  // .ascii 4, '\x00\x00"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00/'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00/\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00/\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x2f000000  // .ascii 4, '\x00\x00\x00/\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\\\x00\x00\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\x00\x00\\\x00\x00\x00\x00\x00\x08\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\\\x00\x00\x00\x00\x00\x08\x00\x00\x00\x0c\x00'
	WORD $0x0000005c  // .ascii 4, '\\\x00\x00\x00\x00\x00\x08\x00\x00\x00\x0c\x00\x00\x00\x00\x00'
	WORD $0x00080000  // .ascii 4, '\x00\x00\x08\x00\x00\x00\x0c\x00\x00\x00\x00\x00\x00\x00\n\x00'
	WORD $0x000c0000  // .ascii 4, '\x00\x00\x0c\x00\x00\x00\x00\x00\x00\x00\n\x00\x00\x00\r\x00'
	WORD $0x00000000  // .ascii 4, '\x00\x00\x00\x00\x00\x00\n\x00\x00\x00\r\x00\t\xff\x00\x00'
	WORD $0x000a0000  // .ascii 4, '\x00\x00\n\x00\x00\x00\r\x00\t\xff\x00\x00'
	WORD $0x000d0000  // .ascii 4, '\x00\x00\r\x00\t\xff\x00\x00'
	WORD $0x0000ff09  // .ascii 4, '\t\xff\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00\x00\x00\x00\x00'
	WORD $0x00000000  // .space 4, '\x00\x00\x00\x00'

TEXT ·__unquote(SB), $0-48
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $160, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_unquote:
	MOVD sp+0(FP), R0
	MOVD nb+8(FP), R1
	MOVD dp+16(FP), R2
	MOVD ep+24(FP), R3
	MOVD flags+32(FP), R4
	CALL ·__unquote_entry__+48(SB)  // _unquote
	MOVD R0, ret+40(FP)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
