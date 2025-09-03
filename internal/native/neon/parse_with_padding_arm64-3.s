LBB5_1017:
    WORD $0xb100047f // cmn    x3, #1
    WORD $0xaa0603e3 // mov    x3, x6
    BNE LBB5_1162
LBB5_1018:
    WORD $0xd1000631 // sub    x17, x17, #1
    WORD $0xd1000484 // sub    x4, x4, #1
    WORD $0x910004c6 // add    x6, x6, #1
    WORD $0xaa1303e7 // mov    x7, x19
    CMP $0, R17
    BEQ LBB5_1238
LBB5_1019:
    WORD $0x3840166d // ldrb    w13, [x19], #1
    WORD $0x5100c1ae // sub    w14, w13, #48
    WORD $0x710029df // cmp    w14, #10
    BLO LBB5_1018
    WORD $0x7100b5bf // cmp    w13, #45
    BLE LBB5_1025
    WORD $0x710195bf // cmp    w13, #101
    BEQ LBB5_1017
    WORD $0x710115bf // cmp    w13, #69
    BEQ LBB5_1017
    WORD $0x7100b9bf // cmp    w13, #46
    BNE LBB5_1038
    WORD $0xb100065f // cmn    x18, #1
    WORD $0xaa0603f2 // mov    x18, x6
    BEQ LBB5_1018
    B LBB5_1162
LBB5_1025:
    WORD $0x7100adbf // cmp    w13, #43
    BEQ LBB5_1027
    WORD $0x7100b5bf // cmp    w13, #45
    BNE LBB5_1038
LBB5_1027:
    WORD $0xb100045f // cmn    x2, #1
    WORD $0xaa0603e2 // mov    x2, x6
    BEQ LBB5_1018
    B LBB5_1162
LBB5_1028:
    WORD $0x71016e5f // cmp    w18, #91
    BEQ LBB5_1055
    WORD $0x71019a5f // cmp    w18, #102
    BNE LBB5_341
    WORD $0xaa0f03f1 // mov    x17, x15
    WORD $0x528001b2 // mov    w18, #13
    WORD $0x3840162d // ldrb    w13, [x17], #1
    WORD $0x710185bf // cmp    w13, #97
    BNE LBB5_1034
    WORD $0x394005ed // ldrb    w13, [x15, #1]
    WORD $0x910009f1 // add    x17, x15, #2
    WORD $0x7101b1bf // cmp    w13, #108
    BNE LBB5_1034
    WORD $0x394009ed // ldrb    w13, [x15, #2]
    WORD $0x91000df1 // add    x17, x15, #3
    WORD $0x7101cdbf // cmp    w13, #115
    BNE LBB5_1034
    WORD $0x39400ded // ldrb    w13, [x15, #3]
    WORD $0x910011f1 // add    x17, x15, #4
    WORD $0x710195bf // cmp    w13, #101
    WORD $0x1a9203f2 // csel    w18, wzr, w18, eq
LBB5_1034:
    WORD $0x5280004d // mov    w13, #2
    WORD $0xaa0003ee // mov    x14, x0
    WORD $0xaa1081ad // orr    x13, x13, x16, lsl #32
    WORD $0xaa1103ef // mov    x15, x17
    WORD $0xf80105cd // str    x13, [x14], #16
    WORD $0xf900510e // str    x14, [x8, #160]
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a9201a2 // csel    w2, w13, w18, eq
    CMP $0, R18_PLATFORM
    BEQ LBB5_1630
    B LBB5_341
LBB5_1035:
    WORD $0xaa0503f1 // mov    x17, x5
    WORD $0xf1000d9f // cmp    x12, #3
    BNE LBB5_1599
LBB5_1036:
    WORD $0xf9405100 // ldr    x0, [x8, #160]
    B LBB5_1240
LBB5_1037:
    WORD $0x8b3440ad // add    x13, x5, w20, uxtw
    WORD $0x8b0401a7 // add    x7, x13, x4
LBB5_1038:
    WORD $0x92800004 // mov    x4, #-1
    CMP $0, R18_PLATFORM
    BEQ LBB5_1226
LBB5_1039:
    CMP $0, R2
    BEQ LBB5_1226
    CMP $0, R3
    BEQ LBB5_1226
    WORD $0xcb0500ed // sub    x13, x7, x5
    WORD $0xd10005af // sub    x15, x13, #1
    WORD $0xeb0f025f // cmp    x18, x15
    BEQ LBB5_1161
    WORD $0xeb0f005f // cmp    x2, x15
    BEQ LBB5_1161
    WORD $0xeb0f007f // cmp    x3, x15
    BEQ LBB5_1161
    WORD $0xf100044e // subs    x14, x2, #1
    BLT LBB5_1223
    WORD $0xeb0e007f // cmp    x3, x14
    BEQ LBB5_1223
    WORD $0xaa2203e4 // mvn    x4, x2
    B LBB5_1226
LBB5_1047:
    WORD $0xaa0f03ed // mov    x13, x15
    WORD $0x528001b2 // mov    w18, #13
    WORD $0x384015ae // ldrb    w14, [x13], #1
    WORD $0x7101c9df // cmp    w14, #114
    BNE LBB5_1050
    WORD $0x394005ee // ldrb    w14, [x15, #1]
    WORD $0x910009ed // add    x13, x15, #2
    WORD $0x7101d5df // cmp    w14, #117
    BNE LBB5_1050
    WORD $0x394009ed // ldrb    w13, [x15, #2]
    WORD $0x710195bf // cmp    w13, #101
    WORD $0x91000ded // add    x13, x15, #3
    WORD $0x1a9203f2 // csel    w18, wzr, w18, eq
LBB5_1050:
    WORD $0x5280014e // mov    w14, #10
    WORD $0xaa0003f1 // mov    x17, x0
    WORD $0xaa1081ce // orr    x14, x14, x16, lsl #32
    WORD $0xaa0d03ef // mov    x15, x13
    WORD $0xf801062e // str    x14, [x17], #16
    WORD $0xf9005111 // str    x17, [x8, #160]
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a9201a2 // csel    w2, w13, w18, eq
    CMP $0, R18_PLATFORM
    BEQ LBB5_1630
    B LBB5_341
LBB5_1051:
    WORD $0xaa0f03ed // mov    x13, x15
    WORD $0x528001b2 // mov    w18, #13
    WORD $0x384015ae // ldrb    w14, [x13], #1
    WORD $0x7101d5df // cmp    w14, #117
    BNE LBB5_1054
    WORD $0x394005ee // ldrb    w14, [x15, #1]
    WORD $0x910009ed // add    x13, x15, #2
    WORD $0x7101b1df // cmp    w14, #108
    BNE LBB5_1054
    WORD $0x394009ed // ldrb    w13, [x15, #2]
    WORD $0x7101b1bf // cmp    w13, #108
    WORD $0x91000ded // add    x13, x15, #3
    WORD $0x1a9203f2 // csel    w18, wzr, w18, eq
LBB5_1054:
    WORD $0xd3607e0e // lsl    x14, x16, #32
    WORD $0xaa0003f0 // mov    x16, x0
    WORD $0xaa0d03ef // mov    x15, x13
    WORD $0xf801060e // str    x14, [x16], #16
    WORD $0xf9005110 // str    x16, [x8, #160]
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a9201a2 // csel    w2, w13, w18, eq
    CMP $0, R18_PLATFORM
    BEQ LBB5_1630
    B LBB5_341
LBB5_1055:
    WORD $0xf940550b // ldr    x11, [x8, #168]
    WORD $0xaa0803f2 // mov    x18, x8
    WORD $0xf900040b // str    x11, [x0, #8]
    WORD $0x528000eb // mov    w11, #7
    WORD $0xf84b8e42 // ldr    x2, [x18, #184]!
    WORD $0xf85e8241 // ldur    x1, [x18, #-24]
    WORD $0xaa10816b // orr    x11, x11, x16, lsl #32
    WORD $0xf940064d // ldr    x13, [x18, #8]
    WORD $0xcb02002e // sub    x14, x1, x2
    WORD $0x91008030 // add    x16, x1, #32
    WORD $0xb10041df // cmn    x14, #16
    WORD $0x9344fdc3 // asr    x3, x14, #4
    WORD $0xfa4d1202 // ccmp    x16, x13, #2, ne
    WORD $0xf85f824d // ldur    x13, [x18, #-8]
    WORD $0xf900000b // str    x11, [x0]
    WORD $0x91004030 // add    x16, x1, #16
    WORD $0xf9005503 // str    x3, [x8, #168]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9a8183eb // csel    x11, xzr, x1, hi
    WORD $0xf81e8250 // stur    x16, [x18, #-24]
    WORD $0xf81f824d // stur    x13, [x18, #-8]
    CMP $0, R11
    BEQ LBB5_1070
    WORD $0x394001ed // ldrb    w13, [x15]
    WORD $0x710081bf // cmp    w13, #32
    BHI LBB5_1067
    WORD $0x5280002e // mov    w14, #1
    WORD $0xd284c000 // mov    x0, #9728
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0xf2c00020 // movk    x0, #1, lsl #32
    WORD $0xea0001df // tst    x14, x0
    BEQ LBB5_1067
    WORD $0x394005ed // ldrb    w13, [x15, #1]
    WORD $0x91000631 // add    x17, x17, #1
    WORD $0x710081bf // cmp    w13, #32
    BHI LBB5_1243
    WORD $0x5280002e // mov    w14, #1
    WORD $0xd284c00f // mov    x15, #9728
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0xf2c0002f // movk    x15, #1, lsl #32
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_1243
    WORD $0xf940490f // ldr    x15, [x8, #144]
    WORD $0xcb0f022d // sub    x13, x17, x15
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_1063
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x92800011 // mov    x17, #-1
    WORD $0x9acd222d // lsl    x13, x17, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_1066
    WORD $0x910101f1 // add    x17, x15, #64
LBB5_1063:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R0
    WORD $0xd101022f // sub    x15, x17, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00003 // ldr    q3, [x0, :lo12:.LCPI5_2]
LBB5_1064:
    WORD $0xadc215e4 // ldp    q4, q5, [x15, #64]!
    WORD $0x4e211c90 // and    v16.16b, v4.16b, v1.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411de6 // ldp    q6, q7, [x15, #32]
    WORD $0x4e211cb1 // and    v17.16b, v5.16b, v1.16b
    WORD $0x4e110011 // tbl    v17.16b, { v0.16b }, v17.16b
    WORD $0x6e248e04 // cmeq    v4.16b, v16.16b, v4.16b
    WORD $0x4e211cd2 // and    v18.16b, v6.16b, v1.16b
    WORD $0x4e120012 // tbl    v18.16b, { v0.16b }, v18.16b
    WORD $0x4e211cf3 // and    v19.16b, v7.16b, v1.16b
    WORD $0x4e130013 // tbl    v19.16b, { v0.16b }, v19.16b
    WORD $0x6e258e25 // cmeq    v5.16b, v17.16b, v5.16b
    WORD $0x6e268e46 // cmeq    v6.16b, v18.16b, v6.16b
    WORD $0x4e221ca5 // and    v5.16b, v5.16b, v2.16b
    WORD $0x4e0300a5 // tbl    v5.16b, { v5.16b }, v3.16b
    WORD $0x4e221c84 // and    v4.16b, v4.16b, v2.16b
    WORD $0x6e278e67 // cmeq    v7.16b, v19.16b, v7.16b
    WORD $0x4e030084 // tbl    v4.16b, { v4.16b }, v3.16b
    WORD $0x4e221cc6 // and    v6.16b, v6.16b, v2.16b
    WORD $0x4e0300c6 // tbl    v6.16b, { v6.16b }, v3.16b
    WORD $0x4e221ce7 // and    v7.16b, v7.16b, v2.16b
    WORD $0x4e71b8a5 // addv    h5, v5.8h
    WORD $0x4e0300e7 // tbl    v7.16b, { v7.16b }, v3.16b
    WORD $0x4e71b884 // addv    h4, v4.8h
    WORD $0x1e2600ad // fmov    w13, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e26008e // fmov    w14, s4
    WORD $0x1e2600b1 // fmov    w17, s5
    WORD $0x1e2600c0 // fmov    w0, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa1181cd // orr    x13, x14, x17, lsl #32
    WORD $0xaa00c1ad // orr    x13, x13, x0, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_1064
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa909350f // stp    x15, x13, [x8, #144]
LBB5_1066:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d01f1 // add    x17, x15, x13
    WORD $0x3840162d // ldrb    w13, [x17], #1
LBB5_1067:
    WORD $0x710175bf // cmp    w13, #93
    BNE LBB5_1244
LBB5_1068:
    WORD $0xb940d10b // ldr    w11, [x8, #208]
    WORD $0x8b03104d // add    x13, x2, x3, lsl #4
    WORD $0x1100056b // add    w11, w11, #1
    WORD $0xb900d10b // str    w11, [x8, #208]
LBB5_1069:
    WORD $0xcb01020b // sub    x11, x16, x1
    WORD $0xf94005ae // ldr    x14, [x13, #8]
    WORD $0xd344fd6b // lsr    x11, x11, #4
    WORD $0xf900550e // str    x14, [x8, #168]
    WORD $0xb9000c2b // str    w11, [x1, #12]
    WORD $0xf940002b // ldr    x11, [x1]
    WORD $0xb90009bf // str    wzr, [x13, #8]
    WORD $0xf940590d // ldr    x13, [x8, #176]
    WORD $0xb940e50e // ldr    w14, [x8, #228]
    WORD $0x92609d6b // and    x11, x11, #0xffffffff000000ff
    WORD $0xeb0e01bf // cmp    x13, x14
    WORD $0xf900002b // str    x11, [x1]
    BLS LBB5_1923
    B LBB5_1701
LBB5_1070:
    WORD $0x52800162 // mov    w2, #11
    B LBB5_341
LBB5_1071:
    WORD $0xf9403911 // ldr    x17, [x8, #112]
    WORD $0x2518e3e0 // ptrue    p0.b
    TST $(1<<5), R17
    BNE LBB5_1616
    WORD $0xaa0f03e1 // mov    x1, x15
    WORD $0x2538cb80 // mov    z0.b, #92
    WORD $0x2538c441 // mov    z1.b, #34
    WORD $0x2518e402 // pfalse    p2.b
    B LBB5_1075
LBB5_1073:
    WORD $0x25904063 // brkb    p3.b, p0/z, p3.b
    WORD $0x25414063 // ands    p3.b, p0/z, p3.b, p1.b
    BNE LBB5_1080
LBB5_1074:
    WORD $0x91008021 // add    x1, x1, #32
LBB5_1075:
    WORD $0xa400a022 // ld1b    { z2.b }, p0/z, [x1]
    WORD $0x2401a043 // cmpeq    p3.b, p0/z, z2.b, z1.b
    WORD $0x2400a041 // cmpeq    p1.b, p0/z, z2.b, z0.b
    WORD $0x25834c64 // mov    p4.b, p3.b
    WORD $0x25c24025 // orrs    p5.b, p0/z, p1.b, p2.b
    BEQ LBB5_1077
    WORD $0x259040a4 // brkb    p4.b, p0/z, p5.b
    WORD $0x25034084 // and    p4.b, p0/z, p4.b, p3.b
LBB5_1077:
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_1627
    WORD $0x2550c060 // ptest    p0, p3.b
    BNE LBB5_1073
    WORD $0x2550c020 // ptest    p0, p1.b
    BEQ LBB5_1074
LBB5_1080:
    WORD $0x25904021 // brkb    p1.b, p0/z, p1.b
    WORD $0x5299fa0e // mov    w14, #53200
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0x52832320 // mov    w0, #6425
    WORD $0x8b0d0037 // add    x23, x1, x13
    WORD $0x5288c8c2 // mov    w2, #17990
    WORD $0x52872724 // mov    w4, #14649
    WORD $0x52848013 // mov    w19, #9216
    WORD $0x72b9f9ee // movk    w14, #53199, lsl #16
    WORD $0x3201c3f2 // mov    w18, #-2139062144
    WORD $0x72a32320 // movk    w0, #6425, lsl #16
    WORD $0x3202c7e1 // mov    w1, #-1061109568
    WORD $0x72a8c8c2 // movk    w2, #17990, lsl #16
    WORD $0x3203cbe3 // mov    w3, #-522133280
    WORD $0x72a72724 // movk    w4, #14649, lsl #16
    WORD $0x3200c3e5 // mov    w5, #16843009
    WORD $0x5297fde6 // mov    w6, #49135
    WORD $0x528017a7 // mov    w7, #189
    WORD $0x72bf9413 // movk    w19, #64672, lsl #16
    WORD $0xaa1703fa // mov    x26, x23
    WORD $0xaa1703f9 // mov    x25, x23
    ADR ESCAPED_TAB, R20
    WORD $0x91000294 // add    x20, x20, :lo12:ESCAPED_TAB
    WORD $0x2538cb80 // mov    z0.b, #92
    WORD $0x2538c441 // mov    z1.b, #34
    WORD $0x2518e401 // pfalse    p1.b
    WORD $0x2538c3e2 // mov    z2.b, #31
LBB5_1081:
    WORD $0x3940074d // ldrb    w13, [x26, #1]
    WORD $0xf101d5bf // cmp    x13, #117
    BEQ LBB5_1084
    WORD $0x386d6a8d // ldrb    w13, [x20, x13]
    CMP $0, R13
    BEQ LBB5_1457
    WORD $0x3800172d // strb    w13, [x25], #1
    WORD $0x91000b55 // add    x21, x26, #2
    WORD $0xaa1903f6 // mov    x22, x25
    B LBB5_1104
LBB5_1084:
    WORD $0xb840234d // ldur    w13, [x26, #2]
    WORD $0x0a2d0255 // bic    w21, w18, w13
    WORD $0x0b0e01b6 // add    w22, w13, w14
    WORD $0x6a1602bf // tst    w21, w22
    BNE LBB5_1681
    WORD $0x0b0001b6 // add    w22, w13, w0
    WORD $0x2a0d02d6 // orr    w22, w22, w13
    WORD $0x7201c2df // tst    w22, #0x80808080
    BNE LBB5_1681
    WORD $0x1200d9b6 // and    w22, w13, #0x7f7f7f7f
    WORD $0x4b160038 // sub    w24, w1, w22
    WORD $0x0b0202db // add    w27, w22, w2
    WORD $0x0a1b0318 // and    w24, w24, w27
    WORD $0x6a15031f // tst    w24, w21
    BNE LBB5_1681
    WORD $0x4b160078 // sub    w24, w3, w22
    WORD $0x0b0402d6 // add    w22, w22, w4
    WORD $0x0a160316 // and    w22, w24, w22
    WORD $0x6a1502df // tst    w22, w21
    BNE LBB5_1681
    WORD $0x5ac009ad // rev    w13, w13
    WORD $0x1200cdb5 // and    w21, w13, #0xf0f0f0f
    WORD $0x0a6d10ad // bic    w13, w5, w13, lsr #4
    WORD $0x2a0d0dad // orr    w13, w13, w13, lsl #3
    WORD $0x0b1501ad // add    w13, w13, w21
    WORD $0x2a4d11ad // orr    w13, w13, w13, lsr #4
    WORD $0x53105db5 // ubfx    w21, w13, #16, #8
    WORD $0x12001dad // and    w13, w13, #0xff
    WORD $0x2a1521b7 // orr    w23, w13, w21, lsl #8
    WORD $0x91001b55 // add    x21, x26, #6
    WORD $0x710202ff // cmp    w23, #128
    BLO LBB5_1136
    WORD $0x91001336 // add    x22, x25, #4
LBB5_1090:
    WORD $0x711ffeff // cmp    w23, #2047
    BLS LBB5_1138
    WORD $0x51403aed // sub    w13, w23, #14, lsl #12
    WORD $0x312005bf // cmn    w13, #2049
    BLS LBB5_1102
    WORD $0x530a7eed // lsr    w13, w23, #10
    WORD $0x7100d9bf // cmp    w13, #54
    BHI LBB5_1139
    WORD $0x394002ad // ldrb    w13, [x21]
    WORD $0x710171bf // cmp    w13, #92
    BNE LBB5_1139
    WORD $0x394006ad // ldrb    w13, [x21, #1]
    WORD $0x7101d5bf // cmp    w13, #117
    BNE LBB5_1139
    WORD $0xb84022ad // ldur    w13, [x21, #2]
    WORD $0x0a2d0258 // bic    w24, w18, w13
    WORD $0x0b0e01b9 // add    w25, w13, w14
    WORD $0x6a19031f // tst    w24, w25
    BNE LBB5_1700
    WORD $0x0b0001b9 // add    w25, w13, w0
    WORD $0x2a0d0339 // orr    w25, w25, w13
    WORD $0x7201c33f // tst    w25, #0x80808080
    BNE LBB5_1700
    WORD $0x1200d9b9 // and    w25, w13, #0x7f7f7f7f
    WORD $0x4b19003a // sub    w26, w1, w25
    WORD $0x0b02033b // add    w27, w25, w2
    WORD $0x0a1b035a // and    w26, w26, w27
    WORD $0x6a18035f // tst    w26, w24
    BNE LBB5_1700
    WORD $0x4b19007a // sub    w26, w3, w25
    WORD $0x0b040339 // add    w25, w25, w4
    WORD $0x0a190359 // and    w25, w26, w25
    WORD $0x6a18033f // tst    w25, w24
    BNE LBB5_1700
    WORD $0x5ac009ad // rev    w13, w13
    WORD $0x91001ab5 // add    x21, x21, #6
    WORD $0x1200cdb8 // and    w24, w13, #0xf0f0f0f
    WORD $0x0a6d10ad // bic    w13, w5, w13, lsr #4
    WORD $0x2a0d0dad // orr    w13, w13, w13, lsl #3
    WORD $0x0b1801ad // add    w13, w13, w24
    WORD $0x2a4d11b8 // orr    w24, w13, w13, lsr #4
    WORD $0x53087f0d // lsr    w13, w24, #8
    WORD $0x12181dad // and    w13, w13, #0xff00
    WORD $0x514039b9 // sub    w25, w13, #14, lsl #12
    WORD $0x33001f0d // bfxil    w13, w24, #0, #8
    WORD $0x3110073f // cmn    w25, #1025
    BHI LBB5_1140
    WORD $0x781fc2c6 // sturh    w6, [x22, #-4]
    WORD $0x2a0d03f7 // mov    w23, w13
    WORD $0x381fe2c7 // sturb    w7, [x22, #-2]
    WORD $0x91000ed6 // add    x22, x22, #3
    WORD $0x710201bf // cmp    w13, #128
    BHS LBB5_1090
    WORD $0xd10012d9 // sub    x25, x22, #4
    B LBB5_1137
LBB5_1102:
    WORD $0x530c7eed // lsr    w13, w23, #12
    WORD $0x52801018 // mov    w24, #128
    WORD $0x52801019 // mov    w25, #128
    WORD $0x321b09ad // orr    w13, w13, #0xe0
    WORD $0x33062ef8 // bfxil    w24, w23, #6, #6
    WORD $0x330016f9 // bfxil    w25, w23, #0, #6
    WORD $0xd10006d7 // sub    x23, x22, #1
    WORD $0x381fc2cd // sturb    w13, [x22, #-4]
    WORD $0x381fd2d8 // sturb    w24, [x22, #-3]
    WORD $0x381fe2d9 // sturb    w25, [x22, #-2]
LBB5_1103:
    WORD $0xaa1703f6 // mov    x22, x23
LBB5_1104:
    WORD $0x394002ad // ldrb    w13, [x21]
    WORD $0xaa1503f7 // mov    x23, x21
    WORD $0xaa1503fa // mov    x26, x21
    WORD $0xaa1603f9 // mov    x25, x22
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1081
    WORD $0xaa1d03fb // mov    x27, x29
    WORD $0xaa1f03f8 // mov    x24, xzr
    WORD $0xa41842a3 // ld1b    { z3.b }, p0/z, [x21, x24]
    WORD $0x25814422 // mov    p2.b, p1.b
    TST $(1<<5), R17
    BEQ LBB5_1109
    B LBB5_1107
LBB5_1106:
    WORD $0xe41842c3 // st1b    { z3.b }, p0, [x22, x24]
    WORD $0x91008318 // add    x24, x24, #32
    WORD $0xa41842a3 // ld1b    { z3.b }, p0/z, [x21, x24]
    WORD $0x25814422 // mov    p2.b, p1.b
    TST $(1<<5), R17
    BEQ LBB5_1109
LBB5_1107:
    WORD $0x2402a062 // cmpeq    p2.b, p0/z, z3.b, z2.b
    WORD $0x2401a064 // cmpeq    p4.b, p0/z, z3.b, z1.b
    WORD $0x2400a063 // cmpeq    p3.b, p0/z, z3.b, z0.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BNE LBB5_1110
LBB5_1108:
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BEQ LBB5_1111
    B LBB5_1668
LBB5_1109:
    WORD $0x2401a064 // cmpeq    p4.b, p0/z, z3.b, z1.b
    WORD $0x2400a063 // cmpeq    p3.b, p0/z, z3.b, z0.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BEQ LBB5_1108
LBB5_1110:
    WORD $0x259040c5 // brkb    p5.b, p0/z, p6.b
    WORD $0x250440a5 // and    p5.b, p0/z, p5.b, p4.b
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_1668
LBB5_1111:
    WORD $0x2550c080 // ptest    p0, p4.b
    WORD $0x25824845 // mov    p5.b, p2.b
    WORD $0x1a9f07ed // cset    w13, ne
    BEQ LBB5_1113
    WORD $0x25904085 // brkb    p5.b, p0/z, p4.b
    WORD $0x250240a5 // and    p5.b, p0/z, p5.b, p2.b
LBB5_1113:
    TST $(1<<5), R17
    BEQ LBB5_1115
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_1696
LBB5_1115:
    CMP $0, R13
    BEQ LBB5_1117
    WORD $0x25904082 // brkb    p2.b, p0/z, p4.b
    WORD $0x25034043 // and    p3.b, p0/z, p2.b, p3.b
LBB5_1117:
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_1106
    WORD $0x8b1802b7 // add    x23, x21, x24
    WORD $0x8b1802d9 // add    x25, x22, x24
    WORD $0xaa1703fa // mov    x26, x23
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x394002ed // ldrb    w13, [x23]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1081
LBB5_1119:
    WORD $0x8b1802d9 // add    x25, x22, x24
    WORD $0x8b1802ba // add    x26, x21, x24
    WORD $0x3900032d // strb    w13, [x25]
    WORD $0x3940074d // ldrb    w13, [x26, #1]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1128
    WORD $0x3900072d // strb    w13, [x25, #1]
    WORD $0x39400b4d // ldrb    w13, [x26, #2]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1129
    WORD $0x8b1802b7 // add    x23, x21, x24
    WORD $0x39000b2d // strb    w13, [x25, #2]
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x39400ef9 // ldrb    w25, [x23, #3]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1130
    WORD $0x39000db9 // strb    w25, [x13, #3]
    WORD $0x394012f9 // ldrb    w25, [x23, #4]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1131
    WORD $0x8b1802b7 // add    x23, x21, x24
    WORD $0x390011b9 // strb    w25, [x13, #4]
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x394016f9 // ldrb    w25, [x23, #5]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1132
    WORD $0x390015b9 // strb    w25, [x13, #5]
    WORD $0x39401af9 // ldrb    w25, [x23, #6]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1133
    WORD $0x8b1802b7 // add    x23, x21, x24
    WORD $0x390019b9 // strb    w25, [x13, #6]
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x39401ef9 // ldrb    w25, [x23, #7]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1134
    WORD $0x39001db9 // strb    w25, [x13, #7]
    WORD $0x91002318 // add    x24, x24, #8
    WORD $0x394022ed // ldrb    w13, [x23, #8]
    WORD $0x710171bf // cmp    w13, #92
    BNE LBB5_1119
    WORD $0x8b1802b7 // add    x23, x21, x24
    WORD $0x8b1802d9 // add    x25, x22, x24
    WORD $0xd10006fa // sub    x26, x23, #1
    B LBB5_1135
LBB5_1128:
    WORD $0x91000757 // add    x23, x26, #1
    WORD $0x91000739 // add    x25, x25, #1
    B LBB5_1135
LBB5_1129:
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0x8b1802d5 // add    x21, x22, x24
    WORD $0x910005ba // add    x26, x13, #1
    WORD $0x910009b7 // add    x23, x13, #2
    WORD $0x91000ab9 // add    x25, x21, #2
    B LBB5_1135
LBB5_1130:
    WORD $0x91000afa // add    x26, x23, #2
    WORD $0x91000ef7 // add    x23, x23, #3
    WORD $0x91000db9 // add    x25, x13, #3
    B LBB5_1135
LBB5_1131:
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0x8b1802d5 // add    x21, x22, x24
    WORD $0x91000dba // add    x26, x13, #3
    WORD $0x910011b7 // add    x23, x13, #4
    WORD $0x910012b9 // add    x25, x21, #4
    B LBB5_1135
LBB5_1132:
    WORD $0x910012fa // add    x26, x23, #4
    WORD $0x910016f7 // add    x23, x23, #5
    WORD $0x910015b9 // add    x25, x13, #5
    B LBB5_1135
LBB5_1133:
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0x8b1802d5 // add    x21, x22, x24
    WORD $0x910015ba // add    x26, x13, #5
    WORD $0x910019b7 // add    x23, x13, #6
    WORD $0x91001ab9 // add    x25, x21, #6
    B LBB5_1135
LBB5_1134:
    WORD $0x91001afa // add    x26, x23, #6
    WORD $0x91001ef7 // add    x23, x23, #7
    WORD $0x91001db9 // add    x25, x13, #7
LBB5_1135:
    WORD $0x9100075a // add    x26, x26, #1
    WORD $0xaa1b03fd // mov    x29, x27
    B LBB5_1081
LBB5_1136:
    WORD $0x2a1703ed // mov    w13, w23
LBB5_1137:
    WORD $0x3800172d // strb    w13, [x25], #1
    WORD $0xaa1903f6 // mov    x22, x25
    B LBB5_1104
LBB5_1138:
    WORD $0x53067eed // lsr    w13, w23, #6
    WORD $0x52801018 // mov    w24, #128
    WORD $0x321a05ad // orr    w13, w13, #0xc0
    WORD $0x330016f8 // bfxil    w24, w23, #0, #6
    WORD $0xd1000ad7 // sub    x23, x22, #2
    WORD $0x381fc2cd // sturb    w13, [x22, #-4]
    WORD $0x381fd2d8 // sturb    w24, [x22, #-3]
    B LBB5_1103
LBB5_1139:
    WORD $0xd10006cd // sub    x13, x22, #1
    WORD $0x781fc2c6 // sturh    w6, [x22, #-4]
    WORD $0x381fe2c7 // sturb    w7, [x22, #-2]
    WORD $0xaa0d03f6 // mov    x22, x13
    B LBB5_1104
LBB5_1140:
    WORD $0x0b1729ad // add    w13, w13, w23, lsl #10
    WORD $0x52801019 // mov    w25, #128
    WORD $0x0b1301ad // add    w13, w13, w19
    WORD $0x5280101a // mov    w26, #128
    WORD $0x53127db7 // lsr    w23, w13, #18
    WORD $0x5280101b // mov    w27, #128
    WORD $0x321c0ef7 // orr    w23, w23, #0xf0
    WORD $0x3300171b // bfxil    w27, w24, #0, #6
    WORD $0x330c45b9 // bfxil    w25, w13, #12, #6
    WORD $0x33062dba // bfxil    w26, w13, #6, #6
    WORD $0x381fc2d7 // sturb    w23, [x22, #-4]
    WORD $0x381fd2d9 // sturb    w25, [x22, #-3]
    WORD $0x381fe2da // sturb    w26, [x22, #-2]
    WORD $0x381ff2db // sturb    w27, [x22, #-1]
    B LBB5_1104
LBB5_1141:
    WORD $0xaa0503f1 // mov    x17, x5
    WORD $0x52800062 // mov    w2, #3
    WORD $0x38402e20 // ldrb    w0, [x17, #2]!
    WORD $0x5100c00d // sub    w13, w0, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1598
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x7100c01f // cmp    w0, #48
    BNE LBB5_1144
LBB5_1143:
    WORD $0x8b0d00ae // add    x14, x5, x13
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x39400dc0 // ldrb    w0, [x14, #3]
    WORD $0x7100c01f // cmp    w0, #48
    BEQ LBB5_1143
LBB5_1144:
    WORD $0x7101141f // cmp    w0, #69
    BEQ LBB5_1164
    WORD $0x7101941f // cmp    w0, #101
    BEQ LBB5_1164
    WORD $0x8b0d00ae // add    x14, x5, x13
    WORD $0xaa1f03f2 // mov    x18, xzr
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0xaa1f03e3 // mov    x3, xzr
    WORD $0x910009d3 // add    x19, x14, #2
    WORD $0x4b0d03e7 // neg    w7, w13
LBB5_1147:
    WORD $0x5280022d // mov    w13, #17
    WORD $0xcb1201b1 // sub    x17, x13, x18
    WORD $0xf100063f // cmp    x17, #1
    BLT LBB5_1163
    WORD $0x4b1201a2 // sub    w2, w13, w18
    WORD $0x5280024d // mov    w13, #18
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x8b110271 // add    x17, x19, x17
    WORD $0xcb1201ad // sub    x13, x13, x18
    WORD $0x52800152 // mov    w18, #10
LBB5_1149:
    WORD $0x39400260 // ldrb    w0, [x19]
    WORD $0x5100c00e // sub    w14, w0, #48
    WORD $0x710025df // cmp    w14, #9
    BHI LBB5_1165
    WORD $0x9b12006e // madd    x14, x3, x18, x0
    WORD $0xd1000694 // sub    x20, x20, #1
    WORD $0x91000673 // add    x19, x19, #1
    WORD $0xd100c1c3 // sub    x3, x14, #48
    WORD $0x8b1401ae // add    x14, x13, x20
    WORD $0xf10005df // cmp    x14, #1
    BGT LBB5_1149
    WORD $0x39400220 // ldrb    w0, [x17]
    B LBB5_1167
LBB5_1152:
    WORD $0x7101167f // cmp    w19, #69
    BEQ LBB5_954
    WORD $0x7101967f // cmp    w19, #101
    BEQ LBB5_954
    CMP $0, R7
    BNE LBB5_1645
    CMP $0, R1
    BEQ LBB5_1239
    WORD $0xb24107ec // mov    x12, #-9223372036854775807
    WORD $0xeb0c007f // cmp    x3, x12
    BLO LBB5_1315
    WORD $0x9e630060 // ucvtf    d0, x3
    WORD $0x2a1f03f2 // mov    w18, wzr
LBB5_1158:
    WORD $0x9e66000c // fmov    x12, d0
    WORD $0xd2410181 // eor    x1, x12, #0x8000000000000000
    B LBB5_1604
LBB5_1159:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa1f03fd // mov    x29, xzr
    TST $(1<<0), R1
    BEQ LBB5_1240
    WORD $0x5280016c // mov    w12, #11
    B LBB5_1602
LBB5_1161:
    WORD $0xcb0d03e4 // neg    x4, x13
LBB5_1162:
    WORD $0x2a1f03f2 // mov    w18, wzr
    WORD $0xaa0403ed // mov    x13, x4
    TST $(1<<63), R4
    BEQ LBB5_1227
    B LBB5_1226
LBB5_1163:
    WORD $0x2a1f03e2 // mov    w2, wzr
    B LBB5_1166
LBB5_1164:
    WORD $0x8b0d00ad // add    x13, x5, x13
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x2a1f03e7 // mov    w7, wzr
    WORD $0xaa1f03e3 // mov    x3, xzr
    WORD $0x910009b1 // add    x17, x13, #2
    B LBB5_954
LBB5_1165:
    WORD $0x4b1403e2 // neg    w2, w20
LBB5_1166:
    WORD $0xaa1303f1 // mov    x17, x19
LBB5_1167:
    WORD $0x4b0200e7 // sub    w7, w7, w2
    WORD $0x5100c00d // sub    w13, w0, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1170
LBB5_1168:
    WORD $0x38401e20 // ldrb    w0, [x17, #1]!
    WORD $0x5100c00d // sub    w13, w0, #48
    WORD $0x710029bf // cmp    w13, #10
    BLO LBB5_1168
    WORD $0x52800024 // mov    w4, #1
LBB5_1170:
    WORD $0x52801bed // mov    w13, #223
    WORD $0x0a0d000d // and    w13, w0, w13
    WORD $0x710115bf // cmp    w13, #69
    BEQ LBB5_954
LBB5_1171:
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x1280000d // mov    w13, #-1
    WORD $0xd374fc6e // lsr    x14, x3, #52
    WORD $0x5a8d15b2 // cneg    w18, w13, eq
    CMP $0, R14
    BNE LBB5_1182
    WORD $0x9e630060 // ucvtf    d0, x3
    WORD $0x531f7e4d // lsr    w13, w18, #31
    WORD $0x9e66000e // fmov    x14, d0
    WORD $0xaa0dfdcd // orr    x13, x14, x13, lsl #63
    WORD $0x9e6701a0 // fmov    d0, x13
    CMP $0, R7
    BEQ LBB5_1597
    CMP $0, R3
    BEQ LBB5_1597
    WORD $0x510004ed // sub    w13, w7, #1
    WORD $0x710091bf // cmp    w13, #36
    BHI LBB5_1180
    WORD $0x2a0703ed // mov    w13, w7
    WORD $0x71005cff // cmp    w7, #23
    BLO LBB5_1177
    WORD $0x510058ed // sub    w13, w7, #22
    ADR P10_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:P10_TAB
    WORD $0xfc6d59c1 // ldr    d1, [x14, w13, uxtw #3]
    WORD $0x528002cd // mov    w13, #22
    WORD $0x1e600820 // fmul    d0, d1, d0
LBB5_1177:
    ADR LCPI5_3, R14
    WORD $0xfd4001c1 // ldr    d1, [x14, :lo12:.LCPI5_3]
    WORD $0x1e612000 // fcmp    d0, d1
    BGT LBB5_1183
    ADR LCPI5_4, R14
    WORD $0xfd4001c1 // ldr    d1, [x14, :lo12:.LCPI5_4]
    WORD $0x1e612000 // fcmp    d0, d1
    BMI LBB5_1183
    ADR P10_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:P10_TAB
    WORD $0xfc6d59c1 // ldr    d1, [x14, w13, uxtw #3]
    B LBB5_1596
LBB5_1180:
    WORD $0x310058ff // cmn    w7, #22
    BLO LBB5_1182
    WORD $0x4b0703ed // neg    w13, w7
    ADR P10_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:P10_TAB
    WORD $0xfc6d59c1 // ldr    d1, [x14, w13, uxtw #3]
    WORD $0x1e611800 // fdiv    d0, d0, d1
    B LBB5_1597
LBB5_1182:
    WORD $0x510570ed // sub    w13, w7, #348
    WORD $0x310ae1bf // cmn    w13, #696
    BLO LBB5_1191
LBB5_1183:
    WORD $0x110570e0 // add    w0, w7, #348
    ADR POW10_M128_TAB, R2
    WORD $0x91000042 // add    x2, x2, :lo12:POW10_M128_TAB
    WORD $0x528a4d4d // mov    w13, #21098
    WORD $0x8b20504e // add    x14, x2, w0, uxtw #4
    WORD $0x72a0006d // movk    w13, #3, lsl #16
    WORD $0xdac01074 // clz    x20, x3
    WORD $0x1b0d7cf3 // mul    w19, w7, w13
    WORD $0xf94005c7 // ldr    x7, [x14, #8]
    WORD $0x9ad4206d // lsl    x13, x3, x20
    WORD $0x13107e6e // asr    w14, w19, #16
    WORD $0x1110fdce // add    w14, w14, #1087
    WORD $0xaa2d03f8 // mvn    x24, x13
    WORD $0x9bcd7cf5 // umulh    x21, x7, x13
    WORD $0x93407dd3 // sxtw    x19, w14
    WORD $0x9b0d7cf6 // mul    x22, x7, x13
    WORD $0xcb140274 // sub    x20, x19, x20
    WORD $0x924022b7 // and    x23, x21, #0x1ff
    WORD $0xeb1802df // cmp    x22, x24
    BLS LBB5_1188
    WORD $0xf107feff // cmp    x23, #511
    BNE LBB5_1188
    WORD $0xd37cec0e // lsl    x14, x0, #4
    WORD $0xf86e684e // ldr    x14, [x2, x14]
    WORD $0x9bcd7dd7 // umulh    x23, x14, x13
    WORD $0x9b0d7dcd // mul    x13, x14, x13
    WORD $0xab1602f6 // adds    x22, x23, x22
    WORD $0x9a9536b5 // cinc    x21, x21, hs
    WORD $0xeb1801bf // cmp    x13, x24
    WORD $0x924022b7 // and    x23, x21, #0x1ff
    BLS LBB5_1188
    WORD $0xb10006df // cmn    x22, #1
    BNE LBB5_1188
    WORD $0xf107feff // cmp    x23, #511
    BEQ LBB5_1191
LBB5_1188:
    WORD $0xd37ffead // lsr    x13, x21, #63
    WORD $0xaa1702d6 // orr    x22, x22, x23
    WORD $0x910025ae // add    x14, x13, #9
    WORD $0x9ace26b5 // lsr    x21, x21, x14
    CMP $0, R22
    BNE LBB5_1190
    WORD $0x924006ae // and    x14, x21, #0x3
    WORD $0xf10005df // cmp    x14, #1
    BEQ LBB5_1191
LBB5_1190:
    WORD $0x924002ae // and    x14, x21, #0x1
    WORD $0x8b0d028d // add    x13, x20, x13
    WORD $0x8b1501d5 // add    x21, x14, x21
    WORD $0xd376feb6 // lsr    x22, x21, #54
    WORD $0xf10002df // cmp    x22, #0
    WORD $0x1a9f17ee // cset    w14, eq
    WORD $0xcb0e01ad // sub    x13, x13, x14
    WORD $0xd11ffdae // sub    x14, x13, #2047
    WORD $0xb11ff9df // cmn    x14, #2046
    BHS LBB5_1228
LBB5_1191:
    WORD $0xf9402903 // ldr    x3, [x8, #80]
    WORD $0xcb050227 // sub    x7, x17, x5
    WORD $0xf9402100 // ldr    x0, [x8, #64]
    CMP $0, R3
    BEQ LBB5_1203
    WORD $0xaa1f03e1 // mov    x1, xzr
    WORD $0x0460e3e2 // cnth    x2
    WORD $0xeb02007f // cmp    x3, x2
    BLO LBB5_1201
    WORD $0xaa1f03e1 // mov    x1, xzr
    WORD $0x04bf504d // rdvl    x13, #2
    WORD $0xeb0d007f // cmp    x3, x13
    BHS LBB5_1197
LBB5_1194:
    WORD $0xcb0203ee // neg    x14, x2
    WORD $0xaa0103ed // mov    x13, x1
    WORD $0x8a0e0061 // and    x1, x3, x14
    WORD $0x2558e3e0 // ptrue    p0.h
    WORD $0x2578c000 // mov    z0.h, #0
LBB5_1195:
    WORD $0xe42d4000 // st1b    { z0.h }, p0, [x0, x13]
    WORD $0x8b0201ad // add    x13, x13, x2
    WORD $0xeb0d003f // cmp    x1, x13
    BNE LBB5_1195
    WORD $0xeb01007f // cmp    x3, x1
    BNE LBB5_1201
    B LBB5_1203
LBB5_1197:
    WORD $0x04bf57ce // rdvl    x14, #-2
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x8a0e0061 // and    x1, x3, x14
    WORD $0x04bf5044 // rdvl    x4, #2
    WORD $0x04205033 // addvl    x19, x0, #1
    WORD $0x2518e3e0 // ptrue    p0.b
    WORD $0x2538c000 // mov    z0.b, #0
LBB5_1198:
    WORD $0xe40d4000 // st1b    { z0.b }, p0, [x0, x13]
    WORD $0xe40d4260 // st1b    { z0.b }, p0, [x19, x13]
    WORD $0x8b0401ad // add    x13, x13, x4
    WORD $0xeb0d003f // cmp    x1, x13
    BNE LBB5_1198
    WORD $0xeb01006d // subs    x13, x3, x1
    BEQ LBB5_1203
    WORD $0xeb0201bf // cmp    x13, x2
    BHS LBB5_1194
LBB5_1201:
    WORD $0xcb01006d // sub    x13, x3, x1
    WORD $0x8b010001 // add    x1, x0, x1
LBB5_1202:
    WORD $0xf10005ad // subs    x13, x13, #1
    WORD $0x3800143f // strb    wzr, [x1], #1
    BNE LBB5_1202
LBB5_1203:
    WORD $0x394000b6 // ldrb    w22, [x5]
    WORD $0xaa1f03f5 // mov    x21, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0xb9001ff9 // str    w25, [sp, #28]
    WORD $0x7100b6df // cmp    w22, #45
    WORD $0x1a9f17f3 // cset    w19, eq
    WORD $0xeb1300ff // cmp    x7, x19
    BLE LBB5_1595
    WORD $0xb9001bf6 // str    w22, [sp, #24]
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x2a1f03f8 // mov    w24, wzr
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0x2a1f03f5 // mov    w21, wzr
    WORD $0x2a1f03f6 // mov    w22, wzr
    WORD $0x2a1f03f7 // mov    w23, wzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x2a1f03f4 // mov    w20, wzr
    WORD $0x5280003a // mov    w26, #1
    B LBB5_1207
LBB5_1205:
    WORD $0x11000718 // add    w24, w24, #1
    WORD $0x382d681b // strb    w27, [x0, x13]
    WORD $0x2a1803f5 // mov    w21, w24
    WORD $0x2a1803f6 // mov    w22, w24
    WORD $0x2a1803f7 // mov    w23, w24
    WORD $0x2a1803f9 // mov    w25, w24
LBB5_1206:
    WORD $0x91000673 // add    x19, x19, #1
    WORD $0xeb07027f // cmp    x19, x7
    WORD $0x1a9fa7fa // cset    w26, lt
    WORD $0xeb1300ff // cmp    x7, x19
    BEQ LBB5_1217
LBB5_1207:
    WORD $0x387368bb // ldrb    w27, [x5, x19]
    WORD $0x5100c36d // sub    w13, w27, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1212
    WORD $0x7100c37f // cmp    w27, #48
    BNE LBB5_1214
    CMP $0, R22
    BEQ LBB5_1216
    WORD $0x93407ead // sxtw    x13, w21
    WORD $0xeb0d007f // cmp    x3, x13
    BHI LBB5_1205
    WORD $0x2a1503f6 // mov    w22, w21
    WORD $0x2a1503f7 // mov    w23, w21
    WORD $0x2a1503f9 // mov    w25, w21
    B LBB5_1206
LBB5_1212:
    WORD $0x7100bb7f // cmp    w27, #46
    BNE LBB5_1218
    WORD $0x52800034 // mov    w20, #1
    WORD $0x2a1903e4 // mov    w4, w25
    B LBB5_1206
LBB5_1214:
    WORD $0x93407eed // sxtw    x13, w23
    WORD $0xeb0d007f // cmp    x3, x13
    BHI LBB5_1205
    WORD $0x52800022 // mov    w2, #1
    WORD $0x2a1703f9 // mov    w25, w23
    B LBB5_1206
LBB5_1216:
    WORD $0x2a1f03f7 // mov    w23, wzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x51000484 // sub    w4, w4, #1
    B LBB5_1206
LBB5_1217:
    WORD $0xaa0703f3 // mov    x19, x7
LBB5_1218:
    WORD $0x7100029f // cmp    w20, #0
    WORD $0xb9401bf6 // ldr    w22, [sp, #24]
    WORD $0x1a8402a4 // csel    w4, w21, w4, eq
    TST $(1<<0), R26
    BEQ LBB5_1255
    WORD $0x387368ad // ldrb    w13, [x5, x19]
    WORD $0x321b01ad // orr    w13, w13, #0x20
    WORD $0x710195bf // cmp    w13, #101
    BNE LBB5_1255
    WORD $0x2a1303ed // mov    w13, w19
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x386d68a1 // ldrb    w1, [x5, x13]
    WORD $0x7100b43f // cmp    w1, #45
    BEQ LBB5_1247
    WORD $0x52800025 // mov    w5, #1
    WORD $0x7100ac3f // cmp    w1, #43
    BNE LBB5_1248
    WORD $0x11000a6d // add    w13, w19, #2
    B LBB5_1248
LBB5_1223:
    WORD $0xaa03024e // orr    x14, x18, x3
    WORD $0xd37ffdcf // lsr    x15, x14, #63
    WORD $0x520001ef // eor    w15, w15, #0x1
    TST $(1<<63), R14
    BNE LBB5_1245
    WORD $0xeb03025f // cmp    x18, x3
    BLT LBB5_1245
    WORD $0xaa3203e4 // mvn    x4, x18
LBB5_1226:
    WORD $0xaa2403e4 // mvn    x4, x4
    WORD $0x52800072 // mov    w18, #3
    WORD $0x9280004d // mov    x13, #-3
LBB5_1227:
    WORD $0x5280036e // mov    w14, #27
    WORD $0x8b2141ad // add    x13, x13, w1, uxtw
    WORD $0xaa1081ce // orr    x14, x14, x16, lsl #32
    WORD $0x8b0400af // add    x15, x5, x4
    WORD $0xa900340e // stp    x14, x13, [x0]
    WORD $0xf9405100 // ldr    x0, [x8, #160]
    WORD $0xb940d90d // ldr    w13, [x8, #216]
    WORD $0x9100400e // add    x14, x0, #16
    WORD $0x110005ad // add    w13, w13, #1
    WORD $0xf900510e // str    x14, [x8, #160]
    WORD $0xb900d90d // str    w13, [x8, #216]
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a9201a2 // csel    w2, w13, w18, eq
    CMP $0, R18_PLATFORM
    BEQ LBB5_1630
    B LBB5_341
LBB5_1228:
    WORD $0xf10002df // cmp    x22, #0
    WORD $0x5280002e // mov    w14, #1
    WORD $0x9a8e05ce // cinc    x14, x14, ne
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x9ace26ae // lsr    x14, x21, x14
    WORD $0xb34c2dae // bfi    x14, x13, #52, #12
    WORD $0xb24101cd // orr    x13, x14, #0x8000000000000000
    WORD $0x9a8e11ad // csel    x13, x13, x14, ne
    WORD $0x9e6701a0 // fmov    d0, x13
    CMP $0, R4
    BEQ LBB5_1597
    WORD $0x9100046d // add    x13, x3, #1
    WORD $0xdac011ae // clz    x14, x13
    WORD $0xcb0e0263 // sub    x3, x19, x14
    WORD $0x9ace21b4 // lsl    x20, x13, x14
    WORD $0xaa3403ed // mvn    x13, x20
    WORD $0x9bc77e84 // umulh    x4, x20, x7
    WORD $0x9b077e87 // mul    x7, x20, x7
    WORD $0x92402093 // and    x19, x4, #0x1ff
    WORD $0xeb0d00ff // cmp    x7, x13
    BLS LBB5_1234
    WORD $0xf107fe7f // cmp    x19, #511
    BNE LBB5_1234
    WORD $0xd37cec0e // lsl    x14, x0, #4
    WORD $0xf86e684e // ldr    x14, [x2, x14]
    WORD $0x9bd47dc0 // umulh    x0, x14, x20
    WORD $0x9b147dce // mul    x14, x14, x20
    WORD $0xab070007 // adds    x7, x0, x7
    WORD $0x9a843484 // cinc    x4, x4, hs
    WORD $0xeb0d01df // cmp    x14, x13
    WORD $0x92402093 // and    x19, x4, #0x1ff
    BLS LBB5_1234
    WORD $0xb10004ff // cmn    x7, #1
    BNE LBB5_1234
    WORD $0xf107fe7f // cmp    x19, #511
    BEQ LBB5_1191
LBB5_1234:
    WORD $0xd37ffc8d // lsr    x13, x4, #63
    WORD $0xaa1300e2 // orr    x2, x7, x19
    WORD $0x910025ae // add    x14, x13, #9
    WORD $0x9ace2480 // lsr    x0, x4, x14
    CMP $0, R2
    BNE LBB5_1236
    WORD $0x9240040e // and    x14, x0, #0x3
    WORD $0xf10005df // cmp    x14, #1
    BEQ LBB5_1191
LBB5_1236:
    WORD $0x9240000e // and    x14, x0, #0x1
    WORD $0x8b0d006d // add    x13, x3, x13
    WORD $0x8b0001c0 // add    x0, x14, x0
    WORD $0xd376fc02 // lsr    x2, x0, #54
    WORD $0xf100005f // cmp    x2, #0
    WORD $0x1a9f17ee // cset    w14, eq
    WORD $0xcb0e01ad // sub    x13, x13, x14
    WORD $0xd11ffdae // sub    x14, x13, #2047
    WORD $0xb11ff9df // cmn    x14, #2046
    BLO LBB5_1191
    WORD $0xf100005f // cmp    x2, #0
    WORD $0x5280002e // mov    w14, #1
    WORD $0x9a8e05ce // cinc    x14, x14, ne
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x9ace240e // lsr    x14, x0, x14
    WORD $0xb34c2dae // bfi    x14, x13, #52, #12
    WORD $0xb24101cd // orr    x13, x14, #0x8000000000000000
    WORD $0x9a8e11ad // csel    x13, x13, x14, ne
    WORD $0x9e6701a1 // fmov    d1, x13
    WORD $0x1e612000 // fcmp    d0, d1
    BEQ LBB5_1597
    B LBB5_1191
LBB5_1238:
    WORD $0xaa0f03e7 // mov    x7, x15
    WORD $0x92800004 // mov    x4, #-1
    CMP $0, R18_PLATFORM
    BNE LBB5_1039
    B LBB5_1226
LBB5_1239:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa0303fd // mov    x29, x3
LBB5_1240:
    WORD $0x5280006c // mov    w12, #3
    B LBB5_1602
LBB5_1241:
    WORD $0x7101f5bf // cmp    w13, #125
    BEQ LBB5_986
LBB5_1242:
    WORD $0x2a0d03e2 // mov    w2, w13
    B LBB5_837
LBB5_1243:
    WORD $0x710175bf // cmp    w13, #93
    BEQ LBB5_1068
LBB5_1244:
    WORD $0x2a0d03e2 // mov    w2, w13
    B LBB5_1703
LBB5_1245:
    WORD $0xd100046e // sub    x14, x3, #1
    WORD $0xeb0e025f // cmp    x18, x14
    WORD $0x1a9f17ee // cset    w14, eq
    WORD $0x6a0e01ff // tst    w15, w14
    WORD $0xda8301a4 // csinv    x4, x13, x3, eq
    B LBB5_1162
LBB5_1246:
    WORD $0x5280002d // mov    w13, #1
    B LBB5_1227
LBB5_1247:
    WORD $0x11000a6d // add    w13, w19, #2
    WORD $0x12800005 // mov    w5, #-1
LBB5_1248:
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x2a1f03f3 // mov    w19, wzr
    WORD $0xeb0d00e7 // subs    x7, x7, x13
    BLE LBB5_1254
    WORD $0x8b0601ad // add    x13, x13, x6
    WORD $0x2a1f03f3 // mov    w19, wzr
    WORD $0x8b0d01ed // add    x13, x15, x13
    WORD $0x5284e1ef // mov    w15, #9999
    WORD $0x52800141 // mov    w1, #10
LBB5_1250:
    WORD $0x384015a6 // ldrb    w6, [x13], #1
    WORD $0x7100c0df // cmp    w6, #48
    BLO LBB5_1254
    WORD $0x7100e4df // cmp    w6, #57
    BHI LBB5_1254
    WORD $0x6b0f027f // cmp    w19, w15
    BGT LBB5_1254
    WORD $0x1b011a6e // madd    w14, w19, w1, w6
    WORD $0xf10004e7 // subs    x7, x7, #1
    WORD $0x5100c1d3 // sub    w19, w14, #48
    BNE LBB5_1250
LBB5_1254:
    WORD $0x1b051264 // madd    w4, w19, w5, w4
LBB5_1255:
    WORD $0xaa1f03f5 // mov    x21, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    CMP $0, R24
    BEQ LBB5_1595
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0xd2effe15 // mov    x21, #9218868437227405312
    WORD $0x7104d89f // cmp    w4, #310
    BGT LBB5_1595
    WORD $0xaa1f03f5 // mov    x21, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x3105289f // cmn    w4, #330
    BLT LBB5_1595
    WORD $0x7100049f // cmp    w4, #1
    BLT LBB5_1319
    WORD $0xb201e7e7 // mov    x7, #-7378697629483820647
    WORD $0x2a1f03e5 // mov    w5, wzr
    WORD $0xf2933347 // movk    x7, #39322
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x52800146 // mov    w6, #10
    WORD $0xf2e03327 // movk    x7, #409, lsl #48
    WORD $0x2a1803fb // mov    w27, w24
    WORD $0x2a1803f5 // mov    w21, w24
    ADR POW_TAB, R19
    WORD $0x91000273 // add    x19, x19, :lo12:POW_TAB
    B LBB5_1263
LBB5_1260:
    CMP $0, R27
    BEQ LBB5_1605
LBB5_1261:
    WORD $0x2a1b03f8 // mov    w24, w27
    WORD $0x2a1b03f5 // mov    w21, w27
LBB5_1262:
    WORD $0x0b050285 // add    w5, w20, w5
    WORD $0x7100009f // cmp    w4, #0
    BLE LBB5_1321
LBB5_1263:
    WORD $0x7100209f // cmp    w4, #8
    BLS LBB5_1266
    WORD $0x52800374 // mov    w20, #27
    CMP $0, R21
    BEQ LBB5_1262
    WORD $0x12800359 // mov    w25, #-27
    B LBB5_1292
LBB5_1266:
    WORD $0xb8645a74 // ldr    w20, [x19, w4, uxtw #2]
    CMP $0, R21
    BEQ LBB5_1262
    WORD $0x4b1403f9 // neg    w25, w20
    WORD $0x3100f73f // cmn    w25, #61
    BLS LBB5_1271
    B LBB5_1292
LBB5_1268:
    WORD $0x2a1f03e4 // mov    w4, wzr
LBB5_1269:
    WORD $0x2a1f03f8 // mov    w24, wzr
LBB5_1270:
    WORD $0x1100f2d9 // add    w25, w22, #60
    WORD $0x2a1803f5 // mov    w21, w24
    WORD $0x3101e2df // cmn    w22, #120
    BGE LBB5_1291
LBB5_1271:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03fa // mov    x26, xzr
    WORD $0x2a1903f6 // mov    w22, w25
    WORD $0x0ab57eb7 // bic    w23, w21, w21, asr #31
LBB5_1272:
    WORD $0xeb0d02ff // cmp    x23, x13
    BEQ LBB5_1275
    WORD $0x386d680e // ldrb    w14, [x0, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b063b4e // madd    x14, x26, x6, x14
    WORD $0xd100c1da // sub    x26, x14, #48
    WORD $0xd37cff4e // lsr    x14, x26, #60
    CMP $0, R14
    BEQ LBB5_1272
    WORD $0xaa1a03f9 // mov    x25, x26
    WORD $0x2a0d03f7 // mov    w23, w13
    B LBB5_1277
LBB5_1275:
    CMP $0, R26
    BEQ LBB5_1269
LBB5_1276:
    WORD $0x8b1a0b4d // add    x13, x26, x26, lsl #2
    WORD $0x110006f7 // add    w23, w23, #1
    WORD $0xd37ff9b9 // lsl    x25, x13, #1
    WORD $0xeb07035f // cmp    x26, x7
    WORD $0xaa1903fa // mov    x26, x25
    BLO LBB5_1276
LBB5_1277:
    WORD $0x6b1502ff // cmp    w23, w21
    BGE LBB5_1281
    WORD $0x2a1703ed // mov    w13, w23
    WORD $0xaa1f03e1 // mov    x1, xzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x93407f15 // sxtw    x21, w24
    WORD $0x8b0d0018 // add    x24, x0, x13
LBB5_1279:
    WORD $0xd37cff2e // lsr    x14, x25, #60
    WORD $0x9240ef39 // and    x25, x25, #0xfffffffffffffff
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x3821680e // strb    w14, [x0, x1]
    WORD $0x38616b0e // ldrb    w14, [x24, x1]
    WORD $0x91000421 // add    x1, x1, #1
    WORD $0x9b063b2e // madd    x14, x25, x6, x14
    WORD $0xd100c1d9 // sub    x25, x14, #48
    WORD $0x8b0101ae // add    x14, x13, x1
    WORD $0xeb1501df // cmp    x14, x21
    BLT LBB5_1279
    WORD $0x2a0103f8 // mov    w24, w1
    CMP $0, R25
    BNE LBB5_1283
    B LBB5_1285
LBB5_1281:
    WORD $0x2a1f03f8 // mov    w24, wzr
    B LBB5_1283
LBB5_1282:
    WORD $0xd37cff2d // lsr    x13, x25, #60
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x9240ef2d // and    x13, x25, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d9 // lsl    x25, x14, #1
    CMP $0, R13
    BEQ LBB5_1285
LBB5_1283:
    WORD $0x93407f0d // sxtw    x13, w24
    WORD $0xeb0d007f // cmp    x3, x13
    BLS LBB5_1282
    WORD $0xd37cff2e // lsr    x14, x25, #60
    WORD $0x11000718 // add    w24, w24, #1
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x382d680e // strb    w14, [x0, x13]
    WORD $0x9240ef2d // and    x13, x25, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d9 // lsl    x25, x14, #1
    CMP $0, R13
    BNE LBB5_1283
LBB5_1285:
    WORD $0x4b17008d // sub    w13, w4, w23
    WORD $0x7100071f // cmp    w24, #1
    WORD $0x110005a4 // add    w4, w13, #1
    BLT LBB5_1290
    WORD $0x2a1803ed // mov    w13, w24
    WORD $0x8b0001ae // add    x14, x13, x0
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_1270
LBB5_1287:
    WORD $0xf10005b8 // subs    x24, x13, #1
    BLS LBB5_1268
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d480e // ldrb    w14, [x0, w13, uxtw]
    WORD $0xaa1803ed // mov    x13, x24
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1287
    B LBB5_1270
LBB5_1290:
    CMP $0, R24
    BNE LBB5_1270
    B LBB5_1268
LBB5_1291:
    WORD $0x2a1803f5 // mov    w21, w24
LBB5_1292:
    WORD $0xaa1f03fb // mov    x27, xzr
    WORD $0xaa1f03f7 // mov    x23, xzr
    WORD $0x4b1903f6 // neg    w22, w25
    WORD $0x0ab57eba // bic    w26, w21, w21, asr #31
LBB5_1293:
    WORD $0xeb1b035f // cmp    x26, x27
    BEQ LBB5_1300
    WORD $0x387b680d // ldrb    w13, [x0, x27]
    WORD $0x9100077b // add    x27, x27, #1
    WORD $0x9b0636ed // madd    x13, x23, x6, x13
    WORD $0xd100c1b7 // sub    x23, x13, #48
    WORD $0x9ad626ed // lsr    x13, x23, x22
    CMP $0, R13
    BEQ LBB5_1293
    WORD $0x2a1b03fa // mov    w26, w27
LBB5_1296:
    WORD $0x9ad621ed // lsl    x13, x15, x22
    WORD $0x6b15035f // cmp    w26, w21
    WORD $0xaa2d03f9 // mvn    x25, x13
    BGE LBB5_1304
    WORD $0x2a1a03ed // mov    w13, w26
    WORD $0xaa1f03fb // mov    x27, xzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x93407f01 // sxtw    x1, w24
    WORD $0x8b0d0015 // add    x21, x0, x13
LBB5_1298:
    WORD $0x9ad626ee // lsr    x14, x23, x22
    WORD $0x8a1902f7 // and    x23, x23, x25
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x383b680e // strb    w14, [x0, x27]
    WORD $0x387b6aae // ldrb    w14, [x21, x27]
    WORD $0x9100077b // add    x27, x27, #1
    WORD $0x9b063aee // madd    x14, x23, x6, x14
    WORD $0xd100c1d7 // sub    x23, x14, #48
    WORD $0x8b1b01ae // add    x14, x13, x27
    WORD $0xeb0101df // cmp    x14, x1
    BLT LBB5_1298
    B LBB5_1305
LBB5_1300:
    CMP $0, R23
    BEQ LBB5_1314
    WORD $0x9ad626ed // lsr    x13, x23, x22
    CMP $0, R13
    BEQ LBB5_1303
    WORD $0x9ad621ed // lsl    x13, x15, x22
    WORD $0x4b1a008e // sub    w14, w4, w26
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0xaa2d03f9 // mvn    x25, x13
    WORD $0x110005c4 // add    w4, w14, #1
    B LBB5_1307
LBB5_1303:
    WORD $0x8b170aed // add    x13, x23, x23, lsl #2
    WORD $0x1100075a // add    w26, w26, #1
    WORD $0xd37ff9b7 // lsl    x23, x13, #1
    WORD $0x9ad626ed // lsr    x13, x23, x22
    CMP $0, R13
    BEQ LBB5_1303
    B LBB5_1296
LBB5_1304:
    WORD $0x2a1f03fb // mov    w27, wzr
LBB5_1305:
    WORD $0x4b1a008d // sub    w13, w4, w26
    WORD $0x110005a4 // add    w4, w13, #1
    CMP $0, R23
    BNE LBB5_1307
    B LBB5_1309
LBB5_1306:
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x8a1902ed // and    x13, x23, x25
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d7 // lsl    x23, x14, #1
    CMP $0, R13
    BEQ LBB5_1309
LBB5_1307:
    WORD $0x9ad626ed // lsr    x13, x23, x22
    WORD $0x93407f61 // sxtw    x1, w27
    WORD $0xeb01007f // cmp    x3, x1
    BLS LBB5_1306
    WORD $0x1100c1ad // add    w13, w13, #48
    WORD $0x1100077b // add    w27, w27, #1
    WORD $0x3821680d // strb    w13, [x0, x1]
    WORD $0x8a1902ed // and    x13, x23, x25
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d7 // lsl    x23, x14, #1
    CMP $0, R13
    BNE LBB5_1307
LBB5_1309:
    WORD $0x7100077f // cmp    w27, #1
    BLT LBB5_1260
    WORD $0x2a1b03ed // mov    w13, w27
    WORD $0x8b0001ae // add    x14, x13, x0
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_1261
LBB5_1311:
    WORD $0xf10005b5 // subs    x21, x13, #1
    BLS LBB5_1316
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d480e // ldrb    w14, [x0, w13, uxtw]
    WORD $0xaa1503ed // mov    x13, x21
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1311
    WORD $0x2a1503f8 // mov    w24, w21
    WORD $0x2a1503fb // mov    w27, w21
    B LBB5_1262
LBB5_1314:
    WORD $0x2a1f03f8 // mov    w24, wzr
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0x2a1f03f5 // mov    w21, wzr
    B LBB5_1262
LBB5_1315:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xcb0303fd // neg    x29, x3
    WORD $0x5280016c // mov    w12, #11
    B LBB5_1602
LBB5_1316:
    WORD $0x510005b8 // sub    w24, w13, #1
LBB5_1317:
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x0b050285 // add    w5, w20, w5
    B LBB5_1320
LBB5_1318:
    WORD $0x5ac001ad // rbit    w13, w13
    WORD $0xaa2403ee // mvn    x14, x4
    WORD $0x5ac011ad // clz    w13, w13
    WORD $0xcb0d01c4 // sub    x4, x14, x13
    B LBB5_1162
LBB5_1319:
    WORD $0x2a1f03e5 // mov    w5, wzr
LBB5_1320:
    WORD $0x2a1803fb // mov    w27, w24
LBB5_1321:
    WORD $0xb201e7f5 // mov    x21, #-7378697629483820647
    WORD $0xb202e7e7 // mov    x7, #-3689348814741910324
    WORD $0xf2933355 // movk    x21, #39322
    WORD $0xf90007fd // str    x29, [sp, #8]
    WORD $0xf29999a7 // movk    x7, #52429
    WORD $0x92800133 // mov    x19, #-10
    WORD $0xb2607ff4 // mov    x20, #-4294967296
    WORD $0xf2e03335 // movk    x21, #409, lsl #48
    WORD $0x52800156 // mov    w22, #10
    WORD $0x2a1803fa // mov    w26, w24
    WORD $0x2a1b03fd // mov    w29, w27
    ADR POW_TAB, R24
    WORD $0x91000318 // add    x24, x24, :lo12:POW_TAB
    B LBB5_1327
LBB5_1322:
    WORD $0x5100075a // sub    w26, w26, #1
LBB5_1323:
    WORD $0x2a1f03e4 // mov    w4, wzr
LBB5_1324:
    TST $(1<<31), R25
    BNE LBB5_1355
LBB5_1325:
    WORD $0x2a1a03fd // mov    w29, w26
    WORD $0x2a1a03fb // mov    w27, w26
LBB5_1326:
    WORD $0x4b1900a5 // sub    w5, w5, w25
LBB5_1327:
    TST $(1<<31), R4
    BNE LBB5_1330
    CMP $0, R4
    BNE LBB5_1414
    WORD $0x3940000d // ldrb    w13, [x0]
    WORD $0x7100d5bf // cmp    w13, #53
    BLO LBB5_1333
    B LBB5_1414
LBB5_1330:
    WORD $0x3100209f // cmn    w4, #8
    BHS LBB5_1333
    WORD $0x52800379 // mov    w25, #27
    CMP $0, R29
    BEQ LBB5_1397
    WORD $0x2a1d03fb // mov    w27, w29
    B LBB5_1334
LBB5_1333:
    WORD $0x4b0403ed // neg    w13, w4
    WORD $0xb86d5b19 // ldr    w25, [x24, w13, uxtw #2]
    CMP $0, R27
    BEQ LBB5_1326
LBB5_1334:
    WORD $0x52800d0e // mov    w14, #104
    ADR LSHIFT_TAB, R13
    WORD $0x910001ad // add    x13, x13, :lo12:LSHIFT_TAB
    WORD $0x93407f7d // sxtw    x29, w27
    WORD $0x9bae372d // umaddl    x13, w25, w14, x13
    WORD $0x2a1903f7 // mov    w23, w25
    WORD $0xaa0003e1 // mov    x1, x0
    WORD $0xaa1d03e6 // mov    x6, x29
    WORD $0xb84045bb // ldr    w27, [x13], #4
    WORD $0xaa0d03ef // mov    x15, x13
LBB5_1335:
    WORD $0x384015ee // ldrb    w14, [x15], #1
    CMP $0, R14
    BEQ LBB5_1340
    WORD $0x3940003e // ldrb    w30, [x1]
    WORD $0x6b0e03df // cmp    w30, w14
    BNE LBB5_1383
    WORD $0xf10004c6 // subs    x6, x6, #1
    WORD $0x91000421 // add    x1, x1, #1
    BNE LBB5_1335
    WORD $0x387d69ad // ldrb    w13, [x13, x29]
    CMP $0, R13
    BEQ LBB5_1340
LBB5_1339:
    WORD $0x5100077b // sub    w27, w27, #1
LBB5_1340:
    WORD $0x710007bf // cmp    w29, #1
    BLT LBB5_1350
    WORD $0x0b1d036d // add    w13, w27, w29
    WORD $0x92407fae // and    x14, x29, #0xffffffff
    WORD $0x93407da1 // sxtw    x1, w13
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x93607dbd // sbfiz    x29, x13, #32, #32
    WORD $0xd100042d // sub    x13, x1, #1
    WORD $0x910005de // add    x30, x14, #1
    B LBB5_1343
LBB5_1342:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x8b1403bd // add    x29, x29, x20
    WORD $0xd100042d // sub    x13, x1, #1
    WORD $0xd10007de // sub    x30, x30, #1
    WORD $0xf10007df // cmp    x30, #1
    BLS LBB5_1345
LBB5_1343:
    WORD $0x51000bce // sub    w14, w30, #2
    WORD $0xaa0d03e1 // mov    x1, x13
    WORD $0xeb0301bf // cmp    x13, x3
    WORD $0x386e480e // ldrb    w14, [x0, w14, uxtw]
    WORD $0xd100c1ce // sub    x14, x14, #48
    WORD $0x9ad721ce // lsl    x14, x14, x23
    WORD $0x8b0f01c6 // add    x6, x14, x15
    WORD $0x9bc77cce // umulh    x14, x6, x7
    WORD $0xd343fdcf // lsr    x15, x14, #3
    WORD $0x9b1319ee // madd    x14, x15, x19, x6
    BHS LBB5_1342
    WORD $0x1100c1cd // add    w13, w14, #48
    WORD $0x3821680d // strb    w13, [x0, x1]
    WORD $0x8b1403bd // add    x29, x29, x20
    WORD $0xd100042d // sub    x13, x1, #1
    WORD $0xd10007de // sub    x30, x30, #1
    WORD $0xf10007df // cmp    x30, #1
    BHI LBB5_1343
LBB5_1345:
    WORD $0xf10028df // cmp    x6, #10
    BLO LBB5_1350
    WORD $0x93407c2d // sxtw    x13, w1
    WORD $0xd10005b7 // sub    x23, x13, #1
    B LBB5_1348
LBB5_1347:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0xd10006f7 // sub    x23, x23, #1
    WORD $0xf10025ff // cmp    x15, #9
    WORD $0xaa0d03ef // mov    x15, x13
    BLS LBB5_1350
LBB5_1348:
    WORD $0x9bc77ded // umulh    x13, x15, x7
    WORD $0xeb0302ff // cmp    x23, x3
    WORD $0xd343fdad // lsr    x13, x13, #3
    WORD $0x9b133dae // madd    x14, x13, x19, x15
    BHS LBB5_1347
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x3837680e // strb    w14, [x0, x23]
    WORD $0xd10006f7 // sub    x23, x23, #1
    WORD $0xf10025ff // cmp    x15, #9
    WORD $0xaa0d03ef // mov    x15, x13
    BHI LBB5_1348
LBB5_1350:
    WORD $0x0b1a036d // add    w13, w27, w26
    WORD $0x0b040364 // add    w4, w27, w4
    WORD $0xeb2dc07f // cmp    x3, w13, sxtw
    WORD $0x1a8381ba // csel    w26, w13, w3, hi
    WORD $0x7100075f // cmp    w26, #1
    BLT LBB5_1382
    WORD $0x8b00034d // add    x13, x26, x0
    WORD $0x385ff1ad // ldurb    w13, [x13, #-1]
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_1324
LBB5_1352:
    WORD $0xf100074d // subs    x13, x26, #1
    BLS LBB5_1322
    WORD $0x51000b4e // sub    w14, w26, #2
    WORD $0xaa0d03fa // mov    x26, x13
    WORD $0x386e480e // ldrb    w14, [x0, w14, uxtw]
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1352
    WORD $0x2a0d03fa // mov    w26, w13
    B LBB5_1324
LBB5_1355:
    WORD $0x3100f73f // cmn    w25, #61
    BHI LBB5_1384
    WORD $0x2a1903fb // mov    w27, w25
    B LBB5_1360
LBB5_1357:
    WORD $0x2a1f03e4 // mov    w4, wzr
LBB5_1358:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_1359:
    WORD $0x1100f36f // add    w15, w27, #60
    WORD $0x3101e37f // cmn    w27, #120
    WORD $0x2a0f03fb // mov    w27, w15
    BGE LBB5_1385
LBB5_1360:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03f7 // mov    x23, xzr
    WORD $0x0aba7f5d // bic    w29, w26, w26, asr #31
LBB5_1361:
    WORD $0xeb0d03bf // cmp    x29, x13
    BEQ LBB5_1364
    WORD $0x386d680e // ldrb    w14, [x0, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b163aee // madd    x14, x23, x22, x14
    WORD $0xd100c1d7 // sub    x23, x14, #48
    WORD $0xd37cfeee // lsr    x14, x23, #60
    CMP $0, R14
    BEQ LBB5_1361
    WORD $0xaa1703ef // mov    x15, x23
    WORD $0x2a0d03fd // mov    w29, w13
    B LBB5_1366
LBB5_1364:
    CMP $0, R23
    BEQ LBB5_1358
LBB5_1365:
    WORD $0x8b170aed // add    x13, x23, x23, lsl #2
    WORD $0x110007bd // add    w29, w29, #1
    WORD $0xd37ff9af // lsl    x15, x13, #1
    WORD $0xeb1502ff // cmp    x23, x21
    WORD $0xaa0f03f7 // mov    x23, x15
    BLO LBB5_1365
LBB5_1366:
    WORD $0x6b1a03bf // cmp    w29, w26
    BGE LBB5_1371
    WORD $0x2a1d03ed // mov    w13, w29
    WORD $0x93407f41 // sxtw    x1, w26
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0xaa0003e6 // mov    x6, x0
    WORD $0xcb0d003a // sub    x26, x1, x13
LBB5_1368:
    WORD $0xd37cfdee // lsr    x14, x15, #60
    WORD $0x9240edef // and    x15, x15, #0xfffffffffffffff
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xeb0101bf // cmp    x13, x1
    WORD $0x390000ce // strb    w14, [x6]
    WORD $0x386d68ce // ldrb    w14, [x6, x13]
    WORD $0x910004c6 // add    x6, x6, #1
    WORD $0x9b1639ee // madd    x14, x15, x22, x14
    WORD $0xd100c1cf // sub    x15, x14, #48
    BNE LBB5_1368
    CMP $0, R15
    BNE LBB5_1372
    B LBB5_1376
LBB5_1371:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_1372:
    B LBB5_1374
LBB5_1373:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x9240eded // and    x13, x15, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9cf // lsl    x15, x14, #1
    CMP $0, R13
    BEQ LBB5_1376
LBB5_1374:
    WORD $0x93407f4d // sxtw    x13, w26
    WORD $0xd37cfdee // lsr    x14, x15, #60
    WORD $0xeb0d007f // cmp    x3, x13
    BLS LBB5_1373
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x1100075a // add    w26, w26, #1
    WORD $0x382d680e // strb    w14, [x0, x13]
    WORD $0x9240eded // and    x13, x15, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9cf // lsl    x15, x14, #1
    CMP $0, R13
    BNE LBB5_1374
LBB5_1376:
    WORD $0x4b1d008d // sub    w13, w4, w29
    WORD $0x7100075f // cmp    w26, #1
    WORD $0x110005a4 // add    w4, w13, #1
    BLT LBB5_1381
    WORD $0x2a1a03ed // mov    w13, w26
    WORD $0x8b0001ae // add    x14, x13, x0
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_1359
LBB5_1378:
    WORD $0xf10005ba // subs    x26, x13, #1
    BLS LBB5_1357
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d480e // ldrb    w14, [x0, w13, uxtw]
    WORD $0xaa1a03ed // mov    x13, x26
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1378
    B LBB5_1359
LBB5_1381:
    CMP $0, R26
    BNE LBB5_1359
    B LBB5_1357
LBB5_1382:
    CMP $0, R26
    BNE LBB5_1324
    B LBB5_1323
LBB5_1383:
    BLO LBB5_1339
    B LBB5_1340
LBB5_1384:
    WORD $0x2a1903ef // mov    w15, w25
LBB5_1385:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03fe // mov    x30, xzr
    WORD $0x4b0f03fb // neg    w27, w15
    WORD $0x0aba7f4f // bic    w15, w26, w26, asr #31
LBB5_1386:
    WORD $0xeb0d01ff // cmp    x15, x13
    BEQ LBB5_1393
    WORD $0x386d680e // ldrb    w14, [x0, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b163bce // madd    x14, x30, x22, x14
    WORD $0xd100c1de // sub    x30, x14, #48
    WORD $0x9adb27ce // lsr    x14, x30, x27
    CMP $0, R14
    BEQ LBB5_1386
    WORD $0x2a0d03ef // mov    w15, w13
LBB5_1389:
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x6b1a01ff // cmp    w15, w26
    WORD $0x9adb21ad // lsl    x13, x13, x27
    WORD $0xaa2d03f7 // mvn    x23, x13
    BGE LBB5_1398
    WORD $0x2a0f03ed // mov    w13, w15
    WORD $0x93407f41 // sxtw    x1, w26
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0xaa0003e6 // mov    x6, x0
    WORD $0xcb0d003d // sub    x29, x1, x13
LBB5_1391:
    WORD $0x9adb27ce // lsr    x14, x30, x27
    WORD $0x8a1703da // and    x26, x30, x23
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xeb0101bf // cmp    x13, x1
    WORD $0x390000ce // strb    w14, [x6]
    WORD $0x386d68ce // ldrb    w14, [x6, x13]
    WORD $0x910004c6 // add    x6, x6, #1
    WORD $0x9b163b4e // madd    x14, x26, x22, x14
    WORD $0xd100c1de // sub    x30, x14, #48
    BNE LBB5_1391
    B LBB5_1399
LBB5_1393:
    CMP $0, R30
    BEQ LBB5_1410
    WORD $0x9adb27cd // lsr    x13, x30, x27
    CMP $0, R13
    BEQ LBB5_1396
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x4b0f008e // sub    w14, w4, w15
    WORD $0x9adb21ad // lsl    x13, x13, x27
    WORD $0x2a1f03fd // mov    w29, wzr
    WORD $0xaa2d03f7 // mvn    x23, x13
    WORD $0x110005c4 // add    w4, w14, #1
    B LBB5_1401
LBB5_1396:
    WORD $0x8b1e0bcd // add    x13, x30, x30, lsl #2
    WORD $0x110005ef // add    w15, w15, #1
    WORD $0xd37ff9be // lsl    x30, x13, #1
    WORD $0x9adb27cd // lsr    x13, x30, x27
    CMP $0, R13
    BEQ LBB5_1396
    B LBB5_1389
LBB5_1397:
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_1327
LBB5_1398:
    WORD $0x2a1f03fd // mov    w29, wzr
LBB5_1399:
    WORD $0x4b0f008d // sub    w13, w4, w15
    WORD $0x110005a4 // add    w4, w13, #1
    CMP $0, R30
    BNE LBB5_1401
    B LBB5_1403
LBB5_1400:
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x8a1703cd // and    x13, x30, x23
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9de // lsl    x30, x14, #1
    CMP $0, R13
    BEQ LBB5_1403
LBB5_1401:
    WORD $0x9adb27cd // lsr    x13, x30, x27
    WORD $0x93407faf // sxtw    x15, w29
    WORD $0xeb0f007f // cmp    x3, x15
    BLS LBB5_1400
    WORD $0x1100c1ad // add    w13, w13, #48
    WORD $0x110007bd // add    w29, w29, #1
    WORD $0x382f680d // strb    w13, [x0, x15]
    WORD $0x8a1703cd // and    x13, x30, x23
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9de // lsl    x30, x14, #1
    CMP $0, R13
    BNE LBB5_1401
LBB5_1403:
    WORD $0x710007bf // cmp    w29, #1
    BLT LBB5_1408
    WORD $0x2a1d03ed // mov    w13, w29
    WORD $0x8b0001ae // add    x14, x13, x0
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_1409
LBB5_1405:
    WORD $0xf10005bb // subs    x27, x13, #1
    BLS LBB5_1411
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d480e // ldrb    w14, [x0, w13, uxtw]
    WORD $0xaa1b03ed // mov    x13, x27
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1405
    WORD $0x2a1b03fa // mov    w26, w27
    WORD $0x2a1b03fd // mov    w29, w27
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_1327
LBB5_1408:
    CMP $0, R29
    BEQ LBB5_1412
LBB5_1409:
    WORD $0x2a1d03fa // mov    w26, w29
    WORD $0x2a1d03fb // mov    w27, w29
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_1327
LBB5_1410:
    WORD $0x2a1f03fa // mov    w26, wzr
    WORD $0x2a1f03fd // mov    w29, wzr
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_1327
LBB5_1411:
    WORD $0x510005ba // sub    w26, w13, #1
    B LBB5_1413
LBB5_1412:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_1413:
    WORD $0x2a1f03e4 // mov    w4, wzr
    B LBB5_1325
LBB5_1414:
    WORD $0x310ff8bf // cmn    w5, #1022
    BGT LBB5_1441
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0xf94007fd // ldr    x29, [sp, #8]
    WORD $0xb9401bf6 // ldr    w22, [sp, #24]
    CMP $0, R27
    BEQ LBB5_1456
    WORD $0x110ff4a6 // add    w6, w5, #1021
    WORD $0x3110e8bf // cmn    w5, #1082
    BHI LBB5_1443
    WORD $0xb201e7e5 // mov    x5, #-7378697629483820647
    WORD $0x52800147 // mov    w7, #10
    WORD $0xf2933345 // movk    x5, #39322
    WORD $0xf2e03325 // movk    x5, #409, lsl #48
    B LBB5_1421
LBB5_1418:
    WORD $0x2a1f03e4 // mov    w4, wzr
LBB5_1419:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_1420:
    WORD $0x1100f0cd // add    w13, w6, #60
    WORD $0x3101e0df // cmn    w6, #120
    WORD $0x2a1a03fb // mov    w27, w26
    WORD $0x2a0d03e6 // mov    w6, w13
    WORD $0x2a1a03ef // mov    w15, w26
    BGE LBB5_1444
LBB5_1421:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x0abb7f73 // bic    w19, w27, w27, asr #31
LBB5_1422:
    WORD $0xeb0d027f // cmp    x19, x13
    BEQ LBB5_1425
    WORD $0x386d680e // ldrb    w14, [x0, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b073a8e // madd    x14, x20, x7, x14
    WORD $0xd100c1d4 // sub    x20, x14, #48
    WORD $0xd37cfe8e // lsr    x14, x20, #60
    CMP $0, R14
    BEQ LBB5_1422
    WORD $0xaa1403ef // mov    x15, x20
    WORD $0x2a0d03f3 // mov    w19, w13
    B LBB5_1427
LBB5_1425:
    CMP $0, R20
    BEQ LBB5_1419
LBB5_1426:
    WORD $0x8b140a8d // add    x13, x20, x20, lsl #2
    WORD $0x11000673 // add    w19, w19, #1
    WORD $0xd37ff9af // lsl    x15, x13, #1
    WORD $0xeb05029f // cmp    x20, x5
    WORD $0xaa0f03f4 // mov    x20, x15
    BLO LBB5_1426
LBB5_1427:
    WORD $0x6b1b027f // cmp    w19, w27
    BGE LBB5_1431
    WORD $0x2a1303ee // mov    w14, w19
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x93407dc1 // sxtw    x1, w14
    WORD $0x93407f54 // sxtw    x20, w26
    WORD $0x8b010015 // add    x21, x0, x1
LBB5_1429:
    WORD $0xd37cfdee // lsr    x14, x15, #60
    WORD $0x9240edef // and    x15, x15, #0xfffffffffffffff
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x382d680e // strb    w14, [x0, x13]
    WORD $0x386d6aae // ldrb    w14, [x21, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b0739ee // madd    x14, x15, x7, x14
    WORD $0xd100c1cf // sub    x15, x14, #48
    WORD $0x8b0d002e // add    x14, x1, x13
    WORD $0xeb1401df // cmp    x14, x20
    BLT LBB5_1429
    WORD $0x2a0d03fa // mov    w26, w13
    CMP $0, R15
    BNE LBB5_1433
    B LBB5_1435
LBB5_1431:
    WORD $0x2a1f03fa // mov    w26, wzr
    B LBB5_1433
LBB5_1432:
    WORD $0xd37cfded // lsr    x13, x15, #60
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x9240eded // and    x13, x15, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9cf // lsl    x15, x14, #1
    CMP $0, R13
    BEQ LBB5_1435
LBB5_1433:
    WORD $0x93407f4d // sxtw    x13, w26
    WORD $0xeb0d007f // cmp    x3, x13
    BLS LBB5_1432
    WORD $0xd37cfdee // lsr    x14, x15, #60
    WORD $0x1100075a // add    w26, w26, #1
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x382d680e // strb    w14, [x0, x13]
    WORD $0x9240eded // and    x13, x15, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9cf // lsl    x15, x14, #1
    CMP $0, R13
    BNE LBB5_1433
LBB5_1435:
    WORD $0x4b13008d // sub    w13, w4, w19
    WORD $0x7100075f // cmp    w26, #1
    WORD $0x110005a4 // add    w4, w13, #1
    BLT LBB5_1440
    WORD $0x2a1a03ed // mov    w13, w26
    WORD $0x8b0001ae // add    x14, x13, x0
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_1420
LBB5_1437:
    WORD $0xf10005ba // subs    x26, x13, #1
    BLS LBB5_1418
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d480e // ldrb    w14, [x0, w13, uxtw]
    WORD $0xaa1a03ed // mov    x13, x26
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1437
    B LBB5_1420
LBB5_1440:
    CMP $0, R26
    BNE LBB5_1420
    B LBB5_1418
LBB5_1441:
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x510004a6 // sub    w6, w5, #1
    WORD $0xd2effe15 // mov    x21, #9218868437227405312
    WORD $0xf94007fd // ldr    x29, [sp, #8]
    WORD $0xb9401bf6 // ldr    w22, [sp, #24]
    WORD $0x711000bf // cmp    w5, #1024
    BLE LBB5_1472
    B LBB5_1595
LBB5_1442:
    WORD $0xaa2403ee // mvn    x14, x4
    WORD $0xcb2d41c4 // sub    x4, x14, w13, uxtw
    B LBB5_1162
LBB5_1443:
    WORD $0x2a1b03ef // mov    w15, w27
    WORD $0x2a0603ed // mov    w13, w6
LBB5_1444:
    WORD $0xaa1f03e5 // mov    x5, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x4b0d03f3 // neg    w19, w13
    WORD $0x0aaf7df5 // bic    w21, w15, w15, asr #31
    WORD $0x5280014d // mov    w13, #10
LBB5_1445:
    WORD $0xeb0502bf // cmp    x21, x5
    BEQ LBB5_1452
    WORD $0x3865680e // ldrb    w14, [x0, x5]
    WORD $0x910004a5 // add    x5, x5, #1
    WORD $0x9b0d3a8e // madd    x14, x20, x13, x14
    WORD $0xd100c1d4 // sub    x20, x14, #48
    WORD $0x9ad3268e // lsr    x14, x20, x19
    CMP $0, R14
    BEQ LBB5_1445
    WORD $0x2a0503f5 // mov    w21, w5
LBB5_1448:
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x6b0f02bf // cmp    w21, w15
    WORD $0x9ad321ad // lsl    x13, x13, x19
    WORD $0xaa2d03e6 // mvn    x6, x13
    BGE LBB5_1458
    WORD $0x2a1503ed // mov    w13, w21
    WORD $0xaa1f03e5 // mov    x5, xzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x93407f4f // sxtw    x15, w26
    WORD $0x8b0d0001 // add    x1, x0, x13
    WORD $0x52800147 // mov    w7, #10
LBB5_1450:
    WORD $0x9ad3268e // lsr    x14, x20, x19
    WORD $0x8a060294 // and    x20, x20, x6
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x3825680e // strb    w14, [x0, x5]
    WORD $0x3865682e // ldrb    w14, [x1, x5]
    WORD $0x910004a5 // add    x5, x5, #1
    WORD $0x9b073a8e // madd    x14, x20, x7, x14
    WORD $0xd100c1d4 // sub    x20, x14, #48
    WORD $0x8b0501ae // add    x14, x13, x5
    WORD $0xeb0f01df // cmp    x14, x15
    BLT LBB5_1450
    B LBB5_1459
LBB5_1452:
    WORD $0x12807fa6 // mov    w6, #-1022
    CMP $0, R20
    BEQ LBB5_1456
    WORD $0x9ad3268d // lsr    x13, x20, x19
    CMP $0, R13
    BEQ LBB5_1455
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x4b15008e // sub    w14, w4, w21
    WORD $0x9ad321ad // lsl    x13, x13, x19
    WORD $0x2a1f03e5 // mov    w5, wzr
    WORD $0xaa2d03e6 // mvn    x6, x13
    WORD $0x110005c4 // add    w4, w14, #1
    B LBB5_1461
LBB5_1455:
    WORD $0x8b140a8d // add    x13, x20, x20, lsl #2
    WORD $0x110006b5 // add    w21, w21, #1
    WORD $0xd37ff9b4 // lsl    x20, x13, #1
    WORD $0x9ad3268d // lsr    x13, x20, x19
    CMP $0, R13
    BEQ LBB5_1455
    B LBB5_1448
LBB5_1456:
    WORD $0x2a1f03e7 // mov    w7, wzr
    B LBB5_1569
LBB5_1457:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x9280002f // mov    x15, #-2
    B LBB5_1628
LBB5_1458:
    WORD $0x2a1f03e5 // mov    w5, wzr
LBB5_1459:
    WORD $0x4b15008d // sub    w13, w4, w21
    WORD $0x110005a4 // add    w4, w13, #1
    CMP $0, R20
    BNE LBB5_1461
    B LBB5_1463
LBB5_1460:
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x8a06028d // and    x13, x20, x6
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d4 // lsl    x20, x14, #1
    CMP $0, R13
    BEQ LBB5_1463
LBB5_1461:
    WORD $0x9ad3268d // lsr    x13, x20, x19
    WORD $0x93407caf // sxtw    x15, w5
    WORD $0xeb0f007f // cmp    x3, x15
    BLS LBB5_1460
    WORD $0x1100c1ad // add    w13, w13, #48
    WORD $0x110004a5 // add    w5, w5, #1
    WORD $0x382f680d // strb    w13, [x0, x15]
    WORD $0x8a06028d // and    x13, x20, x6
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d4 // lsl    x20, x14, #1
    CMP $0, R13
    BNE LBB5_1461
LBB5_1463:
    WORD $0x710004bf // cmp    w5, #1
    BLT LBB5_1467
    WORD $0x2a0503ef // mov    w15, w5
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0x8b0001ed // add    x13, x15, x0
    WORD $0x385ff1ad // ldurb    w13, [x13, #-1]
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_1468
LBB5_1465:
    WORD $0xaa0f03ed // mov    x13, x15
    WORD $0xf10005ef // subs    x15, x15, #1
    BLS LBB5_1470
    WORD $0x510009ae // sub    w14, w13, #2
    WORD $0x386e480e // ldrb    w14, [x0, w14, uxtw]
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1465
    B LBB5_1471
LBB5_1467:
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0x2a0503ed // mov    w13, w5
    CMP $0, R5
    BNE LBB5_1473
    B LBB5_1594
LBB5_1468:
    WORD $0x2a0503ed // mov    w13, w5
    B LBB5_1473
LBB5_1469:
    WORD $0x52800029 // mov    w9, #1
    WORD $0x52800162 // mov    w2, #11
    WORD $0x39032109 // strb    w9, [x8, #200]
    B LBB5_341
LBB5_1470:
    WORD $0x2a1f03e4 // mov    w4, wzr
LBB5_1471:
    WORD $0x510005ba // sub    w26, w13, #1
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0x2a1a03fb // mov    w27, w26
LBB5_1472:
    WORD $0x2a1f03e7 // mov    w7, wzr
    WORD $0x2a1a03e5 // mov    w5, w26
    WORD $0x2a1b03ed // mov    w13, w27
    CMP $0, R27
    BEQ LBB5_1569
LBB5_1473:
    WORD $0x39400013 // ldrb    w19, [x0]
    WORD $0x93407da7 // sxtw    x7, w13
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_1550
    WORD $0x710004ff // cmp    w7, #1
    BNE LBB5_1476
LBB5_1475:
    WORD $0x5282b18d // mov    w13, #5516
    ADR LSHIFT_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:LSHIFT_TAB
    WORD $0x52800215 // mov    w21, #16
    WORD $0x8b0701ce // add    x14, x14, x7
    WORD $0x386d69cd // ldrb    w13, [x14, x13]
    CMP $0, R13
    BNE LBB5_1551
    B LBB5_1553
LBB5_1476:
    WORD $0x39400413 // ldrb    w19, [x0, #1]
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_1550
    WORD $0x710008ff // cmp    w7, #2
    BEQ LBB5_1475
    WORD $0x39400813 // ldrb    w19, [x0, #2]
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_1550
    WORD $0x71000cff // cmp    w7, #3
    BEQ LBB5_1475
    WORD $0x39400c13 // ldrb    w19, [x0, #3]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_1550
    WORD $0x710010ff // cmp    w7, #4
    BEQ LBB5_1475
    WORD $0x39401013 // ldrb    w19, [x0, #4]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710014ff // cmp    w7, #5
    BEQ LBB5_1475
    WORD $0x39401413 // ldrb    w19, [x0, #5]
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710018ff // cmp    w7, #6
    BEQ LBB5_1475
    WORD $0x39401813 // ldrb    w19, [x0, #6]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_1550
    WORD $0x71001cff // cmp    w7, #7
    BEQ LBB5_1475
    WORD $0x39401c13 // ldrb    w19, [x0, #7]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_1550
    WORD $0x710020ff // cmp    w7, #8
    BEQ LBB5_1475
    WORD $0x39402013 // ldrb    w19, [x0, #8]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710024ff // cmp    w7, #9
    BEQ LBB5_1475
    WORD $0x39402413 // ldrb    w19, [x0, #9]
    WORD $0x52800694 // mov    w20, #52
    WORD $0x7100d27f // cmp    w19, #52
    BNE LBB5_1550
    WORD $0x710028ff // cmp    w7, #10
    BEQ LBB5_1475
    WORD $0x39402813 // ldrb    w19, [x0, #10]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_1550
    WORD $0x71002cff // cmp    w7, #11
    BEQ LBB5_1475
    WORD $0x39402c13 // ldrb    w19, [x0, #11]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710030ff // cmp    w7, #12
    BEQ LBB5_1475
    WORD $0x39403013 // ldrb    w19, [x0, #12]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_1550
    WORD $0x710034ff // cmp    w7, #13
    BEQ LBB5_1475
    WORD $0x39403413 // ldrb    w19, [x0, #13]
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_1550
    WORD $0x710038ff // cmp    w7, #14
    BEQ LBB5_1475
    WORD $0x39403813 // ldrb    w19, [x0, #14]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_1550
    WORD $0x71003cff // cmp    w7, #15
    BEQ LBB5_1475
    WORD $0x39403c13 // ldrb    w19, [x0, #15]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_1550
    WORD $0x710040ff // cmp    w7, #16
    BEQ LBB5_1475
    WORD $0x39404013 // ldrb    w19, [x0, #16]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_1550
    WORD $0x710044ff // cmp    w7, #17
    BEQ LBB5_1475
    WORD $0x39404413 // ldrb    w19, [x0, #17]
    WORD $0x52800694 // mov    w20, #52
    WORD $0x7100d27f // cmp    w19, #52
    BNE LBB5_1550
    WORD $0x710048ff // cmp    w7, #18
    BEQ LBB5_1475
    WORD $0x39404813 // ldrb    w19, [x0, #18]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_1550
    WORD $0x71004cff // cmp    w7, #19
    BEQ LBB5_1475
    WORD $0x39404c13 // ldrb    w19, [x0, #19]
    WORD $0x52800694 // mov    w20, #52
    WORD $0x7100d27f // cmp    w19, #52
    BNE LBB5_1550
    WORD $0x710050ff // cmp    w7, #20
    BEQ LBB5_1475
    WORD $0x39405013 // ldrb    w19, [x0, #20]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710054ff // cmp    w7, #21
    BEQ LBB5_1475
    WORD $0x39405413 // ldrb    w19, [x0, #21]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_1550
    WORD $0x710058ff // cmp    w7, #22
    BEQ LBB5_1475
    WORD $0x39405813 // ldrb    w19, [x0, #22]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_1550
    WORD $0x71005cff // cmp    w7, #23
    BEQ LBB5_1475
    WORD $0x39405c13 // ldrb    w19, [x0, #23]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_1550
    WORD $0x710060ff // cmp    w7, #24
    BEQ LBB5_1475
    WORD $0x39406013 // ldrb    w19, [x0, #24]
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_1550
    WORD $0x710064ff // cmp    w7, #25
    BEQ LBB5_1475
    WORD $0x39406413 // ldrb    w19, [x0, #25]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_1550
    WORD $0x710068ff // cmp    w7, #26
    BEQ LBB5_1475
    WORD $0x39406813 // ldrb    w19, [x0, #26]
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_1550
    WORD $0x71006cff // cmp    w7, #27
    BEQ LBB5_1475
    WORD $0x39406c13 // ldrb    w19, [x0, #27]
    WORD $0x52800714 // mov    w20, #56
    WORD $0x7100e27f // cmp    w19, #56
    BNE LBB5_1550
    WORD $0x710070ff // cmp    w7, #28
    BEQ LBB5_1475
    WORD $0x39407013 // ldrb    w19, [x0, #28]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_1550
    WORD $0x710074ff // cmp    w7, #29
    BEQ LBB5_1475
    WORD $0x39407413 // ldrb    w19, [x0, #29]
    WORD $0x52800734 // mov    w20, #57
    WORD $0x7100e67f // cmp    w19, #57
    BNE LBB5_1550
    WORD $0x710078ff // cmp    w7, #30
    BEQ LBB5_1475
    WORD $0x39407813 // ldrb    w19, [x0, #30]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_1550
    WORD $0x71007cff // cmp    w7, #31
    BEQ LBB5_1475
    WORD $0x39407c13 // ldrb    w19, [x0, #31]
    WORD $0x52800714 // mov    w20, #56
    WORD $0x7100e27f // cmp    w19, #56
    BNE LBB5_1550
    WORD $0x710080ff // cmp    w7, #32
    BEQ LBB5_1475
    WORD $0x39408013 // ldrb    w19, [x0, #32]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710084ff // cmp    w7, #33
    BEQ LBB5_1475
    WORD $0x39408413 // ldrb    w19, [x0, #33]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_1550
    WORD $0x710088ff // cmp    w7, #34
    BEQ LBB5_1475
    WORD $0x39408813 // ldrb    w19, [x0, #34]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_1550
    WORD $0x71008cff // cmp    w7, #35
    BEQ LBB5_1475
    WORD $0x39408c13 // ldrb    w19, [x0, #35]
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_1550
    WORD $0x710090ff // cmp    w7, #36
    BEQ LBB5_1475
    WORD $0x39409013 // ldrb    w19, [x0, #36]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_1550
    WORD $0x710094ff // cmp    w7, #37
    BEQ LBB5_1475
    WORD $0x39409413 // ldrb    w19, [x0, #37]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_1550
    WORD $0x52800215 // mov    w21, #16
    WORD $0x710098ff // cmp    w7, #38
    BEQ LBB5_1475
    B LBB5_1552
LBB5_1550:
    WORD $0x52800215 // mov    w21, #16
    WORD $0x6b14027f // cmp    w19, w20
    BHS LBB5_1552
LBB5_1551:
    WORD $0x528001f5 // mov    w21, #15
LBB5_1552:
    WORD $0x710004ff // cmp    w7, #1
    BLT LBB5_1563
LBB5_1553:
    WORD $0x0b0702b4 // add    w20, w21, w7
    WORD $0x92407ced // and    x13, x7, #0xffffffff
    WORD $0x93407e8e // sxtw    x14, w20
    WORD $0xb202e7f7 // mov    x23, #-3689348814741910324
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x910005a7 // add    x7, x13, #1
    WORD $0xd10005d3 // sub    x19, x14, #1
    WORD $0xd2ff4016 // mov    x22, #-432345564227567616
    WORD $0xf29999b7 // movk    x23, #52429
    WORD $0x92800138 // mov    x24, #-10
    B LBB5_1555
LBB5_1554:
    WORD $0xf100003f // cmp    x1, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0x51000694 // sub    w20, w20, #1
    WORD $0xd1000673 // sub    x19, x19, #1
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf10004ff // cmp    x7, #1
    BLS LBB5_1557
LBB5_1555:
    WORD $0x510008ed // sub    w13, w7, #2
    WORD $0xeb03027f // cmp    x19, x3
    WORD $0x386d480d // ldrb    w13, [x0, w13, uxtw]
    WORD $0x8b0dd5ed // add    x13, x15, x13, lsl #53
    WORD $0x8b1601ad // add    x13, x13, x22
    WORD $0x9bd77dae // umulh    x14, x13, x23
    WORD $0xd343fdcf // lsr    x15, x14, #3
    WORD $0x9b1835e1 // madd    x1, x15, x24, x13
    BHS LBB5_1554
    WORD $0x1100c02e // add    w14, w1, #48
    WORD $0x3833680e // strb    w14, [x0, x19]
    WORD $0x51000694 // sub    w20, w20, #1
    WORD $0xd1000673 // sub    x19, x19, #1
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf10004ff // cmp    x7, #1
    BHI LBB5_1555
LBB5_1557:
    WORD $0xf10029bf // cmp    x13, #10
    BHS LBB5_1559
    WORD $0xb9401bf6 // ldr    w22, [sp, #24]
    B LBB5_1563
LBB5_1559:
    WORD $0x93407e8d // sxtw    x13, w20
    WORD $0x92800133 // mov    x19, #-10
    WORD $0xd10005a7 // sub    x7, x13, #1
    WORD $0xb202e7ed // mov    x13, #-3689348814741910324
    WORD $0xf29999ad // movk    x13, #52429
    WORD $0xb9401bf6 // ldr    w22, [sp, #24]
    B LBB5_1561
LBB5_1560:
    WORD $0xf100029f // cmp    x20, #0
    WORD $0x1a9f0442 // csinc    w2, w2, wzr, eq
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf10025ff // cmp    x15, #9
    WORD $0xaa0103ef // mov    x15, x1
    BLS LBB5_1563
LBB5_1561:
    WORD $0x9bcd7dee // umulh    x14, x15, x13
    WORD $0xeb0300ff // cmp    x7, x3
    WORD $0xd343fdc1 // lsr    x1, x14, #3
    WORD $0x9b133c34 // madd    x20, x1, x19, x15
    BHS LBB5_1560
