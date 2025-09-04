    WORD $0x4e02011f // tbl    v31.16b, { v8.16b }, v2.16b
    WORD $0x2a2703f8 // mvn    w24, w7
    WORD $0x4e0203de // tbl    v30.16b, { v30.16b }, v2.16b
    WORD $0x1a9f97f5 // cset    w21, hi
    WORD $0x4e0203bd // tbl    v29.16b, { v29.16b }, v2.16b
    WORD $0x4e02039c // tbl    v28.16b, { v28.16b }, v2.16b
    WORD $0x4e71bbff // addv    h31, v31.8h
    WORD $0x4e71bbde // addv    h30, v30.8h
    WORD $0x4e71bbbd // addv    h29, v29.8h
    WORD $0x4e71bb9c // addv    h28, v28.8h
    WORD $0x1e2603d6 // fmov    w22, s30
    WORD $0x1e2603b7 // fmov    w23, s29
    WORD $0x1e2603f4 // fmov    w20, s31
    WORD $0x1e260393 // fmov    w19, s28
    WORD $0x33103ed4 // bfi    w20, w22, #16, #16
    WORD $0x33103ef3 // bfi    w19, w23, #16, #16
    WORD $0x8a140316 // and    x22, x24, x20
    WORD $0xea130317 // ands    x23, x24, x19
    BEQ LBB0_464
LBB0_462:
    WORD $0xd10006f8 // sub    x24, x23, #1
    WORD $0x8a160319 // and    x25, x24, x22
    WORD $0x9e67033c // fmov    d28, x25
    WORD $0x0e205b9c // cnt    v28.8b, v28.8b
    WORD $0x2e303b9c // uaddlv    h28, v28.8b
    WORD $0x1e260399 // fmov    w25, s28
    WORD $0x8b060339 // add    x25, x25, x6
    WORD $0xeb02033f // cmp    x25, x2
    BLS LBB0_468
    WORD $0xea170317 // ands    x23, x24, x23
    WORD $0x91000442 // add    x2, x2, #1
    BNE LBB0_462
LBB0_464:
    WORD $0x9e6702dc // fmov    d28, x22
    WORD $0xd2407eb5 // eor    x21, x21, #0xffffffff
    WORD $0x8a1402b4 // and    x20, x21, x20
    WORD $0xea1302b3 // ands    x19, x21, x19
    WORD $0x0e205b9c // cnt    v28.8b, v28.8b
    WORD $0x2e303b9c // uaddlv    h28, v28.8b
    WORD $0x1e260396 // fmov    w22, s28
    WORD $0x8b0602c6 // add    x6, x22, x6
    BEQ LBB0_467
LBB0_465:
    WORD $0xd1000675 // sub    x21, x19, #1
    WORD $0x8a1402b6 // and    x22, x21, x20
    WORD $0x9e6702dc // fmov    d28, x22
    WORD $0x0e205b9c // cnt    v28.8b, v28.8b
    WORD $0x2e303b9c // uaddlv    h28, v28.8b
    WORD $0x1e260396 // fmov    w22, s28
    WORD $0x8b0602d6 // add    x22, x22, x6
    WORD $0xeb0202df // cmp    x22, x2
    BLS LBB0_470
    WORD $0xea1302b3 // ands    x19, x21, x19
    WORD $0x91000442 // add    x2, x2, #1
    BNE LBB0_465
LBB0_467:
    WORD $0x9e67029c // fmov    d28, x20
    WORD $0x937ffce7 // asr    x7, x7, #63
    WORD $0x91010210 // add    x16, x16, #64
    WORD $0x0e205b9c // cnt    v28.8b, v28.8b
    WORD $0x2e303b9c // uaddlv    h28, v28.8b
    WORD $0x1e260385 // fmov    w5, s28
    WORD $0x8b0600a6 // add    x6, x5, x6
    WORD $0xaa1203e5 // mov    x5, x18
    WORD $0xf1010252 // subs    x18, x18, #64
    BGE LBB0_458
    B LBB0_442
LBB0_468:
    WORD $0xf9400410 // ldr    x16, [x0, #8]
    WORD $0xdac002f2 // rbit    x18, x23
    B LBB0_471
LBB0_469:
    WORD $0x5ac00290 // rbit    w16, w20
    WORD $0x8b040052 // add    x18, x2, x4
    WORD $0x5ac01210 // clz    w16, w16
    WORD $0x8b100250 // add    x16, x18, x16
    WORD $0x91000a12 // add    x18, x16, #2
    WORD $0xf9000032 // str    x18, [x1]
    B LBB0_405
LBB0_470:
    WORD $0xf9400410 // ldr    x16, [x0, #8]
    WORD $0xdac00272 // rbit    x18, x19
LBB0_471:
    WORD $0xdac01252 // clz    x18, x18
    WORD $0xcb050252 // sub    x18, x18, x5
    WORD $0x8b100250 // add    x16, x18, x16
    WORD $0x91000612 // add    x18, x16, #1
    WORD $0xf9000032 // str    x18, [x1]
    WORD $0xf9400402 // ldr    x2, [x0, #8]
    WORD $0xeb02025f // cmp    x18, x2
    WORD $0x9a902452 // csinc    x18, x2, x16, hs
    WORD $0xf9000032 // str    x18, [x1]
    B LBB0_405
LBB0_472:
    WORD $0x92800025 // mov    x5, #-2
    WORD $0x52800046 // mov    w6, #2
    WORD $0x8b060042 // add    x2, x2, x6
    WORD $0xab0400a4 // adds    x4, x5, x4
    BLE LBB0_405
LBB0_473:
    WORD $0x39400045 // ldrb    w5, [x2]
    WORD $0x710170bf // cmp    w5, #92
    BEQ LBB0_472
    WORD $0x710088bf // cmp    w5, #34
    BEQ LBB0_476
    WORD $0x92800005 // mov    x5, #-1
    WORD $0x52800026 // mov    w6, #1
    WORD $0x8b060042 // add    x2, x2, x6
    WORD $0xab0400a4 // adds    x4, x5, x4
    BGT LBB0_473
    B LBB0_405
LBB0_476:
    WORD $0xcb100050 // sub    x16, x2, x16
    WORD $0x91000612 // add    x18, x16, #1
    WORD $0xf9000032 // str    x18, [x1]
    B LBB0_405
LBB0_477:
    WORD $0x8b120202 // add    x2, x16, x18
    B LBB0_440
LBB0_478:
    WORD $0xf9400412 // ldr    x18, [x0, #8]
    WORD $0xf9000032 // str    x18, [x1]
    B LBB0_405
LBB0_479:
    WORD $0xd10004a6 // sub    x6, x5, #1
    WORD $0xeb0400df // cmp    x6, x4
    BEQ LBB0_405
    WORD $0x8b020202 // add    x2, x16, x2
    WORD $0x8b040042 // add    x2, x2, x4
    WORD $0xcb0400a4 // sub    x4, x5, x4
    WORD $0x91000842 // add    x2, x2, #2
    WORD $0xd1000884 // sub    x4, x4, #2
    B LBB0_440
LBB0_481:
    WORD $0x91004108 // add    x8, x8, #16
    WORD $0xeb0a011f // cmp    x8, x10
    BNE LBB0_2
LBB0_482:
    CMP $0, R3
    BEQ LBB0_786
    ADR LCPI0_10, R8
    ADR LCPI0_1, R17
    WORD $0x5280004a // mov    w10, #2
    WORD $0xd284c00d // mov    x13, #9728
    WORD $0x4f01e440 // movi    v0.16b, #34
    WORD $0xb20903f0 // mov    x16, #36028797027352576
    WORD $0x3dc00101 // ldr    q1, [x8, :lo12:.LCPI0_10]
    ADR LCPI0_0, R8
    WORD $0x4f01e5c3 // movi    v3.16b, #46
    WORD $0x528d8c35 // mov    w21, #27745
    WORD $0x4f01e565 // movi    v5.16b, #43
    WORD $0x528eadd8 // mov    w24, #30062
    WORD $0x3d800061 // str    q1, [x3]
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0xf9400009 // ldr    x9, [x0]
    WORD $0x9280000b // mov    x11, #-1
    WORD $0x4f02e781 // movi    v1.16b, #92
    WORD $0x5280002c // mov    w12, #1
    WORD $0x4f01e5a6 // movi    v6.16b, #45
    WORD $0xf2c0002d // movk    x13, #1, lsl #32
    WORD $0x4f06e607 // movi    v7.16b, #208
    WORD $0x5280006e // mov    w14, #3
    WORD $0x4f00e550 // movi    v16.16b, #10
    WORD $0x5280008f // mov    w15, #4
    WORD $0x4f06e7f1 // movi    v17.16b, #223
    WORD $0xf2800030 // movk    x16, #1
    WORD $0x4f02e4b2 // movi    v18.16b, #69
    WORD $0x3dc00102 // ldr    q2, [x8, :lo12:.LCPI0_0]
    WORD $0x3dc00224 // ldr    q4, [x17, :lo12:.LCPI0_1]
    WORD $0xaa2903f1 // mvn    x17, x9
    WORD $0xd1000532 // sub    x18, x9, #1
    WORD $0xcb0903e2 // neg    x2, x9
    WORD $0xcb090144 // sub    x4, x10, x9
    WORD $0x12800005 // mov    w5, #-1
    WORD $0x528000be // mov    w30, #5
    WORD $0x72acae75 // movk    w21, #25971, lsl #16
    WORD $0x72ad8d98 // movk    w24, #27756, lsl #16
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    WORD $0x528000d9 // mov    w25, #6
LBB0_484:
    WORD $0xf9400408 // ldr    x8, [x0, #8]
    WORD $0xf9400026 // ldr    x6, [x1]
    WORD $0xeb0800df // cmp    x6, x8
    BHS LBB0_488
    WORD $0x38666927 // ldrb    w7, [x9, x6]
    WORD $0x710034ff // cmp    w7, #13
    BEQ LBB0_488
    WORD $0x710080ff // cmp    w7, #32
    BEQ LBB0_488
    WORD $0xaa0603f6 // mov    x22, x6
    WORD $0x51002ce7 // sub    w7, w7, #11
    WORD $0x310008ff // cmn    w7, #2
    BLO LBB0_504
LBB0_488:
    WORD $0x910004d6 // add    x22, x6, #1
    WORD $0xeb0802df // cmp    x22, x8
    BHS LBB0_492
    WORD $0x38766927 // ldrb    w7, [x9, x22]
    WORD $0x710034ff // cmp    w7, #13
    BEQ LBB0_492
    WORD $0x710080ff // cmp    w7, #32
    BEQ LBB0_492
    WORD $0x51002ce7 // sub    w7, w7, #11
    WORD $0x310008ff // cmn    w7, #2
    BLO LBB0_504
LBB0_492:
    WORD $0x910008d6 // add    x22, x6, #2
    WORD $0xeb0802df // cmp    x22, x8
    BHS LBB0_496
    WORD $0x38766927 // ldrb    w7, [x9, x22]
    WORD $0x710034ff // cmp    w7, #13
    BEQ LBB0_496
    WORD $0x710080ff // cmp    w7, #32
    BEQ LBB0_496
    WORD $0x51002ce7 // sub    w7, w7, #11
    WORD $0x310008ff // cmn    w7, #2
    BLO LBB0_504
LBB0_496:
    WORD $0x91000cd6 // add    x22, x6, #3
    WORD $0xeb0802df // cmp    x22, x8
    BHS LBB0_500
    WORD $0x38766927 // ldrb    w7, [x9, x22]
    WORD $0x710034ff // cmp    w7, #13
    BEQ LBB0_500
    WORD $0x710080ff // cmp    w7, #32
    BEQ LBB0_500
    WORD $0x51002ce7 // sub    w7, w7, #11
    WORD $0x310008ff // cmn    w7, #2
    BLO LBB0_504
LBB0_500:
    WORD $0x910010d6 // add    x22, x6, #4
    WORD $0xeb0802df // cmp    x22, x8
    BHS LBB0_869
LBB0_501:
    WORD $0x38766926 // ldrb    w6, [x9, x22]
    WORD $0x9ac62187 // lsl    x7, x12, x6
    WORD $0x710080df // cmp    w6, #32
    WORD $0x8a0d00e6 // and    x6, x7, x13
    WORD $0xfa4098c4 // ccmp    x6, #0, #4, ls
    BEQ LBB0_503
    WORD $0x910006d6 // add    x22, x22, #1
    WORD $0xeb16011f // cmp    x8, x22
    BNE LBB0_501
    B LBB0_806
LBB0_503:
    WORD $0xeb0802df // cmp    x22, x8
    BHS LBB0_806
LBB0_504:
    WORD $0x910006c8 // add    x8, x22, #1
    WORD $0xf9000028 // str    x8, [x1]
    WORD $0x38766926 // ldrb    w6, [x9, x22]
    CMP $0, R6
    BEQ LBB0_806
    WORD $0xf9400073 // ldr    x19, [x3]
    WORD $0xb100057f // cmn    x11, #1
    WORD $0x9a8b02cb // csel    x11, x22, x11, eq
    WORD $0xd1000667 // sub    x7, x19, #1
    WORD $0x8b070c68 // add    x8, x3, x7, lsl #3
    WORD $0xb8408d14 // ldr    w20, [x8, #8]!
    WORD $0x71000e9f // cmp    w20, #3
    BGT LBB0_522
    WORD $0x7100069f // cmp    w20, #1
    BEQ LBB0_541
    WORD $0x71000a9f // cmp    w20, #2
    BEQ LBB0_543
    WORD $0x71000e9f // cmp    w20, #3
    BNE LBB0_551
    WORD $0x710088df // cmp    w6, #34
    BNE LBB0_951
    WORD $0xf900010f // str    x15, [x8]
    WORD $0xf9400037 // ldr    x23, [x1]
    WORD $0xf9400416 // ldr    x22, [x0, #8]
    WORD $0xeb1702c8 // subs    x8, x22, x23
    BEQ LBB0_957
    WORD $0xf101011f // cmp    x8, #64
    BLO LBB0_722
    WORD $0xaa1f03f9 // mov    x25, xzr
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xaa1703fa // mov    x26, x23
LBB0_513:
    WORD $0x8b1a0126 // add    x6, x9, x26
    WORD $0xad4050d3 // ldp    q19, q20, [x6]
    WORD $0x6e208e77 // cmeq    v23.16b, v19.16b, v0.16b
    WORD $0x6e218e73 // cmeq    v19.16b, v19.16b, v1.16b
    WORD $0xad4158d5 // ldp    q21, q22, [x6, #32]
    WORD $0x6e208e98 // cmeq    v24.16b, v20.16b, v0.16b
    WORD $0x6e218e94 // cmeq    v20.16b, v20.16b, v1.16b
    WORD $0x4e221f18 // and    v24.16b, v24.16b, v2.16b
    WORD $0x6e208eb9 // cmeq    v25.16b, v21.16b, v0.16b
    WORD $0x6e218eb5 // cmeq    v21.16b, v21.16b, v1.16b
    WORD $0x6e208eda // cmeq    v26.16b, v22.16b, v0.16b
    WORD $0x4e221f39 // and    v25.16b, v25.16b, v2.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e221f5a // and    v26.16b, v26.16b, v2.16b
    WORD $0x4e040339 // tbl    v25.16b, { v25.16b }, v4.16b
    WORD $0x6e218ed6 // cmeq    v22.16b, v22.16b, v1.16b
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e04035a // tbl    v26.16b, { v26.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e221ef7 // and    v23.16b, v23.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e040318 // tbl    v24.16b, { v24.16b }, v4.16b
    WORD $0x4e71bb39 // addv    h25, v25.8h
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402f7 // tbl    v23.16b, { v23.16b }, v4.16b
    WORD $0x4e71bb5a // addv    h26, v26.8h
    WORD $0x4e71bab5 // addv    h21, v21.8h
    WORD $0x1e260327 // fmov    w7, s25
    WORD $0x1e260294 // fmov    w20, s20
    WORD $0x1e260353 // fmov    w19, s26
    WORD $0x4e71bad4 // addv    h20, v22.8h
    WORD $0x1e2602b5 // fmov    w21, s21
    WORD $0x4e71bb18 // addv    h24, v24.8h
    WORD $0xd3607ce7 // lsl    x7, x7, #32
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0xaa13c0e7 // orr    x7, x7, x19, lsl #48
    WORD $0x4e71baf5 // addv    h21, v23.8h
    WORD $0xd3607eb5 // lsl    x21, x21, #32
    WORD $0x1e26029b // fmov    w27, s20
    WORD $0x53103e94 // lsl    w20, w20, #16
    WORD $0x1e260306 // fmov    w6, s24
    WORD $0x1e260273 // fmov    w19, s19
    WORD $0x1e2602bd // fmov    w29, s21
    WORD $0xaa1bc2b5 // orr    x21, x21, x27, lsl #48
    WORD $0x53103cc6 // lsl    w6, w6, #16
    WORD $0xaa1302b3 // orr    x19, x21, x19
    WORD $0xaa1d00f5 // orr    x21, x7, x29
    WORD $0xaa140267 // orr    x7, x19, x20
    WORD $0xaa0602a6 // orr    x6, x21, x6
    CMP $0, R7
    BNE LBB0_517
    CMP $0, R25
    BNE LBB0_519
    CMP $0, R6
    BNE LBB0_520
LBB0_516:
    WORD $0xd1010108 // sub    x8, x8, #64
    WORD $0x9101035a // add    x26, x26, #64
    WORD $0xf100fd1f // cmp    x8, #63
    BHI LBB0_513
    B LBB0_702
LBB0_517:
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_519
    WORD $0xdac000f3 // rbit    x19, x7
    WORD $0xdac01273 // clz    x19, x19
    WORD $0x8b1a0278 // add    x24, x19, x26
LBB0_519:
    WORD $0x8a3900f3 // bic    x19, x7, x25
    WORD $0xaa130734 // orr    x20, x25, x19, lsl #1
    WORD $0x8a3400e7 // bic    x7, x7, x20
    WORD $0x9201f0e7 // and    x7, x7, #0xaaaaaaaaaaaaaaaa
    WORD $0xab1300e7 // adds    x7, x7, x19
    WORD $0xd37ff8e7 // lsl    x7, x7, #1
    WORD $0x1a9f37f9 // cset    w25, hs
    WORD $0xd200f0e7 // eor    x7, x7, #0x5555555555555555
    WORD $0x8a1400e7 // and    x7, x7, x20
    WORD $0x8a2700c6 // bic    x6, x6, x7
    CMP $0, R6
    BEQ LBB0_516
LBB0_520:
    WORD $0xdac000c8 // rbit    x8, x6
    WORD $0xdac01108 // clz    x8, x8
    WORD $0x8b1a0108 // add    x8, x8, x26
    WORD $0x91000508 // add    x8, x8, #1
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BNE LBB0_924
LBB0_521:
    WORD $0x528d8c35 // mov    w21, #27745
    WORD $0x528eadd8 // mov    w24, #30062
    WORD $0xf9000028 // str    x8, [x1]
    WORD $0xf10006e8 // subs    x8, x23, #1
    WORD $0x72acae75 // movk    w21, #25971, lsl #16
    WORD $0x72ad8d98 // movk    w24, #27756, lsl #16
    BGE LBB0_618
    B LBB0_868
LBB0_522:
    WORD $0x7100129f // cmp    w20, #4
    BEQ LBB0_547
    WORD $0x7100169f // cmp    w20, #5
    BEQ LBB0_549
    WORD $0x71001a9f // cmp    w20, #6
    BNE LBB0_551
    WORD $0x7101f4df // cmp    w6, #125
    BEQ LBB0_550
    WORD $0x710088df // cmp    w6, #34
    BNE LBB0_951
    WORD $0xf900010a // str    x10, [x8]
    WORD $0xf9400037 // ldr    x23, [x1]
    WORD $0xf9400416 // ldr    x22, [x0, #8]
    WORD $0xeb1702c8 // subs    x8, x22, x23
    BEQ LBB0_957
    WORD $0xf101011f // cmp    x8, #64
    BLO LBB0_727
    WORD $0xaa1f03f9 // mov    x25, xzr
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xaa1703fa // mov    x26, x23
LBB0_530:
    WORD $0x8b1a0126 // add    x6, x9, x26
    WORD $0xad4050d3 // ldp    q19, q20, [x6]
    WORD $0x6e208e77 // cmeq    v23.16b, v19.16b, v0.16b
    WORD $0x6e218e73 // cmeq    v19.16b, v19.16b, v1.16b
    WORD $0xad4158d5 // ldp    q21, q22, [x6, #32]
    WORD $0x6e208e98 // cmeq    v24.16b, v20.16b, v0.16b
    WORD $0x6e218e94 // cmeq    v20.16b, v20.16b, v1.16b
    WORD $0x4e221f18 // and    v24.16b, v24.16b, v2.16b
    WORD $0x6e208eb9 // cmeq    v25.16b, v21.16b, v0.16b
    WORD $0x6e218eb5 // cmeq    v21.16b, v21.16b, v1.16b
    WORD $0x6e208eda // cmeq    v26.16b, v22.16b, v0.16b
    WORD $0x4e221f39 // and    v25.16b, v25.16b, v2.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e221f5a // and    v26.16b, v26.16b, v2.16b
    WORD $0x4e040339 // tbl    v25.16b, { v25.16b }, v4.16b
    WORD $0x6e218ed6 // cmeq    v22.16b, v22.16b, v1.16b
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e04035a // tbl    v26.16b, { v26.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e221ef7 // and    v23.16b, v23.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e040318 // tbl    v24.16b, { v24.16b }, v4.16b
    WORD $0x4e71bb39 // addv    h25, v25.8h
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402f7 // tbl    v23.16b, { v23.16b }, v4.16b
    WORD $0x4e71bb5a // addv    h26, v26.8h
    WORD $0x4e71bab5 // addv    h21, v21.8h
    WORD $0x1e260327 // fmov    w7, s25
    WORD $0x1e260294 // fmov    w20, s20
    WORD $0x1e260353 // fmov    w19, s26
    WORD $0x4e71bad4 // addv    h20, v22.8h
    WORD $0x1e2602b5 // fmov    w21, s21
    WORD $0x4e71bb18 // addv    h24, v24.8h
    WORD $0xd3607ce7 // lsl    x7, x7, #32
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0xaa13c0e7 // orr    x7, x7, x19, lsl #48
    WORD $0x4e71baf5 // addv    h21, v23.8h
    WORD $0xd3607eb5 // lsl    x21, x21, #32
    WORD $0x1e26029b // fmov    w27, s20
    WORD $0x53103e94 // lsl    w20, w20, #16
    WORD $0x1e260306 // fmov    w6, s24
    WORD $0x1e260273 // fmov    w19, s19
    WORD $0x1e2602bd // fmov    w29, s21
    WORD $0xaa1bc2b5 // orr    x21, x21, x27, lsl #48
    WORD $0x53103cc6 // lsl    w6, w6, #16
    WORD $0xaa1302b3 // orr    x19, x21, x19
    WORD $0xaa1d00f5 // orr    x21, x7, x29
    WORD $0xaa140267 // orr    x7, x19, x20
    WORD $0xaa0602a6 // orr    x6, x21, x6
    CMP $0, R7
    BNE LBB0_534
    CMP $0, R25
    BNE LBB0_536
    CMP $0, R6
    BNE LBB0_537
LBB0_533:
    WORD $0xd1010108 // sub    x8, x8, #64
    WORD $0x9101035a // add    x26, x26, #64
    WORD $0xf100fd1f // cmp    x8, #63
    BHI LBB0_530
    B LBB0_708
LBB0_534:
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_536
    WORD $0xdac000f3 // rbit    x19, x7
    WORD $0xdac01273 // clz    x19, x19
    WORD $0x8b1a0278 // add    x24, x19, x26
LBB0_536:
    WORD $0x8a3900f3 // bic    x19, x7, x25
    WORD $0xaa130734 // orr    x20, x25, x19, lsl #1
    WORD $0x8a3400e7 // bic    x7, x7, x20
    WORD $0x9201f0e7 // and    x7, x7, #0xaaaaaaaaaaaaaaaa
    WORD $0xab1300e7 // adds    x7, x7, x19
    WORD $0xd37ff8e7 // lsl    x7, x7, #1
    WORD $0x1a9f37f9 // cset    w25, hs
    WORD $0xd200f0e7 // eor    x7, x7, #0x5555555555555555
    WORD $0x8a1400e7 // and    x7, x7, x20
    WORD $0x8a2700c6 // bic    x6, x6, x7
    CMP $0, R6
    BEQ LBB0_533
LBB0_537:
    WORD $0xdac000c8 // rbit    x8, x6
    WORD $0xdac01108 // clz    x8, x8
    WORD $0x8b1a0108 // add    x8, x8, x26
    WORD $0x91000508 // add    x8, x8, #1
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BNE LBB0_924
LBB0_538:
    WORD $0xf9000028 // str    x8, [x1]
    WORD $0xf10006e8 // subs    x8, x23, #1
    BLT LBB0_868
    WORD $0xf9400068 // ldr    x8, [x3]
    WORD $0xf13ffd1f // cmp    x8, #4095
    BGT LBB0_923
    WORD $0x91000506 // add    x6, x8, #1
    WORD $0x8b080c68 // add    x8, x3, x8, lsl #3
    WORD $0x528d8c35 // mov    w21, #27745
    WORD $0x528eadd8 // mov    w24, #30062
    WORD $0x72acae75 // movk    w21, #25971, lsl #16
    WORD $0x72ad8d98 // movk    w24, #27756, lsl #16
    WORD $0xf9000066 // str    x6, [x3]
    WORD $0xf900050f // str    x15, [x8, #8]
    B LBB0_618
LBB0_541:
    WORD $0x7100b0df // cmp    w6, #44
    BEQ LBB0_559
    WORD $0x710174df // cmp    w6, #93
    BEQ LBB0_550
    B LBB0_951
LBB0_543:
    WORD $0x7101f4df // cmp    w6, #125
    BEQ LBB0_550
    WORD $0x7100b0df // cmp    w6, #44
    BNE LBB0_951
    WORD $0xf13ffe7f // cmp    x19, #4095
    BGT LBB0_923
    WORD $0x91000668 // add    x8, x19, #1
    WORD $0x8b130c66 // add    x6, x3, x19, lsl #3
    WORD $0xf9000068 // str    x8, [x3]
    WORD $0xf90004ce // str    x14, [x6, #8]
    B LBB0_618
LBB0_547:
    WORD $0x7100e8df // cmp    w6, #58
    BNE LBB0_951
    WORD $0xf900011f // str    xzr, [x8]
    B LBB0_618
LBB0_549:
    WORD $0x710174df // cmp    w6, #93
    BNE LBB0_552
LBB0_550:
    WORD $0xaa0b03e8 // mov    x8, x11
    WORD $0xf9000067 // str    x7, [x3]
    CMP $0, R7
    BNE LBB0_484
    B LBB0_868
LBB0_551:
    WORD $0xf9000067 // str    x7, [x3]
    B LBB0_553
LBB0_552:
    WORD $0xf900010c // str    x12, [x8]
LBB0_553:
    WORD $0x92800028 // mov    x8, #-2
    WORD $0x710168df // cmp    w6, #90
    BGT LBB0_561
    WORD $0x5100c0c7 // sub    w7, w6, #48
    WORD $0x710028ff // cmp    w7, #10
    BHS LBB0_619
    WORD $0xf9400036 // ldr    x22, [x1]
    WORD $0xf9400406 // ldr    x6, [x0, #8]
    WORD $0xd10006c8 // sub    x8, x22, #1
    WORD $0xeb0800d7 // subs    x23, x6, x8
    BEQ LBB0_927
    WORD $0x8b080126 // add    x6, x9, x8
    WORD $0x394000c7 // ldrb    w7, [x6]
    WORD $0x7100c0ff // cmp    w7, #48
    BNE LBB0_569
    WORD $0xf10006ff // cmp    x23, #1
    BNE LBB0_567
LBB0_558:
    WORD $0x52800027 // mov    w7, #1
    B LBB0_617
LBB0_559:
    WORD $0xf13ffe7f // cmp    x19, #4095
    BGT LBB0_923
    WORD $0x91000668 // add    x8, x19, #1
    WORD $0x8b130c66 // add    x6, x3, x19, lsl #3
    WORD $0xf9000068 // str    x8, [x3]
    WORD $0xf90004df // str    xzr, [x6, #8]
    B LBB0_618
LBB0_561:
    WORD $0x7101b4df // cmp    w6, #109
    BLE LBB0_625
    WORD $0x7101b8df // cmp    w6, #110
    BEQ LBB0_630
    WORD $0x7101d0df // cmp    w6, #116
    BEQ LBB0_633
    WORD $0x7101ecdf // cmp    w6, #123
    BNE LBB0_868
    WORD $0xf9400068 // ldr    x8, [x3]
    WORD $0xf13ffd1f // cmp    x8, #4095
    BGT LBB0_923
    WORD $0x91000506 // add    x6, x8, #1
    WORD $0x8b080c68 // add    x8, x3, x8, lsl #3
    WORD $0xf9000066 // str    x6, [x3]
    WORD $0xf9000519 // str    x25, [x8, #8]
    B LBB0_618
LBB0_567:
    WORD $0x38766927 // ldrb    w7, [x9, x22]
    WORD $0x5100b8e7 // sub    w7, w7, #46
    WORD $0x7100dcff // cmp    w7, #55
    BHI LBB0_558
    WORD $0x9ac72193 // lsl    x19, x12, x7
    WORD $0x52800027 // mov    w7, #1
    WORD $0xea10027f // tst    x19, x16
    BEQ LBB0_617
LBB0_569:
    WORD $0xf10042ff // cmp    x23, #16
    BLO LBB0_723
    WORD $0xaa1f03fe // mov    x30, xzr
    WORD $0x9280001a // mov    x26, #-1
    WORD $0x92800019 // mov    x25, #-1
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xaa1703fb // mov    x27, x23
LBB0_571:
    WORD $0x3cfe68d3 // ldr    q19, [x6, x30]
    WORD $0x6e238e74 // cmeq    v20.16b, v19.16b, v3.16b
    WORD $0x6e258e75 // cmeq    v21.16b, v19.16b, v5.16b
    WORD $0x6e268e76 // cmeq    v22.16b, v19.16b, v6.16b
    WORD $0x4e278677 // add    v23.16b, v19.16b, v7.16b
    WORD $0x4e311e73 // and    v19.16b, v19.16b, v17.16b
    WORD $0x6e328e73 // cmeq    v19.16b, v19.16b, v18.16b
    WORD $0x6e373617 // cmhi    v23.16b, v16.16b, v23.16b
    WORD $0x4eb51ed5 // orr    v21.16b, v22.16b, v21.16b
    WORD $0x4eb31e96 // orr    v22.16b, v20.16b, v19.16b
    WORD $0x4eb61ef6 // orr    v22.16b, v23.16b, v22.16b
    WORD $0x4eb51ed6 // orr    v22.16b, v22.16b, v21.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e71bad6 // addv    h22, v22.8h
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x1e2602c7 // fmov    w7, s22
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x2a2703e7 // mvn    w7, w7
    WORD $0x4e71bab5 // addv    h21, v21.8h
    WORD $0x32103ce7 // orr    w7, w7, #0xffff0000
    WORD $0x5ac000e7 // rbit    w7, w7
    WORD $0x1e260294 // fmov    w20, s20
    WORD $0x1e260273 // fmov    w19, s19
    WORD $0x5ac010e7 // clz    w7, w7
    WORD $0x1e2602bd // fmov    w29, s21
    WORD $0x710040ff // cmp    w7, #16
    BEQ LBB0_573
    WORD $0x1ac720b5 // lsl    w21, w5, w7
    WORD $0x0a350294 // bic    w20, w20, w21
    WORD $0x0a350273 // bic    w19, w19, w21
    WORD $0x0a3503bd // bic    w29, w29, w21
LBB0_573:
    WORD $0x51000695 // sub    w21, w20, #1
    WORD $0x6a1402b5 // ands    w21, w21, w20
    BNE LBB0_707
    WORD $0x51000675 // sub    w21, w19, #1
    WORD $0x6a1302b5 // ands    w21, w21, w19
    BNE LBB0_707
    WORD $0x510007b5 // sub    w21, w29, #1
    WORD $0x6a1d02b5 // ands    w21, w21, w29
    BNE LBB0_707
    CMP $0, R20
    BEQ LBB0_579
    WORD $0x5ac00294 // rbit    w20, w20
    WORD $0xb100071f // cmn    x24, #1
    WORD $0x5ac01294 // clz    w20, w20
    BNE LBB0_715
    WORD $0x8b1403d8 // add    x24, x30, x20
LBB0_579:
    CMP $0, R19
    BEQ LBB0_582
    WORD $0x5ac00273 // rbit    w19, w19
    WORD $0xb100073f // cmn    x25, #1
    WORD $0x5ac01273 // clz    w19, w19
    BNE LBB0_713
    WORD $0x8b1303d9 // add    x25, x30, x19
LBB0_582:
    CMP $0, R29
    BEQ LBB0_585
    WORD $0x5ac003b3 // rbit    w19, w29
    WORD $0xb100075f // cmn    x26, #1
    WORD $0x5ac01273 // clz    w19, w19
    BNE LBB0_713
    WORD $0x8b1303da // add    x26, x30, x19
LBB0_585:
    WORD $0x710040ff // cmp    w7, #16
    BNE LBB0_603
    WORD $0xd100437b // sub    x27, x27, #16
    WORD $0x910043de // add    x30, x30, #16
    WORD $0xf1003f7f // cmp    x27, #15
    BHI LBB0_571
    WORD $0x8b1e00dd // add    x29, x6, x30
    WORD $0xeb1e02ff // cmp    x23, x30
    BEQ LBB0_604
LBB0_588:
    WORD $0x8b1d0093 // add    x19, x4, x29
    WORD $0x8b1b03a7 // add    x7, x29, x27
    WORD $0xcb16027e // sub    x30, x19, x22
    WORD $0xaa1d03f7 // mov    x23, x29
    B LBB0_590
LBB0_589:
    WORD $0xd100077b // sub    x27, x27, #1
    WORD $0x910007de // add    x30, x30, #1
    WORD $0xaa1703fd // mov    x29, x23
    CMP $0, R27
    BEQ LBB0_651
LBB0_590:
    WORD $0x384016f3 // ldrb    w19, [x23], #1
    WORD $0x5100c274 // sub    w20, w19, #48
    WORD $0x71002a9f // cmp    w20, #10
    BLO LBB0_589
    WORD $0x7100b67f // cmp    w19, #45
    BLE LBB0_597
    WORD $0x7101967f // cmp    w19, #101
    BEQ LBB0_601
    WORD $0x7101167f // cmp    w19, #69
    BEQ LBB0_601
    WORD $0x7100ba7f // cmp    w19, #46
    BNE LBB0_604
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_701
    WORD $0xd10007d8 // sub    x24, x30, #1
    B LBB0_589
LBB0_597:
    WORD $0x7100ae7f // cmp    w19, #43
    BEQ LBB0_599
    WORD $0x7100b67f // cmp    w19, #45
    BNE LBB0_604
LBB0_599:
    WORD $0xb100075f // cmn    x26, #1
    BNE LBB0_701
    WORD $0xd10007da // sub    x26, x30, #1
    B LBB0_589
LBB0_601:
    WORD $0xb100073f // cmn    x25, #1
    BNE LBB0_701
    WORD $0xd10007d9 // sub    x25, x30, #1
    B LBB0_589
LBB0_603:
    WORD $0x8b2740c7 // add    x7, x6, w7, uxtw
    WORD $0x8b1e00fd // add    x29, x7, x30
LBB0_604:
    WORD $0x92800007 // mov    x7, #-1
    CMP $0, R24
    BEQ LBB0_933
LBB0_605:
    WORD $0x528000be // mov    w30, #5
    CMP $0, R26
    BEQ LBB0_933
    CMP $0, R25
    BEQ LBB0_933
    WORD $0xcb0603a6 // sub    x6, x29, x6
    WORD $0xd10004c7 // sub    x7, x6, #1
    WORD $0xeb07031f // cmp    x24, x7
    BEQ LBB0_615
    WORD $0xeb07035f // cmp    x26, x7
    BEQ LBB0_615
    WORD $0xeb07033f // cmp    x25, x7
    BEQ LBB0_615
    WORD $0xf1000747 // subs    x7, x26, #1
    BLT LBB0_612
    WORD $0xeb07033f // cmp    x25, x7
    BNE LBB0_930
LBB0_612:
    WORD $0xaa190307 // orr    x7, x24, x25
    TST $(1<<63), R7
    BNE LBB0_614
    WORD $0xeb19031f // cmp    x24, x25
    BGE LBB0_932
LBB0_614:
    WORD $0xd1000733 // sub    x19, x25, #1
    WORD $0xd37ffce7 // lsr    x7, x7, #63
    WORD $0xeb13031f // cmp    x24, x19
    WORD $0x520000e7 // eor    w7, w7, #0x1
    WORD $0x1a9f17f3 // cset    w19, eq
    WORD $0x6a1300ff // tst    w7, w19
    WORD $0xda9900c7 // csinv    x7, x6, x25, eq
    B LBB0_616
LBB0_615:
    WORD $0xcb0603e7 // neg    x7, x6
LBB0_616:
    TST $(1<<63), R7
    BNE LBB0_933
LBB0_617:
    WORD $0x528d8c35 // mov    w21, #27745
    WORD $0x528eadd8 // mov    w24, #30062
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x8b0800e6 // add    x6, x7, x8
    WORD $0x72acae75 // movk    w21, #25971, lsl #16
    WORD $0x72ad8d98 // movk    w24, #27756, lsl #16
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    WORD $0x528000d9 // mov    w25, #6
    WORD $0xf9000026 // str    x6, [x1]
    TST $(1<<63), R8
    BNE LBB0_868
LBB0_618:
    WORD $0xf9400066 // ldr    x6, [x3]
    WORD $0xaa0b03e8 // mov    x8, x11
    CMP $0, R6
    BNE LBB0_484
    B LBB0_868
LBB0_619:
    WORD $0x710088df // cmp    w6, #34
    BEQ LBB0_637
    WORD $0x7100b4df // cmp    w6, #45
    BNE LBB0_868
    WORD $0xf9400406 // ldr    x6, [x0, #8]
    WORD $0xf9400028 // ldr    x8, [x1]
    WORD $0xeb0800db // subs    x27, x6, x8
    BEQ LBB0_936
    WORD $0x8b080136 // add    x22, x9, x8
    WORD $0x394002c6 // ldrb    w6, [x22]
    WORD $0x7100c0df // cmp    w6, #48
    BNE LBB0_654
    WORD $0xf100077f // cmp    x27, #1
    BNE LBB0_652
LBB0_624:
    WORD $0x5280003d // mov    w29, #1
    B LBB0_700
LBB0_625:
    WORD $0x71016cdf // cmp    w6, #91
    BEQ LBB0_649
    WORD $0x710198df // cmp    w6, #102
    BNE LBB0_868
    WORD $0xf9400406 // ldr    x6, [x0, #8]
    WORD $0xf9400028 // ldr    x8, [x1]
    WORD $0xd10010c7 // sub    x7, x6, #4
    WORD $0xeb07011f // cmp    x8, x7
    BHI LBB0_931
    WORD $0xb8686926 // ldr    w6, [x9, x8]
    WORD $0x6b1500df // cmp    w6, w21
    BNE LBB0_937
    WORD $0x91001106 // add    x6, x8, #4
    B LBB0_636
LBB0_630:
    WORD $0xf9400406 // ldr    x6, [x0, #8]
    WORD $0xf9400028 // ldr    x8, [x1]
    WORD $0xd1000cc7 // sub    x7, x6, #3
    WORD $0xeb07011f // cmp    x8, x7
    BHI LBB0_931
    WORD $0xb8686a46 // ldr    w6, [x18, x8]
    WORD $0x6b1800df // cmp    w6, w24
    BNE LBB0_942
    WORD $0x91000d06 // add    x6, x8, #3
    WORD $0xf100011f // cmp    x8, #0
    WORD $0xf9000026 // str    x6, [x1]
    BGT LBB0_618
    B LBB0_929
LBB0_633:
    WORD $0xf9400406 // ldr    x6, [x0, #8]
    WORD $0xf9400028 // ldr    x8, [x1]
    WORD $0xd1000cc7 // sub    x7, x6, #3
    WORD $0xeb07011f // cmp    x8, x7
    BHI LBB0_931
    WORD $0xb8686a46 // ldr    w6, [x18, x8]
    WORD $0x6b1a00df // cmp    w6, w26
    BNE LBB0_946
    WORD $0x91000d06 // add    x6, x8, #3
LBB0_636:
    WORD $0xf100011f // cmp    x8, #0
    WORD $0xf9000026 // str    x6, [x1]
    BGT LBB0_618
    B LBB0_929
LBB0_637:
    WORD $0xf9400037 // ldr    x23, [x1]
    WORD $0xf9400416 // ldr    x22, [x0, #8]
    WORD $0xeb1702c8 // subs    x8, x22, x23
    BEQ LBB0_957
    WORD $0xf101011f // cmp    x8, #64
    BLO LBB0_728
    WORD $0xaa1f03f9 // mov    x25, xzr
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xaa1703fa // mov    x26, x23
LBB0_640:
    WORD $0x8b1a0126 // add    x6, x9, x26
    WORD $0xad4050d3 // ldp    q19, q20, [x6]
    WORD $0x6e208e77 // cmeq    v23.16b, v19.16b, v0.16b
    WORD $0x6e218e73 // cmeq    v19.16b, v19.16b, v1.16b
    WORD $0xad4158d5 // ldp    q21, q22, [x6, #32]
    WORD $0x6e208e98 // cmeq    v24.16b, v20.16b, v0.16b
    WORD $0x6e218e94 // cmeq    v20.16b, v20.16b, v1.16b
    WORD $0x4e221f18 // and    v24.16b, v24.16b, v2.16b
    WORD $0x6e208eb9 // cmeq    v25.16b, v21.16b, v0.16b
    WORD $0x6e218eb5 // cmeq    v21.16b, v21.16b, v1.16b
    WORD $0x6e208eda // cmeq    v26.16b, v22.16b, v0.16b
    WORD $0x4e221f39 // and    v25.16b, v25.16b, v2.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e221f5a // and    v26.16b, v26.16b, v2.16b
    WORD $0x4e040339 // tbl    v25.16b, { v25.16b }, v4.16b
    WORD $0x6e218ed6 // cmeq    v22.16b, v22.16b, v1.16b
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e04035a // tbl    v26.16b, { v26.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e221ef7 // and    v23.16b, v23.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e040318 // tbl    v24.16b, { v24.16b }, v4.16b
    WORD $0x4e71bb39 // addv    h25, v25.8h
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402f7 // tbl    v23.16b, { v23.16b }, v4.16b
    WORD $0x4e71bb5a // addv    h26, v26.8h
    WORD $0x4e71bab5 // addv    h21, v21.8h
    WORD $0x1e260327 // fmov    w7, s25
    WORD $0x1e260294 // fmov    w20, s20
    WORD $0x1e260353 // fmov    w19, s26
    WORD $0x4e71bad4 // addv    h20, v22.8h
    WORD $0x1e2602b5 // fmov    w21, s21
    WORD $0x4e71bb18 // addv    h24, v24.8h
    WORD $0xd3607ce7 // lsl    x7, x7, #32
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0xaa13c0e7 // orr    x7, x7, x19, lsl #48
    WORD $0x4e71baf5 // addv    h21, v23.8h
    WORD $0xd3607eb5 // lsl    x21, x21, #32
    WORD $0x1e26029b // fmov    w27, s20
    WORD $0x53103e94 // lsl    w20, w20, #16
    WORD $0x1e260306 // fmov    w6, s24
    WORD $0x1e260273 // fmov    w19, s19
    WORD $0x1e2602bd // fmov    w29, s21
    WORD $0xaa1bc2b5 // orr    x21, x21, x27, lsl #48
    WORD $0x53103cc6 // lsl    w6, w6, #16
    WORD $0xaa1302b3 // orr    x19, x21, x19
    WORD $0xaa1d00f5 // orr    x21, x7, x29
    WORD $0xaa140267 // orr    x7, x19, x20
    WORD $0xaa0602a6 // orr    x6, x21, x6
    CMP $0, R7
    BNE LBB0_644
    CMP $0, R25
    BNE LBB0_646
    CMP $0, R6
    BNE LBB0_647
LBB0_643:
    WORD $0xd1010108 // sub    x8, x8, #64
    WORD $0x9101035a // add    x26, x26, #64
    WORD $0xf100fd1f // cmp    x8, #63
    BHI LBB0_640
    B LBB0_717
LBB0_644:
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_646
    WORD $0xdac000f3 // rbit    x19, x7
    WORD $0xdac01273 // clz    x19, x19
    WORD $0x8b1a0278 // add    x24, x19, x26
LBB0_646:
    WORD $0x8a3900f3 // bic    x19, x7, x25
    WORD $0x528000be // mov    w30, #5
    WORD $0xaa130734 // orr    x20, x25, x19, lsl #1
    WORD $0x8a3400e7 // bic    x7, x7, x20
    WORD $0x9201f0e7 // and    x7, x7, #0xaaaaaaaaaaaaaaaa
    WORD $0xab1300e7 // adds    x7, x7, x19
    WORD $0xd37ff8e7 // lsl    x7, x7, #1
    WORD $0x1a9f37f9 // cset    w25, hs
    WORD $0xd200f0e7 // eor    x7, x7, #0x5555555555555555
    WORD $0x8a1400e7 // and    x7, x7, x20
    WORD $0x8a2700c6 // bic    x6, x6, x7
    CMP $0, R6
    BEQ LBB0_643
LBB0_647:
    WORD $0xdac000c8 // rbit    x8, x6
    WORD $0xdac01108 // clz    x8, x8
    WORD $0x8b1a0108 // add    x8, x8, x26
    WORD $0x91000508 // add    x8, x8, #1
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BNE LBB0_924
LBB0_648:
    WORD $0x528d8c35 // mov    w21, #27745
    WORD $0x528eadd8 // mov    w24, #30062
    WORD $0x72acae75 // movk    w21, #25971, lsl #16
    WORD $0x72ad8d98 // movk    w24, #27756, lsl #16
    WORD $0xf10002ff // cmp    x23, #0
    WORD $0xf9000028 // str    x8, [x1]
    BGT LBB0_618
    B LBB0_952
LBB0_649:
    WORD $0xf9400068 // ldr    x8, [x3]
    WORD $0xf13ffd1f // cmp    x8, #4095
    BGT LBB0_923
    WORD $0x91000506 // add    x6, x8, #1
    WORD $0x8b080c68 // add    x8, x3, x8, lsl #3
    WORD $0xf9000066 // str    x6, [x3]
    WORD $0xf900051e // str    x30, [x8, #8]
    B LBB0_618
LBB0_651:
    WORD $0xaa0703fd // mov    x29, x7
    WORD $0x92800007 // mov    x7, #-1
    CMP $0, R24
    BNE LBB0_605
    B LBB0_933
LBB0_652:
    WORD $0x394006c6 // ldrb    w6, [x22, #1]
    WORD $0x5100b8c6 // sub    w6, w6, #46
    WORD $0x7100dcdf // cmp    w6, #55
    BHI LBB0_624
    WORD $0x5280003d // mov    w29, #1
    WORD $0x9ac62186 // lsl    x6, x12, x6
    WORD $0xea1000df // tst    x6, x16
    BEQ LBB0_700
LBB0_654:
    WORD $0xf100437f // cmp    x27, #16
    BLO LBB0_729
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x92800019 // mov    x25, #-1
    WORD $0x92800018 // mov    x24, #-1
    WORD $0x92800017 // mov    x23, #-1
    WORD $0xaa1b03fa // mov    x26, x27
LBB0_656:
    WORD $0x3cfd6ad3 // ldr    q19, [x22, x29]
    WORD $0x6e238e74 // cmeq    v20.16b, v19.16b, v3.16b
    WORD $0x6e258e75 // cmeq    v21.16b, v19.16b, v5.16b
    WORD $0x6e268e76 // cmeq    v22.16b, v19.16b, v6.16b
    WORD $0x4e278677 // add    v23.16b, v19.16b, v7.16b
    WORD $0x4e311e73 // and    v19.16b, v19.16b, v17.16b
    WORD $0x6e328e73 // cmeq    v19.16b, v19.16b, v18.16b
    WORD $0x6e373617 // cmhi    v23.16b, v16.16b, v23.16b
    WORD $0x4eb51ed5 // orr    v21.16b, v22.16b, v21.16b
    WORD $0x4eb31e96 // orr    v22.16b, v20.16b, v19.16b
    WORD $0x4eb61ef6 // orr    v22.16b, v23.16b, v22.16b
    WORD $0x4eb51ed6 // orr    v22.16b, v22.16b, v21.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e71bad6 // addv    h22, v22.8h
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x1e2602c6 // fmov    w6, s22
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x2a2603e6 // mvn    w6, w6
    WORD $0x4e71bab5 // addv    h21, v21.8h
    WORD $0x32103cc6 // orr    w6, w6, #0xffff0000
    WORD $0x5ac000c6 // rbit    w6, w6
    WORD $0x1e260294 // fmov    w20, s20
    WORD $0x1e260273 // fmov    w19, s19
    WORD $0x5ac010c6 // clz    w6, w6
    WORD $0x1e2602a7 // fmov    w7, s21
    WORD $0x710040df // cmp    w6, #16
    BEQ LBB0_658
    WORD $0x1ac620b5 // lsl    w21, w5, w6
    WORD $0x0a350294 // bic    w20, w20, w21
    WORD $0x0a350273 // bic    w19, w19, w21
    WORD $0x0a3500e7 // bic    w7, w7, w21
LBB0_658:
    WORD $0x51000695 // sub    w21, w20, #1
    WORD $0x6a1402be // ands    w30, w21, w20
    BNE LBB0_716
    WORD $0x51000675 // sub    w21, w19, #1
    WORD $0x6a1302be // ands    w30, w21, w19
    BNE LBB0_716
    WORD $0x510004f5 // sub    w21, w7, #1
    WORD $0x6a0702be // ands    w30, w21, w7
    BNE LBB0_716
    CMP $0, R20
    BEQ LBB0_664
    WORD $0x5ac00294 // rbit    w20, w20
    WORD $0xb10006ff // cmn    x23, #1
    WORD $0x5ac01294 // clz    w20, w20
    BNE LBB0_724
    WORD $0x8b1403b7 // add    x23, x29, x20
LBB0_664:
    CMP $0, R19
    BEQ LBB0_667
    WORD $0x5ac00273 // rbit    w19, w19
    WORD $0xb100071f // cmn    x24, #1
    WORD $0x5ac01273 // clz    w19, w19
    BNE LBB0_725
    WORD $0x8b1303b8 // add    x24, x29, x19
LBB0_667:
    CMP $0, R7
    BEQ LBB0_670
    WORD $0x5ac000e7 // rbit    w7, w7
    WORD $0xb100073f // cmn    x25, #1
    WORD $0x5ac010e7 // clz    w7, w7
    BNE LBB0_726
    WORD $0x8b0703b9 // add    x25, x29, x7
LBB0_670:
    WORD $0x710040df // cmp    w6, #16
    BNE LBB0_686
    WORD $0xd100435a // sub    x26, x26, #16
    WORD $0x910043bd // add    x29, x29, #16
    WORD $0xf1003f5f // cmp    x26, #15
    BHI LBB0_656
    WORD $0x8b1d02de // add    x30, x22, x29
    WORD $0xeb1d037f // cmp    x27, x29
    BEQ LBB0_687
LBB0_673:
    WORD $0x8b080247 // add    x7, x18, x8
    WORD $0x8b1a03c6 // add    x6, x30, x26
    WORD $0xcb1e00fd // sub    x29, x7, x30
    WORD $0xcb1603c7 // sub    x7, x30, x22
    WORD $0xaa1e03fb // mov    x27, x30
    B LBB0_676
LBB0_674:
    WORD $0xb100073f // cmn    x25, #1
    WORD $0xaa0703f9 // mov    x25, x7
    BNE LBB0_685
LBB0_675:
    WORD $0xd100075a // sub    x26, x26, #1
    WORD $0xd10007bd // sub    x29, x29, #1
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0xaa1b03fe // mov    x30, x27
    CMP $0, R26
    BEQ LBB0_714
LBB0_676:
    WORD $0x38401773 // ldrb    w19, [x27], #1
    WORD $0x5100c274 // sub    w20, w19, #48
    WORD $0x71002a9f // cmp    w20, #10
    BLO LBB0_675
    WORD $0x7100b67f // cmp    w19, #45
    BLE LBB0_682
    WORD $0x7101967f // cmp    w19, #101
    BEQ LBB0_684
    WORD $0x7101167f // cmp    w19, #69
    BEQ LBB0_684
    WORD $0x7100ba7f // cmp    w19, #46
    BNE LBB0_687
    WORD $0xb10006ff // cmn    x23, #1
    WORD $0xaa0703f7 // mov    x23, x7
    BEQ LBB0_675
    B LBB0_685
LBB0_682:
    WORD $0x7100ae7f // cmp    w19, #43
    BEQ LBB0_674
    WORD $0x7100b67f // cmp    w19, #45
    BEQ LBB0_674
    B LBB0_687
LBB0_684:
    WORD $0xb100071f // cmn    x24, #1
    WORD $0xaa0703f8 // mov    x24, x7
    BEQ LBB0_675
LBB0_685:
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    B LBB0_699
LBB0_686:
    WORD $0x8b2642c6 // add    x6, x22, w6, uxtw
    WORD $0x8b1d00de // add    x30, x6, x29
LBB0_687:
    WORD $0x9280001d // mov    x29, #-1
    CMP $0, R23
    BEQ LBB0_956
LBB0_688:
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    CMP $0, R25
    BEQ LBB0_956
    CMP $0, R24
    BEQ LBB0_956
    WORD $0xcb1603c6 // sub    x6, x30, x22
    WORD $0xd10004c7 // sub    x7, x6, #1
    WORD $0xeb0702ff // cmp    x23, x7
    BEQ LBB0_698
    WORD $0xeb07033f // cmp    x25, x7
    BEQ LBB0_698
    WORD $0xeb07031f // cmp    x24, x7
    BEQ LBB0_698
    WORD $0xf1000727 // subs    x7, x25, #1
    BLT LBB0_695
    WORD $0xeb07031f // cmp    x24, x7
    BNE LBB0_954
LBB0_695:
    WORD $0xaa1802e7 // orr    x7, x23, x24
    TST $(1<<63), R7
    BNE LBB0_697
    WORD $0xeb1802ff // cmp    x23, x24
    BGE LBB0_955
LBB0_697:
    WORD $0xd1000713 // sub    x19, x24, #1
    WORD $0xd37ffce7 // lsr    x7, x7, #63
    WORD $0xeb1302ff // cmp    x23, x19
    WORD $0x520000e7 // eor    w7, w7, #0x1
    WORD $0x1a9f17f3 // cset    w19, eq
    WORD $0x6a1300ff // tst    w7, w19
    WORD $0xda9800dd // csinv    x29, x6, x24, eq
    B LBB0_699
LBB0_698:
    WORD $0xcb0603fd // neg    x29, x6
LBB0_699:
    TST $(1<<63), R29
    BNE LBB0_956
LBB0_700:
    WORD $0x528d8c35 // mov    w21, #27745
    WORD $0x528eadd8 // mov    w24, #30062
    WORD $0x8b0803a6 // add    x6, x29, x8
    WORD $0x528000be // mov    w30, #5
    WORD $0x72acae75 // movk    w21, #25971, lsl #16
    WORD $0x72ad8d98 // movk    w24, #27756, lsl #16
    WORD $0x528000d9 // mov    w25, #6
    WORD $0xf100011f // cmp    x8, #0
    WORD $0xf9000026 // str    x6, [x1]
    BGT LBB0_618
    B LBB0_929
LBB0_701:
    WORD $0xcb1e03e7 // neg    x7, x30
    WORD $0x528000be // mov    w30, #5
    B LBB0_616
LBB0_702:
    WORD $0x8b1a013a // add    x26, x9, x26
    WORD $0xf1008106 // subs    x6, x8, #32
    BLO LBB0_734
LBB0_703:
    WORD $0xad405353 // ldp    q19, q20, [x26]
    WORD $0x6e208e75 // cmeq    v21.16b, v19.16b, v0.16b
    WORD $0x6e218e73 // cmeq    v19.16b, v19.16b, v1.16b
    WORD $0x6e208e96 // cmeq    v22.16b, v20.16b, v0.16b
    WORD $0x6e218e94 // cmeq    v20.16b, v20.16b, v1.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e71bad6 // addv    h22, v22.8h
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x1e260293 // fmov    w19, s20
    WORD $0x4e71bab4 // addv    h20, v21.8h
    WORD $0x1e2602d4 // fmov    w20, s22
    WORD $0x1e260267 // fmov    w7, s19
    WORD $0x1e260288 // fmov    w8, s20
    WORD $0x33103e67 // bfi    w7, w19, #16, #16
    WORD $0x33103e88 // bfi    w8, w20, #16, #16
    CMP $0, R7
    BNE LBB0_730
    CMP $0, R25
    BNE LBB0_732
    CMP $0, R8
    BEQ LBB0_733
LBB0_706:
    WORD $0xdac00108 // rbit    x8, x8
    WORD $0xcb090346 // sub    x6, x26, x9
    WORD $0xdac01108 // clz    x8, x8
    WORD $0x8b0800c8 // add    x8, x6, x8
    WORD $0x91000508 // add    x8, x8, #1
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BEQ LBB0_521
    B LBB0_924
LBB0_707:
    WORD $0x5ac002a6 // rbit    w6, w21
    WORD $0xaa3e03e7 // mvn    x7, x30
    WORD $0x5ac010c6 // clz    w6, w6
    WORD $0x528000be // mov    w30, #5
    WORD $0xcb0600e7 // sub    x7, x7, x6
    B LBB0_616
LBB0_708:
    WORD $0x8b1a013a // add    x26, x9, x26
    WORD $0xf1008106 // subs    x6, x8, #32
    BLO LBB0_751
LBB0_709:
    WORD $0xad405353 // ldp    q19, q20, [x26]
    WORD $0x6e208e75 // cmeq    v21.16b, v19.16b, v0.16b
    WORD $0x6e218e73 // cmeq    v19.16b, v19.16b, v1.16b
    WORD $0x6e208e96 // cmeq    v22.16b, v20.16b, v0.16b
    WORD $0x6e218e94 // cmeq    v20.16b, v20.16b, v1.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e71bad6 // addv    h22, v22.8h
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x1e260293 // fmov    w19, s20
    WORD $0x4e71bab4 // addv    h20, v21.8h
    WORD $0x1e2602d4 // fmov    w20, s22
    WORD $0x1e260267 // fmov    w7, s19
    WORD $0x1e260288 // fmov    w8, s20
    WORD $0x33103e67 // bfi    w7, w19, #16, #16
    WORD $0x33103e88 // bfi    w8, w20, #16, #16
    CMP $0, R7
    BNE LBB0_747
    CMP $0, R25
    BNE LBB0_749
    CMP $0, R8
    BEQ LBB0_750
LBB0_712:
    WORD $0xdac00108 // rbit    x8, x8
    WORD $0xcb090346 // sub    x6, x26, x9
    WORD $0xdac01108 // clz    x8, x8
    WORD $0x8b0800c8 // add    x8, x6, x8
    WORD $0x91000508 // add    x8, x8, #1
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BEQ LBB0_538
    B LBB0_924
LBB0_713:
    WORD $0xaa3e03e6 // mvn    x6, x30
    WORD $0x528000be // mov    w30, #5
    WORD $0xcb3340c7 // sub    x7, x6, w19, uxtw
    B LBB0_616
LBB0_714:
    WORD $0xaa0603fe // mov    x30, x6
    WORD $0x9280001d // mov    x29, #-1
    CMP $0, R23
    BNE LBB0_688
    B LBB0_956
LBB0_715:
    WORD $0xaa3e03e6 // mvn    x6, x30
    WORD $0x528000be // mov    w30, #5
    WORD $0xcb3440c7 // sub    x7, x6, w20, uxtw
    B LBB0_616
LBB0_716:
    WORD $0x5ac003c6 // rbit    w6, w30
    WORD $0xaa3d03e7 // mvn    x7, x29
    WORD $0x5ac010c6 // clz    w6, w6
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0xcb0600fd // sub    x29, x7, x6
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    B LBB0_699
LBB0_717:
    WORD $0x8b1a013a // add    x26, x9, x26
    WORD $0xf1008106 // subs    x6, x8, #32
    BLO LBB0_771
LBB0_718:
    WORD $0xad405353 // ldp    q19, q20, [x26]
    WORD $0x6e208e75 // cmeq    v21.16b, v19.16b, v0.16b
    WORD $0x6e218e73 // cmeq    v19.16b, v19.16b, v1.16b
    WORD $0x6e208e96 // cmeq    v22.16b, v20.16b, v0.16b
    WORD $0x6e218e94 // cmeq    v20.16b, v20.16b, v1.16b
    WORD $0x4e221eb5 // and    v21.16b, v21.16b, v2.16b
    WORD $0x4e221e94 // and    v20.16b, v20.16b, v2.16b
    WORD $0x4e040294 // tbl    v20.16b, { v20.16b }, v4.16b
    WORD $0x4e221ed6 // and    v22.16b, v22.16b, v2.16b
    WORD $0x4e221e73 // and    v19.16b, v19.16b, v2.16b
    WORD $0x4e0402d6 // tbl    v22.16b, { v22.16b }, v4.16b
    WORD $0x4e040273 // tbl    v19.16b, { v19.16b }, v4.16b
    WORD $0x4e0402b5 // tbl    v21.16b, { v21.16b }, v4.16b
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e71bad6 // addv    h22, v22.8h
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x1e260293 // fmov    w19, s20
    WORD $0x4e71bab4 // addv    h20, v21.8h
    WORD $0x1e2602d4 // fmov    w20, s22
    WORD $0x1e260267 // fmov    w7, s19
    WORD $0x1e260288 // fmov    w8, s20
    WORD $0x33103e67 // bfi    w7, w19, #16, #16
    WORD $0x33103e88 // bfi    w8, w20, #16, #16
    CMP $0, R7
    BNE LBB0_767
    CMP $0, R25
    BNE LBB0_769
    CMP $0, R8
    BEQ LBB0_770
LBB0_721:
    WORD $0xdac00108 // rbit    x8, x8
    WORD $0xcb090346 // sub    x6, x26, x9
    WORD $0xdac01108 // clz    x8, x8
    WORD $0x8b0800c8 // add    x8, x6, x8
    WORD $0x91000508 // add    x8, x8, #1
    B LBB0_781
LBB0_722:
    WORD $0xaa1f03f9 // mov    x25, xzr
    WORD $0x8b17013a // add    x26, x9, x23
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xf1008106 // subs    x6, x8, #32
    BHS LBB0_703
    B LBB0_734
LBB0_723:
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xaa0603fd // mov    x29, x6
    WORD $0xaa1703fb // mov    x27, x23
    WORD $0x92800019 // mov    x25, #-1
    WORD $0x9280001a // mov    x26, #-1
    B LBB0_588
LBB0_724:
    WORD $0xaa3d03e6 // mvn    x6, x29
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0xcb3440dd // sub    x29, x6, w20, uxtw
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    B LBB0_699
LBB0_725:
    WORD $0xaa3d03e6 // mvn    x6, x29
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0xcb3340dd // sub    x29, x6, w19, uxtw
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    B LBB0_699
LBB0_726:
    WORD $0xaa3d03e6 // mvn    x6, x29
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0xcb2740dd // sub    x29, x6, w7, uxtw
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    B LBB0_699
LBB0_727:
    WORD $0xaa1f03f9 // mov    x25, xzr
    WORD $0x8b17013a // add    x26, x9, x23
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xf1008106 // subs    x6, x8, #32
    BHS LBB0_709
    B LBB0_751
LBB0_728:
    WORD $0xaa1f03f9 // mov    x25, xzr
    WORD $0x8b17013a // add    x26, x9, x23
    WORD $0x92800018 // mov    x24, #-1
    WORD $0xf1008106 // subs    x6, x8, #32
    BHS LBB0_718
    B LBB0_771
LBB0_729:
    WORD $0x92800017 // mov    x23, #-1
    WORD $0xaa1603fe // mov    x30, x22
    WORD $0xaa1b03fa // mov    x26, x27
    WORD $0x92800018 // mov    x24, #-1
    WORD $0x92800019 // mov    x25, #-1
    B LBB0_673
LBB0_730:
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_732
    WORD $0xdac000f3 // rbit    x19, x7
    WORD $0xcb090354 // sub    x20, x26, x9
    WORD $0xdac01273 // clz    x19, x19
    WORD $0x8b140278 // add    x24, x19, x20
LBB0_732:
    WORD $0x0a3900f3 // bic    w19, w7, w25
    WORD $0x528000be // mov    w30, #5
    WORD $0x531f7a74 // lsl    w20, w19, #1
    WORD $0x0a3400e7 // bic    w7, w7, w20
    WORD $0x1201f0e7 // and    w7, w7, #0xaaaaaaaa
    WORD $0x331f7a79 // bfi    w25, w19, #1, #31
    WORD $0x2b1300e7 // adds    w7, w7, w19
    WORD $0x531f78e7 // lsl    w7, w7, #1
    WORD $0x1a9f37f3 // cset    w19, hs
    WORD $0x5200f0e7 // eor    w7, w7, #0x55555555
    WORD $0x0a1900e7 // and    w7, w7, w25
    WORD $0xaa1303f9 // mov    x25, x19
    WORD $0x2a2703e7 // mvn    w7, w7
    WORD $0x8a0800e8 // and    x8, x7, x8
    CMP $0, R8
    BNE LBB0_706
LBB0_733:
    WORD $0x9100835a // add    x26, x26, #32
    WORD $0xaa0603e8 // mov    x8, x6
LBB0_734:
    CMP $0, R25
    BNE LBB0_763
    WORD $0xaa1803e6 // mov    x6, x24
    CMP $0, R8
    BEQ LBB0_744
LBB0_736:
    WORD $0xaa1f03e7 // mov    x7, xzr
LBB0_737:
    WORD $0x38676b53 // ldrb    w19, [x26, x7]
    WORD $0x71008a7f // cmp    w19, #34
    BEQ LBB0_742
    WORD $0x7101727f // cmp    w19, #92
    BEQ LBB0_740
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0xeb07011f // cmp    x8, x7
    BNE LBB0_737
    B LBB0_745
LBB0_740:
    WORD $0xd1000513 // sub    x19, x8, #1
    WORD $0xeb07027f // cmp    x19, x7
    BEQ LBB0_925
    WORD $0x8b070353 // add    x19, x26, x7
    WORD $0xcb070114 // sub    x20, x8, x7
    WORD $0xd1000915 // sub    x21, x8, #2
    WORD $0xd1000a88 // sub    x8, x20, #2
    WORD $0x8b020274 // add    x20, x19, x2
    WORD $0xb10004df // cmn    x6, #1
    WORD $0x9a980298 // csel    x24, x20, x24, eq
    WORD $0x9a860286 // csel    x6, x20, x6, eq
    WORD $0x91000a7a // add    x26, x19, #2
    WORD $0xeb0702bf // cmp    x21, x7
    BNE LBB0_736
    B LBB0_925
LBB0_742:
    WORD $0x8b070348 // add    x8, x26, x7
    WORD $0x9100051a // add    x26, x8, #1
LBB0_743:
    WORD $0x528000be // mov    w30, #5
LBB0_744:
    WORD $0x528000d9 // mov    w25, #6
    WORD $0xcb090348 // sub    x8, x26, x9
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BEQ LBB0_521
    B LBB0_924
LBB0_745:
    WORD $0x71008a7f // cmp    w19, #34
    BNE LBB0_925
    WORD $0x8b08035a // add    x26, x26, x8
    B LBB0_743
LBB0_747:
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_749
    WORD $0xdac000f3 // rbit    x19, x7
    WORD $0xcb090354 // sub    x20, x26, x9
    WORD $0xdac01273 // clz    x19, x19
    WORD $0x8b140278 // add    x24, x19, x20
LBB0_749:
    WORD $0x0a3900f3 // bic    w19, w7, w25
    WORD $0x528000be // mov    w30, #5
    WORD $0x531f7a74 // lsl    w20, w19, #1
    WORD $0x0a3400e7 // bic    w7, w7, w20
    WORD $0x1201f0e7 // and    w7, w7, #0xaaaaaaaa
    WORD $0x331f7a79 // bfi    w25, w19, #1, #31
    WORD $0x2b1300e7 // adds    w7, w7, w19
    WORD $0x531f78e7 // lsl    w7, w7, #1
    WORD $0x1a9f37f3 // cset    w19, hs
    WORD $0x5200f0e7 // eor    w7, w7, #0x55555555
    WORD $0x0a1900e7 // and    w7, w7, w25
    WORD $0xaa1303f9 // mov    x25, x19
    WORD $0x2a2703e7 // mvn    w7, w7
    WORD $0x8a0800e8 // and    x8, x7, x8
    CMP $0, R8
    BNE LBB0_712
LBB0_750:
    WORD $0x9100835a // add    x26, x26, #32
    WORD $0xaa0603e8 // mov    x8, x6
LBB0_751:
    CMP $0, R25
    BNE LBB0_765
    WORD $0xaa1803e6 // mov    x6, x24
    CMP $0, R8
    BEQ LBB0_760
LBB0_753:
    WORD $0xaa1f03e7 // mov    x7, xzr
LBB0_754:
    WORD $0x38676b53 // ldrb    w19, [x26, x7]
    WORD $0x71008a7f // cmp    w19, #34
    BEQ LBB0_759
    WORD $0x7101727f // cmp    w19, #92
    BEQ LBB0_757
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0xeb07011f // cmp    x8, x7
    BNE LBB0_754
    B LBB0_761
LBB0_757:
    WORD $0xd1000513 // sub    x19, x8, #1
    WORD $0xeb07027f // cmp    x19, x7
    BEQ LBB0_925
    WORD $0x8b070353 // add    x19, x26, x7
    WORD $0xcb070114 // sub    x20, x8, x7
    WORD $0xd1000915 // sub    x21, x8, #2
    WORD $0xd1000a88 // sub    x8, x20, #2
    WORD $0x8b020274 // add    x20, x19, x2
    WORD $0xb10004df // cmn    x6, #1
    WORD $0x9a980298 // csel    x24, x20, x24, eq
    WORD $0x9a860286 // csel    x6, x20, x6, eq
    WORD $0x91000a7a // add    x26, x19, #2
    WORD $0xeb0702bf // cmp    x21, x7
    BNE LBB0_753
    B LBB0_925
LBB0_759:
    WORD $0x8b070348 // add    x8, x26, x7
    WORD $0x9100051a // add    x26, x8, #1
LBB0_760:
    WORD $0xcb090348 // sub    x8, x26, x9
    WORD $0x528000be // mov    w30, #5
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BEQ LBB0_538
    B LBB0_924
LBB0_761:
    WORD $0x71008a7f // cmp    w19, #34
    BNE LBB0_925
    WORD $0x8b08035a // add    x26, x26, x8
    B LBB0_760
LBB0_763:
    CMP $0, R8
    BEQ LBB0_925
    WORD $0x8b110346 // add    x6, x26, x17
    WORD $0xb100071f // cmn    x24, #1
    WORD $0x9a9800c7 // csel    x7, x6, x24, eq
    WORD $0x9a9800c6 // csel    x6, x6, x24, eq
    WORD $0x9100075a // add    x26, x26, #1
    WORD $0xd1000508 // sub    x8, x8, #1
    WORD $0xaa0703f8 // mov    x24, x7
    WORD $0x528000be // mov    w30, #5
    CMP $0, R8
    BNE LBB0_736
    B LBB0_744
LBB0_765:
    CMP $0, R8
    BEQ LBB0_925
    WORD $0x8b110346 // add    x6, x26, x17
    WORD $0xb100071f // cmn    x24, #1
    WORD $0x9a9800c7 // csel    x7, x6, x24, eq
    WORD $0x9a9800c6 // csel    x6, x6, x24, eq
    WORD $0x9100075a // add    x26, x26, #1
    WORD $0xd1000508 // sub    x8, x8, #1
    WORD $0xaa0703f8 // mov    x24, x7
    CMP $0, R8
    BNE LBB0_753
    B LBB0_760
LBB0_767:
    WORD $0xb100071f // cmn    x24, #1
    BNE LBB0_769
    WORD $0xdac000f3 // rbit    x19, x7
    WORD $0xcb090354 // sub    x20, x26, x9
    WORD $0xdac01273 // clz    x19, x19
    WORD $0x8b140278 // add    x24, x19, x20
LBB0_769:
    WORD $0x0a3900f3 // bic    w19, w7, w25
    WORD $0x531f7a74 // lsl    w20, w19, #1
    WORD $0x0a3400e7 // bic    w7, w7, w20
    WORD $0x1201f0e7 // and    w7, w7, #0xaaaaaaaa
    WORD $0x331f7a79 // bfi    w25, w19, #1, #31
    WORD $0x2b1300e7 // adds    w7, w7, w19
    WORD $0x531f78e7 // lsl    w7, w7, #1
    WORD $0x1a9f37f3 // cset    w19, hs
    WORD $0x5200f0e7 // eor    w7, w7, #0x55555555
    WORD $0x0a1900e7 // and    w7, w7, w25
    WORD $0xaa1303f9 // mov    x25, x19
    WORD $0x2a2703e7 // mvn    w7, w7
    WORD $0x8a0800e8 // and    x8, x7, x8
    CMP $0, R8
    BNE LBB0_721
LBB0_770:
    WORD $0x9100835a // add    x26, x26, #32
    WORD $0xaa0603e8 // mov    x8, x6
LBB0_771:
    CMP $0, R25
    BNE LBB0_784
    WORD $0xaa1803e6 // mov    x6, x24
    CMP $0, R8
    BEQ LBB0_780
LBB0_773:
    WORD $0xaa1f03e7 // mov    x7, xzr
LBB0_774:
    WORD $0x38676b53 // ldrb    w19, [x26, x7]
    WORD $0x71008a7f // cmp    w19, #34
    BEQ LBB0_779
    WORD $0x7101727f // cmp    w19, #92
    BEQ LBB0_777
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0xeb07011f // cmp    x8, x7
    BNE LBB0_774
    B LBB0_782
LBB0_777:
    WORD $0xd1000513 // sub    x19, x8, #1
    WORD $0xeb07027f // cmp    x19, x7
    BEQ LBB0_925
    WORD $0x8b070353 // add    x19, x26, x7
    WORD $0xcb070114 // sub    x20, x8, x7
    WORD $0xd1000915 // sub    x21, x8, #2
    WORD $0xd1000a88 // sub    x8, x20, #2
    WORD $0x8b020274 // add    x20, x19, x2
    WORD $0xb10004df // cmn    x6, #1
    WORD $0x9a980298 // csel    x24, x20, x24, eq
    WORD $0x9a860286 // csel    x6, x20, x6, eq
    WORD $0x91000a7a // add    x26, x19, #2
    WORD $0xeb0702bf // cmp    x21, x7
    BNE LBB0_773
    B LBB0_925
LBB0_779:
    WORD $0x8b070348 // add    x8, x26, x7
    WORD $0x9100051a // add    x26, x8, #1
LBB0_780:
    WORD $0xcb090348 // sub    x8, x26, x9
LBB0_781:
    WORD $0x528000be // mov    w30, #5
    WORD $0x528e4e9a // mov    w26, #29300
    WORD $0x528000d9 // mov    w25, #6
    WORD $0x72acaeba // movk    w26, #25973, lsl #16
    TST $(1<<63), R8
    BEQ LBB0_648
    B LBB0_924
LBB0_782:
    WORD $0x71008a7f // cmp    w19, #34
    BNE LBB0_925
    WORD $0x8b08035a // add    x26, x26, x8
    B LBB0_780
LBB0_784:
    CMP $0, R8
    BEQ LBB0_925
    WORD $0x8b110346 // add    x6, x26, x17
    WORD $0xb100071f // cmn    x24, #1
    WORD $0x9a9800c7 // csel    x7, x6, x24, eq
    WORD $0x9a9800c6 // csel    x6, x6, x24, eq
    WORD $0x9100075a // add    x26, x26, #1
    WORD $0xd1000508 // sub    x8, x8, #1
    WORD $0xaa0703f8 // mov    x24, x7
    CMP $0, R8
    BNE LBB0_773
    B LBB0_780
LBB0_786:
    WORD $0xa940200a // ldp    x10, x8, [x0]
    WORD $0xf940002b // ldr    x11, [x1]
    WORD $0xeb08017f // cmp    x11, x8
    BHS LBB0_790
    WORD $0x386b6949 // ldrb    w9, [x10, x11]
    WORD $0x7100353f // cmp    w9, #13
    BEQ LBB0_790
    WORD $0x7100813f // cmp    w9, #32
    BEQ LBB0_790
    WORD $0x51002d2c // sub    w12, w9, #11
    WORD $0xaa0b03e9 // mov    x9, x11
    WORD $0x3100099f // cmn    w12, #2
    BLO LBB0_813
LBB0_790:
    WORD $0x91000569 // add    x9, x11, #1
    WORD $0xeb08013f // cmp    x9, x8
    BHS LBB0_794
    WORD $0x3869694c // ldrb    w12, [x10, x9]
    WORD $0x7100359f // cmp    w12, #13
    BEQ LBB0_794
    WORD $0x7100819f // cmp    w12, #32
    BEQ LBB0_794
    WORD $0x51002d8c // sub    w12, w12, #11
    WORD $0x3100099f // cmn    w12, #2
    BLO LBB0_813
LBB0_794:
    WORD $0x91000969 // add    x9, x11, #2
    WORD $0xeb08013f // cmp    x9, x8
    BHS LBB0_798
    WORD $0x3869694c // ldrb    w12, [x10, x9]
    WORD $0x7100359f // cmp    w12, #13
    BEQ LBB0_798
    WORD $0x7100819f // cmp    w12, #32
    BEQ LBB0_798
    WORD $0x51002d8c // sub    w12, w12, #11
    WORD $0x3100099f // cmn    w12, #2
    BLO LBB0_813
LBB0_798:
    WORD $0x91000d69 // add    x9, x11, #3
    WORD $0xeb08013f // cmp    x9, x8
    BHS LBB0_802
    WORD $0x3869694c // ldrb    w12, [x10, x9]
    WORD $0x7100359f // cmp    w12, #13
    BEQ LBB0_802
    WORD $0x7100819f // cmp    w12, #32
    BEQ LBB0_802
    WORD $0x51002d8c // sub    w12, w12, #11
    WORD $0x3100099f // cmn    w12, #2
    BLO LBB0_813
LBB0_802:
    WORD $0x91001169 // add    x9, x11, #4
    WORD $0xeb08013f // cmp    x9, x8
    BHS LBB0_807
    WORD $0xd284c00c // mov    x12, #9728
    WORD $0x5280002b // mov    w11, #1
    WORD $0xf2c0002c // movk    x12, #1, lsl #32
LBB0_804:
    WORD $0x3869694d // ldrb    w13, [x10, x9]
    WORD $0x9acd216e // lsl    x14, x11, x13
    WORD $0x710081bf // cmp    w13, #32
    WORD $0x8a0c01cd // and    x13, x14, x12
    WORD $0xfa4099a4 // ccmp    x13, #0, #4, ls
    BEQ LBB0_812
    WORD $0x91000529 // add    x9, x9, #1
    WORD $0xeb09011f // cmp    x8, x9
    BNE LBB0_804
LBB0_806:
    WORD $0x92800008 // mov    x8, #-1
    B LBB0_868
LBB0_807:
    WORD $0x92800008 // mov    x8, #-1
    B LBB0_867
LBB0_808:
    WORD $0xd1000489 // sub    x9, x4, #1
    WORD $0x92800428 // mov    x8, #-34
    B LBB0_867
LBB0_809:
    WORD $0xf9000022 // str    x2, [x1]
    B LBB0_865
LBB0_810:
    WORD $0x7101749f // cmp    w4, #93
    BNE LBB0_865
LBB0_811:
    WORD $0x92800408 // mov    x8, #-33
    WORD $0xf9000022 // str    x2, [x1]
    B LBB0_868
LBB0_812:
    WORD $0xeb08013f // cmp    x9, x8
    BHS LBB0_806
LBB0_813:
    WORD $0x9100052d // add    x13, x9, #1
    WORD $0xf900002d // str    x13, [x1]
    WORD $0x38696948 // ldrb    w8, [x10, x9]
    WORD $0x7101691f // cmp    w8, #90
    BGT LBB0_833
    WORD $0x7100bd1f // cmp    w8, #47
    BLE LBB0_870
    WORD $0x5100c108 // sub    w8, w8, #48
    WORD $0x7100291f // cmp    w8, #10
    BHS LBB0_866
LBB0_816:
    WORD $0xf9400408 // ldr    x8, [x0, #8]
    WORD $0xcb0d010b // sub    x11, x8, x13
    WORD $0x8b0d0148 // add    x8, x10, x13
    WORD $0xf100417f // cmp    x11, #16
    BLO LBB0_820
    ADR LCPI0_0, R13
    ADR LCPI0_1, R14
    WORD $0x4f01e580 // movi    v0.16b, #44
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0x4f06e7e1 // movi    v1.16b, #223
    WORD $0x4f02e7a2 // movi    v2.16b, #93
    WORD $0x3dc001a3 // ldr    q3, [x13, :lo12:.LCPI0_0]
    WORD $0x3dc001c4 // ldr    q4, [x14, :lo12:.LCPI0_1]
LBB0_818:
    WORD $0x3dc00105 // ldr    q5, [x8]
    WORD $0x6e208ca6 // cmeq    v6.16b, v5.16b, v0.16b
    WORD $0x4e211ca5 // and    v5.16b, v5.16b, v1.16b
    WORD $0x6e228ca5 // cmeq    v5.16b, v5.16b, v2.16b
    WORD $0x4ea61ca5 // orr    v5.16b, v5.16b, v6.16b
    WORD $0x4e231ca5 // and    v5.16b, v5.16b, v3.16b
    WORD $0x4e0400a5 // tbl    v5.16b, { v5.16b }, v4.16b
    WORD $0x4e71b8a5 // addv    h5, v5.8h
    WORD $0x1e2600ad // fmov    w13, s5
    CMP $0, R13
    BNE LBB0_828
    WORD $0x91004108 // add    x8, x8, #16
    WORD $0xd100416b // sub    x11, x11, #16
    WORD $0x9100418c // add    x12, x12, #16
    WORD $0xf1003d7f // cmp    x11, #15
    BHI LBB0_818
LBB0_820:
    CMP $0, R11
    BEQ LBB0_928
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x8b0b010c // add    x12, x8, x11
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c2002e // movk    x14, #4097, lsl #32
LBB0_822:
    WORD $0x3940010f // ldrb    w15, [x8]
    WORD $0x7100b1ff // cmp    w15, #44
    BHI LBB0_824
    WORD $0x9acf21b0 // lsl    x16, x13, x15
    WORD $0xea0e021f // tst    x16, x14
    BNE LBB0_928
LBB0_824:
    WORD $0x710175ff // cmp    w15, #93
    BEQ LBB0_928
    WORD $0x7101f5ff // cmp    w15, #125
    BEQ LBB0_928
    WORD $0x91000508 // add    x8, x8, #1
    WORD $0xf100056b // subs    x11, x11, #1
    BNE LBB0_822
    WORD $0xaa0c03e8 // mov    x8, x12
    WORD $0xcb0a018a // sub    x10, x12, x10
    B LBB0_916
LBB0_828:
    WORD $0x5ac001ab // rbit    w11, w13
    WORD $0xcb0a010d // sub    x13, x8, x10
    WORD $0x5ac01168 // clz    w8, w11
    WORD $0x8b0801ab // add    x11, x13, x8
    WORD $0xf100057f // cmp    x11, #1
    WORD $0xf900002b // str    x11, [x1]
    BLT LBB0_832
    WORD $0x8b0c0108 // add    x8, x8, x12
    WORD $0xd284c00c // mov    x12, #9728
    WORD $0x8b09014a // add    x10, x10, x9
    WORD $0x5280002b // mov    w11, #1
    WORD $0xf2c0002c // movk    x12, #1, lsl #32
LBB0_830:
    WORD $0x3868694d // ldrb    w13, [x10, x8]
    WORD $0x9acd216e // lsl    x14, x11, x13
    WORD $0x710081bf // cmp    w13, #32
    WORD $0x8a0c01cd // and    x13, x14, x12
    WORD $0xfa4099a4 // ccmp    x13, #0, #4, ls
    BEQ LBB0_832
    WORD $0x8b08012d // add    x13, x9, x8
    WORD $0xd1000508 // sub    x8, x8, #1
    WORD $0x8b08012e // add    x14, x9, x8
    WORD $0x910009ce // add    x14, x14, #2
    WORD $0xf10005df // cmp    x14, #1
    WORD $0xf900002d // str    x13, [x1]
    BHI LBB0_830
LBB0_832:
    WORD $0xaa0903e8 // mov    x8, x9
    B LBB0_868
LBB0_833:
    WORD $0x7101b51f // cmp    w8, #109
    BLE LBB0_873
    WORD $0x7101b91f // cmp    w8, #110
    BEQ LBB0_876
    WORD $0x7101d11f // cmp    w8, #116
    BEQ LBB0_876
    WORD $0x7101ed1f // cmp    w8, #123
    BNE LBB0_866
    WORD $0xf940040c // ldr    x12, [x0, #8]
    ADR LCPI0_0, R15
    WORD $0x8b0d014a // add    x10, x10, x13
    ADR LCPI0_1, R17
    ADR LCPI0_2, R18_PLATFORM
    WORD $0xaa1f03f0 // mov    x16, xzr
    WORD $0xcb0d018c // sub    x12, x12, x13
    ADR LCPI0_7, R13
    WORD $0x3dc001e0 // ldr    q0, [x15, :lo12:.LCPI0_0]
    ADR LCPI0_3, R15
    WORD $0x3dc00221 // ldr    q1, [x17, :lo12:.LCPI0_1]
    ADR LCPI0_4, R17
    WORD $0x3dc001b0 // ldr    q16, [x13, :lo12:.LCPI0_7]
    ADR LCPI0_8, R13
    WORD $0x3dc001e3 // ldr    q3, [x15, :lo12:.LCPI0_3]
    ADR LCPI0_6, R15
    WORD $0x3dc00242 // ldr    q2, [x18, :lo12:.LCPI0_2]
    ADR LCPI0_5, R18_PLATFORM
    WORD $0x3dc001b3 // ldr    q19, [x13, :lo12:.LCPI0_8]
    ADR LCPI0_9, R13
    WORD $0x3dc001e6 // ldr    q6, [x15, :lo12:.LCPI0_6]
    WORD $0x910003ef // mov    x15, sp
    WORD $0x4f02e787 // movi    v7.16b, #92
    WORD $0xaa1f03eb // mov    x11, xzr
    WORD $0x4f01e451 // movi    v17.16b, #34
    WORD $0xaa1f03ee // mov    x14, xzr
    WORD $0x4f02e492 // movi    v18.16b, #68
    WORD $0xaa1f03e8 // mov    x8, xzr
    WORD $0x4f03e774 // movi    v20.16b, #123
    WORD $0x3dc00224 // ldr    q4, [x17, :lo12:.LCPI0_4]
    WORD $0x4f03e7b5 // movi    v21.16b, #125
    WORD $0x3dc00245 // ldr    q5, [x18, :lo12:.LCPI0_5]
    WORD $0x6f00e416 // movi    v22.2d, #0000000000000000
    WORD $0x3dc001b7 // ldr    q23, [x13, :lo12:.LCPI0_9]
    WORD $0x910081ed // add    x13, x15, #32
    WORD $0x25d8e040 // ptrue    p0.d, vl2
    WORD $0xf101018f // subs    x15, x12, #64
    BGE LBB0_854
LBB0_838:
    WORD $0xf100019f // cmp    x12, #0
    BLE LBB0_961
    WORD $0x92402d51 // and    x17, x10, #0xfff
    WORD $0xad005bf6 // stp    q22, q22, [sp]
    WORD $0xf13f063f // cmp    x17, #4033
    WORD $0xad015bf6 // stp    q22, q22, [sp, #32]
    BLO LBB0_854
    WORD $0xf1008192 // subs    x18, x12, #32
    BLO LBB0_842
    WORD $0xa9410943 // ldp    x3, x2, [x10, #16]
    WORD $0xaa0d03f1 // mov    x17, x13
    WORD $0x3cc20558 // ldr    q24, [x10], #32
    WORD $0x3d8003f8 // str    q24, [sp]
    WORD $0xa9010be3 // stp    x3, x2, [sp, #16]
    B LBB0_843
LBB0_842:
    WORD $0x910003f1 // mov    x17, sp
    WORD $0xaa0c03f2 // mov    x18, x12
LBB0_843:
    WORD $0xf1004242 // subs    x2, x18, #16
    BLO LBB0_849
    WORD $0xa8c14943 // ldp    x3, x18, [x10], #16
    WORD $0xa8814a23 // stp    x3, x18, [x17], #16
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0xf1002042 // subs    x2, x2, #8
    BHS LBB0_850
LBB0_845:
    WORD $0xf1001242 // subs    x2, x18, #4
    BLO LBB0_851
LBB0_846:
    WORD $0xb8404543 // ldr    w3, [x10], #4
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0xb8004623 // str    w3, [x17], #4
    WORD $0xf1000842 // subs    x2, x2, #2
    BHS LBB0_852
LBB0_847:
    CMP $0, R18_PLATFORM
    BEQ LBB0_853
LBB0_848:
    WORD $0x39400152 // ldrb    w18, [x10]
    WORD $0x910003ea // mov    x10, sp
    WORD $0x39000232 // strb    w18, [x17]
    B LBB0_854
LBB0_849:
    WORD $0xf1002242 // subs    x2, x18, #8
    BLO LBB0_845
LBB0_850:
    WORD $0xf8408543 // ldr    x3, [x10], #8
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0xf8008623 // str    x3, [x17], #8
    WORD $0xf1001042 // subs    x2, x2, #4
    BHS LBB0_846
LBB0_851:
    WORD $0xf1000a42 // subs    x2, x18, #2
    BLO LBB0_847
LBB0_852:
    WORD $0x78402543 // ldrh    w3, [x10], #2
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0x78002623 // strh    w3, [x17], #2
    CMP $0, R2
    BNE LBB0_848
LBB0_853:
    WORD $0x910003ea // mov    x10, sp
LBB0_854:
    WORD $0xad416d5a // ldp    q26, q27, [x10, #32]
    WORD $0x6e278f5e // cmeq    v30.16b, v26.16b, v7.16b
    WORD $0xad406558 // ldp    q24, q25, [x10]
    WORD $0x6e278f7f // cmeq    v31.16b, v27.16b, v7.16b
    WORD $0x4e201fde // and    v30.16b, v30.16b, v0.16b
    WORD $0x4e0103de // tbl    v30.16b, { v30.16b }, v1.16b
    WORD $0x6e278f1c // cmeq    v28.16b, v24.16b, v7.16b
    WORD $0x4e201fff // and    v31.16b, v31.16b, v0.16b
    WORD $0x6e278f3d // cmeq    v29.16b, v25.16b, v7.16b
    WORD $0x4e201f9c // and    v28.16b, v28.16b, v0.16b
    WORD $0x4e201fbd // and    v29.16b, v29.16b, v0.16b
    WORD $0x4e0103ff // tbl    v31.16b, { v31.16b }, v1.16b
    WORD $0x4e01039c // tbl    v28.16b, { v28.16b }, v1.16b
    WORD $0x4e0103bd // tbl    v29.16b, { v29.16b }, v1.16b
    WORD $0x4e71bbde // addv    h30, v30.8h
    WORD $0x4e71bbff // addv    h31, v31.8h
    WORD $0x4e71bb9c // addv    h28, v28.8h
    WORD $0x4e71bbbd // addv    h29, v29.8h
    WORD $0x1e2603d1 // fmov    w17, s30
    WORD $0x1e2603f2 // fmov    w18, s31
    WORD $0x1e260382 // fmov    w2, s28
    WORD $0x1e2603a3 // fmov    w3, s29
    WORD $0xd3607e31 // lsl    x17, x17, #32
    WORD $0xaa12c231 // orr    x17, x17, x18, lsl #48
    WORD $0x53103c72 // lsl    w18, w3, #16
    WORD $0xaa020231 // orr    x17, x17, x2
    WORD $0xaa120231 // orr    x17, x17, x18
    WORD $0xaa0b0232 // orr    x18, x17, x11
    CMP $0, R18_PLATFORM
    BNE LBB0_856
    WORD $0xaa1f03eb // mov    x11, xzr
    WORD $0x92800011 // mov    x17, #-1
    B LBB0_857
LBB0_856:
    WORD $0x8a2b0232 // bic    x18, x17, x11
    WORD $0xaa12056b // orr    x11, x11, x18, lsl #1
    WORD $0x8a2b0231 // bic    x17, x17, x11
    WORD $0x9201f231 // and    x17, x17, #0xaaaaaaaaaaaaaaaa
    WORD $0xab120231 // adds    x17, x17, x18
    WORD $0xd37ffa31 // lsl    x17, x17, #1
    WORD $0xd200f231 // eor    x17, x17, #0x5555555555555555
    WORD $0x8a0b0231 // and    x17, x17, x11
    WORD $0x1a9f37eb // cset    w11, hs
    WORD $0xaa3103f1 // mvn    x17, x17
LBB0_857:
    WORD $0x6e318f7b // cmeq    v27.16b, v27.16b, v17.16b
    WORD $0x6e318f5a // cmeq    v26.16b, v26.16b, v17.16b
    WORD $0x4e201f7b // and    v27.16b, v27.16b, v0.16b
    WORD $0x4e01037b // tbl    v27.16b, { v27.16b }, v1.16b
    WORD $0x4e201f5a // and    v26.16b, v26.16b, v0.16b
    WORD $0x6e318f1c // cmeq    v28.16b, v24.16b, v17.16b
    WORD $0x4e01035a // tbl    v26.16b, { v26.16b }, v1.16b
    WORD $0x6e318f3d // cmeq    v29.16b, v25.16b, v17.16b
    WORD $0x4e201f9c // and    v28.16b, v28.16b, v0.16b
    WORD $0x4e201fbd // and    v29.16b, v29.16b, v0.16b
    WORD $0x4e71bb7b // addv    h27, v27.8h
    WORD $0x4e01039c // tbl    v28.16b, { v28.16b }, v1.16b
    WORD $0x4e0103bd // tbl    v29.16b, { v29.16b }, v1.16b
    WORD $0x4e71bb5a // addv    h26, v26.8h
    WORD $0x1e260372 // fmov    w18, s27
    WORD $0x4e71bb9c // addv    h28, v28.8h
    WORD $0x4e71bbbb // addv    h27, v29.8h
    WORD $0x1e260342 // fmov    w2, s26
    WORD $0xd3503e52 // lsl    x18, x18, #48
    WORD $0x1e260383 // fmov    w3, s28
    WORD $0xaa028252 // orr    x18, x18, x2, lsl #32
    WORD $0x1e260362 // fmov    w2, s27
    WORD $0xaa030252 // orr    x18, x18, x3
    WORD $0x53103c42 // lsl    w2, w2, #16
    WORD $0xaa020252 // orr    x18, x18, x2
    WORD $0x8a110251 // and    x17, x18, x17
    WORD $0x9202e232 // and    x18, x17, #0x4444444444444444
    WORD $0x4e080e3a // dup    v26.2d, x17
    WORD $0x9201e231 // and    x17, x17, #0x8888888888888888
    WORD $0x4e080e5b // dup    v27.2d, x18
    WORD $0x0420bf7d // movprfx    z29, z27
    WORD $0x04d0005d // mul    z29.d, p0/m, z29.d, z2.d
    WORD $0x4e221f5a // and    v26.16b, v26.16b, v2.16b
    WORD $0x0420bf5c // movprfx    z28, z26
    WORD $0x04d0025c // mul    z28.d, p0/m, z28.d, z18.d
    WORD $0x6e3d1f9c // eor    v28.16b, v28.16b, v29.16b
    WORD $0x04d000db // mul    z27.d, p0/m, z27.d, z6.d
    WORD $0x4e080e3d // dup    v29.2d, x17
    WORD $0x0420bfa8 // movprfx    z8, z29
    WORD $0x04d00268 // mul    z8.d, p0/m, z8.d, z19.d
    WORD $0x4e08075e // dup    v30.2d, v26.d[0]
    WORD $0x6e1a435f // ext    v31.16b, v26.16b, v26.16b, #8
    WORD $0x4e18075a // dup    v26.2d, v26.d[1]
    WORD $0x04d0021d // mul    z29.d, p0/m, z29.d, z16.d
    WORD $0x04d0009e // mul    z30.d, p0/m, z30.d, z4.d
    WORD $0x04d000ba // mul    z26.d, p0/m, z26.d, z5.d
    WORD $0x6e281f7b // eor    v27.16b, v27.16b, v8.16b
    WORD $0x6e3d1f9c // eor    v28.16b, v28.16b, v29.16b
    WORD $0x0420bffd // movprfx    z29, z31
    WORD $0x04d0007d // mul    z29.d, p0/m, z29.d, z3.d
    WORD $0x6e3a1fda // eor    v26.16b, v30.16b, v26.16b
    WORD $0x6e3b1f5a // eor    v26.16b, v26.16b, v27.16b
    WORD $0x6e3c1fbb // eor    v27.16b, v29.16b, v28.16b
    WORD $0x4e371f7b // and    v27.16b, v27.16b, v23.16b
    WORD $0x4e241f5a // and    v26.16b, v26.16b, v4.16b
    WORD $0x4ebb1f5a // orr    v26.16b, v26.16b, v27.16b
    WORD $0x6e1a435b // ext    v27.16b, v26.16b, v26.16b, #8
    WORD $0x6e348f1c // cmeq    v28.16b, v24.16b, v20.16b
    WORD $0x6e358f18 // cmeq    v24.16b, v24.16b, v21.16b
    WORD $0x4e201f9c // and    v28.16b, v28.16b, v0.16b
    WORD $0x0ebb1f5a // orr    v26.8b, v26.8b, v27.8b
    WORD $0x6e348f3b // cmeq    v27.16b, v25.16b, v20.16b
    WORD $0x9e660351 // fmov    x17, d26
    WORD $0x6e358f39 // cmeq    v25.16b, v25.16b, v21.16b
    WORD $0x4e201f7a // and    v26.16b, v27.16b, v0.16b
    WORD $0x4e201f39 // and    v25.16b, v25.16b, v0.16b
    WORD $0xca100230 // eor    x16, x17, x16
    WORD $0x4e201f18 // and    v24.16b, v24.16b, v0.16b
    WORD $0xf100821f // cmp    x16, #32
    WORD $0x4e01039b // tbl    v27.16b, { v28.16b }, v1.16b
    WORD $0x2a3003e5 // mvn    w5, w16
    WORD $0x4e01035a // tbl    v26.16b, { v26.16b }, v1.16b
    WORD $0x1a9f97e2 // cset    w2, hi
    WORD $0x4e010339 // tbl    v25.16b, { v25.16b }, v1.16b
    WORD $0x4e010318 // tbl    v24.16b, { v24.16b }, v1.16b
    WORD $0x4e71bb7b // addv    h27, v27.8h
    WORD $0x4e71bb5a // addv    h26, v26.8h
    WORD $0x4e71bb39 // addv    h25, v25.8h
    WORD $0x4e71bb18 // addv    h24, v24.8h
    WORD $0x1e260343 // fmov    w3, s26
    WORD $0x1e260324 // fmov    w4, s25
    WORD $0x1e260372 // fmov    w18, s27
    WORD $0x1e260311 // fmov    w17, s24
    WORD $0x33103c72 // bfi    w18, w3, #16, #16
    WORD $0x33103c91 // bfi    w17, w4, #16, #16
    WORD $0x8a1200a3 // and    x3, x5, x18
    WORD $0xea1100a4 // ands    x4, x5, x17
    BEQ LBB0_860
LBB0_858:
    WORD $0xd1000485 // sub    x5, x4, #1
    WORD $0x8a0300a6 // and    x6, x5, x3
    WORD $0x9e6700d8 // fmov    d24, x6
    WORD $0x0e205b18 // cnt    v24.8b, v24.8b
    WORD $0x2e303b18 // uaddlv    h24, v24.8b
    WORD $0x1e260306 // fmov    w6, s24
    WORD $0x8b0e00c6 // add    x6, x6, x14
    WORD $0xeb0800df // cmp    x6, x8
    BLS LBB0_914
    WORD $0xea0400a4 // ands    x4, x5, x4
    WORD $0x91000508 // add    x8, x8, #1
    BNE LBB0_858
LBB0_860:
    WORD $0x9e670078 // fmov    d24, x3
    WORD $0xd2407c42 // eor    x2, x2, #0xffffffff
    WORD $0x8a120052 // and    x18, x2, x18
    WORD $0xea110051 // ands    x17, x2, x17
    WORD $0x0e205b18 // cnt    v24.8b, v24.8b
    WORD $0x2e303b18 // uaddlv    h24, v24.8b
    WORD $0x1e260303 // fmov    w3, s24
    WORD $0x8b0e006e // add    x14, x3, x14
    BEQ LBB0_863
LBB0_861:
    WORD $0xd1000622 // sub    x2, x17, #1
    WORD $0x8a120043 // and    x3, x2, x18
    WORD $0x9e670078 // fmov    d24, x3
    WORD $0x0e205b18 // cnt    v24.8b, v24.8b
    WORD $0x2e303b18 // uaddlv    h24, v24.8b
    WORD $0x1e260303 // fmov    w3, s24
    WORD $0x8b0e0063 // add    x3, x3, x14
    WORD $0xeb08007f // cmp    x3, x8
    BLS LBB0_917
    WORD $0xea110051 // ands    x17, x2, x17
    WORD $0x91000508 // add    x8, x8, #1
    BNE LBB0_861
LBB0_863:
    WORD $0x9e670258 // fmov    d24, x18
    WORD $0x937ffe10 // asr    x16, x16, #63
    WORD $0x9101014a // add    x10, x10, #64
    WORD $0x0e205b18 // cnt    v24.8b, v24.8b
    WORD $0x2e303b18 // uaddlv    h24, v24.8b
    WORD $0x1e26030c // fmov    w12, s24
    WORD $0x8b0e018e // add    x14, x12, x14
    WORD $0xaa0f03ec // mov    x12, x15
    WORD $0xf10101ef // subs    x15, x15, #64
    BGE LBB0_854
    B LBB0_838
LBB0_864:
    WORD $0x7101f63f // cmp    w17, #125
    BEQ LBB0_811
LBB0_865:
    WORD $0xf9400028 // ldr    x8, [x1]
    WORD $0xd1000509 // sub    x9, x8, #1
LBB0_866:
    WORD $0x92800028 // mov    x8, #-2
LBB0_867:
    WORD $0xf9000029 // str    x9, [x1]
LBB0_868:
    WORD $0xa94b4ff4 // ldp    x20, x19, [sp, #176]
    WORD $0xaa0803e0 // mov    x0, x8
    WORD $0xa94a57f6 // ldp    x22, x21, [sp, #160]
    WORD $0xa9495ff8 // ldp    x24, x23, [sp, #144]
    WORD $0xa94867fa // ldp    x26, x25, [sp, #128]
    WORD $0xa9476ffe // ldp    x30, x27, [sp, #112]
    WORD $0x6d45a3e9 // ldp    d9, d8, [sp, #88]
    WORD $0x6d44abeb // ldp    d11, d10, [sp, #72]
    WORD $0xf94037fd // ldr    x29, [sp, #104]
    WORD $0xfd4023ec // ldr    d12, [sp, #64]
    WORD $0x910303ff // add    sp, sp, #192
    WORD $0xd65f03c0 // ret
LBB0_869:
    WORD $0x92800008 // mov    x8, #-1
    WORD $0xf9000036 // str    x22, [x1]
    B LBB0_868
LBB0_870:
    CMP $0, R8
    BEQ LBB0_806
    WORD $0x7100891f // cmp    w8, #34
    BEQ LBB0_877
    WORD $0x7100b51f // cmp    w8, #45
    BEQ LBB0_816
    B LBB0_866
LBB0_873:
    WORD $0x71016d1f // cmp    w8, #91
    BEQ LBB0_887
    WORD $0x7101991f // cmp    w8, #102
    BNE LBB0_866
    WORD $0xf9400408 // ldr    x8, [x0, #8]
    WORD $0x9100152a // add    x10, x9, #5
    WORD $0xeb08015f // cmp    x10, x8
    BHI LBB0_806
    B LBB0_916
LBB0_876:
    WORD $0xf9400408 // ldr    x8, [x0, #8]
    WORD $0x9100112a // add    x10, x9, #4
    WORD $0xeb08015f // cmp    x10, x8
    BHI LBB0_806
    B LBB0_916
LBB0_877:
    WORD $0xf940040e // ldr    x14, [x0, #8]
    WORD $0xcb0d01cb // sub    x11, x14, x13
    WORD $0xf100817f // cmp    x11, #32
    BLT LBB0_960
    ADR LCPI0_0, R15
    ADR LCPI0_1, R16
    WORD $0x4f01e440 // movi    v0.16b, #34
    WORD $0xaa1f03e8 // mov    x8, xzr
    WORD $0x4f02e782 // movi    v2.16b, #92
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x8b09014c // add    x12, x10, x9
    WORD $0xcb0901cb // sub    x11, x14, x9
    WORD $0x3dc001e1 // ldr    q1, [x15, :lo12:.LCPI0_0]
    WORD $0x528003ee // mov    w14, #31
    WORD $0x3dc00203 // ldr    q3, [x16, :lo12:.LCPI0_1]
LBB0_879:
    WORD $0x8b08018f // add    x15, x12, x8
    WORD $0x3cc111e4 // ldur    q4, [x15, #17]
    WORD $0x3cc011e6 // ldur    q6, [x15, #1]
    WORD $0x6e228c85 // cmeq    v5.16b, v4.16b, v2.16b
    WORD $0x6e208c84 // cmeq    v4.16b, v4.16b, v0.16b
    WORD $0x4e211ca5 // and    v5.16b, v5.16b, v1.16b
    WORD $0x4e0300a5 // tbl    v5.16b, { v5.16b }, v3.16b
    WORD $0x4e211c84 // and    v4.16b, v4.16b, v1.16b
    WORD $0x4e030084 // tbl    v4.16b, { v4.16b }, v3.16b
    WORD $0x4e71b8a5 // addv    h5, v5.8h
    WORD $0x4e71b884 // addv    h4, v4.8h
    WORD $0x1e2600af // fmov    w15, s5
    WORD $0x6e208cc5 // cmeq    v5.16b, v6.16b, v0.16b
    WORD $0x6e228cc6 // cmeq    v6.16b, v6.16b, v2.16b
    WORD $0x4e211ca5 // and    v5.16b, v5.16b, v1.16b
    WORD $0x4e211cc6 // and    v6.16b, v6.16b, v1.16b
    WORD $0x4e0300c6 // tbl    v6.16b, { v6.16b }, v3.16b
    WORD $0x4e0300a5 // tbl    v5.16b, { v5.16b }, v3.16b
    WORD $0x1e260091 // fmov    w17, s4
    WORD $0x4e71b8c6 // addv    h6, v6.8h
    WORD $0x4e71b8a5 // addv    h5, v5.8h
    WORD $0x1e2600d0 // fmov    w16, s6
    WORD $0x33103df0 // bfi    w16, w15, #16, #16
    WORD $0x1e2600af // fmov    w15, s5
    WORD $0x7100021f // cmp    w16, #0
    WORD $0xfa4009a0 // ccmp    x13, #0, #0, eq
    WORD $0x33103e2f // bfi    w15, w17, #16, #16
    BEQ LBB0_881
    WORD $0x0a2d0211 // bic    w17, w16, w13
    WORD $0x2a1105ad // orr    w13, w13, w17, lsl #1
    WORD $0x0a2d0210 // bic    w16, w16, w13
    WORD $0x1201f210 // and    w16, w16, #0xaaaaaaaa
    WORD $0x2b110210 // adds    w16, w16, w17
    WORD $0x531f7a10 // lsl    w16, w16, #1
    WORD $0x5200f210 // eor    w16, w16, #0x55555555
    WORD $0x0a0d020d // and    w13, w16, w13
    WORD $0x2a2d03f0 // mvn    w16, w13
    WORD $0x1a9f37ed // cset    w13, hs
    WORD $0x8a0f020f // and    x15, x16, x15
    B LBB0_882
LBB0_881:
    WORD $0xaa1f03ed // mov    x13, xzr
LBB0_882:
    CMP $0, R15
    BNE LBB0_915
    WORD $0xd10081ce // sub    x14, x14, #32
    WORD $0x91008108 // add    x8, x8, #32
    WORD $0x8b0e016f // add    x15, x11, x14
    WORD $0xf100fdff // cmp    x15, #63
    BGT LBB0_879
    CMP $0, R13
    BNE LBB0_962
    WORD $0x8b09014c // add    x12, x10, x9
    WORD $0x8b08018c // add    x12, x12, x8
    WORD $0xaa2803e8 // mvn    x8, x8
    WORD $0x9100058c // add    x12, x12, #1
    WORD $0x8b0b010b // add    x11, x8, x11
LBB0_886:
    WORD $0xf100057f // cmp    x11, #1
    BGE LBB0_920
    B LBB0_806
LBB0_887:
    WORD $0xf940040c // ldr    x12, [x0, #8]
    ADR LCPI0_0, R15
    WORD $0x8b0d014a // add    x10, x10, x13
    ADR LCPI0_1, R17
    ADR LCPI0_2, R18_PLATFORM
    WORD $0xaa1f03f0 // mov    x16, xzr
    WORD $0xcb0d018c // sub    x12, x12, x13
    ADR LCPI0_7, R13
    WORD $0x3dc001e0 // ldr    q0, [x15, :lo12:.LCPI0_0]
    ADR LCPI0_3, R15
    WORD $0x3dc00221 // ldr    q1, [x17, :lo12:.LCPI0_1]
    ADR LCPI0_4, R17
    WORD $0x3dc001b0 // ldr    q16, [x13, :lo12:.LCPI0_7]
    ADR LCPI0_8, R13
    WORD $0x3dc001e3 // ldr    q3, [x15, :lo12:.LCPI0_3]
    ADR LCPI0_6, R15
    WORD $0x3dc00242 // ldr    q2, [x18, :lo12:.LCPI0_2]
    ADR LCPI0_5, R18_PLATFORM
    WORD $0x3dc001b3 // ldr    q19, [x13, :lo12:.LCPI0_8]
    ADR LCPI0_9, R13
    WORD $0x3dc001e6 // ldr    q6, [x15, :lo12:.LCPI0_6]
    WORD $0x910003ef // mov    x15, sp
    WORD $0x4f02e787 // movi    v7.16b, #92
    WORD $0xaa1f03eb // mov    x11, xzr
    WORD $0x4f01e451 // movi    v17.16b, #34
    WORD $0xaa1f03ee // mov    x14, xzr
    WORD $0x4f02e492 // movi    v18.16b, #68
    WORD $0xaa1f03e8 // mov    x8, xzr
    WORD $0x4f02e774 // movi    v20.16b, #91
    WORD $0x3dc00224 // ldr    q4, [x17, :lo12:.LCPI0_4]
    WORD $0x4f02e7b5 // movi    v21.16b, #93
    WORD $0x3dc00245 // ldr    q5, [x18, :lo12:.LCPI0_5]
    WORD $0x6f00e416 // movi    v22.2d, #0000000000000000
    WORD $0x3dc001b7 // ldr    q23, [x13, :lo12:.LCPI0_9]
    WORD $0x910081ed // add    x13, x15, #32
    WORD $0x25d8e040 // ptrue    p0.d, vl2
    WORD $0xf101018f // subs    x15, x12, #64
    BGE LBB0_904
LBB0_888:
    WORD $0xf100019f // cmp    x12, #0
    BLE LBB0_961
    WORD $0x92402d51 // and    x17, x10, #0xfff
    WORD $0xad005bf6 // stp    q22, q22, [sp]
    WORD $0xf13f063f // cmp    x17, #4033
    WORD $0xad015bf6 // stp    q22, q22, [sp, #32]
    BLO LBB0_904
    WORD $0xf1008192 // subs    x18, x12, #32
    BLO LBB0_892
    WORD $0xa9410943 // ldp    x3, x2, [x10, #16]
    WORD $0xaa0d03f1 // mov    x17, x13
    WORD $0x3cc20558 // ldr    q24, [x10], #32
    WORD $0x3d8003f8 // str    q24, [sp]
    WORD $0xa9010be3 // stp    x3, x2, [sp, #16]
    B LBB0_893
LBB0_892:
    WORD $0x910003f1 // mov    x17, sp
    WORD $0xaa0c03f2 // mov    x18, x12
LBB0_893:
    WORD $0xf1004242 // subs    x2, x18, #16
    BLO LBB0_899
    WORD $0xa8c14943 // ldp    x3, x18, [x10], #16
    WORD $0xa8814a23 // stp    x3, x18, [x17], #16
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0xf1002042 // subs    x2, x2, #8
    BHS LBB0_900
LBB0_895:
    WORD $0xf1001242 // subs    x2, x18, #4
    BLO LBB0_901
LBB0_896:
    WORD $0xb8404543 // ldr    w3, [x10], #4
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0xb8004623 // str    w3, [x17], #4
    WORD $0xf1000842 // subs    x2, x2, #2
    BHS LBB0_902
LBB0_897:
    CMP $0, R18_PLATFORM
    BEQ LBB0_903
LBB0_898:
    WORD $0x39400152 // ldrb    w18, [x10]
    WORD $0x910003ea // mov    x10, sp
    WORD $0x39000232 // strb    w18, [x17]
    B LBB0_904
LBB0_899:
    WORD $0xf1002242 // subs    x2, x18, #8
    BLO LBB0_895
LBB0_900:
    WORD $0xf8408543 // ldr    x3, [x10], #8
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0xf8008623 // str    x3, [x17], #8
    WORD $0xf1001042 // subs    x2, x2, #4
    BHS LBB0_896
LBB0_901:
    WORD $0xf1000a42 // subs    x2, x18, #2
    BLO LBB0_897
LBB0_902:
    WORD $0x78402543 // ldrh    w3, [x10], #2
    WORD $0xaa0203f2 // mov    x18, x2
    WORD $0x78002623 // strh    w3, [x17], #2
    CMP $0, R2
    BNE LBB0_898
LBB0_903:
    WORD $0x910003ea // mov    x10, sp
LBB0_904:
    WORD $0xad416d5a // ldp    q26, q27, [x10, #32]
    WORD $0x6e278f5e // cmeq    v30.16b, v26.16b, v7.16b
    WORD $0xad406558 // ldp    q24, q25, [x10]
    WORD $0x6e278f7f // cmeq    v31.16b, v27.16b, v7.16b
    WORD $0x4e201fde // and    v30.16b, v30.16b, v0.16b
    WORD $0x4e0103de // tbl    v30.16b, { v30.16b }, v1.16b
    WORD $0x6e278f1c // cmeq    v28.16b, v24.16b, v7.16b
    WORD $0x4e201fff // and    v31.16b, v31.16b, v0.16b
    WORD $0x6e278f3d // cmeq    v29.16b, v25.16b, v7.16b
    WORD $0x4e201f9c // and    v28.16b, v28.16b, v0.16b
    WORD $0x4e201fbd // and    v29.16b, v29.16b, v0.16b
    WORD $0x4e0103ff // tbl    v31.16b, { v31.16b }, v1.16b
    WORD $0x4e01039c // tbl    v28.16b, { v28.16b }, v1.16b
    WORD $0x4e0103bd // tbl    v29.16b, { v29.16b }, v1.16b
    WORD $0x4e71bbde // addv    h30, v30.8h
    WORD $0x4e71bbff // addv    h31, v31.8h
    WORD $0x4e71bb9c // addv    h28, v28.8h
    WORD $0x4e71bbbd // addv    h29, v29.8h
    WORD $0x1e2603d1 // fmov    w17, s30
    WORD $0x1e2603f2 // fmov    w18, s31
    WORD $0x1e260382 // fmov    w2, s28
    WORD $0x1e2603a3 // fmov    w3, s29
    WORD $0xd3607e31 // lsl    x17, x17, #32
    WORD $0xaa12c231 // orr    x17, x17, x18, lsl #48
    WORD $0x53103c72 // lsl    w18, w3, #16
    WORD $0xaa020231 // orr    x17, x17, x2
    WORD $0xaa120231 // orr    x17, x17, x18
    WORD $0xaa0b0232 // orr    x18, x17, x11
    CMP $0, R18_PLATFORM
    BNE LBB0_906
    WORD $0xaa1f03eb // mov    x11, xzr
    WORD $0x92800011 // mov    x17, #-1
    B LBB0_907
LBB0_906:
    WORD $0x8a2b0232 // bic    x18, x17, x11
    WORD $0xaa12056b // orr    x11, x11, x18, lsl #1
    WORD $0x8a2b0231 // bic    x17, x17, x11
    WORD $0x9201f231 // and    x17, x17, #0xaaaaaaaaaaaaaaaa
    WORD $0xab120231 // adds    x17, x17, x18
    WORD $0xd37ffa31 // lsl    x17, x17, #1
    WORD $0xd200f231 // eor    x17, x17, #0x5555555555555555
    WORD $0x8a0b0231 // and    x17, x17, x11
    WORD $0x1a9f37eb // cset    w11, hs
    WORD $0xaa3103f1 // mvn    x17, x17
LBB0_907:
    WORD $0x6e318f7b // cmeq    v27.16b, v27.16b, v17.16b
    WORD $0x6e318f5a // cmeq    v26.16b, v26.16b, v17.16b
    WORD $0x4e201f7b // and    v27.16b, v27.16b, v0.16b
    WORD $0x4e01037b // tbl    v27.16b, { v27.16b }, v1.16b
    WORD $0x4e201f5a // and    v26.16b, v26.16b, v0.16b
    WORD $0x6e318f1c // cmeq    v28.16b, v24.16b, v17.16b
    WORD $0x4e01035a // tbl    v26.16b, { v26.16b }, v1.16b
    WORD $0x6e318f3d // cmeq    v29.16b, v25.16b, v17.16b
    WORD $0x4e201f9c // and    v28.16b, v28.16b, v0.16b
    WORD $0x4e201fbd // and    v29.16b, v29.16b, v0.16b
    WORD $0x4e71bb7b // addv    h27, v27.8h
    WORD $0x4e01039c // tbl    v28.16b, { v28.16b }, v1.16b
    WORD $0x4e0103bd // tbl    v29.16b, { v29.16b }, v1.16b
    WORD $0x4e71bb5a // addv    h26, v26.8h
    WORD $0x1e260372 // fmov    w18, s27
    WORD $0x4e71bb9c // addv    h28, v28.8h
    WORD $0x4e71bbbb // addv    h27, v29.8h
    WORD $0x1e260342 // fmov    w2, s26
    WORD $0xd3503e52 // lsl    x18, x18, #48
    WORD $0x1e260383 // fmov    w3, s28
    WORD $0xaa028252 // orr    x18, x18, x2, lsl #32
    WORD $0x1e260362 // fmov    w2, s27
    WORD $0xaa030252 // orr    x18, x18, x3
    WORD $0x53103c42 // lsl    w2, w2, #16
    WORD $0xaa020252 // orr    x18, x18, x2
    WORD $0x8a110251 // and    x17, x18, x17
    WORD $0x9202e232 // and    x18, x17, #0x4444444444444444
    WORD $0x4e080e3a // dup    v26.2d, x17
    WORD $0x9201e231 // and    x17, x17, #0x8888888888888888
    WORD $0x4e080e5b // dup    v27.2d, x18
    WORD $0x0420bf7d // movprfx    z29, z27
    WORD $0x04d0005d // mul    z29.d, p0/m, z29.d, z2.d
    WORD $0x4e221f5a // and    v26.16b, v26.16b, v2.16b
    WORD $0x0420bf5c // movprfx    z28, z26
    WORD $0x04d0025c // mul    z28.d, p0/m, z28.d, z18.d
    WORD $0x6e3d1f9c // eor    v28.16b, v28.16b, v29.16b
    WORD $0x04d000db // mul    z27.d, p0/m, z27.d, z6.d
    WORD $0x4e080e3d // dup    v29.2d, x17
    WORD $0x0420bfa8 // movprfx    z8, z29
    WORD $0x04d00268 // mul    z8.d, p0/m, z8.d, z19.d
    WORD $0x4e08075e // dup    v30.2d, v26.d[0]
    WORD $0x6e1a435f // ext    v31.16b, v26.16b, v26.16b, #8
    WORD $0x4e18075a // dup    v26.2d, v26.d[1]
    WORD $0x04d0021d // mul    z29.d, p0/m, z29.d, z16.d
