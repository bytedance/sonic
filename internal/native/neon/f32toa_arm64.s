// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__f32toa_entry__(SB), NOSPLIT, $16
	NO_LOCAL_POINTERS
	WORD $0x100000a0 // adr x0, .+20
	MOVD R0, ret(FP)
	RET
	  // .p2align 2, 0x00
_f32toa:
	WORD $0xd10083ff  // sub	sp, sp, #32
	WORD $0xa900fbfd  // stp	fp, lr, [sp, #8]
	WORD $0xa93ffbfd  // stp	fp, lr, [sp, #-8]
	WORD $0xd10023fd  // sub	fp, sp, #8
	WORD $0x1e260009  // fmov	w9, s0
	WORD $0x5317792b  // ubfx	w11, w9, #23, #8
	WORD $0x7103fd7f  // cmp	w11, #255
	WORD $0x54005b20  // b.eq	LBB0_130 $2916(%rip)
	WORD $0x528005a8  // mov	w8, #45
	WORD $0x39000008  // strb	w8, [x0]
	WORD $0x531f7d2a  // lsr	w10, w9, #31
	WORD $0x8b0a0008  // add	x8, x0, x10
	WORD $0x1e26000c  // fmov	w12, s0
	WORD $0x7200799f  // tst	w12, #0x7fffffff
	WORD $0x54000e00  // b.eq	LBB0_10 $448(%rip)
	WORD $0x1200592e  // and	w14, w9, #0x7fffff
	WORD $0x5290d3e9  // mov	w9, #34463
	WORD $0x72a00029  // movk	w9, #1, lsl #16
	WORD $0x34005a4b  // cbz	w11, LBB0_131 $2888(%rip)
	WORD $0x320901cd  // orr	w13, w14, #0x800000
	WORD $0x5102596c  // sub	w12, w11, #150
	WORD $0x5101fd6f  // sub	w15, w11, #127
	WORD $0x71005dff  // cmp	w15, #23
	WORD $0x54000108  // b.hi	LBB0_5 $32(%rip)
	WORD $0x528012cf  // mov	w15, #150
	WORD $0x4b0b01ef  // sub	w15, w15, w11
	WORD $0x92800010  // mov	x16, #-1
	WORD $0x9acf2210  // lsl	x16, x16, x15
	WORD $0x2a3003f0  // mvn	w16, w16
	WORD $0xea0d021f  // tst	x16, x13
	WORD $0x54001120  // b.eq	LBB0_17 $548(%rip)
LBB0_5:
	WORD $0x120001af  // and	w15, w13, #0x1
	WORD $0x710001df  // cmp	w14, #0
	WORD $0x1a9f17ee  // cset	w14, eq
	WORD $0x7100057f  // cmp	w11, #1
	WORD $0x1a9f97eb  // cset	w11, hi
	WORD $0x531e75b0  // lsl	w16, w13, #2
	WORD $0x6a0b01cb  // ands	w11, w14, w11
	WORD $0x2a0b020b  // orr	w11, w16, w11
	WORD $0x52800051  // mov	w17, #2
	WORD $0x331e5db1  // bfi	w17, w13, #2, #24
	WORD $0x5288826d  // mov	w13, #17427
	WORD $0x72a0026d  // movk	w13, #19, lsl #16
	WORD $0x5280202e  // mov	w14, #257
	WORD $0x72bfff0e  // movk	w14, #65528, lsl #16
	WORD $0x1a9f11ce  // csel	w14, w14, wzr, ne
	WORD $0x51000961  // sub	w1, w11, #2
	WORD $0x1b0d398b  // madd	w11, w12, w13, w14
	WORD $0x13167d6b  // asr	w11, w11, #22
	WORD $0x528d962d  // mov	w13, #27825
	WORD $0x72bffcad  // movk	w13, #65509, lsl #16
	WORD $0x1b0d7d6d  // mul	w13, w11, w13
	WORD $0x0b8d4d8c  // add	w12, w12, w13, asr #19
	WORD $0x1100058c  // add	w12, w12, #1
	WORD $0x528003ed  // mov	w13, #31
Lloh0:
	WORD $0x10005c6e  // adr	x14, _pow10_ceil_sig_f32.g $2956(%rip)
Lloh1:
	WORD $0x910001ce  // add	x14, x14, _pow10_ceil_sig_f32.g@PAGEOFF $0(%rip)
	WORD $0x4b0b01ad  // sub	w13, w13, w11
	WORD $0xf86d59c2  // ldr	x2, [x14, w13, uxtw #3]
	WORD $0x1acc202d  // lsl	w13, w1, w12
	WORD $0x9b027dae  // mul	x14, x13, x2
	WORD $0x9bc27dad  // umulh	x13, x13, x2
	WORD $0xf25f79df  // tst	x14, #0xfffffffe00000000
	WORD $0x1a9f07ee  // cset	w14, ne
	WORD $0x2a0d01c1  // orr	w1, w14, w13
	WORD $0x1acc220d  // lsl	w13, w16, w12
	WORD $0x9b027dae  // mul	x14, x13, x2
	WORD $0x9bc27dad  // umulh	x13, x13, x2
	WORD $0xf25f79df  // tst	x14, #0xfffffffe00000000
	WORD $0x1a9f07ee  // cset	w14, ne
	WORD $0x2a0d01ce  // orr	w14, w14, w13
	WORD $0x1acc222c  // lsl	w12, w17, w12
	WORD $0x9b027d90  // mul	x16, x12, x2
	WORD $0x9bc27d8c  // umulh	x12, x12, x2
	WORD $0xf25f7a1f  // tst	x16, #0xfffffffe00000000
	WORD $0x1a9f07f0  // cset	w16, ne
	WORD $0x2a0c020c  // orr	w12, w16, w12
	WORD $0x0b0f0030  // add	w16, w1, w15
	WORD $0x4b0f018f  // sub	w15, w12, w15
	WORD $0x7100a1df  // cmp	w14, #40
	WORD $0x540001e3  // b.lo	LBB0_7 $60(%rip)
	WORD $0x529999ac  // mov	w12, #52429
	WORD $0x72b9998c  // movk	w12, #52428, lsl #16
	WORD $0x9bac7dac  // umull	x12, w13, w12
	WORD $0xd365fd8c  // lsr	x12, x12, #37
	WORD $0x8b0c0991  // add	x17, x12, x12, lsl #2
	WORD $0xd37df231  // lsl	x17, x17, #3
	WORD $0x9100a221  // add	x1, x17, #40
	WORD $0xeb30423f  // cmp	x17, w16, uxtw
	WORD $0x1a9f27f1  // cset	w17, lo
	WORD $0xeb2f403f  // cmp	x1, w15, uxtw
	WORD $0x1a9f87e1  // cset	w1, ls
	WORD $0x1a8c858c  // cinc	w12, w12, ls
	WORD $0x6b01023f  // cmp	w17, w1
	WORD $0x540004e0  // b.eq	LBB0_11 $156(%rip)
LBB0_7:
	WORD $0xd3427dac  // ubfx	x12, x13, #2, #30
	WORD $0x121e75b1  // and	w17, w13, #0xfffffffc
	WORD $0x11001221  // add	w1, w17, #4
	WORD $0x6b0f003f  // cmp	w1, w15
	WORD $0x1a9f87e2  // cset	w2, ls
	WORD $0x6b11021f  // cmp	w16, w17
	WORD $0x1a9f97f0  // cset	w16, hi
	WORD $0x4a020210  // eor	w16, w16, w2
	WORD $0x321f0231  // orr	w17, w17, #0x2
	WORD $0x52800022  // mov	w2, #1
	WORD $0x6b1101df  // cmp	w14, w17
	WORD $0x1a9f17ee  // cset	w14, eq
	WORD $0x0a4d09cd  // and	w13, w14, w13, lsr #2
	WORD $0x1a8d804d  // csel	w13, w2, w13, hi
	WORD $0x0b0c01ad  // add	w13, w13, w12
	WORD $0x6b0f003f  // cmp	w1, w15
	WORD $0x1a8c858c  // cinc	w12, w12, ls
	WORD $0x7200021f  // tst	w16, #0x1
	WORD $0x1a8c11ac  // csel	w12, w13, w12, ne
	WORD $0x6b09019f  // cmp	w12, w9
	WORD $0x540002a9  // b.ls	LBB0_12 $84(%rip)
LBB0_8:
	WORD $0x52884809  // mov	w9, #16960
	WORD $0x72a001e9  // movk	w9, #15, lsl #16
	WORD $0x6b09019f  // cmp	w12, w9
	WORD $0x54000322  // b.hs	LBB0_14 $100(%rip)
	WORD $0x528000cd  // mov	w13, #6
	WORD $0x0b0b01a9  // add	w9, w13, w11
	WORD $0x5100592e  // sub	w14, w9, #22
	WORD $0x310071df  // cmn	w14, #28
	WORD $0x54000ae8  // b.hi	LBB0_22 $348(%rip)
	WORD $0x140000bc  // b	LBB0_43 $752(%rip)
LBB0_10:
	WORD $0x52800609  // mov	w9, #48
	WORD $0x39000109  // strb	w9, [x8]
	WORD $0x4b000108  // sub	w8, w8, w0
	WORD $0x11000500  // add	w0, w8, #1
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_11:
	WORD $0x1100056b  // add	w11, w11, #1
	WORD $0x6b09019f  // cmp	w12, w9
	WORD $0x54fffda8  // b.hi	LBB0_8 $-76(%rip)
LBB0_12:
	WORD $0x7100299f  // cmp	w12, #10
	WORD $0x540002c2  // b.hs	LBB0_15 $88(%rip)
	WORD $0x5280002d  // mov	w13, #1
	WORD $0x0b0b01a9  // add	w9, w13, w11
	WORD $0x5100592e  // sub	w14, w9, #22
	WORD $0x310071df  // cmn	w14, #28
	WORD $0x540008a8  // b.hi	LBB0_22 $276(%rip)
	WORD $0x140000aa  // b	LBB0_43 $680(%rip)
LBB0_14:
	WORD $0x528000e9  // mov	w9, #7
	WORD $0x5292d00d  // mov	w13, #38528
	WORD $0x72a0130d  // movk	w13, #152, lsl #16
	WORD $0x529c200e  // mov	w14, #57600
	WORD $0x72a0beae  // movk	w14, #1525, lsl #16
	WORD $0x6b0e019f  // cmp	w12, w14
	WORD $0x5280010e  // mov	w14, #8
	WORD $0x1a8e35ce  // cinc	w14, w14, hs
	WORD $0x6b0d019f  // cmp	w12, w13
	WORD $0x1a8e312d  // csel	w13, w9, w14, lo
	WORD $0x0b0b01a9  // add	w9, w13, w11
	WORD $0x5100592e  // sub	w14, w9, #22
	WORD $0x310071df  // cmn	w14, #28
	WORD $0x540006c8  // b.hi	LBB0_22 $216(%rip)
	WORD $0x1400009b  // b	LBB0_43 $620(%rip)
LBB0_15:
	WORD $0x7101919f  // cmp	w12, #100
	WORD $0x54000582  // b.hs	LBB0_20 $176(%rip)
	WORD $0x5280004d  // mov	w13, #2
	WORD $0x0b0b01a9  // add	w9, w13, w11
	WORD $0x5100592e  // sub	w14, w9, #22
	WORD $0x310071df  // cmn	w14, #28
	WORD $0x540005c8  // b.hi	LBB0_22 $184(%rip)
	WORD $0x14000093  // b	LBB0_43 $588(%rip)
LBB0_17:
	WORD $0x1acf25ab  // lsr	w11, w13, w15
Lloh2:
	WORD $0x100047ea  // adr	x10, _Digits $2300(%rip)
Lloh3:
	WORD $0x9100014a  // add	x10, x10, _Digits@PAGEOFF $0(%rip)
	WORD $0x6b09017f  // cmp	w11, w9
	WORD $0x54001049  // b.ls	LBB0_40 $520(%rip)
	WORD $0x5292d009  // mov	w9, #38528
	WORD $0x72a01309  // movk	w9, #152, lsl #16
	WORD $0x6b09017f  // cmp	w11, w9
	WORD $0x528000e9  // mov	w9, #7
	WORD $0x9a893529  // cinc	x9, x9, hs
	WORD $0x5288480c  // mov	w12, #16960
	WORD $0x72a001ec  // movk	w12, #15, lsl #16
	WORD $0x6b0c017f  // cmp	w11, w12
	WORD $0x528000cc  // mov	w12, #6
	WORD $0x9a893189  // csel	x9, x12, x9, lo
	WORD $0x8b09010c  // add	x12, x8, x9
LBB0_19:
	WORD $0x5286dc69  // mov	w9, #14051
	WORD $0x72a00349  // movk	w9, #26, lsl #16
	WORD $0x9ba97d69  // umull	x9, w11, w9
	WORD $0xd362fd2d  // lsr	x13, x9, #34
	WORD $0x1284e1e9  // mov	w9, #-10000
	WORD $0x1b092da9  // madd	w9, w13, w9, w11
	WORD $0x5290a3eb  // mov	w11, #34079
	WORD $0x72aa3d6b  // movk	w11, #20971, lsl #16
	WORD $0x9bab7d2b  // umull	x11, w9, w11
	WORD $0xd365fd6b  // lsr	x11, x11, #37
	WORD $0x52800c8e  // mov	w14, #100
	WORD $0x1b0ea569  // msub	w9, w11, w14, w9
	WORD $0x7869594e  // ldrh	w14, [x10, w9, uxtw #1]
	WORD $0x786b794b  // ldrh	w11, [x10, x11, lsl #1]
	WORD $0xaa0c03e9  // mov	x9, x12
	WORD $0x781fcd8b  // strh	w11, [x12, #-4]!
	WORD $0x7900058e  // strh	w14, [x12, #2]
	WORD $0xaa0d03eb  // mov	x11, x13
	WORD $0x7101917f  // cmp	w11, #100
	WORD $0x540022e2  // b.hs	LBB0_77 $1116(%rip)
	WORD $0x14000124  // b	LBB0_79 $1168(%rip)
LBB0_20:
	WORD $0x710fa19f  // cmp	w12, #1000
	WORD $0x54000c82  // b.hs	LBB0_42 $400(%rip)
	WORD $0x5280006d  // mov	w13, #3
	WORD $0x0b0b01a9  // add	w9, w13, w11
	WORD $0x5100592e  // sub	w14, w9, #22
	WORD $0x310071df  // cmn	w14, #28
	WORD $0x54000ce9  // b.ls	LBB0_43 $412(%rip)
LBB0_22:
	WORD $0x37f803cb  // tbnz	w11, #31, LBB0_26 $120(%rip)
	WORD $0x2a0d03ee  // mov	w14, w13
	WORD $0x8b0e010b  // add	x11, x8, x14
	WORD $0x53047d8f  // lsr	w15, w12, #4
	WORD $0x7109c5ff  // cmp	w15, #625
	WORD $0x54000483  // b.lo	LBB0_30 $144(%rip)
	WORD $0x5282eb2f  // mov	w15, #5977
	WORD $0x72ba36ef  // movk	w15, #53687, lsl #16
	WORD $0x9baf7d8f  // umull	x15, w12, w15
	WORD $0xd36dfdf0  // lsr	x16, x15, #45
	WORD $0x1284e1ef  // mov	w15, #-10000
	WORD $0x1b0f320c  // madd	w12, w16, w15, w12
	WORD $0x5290a3ef  // mov	w15, #34079
	WORD $0x72aa3d6f  // movk	w15, #20971, lsl #16
	WORD $0x9baf7d8f  // umull	x15, w12, w15
	WORD $0xd365fdef  // lsr	x15, x15, #37
	WORD $0x52800c91  // mov	w17, #100
	WORD $0x1b11b1ec  // msub	w12, w15, w17, w12
Lloh4:
	WORD $0x10004051  // adr	x17, _Digits $2056(%rip)
Lloh5:
	WORD $0x91000231  // add	x17, x17, _Digits@PAGEOFF $0(%rip)
	WORD $0x786c5a2c  // ldrh	w12, [x17, w12, uxtw #1]
	WORD $0x786f7a31  // ldrh	w17, [x17, x15, lsl #1]
	WORD $0xaa0b03ef  // mov	x15, x11
	WORD $0x781fcdf1  // strh	w17, [x15, #-4]!
	WORD $0x790005ec  // strh	w12, [x15, #2]
	WORD $0xaa1003ec  // mov	x12, x16
	WORD $0x7101919f  // cmp	w12, #100
	WORD $0x54000222  // b.hs	LBB0_31 $68(%rip)
LBB0_25:
	WORD $0xaa0c03ef  // mov	x15, x12
	WORD $0x1400001f  // b	LBB0_33 $124(%rip)
LBB0_26:
	WORD $0x7100013f  // cmp	w9, #0
	WORD $0x540024cc  // b.gt	LBB0_88 $1176(%rip)
	WORD $0x5285c60e  // mov	w14, #11824
	WORD $0x7800250e  // strh	w14, [x8], #2
	WORD $0x36f82469  // tbz	w9, #31, LBB0_88 $1164(%rip)
	WORD $0x2a2d03ee  // mvn	w14, w13
	WORD $0x4b0b01ce  // sub	w14, w14, w11
	WORD $0x7100fddf  // cmp	w14, #63
	WORD $0x54002182  // b.hs	LBB0_83 $1072(%rip)
	WORD $0x5280000e  // mov	w14, #0
	WORD $0x14000117  // b	LBB0_86 $1116(%rip)
LBB0_30:
	WORD $0xaa0b03ef  // mov	x15, x11
	WORD $0x7101919f  // cmp	w12, #100
	WORD $0x54fffe23  // b.lo	LBB0_25 $-60(%rip)
LBB0_31:
	WORD $0xd10005f0  // sub	x16, x15, #1
	WORD $0x5290a3f1  // mov	w17, #34079
	WORD $0x72aa3d71  // movk	w17, #20971, lsl #16
	WORD $0x52800c81  // mov	w1, #100
Lloh6:
	WORD $0x10003c82  // adr	x2, _Digits $1936(%rip)
Lloh7:
	WORD $0x91000042  // add	x2, x2, _Digits@PAGEOFF $0(%rip)
LBB0_32:
	WORD $0x9bb17d8f  // umull	x15, w12, w17
	WORD $0xd365fdef  // lsr	x15, x15, #37
	WORD $0x1b01b1e3  // msub	w3, w15, w1, w12
	WORD $0x78635843  // ldrh	w3, [x2, w3, uxtw #1]
	WORD $0x781ff203  // sturh	w3, [x16, #-1]
	WORD $0xd1000a10  // sub	x16, x16, #2
	WORD $0x53047d83  // lsr	w3, w12, #4
	WORD $0xaa0f03ec  // mov	x12, x15
	WORD $0x7109c07f  // cmp	w3, #624
	WORD $0x54fffee8  // b.hi	LBB0_32 $-36(%rip)
LBB0_33:
	WORD $0x8b09010c  // add	x12, x8, x9
	WORD $0x710029ff  // cmp	w15, #10
	WORD $0x54000163  // b.lo	LBB0_36 $44(%rip)
Lloh8:
	WORD $0x10003ab0  // adr	x16, _Digits $1876(%rip)
Lloh9:
	WORD $0x91000210  // add	x16, x16, _Digits@PAGEOFF $0(%rip)
	WORD $0x786f5a0f  // ldrh	w15, [x16, w15, uxtw #1]
	WORD $0x7900010f  // strh	w15, [x8]
	WORD $0x6b0901bf  // cmp	w13, w9
	WORD $0x54000123  // b.lo	LBB0_37 $36(%rip)
LBB0_35:
	WORD $0x4b000180  // sub	w0, w12, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_36:
	WORD $0x321c05ef  // orr	w15, w15, #0x30
	WORD $0x3900010f  // strb	w15, [x8]
	WORD $0x6b0901bf  // cmp	w13, w9
	WORD $0x54ffff22  // b.hs	LBB0_35 $-28(%rip)
LBB0_37:
	WORD $0x8b0a0008  // add	x8, x0, x10
	WORD $0x8b0e010d  // add	x13, x8, x14
	WORD $0x910005af  // add	x15, x13, #1
	WORD $0x8b090108  // add	x8, x8, x9
	WORD $0xeb0801ff  // cmp	x15, x8
	WORD $0x9a8d9508  // csinc	x8, x8, x13, ls
	WORD $0xcb0d0108  // sub	x8, x8, x13
	WORD $0xf100211f  // cmp	x8, #8
	WORD $0x54003663  // b.lo	LBB0_128 $1740(%rip)
	WORD $0xf101011f  // cmp	x8, #64
	WORD $0x540012a2  // b.hs	LBB0_67 $596(%rip)
	WORD $0xd2800009  // mov	x9, #0
	WORD $0x140000a1  // b	LBB0_71 $644(%rip)
LBB0_40:
	WORD $0x7100297f  // cmp	w11, #10
	WORD $0x540011a2  // b.hs	LBB0_65 $564(%rip)
	WORD $0x52800029  // mov	w9, #1
	WORD $0x140000ad  // b	LBB0_76 $692(%rip)
LBB0_42:
	WORD $0x53047d89  // lsr	w9, w12, #4
	WORD $0x7109c53f  // cmp	w9, #625
	WORD $0x52800089  // mov	w9, #4
	WORD $0x1a89352d  // cinc	w13, w9, hs
	WORD $0x0b0b01a9  // add	w9, w13, w11
	WORD $0x5100592e  // sub	w14, w9, #22
	WORD $0x310071df  // cmn	w14, #28
	WORD $0x54fff368  // b.hi	LBB0_22 $-404(%rip)
LBB0_43:
	WORD $0x9100050b  // add	x11, x8, #1
	WORD $0x2a0d03ed  // mov	w13, w13
	WORD $0x8b0d0170  // add	x16, x11, x13
	WORD $0x53047d8e  // lsr	w14, w12, #4
	WORD $0x7109c5df  // cmp	w14, #625
	WORD $0x540002c3  // b.lo	LBB0_46 $88(%rip)
	WORD $0x5282eb2e  // mov	w14, #5977
	WORD $0x72ba36ee  // movk	w14, #53687, lsl #16
	WORD $0x9bae7d8e  // umull	x14, w12, w14
	WORD $0xd36dfdcf  // lsr	x15, x14, #45
	WORD $0x1284e1ee  // mov	w14, #-10000
	WORD $0x1b0e31ec  // madd	w12, w15, w14, w12
	WORD $0x340002ac  // cbz	w12, LBB0_48 $84(%rip)
	WORD $0xd280000e  // mov	x14, #0
	WORD $0x5290a3f1  // mov	w17, #34079
	WORD $0x72aa3d71  // movk	w17, #20971, lsl #16
	WORD $0x9bb17d91  // umull	x17, w12, w17
	WORD $0xd365fe31  // lsr	x17, x17, #37
	WORD $0x52800c81  // mov	w1, #100
	WORD $0x1b01b22c  // msub	w12, w17, w1, w12
Lloh10:
	WORD $0x10003341  // adr	x1, _Digits $1640(%rip)
Lloh11:
	WORD $0x91000021  // add	x1, x1, _Digits@PAGEOFF $0(%rip)
	WORD $0x786c582c  // ldrh	w12, [x1, w12, uxtw #1]
	WORD $0x781fe20c  // sturh	w12, [x16, #-2]
	WORD $0x7871782c  // ldrh	w12, [x1, x17, lsl #1]
	WORD $0x781fc20c  // sturh	w12, [x16, #-4]
	WORD $0x14000008  // b	LBB0_49 $32(%rip)
LBB0_46:
	WORD $0xd280000e  // mov	x14, #0
	WORD $0xaa0c03ef  // mov	x15, x12
	WORD $0x710191ff  // cmp	w15, #100
	WORD $0x540000e2  // b.hs	LBB0_50 $28(%rip)
LBB0_47:
	WORD $0xaa0f03ec  // mov	x12, x15
	WORD $0x14000015  // b	LBB0_52 $84(%rip)
LBB0_48:
	WORD $0x9280006e  // mov	x14, #-4
LBB0_49:
	WORD $0xd1001210  // sub	x16, x16, #4
	WORD $0x710191ff  // cmp	w15, #100
	WORD $0x54ffff63  // b.lo	LBB0_47 $-20(%rip)
LBB0_50:
	WORD $0xd1000610  // sub	x16, x16, #1
	WORD $0x5290a3f1  // mov	w17, #34079
	WORD $0x72aa3d71  // movk	w17, #20971, lsl #16
	WORD $0x52800c81  // mov	w1, #100
Lloh12:
	WORD $0x100030a2  // adr	x2, _Digits $1556(%rip)
Lloh13:
	WORD $0x91000042  // add	x2, x2, _Digits@PAGEOFF $0(%rip)
LBB0_51:
	WORD $0x9bb17dec  // umull	x12, w15, w17
	WORD $0xd365fd8c  // lsr	x12, x12, #37
	WORD $0x1b01bd83  // msub	w3, w12, w1, w15
	WORD $0x78635843  // ldrh	w3, [x2, w3, uxtw #1]
	WORD $0x781ff203  // sturh	w3, [x16, #-1]
	WORD $0xd1000a10  // sub	x16, x16, #2
	WORD $0x53047de3  // lsr	w3, w15, #4
	WORD $0xaa0c03ef  // mov	x15, x12
	WORD $0x7109c07f  // cmp	w3, #624
	WORD $0x54fffee8  // b.hi	LBB0_51 $-36(%rip)
LBB0_52:
	WORD $0x7100299f  // cmp	w12, #10
	WORD $0x54000123  // b.lo	LBB0_54 $36(%rip)
Lloh14:
	WORD $0x10002eef  // adr	x15, _Digits $1500(%rip)
Lloh15:
	WORD $0x910001ef  // add	x15, x15, _Digits@PAGEOFF $0(%rip)
	WORD $0x8b2c45ef  // add	x15, x15, w12, uxtw #1
	WORD $0x394001ec  // ldrb	w12, [x15]
	WORD $0x3900050c  // strb	w12, [x8, #1]
	WORD $0x394005ef  // ldrb	w15, [x15, #1]
	WORD $0x3900090f  // strb	w15, [x8, #2]
	WORD $0x14000003  // b	LBB0_55 $12(%rip)
LBB0_54:
	WORD $0x321c058c  // orr	w12, w12, #0x30
	WORD $0x3900016c  // strb	w12, [x11]
LBB0_55:
	WORD $0x8b0a01ca  // add	x10, x14, x10
	WORD $0x8b0a000a  // add	x10, x0, x10
	WORD $0x910005ce  // add	x14, x14, #1
LBB0_56:
	WORD $0x386d694f  // ldrb	w15, [x10, x13]
	WORD $0xd100054a  // sub	x10, x10, #1
	WORD $0xd10005ce  // sub	x14, x14, #1
	WORD $0x7100c1ff  // cmp	w15, #48
	WORD $0x54ffff80  // b.eq	LBB0_56 $-16(%rip)
	WORD $0x3900010c  // strb	w12, [x8]
	WORD $0x8b0e01ac  // add	x12, x13, x14
	WORD $0x8b0d0148  // add	x8, x10, x13
	WORD $0xf100099f  // cmp	x12, #2
	WORD $0x540000ab  // b.lt	LBB0_59 $20(%rip)
	WORD $0x91000908  // add	x8, x8, #2
	WORD $0x528005ca  // mov	w10, #46
	WORD $0x3900016a  // strb	w10, [x11]
	WORD $0x14000002  // b	LBB0_60 $8(%rip)
LBB0_59:
	WORD $0x91000508  // add	x8, x8, #1
LBB0_60:
	WORD $0x52800caa  // mov	w10, #101
	WORD $0x3900010a  // strb	w10, [x8]
	WORD $0x5280002a  // mov	w10, #1
	WORD $0x4b09014a  // sub	w10, w10, w9
	WORD $0x71000529  // subs	w9, w9, #1
	WORD $0x5280056b  // mov	w11, #43
	WORD $0x528005ac  // mov	w12, #45
	WORD $0x1a8bb18b  // csel	w11, w12, w11, lt
	WORD $0x1a89b149  // csel	w9, w10, w9, lt
	WORD $0x3900050b  // strb	w11, [x8, #1]
	WORD $0x7101913f  // cmp	w9, #100
	WORD $0x54000243  // b.lo	LBB0_62 $72(%rip)
	WORD $0x529999aa  // mov	w10, #52429
	WORD $0x72b9998a  // movk	w10, #52428, lsl #16
	WORD $0x9baa7d2a  // umull	x10, w9, w10
	WORD $0xd363fd4a  // lsr	x10, x10, #35
	WORD $0x5280014b  // mov	w11, #10
	WORD $0x1b0ba549  // msub	w9, w10, w11, w9
Lloh16:
	WORD $0x1000292b  // adr	x11, _Digits $1316(%rip)
Lloh17:
	WORD $0x9100016b  // add	x11, x11, _Digits@PAGEOFF $0(%rip)
	WORD $0x786a796a  // ldrh	w10, [x11, x10, lsl #1]
	WORD $0x7900050a  // strh	w10, [x8, #2]
	WORD $0x321c0529  // orr	w9, w9, #0x30
	WORD $0x39001109  // strb	w9, [x8, #4]
	WORD $0x9100150c  // add	x12, x8, #5
	WORD $0x4b000180  // sub	w0, w12, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_62:
	WORD $0x7100293f  // cmp	w9, #10
	WORD $0x54000143  // b.lo	LBB0_64 $40(%rip)
Lloh18:
	WORD $0x1000278a  // adr	x10, _Digits $1264(%rip)
Lloh19:
	WORD $0x9100014a  // add	x10, x10, _Digits@PAGEOFF $0(%rip)
	WORD $0x78695949  // ldrh	w9, [x10, w9, uxtw #1]
	WORD $0x79000509  // strh	w9, [x8, #2]
	WORD $0x9100110c  // add	x12, x8, #4
	WORD $0x4b000180  // sub	w0, w12, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_64:
	WORD $0x321c0529  // orr	w9, w9, #0x30
	WORD $0x91000d0c  // add	x12, x8, #3
	WORD $0x39000909  // strb	w9, [x8, #2]
	WORD $0x4b000180  // sub	w0, w12, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_65:
	WORD $0x7101917f  // cmp	w11, #100
	WORD $0x540003c2  // b.hs	LBB0_74 $120(%rip)
	WORD $0x52800049  // mov	w9, #2
	WORD $0x1400001f  // b	LBB0_76 $124(%rip)
LBB0_67:
	WORD $0x927ae509  // and	x9, x8, #0xffffffffffffffc0
	WORD $0x8b0e014d  // add	x13, x10, x14
	WORD $0x8b0001ad  // add	x13, x13, x0
	WORD $0x910081ad  // add	x13, x13, #32
	WORD $0x4f01e600  // movi.16b	v0, #48
	WORD $0xaa0903ef  // mov	x15, x9
LBB0_68:
	WORD $0xad3f01a0  // stp	q0, q0, [x13, #-32]
	WORD $0xac8201a0  // stp	q0, q0, [x13], #64
	WORD $0xf10101ef  // subs	x15, x15, #64
	WORD $0x54ffffa1  // b.ne	LBB0_68 $-12(%rip)
	WORD $0xeb09011f  // cmp	x8, x9
	WORD $0x54ffe9c0  // b.eq	LBB0_35 $-712(%rip)
	WORD $0xf27d091f  // tst	x8, #0x38
	WORD $0x540021c0  // b.eq	LBB0_127 $1080(%rip)
LBB0_71:
	WORD $0x927df10d  // and	x13, x8, #0xfffffffffffffff8
	WORD $0x8b0d016b  // add	x11, x11, x13
	WORD $0x8b0a012a  // add	x10, x9, x10
	WORD $0x8b0e014a  // add	x10, x10, x14
	WORD $0x8b0a000a  // add	x10, x0, x10
	WORD $0xcb0d0129  // sub	x9, x9, x13
	WORD $0x0f01e600  // movi.8b	v0, #48
LBB0_72:
	WORD $0xfc008540  // str	d0, [x10], #8
	WORD $0xb1002129  // adds	x9, x9, #8
	WORD $0x54ffffc1  // b.ne	LBB0_72 $-8(%rip)
	WORD $0xeb0d011f  // cmp	x8, x13
	WORD $0x54ffe800  // b.eq	LBB0_35 $-768(%rip)
	WORD $0x14000102  // b	LBB0_128 $1032(%rip)
LBB0_74:
	WORD $0x710f9d7f  // cmp	w11, #999
	WORD $0x54000448  // b.hi	LBB0_82 $136(%rip)
	WORD $0x52800069  // mov	w9, #3
LBB0_76:
	WORD $0x8b090109  // add	x9, x8, x9
	WORD $0xaa0903ec  // mov	x12, x9
	WORD $0x7101917f  // cmp	w11, #100
	WORD $0x540001e3  // b.lo	LBB0_79 $60(%rip)
LBB0_77:
	WORD $0xd100058c  // sub	x12, x12, #1
	WORD $0x5290a3ed  // mov	w13, #34079
	WORD $0x72aa3d6d  // movk	w13, #20971, lsl #16
	WORD $0x52800c8e  // mov	w14, #100
LBB0_78:
	WORD $0xaa0b03ef  // mov	x15, x11
	WORD $0x9bad7d6b  // umull	x11, w11, w13
	WORD $0xd365fd6b  // lsr	x11, x11, #37
	WORD $0x1b0ebd70  // msub	w16, w11, w14, w15
	WORD $0x78705950  // ldrh	w16, [x10, w16, uxtw #1]
	WORD $0x781ff190  // sturh	w16, [x12, #-1]
	WORD $0xd100098c  // sub	x12, x12, #2
	WORD $0x53047def  // lsr	w15, w15, #4
	WORD $0x7109c1ff  // cmp	w15, #624
	WORD $0x54fffee8  // b.hi	LBB0_78 $-36(%rip)
LBB0_79:
	WORD $0x7100297f  // cmp	w11, #10
	WORD $0x540000e3  // b.lo	LBB0_81 $28(%rip)
	WORD $0x786b594a  // ldrh	w10, [x10, w11, uxtw #1]
	WORD $0x7900010a  // strh	w10, [x8]
	WORD $0x4b000120  // sub	w0, w9, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_81:
	WORD $0x321c056a  // orr	w10, w11, #0x30
	WORD $0x3900010a  // strb	w10, [x8]
	WORD $0x4b000120  // sub	w0, w9, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_82:
	WORD $0x53047d69  // lsr	w9, w11, #4
	WORD $0x5280008c  // mov	w12, #4
	WORD $0x7109c53f  // cmp	w9, #625
	WORD $0x9a8c3589  // cinc	x9, x12, hs
	WORD $0x8b09010c  // add	x12, x8, x9
	WORD $0xaa0c03e9  // mov	x9, x12
	WORD $0x54ffd682  // b.hs	LBB0_19 $-1328(%rip)
	WORD $0x17ffffdd  // b	LBB0_77 $-140(%rip)
LBB0_83:
	WORD $0x910005cf  // add	x15, x14, #1
	WORD $0x927a69ee  // and	x14, x15, #0x1ffffffc0
	WORD $0x8b0e0108  // add	x8, x8, x14
	WORD $0x8b00014a  // add	x10, x10, x0
	WORD $0x9100894a  // add	x10, x10, #34
	WORD $0x4f01e600  // movi.16b	v0, #48
	WORD $0xaa0e03f0  // mov	x16, x14
LBB0_84:
	WORD $0xad3f0140  // stp	q0, q0, [x10, #-32]
	WORD $0xac820140  // stp	q0, q0, [x10], #64
	WORD $0xf1010210  // subs	x16, x16, #64
	WORD $0x54ffffa1  // b.ne	LBB0_84 $-12(%rip)
	WORD $0xeb0e01ff  // cmp	x15, x14
	WORD $0x540000e0  // b.eq	LBB0_88 $28(%rip)
LBB0_86:
	WORD $0x0b0901ca  // add	w10, w14, w9
	WORD $0x4b0a03ea  // neg	w10, w10
	WORD $0x5280060e  // mov	w14, #48
LBB0_87:
	WORD $0x3800150e  // strb	w14, [x8], #1
	WORD $0x7100054a  // subs	w10, w10, #1
	WORD $0x54ffffc1  // b.ne	LBB0_87 $-8(%rip)
LBB0_88:
	WORD $0x2a0d03ee  // mov	w14, w13
	WORD $0x8b0e010d  // add	x13, x8, x14
	WORD $0x53047d8a  // lsr	w10, w12, #4
	WORD $0x7109c55f  // cmp	w10, #625
	WORD $0x540002c3  // b.lo	LBB0_91 $88(%rip)
	WORD $0x5282eb2a  // mov	w10, #5977
	WORD $0x72ba36ea  // movk	w10, #53687, lsl #16
	WORD $0x9baa7d8a  // umull	x10, w12, w10
	WORD $0xd36dfd4a  // lsr	x10, x10, #45
	WORD $0x1284e1ef  // mov	w15, #-10000
	WORD $0x1b0f314c  // madd	w12, w10, w15, w12
	WORD $0x340002cc  // cbz	w12, LBB0_93 $88(%rip)
	WORD $0xd280000f  // mov	x15, #0
	WORD $0x5290a3f0  // mov	w16, #34079
	WORD $0x72aa3d70  // movk	w16, #20971, lsl #16
	WORD $0x9bb07d90  // umull	x16, w12, w16
	WORD $0xd365fe10  // lsr	x16, x16, #37
	WORD $0x52800c91  // mov	w17, #100
	WORD $0x1b11b20c  // msub	w12, w16, w17, w12
Lloh20:
	WORD $0x10001791  // adr	x17, _Digits $752(%rip)
Lloh21:
	WORD $0x91000231  // add	x17, x17, _Digits@PAGEOFF $0(%rip)
	WORD $0x786c5a2c  // ldrh	w12, [x17, w12, uxtw #1]
	WORD $0x781fe1ac  // sturh	w12, [x13, #-2]
	WORD $0x78707a2c  // ldrh	w12, [x17, x16, lsl #1]
	WORD $0x781fc1ac  // sturh	w12, [x13, #-4]
	WORD $0x14000009  // b	LBB0_94 $36(%rip)
LBB0_91:
	WORD $0xd280000f  // mov	x15, #0
	WORD $0xaa0d03f0  // mov	x16, x13
	WORD $0xaa0c03ea  // mov	x10, x12
	WORD $0x7101915f  // cmp	w10, #100
	WORD $0x540000e2  // b.hs	LBB0_95 $28(%rip)
LBB0_92:
	WORD $0xaa0a03f0  // mov	x16, x10
	WORD $0x14000015  // b	LBB0_97 $84(%rip)
LBB0_93:
	WORD $0x9280006f  // mov	x15, #-4
LBB0_94:
	WORD $0xd10011b0  // sub	x16, x13, #4
	WORD $0x7101915f  // cmp	w10, #100
	WORD $0x54ffff63  // b.lo	LBB0_92 $-20(%rip)
LBB0_95:
	WORD $0xd100060c  // sub	x12, x16, #1
	WORD $0x5290a3f1  // mov	w17, #34079
	WORD $0x72aa3d71  // movk	w17, #20971, lsl #16
	WORD $0x52800c81  // mov	w1, #100
Lloh22:
	WORD $0x100014c2  // adr	x2, _Digits $664(%rip)
Lloh23:
	WORD $0x91000042  // add	x2, x2, _Digits@PAGEOFF $0(%rip)
LBB0_96:
	WORD $0x9bb17d50  // umull	x16, w10, w17
	WORD $0xd365fe10  // lsr	x16, x16, #37
	WORD $0x1b01aa03  // msub	w3, w16, w1, w10
	WORD $0x78635843  // ldrh	w3, [x2, w3, uxtw #1]
	WORD $0x781ff183  // sturh	w3, [x12, #-1]
	WORD $0xd100098c  // sub	x12, x12, #2
	WORD $0x53047d43  // lsr	w3, w10, #4
	WORD $0xaa1003ea  // mov	x10, x16
	WORD $0x7109c07f  // cmp	w3, #624
	WORD $0x54fffee8  // b.hi	LBB0_96 $-36(%rip)
LBB0_97:
	WORD $0x71002a1f  // cmp	w16, #10
	WORD $0x540000c3  // b.lo	LBB0_99 $24(%rip)
Lloh24:
	WORD $0x1000130a  // adr	x10, _Digits $608(%rip)
Lloh25:
	WORD $0x9100014a  // add	x10, x10, _Digits@PAGEOFF $0(%rip)
	WORD $0x7870594a  // ldrh	w10, [x10, w16, uxtw #1]
	WORD $0x7900010a  // strh	w10, [x8]
	WORD $0x14000003  // b	LBB0_100 $12(%rip)
LBB0_99:
	WORD $0x321c060a  // orr	w10, w16, #0x30
	WORD $0x3900010a  // strb	w10, [x8]
LBB0_100:
	WORD $0xd280000a  // mov	x10, #0
	WORD $0x8b0f01ad  // add	x13, x13, x15
	WORD $0x4b0f016c  // sub	w12, w11, w15
	WORD $0x51000581  // sub	w1, w12, #1
	WORD $0x51000991  // sub	w17, w12, #2
	WORD $0xaa1103f0  // mov	x16, x17
LBB0_101:
	WORD $0x8b0a01ac  // add	x12, x13, x10
	WORD $0x385ff18c  // ldurb	w12, [x12, #-1]
	WORD $0xd100054a  // sub	x10, x10, #1
	WORD $0x11000610  // add	w16, w16, #1
	WORD $0x7100c19f  // cmp	w12, #48
	WORD $0x54ffff60  // b.eq	LBB0_101 $-20(%rip)
	WORD $0x8b0a01ac  // add	x12, x13, x10
	WORD $0x9100058c  // add	x12, x12, #1
	WORD $0x7100053f  // cmp	w9, #1
	WORD $0x54ffd66b  // b.lt	LBB0_35 $-1332(%rip)
	WORD $0x0b0e01e2  // add	w2, w15, w14
	WORD $0x0b0a0042  // add	w2, w2, w10
	WORD $0x11000442  // add	w2, w2, #1
	WORD $0x6b02013f  // cmp	w9, w2
	WORD $0x5400016a  // b.ge	LBB0_107 $44(%rip)
	WORD $0x4b0b01ee  // sub	w14, w15, w11
	WORD $0x110005cb  // add	w11, w14, #1
	WORD $0x8b0a016c  // add	x12, x11, x10
	WORD $0x7100059f  // cmp	w12, #1
	WORD $0x54000c6b  // b.lt	LBB0_126 $396(%rip)
	WORD $0x92407d8b  // and	x11, x12, #0xffffffff
	WORD $0x7100219f  // cmp	w12, #8
	WORD $0x540001a2  // b.hs	LBB0_110 $52(%rip)
	WORD $0xd280000c  // mov	x12, #0
	WORD $0x14000055  // b	LBB0_124 $340(%rip)
LBB0_107:
	WORD $0xcb0a0029  // sub	x9, x1, x10
	WORD $0x7100053f  // cmp	w9, #1
	WORD $0x54ffd42b  // b.lt	LBB0_35 $-1404(%rip)
	WORD $0x4b0f016b  // sub	w11, w11, w15
	WORD $0x4b0a016b  // sub	w11, w11, w10
	WORD $0x5100096b  // sub	w11, w11, #2
	WORD $0x7100fd7f  // cmp	w11, #63
	WORD $0x540000e2  // b.hs	LBB0_112 $28(%rip)
	WORD $0x5280000b  // mov	w11, #0
	WORD $0x1400001f  // b	LBB0_115 $124(%rip)
LBB0_110:
	WORD $0x7101019f  // cmp	w12, #64
	WORD $0x54000462  // b.hs	LBB0_117 $140(%rip)
	WORD $0xd280000c  // mov	x12, #0
	WORD $0x14000038  // b	LBB0_121 $224(%rip)
LBB0_112:
	WORD $0xd2800001  // mov	x1, #0
	WORD $0xcb0a0222  // sub	x2, x17, x10
	WORD $0x91000571  // add	x17, x11, #1
	WORD $0x927a6a2b  // and	x11, x17, #0x1ffffffc0
	WORD $0x9100060c  // add	x12, x16, #1
	WORD $0x927a698c  // and	x12, x12, #0x1ffffffc0
	WORD $0x8b0e01ee  // add	x14, x15, x14
	WORD $0x8b0e0108  // add	x8, x8, x14
	WORD $0x8b0a0108  // add	x8, x8, x10
	WORD $0x8b0c0108  // add	x8, x8, x12
	WORD $0x9100050c  // add	x12, x8, #1
	WORD $0x92407c48  // and	x8, x2, #0xffffffff
	WORD $0x91000508  // add	x8, x8, #1
	WORD $0x927a6908  // and	x8, x8, #0x1ffffffc0
	WORD $0x4f01e600  // movi.16b	v0, #48
LBB0_113:
	WORD $0x8b0101ae  // add	x14, x13, x1
	WORD $0x8b0a01ce  // add	x14, x14, x10
	WORD $0x3c8011c0  // stur	q0, [x14, #1]
	WORD $0x3c8111c0  // stur	q0, [x14, #17]
	WORD $0x3c8211c0  // stur	q0, [x14, #33]
	WORD $0x3c8311c0  // stur	q0, [x14, #49]
	WORD $0x91010021  // add	x1, x1, #64
	WORD $0xeb01011f  // cmp	x8, x1
	WORD $0x54ffff01  // b.ne	LBB0_113 $-32(%rip)
	WORD $0xeb0b023f  // cmp	x17, x11
	WORD $0x54ffcf80  // b.eq	LBB0_35 $-1552(%rip)
LBB0_115:
	WORD $0x52800608  // mov	w8, #48
LBB0_116:
	WORD $0x38001588  // strb	w8, [x12], #1
	WORD $0x1100056b  // add	w11, w11, #1
	WORD $0x6b09017f  // cmp	w11, w9
	WORD $0x54ffffab  // b.lt	LBB0_116 $-12(%rip)
	WORD $0x17fffe76  // b	LBB0_35 $-1576(%rip)
LBB0_117:
	WORD $0xd280000f  // mov	x15, #0
	WORD $0x927a616c  // and	x12, x11, #0x7fffffc0
	WORD $0x0b0a01d0  // add	w16, w14, w10
	WORD $0x11000610  // add	w16, w16, #1
	WORD $0x927a6210  // and	x16, x16, #0x7fffffc0
	WORD $0xcb1003f0  // neg	x16, x16
	WORD $0x8b0a01b1  // add	x17, x13, x10
LBB0_118:
	WORD $0x8b0f0221  // add	x1, x17, x15
	WORD $0x3cdf1020  // ldur	q0, [x1, #-15]
	WORD $0x3cde1021  // ldur	q1, [x1, #-31]
	WORD $0x3cdd1022  // ldur	q2, [x1, #-47]
	WORD $0x3cdc1023  // ldur	q3, [x1, #-63]
	WORD $0x3c9f2020  // stur	q0, [x1, #-14]
	WORD $0x3c9e2021  // stur	q1, [x1, #-30]
	WORD $0x3c9d2022  // stur	q2, [x1, #-46]
	WORD $0x3c9c2023  // stur	q3, [x1, #-62]
	WORD $0xd10101ef  // sub	x15, x15, #64
	WORD $0xeb0f021f  // cmp	x16, x15
	WORD $0x54fffea1  // b.ne	LBB0_118 $-44(%rip)
	WORD $0xeb0b019f  // cmp	x12, x11
	WORD $0x54000360  // b.eq	LBB0_126 $108(%rip)
	WORD $0xf27d097f  // tst	x11, #0x38
	WORD $0x54000200  // b.eq	LBB0_124 $64(%rip)
LBB0_121:
	WORD $0xcb0c03ef  // neg	x15, x12
	WORD $0x927d6d6c  // and	x12, x11, #0x7ffffff8
	WORD $0x0b0a01ce  // add	w14, w14, w10
	WORD $0x110005ce  // add	w14, w14, #1
	WORD $0x927d6dce  // and	x14, x14, #0x7ffffff8
	WORD $0xcb0e03ee  // neg	x14, x14
	WORD $0x8b0a01b0  // add	x16, x13, x10
LBB0_122:
	WORD $0x8b0f0211  // add	x17, x16, x15
	WORD $0xfc5f9220  // ldur	d0, [x17, #-7]
	WORD $0xfc1fa220  // stur	d0, [x17, #-6]
	WORD $0xd10021ef  // sub	x15, x15, #8
	WORD $0xeb0f01df  // cmp	x14, x15
	WORD $0x54ffff61  // b.ne	LBB0_122 $-20(%rip)
	WORD $0xeb0b019f  // cmp	x12, x11
	WORD $0x54000140  // b.eq	LBB0_126 $40(%rip)
LBB0_124:
	WORD $0xcb0c03ee  // neg	x14, x12
	WORD $0x8b0a01af  // add	x15, x13, x10
LBB0_125:
	WORD $0x8b0e01f0  // add	x16, x15, x14
	WORD $0x386e69f1  // ldrb	w17, [x15, x14]
	WORD $0x39000611  // strb	w17, [x16, #1]
	WORD $0x9100058c  // add	x12, x12, #1
	WORD $0xd10005ce  // sub	x14, x14, #1
	WORD $0xeb0b019f  // cmp	x12, x11
	WORD $0x54ffff43  // b.lo	LBB0_125 $-24(%rip)
LBB0_126:
	WORD $0x528005cb  // mov	w11, #46
	WORD $0x3829490b  // strb	w11, [x8, w9, uxtw]
	WORD $0x8b0a01a8  // add	x8, x13, x10
	WORD $0x9100090c  // add	x12, x8, #2
	WORD $0x4b000180  // sub	w0, w12, w0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_127:
	WORD $0x8b09016b  // add	x11, x11, x9
LBB0_128:
	WORD $0x52800608  // mov	w8, #48
LBB0_129:
	WORD $0x38001568  // strb	w8, [x11], #1
	WORD $0xeb0c017f  // cmp	x11, x12
	WORD $0x54ffffc3  // b.lo	LBB0_129 $-8(%rip)
	WORD $0x17fffe39  // b	LBB0_35 $-1820(%rip)
LBB0_130:
	WORD $0x52800000  // mov	w0, #0
	WORD $0xa940fbfd  // ldp	fp, lr, [sp, #8]
	WORD $0x910083ff  // add	sp, sp, #32
	WORD $0xd65f03c0  // ret
LBB0_131:
	WORD $0x1280128c  // mov	w12, #-149
	WORD $0xaa0e03ed  // mov	x13, x14
	WORD $0x17fffd39  // b	LBB0_5 $-2844(%rip)
_Digits:
	WORD $0x31303030  // .ascii 4, '0001020304050607'
	WORD $0x33303230  // .ascii 4, '0203040506070809'
	WORD $0x35303430  // .ascii 4, '0405060708091011'
	WORD $0x37303630  // .ascii 4, '0607080910111213'
	WORD $0x39303830  // .ascii 4, '0809101112131415'
	WORD $0x31313031  // .ascii 4, '1011121314151617'
	WORD $0x33313231  // .ascii 4, '1213141516171819'
	WORD $0x35313431  // .ascii 4, '1415161718192021'
	WORD $0x37313631  // .ascii 4, '1617181920212223'
	WORD $0x39313831  // .ascii 4, '1819202122232425'
	WORD $0x31323032  // .ascii 4, '2021222324252627'
	WORD $0x33323232  // .ascii 4, '2223242526272829'
	WORD $0x35323432  // .ascii 4, '2425262728293031'
	WORD $0x37323632  // .ascii 4, '2627282930313233'
	WORD $0x39323832  // .ascii 4, '2829303132333435'
	WORD $0x31333033  // .ascii 4, '3031323334353637'
	WORD $0x33333233  // .ascii 4, '3233343536373839'
	WORD $0x35333433  // .ascii 4, '3435363738394041'
	WORD $0x37333633  // .ascii 4, '3637383940414243'
	WORD $0x39333833  // .ascii 4, '3839404142434445'
	WORD $0x31343034  // .ascii 4, '4041424344454647'
	WORD $0x33343234  // .ascii 4, '4243444546474849'
	WORD $0x35343434  // .ascii 4, '4445464748495051'
	WORD $0x37343634  // .ascii 4, '4647484950515253'
	WORD $0x39343834  // .ascii 4, '4849505152535455'
	WORD $0x31353035  // .ascii 4, '5051525354555657'
	WORD $0x33353235  // .ascii 4, '5253545556575859'
	WORD $0x35353435  // .ascii 4, '5455565758596061'
	WORD $0x37353635  // .ascii 4, '5657585960616263'
	WORD $0x39353835  // .ascii 4, '5859606162636465'
	WORD $0x31363036  // .ascii 4, '6061626364656667'
	WORD $0x33363236  // .ascii 4, '6263646566676869'
	WORD $0x35363436  // .ascii 4, '6465666768697071'
	WORD $0x37363636  // .ascii 4, '6667686970717273'
	WORD $0x39363836  // .ascii 4, '6869707172737475'
	WORD $0x31373037  // .ascii 4, '7071727374757677'
	WORD $0x33373237  // .ascii 4, '7273747576777879'
	WORD $0x35373437  // .ascii 4, '7475767778798081'
	WORD $0x37373637  // .ascii 4, '7677787980818283'
	WORD $0x39373837  // .ascii 4, '7879808182838485'
	WORD $0x31383038  // .ascii 4, '8081828384858687'
	WORD $0x33383238  // .ascii 4, '8283848586878889'
	WORD $0x35383438  // .ascii 4, '8485868788899091'
	WORD $0x37383638  // .ascii 4, '8687888990919293'
	WORD $0x39383838  // .ascii 4, '8889909192939495'
	WORD $0x31393039  // .ascii 4, '9091929394959697'
	WORD $0x33393239  // .ascii 4, '9293949596979899'
	WORD $0x35393439  // .ascii 4, '949596979899'
	WORD $0x37393639  // .ascii 4, '96979899'
	WORD $0x39393839  // .ascii 4, '9899'
	WORD $0x00000000  // .p2align 3, 0x00
_LB_d22aa7db: // _pow10_ceil_sig_f32.g
	WORD $0x4b43fcf5; WORD $0x81ceb32c  // .quad -9093133594791772939
	WORD $0x5e14fc32; WORD $0xa2425ff7  // .quad -6754730975062328270
	WORD $0x359a3b3f; WORD $0xcad2f7f5  // .quad -3831727700400522433
	WORD $0x8300ca0e; WORD $0xfd87b5f2  // .quad -177973607073265138
	WORD $0x91e07e49; WORD $0x9e74d1b7  // .quad -7028762532061872567
	WORD $0x76589ddb; WORD $0xc6120625  // .quad -4174267146649952805
	WORD $0xd3eec552; WORD $0xf79687ae  // .quad -606147914885053102
	WORD $0x44753b53; WORD $0x9abe14cd  // .quad -7296371474444240045
	WORD $0x95928a28; WORD $0xc16d9a00  // .quad -4508778324627912152
	WORD $0xbaf72cb2; WORD $0xf1c90080  // .quad -1024286887357502286
	WORD $0x74da7bef; WORD $0x971da050  // .quad -7557708332239520785
	WORD $0x92111aeb; WORD $0xbce50864  // .quad -4835449396872013077
	WORD $0xb69561a6; WORD $0xec1e4a7d  // .quad -1432625727662628442
	WORD $0x921d5d08; WORD $0x9392ee8e  // .quad -7812920107430224632
	WORD $0x36a4b44a; WORD $0xb877aa32  // .quad -5154464115860392886
	WORD $0xc44de15c; WORD $0xe69594be  // .quad -1831394126398103204
	WORD $0x3ab0acda; WORD $0x901d7cf7  // .quad -8062150356639896358
	WORD $0x095cd810; WORD $0xb424dc35  // .quad -5466001927372482544
	WORD $0x4bb40e14; WORD $0xe12e1342  // .quad -2220816390788215276
	WORD $0x6f5088cc; WORD $0x8cbccc09  // .quad -8305539271883716404
	WORD $0xcb24aaff; WORD $0xafebff0b  // .quad -5770238071427257601
	WORD $0xbdedd5bf; WORD $0xdbe6fece  // .quad -2601111570856684097
	WORD $0x36b4a598; WORD $0x89705f41  // .quad -8543223759426509416
	WORD $0x8461cefd; WORD $0xabcc7711  // .quad -6067343680855748867
	WORD $0xe57a42bd; WORD $0xd6bf94d5  // .quad -2972493582642298179
	WORD $0xaf6c69b6; WORD $0x8637bd05  // .quad -8775337516792518218
	WORD $0x1b478424; WORD $0xa7c5ac47  // .quad -6357485877563259868
	WORD $0xe219652c; WORD $0xd1b71758  // .quad -3335171328526686932
	WORD $0x8d4fdf3c; WORD $0x83126e97  // .quad -9002011107970261188
	WORD $0x70a3d70b; WORD $0xa3d70a3d  // .quad -6640827866535438581
	WORD $0xcccccccd; WORD $0xcccccccc  // .quad -3689348814741910323
	WORD $0x00000000; WORD $0x80000000  // .quad -9223372036854775808
	WORD $0x00000000; WORD $0xa0000000  // .quad -6917529027641081856
	WORD $0x00000000; WORD $0xc8000000  // .quad -4035225266123964416
	WORD $0x00000000; WORD $0xfa000000  // .quad -432345564227567616
	WORD $0x00000000; WORD $0x9c400000  // .quad -7187745005283311616
	WORD $0x00000000; WORD $0xc3500000  // .quad -4372995238176751616
	WORD $0x00000000; WORD $0xf4240000  // .quad -854558029293551616
	WORD $0x00000000; WORD $0x98968000  // .quad -7451627795949551616
	WORD $0x00000000; WORD $0xbebc2000  // .quad -4702848726509551616
	WORD $0x00000000; WORD $0xee6b2800  // .quad -1266874889709551616
	WORD $0x00000000; WORD $0x9502f900  // .quad -7709325833709551616
	WORD $0x00000000; WORD $0xba43b740  // .quad -5024971273709551616
	WORD $0x00000000; WORD $0xe8d4a510  // .quad -1669528073709551616
	WORD $0x00000000; WORD $0x9184e72a  // .quad -7960984073709551616
	WORD $0x80000000; WORD $0xb5e620f4  // .quad -5339544073709551616
	WORD $0xa0000000; WORD $0xe35fa931  // .quad -2062744073709551616
	WORD $0x04000000; WORD $0x8e1bc9bf  // .quad -8206744073709551616
	WORD $0xc5000000; WORD $0xb1a2bc2e  // .quad -5646744073709551616
	WORD $0x76400000; WORD $0xde0b6b3a  // .quad -2446744073709551616
	WORD $0x89e80000; WORD $0x8ac72304  // .quad -8446744073709551616
	WORD $0xac620000; WORD $0xad78ebc5  // .quad -5946744073709551616
	WORD $0x177a8000; WORD $0xd8d726b7  // .quad -2821744073709551616
	WORD $0x6eac9000; WORD $0x87867832  // .quad -8681119073709551616
	WORD $0x0a57b400; WORD $0xa968163f  // .quad -6239712823709551616
	WORD $0xcceda100; WORD $0xd3c21bce  // .quad -3187955011209551616
	WORD $0x401484a0; WORD $0x84595161  // .quad -8910000909647051616
	WORD $0x9019a5c8; WORD $0xa56fa5b9  // .quad -6525815118631426616
	WORD $0xf4200f3a; WORD $0xcecb8f27  // .quad -3545582879861895366
	WORD $0xf8940985; WORD $0x813f3978  // .quad -9133518327554766459
	WORD $0x36b90be6; WORD $0xa18f07d7  // .quad -6805211891016070170
	WORD $0x04674edf; WORD $0xc9f2c9cd  // .quad -3894828845342699809
	WORD $0x45812297; WORD $0xfc6f7c40  // .quad -256850038250986857
	WORD $0x2b70b59e; WORD $0x9dc5ada8  // .quad -7078060301547948642
	WORD $0x364ce306; WORD $0xc5371912  // .quad -4235889358507547898
	WORD $0xc3e01bc7; WORD $0xf684df56  // .quad -683175679707046969
	WORD $0x3a6c115d; WORD $0x9a130b96  // .quad -7344513827457986211
	WORD $0xc90715b4; WORD $0xc097ce7b  // .quad -4568956265895094860
	WORD $0xbb48db21; WORD $0xf0bdc21a  // .quad -1099509313941480671
	WORD $0xb50d88f5; WORD $0x96769950  // .quad -7604722348854507275
	WORD $0xe250eb32; WORD $0xbc143fa4  // .quad -4894216917640746190
	WORD $0x1ae525fe; WORD $0xeb194f8e  // .quad -1506085128623544834
	WORD $0xd0cf37bf; WORD $0x92efd1b8  // .quad -7858832233030797377
	WORD $0x050305ae; WORD $0xb7abc627  // .quad -5211854272861108818
	WORD $0xc643c71a; WORD $0xe596b7b0  // .quad -1903131822648998118
	WORD $0x7bea5c70; WORD $0x8f7e32ce  // .quad -8106986416796705680
	WORD $0x1ae4f38c; WORD $0xb35dbf82  // .quad -5522047002568494196

TEXT ·__f32toa(SB), NOSPLIT, $0-24
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $96, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_f32toa:
	MOVD out+0(FP), R0
	FMOVD val+8(FP), F0
	MOVD ·_subr__f32toa(SB), R11
	WORD $0x1000005e // adr x30, .+8
	JMP (R11)
	MOVD R0, ret+16(FP)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
