    WORD $0x1100c28e // add    w14, w20, #48
    WORD $0x3827680e // strb    w14, [x0, x7]
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf10025ff // cmp    x15, #9
    WORD $0xaa0103ef // mov    x15, x1
    BHI LBB5_1561
LBB5_1563:
    WORD $0x0b0502ad // add    w13, w21, w5
    WORD $0x0b0402a4 // add    w4, w21, w4
    WORD $0xeb2dc07f // cmp    x3, w13, sxtw
    WORD $0x1a8381a7 // csel    w7, w13, w3, hi
    WORD $0x710004ff // cmp    w7, #1
    BLT LBB5_1568
    WORD $0x8b0000ed // add    x13, x7, x0
    WORD $0x385ff1ad // ldurb    w13, [x13, #-1]
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_1569
LBB5_1565:
    WORD $0xf10004ed // subs    x13, x7, #1
    BLS LBB5_1583
    WORD $0x510008ee // sub    w14, w7, #2
    WORD $0xaa0d03e7 // mov    x7, x13
    WORD $0x386e480e // ldrb    w14, [x0, w14, uxtw]
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_1565
    WORD $0x2a0d03e7 // mov    w7, w13
    B LBB5_1569
LBB5_1568:
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x2a1f03e3 // mov    w3, wzr
    CMP $0, R7
    BEQ LBB5_1592
LBB5_1569:
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x7100509f // cmp    w4, #20
    BGT LBB5_1594
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x7100049f // cmp    w4, #1
    BLT LBB5_1575
    WORD $0x2a0403ed // mov    w13, w4
    WORD $0x0aa77ce5 // bic    w5, w7, w7, asr #31
    WORD $0xd10005ae // sub    x14, x13, #1
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0xeb0501df // cmp    x14, x5
    WORD $0x52800153 // mov    w19, #10
    WORD $0x9a8531c3 // csel    x3, x14, x5, lo
    WORD $0xaa0003f4 // mov    x20, x0
    WORD $0x91000461 // add    x1, x3, #1
LBB5_1572:
    CMP $0, R5
    BEQ LBB5_1575
    WORD $0x3840168e // ldrb    w14, [x20], #1
    WORD $0xd10005ad // sub    x13, x13, #1
    WORD $0xd10004a5 // sub    x5, x5, #1
    WORD $0x9b1339ee // madd    x14, x15, x19, x14
    WORD $0xd100c1cf // sub    x15, x14, #48
    CMP $0, R13
    BNE LBB5_1572
    WORD $0xaa0103e3 // mov    x3, x1
LBB5_1575:
    WORD $0x6b030085 // subs    w5, w4, w3
    BLE LBB5_1582
    WORD $0x710010bf // cmp    w5, #4
    BLO LBB5_1580
    WORD $0x5280002d // mov    w13, #1
    WORD $0x5280014e // mov    w14, #10
    WORD $0x121e74b3 // and    w19, w5, #0xfffffffc
    WORD $0x25d8e040 // ptrue    p0.d, vl2
    WORD $0x0b130063 // add    w3, w3, w19
    WORD $0x4e080da0 // dup    v0.2d, x13
    WORD $0x2a1303ed // mov    w13, w19
    WORD $0x4e080dc2 // dup    v2.2d, x14
    WORD $0x4ea01c01 // mov    v1.16b, v0.16b
    WORD $0x4e181de1 // mov    v1.d[1], x15
LBB5_1578:
    WORD $0x710011ad // subs    w13, w13, #4
    WORD $0x04d00040 // mul    z0.d, p0/m, z0.d, z2.d
    WORD $0x04d00041 // mul    z1.d, p0/m, z1.d, z2.d
    BNE LBB5_1578
    WORD $0x4ec07822 // zip2    v2.2d, v1.2d, v0.2d
    WORD $0x6b1300bf // cmp    w5, w19
    WORD $0x4ec03820 // zip1    v0.2d, v1.2d, v0.2d
    WORD $0x04d00040 // mul    z0.d, p0/m, z0.d, z2.d
    WORD $0x25d8e020 // ptrue    p0.d, vl1
    WORD $0x6e004001 // ext    v1.16b, v0.16b, v0.16b, #8
    WORD $0x04d00020 // mul    z0.d, p0/m, z0.d, z1.d
    WORD $0x9e66000f // fmov    x15, d0
    BEQ LBB5_1582
LBB5_1580:
    WORD $0x4b03008d // sub    w13, w4, w3
LBB5_1581:
    WORD $0x8b0f09ee // add    x14, x15, x15, lsl #2
    WORD $0x710005ad // subs    w13, w13, #1
    WORD $0xd37ff9cf // lsl    x15, x14, #1
    BNE LBB5_1581
LBB5_1582:
    WORD $0x2a1f03e3 // mov    w3, wzr
    TST $(1<<31), R4
    BEQ LBB5_1584
    B LBB5_1592
LBB5_1583:
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x510004e7 // sub    w7, w7, #1
LBB5_1584:
    WORD $0x6b0400ff // cmp    w7, w4
    BLE LBB5_1590
    WORD $0x3864480d // ldrb    w13, [x0, w4, uxtw]
    WORD $0x7100d5bf // cmp    w13, #53
    BNE LBB5_1591
    WORD $0x1100048e // add    w14, w4, #1
    WORD $0x6b0701df // cmp    w14, w7
    BNE LBB5_1591
    WORD $0x52800023 // mov    w3, #1
    CMP $0, R2
    BNE LBB5_1592
    CMP $0, R4
    BEQ LBB5_1590
    WORD $0x5100048d // sub    w13, w4, #1
    WORD $0x386d480d // ldrb    w13, [x0, w13, uxtw]
    WORD $0x120001a3 // and    w3, w13, #0x1
    WORD $0x8b2341ef // add    x15, x15, w3, uxtw
    WORD $0xd2e0040d // mov    x13, #9007199254740992
    WORD $0xeb0d01ff // cmp    x15, x13
    BEQ LBB5_1593
    B LBB5_1594
LBB5_1590:
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0x8b3f41ef // add    x15, x15, wzr, uxtw
    WORD $0xd2e0040d // mov    x13, #9007199254740992
    WORD $0xeb0d01ff // cmp    x15, x13
    BEQ LBB5_1593
    B LBB5_1594
LBB5_1591:
    WORD $0x7100d1bf // cmp    w13, #52
    WORD $0x1a9f97e3 // cset    w3, hi
LBB5_1592:
    WORD $0x8b2341ef // add    x15, x15, w3, uxtw
    WORD $0xd2e0040d // mov    x13, #9007199254740992
    WORD $0xeb0d01ff // cmp    x15, x13
    BNE LBB5_1594
LBB5_1593:
    WORD $0x110004cd // add    w13, w6, #1
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x710ff8df // cmp    w6, #1022
    WORD $0xd2e0020f // mov    x15, #4503599627370496
    WORD $0xd2effe15 // mov    x21, #9218868437227405312
    WORD $0x2a0d03e6 // mov    w6, w13
    BGT LBB5_1595
LBB5_1594:
    WORD $0x110ffccd // add    w13, w6, #1023
    WORD $0x9374d1ee // sbfx    x14, x15, #52, #1
    WORD $0x120029ad // and    w13, w13, #0x7ff
    WORD $0xaa0f03f4 // mov    x20, x15
    WORD $0x8a0dd1d5 // and    x21, x14, x13, lsl #52
LBB5_1595:
    WORD $0x9240ce8d // and    x13, x20, #0xfffffffffffff
    WORD $0x7100b6df // cmp    w22, #45
    WORD $0xaa1501ad // orr    x13, x13, x21
    WORD $0x1e620240 // scvtf    d0, w18
    WORD $0xb24101ae // orr    x14, x13, #0x8000000000000000
    WORD $0xb9401ff9 // ldr    w25, [sp, #28]
    WORD $0x9a8d01cd // csel    x13, x14, x13, eq
    WORD $0x9e6701a1 // fmov    d1, x13
LBB5_1596:
    WORD $0x1e610800 // fmul    d0, d0, d1
LBB5_1597:
    WORD $0x9e660001 // fmov    x1, d0
    WORD $0x2a1f03f2 // mov    w18, wzr
    WORD $0x52800082 // mov    w2, #4
    WORD $0xd2effe0e // mov    x14, #9218868437227405312
    WORD $0x9240f82d // and    x13, x1, #0x7fffffffffffffff
    WORD $0xeb0e01bf // cmp    x13, x14
    BNE LBB5_1604
LBB5_1598:
    WORD $0xf1000d9f // cmp    x12, #3
    BEQ LBB5_1036
LBB5_1599:
    WORD $0xf1004d9f // cmp    x12, #19
    BEQ LBB5_1603
    WORD $0xf1002d9f // cmp    x12, #11
    BNE LBB5_2444
    WORD $0xf9405100 // ldr    x0, [x8, #160]
    WORD $0x5280016c // mov    w12, #11
LBB5_1602:
    WORD $0xaa10818d // orr    x13, x12, x16, lsl #32
    WORD $0xf900041d // str    x29, [x0, #8]
    WORD $0xb940d90e // ldr    w14, [x8, #216]
    WORD $0xaa1103ef // mov    x15, x17
    WORD $0x2a0203f2 // mov    w18, w2
    WORD $0xf900000d // str    x13, [x0]
    WORD $0xf9405100 // ldr    x0, [x8, #160]
    WORD $0x110005cd // add    w13, w14, #1
    WORD $0x9100400e // add    x14, x0, #16
    WORD $0xb900d90d // str    w13, [x8, #216]
    WORD $0xf900510e // str    x14, [x8, #160]
    WORD $0x7100005f // cmp    w2, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a8201a2 // csel    w2, w13, w2, eq
    CMP $0, R18_PLATFORM
    BEQ LBB5_1630
    B LBB5_341
LBB5_1603:
    WORD $0x2a0203f2 // mov    w18, w2
    WORD $0xaa1d03e1 // mov    x1, x29
LBB5_1604:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x5280026c // mov    w12, #19
    WORD $0xaa10818f // orr    x15, x12, x16, lsl #32
    WORD $0xaa0103fd // mov    x29, x1
    WORD $0xf90005a1 // str    x1, [x13, #8]
    WORD $0xb940d90e // ldr    w14, [x8, #216]
    WORD $0xf90001af // str    x15, [x13]
    WORD $0xf9405100 // ldr    x0, [x8, #160]
    WORD $0xaa1103ef // mov    x15, x17
    WORD $0x110005ce // add    w14, w14, #1
    WORD $0x9100400d // add    x13, x0, #16
    WORD $0xb900d90e // str    w14, [x8, #216]
    WORD $0xf900510d // str    x13, [x8, #160]
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a9201a2 // csel    w2, w13, w18, eq
    CMP $0, R18_PLATFORM
    BEQ LBB5_1630
    B LBB5_341
LBB5_1605:
    WORD $0x2a1f03f8 // mov    w24, wzr
    B LBB5_1317
LBB5_1606:
    WORD $0xaa0503e7 // mov    x7, x5
    WORD $0x92800003 // mov    x3, #-1
    WORD $0x92800002 // mov    x2, #-1
    B LBB5_1016
LBB5_1607:
    WORD $0xaa0003f1 // mov    x17, x0
    WORD $0xf1000d9f // cmp    x12, #3
    BNE LBB5_1599
    B LBB5_1036
LBB5_1608:
    WORD $0x2a1f03e7 // mov    w7, wzr
    WORD $0x5284e202 // mov    w2, #10000
    B LBB5_961
LBB5_1609:
    WORD $0xcb120231 // sub    x17, x17, x18
    WORD $0x39400233 // ldrb    w19, [x17]
    WORD $0x5100c26d // sub    w13, w19, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1643
    WORD $0xaa1f03e3 // mov    x3, xzr
    WORD $0xaa1f03f2 // mov    x18, xzr
    WORD $0x5280014d // mov    w13, #10
LBB5_1611:
    WORD $0x8b1200ae // add    x14, x5, x18
    WORD $0x9b0d7c71 // mul    x17, x3, x13
    WORD $0x8b334231 // add    x17, x17, w19, uxtw
    WORD $0x394005d3 // ldrb    w19, [x14, #1]
    WORD $0xd100c223 // sub    x3, x17, #48
    WORD $0x5100c262 // sub    w2, w19, #48
    WORD $0x7100245f // cmp    w2, #9
    WORD $0xfa529a42 // ccmp    x18, #18, #2, ls
    WORD $0x91000652 // add    x18, x18, #1
    BLO LBB5_1611
    WORD $0x8b1200b1 // add    x17, x5, x18
    WORD $0x7100245f // cmp    w2, #9
    BHI LBB5_1644
    WORD $0xaa1f03e7 // mov    x7, xzr
LBB5_1614:
    WORD $0x8b0700ad // add    x13, x5, x7
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0x8b1201ad // add    x13, x13, x18
    WORD $0x394005b3 // ldrb    w19, [x13, #1]
    WORD $0x5100c26d // sub    w13, w19, #48
    WORD $0x710029bf // cmp    w13, #10
    BLO LBB5_1614
    WORD $0x8b1200ad // add    x13, x5, x18
    WORD $0x52800024 // mov    w4, #1
    WORD $0x8b0701b1 // add    x17, x13, x7
    B LBB5_966
LBB5_1616:
    WORD $0xaa1d03fb // mov    x27, x29
    WORD $0xaa0f03e1 // mov    x1, x15
    WORD $0x2538cb80 // mov    z0.b, #92
    WORD $0x2538c441 // mov    z1.b, #34
    WORD $0x2538c3e2 // mov    z2.b, #31
LBB5_1617:
    WORD $0xa400a023 // ld1b    { z3.b }, p0/z, [x1]
    WORD $0x2400a061 // cmpeq    p1.b, p0/z, z3.b, z0.b
    WORD $0x2401a063 // cmpeq    p3.b, p0/z, z3.b, z1.b
    WORD $0x2402a062 // cmpeq    p2.b, p0/z, z3.b, z2.b
    WORD $0x25834c64 // mov    p4.b, p3.b
    WORD $0x25c24025 // orrs    p5.b, p0/z, p1.b, p2.b
    BEQ LBB5_1619
    WORD $0x259040a4 // brkb    p4.b, p0/z, p5.b
    WORD $0x25034084 // and    p4.b, p0/z, p4.b, p3.b
LBB5_1619:
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_1626
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_1623
    WORD $0x25904063 // brkb    p3.b, p0/z, p3.b
    WORD $0x25414064 // ands    p4.b, p0/z, p3.b, p1.b
    BNE LBB5_1661
    WORD $0x25024063 // and    p3.b, p0/z, p3.b, p2.b
    B LBB5_1624
LBB5_1623:
    WORD $0x25824843 // mov    p3.b, p2.b
    WORD $0x2550c020 // ptest    p0, p1.b
    BNE LBB5_1661
LBB5_1624:
    WORD $0x2550c060 // ptest    p0, p3.b
    BNE LBB5_1697
    WORD $0x91008021 // add    x1, x1, #32
    B LBB5_1617
LBB5_1626:
    WORD $0xaa1b03fd // mov    x29, x27
LBB5_1627:
    WORD $0x25904061 // brkb    p1.b, p0/z, p3.b
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0x8b0101ad // add    x13, x13, x1
    WORD $0x910005b7 // add    x23, x13, #1
    WORD $0xaa2f03ed // mvn    x13, x15
    WORD $0x8b0d02ef // add    x15, x23, x13
    WORD $0x2a1f03f2 // mov    w18, wzr
    TST $(1<<63), R15
    BEQ LBB5_1629
LBB5_1628:
    WORD $0x4b0f03f2 // neg    w18, w15
LBB5_1629:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x7100033f // cmp    w25, #0
    WORD $0x5280018e // mov    w14, #12
    WORD $0x52800091 // mov    w17, #4
    WORD $0x9a8e022e // csel    x14, x17, x14, eq
    WORD $0xf90005af // str    x15, [x13, #8]
    WORD $0xaa1081ce // orr    x14, x14, x16, lsl #32
    WORD $0xf9405100 // ldr    x0, [x8, #160]
    WORD $0xd2c0002f // mov    x15, #4294967296
    WORD $0xb940d510 // ldr    w16, [x8, #212]
    WORD $0x8b0f01ce // add    x14, x14, x15
    WORD $0xaa1703ef // mov    x15, x23
    WORD $0x91004011 // add    x17, x0, #16
    WORD $0x11000610 // add    w16, w16, #1
    WORD $0xf90001ae // str    x14, [x13]
    WORD $0xf9005111 // str    x17, [x8, #160]
    WORD $0xb900d510 // str    w16, [x8, #212]
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a9201a2 // csel    w2, w13, w18, eq
    CMP $0, R18_PLATFORM
    BNE LBB5_341
LBB5_1630:
    WORD $0xf940610d // ldr    x13, [x8, #192]
    WORD $0x9100800e // add    x14, x0, #32
    WORD $0xeb0d01df // cmp    x14, x13
    BHI LBB5_341
    WORD $0xaa0f03f1 // mov    x17, x15
    WORD $0x38401632 // ldrb    w18, [x17], #1
    WORD $0x7100825f // cmp    w18, #32
    BHI LBB5_1642
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ad221ad // lsl    x13, x13, x18
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1642
    WORD $0x394005f2 // ldrb    w18, [x15, #1]
    WORD $0x910009f1 // add    x17, x15, #2
    WORD $0x7100825f // cmp    w18, #32
    BHI LBB5_1642
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ad221ad // lsl    x13, x13, x18
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1642
    WORD $0xf940490f // ldr    x15, [x8, #144]
    WORD $0xcb0f022d // sub    x13, x17, x15
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_1638
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x92800010 // mov    x16, #-1
    WORD $0x9acd220d // lsl    x13, x16, x13
    WORD $0xea0d01d0 // ands    x16, x14, x13
    BNE LBB5_1641
    WORD $0x910101f1 // add    x17, x15, #64
LBB5_1638:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R16
    WORD $0xd101022f // sub    x15, x17, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00203 // ldr    q3, [x16, :lo12:.LCPI5_2]
LBB5_1639:
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
    WORD $0x1e2600b0 // fmov    w16, s5
    WORD $0x1e2600d1 // fmov    w17, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa1081cd // orr    x13, x14, x16, lsl #32
    WORD $0xaa11c1ad // orr    x13, x13, x17, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_1639
LBB5_1640:
    WORD $0xaa2d03f0 // mvn    x16, x13
    WORD $0xa909410f // stp    x15, x16, [x8, #144]
LBB5_1641:
    WORD $0xdac0020d // rbit    x13, x16
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d01f1 // add    x17, x15, x13
    WORD $0x38401632 // ldrb    w18, [x17], #1
    B LBB5_1938
LBB5_1642:
    B LBB5_1938
LBB5_1643:
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0xaa1f03f2 // mov    x18, xzr
    WORD $0x2a1f03e7 // mov    w7, wzr
    WORD $0xaa1f03e3 // mov    x3, xzr
    B LBB5_966
LBB5_1644:
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x2a1f03e7 // mov    w7, wzr
    B LBB5_966
LBB5_1645:
    WORD $0x710004ff // cmp    w7, #1
    BNE LBB5_1171
    WORD $0x5280014d // mov    w13, #10
    WORD $0x9bcd7c6d // umulh    x13, x3, x13
    WORD $0xeb0d03ff // cmp    xzr, x13
    BEQ LBB5_1692
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x1280000d // mov    w13, #-1
    WORD $0x5a8d15b2 // cneg    w18, w13, eq
    WORD $0x52800027 // mov    w7, #1
    B LBB5_1183
LBB5_1648:
    WORD $0x8b1802ce // add    x14, x22, x24
    WORD $0x8b1802f2 // add    x18, x23, x24
    WORD $0x394001cd // ldrb    w13, [x14]
    WORD $0x710089bf // cmp    w13, #34
    BNE LBB5_1651
LBB5_1649:
    WORD $0x910005cf // add    x15, x14, #1
    WORD $0xcb110251 // sub    x17, x18, x17
LBB5_1650:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x52800192 // mov    w18, #12
    WORD $0xaa0f03e0 // mov    x0, x15
    WORD $0xaa1b03fd // mov    x29, x27
    TST $(1<<63), R17
    BEQ LBB5_921
    B LBB5_1684
LBB5_1651:
    WORD $0x8b1802ee // add    x14, x23, x24
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x390001cd // strb    w13, [x14]
    WORD $0x394005ed // ldrb    w13, [x15, #1]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_1662
    WORD $0x390005cd // strb    w13, [x14, #1]
    WORD $0x394009ef // ldrb    w15, [x15, #2]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_1663
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x390009cf // strb    w15, [x14, #2]
    WORD $0x39400daf // ldrb    w15, [x13, #3]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_1664
    WORD $0x39000dcf // strb    w15, [x14, #3]
    WORD $0x394011af // ldrb    w15, [x13, #4]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_1665
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x390011cf // strb    w15, [x14, #4]
    WORD $0x394015af // ldrb    w15, [x13, #5]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_1666
    WORD $0x390015cf // strb    w15, [x14, #5]
    WORD $0x394019af // ldrb    w15, [x13, #6]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_1667
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x390019cf // strb    w15, [x14, #6]
    WORD $0x39401daf // ldrb    w15, [x13, #7]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_1680
    WORD $0x39001dcf // strb    w15, [x14, #7]
    WORD $0x91002318 // add    x24, x24, #8
    WORD $0x394021ad // ldrb    w13, [x13, #8]
    WORD $0x710089bf // cmp    w13, #34
    BNE LBB5_1651
    WORD $0x8b1802ce // add    x14, x22, x24
    WORD $0x8b1802f2 // add    x18, x23, x24
    B LBB5_1649
LBB5_1660:
    WORD $0x92800171 // mov    x17, #-12
    WORD $0x4b1103e2 // neg    w2, w17
    B LBB5_341
LBB5_1661:
    WORD $0xaa1b03fd // mov    x29, x27
    B LBB5_1080
LBB5_1662:
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x910009ef // add    x15, x15, #2
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910005b1 // add    x17, x13, #1
    B LBB5_1650
LBB5_1663:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x91000daf // add    x15, x13, #3
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910009b1 // add    x17, x13, #2
    B LBB5_1650
LBB5_1664:
    WORD $0x910011af // add    x15, x13, #4
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x91000db1 // add    x17, x13, #3
    B LBB5_1650
LBB5_1665:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x910015af // add    x15, x13, #5
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910011b1 // add    x17, x13, #4
    B LBB5_1650
LBB5_1666:
    WORD $0x910019af // add    x15, x13, #6
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910015b1 // add    x17, x13, #5
    B LBB5_1650
LBB5_1667:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x91001daf // add    x15, x13, #7
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910019b1 // add    x17, x13, #6
    B LBB5_1650
LBB5_1668:
    WORD $0x8b1802ae // add    x14, x21, x24
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x394001cd // ldrb    w13, [x14]
    WORD $0x710089bf // cmp    w13, #34
    BNE LBB5_1671
LBB5_1669:
    WORD $0x910005d7 // add    x23, x14, #1
    WORD $0xcb0f022f // sub    x15, x17, x15
LBB5_1670:
    WORD $0x52800039 // mov    w25, #1
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x2a1f03f2 // mov    w18, wzr
    TST $(1<<63), R15
    BEQ LBB5_1629
    B LBB5_1628
LBB5_1671:
    WORD $0x8b1802ce // add    x14, x22, x24
    WORD $0x8b1802b1 // add    x17, x21, x24
    WORD $0x390001cd // strb    w13, [x14]
    WORD $0x3940062d // ldrb    w13, [x17, #1]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_1685
    WORD $0x390005cd // strb    w13, [x14, #1]
    WORD $0x39400a31 // ldrb    w17, [x17, #2]
    WORD $0x71008a3f // cmp    w17, #34
    BEQ LBB5_1686
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0x390009d1 // strb    w17, [x14, #2]
    WORD $0x39400db1 // ldrb    w17, [x13, #3]
    WORD $0x71008a3f // cmp    w17, #34
    BEQ LBB5_1687
    WORD $0x39000dd1 // strb    w17, [x14, #3]
    WORD $0x394011b1 // ldrb    w17, [x13, #4]
    WORD $0x71008a3f // cmp    w17, #34
    BEQ LBB5_1688
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0x390011d1 // strb    w17, [x14, #4]
    WORD $0x394015b1 // ldrb    w17, [x13, #5]
    WORD $0x71008a3f // cmp    w17, #34
    BEQ LBB5_1689
    WORD $0x390015d1 // strb    w17, [x14, #5]
    WORD $0x394019b1 // ldrb    w17, [x13, #6]
    WORD $0x71008a3f // cmp    w17, #34
    BEQ LBB5_1690
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0x390019d1 // strb    w17, [x14, #6]
    WORD $0x39401db1 // ldrb    w17, [x13, #7]
    WORD $0x71008a3f // cmp    w17, #34
    BEQ LBB5_1691
    WORD $0x39001dd1 // strb    w17, [x14, #7]
    WORD $0x91002318 // add    x24, x24, #8
    WORD $0x394021ad // ldrb    w13, [x13, #8]
    WORD $0x710089bf // cmp    w13, #34
    BNE LBB5_1671
    WORD $0x8b1802ae // add    x14, x21, x24
    WORD $0x8b1802d1 // add    x17, x22, x24
    B LBB5_1669
LBB5_1680:
    WORD $0x910021af // add    x15, x13, #8
    WORD $0xcb1102ed // sub    x13, x23, x17
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x91001db1 // add    x17, x13, #7
    B LBB5_1650
LBB5_1681:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x9280016f // mov    x15, #-12
    B LBB5_1628
LBB5_1682:
    WORD $0x8b1802cf // add    x15, x22, x24
LBB5_1683:
    WORD $0x25904041 // brkb    p1.b, p0/z, p2.b
    WORD $0x92800011 // mov    x17, #-1
    WORD $0x25208029 // cntp    x9, p0, p1.b
    WORD $0x8b0901ef // add    x15, x15, x9
LBB5_1684:
    WORD $0x4b1103e2 // neg    w2, w17
    B LBB5_341
LBB5_1685:
    WORD $0xcb0f02cd // sub    x13, x22, x15
    WORD $0x91000a37 // add    x23, x17, #2
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910005af // add    x15, x13, #1
    B LBB5_1670
LBB5_1686:
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0xcb0f02ce // sub    x14, x22, x15
    WORD $0x91000db7 // add    x23, x13, #3
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910009af // add    x15, x13, #2
    B LBB5_1670
LBB5_1687:
    WORD $0xcb0f02ce // sub    x14, x22, x15
    WORD $0x910011b7 // add    x23, x13, #4
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x91000daf // add    x15, x13, #3
    B LBB5_1670
LBB5_1688:
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0xcb0f02ce // sub    x14, x22, x15
    WORD $0x910015b7 // add    x23, x13, #5
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910011af // add    x15, x13, #4
    B LBB5_1670
LBB5_1689:
    WORD $0xcb0f02ce // sub    x14, x22, x15
    WORD $0x910019b7 // add    x23, x13, #6
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910015af // add    x15, x13, #5
    B LBB5_1670
LBB5_1690:
    WORD $0x8b1802ad // add    x13, x21, x24
    WORD $0xcb0f02ce // sub    x14, x22, x15
    WORD $0x91001db7 // add    x23, x13, #7
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910019af // add    x15, x13, #6
    B LBB5_1670
LBB5_1691:
    WORD $0xcb0f02ce // sub    x14, x22, x15
    WORD $0x910021b7 // add    x23, x13, #8
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x91001daf // add    x15, x13, #7
    B LBB5_1670
LBB5_1692:
    WORD $0x385ff22d // ldurb    w13, [x17, #-1]
    WORD $0x8b03086e // add    x14, x3, x3, lsl #2
    WORD $0xd37ff9ce // lsl    x14, x14, #1
    WORD $0x52800027 // mov    w7, #1
    WORD $0x5100c1ad // sub    w13, w13, #48
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x937ffdb2 // asr    x18, x13, #63
    WORD $0xab0d01cd // adds    x13, x14, x13
    WORD $0x9a923652 // cinc    x18, x18, hs
    WORD $0x9340024e // sbfx    x14, x18, #0, #1
    WORD $0xca1201d2 // eor    x18, x14, x18
    CMP $0, R18_PLATFORM
    BNE LBB5_1171
    TST $(1<<63), R14
    BNE LBB5_1171
    CMP $0, R1
    BEQ LBB5_1698
    WORD $0x2a1f03f2 // mov    w18, wzr
    WORD $0x9e6301a0 // ucvtf    d0, x13
    B LBB5_1158
LBB5_1696:
    WORD $0x8b1802a1 // add    x1, x21, x24
    WORD $0x52800039 // mov    w25, #1
LBB5_1697:
    WORD $0x25904041 // brkb    p1.b, p0/z, p2.b
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x8b0d0037 // add    x23, x1, x13
    B LBB5_1628
LBB5_1698:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa0d03fd // mov    x29, x13
    B LBB5_1240
LBB5_1699:
    WORD $0x92800171 // mov    x17, #-12
    WORD $0xaa1603ef // mov    x15, x22
    WORD $0x4b1103e2 // neg    w2, w17
    B LBB5_341
LBB5_1700:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x9280016f // mov    x15, #-12
    WORD $0xaa1503f7 // mov    x23, x21
    B LBB5_1628
LBB5_1701:
    WORD $0xf14005bf // cmp    x13, #1, lsl #12
    WORD $0xb900e50d // str    w13, [x8, #228]
    BHI LBB5_834
    B LBB5_1923
LBB5_1702:
LBB5_1703:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R16
    WORD $0xd284c001 // mov    x1, #9728
    WORD $0x4f04e5e3 // movi    v3.16b, #143
    WORD $0x9102e112 // add    x18, x8, #184
    WORD $0x528000ef // mov    w15, #7
    WORD $0x52800020 // mov    w0, #1
    WORD $0xf2c00021 // movk    x1, #1, lsl #32
    WORD $0x92800003 // mov    x3, #-1
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x2a0203e6 // mov    w6, w2
    WORD $0x3dc001c1 // ldr    q1, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00202 // ldr    q2, [x16, :lo12:.LCPI5_2]
LBB5_1704:
    WORD $0xaa2a03ed // mvn    x13, x10
    WORD $0xaa1103e4 // mov    x4, x17
    WORD $0x8b1101b0 // add    x16, x13, x17
    WORD $0x528000c2 // mov    w2, #6
    WORD $0x710168df // cmp    w6, #90
    BLE LBB5_1722
    WORD $0x91000491 // add    x17, x4, #1
    WORD $0x710194df // cmp    w6, #101
    BGT LBB5_1737
    WORD $0x71016cdf // cmp    w6, #91
    BNE LBB5_1897
    WORD $0xa94a350b // ldp    x11, x13, [x8, #160]
    WORD $0xf900056d // str    x13, [x11, #8]
    WORD $0xaa1081ed // orr    x13, x15, x16, lsl #32
    WORD $0xa94bb905 // ldp    x5, x14, [x8, #184]
    WORD $0xf9405102 // ldr    x2, [x8, #160]
    WORD $0xf900016d // str    x13, [x11]
    WORD $0xf9405906 // ldr    x6, [x8, #176]
    WORD $0xcb050050 // sub    x16, x2, x5
    WORD $0x91008047 // add    x7, x2, #32
    WORD $0xb100421f // cmn    x16, #16
    WORD $0x9344fe10 // asr    x16, x16, #4
    WORD $0xfa4e10e2 // ccmp    x7, x14, #2, ne
    WORD $0x9100404d // add    x13, x2, #16
    WORD $0x910004ce // add    x14, x6, #1
    WORD $0x9a8283eb // csel    x11, xzr, x2, hi
    WORD $0xa90a410d // stp    x13, x16, [x8, #160]
    WORD $0xf900590e // str    x14, [x8, #176]
    CMP $0, R11
    BEQ LBB5_1815
    WORD $0x39400086 // ldrb    w6, [x4]
    WORD $0x710080df // cmp    w6, #32
    BHI LBB5_1719
    WORD $0x9ac6200d // lsl    x13, x0, x6
    WORD $0xea0101bf // tst    x13, x1
    BEQ LBB5_1719
    WORD $0x39400486 // ldrb    w6, [x4, #1]
    WORD $0x91000631 // add    x17, x17, #1
    WORD $0x710080df // cmp    w6, #32
    BHI LBB5_1720
    WORD $0x9ac6200d // lsl    x13, x0, x6
    WORD $0xea0101bf // tst    x13, x1
    BEQ LBB5_1720
    WORD $0xf9404904 // ldr    x4, [x8, #144]
    WORD $0xcb04022d // sub    x13, x17, x4
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_1715
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x9acd206d // lsl    x13, x3, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_1718
    WORD $0x91010091 // add    x17, x4, #64
LBB5_1715:
    WORD $0xd1010224 // sub    x4, x17, #64
LBB5_1716:
    WORD $0xadc21484 // ldp    q4, q5, [x4, #64]!
    WORD $0x4e231c90 // and    v16.16b, v4.16b, v3.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411c86 // ldp    q6, q7, [x4, #32]
    WORD $0x4e231cb1 // and    v17.16b, v5.16b, v3.16b
    WORD $0x4e110011 // tbl    v17.16b, { v0.16b }, v17.16b
    WORD $0x6e248e04 // cmeq    v4.16b, v16.16b, v4.16b
    WORD $0x4e231cd2 // and    v18.16b, v6.16b, v3.16b
    WORD $0x4e120012 // tbl    v18.16b, { v0.16b }, v18.16b
    WORD $0x4e231cf3 // and    v19.16b, v7.16b, v3.16b
    WORD $0x4e130013 // tbl    v19.16b, { v0.16b }, v19.16b
    WORD $0x6e258e25 // cmeq    v5.16b, v17.16b, v5.16b
    WORD $0x6e268e46 // cmeq    v6.16b, v18.16b, v6.16b
    WORD $0x4e211ca5 // and    v5.16b, v5.16b, v1.16b
    WORD $0x4e0200a5 // tbl    v5.16b, { v5.16b }, v2.16b
    WORD $0x4e211c84 // and    v4.16b, v4.16b, v1.16b
    WORD $0x6e278e67 // cmeq    v7.16b, v19.16b, v7.16b
    WORD $0x4e020084 // tbl    v4.16b, { v4.16b }, v2.16b
    WORD $0x4e211cc6 // and    v6.16b, v6.16b, v1.16b
    WORD $0x4e0200c6 // tbl    v6.16b, { v6.16b }, v2.16b
    WORD $0x4e211ce7 // and    v7.16b, v7.16b, v1.16b
    WORD $0x4e71b8a5 // addv    h5, v5.8h
    WORD $0x4e0200e7 // tbl    v7.16b, { v7.16b }, v2.16b
    WORD $0x4e71b884 // addv    h4, v4.8h
    WORD $0x1e2600ad // fmov    w13, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e26008e // fmov    w14, s4
    WORD $0x1e2600b1 // fmov    w17, s5
    WORD $0x1e2600c6 // fmov    w6, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa1181cd // orr    x13, x14, x17, lsl #32
    WORD $0xaa06c1ad // orr    x13, x13, x6, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_1716
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa9093504 // stp    x4, x13, [x8, #144]
LBB5_1718:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d0091 // add    x17, x4, x13
    WORD $0x38401626 // ldrb    w6, [x17], #1
LBB5_1719:
    WORD $0x710174df // cmp    w6, #93
    BNE LBB5_1704
    B LBB5_1721
LBB5_1720:
    WORD $0x710174df // cmp    w6, #93
    BNE LBB5_1704
LBB5_1721:
    WORD $0xb940d10b // ldr    w11, [x8, #208]
    WORD $0x8b1010ad // add    x13, x5, x16, lsl #4
    WORD $0xcb02004e // sub    x14, x2, x2
    WORD $0x1100056b // add    w11, w11, #1
    WORD $0xb900d10b // str    w11, [x8, #208]
    WORD $0x910041cb // add    x11, x14, #16
    WORD $0xf94005ae // ldr    x14, [x13, #8]
    WORD $0xd344fd6b // lsr    x11, x11, #4
    WORD $0xf900550e // str    x14, [x8, #168]
    WORD $0xb9000c4b // str    w11, [x2, #12]
    WORD $0xf940004b // ldr    x11, [x2]
    WORD $0xb90009bf // str    wzr, [x13, #8]
    WORD $0xf940590d // ldr    x13, [x8, #176]
    WORD $0xb940e50e // ldr    w14, [x8, #228]
    WORD $0x92609d6b // and    x11, x11, #0xffffffff000000ff
    WORD $0xeb0e01bf // cmp    x13, x14
    WORD $0xf900004b // str    x11, [x2]
    BHI LBB5_1701
    B LBB5_1923
LBB5_1722:
    WORD $0x5100c0cd // sub    w13, w6, #48
    WORD $0x710029bf // cmp    w13, #10
    BHS LBB5_1750
    WORD $0x2a1f03f2 // mov    w18, wzr
    WORD $0x520003ed // eor    w13, wzr, #0x1
    WORD $0x3941c10e // ldrb    w14, [x8, #112]
    WORD $0x934001a6 // sbfx    x6, x13, #0, #1
    WORD $0xcb0d0085 // sub    x5, x4, x13
    TST $(1<<1), R14
    BNE LBB5_1753
LBB5_1724:
    WORD $0x394000a7 // ldrb    w7, [x5]
    WORD $0x7100c0ed // subs    w13, w7, #48
    BNE LBB5_1744
    WORD $0xaa0503ef // mov    x15, x5
    WORD $0x38401ded // ldrb    w13, [x15, #1]!
    WORD $0x7100b9bf // cmp    w13, #46
    BEQ LBB5_1898
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x710115bf // cmp    w13, #69
    BEQ LBB5_1728
    WORD $0x710195bf // cmp    w13, #101
    BNE LBB5_1916
LBB5_1728:
    WORD $0x2a0103e3 // mov    w3, w1
LBB5_1729:
    WORD $0xaa0f03e7 // mov    x7, x15
    WORD $0x12800011 // mov    w17, #-1
    WORD $0x38401cf3 // ldrb    w19, [x7, #1]!
    WORD $0x7100b67f // cmp    w19, #45
    BEQ LBB5_1731
    WORD $0x52800031 // mov    w17, #1
    WORD $0x7100ae7f // cmp    w19, #43
    BNE LBB5_1732
LBB5_1731:
    WORD $0x38402df3 // ldrb    w19, [x15, #2]!
    WORD $0xaa0f03e7 // mov    x7, x15
LBB5_1732:
    WORD $0x52800062 // mov    w2, #3
    WORD $0x5100c26d // sub    w13, w19, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_2421
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x5280014f // mov    w15, #10
LBB5_1734:
    WORD $0x8b0d00ee // add    x14, x7, x13
    WORD $0x1b0f4c42 // madd    w2, w2, w15, w19
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x5100c042 // sub    w2, w2, #48
    WORD $0x394005d3 // ldrb    w19, [x14, #1]
    WORD $0x5100c26e // sub    w14, w19, #48
    WORD $0x710029df // cmp    w14, #10
    BLO LBB5_1734
    WORD $0x8b0d00ef // add    x15, x7, x13
    WORD $0xd10005ae // sub    x14, x13, #1
    WORD $0xf10025df // cmp    x14, #9
    BHS LBB5_2422
LBB5_1736:
    WORD $0x1b110c43 // madd    w3, w2, w17, w3
    B LBB5_1958
LBB5_1737:
    WORD $0x7101ccdf // cmp    w6, #115
    BGT LBB5_1789
    WORD $0x710198df // cmp    w6, #102
    BEQ LBB5_1816
    WORD $0x7101b8df // cmp    w6, #110
    BNE LBB5_1897
    WORD $0xaa0403f1 // mov    x17, x4
    WORD $0x528001a1 // mov    w1, #13
    WORD $0x3840162d // ldrb    w13, [x17], #1
    WORD $0x7101d5bf // cmp    w13, #117
    BNE LBB5_1743
    WORD $0x3940048d // ldrb    w13, [x4, #1]
    WORD $0x91000891 // add    x17, x4, #2
    WORD $0x7101b1bf // cmp    w13, #108
    BNE LBB5_1743
    WORD $0x3940088d // ldrb    w13, [x4, #2]
    WORD $0x91000c91 // add    x17, x4, #3
    WORD $0x7101b1bf // cmp    w13, #108
    WORD $0x1a8103e1 // csel    w1, wzr, w1, eq
LBB5_1743:
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0xd3607e0d // lsl    x13, x16, #32
    B LBB5_1826
LBB5_1744:
    WORD $0x52800062 // mov    w2, #3
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1803
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0xaa1f03f1 // mov    x17, xzr
    WORD $0x5280014d // mov    w13, #10
LBB5_1746:
    WORD $0x8b1100ae // add    x14, x5, x17
    WORD $0x9b0d7c0f // mul    x15, x0, x13
    WORD $0x91000631 // add    x17, x17, #1
    WORD $0x8b2741ef // add    x15, x15, w7, uxtw
    WORD $0x394005c7 // ldrb    w7, [x14, #1]
    WORD $0xd100c1e0 // sub    x0, x15, #48
    WORD $0x5100c0ee // sub    w14, w7, #48
    WORD $0x710029df // cmp    w14, #10
    BLO LBB5_1746
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0x8b1100af // add    x15, x5, x17
    WORD $0xd100062d // sub    x13, x17, #1
    WORD $0xf1004dbf // cmp    x13, #19
    BHS LBB5_2423
LBB5_1748:
    WORD $0x7100b8ff // cmp    w7, #46
    BNE LBB5_1909
    WORD $0x38401de7 // ldrb    w7, [x15, #1]!
    WORD $0x52800062 // mov    w2, #3
    WORD $0xaa0f03f3 // mov    x19, x15
    WORD $0x5100c0ed // sub    w13, w7, #48
    WORD $0x710029bf // cmp    w13, #10
    BLO LBB5_1904
    B LBB5_2380
LBB5_1750:
    WORD $0x710088df // cmp    w6, #34
    BEQ LBB5_1827
    WORD $0x7100b4df // cmp    w6, #45
    BNE LBB5_1897
    WORD $0x52800032 // mov    w18, #1
    WORD $0x5200024d // eor    w13, w18, #0x1
    WORD $0x3941c10e // ldrb    w14, [x8, #112]
    WORD $0x934001a6 // sbfx    x6, x13, #0, #1
    WORD $0xcb0d0085 // sub    x5, x4, x13
    TST $(1<<1), R14
    BEQ LBB5_1724
LBB5_1753:
    WORD $0xcb05012d // sub    x13, x9, x5
    WORD $0x92800002 // mov    x2, #-1
    WORD $0xeb0601b1 // subs    x17, x13, x6
    BEQ LBB5_2013
    WORD $0x394000ad // ldrb    w13, [x5]
    WORD $0x924000c3 // and    x3, x6, #0x1
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_1758
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x52800022 // mov    w2, #1
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf100063f // cmp    x17, #1
    BEQ LBB5_2014
    WORD $0x394004ad // ldrb    w13, [x5, #1]
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x52800022 // mov    w2, #1
    WORD $0x5100b9ad // sub    w13, w13, #46
    WORD $0x7100ddbf // cmp    w13, #55
    BHI LBB5_2028
    WORD $0x5280002e // mov    w14, #1
    WORD $0xb20903ef // mov    x15, #36028797027352576
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf280002f // movk    x15, #1
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_2014
LBB5_1758:
    WORD $0x9280000f // mov    x15, #-1
    WORD $0xf100423f // cmp    x17, #16
    BLO LBB5_2420
    WORD $0x4f01e5c3 // movi    v3.16b, #46
    WORD $0xaa1f03e2 // mov    x2, xzr
    WORD $0x4f01e564 // movi    v4.16b, #43
    WORD $0x92800001 // mov    x1, #-1
    WORD $0x4f01e5a5 // movi    v5.16b, #45
    WORD $0x12800007 // mov    w7, #-1
    WORD $0x4f06e606 // movi    v6.16b, #208
    WORD $0x92800000 // mov    x0, #-1
    WORD $0x4f00e547 // movi    v7.16b, #10
    WORD $0x4f06e7f0 // movi    v16.16b, #223
    WORD $0x4f02e4b1 // movi    v17.16b, #69
LBB5_1760:
    WORD $0x3ce268b2 // ldr    q18, [x5, x2]
    WORD $0x6e238e53 // cmeq    v19.16b, v18.16b, v3.16b
    WORD $0x6e248e54 // cmeq    v20.16b, v18.16b, v4.16b
    WORD $0x6e258e55 // cmeq    v21.16b, v18.16b, v5.16b
    WORD $0x4e268656 // add    v22.16b, v18.16b, v6.16b
    WORD $0x4e301e52 // and    v18.16b, v18.16b, v16.16b
    WORD $0x6e318e52 // cmeq    v18.16b, v18.16b, v17.16b
    WORD $0x6e3634f6 // cmhi    v22.16b, v7.16b, v22.16b
    WORD $0x4eb41eb4 // orr    v20.16b, v21.16b, v20.16b
    WORD $0x4eb21e75 // orr    v21.16b, v19.16b, v18.16b
    WORD $0x4eb51ed5 // orr    v21.16b, v22.16b, v21.16b
    WORD $0x4eb41eb5 // orr    v21.16b, v21.16b, v20.16b
    WORD $0x4e211eb5 // and    v21.16b, v21.16b, v1.16b
    WORD $0x4e0202b5 // tbl    v21.16b, { v21.16b }, v2.16b
    WORD $0x4e211e73 // and    v19.16b, v19.16b, v1.16b
    WORD $0x4e211e52 // and    v18.16b, v18.16b, v1.16b
    WORD $0x4e211e94 // and    v20.16b, v20.16b, v1.16b
    WORD $0x4e71bab5 // addv    h21, v21.8h
    WORD $0x4e020273 // tbl    v19.16b, { v19.16b }, v2.16b
    WORD $0x4e020252 // tbl    v18.16b, { v18.16b }, v2.16b
    WORD $0x4e020294 // tbl    v20.16b, { v20.16b }, v2.16b
    WORD $0x1e2602ad // fmov    w13, s21
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x4e71ba52 // addv    h18, v18.8h
    WORD $0x2a2d03ed // mvn    w13, w13
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x32103dad // orr    w13, w13, #0xffff0000
    WORD $0x5ac001ad // rbit    w13, w13
    WORD $0x1e260276 // fmov    w22, s19
    WORD $0x1e260255 // fmov    w21, s18
    WORD $0x5ac011b3 // clz    w19, w13
    WORD $0x1e260294 // fmov    w20, s20
    WORD $0x7100427f // cmp    w19, #16
    BEQ LBB5_1762
    WORD $0x1ad320ed // lsl    w13, w7, w19
    WORD $0x0a2d02d6 // bic    w22, w22, w13
    WORD $0x0a2d02b5 // bic    w21, w21, w13
    WORD $0x0a2d0294 // bic    w20, w20, w13
LBB5_1762:
    WORD $0x510006cd // sub    w13, w22, #1
    WORD $0x6a1601ad // ands    w13, w13, w22
    BNE LBB5_2100
    WORD $0x510006ad // sub    w13, w21, #1
    WORD $0x6a1501ad // ands    w13, w13, w21
    BNE LBB5_2100
    WORD $0x5100068d // sub    w13, w20, #1
    WORD $0x6a1401ad // ands    w13, w13, w20
    BNE LBB5_2100
    CMP $0, R22
    BEQ LBB5_1768
    WORD $0x5ac002cd // rbit    w13, w22
    WORD $0xb10005ff // cmn    x15, #1
    WORD $0x5ac011ad // clz    w13, w13
    BNE LBB5_2224
    WORD $0x8b0d004f // add    x15, x2, x13
LBB5_1768:
    CMP $0, R21
    BEQ LBB5_1771
    WORD $0x5ac002ad // rbit    w13, w21
    WORD $0xb100041f // cmn    x0, #1
    WORD $0x5ac011ad // clz    w13, w13
    BNE LBB5_2224
    WORD $0x8b0d0040 // add    x0, x2, x13
LBB5_1771:
    CMP $0, R20
    BEQ LBB5_1774
    WORD $0x5ac0028d // rbit    w13, w20
    WORD $0xb100043f // cmn    x1, #1
    WORD $0x5ac011ad // clz    w13, w13
    BNE LBB5_2224
    WORD $0x8b0d0041 // add    x1, x2, x13
LBB5_1774:
    WORD $0x7100427f // cmp    w19, #16
    BNE LBB5_1805
    WORD $0xd1004231 // sub    x17, x17, #16
    WORD $0x91004042 // add    x2, x2, #16
    WORD $0xf1003e3f // cmp    x17, #15
    BHI LBB5_1760
    WORD $0x8b03012d // add    x13, x9, x3
    WORD $0x8b0200a3 // add    x3, x5, x2
    WORD $0xcb0601ad // sub    x13, x13, x6
    WORD $0xcb0401ad // sub    x13, x13, x4
    WORD $0xeb0201bf // cmp    x13, x2
    BEQ LBB5_1806
LBB5_1777:
    WORD $0xaa2303ed // mvn    x13, x3
    WORD $0x8b0400ce // add    x14, x6, x4
    WORD $0x8b110064 // add    x4, x3, x17
    WORD $0x8b0e01a2 // add    x2, x13, x14
    WORD $0xcb050066 // sub    x6, x3, x5
    WORD $0xaa0303e7 // mov    x7, x3
    B LBB5_1780
LBB5_1778:
    WORD $0xb100041f // cmn    x0, #1
    WORD $0xaa0603e0 // mov    x0, x6
    BNE LBB5_1918
LBB5_1779:
    WORD $0xd1000631 // sub    x17, x17, #1
    WORD $0xd1000442 // sub    x2, x2, #1
    WORD $0x910004c6 // add    x6, x6, #1
    WORD $0xaa0703e3 // mov    x3, x7
    CMP $0, R17
    BEQ LBB5_2025
LBB5_1780:
    WORD $0x384014ed // ldrb    w13, [x7], #1
    WORD $0x5100c1ae // sub    w14, w13, #48
    WORD $0x710029df // cmp    w14, #10
    BLO LBB5_1779
    WORD $0x7100b5bf // cmp    w13, #45
    BLE LBB5_1786
    WORD $0x710195bf // cmp    w13, #101
    BEQ LBB5_1778
    WORD $0x710115bf // cmp    w13, #69
    BEQ LBB5_1778
    WORD $0x7100b9bf // cmp    w13, #46
    BNE LBB5_1806
    WORD $0xb10005ff // cmn    x15, #1
    WORD $0xaa0603ef // mov    x15, x6
    BEQ LBB5_1779
    B LBB5_1918
LBB5_1786:
    WORD $0x7100adbf // cmp    w13, #43
    BEQ LBB5_1788
    WORD $0x7100b5bf // cmp    w13, #45
    BNE LBB5_1806
LBB5_1788:
    WORD $0xb100043f // cmn    x1, #1
    WORD $0xaa0603e1 // mov    x1, x6
    BEQ LBB5_1779
    B LBB5_1918
LBB5_1789:
    WORD $0x7101d0df // cmp    w6, #116
    BEQ LBB5_1821
    WORD $0x7101ecdf // cmp    w6, #123
    BNE LBB5_1897
    WORD $0xa94a350b // ldp    x11, x13, [x8, #160]
    WORD $0xf900056d // str    x13, [x11, #8]
    WORD $0x528000cd // mov    w13, #6
    WORD $0xa94bb900 // ldp    x0, x14, [x8, #184]
    WORD $0xaa1081ad // orr    x13, x13, x16, lsl #32
    WORD $0xf940510f // ldr    x15, [x8, #160]
    WORD $0xf9405902 // ldr    x2, [x8, #176]
    WORD $0xf900016d // str    x13, [x11]
    WORD $0xcb0001f0 // sub    x16, x15, x0
    WORD $0x910081e1 // add    x1, x15, #32
    WORD $0xb100421f // cmn    x16, #16
    WORD $0x9100044d // add    x13, x2, #1
    WORD $0xfa4e1022 // ccmp    x1, x14, #2, ne
    WORD $0x9344fe01 // asr    x1, x16, #4
    WORD $0x910041f0 // add    x16, x15, #16
    WORD $0xf900590d // str    x13, [x8, #176]
    WORD $0x9a8f83eb // csel    x11, xzr, x15, hi
    WORD $0xa90a0510 // stp    x16, x1, [x8, #160]
    CMP $0, R11
    BEQ LBB5_1815
    WORD $0x39400082 // ldrb    w2, [x4]
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_1920
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ac221ad // lsl    x13, x13, x2
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1920
    WORD $0x38401622 // ldrb    w2, [x17], #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_1920
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ac221ad // lsl    x13, x13, x2
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1920
    WORD $0xf9404902 // ldr    x2, [x8, #144]
    WORD $0xcb02022d // sub    x13, x17, x2
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_1799
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x92800011 // mov    x17, #-1
    WORD $0x9acd222d // lsl    x13, x17, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_1802
    WORD $0x91010051 // add    x17, x2, #64
LBB5_1799:
    WORD $0x4f04e5e3 // movi    v3.16b, #143
    WORD $0xd1010222 // sub    x2, x17, #64
LBB5_1800:
    WORD $0xadc21444 // ldp    q4, q5, [x2, #64]!
    WORD $0x4e231c90 // and    v16.16b, v4.16b, v3.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411c46 // ldp    q6, q7, [x2, #32]
    WORD $0x4e231cb1 // and    v17.16b, v5.16b, v3.16b
    WORD $0x4e110011 // tbl    v17.16b, { v0.16b }, v17.16b
    WORD $0x6e248e04 // cmeq    v4.16b, v16.16b, v4.16b
    WORD $0x4e231cd2 // and    v18.16b, v6.16b, v3.16b
    WORD $0x4e120012 // tbl    v18.16b, { v0.16b }, v18.16b
    WORD $0x4e231cf3 // and    v19.16b, v7.16b, v3.16b
    WORD $0x4e130013 // tbl    v19.16b, { v0.16b }, v19.16b
    WORD $0x6e258e25 // cmeq    v5.16b, v17.16b, v5.16b
    WORD $0x6e268e46 // cmeq    v6.16b, v18.16b, v6.16b
    WORD $0x4e211ca5 // and    v5.16b, v5.16b, v1.16b
    WORD $0x4e0200a5 // tbl    v5.16b, { v5.16b }, v2.16b
    WORD $0x4e211c84 // and    v4.16b, v4.16b, v1.16b
    WORD $0x6e278e67 // cmeq    v7.16b, v19.16b, v7.16b
    WORD $0x4e020084 // tbl    v4.16b, { v4.16b }, v2.16b
    WORD $0x4e211cc6 // and    v6.16b, v6.16b, v1.16b
    WORD $0x4e0200c6 // tbl    v6.16b, { v6.16b }, v2.16b
    WORD $0x4e211ce7 // and    v7.16b, v7.16b, v1.16b
    WORD $0x4e71b8a5 // addv    h5, v5.8h
    WORD $0x4e0200e7 // tbl    v7.16b, { v7.16b }, v2.16b
    WORD $0x4e71b884 // addv    h4, v4.8h
    WORD $0x1e2600ad // fmov    w13, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e26008e // fmov    w14, s4
    WORD $0x1e2600b1 // fmov    w17, s5
    WORD $0x1e2600c3 // fmov    w3, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa1181cd // orr    x13, x14, x17, lsl #32
    WORD $0xaa03c1ad // orr    x13, x13, x3, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_1800
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa9093502 // stp    x2, x13, [x8, #144]
LBB5_1802:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d0051 // add    x17, x2, x13
    WORD $0x38401622 // ldrb    w2, [x17], #1
    B LBB5_1921
LBB5_1803:
    WORD $0xaa0503ef // mov    x15, x5
    WORD $0xf1004d9f // cmp    x12, #19
    BNE LBB5_2381
LBB5_1804:
    WORD $0x2a0203e1 // mov    w1, w2
    WORD $0xaa1d03e0 // mov    x0, x29
    B LBB5_2384
LBB5_1805:
    WORD $0x8b3340ad // add    x13, x5, w19, uxtw
    WORD $0x8b0201a3 // add    x3, x13, x2
LBB5_1806:
    WORD $0x92800002 // mov    x2, #-1
    CMP $0, R15
    BEQ LBB5_2013
LBB5_1807:
    CMP $0, R1
    BEQ LBB5_2013
    CMP $0, R0
    BEQ LBB5_2013
    WORD $0xcb05006d // sub    x13, x3, x5
    WORD $0xd10005b1 // sub    x17, x13, #1
    WORD $0xeb1101ff // cmp    x15, x17
    BEQ LBB5_1917
    WORD $0xeb11003f // cmp    x1, x17
    BEQ LBB5_1917
    WORD $0xeb11001f // cmp    x0, x17
    BEQ LBB5_1917
    WORD $0xf100042e // subs    x14, x1, #1
    BLT LBB5_2010
    WORD $0xeb0e001f // cmp    x0, x14
    BEQ LBB5_2010
    WORD $0xaa2103e2 // mvn    x2, x1
    B LBB5_2013
LBB5_1815:
    WORD $0x52800162 // mov    w2, #11
    WORD $0xaa0403ef // mov    x15, x4
    B LBB5_341
LBB5_1816:
    WORD $0xaa0403f1 // mov    x17, x4
    WORD $0x528001a1 // mov    w1, #13
    WORD $0x3840162d // ldrb    w13, [x17], #1
    WORD $0x710185bf // cmp    w13, #97
    BNE LBB5_1820
    WORD $0x3940048d // ldrb    w13, [x4, #1]
    WORD $0x91000891 // add    x17, x4, #2
    WORD $0x7101b1bf // cmp    w13, #108
    BNE LBB5_1820
    WORD $0x3940088d // ldrb    w13, [x4, #2]
    WORD $0x91000c91 // add    x17, x4, #3
    WORD $0x7101cdbf // cmp    w13, #115
    BNE LBB5_1820
    WORD $0x39400c8d // ldrb    w13, [x4, #3]
    WORD $0x91001091 // add    x17, x4, #4
    WORD $0x710195bf // cmp    w13, #101
    WORD $0x1a8103e1 // csel    w1, wzr, w1, eq
LBB5_1820:
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0x5280004d // mov    w13, #2
    B LBB5_1825
LBB5_1821:
    WORD $0xaa0403f1 // mov    x17, x4
    WORD $0x528001a1 // mov    w1, #13
    WORD $0x3840162d // ldrb    w13, [x17], #1
    WORD $0x7101c9bf // cmp    w13, #114
    BNE LBB5_1824
    WORD $0x3940048d // ldrb    w13, [x4, #1]
    WORD $0x91000891 // add    x17, x4, #2
    WORD $0x7101d5bf // cmp    w13, #117
    BNE LBB5_1824
    WORD $0x3940088d // ldrb    w13, [x4, #2]
    WORD $0x91000c91 // add    x17, x4, #3
    WORD $0x710195bf // cmp    w13, #101
    WORD $0x1a8103e1 // csel    w1, wzr, w1, eq
LBB5_1824:
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0x5280014d // mov    w13, #10
LBB5_1825:
    WORD $0xaa1081ad // orr    x13, x13, x16, lsl #32
LBB5_1826:
    WORD $0xaa1203ee // mov    x14, x18
    WORD $0xf80105cd // str    x13, [x14], #16
    WORD $0xf900510e // str    x14, [x8, #160]
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a8101a2 // csel    w2, w13, w1, eq
    CMP $0, R1
    BEQ LBB5_2388
    B LBB5_2444
LBB5_1827:
    WORD $0xf940390f // ldr    x15, [x8, #112]
    WORD $0x2518e3e0 // ptrue    p0.b
    WORD $0x2518e401 // pfalse    p1.b
    TST $(1<<5), R15
    BNE LBB5_2430
    WORD $0xaa0403f1 // mov    x17, x4
    WORD $0x2538cb83 // mov    z3.b, #92
    WORD $0x2538c444 // mov    z4.b, #34
    WORD $0x2518e402 // pfalse    p2.b
    B LBB5_1831
LBB5_1829:
    WORD $0x25904084 // brkb    p4.b, p0/z, p4.b
    WORD $0x25434084 // ands    p4.b, p0/z, p4.b, p3.b
    BNE LBB5_1836
LBB5_1830:
    WORD $0x91008231 // add    x17, x17, #32
LBB5_1831:
    WORD $0xa400a225 // ld1b    { z5.b }, p0/z, [x17]
    WORD $0x2404a0a4 // cmpeq    p4.b, p0/z, z5.b, z4.b
    WORD $0x2403a0a3 // cmpeq    p3.b, p0/z, z5.b, z3.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BEQ LBB5_1833
    WORD $0x259040c5 // brkb    p5.b, p0/z, p6.b
    WORD $0x250440a5 // and    p5.b, p0/z, p5.b, p4.b
LBB5_1833:
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_2441
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_1829
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_1830
LBB5_1836:
    WORD $0x25904062 // brkb    p2.b, p0/z, p3.b
    WORD $0x5299fa0e // mov    w14, #53200
    WORD $0x2520804d // cntp    x13, p0, p2.b
    WORD $0x52832320 // mov    w0, #6425
    WORD $0x8b0d0231 // add    x17, x17, x13
    WORD $0x5288c8c2 // mov    w2, #17990
    WORD $0x52872725 // mov    w5, #14649
    WORD $0x52848014 // mov    w20, #9216
    WORD $0x72b9f9ee // movk    w14, #53199, lsl #16
    WORD $0x3201c3f2 // mov    w18, #-2139062144
    WORD $0x72a32320 // movk    w0, #6425, lsl #16
    WORD $0x3202c7e1 // mov    w1, #-1061109568
    WORD $0x72a8c8c2 // movk    w2, #17990, lsl #16
    WORD $0x3203cbe3 // mov    w3, #-522133280
    WORD $0x72a72725 // movk    w5, #14649, lsl #16
    WORD $0x3200c3e6 // mov    w6, #16843009
    WORD $0x5297fde7 // mov    w7, #49135
    WORD $0x528017b3 // mov    w19, #189
    WORD $0x72bf9414 // movk    w20, #64672, lsl #16
    WORD $0xaa1103fa // mov    x26, x17
    WORD $0xaa1103f9 // mov    x25, x17
    ADR ESCAPED_TAB, R21
    WORD $0x910002b5 // add    x21, x21, :lo12:ESCAPED_TAB
    WORD $0x2538cb83 // mov    z3.b, #92
    WORD $0x2538c444 // mov    z4.b, #34
    WORD $0x2538c3e5 // mov    z5.b, #31
LBB5_1837:
    WORD $0x3940074d // ldrb    w13, [x26, #1]
    WORD $0xf101d5bf // cmp    x13, #117
    BEQ LBB5_1840
    WORD $0x386d6aad // ldrb    w13, [x21, x13]
    CMP $0, R13
    BEQ LBB5_2239
    WORD $0x3800172d // strb    w13, [x25], #1
    WORD $0x91000b56 // add    x22, x26, #2
    WORD $0xaa1903f7 // mov    x23, x25
    B LBB5_1860
LBB5_1840:
    WORD $0xb840234d // ldur    w13, [x26, #2]
    WORD $0x0a2d0256 // bic    w22, w18, w13
    WORD $0x0b0e01b7 // add    w23, w13, w14
    WORD $0x6a1702df // tst    w22, w23
    BNE LBB5_2465
    WORD $0x0b0001b7 // add    w23, w13, w0
    WORD $0x2a0d02f7 // orr    w23, w23, w13
    WORD $0x7201c2ff // tst    w23, #0x80808080
    BNE LBB5_2465
    WORD $0x1200d9b7 // and    w23, w13, #0x7f7f7f7f
    WORD $0x4b170038 // sub    w24, w1, w23
    WORD $0x0b0202fb // add    w27, w23, w2
    WORD $0x0a1b0318 // and    w24, w24, w27
    WORD $0x6a16031f // tst    w24, w22
    BNE LBB5_2465
    WORD $0x4b170078 // sub    w24, w3, w23
    WORD $0x0b0502f7 // add    w23, w23, w5
    WORD $0x0a170317 // and    w23, w24, w23
    WORD $0x6a1602ff // tst    w23, w22
    BNE LBB5_2465
    WORD $0x5ac009ad // rev    w13, w13
    WORD $0x91001b56 // add    x22, x26, #6
    WORD $0x1200cdb1 // and    w17, w13, #0xf0f0f0f
    WORD $0x0a6d10cd // bic    w13, w6, w13, lsr #4
    WORD $0x2a0d0dad // orr    w13, w13, w13, lsl #3
    WORD $0x0b1101ad // add    w13, w13, w17
    WORD $0x2a4d11ad // orr    w13, w13, w13, lsr #4
    WORD $0x53105db1 // ubfx    w17, w13, #16, #8
    WORD $0x12001dad // and    w13, w13, #0xff
    WORD $0x2a1121b1 // orr    w17, w13, w17, lsl #8
    WORD $0x7102023f // cmp    w17, #128
    BLO LBB5_1892
    WORD $0x91001337 // add    x23, x25, #4
LBB5_1846:
    WORD $0x711ffe3f // cmp    w17, #2047
    BLS LBB5_1894
    WORD $0x51403a2d // sub    w13, w17, #14, lsl #12
    WORD $0x312005bf // cmn    w13, #2049
    BLS LBB5_1858
    WORD $0x530a7e2d // lsr    w13, w17, #10
    WORD $0x7100d9bf // cmp    w13, #54
    BHI LBB5_1895
    WORD $0x394002cd // ldrb    w13, [x22]
    WORD $0x710171bf // cmp    w13, #92
    BNE LBB5_1895
    WORD $0x394006cd // ldrb    w13, [x22, #1]
    WORD $0x7101d5bf // cmp    w13, #117
    BNE LBB5_1895
    WORD $0xb84022cd // ldur    w13, [x22, #2]
    WORD $0x0a2d0258 // bic    w24, w18, w13
    WORD $0x0b0e01b9 // add    w25, w13, w14
    WORD $0x6a19031f // tst    w24, w25
    BNE LBB5_2480
    WORD $0x0b0001b9 // add    w25, w13, w0
    WORD $0x2a0d0339 // orr    w25, w25, w13
    WORD $0x7201c33f // tst    w25, #0x80808080
    BNE LBB5_2480
    WORD $0x1200d9b9 // and    w25, w13, #0x7f7f7f7f
    WORD $0x4b19003a // sub    w26, w1, w25
    WORD $0x0b02033b // add    w27, w25, w2
    WORD $0x0a1b035a // and    w26, w26, w27
    WORD $0x6a18035f // tst    w26, w24
    BNE LBB5_2480
    WORD $0x4b19007a // sub    w26, w3, w25
    WORD $0x0b050339 // add    w25, w25, w5
    WORD $0x0a190359 // and    w25, w26, w25
    WORD $0x6a18033f // tst    w25, w24
    BNE LBB5_2480
    WORD $0x5ac009ad // rev    w13, w13
    WORD $0x91001ad6 // add    x22, x22, #6
    WORD $0x1200cdb8 // and    w24, w13, #0xf0f0f0f
    WORD $0x0a6d10cd // bic    w13, w6, w13, lsr #4
    WORD $0x2a0d0dad // orr    w13, w13, w13, lsl #3
    WORD $0x0b1801ad // add    w13, w13, w24
    WORD $0x2a4d11b8 // orr    w24, w13, w13, lsr #4
    WORD $0x53087f0d // lsr    w13, w24, #8
    WORD $0x12181dad // and    w13, w13, #0xff00
    WORD $0x514039b9 // sub    w25, w13, #14, lsl #12
    WORD $0x33001f0d // bfxil    w13, w24, #0, #8
    WORD $0x3110073f // cmn    w25, #1025
    BHI LBB5_1896
    WORD $0x781fc2e7 // sturh    w7, [x23, #-4]
    WORD $0x2a0d03f1 // mov    w17, w13
    WORD $0x381fe2f3 // sturb    w19, [x23, #-2]
    WORD $0x91000ef7 // add    x23, x23, #3
    WORD $0x710201bf // cmp    w13, #128
    BHS LBB5_1846
    WORD $0xd10012f9 // sub    x25, x23, #4
    B LBB5_1893
LBB5_1858:
    WORD $0x530c7e2d // lsr    w13, w17, #12
    WORD $0x52801018 // mov    w24, #128
    WORD $0x52801019 // mov    w25, #128
    WORD $0x321b09ad // orr    w13, w13, #0xe0
    WORD $0x33062e38 // bfxil    w24, w17, #6, #6
    WORD $0x33001639 // bfxil    w25, w17, #0, #6
    WORD $0xd10006f1 // sub    x17, x23, #1
    WORD $0x381fc2ed // sturb    w13, [x23, #-4]
    WORD $0x381fd2f8 // sturb    w24, [x23, #-3]
    WORD $0x381fe2f9 // sturb    w25, [x23, #-2]
LBB5_1859:
    WORD $0xaa1103f7 // mov    x23, x17
LBB5_1860:
    WORD $0x394002cd // ldrb    w13, [x22]
    WORD $0xaa1603f1 // mov    x17, x22
    WORD $0xaa1603fa // mov    x26, x22
    WORD $0xaa1703f9 // mov    x25, x23
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1837
    WORD $0xaa1d03fb // mov    x27, x29
    WORD $0xaa1f03f8 // mov    x24, xzr
    WORD $0xa41842c6 // ld1b    { z6.b }, p0/z, [x22, x24]
    WORD $0x25814422 // mov    p2.b, p1.b
    TST $(1<<5), R15
    BEQ LBB5_1865
    B LBB5_1863
LBB5_1862:
    WORD $0xe41842e6 // st1b    { z6.b }, p0, [x23, x24]
    WORD $0x91008318 // add    x24, x24, #32
    WORD $0xa41842c6 // ld1b    { z6.b }, p0/z, [x22, x24]
    WORD $0x25814422 // mov    p2.b, p1.b
    TST $(1<<5), R15
    BEQ LBB5_1865
LBB5_1863:
    WORD $0x2405a0c2 // cmpeq    p2.b, p0/z, z6.b, z5.b
    WORD $0x2404a0c4 // cmpeq    p4.b, p0/z, z6.b, z4.b
    WORD $0x2403a0c3 // cmpeq    p3.b, p0/z, z6.b, z3.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BNE LBB5_1866
LBB5_1864:
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BEQ LBB5_1867
    B LBB5_2453
LBB5_1865:
    WORD $0x2404a0c4 // cmpeq    p4.b, p0/z, z6.b, z4.b
    WORD $0x2403a0c3 // cmpeq    p3.b, p0/z, z6.b, z3.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BEQ LBB5_1864
LBB5_1866:
    WORD $0x259040c5 // brkb    p5.b, p0/z, p6.b
    WORD $0x250440a5 // and    p5.b, p0/z, p5.b, p4.b
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_2453
LBB5_1867:
    WORD $0x2550c080 // ptest    p0, p4.b
    WORD $0x25824845 // mov    p5.b, p2.b
    WORD $0x1a9f07ed // cset    w13, ne
    BEQ LBB5_1869
    WORD $0x25904085 // brkb    p5.b, p0/z, p4.b
    WORD $0x250240a5 // and    p5.b, p0/z, p5.b, p2.b
LBB5_1869:
    TST $(1<<5), R15
    BEQ LBB5_1871
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_2477
LBB5_1871:
    CMP $0, R13
    BEQ LBB5_1873
    WORD $0x25904082 // brkb    p2.b, p0/z, p4.b
    WORD $0x25034043 // and    p3.b, p0/z, p2.b, p3.b
LBB5_1873:
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_1862
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x8b1802f9 // add    x25, x23, x24
    WORD $0xaa1103fa // mov    x26, x17
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x3940022d // ldrb    w13, [x17]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1837
LBB5_1875:
    WORD $0x8b1802f9 // add    x25, x23, x24
    WORD $0x8b1802da // add    x26, x22, x24
    WORD $0x3900032d // strb    w13, [x25]
    WORD $0x3940074d // ldrb    w13, [x26, #1]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1884
    WORD $0x3900072d // strb    w13, [x25, #1]
    WORD $0x39400b4d // ldrb    w13, [x26, #2]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_1885
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x39000b2d // strb    w13, [x25, #2]
    WORD $0x8b1802ed // add    x13, x23, x24
    WORD $0x39400e39 // ldrb    w25, [x17, #3]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1886
    WORD $0x39000db9 // strb    w25, [x13, #3]
    WORD $0x39401239 // ldrb    w25, [x17, #4]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1887
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x390011b9 // strb    w25, [x13, #4]
    WORD $0x8b1802ed // add    x13, x23, x24
    WORD $0x39401639 // ldrb    w25, [x17, #5]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1888
    WORD $0x390015b9 // strb    w25, [x13, #5]
    WORD $0x39401a39 // ldrb    w25, [x17, #6]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1889
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x390019b9 // strb    w25, [x13, #6]
    WORD $0x8b1802ed // add    x13, x23, x24
    WORD $0x39401e39 // ldrb    w25, [x17, #7]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_1890
    WORD $0x39001db9 // strb    w25, [x13, #7]
    WORD $0x91002318 // add    x24, x24, #8
    WORD $0x3940222d // ldrb    w13, [x17, #8]
    WORD $0x710171bf // cmp    w13, #92
    BNE LBB5_1875
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x8b1802f9 // add    x25, x23, x24
    WORD $0xd100063a // sub    x26, x17, #1
    B LBB5_1891
LBB5_1884:
    WORD $0x91000751 // add    x17, x26, #1
    WORD $0x91000739 // add    x25, x25, #1
    B LBB5_1891
LBB5_1885:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x8b1802f6 // add    x22, x23, x24
    WORD $0x910005ba // add    x26, x13, #1
    WORD $0x910009b1 // add    x17, x13, #2
    WORD $0x91000ad9 // add    x25, x22, #2
    B LBB5_1891
LBB5_1886:
    WORD $0x91000a3a // add    x26, x17, #2
    WORD $0x91000e31 // add    x17, x17, #3
    WORD $0x91000db9 // add    x25, x13, #3
    B LBB5_1891
LBB5_1887:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x8b1802f6 // add    x22, x23, x24
    WORD $0x91000dba // add    x26, x13, #3
    WORD $0x910011b1 // add    x17, x13, #4
    WORD $0x910012d9 // add    x25, x22, #4
    B LBB5_1891
LBB5_1888:
    WORD $0x9100123a // add    x26, x17, #4
    WORD $0x91001631 // add    x17, x17, #5
    WORD $0x910015b9 // add    x25, x13, #5
    B LBB5_1891
LBB5_1889:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x8b1802f6 // add    x22, x23, x24
    WORD $0x910015ba // add    x26, x13, #5
    WORD $0x910019b1 // add    x17, x13, #6
    WORD $0x91001ad9 // add    x25, x22, #6
    B LBB5_1891
LBB5_1890:
    WORD $0x91001a3a // add    x26, x17, #6
    WORD $0x91001e31 // add    x17, x17, #7
    WORD $0x91001db9 // add    x25, x13, #7
LBB5_1891:
    WORD $0x9100075a // add    x26, x26, #1
    WORD $0xaa1b03fd // mov    x29, x27
    B LBB5_1837
LBB5_1892:
    WORD $0x2a1103ed // mov    w13, w17
LBB5_1893:
    WORD $0x3800172d // strb    w13, [x25], #1
    WORD $0xaa1903f7 // mov    x23, x25
    B LBB5_1860
LBB5_1894:
    WORD $0x53067e2d // lsr    w13, w17, #6
    WORD $0x52801018 // mov    w24, #128
    WORD $0x321a05ad // orr    w13, w13, #0xc0
    WORD $0x33001638 // bfxil    w24, w17, #0, #6
    WORD $0xd1000af1 // sub    x17, x23, #2
    WORD $0x381fc2ed // sturb    w13, [x23, #-4]
    WORD $0x381fd2f8 // sturb    w24, [x23, #-3]
    B LBB5_1859
LBB5_1895:
    WORD $0xd10006ed // sub    x13, x23, #1
    WORD $0x781fc2e7 // sturh    w7, [x23, #-4]
    WORD $0x381fe2f3 // sturb    w19, [x23, #-2]
    WORD $0xaa0d03f7 // mov    x23, x13
    B LBB5_1860
LBB5_1896:
    WORD $0x0b1129ad // add    w13, w13, w17, lsl #10
    WORD $0x52801019 // mov    w25, #128
    WORD $0x0b1401ad // add    w13, w13, w20
    WORD $0x5280101a // mov    w26, #128
    WORD $0x53127db1 // lsr    w17, w13, #18
    WORD $0x5280101b // mov    w27, #128
    WORD $0x321c0e31 // orr    w17, w17, #0xf0
    WORD $0x3300171b // bfxil    w27, w24, #0, #6
    WORD $0x330c45b9 // bfxil    w25, w13, #12, #6
    WORD $0x33062dba // bfxil    w26, w13, #6, #6
    WORD $0x381fc2f1 // sturb    w17, [x23, #-4]
    WORD $0x381fd2f9 // sturb    w25, [x23, #-3]
    WORD $0x381fe2fa // sturb    w26, [x23, #-2]
    WORD $0x381ff2fb // sturb    w27, [x23, #-1]
    B LBB5_1860
LBB5_1897:
    WORD $0xaa0403ef // mov    x15, x4
    B LBB5_341
LBB5_1898:
    WORD $0xaa0503ef // mov    x15, x5
    WORD $0x52800062 // mov    w2, #3
    WORD $0x38402de7 // ldrb    w7, [x15, #2]!
    WORD $0x5100c0ed // sub    w13, w7, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_2380
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x7100c0ff // cmp    w7, #48
    BNE LBB5_1901
LBB5_1900:
    WORD $0x8b0d00ae // add    x14, x5, x13
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x39400dc7 // ldrb    w7, [x14, #3]
    WORD $0x7100c0ff // cmp    w7, #48
    BEQ LBB5_1900
LBB5_1901:
    WORD $0x710114ff // cmp    w7, #69
    BEQ LBB5_1951
    WORD $0x710194ff // cmp    w7, #101
    BEQ LBB5_1951
    WORD $0x8b0d00ae // add    x14, x5, x13
    WORD $0xaa1f03f1 // mov    x17, xzr
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0x910009d3 // add    x19, x14, #2
    WORD $0x4b0d03e3 // neg    w3, w13
LBB5_1904:
    WORD $0x5280022d // mov    w13, #17
    WORD $0xcb1101af // sub    x15, x13, x17
    WORD $0xf10005ff // cmp    x15, #1
    BLT LBB5_1919
    WORD $0x4b1101a2 // sub    w2, w13, w17
    WORD $0x5280024d // mov    w13, #18
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x8b0f026f // add    x15, x19, x15
    WORD $0xcb1101ad // sub    x13, x13, x17
    WORD $0x52800151 // mov    w17, #10
LBB5_1906:
    WORD $0x39400267 // ldrb    w7, [x19]
    WORD $0x5100c0ee // sub    w14, w7, #48
    WORD $0x710025df // cmp    w14, #9
    BHI LBB5_1952
    WORD $0x9b111c0e // madd    x14, x0, x17, x7
    WORD $0xd1000694 // sub    x20, x20, #1
    WORD $0x91000673 // add    x19, x19, #1
    WORD $0xd100c1c0 // sub    x0, x14, #48
    WORD $0x8b1401ae // add    x14, x13, x20
    WORD $0xf10005df // cmp    x14, #1
    BGT LBB5_1906
    WORD $0x394001e7 // ldrb    w7, [x15]
    B LBB5_1954
LBB5_1909:
    WORD $0x710114ff // cmp    w7, #69
    BEQ LBB5_1729
    WORD $0x710194ff // cmp    w7, #101
    BEQ LBB5_1729
    CMP $0, R3
    BNE LBB5_2447
    CMP $0, R18_PLATFORM
    BEQ LBB5_2026
    WORD $0xb24107ec // mov    x12, #-9223372036854775807
    WORD $0xeb0c001f // cmp    x0, x12
    BLO LBB5_2097
    WORD $0x9e630003 // ucvtf    d3, x0
    WORD $0x2a1f03e1 // mov    w1, wzr
LBB5_1915:
    WORD $0x9e66006c // fmov    x12, d3
    WORD $0xd2410180 // eor    x0, x12, #0x8000000000000000
    B LBB5_2384
LBB5_1916:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa1f03fd // mov    x29, xzr
    TST $(1<<0), R18_PLATFORM
    BEQ LBB5_2383
    B LBB5_2385
LBB5_1917:
    WORD $0xcb0d03e2 // neg    x2, x13
LBB5_1918:
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0xaa0203ed // mov    x13, x2
    TST $(1<<63), R2
    BEQ LBB5_2014
    B LBB5_2013
LBB5_1919:
    WORD $0x2a1f03e2 // mov    w2, wzr
    B LBB5_1953
LBB5_1920:
LBB5_1921:
    WORD $0x7101f45f // cmp    w2, #125
    BNE LBB5_837
    WORD $0xb940cd0b // ldr    w11, [x8, #204]
    WORD $0x8b01100d // add    x13, x0, x1, lsl #4
    WORD $0x1100056b // add    w11, w11, #1
    WORD $0xb900cd0b // str    w11, [x8, #204]
    WORD $0xcb0f020b // sub    x11, x16, x15
    WORD $0xf94005ae // ldr    x14, [x13, #8]
    WORD $0xd344fd6b // lsr    x11, x11, #4
    WORD $0xf900550e // str    x14, [x8, #168]
    WORD $0xb9000deb // str    w11, [x15, #12]
    WORD $0xf94001eb // ldr    x11, [x15]
    WORD $0xb90009bf // str    wzr, [x13, #8]
    WORD $0xf940590d // ldr    x13, [x8, #176]
    WORD $0xb940e50e // ldr    w14, [x8, #228]
    WORD $0x92609d6b // and    x11, x11, #0xffffffff000000ff
    WORD $0xeb0e01bf // cmp    x13, x14
    WORD $0xf90001eb // str    x11, [x15]
    BHI LBB5_1701
LBB5_1923:
    WORD $0xf9405510 // ldr    x16, [x8, #168]
    WORD $0xd10005ad // sub    x13, x13, #1
    WORD $0xaa1f03eb // mov    x11, xzr
    WORD $0xb100061f // cmn    x16, #1
    WORD $0xf900590d // str    x13, [x8, #176]
    BEQ LBB5_1925
LBB5_1924:
    WORD $0xf940024b // ldr    x11, [x18]
    WORD $0x8b10116b // add    x11, x11, x16, lsl #4
LBB5_1925:
    CMP $0, R11
    BEQ LBB5_835
LBB5_1926:
    WORD $0xaa1103ef // mov    x15, x17
    WORD $0x384015e0 // ldrb    w0, [x15], #1
    WORD $0x7100801f // cmp    w0, #32
    BHI LBB5_1937
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ac021ad // lsl    x13, x13, x0
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1937
    WORD $0x39400620 // ldrb    w0, [x17, #1]
    WORD $0x91000a2f // add    x15, x17, #2
    WORD $0x7100801f // cmp    w0, #32
    BHI LBB5_1950
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ac021ad // lsl    x13, x13, x0
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1950
    WORD $0xf9404910 // ldr    x16, [x8, #144]
    WORD $0xcb1001ed // sub    x13, x15, x16
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_1933
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x9acd21ed // lsl    x13, x15, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_1936
    WORD $0x9101020f // add    x15, x16, #64
LBB5_1933:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R17
    WORD $0xd10101f0 // sub    x16, x15, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00223 // ldr    q3, [x17, :lo12:.LCPI5_2]
LBB5_1934:
    WORD $0xadc21604 // ldp    q4, q5, [x16, #64]!
    WORD $0x4e211c90 // and    v16.16b, v4.16b, v1.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411e06 // ldp    q6, q7, [x16, #32]
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
    WORD $0x1e2600af // fmov    w15, s5
    WORD $0x1e2600d1 // fmov    w17, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa0f81cd // orr    x13, x14, x15, lsl #32
    WORD $0xaa11c1ad // orr    x13, x13, x17, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_1934
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa9093510 // stp    x16, x13, [x8, #144]
LBB5_1936:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d020f // add    x15, x16, x13
    WORD $0x384015e0 // ldrb    w0, [x15], #1
LBB5_1937:
    WORD $0x3940016d // ldrb    w13, [x11]
    WORD $0xaa0f03f1 // mov    x17, x15
    WORD $0x2a0003f2 // mov    w18, w0
    WORD $0xf10019bf // cmp    x13, #6
    BNE LBB5_2400
LBB5_1938:
    WORD $0xf940016d // ldr    x13, [x11]
    WORD $0x12001e4e // and    w14, w18, #0xff
    WORD $0x7100b1df // cmp    w14, #44
    WORD $0x910401ad // add    x13, x13, #256
    WORD $0xf900016d // str    x13, [x11]
    BNE LBB5_2251
    WORD $0x38401622 // ldrb    w2, [x17], #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_836
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ac221ad // lsl    x13, x13, x2
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_836
    WORD $0x38401622 // ldrb    w2, [x17], #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_836
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ac221ad // lsl    x13, x13, x2
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_836
    WORD $0xf940490e // ldr    x14, [x8, #144]
    WORD $0xcb0e022d // sub    x13, x17, x14
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_1946
    WORD $0xf9404d0f // ldr    x15, [x8, #152]
    WORD $0x92800010 // mov    x16, #-1
    WORD $0x9acd220d // lsl    x13, x16, x13
    WORD $0xea0d01ed // ands    x13, x15, x13
    BNE LBB5_1949
    WORD $0x910101d1 // add    x17, x14, #64
LBB5_1946:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R15
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0xd101022e // sub    x14, x17, #64
    WORD $0x3dc001e3 // ldr    q3, [x15, :lo12:.LCPI5_2]
LBB5_1947:
    WORD $0xadc215c4 // ldp    q4, q5, [x14, #64]!
    WORD $0x4e211c90 // and    v16.16b, v4.16b, v1.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411dc6 // ldp    q6, q7, [x14, #32]
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
    WORD $0x1e26008f // fmov    w15, s4
    WORD $0x1e2600b0 // fmov    w16, s5
    WORD $0x1e2600d1 // fmov    w17, s6
    WORD $0x33103daf // bfi    w15, w13, #16, #16
    WORD $0xaa1081ed // orr    x13, x15, x16, lsl #32
    WORD $0xaa11c1ad // orr    x13, x13, x17, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_1947
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa909350e // stp    x14, x13, [x8, #144]
LBB5_1949:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d01d1 // add    x17, x14, x13
    WORD $0x38401622 // ldrb    w2, [x17], #1
    B LBB5_837
LBB5_1950:
    WORD $0x3940016d // ldrb    w13, [x11]
    WORD $0xaa0f03f1 // mov    x17, x15
    WORD $0x2a0003f2 // mov    w18, w0
    WORD $0xf10019bf // cmp    x13, #6
    BEQ LBB5_1938
    B LBB5_2400
LBB5_1951:
    WORD $0x8b0d00ad // add    x13, x5, x13
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0x910009af // add    x15, x13, #2
    B LBB5_1729
LBB5_1952:
    WORD $0x4b1403e2 // neg    w2, w20
LBB5_1953:
    WORD $0xaa1303ef // mov    x15, x19
LBB5_1954:
    WORD $0x4b020063 // sub    w3, w3, w2
    WORD $0x5100c0ed // sub    w13, w7, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1957
LBB5_1955:
    WORD $0x38401de7 // ldrb    w7, [x15, #1]!
    WORD $0x5100c0ed // sub    w13, w7, #48
    WORD $0x710029bf // cmp    w13, #10
    BLO LBB5_1955
    WORD $0x52800021 // mov    w1, #1
LBB5_1957:
    WORD $0x52801bed // mov    w13, #223
    WORD $0x0a0d00ed // and    w13, w7, w13
    WORD $0x710115bf // cmp    w13, #69
    BEQ LBB5_1729
LBB5_1958:
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x1280000d // mov    w13, #-1
    WORD $0xd374fc0e // lsr    x14, x0, #52
    WORD $0x5a8d15b1 // cneg    w17, w13, eq
    CMP $0, R14
    BNE LBB5_1969
    WORD $0x9e630003 // ucvtf    d3, x0
    WORD $0x531f7e2d // lsr    w13, w17, #31
    WORD $0x9e66006e // fmov    x14, d3
    WORD $0xaa0dfdcd // orr    x13, x14, x13, lsl #63
    WORD $0x9e6701a3 // fmov    d3, x13
    CMP $0, R3
    BEQ LBB5_2379
    CMP $0, R0
    BEQ LBB5_2379
    WORD $0x5100046d // sub    w13, w3, #1
    WORD $0x710091bf // cmp    w13, #36
    BHI LBB5_1967
    WORD $0x2a0303ed // mov    w13, w3
    WORD $0x71005c7f // cmp    w3, #23
    BLO LBB5_1964
    WORD $0x5100586d // sub    w13, w3, #22
    ADR P10_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:P10_TAB
    WORD $0xfc6d59c4 // ldr    d4, [x14, w13, uxtw #3]
    WORD $0x528002cd // mov    w13, #22
    WORD $0x1e630883 // fmul    d3, d4, d3
LBB5_1964:
    ADR LCPI5_3, R14
    WORD $0xfd4001c4 // ldr    d4, [x14, :lo12:.LCPI5_3]
    WORD $0x1e642060 // fcmp    d3, d4
    BGT LBB5_1970
    ADR LCPI5_4, R14
    WORD $0xfd4001c4 // ldr    d4, [x14, :lo12:.LCPI5_4]
    WORD $0x1e642060 // fcmp    d3, d4
    BMI LBB5_1970
    ADR P10_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:P10_TAB
    WORD $0xfc6d59c4 // ldr    d4, [x14, w13, uxtw #3]
    B LBB5_2378
LBB5_1967:
    WORD $0x3100587f // cmn    w3, #22
    BLO LBB5_1969
    WORD $0x4b0303ed // neg    w13, w3
    ADR P10_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:P10_TAB
    WORD $0xfc6d59c4 // ldr    d4, [x14, w13, uxtw #3]
    WORD $0x1e641863 // fdiv    d3, d3, d4
    B LBB5_2379
LBB5_1969:
    WORD $0x5105706d // sub    w13, w3, #348
    WORD $0x310ae1bf // cmn    w13, #696
    BLO LBB5_1978
LBB5_1970:
    WORD $0x11057062 // add    w2, w3, #348
    ADR POW10_M128_TAB, R7
    WORD $0x910000e7 // add    x7, x7, :lo12:POW10_M128_TAB
    WORD $0x528a4d4d // mov    w13, #21098
    WORD $0x8b2250ee // add    x14, x7, w2, uxtw #4
    WORD $0x72a0006d // movk    w13, #3, lsl #16
    WORD $0xdac01014 // clz    x20, x0
    WORD $0x1b0d7c73 // mul    w19, w3, w13
    WORD $0xf94005c3 // ldr    x3, [x14, #8]
    WORD $0x9ad4200d // lsl    x13, x0, x20
    WORD $0x13107e6e // asr    w14, w19, #16
    WORD $0x1110fdce // add    w14, w14, #1087
    WORD $0xaa2d03f8 // mvn    x24, x13
    WORD $0x9bcd7c75 // umulh    x21, x3, x13
    WORD $0x93407dd3 // sxtw    x19, w14
    WORD $0x9b0d7c76 // mul    x22, x3, x13
    WORD $0xcb140274 // sub    x20, x19, x20
    WORD $0x924022b7 // and    x23, x21, #0x1ff
    WORD $0xeb1802df // cmp    x22, x24
    BLS LBB5_1975
    WORD $0xf107feff // cmp    x23, #511
    BNE LBB5_1975
    WORD $0xd37cec4e // lsl    x14, x2, #4
    WORD $0xf86e68ee // ldr    x14, [x7, x14]
    WORD $0x9bcd7dd7 // umulh    x23, x14, x13
    WORD $0x9b0d7dcd // mul    x13, x14, x13
    WORD $0xab1602f6 // adds    x22, x23, x22
    WORD $0x9a9536b5 // cinc    x21, x21, hs
    WORD $0xeb1801bf // cmp    x13, x24
    WORD $0x924022b7 // and    x23, x21, #0x1ff
    BLS LBB5_1975
    WORD $0xb10006df // cmn    x22, #1
    BNE LBB5_1975
    WORD $0xf107feff // cmp    x23, #511
    BEQ LBB5_1978
LBB5_1975:
    WORD $0xd37ffead // lsr    x13, x21, #63
    WORD $0xaa1702d6 // orr    x22, x22, x23
    WORD $0x910025ae // add    x14, x13, #9
    WORD $0x9ace26b5 // lsr    x21, x21, x14
    CMP $0, R22
    BNE LBB5_1977
    WORD $0x924006ae // and    x14, x21, #0x3
    WORD $0xf10005df // cmp    x14, #1
    BEQ LBB5_1978
LBB5_1977:
    WORD $0x924002ae // and    x14, x21, #0x1
    WORD $0x8b0d028d // add    x13, x20, x13
    WORD $0x8b1501d5 // add    x21, x14, x21
    WORD $0xd376feb6 // lsr    x22, x21, #54
    WORD $0xf10002df // cmp    x22, #0
    WORD $0x1a9f17ee // cset    w14, eq
    WORD $0xcb0e01ad // sub    x13, x13, x14
    WORD $0xd11ffdae // sub    x14, x13, #2047
    WORD $0xb11ff9df // cmn    x14, #2046
    BHS LBB5_2015
LBB5_1978:
    WORD $0xf9402902 // ldr    x2, [x8, #80]
    WORD $0xcb0501e7 // sub    x7, x15, x5
    WORD $0xf9402112 // ldr    x18, [x8, #64]
    CMP $0, R2
    BEQ LBB5_1990
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0x0460e3e1 // cnth    x1
    WORD $0xeb01005f // cmp    x2, x1
    BLO LBB5_1988
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0x04bf504d // rdvl    x13, #2
    WORD $0xeb0d005f // cmp    x2, x13
    BHS LBB5_1984
LBB5_1981:
    WORD $0xcb0103ee // neg    x14, x1
    WORD $0xaa0003ed // mov    x13, x0
    WORD $0x8a0e0040 // and    x0, x2, x14
    WORD $0x2558e3e0 // ptrue    p0.h
    WORD $0x2578c003 // mov    z3.h, #0
LBB5_1982:
    WORD $0xe42d4243 // st1b    { z3.h }, p0, [x18, x13]
    WORD $0x8b0101ad // add    x13, x13, x1
    WORD $0xeb0d001f // cmp    x0, x13
    BNE LBB5_1982
    WORD $0xeb00005f // cmp    x2, x0
    BNE LBB5_1988
    B LBB5_1990
LBB5_1984:
    WORD $0x04bf57ce // rdvl    x14, #-2
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x8a0e0040 // and    x0, x2, x14
    WORD $0x04bf5043 // rdvl    x3, #2
    WORD $0x04325033 // addvl    x19, x18, #1
    WORD $0x2518e3e0 // ptrue    p0.b
    WORD $0x2538c003 // mov    z3.b, #0
LBB5_1985:
    WORD $0xe40d4243 // st1b    { z3.b }, p0, [x18, x13]
    WORD $0xe40d4263 // st1b    { z3.b }, p0, [x19, x13]
    WORD $0x8b0301ad // add    x13, x13, x3
    WORD $0xeb0d001f // cmp    x0, x13
    BNE LBB5_1985
    WORD $0xeb00004d // subs    x13, x2, x0
    BEQ LBB5_1990
    WORD $0xeb0101bf // cmp    x13, x1
    BHS LBB5_1981
LBB5_1988:
    WORD $0xcb00004d // sub    x13, x2, x0
    WORD $0x8b000240 // add    x0, x18, x0
LBB5_1989:
    WORD $0xf10005ad // subs    x13, x13, #1
    WORD $0x3800141f // strb    wzr, [x0], #1
    BNE LBB5_1989
LBB5_1990:
    WORD $0x394000a0 // ldrb    w0, [x5]
    WORD $0xaa1f03f5 // mov    x21, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0xb9001ff9 // str    w25, [sp, #28]
    WORD $0x7100b41f // cmp    w0, #45
    WORD $0x1a9f17f3 // cset    w19, eq
    WORD $0xeb1300ff // cmp    x7, x19
    BLE LBB5_2377
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x2a1f03f5 // mov    w21, wzr
    WORD $0x2a1f03f6 // mov    w22, wzr
    WORD $0x2a1f03f7 // mov    w23, wzr
    WORD $0x2a1f03f8 // mov    w24, wzr
    WORD $0x2a1f03f4 // mov    w20, wzr
    WORD $0x5280003a // mov    w26, #1
    B LBB5_1994
LBB5_1992:
    WORD $0x11000739 // add    w25, w25, #1
    WORD $0x382d6a5b // strb    w27, [x18, x13]
    WORD $0x2a1903f5 // mov    w21, w25
    WORD $0x2a1903f6 // mov    w22, w25
    WORD $0x2a1903f7 // mov    w23, w25
    WORD $0x2a1903f8 // mov    w24, w25
LBB5_1993:
    WORD $0x91000673 // add    x19, x19, #1
    WORD $0xeb07027f // cmp    x19, x7
    WORD $0x1a9fa7fa // cset    w26, lt
    WORD $0xeb1300ff // cmp    x7, x19
    BEQ LBB5_2004
LBB5_1994:
    WORD $0x387368bb // ldrb    w27, [x5, x19]
    WORD $0x5100c36d // sub    w13, w27, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1999
    WORD $0x7100c37f // cmp    w27, #48
    BNE LBB5_2001
    CMP $0, R22
    BEQ LBB5_2003
    WORD $0x93407ead // sxtw    x13, w21
    WORD $0xeb0d005f // cmp    x2, x13
    BHI LBB5_1992
    WORD $0x2a1503f6 // mov    w22, w21
    WORD $0x2a1503f7 // mov    w23, w21
    WORD $0x2a1503f8 // mov    w24, w21
    B LBB5_1993
LBB5_1999:
    WORD $0x7100bb7f // cmp    w27, #46
    BNE LBB5_2005
    WORD $0x52800034 // mov    w20, #1
    WORD $0x2a1803e3 // mov    w3, w24
    B LBB5_1993
LBB5_2001:
    WORD $0x93407eed // sxtw    x13, w23
    WORD $0xeb0d005f // cmp    x2, x13
    BHI LBB5_1992
    WORD $0x52800021 // mov    w1, #1
    WORD $0x2a1703f8 // mov    w24, w23
    B LBB5_1993
LBB5_2003:
    WORD $0x2a1f03f7 // mov    w23, wzr
    WORD $0x2a1f03f8 // mov    w24, wzr
    WORD $0x51000463 // sub    w3, w3, #1
    B LBB5_1993
LBB5_2004:
    WORD $0xaa0703f3 // mov    x19, x7
LBB5_2005:
    WORD $0x7100029f // cmp    w20, #0
    WORD $0x1a8302a3 // csel    w3, w21, w3, eq
    TST $(1<<0), R26
    BEQ LBB5_2037
    WORD $0x387368ad // ldrb    w13, [x5, x19]
    WORD $0x321b01ad // orr    w13, w13, #0x20
    WORD $0x710195bf // cmp    w13, #101
    BNE LBB5_2037
    WORD $0x2a1303ed // mov    w13, w19
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x386d68b5 // ldrb    w21, [x5, x13]
    WORD $0x7100b6bf // cmp    w21, #45
    BEQ LBB5_2029
    WORD $0x52800034 // mov    w20, #1
    WORD $0x7100aebf // cmp    w21, #43
    BNE LBB5_2030
    WORD $0x11000a6d // add    w13, w19, #2
    B LBB5_2030
LBB5_2010:
    WORD $0xaa0001ee // orr    x14, x15, x0
    WORD $0xd37ffdd1 // lsr    x17, x14, #63
    WORD $0x52000231 // eor    w17, w17, #0x1
    TST $(1<<63), R14
    BNE LBB5_2027
    WORD $0xeb0001ff // cmp    x15, x0
    BLT LBB5_2027
    WORD $0xaa2f03e2 // mvn    x2, x15
LBB5_2013:
    WORD $0xaa2203e2 // mvn    x2, x2
    WORD $0x52800061 // mov    w1, #3
    WORD $0x9280004d // mov    x13, #-3
LBB5_2014:
    WORD $0x8b3241ad // add    x13, x13, w18, uxtw
    WORD $0xf940510e // ldr    x14, [x8, #160]
    WORD $0x8b0200b1 // add    x17, x5, x2
    WORD $0xf90005cd // str    x13, [x14, #8]
    WORD $0x5280036d // mov    w13, #27
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0xaa1081ad // orr    x13, x13, x16, lsl #32
    WORD $0xb940d90f // ldr    w15, [x8, #216]
    WORD $0x91004250 // add    x16, x18, #16
    WORD $0xf90001cd // str    x13, [x14]
    WORD $0x110005ef // add    w15, w15, #1
    WORD $0xf9005110 // str    x16, [x8, #160]
    WORD $0xb900d90f // str    w15, [x8, #216]
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a8101a2 // csel    w2, w13, w1, eq
    CMP $0, R1
    BEQ LBB5_2388
    B LBB5_2444
LBB5_2015:
    WORD $0xf10002df // cmp    x22, #0
    WORD $0x5280002e // mov    w14, #1
    WORD $0x9a8e05ce // cinc    x14, x14, ne
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x9ace26ae // lsr    x14, x21, x14
    WORD $0xb34c2dae // bfi    x14, x13, #52, #12
    WORD $0xb24101cd // orr    x13, x14, #0x8000000000000000
    WORD $0x9a8e11ad // csel    x13, x13, x14, ne
    WORD $0x9e6701a3 // fmov    d3, x13
    CMP $0, R1
    BEQ LBB5_2379
    WORD $0x9100040d // add    x13, x0, #1
    WORD $0xdac011ae // clz    x14, x13
    WORD $0xcb0e0260 // sub    x0, x19, x14
    WORD $0x9ace21b4 // lsl    x20, x13, x14
    WORD $0xaa3403ed // mvn    x13, x20
    WORD $0x9bc37e81 // umulh    x1, x20, x3
    WORD $0x9b037e83 // mul    x3, x20, x3
    WORD $0x92402033 // and    x19, x1, #0x1ff
    WORD $0xeb0d007f // cmp    x3, x13
    BLS LBB5_2021
    WORD $0xf107fe7f // cmp    x19, #511
    BNE LBB5_2021
    WORD $0xd37cec4e // lsl    x14, x2, #4
    WORD $0xf86e68ee // ldr    x14, [x7, x14]
    WORD $0x9bd47dc2 // umulh    x2, x14, x20
    WORD $0x9b147dce // mul    x14, x14, x20
    WORD $0xab030043 // adds    x3, x2, x3
    WORD $0x9a813421 // cinc    x1, x1, hs
    WORD $0xeb0d01df // cmp    x14, x13
    WORD $0x92402033 // and    x19, x1, #0x1ff
    BLS LBB5_2021
    WORD $0xb100047f // cmn    x3, #1
    BNE LBB5_2021
    WORD $0xf107fe7f // cmp    x19, #511
    BEQ LBB5_1978
LBB5_2021:
    WORD $0xd37ffc2d // lsr    x13, x1, #63
    WORD $0xaa130062 // orr    x2, x3, x19
    WORD $0x910025ae // add    x14, x13, #9
    WORD $0x9ace2421 // lsr    x1, x1, x14
    CMP $0, R2
    BNE LBB5_2023
    WORD $0x9240042e // and    x14, x1, #0x3
    WORD $0xf10005df // cmp    x14, #1
    BEQ LBB5_1978
LBB5_2023:
    WORD $0x9240002e // and    x14, x1, #0x1
    WORD $0x8b0d000d // add    x13, x0, x13
    WORD $0x8b0101c1 // add    x1, x14, x1
    WORD $0xd376fc22 // lsr    x2, x1, #54
    WORD $0xf100005f // cmp    x2, #0
    WORD $0x1a9f17ee // cset    w14, eq
    WORD $0xcb0e01ad // sub    x13, x13, x14
    WORD $0xd11ffdae // sub    x14, x13, #2047
    WORD $0xb11ff9df // cmn    x14, #2046
    BLO LBB5_1978
    WORD $0xf100005f // cmp    x2, #0
    WORD $0x5280002e // mov    w14, #1
    WORD $0x9a8e05ce // cinc    x14, x14, ne
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x9ace242e // lsr    x14, x1, x14
    WORD $0xb34c2dae // bfi    x14, x13, #52, #12
    WORD $0xb24101cd // orr    x13, x14, #0x8000000000000000
    WORD $0x9a8e11ad // csel    x13, x13, x14, ne
    WORD $0x9e6701a4 // fmov    d4, x13
    WORD $0x1e642060 // fcmp    d3, d4
    BEQ LBB5_2379
    B LBB5_1978
LBB5_2025:
    WORD $0xaa0403e3 // mov    x3, x4
    WORD $0x92800002 // mov    x2, #-1
    CMP $0, R15
    BNE LBB5_1807
    B LBB5_2013
LBB5_2026:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa0003fd // mov    x29, x0
    B LBB5_2383
LBB5_2027:
    WORD $0xd100040e // sub    x14, x0, #1
    WORD $0xeb0e01ff // cmp    x15, x14
    WORD $0x1a9f17ee // cset    w14, eq
    WORD $0x6a0e023f // tst    w17, w14
    WORD $0xda8001a2 // csinv    x2, x13, x0, eq
    B LBB5_1918
LBB5_2028:
    WORD $0x5280002d // mov    w13, #1
    B LBB5_2014
LBB5_2029:
    WORD $0x11000a6d // add    w13, w19, #2
    WORD $0x12800014 // mov    w20, #-1
LBB5_2030:
    WORD $0x2a1f03f3 // mov    w19, wzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0xeb0d00ff // cmp    x7, x13
    BLE LBB5_2036
    WORD $0xcb0d01ee // sub    x14, x15, x13
    WORD $0x8b0601a6 // add    x6, x13, x6
    WORD $0x2a1f03f3 // mov    w19, wzr
    WORD $0xcb0501cd // sub    x13, x14, x5
    WORD $0x8b060084 // add    x4, x4, x6
    WORD $0x5284e1e5 // mov    w5, #9999
    WORD $0x52800146 // mov    w6, #10
LBB5_2032:
    WORD $0x38401487 // ldrb    w7, [x4], #1
    WORD $0x7100c0ff // cmp    w7, #48
    BLO LBB5_2036
    WORD $0x7100e4ff // cmp    w7, #57
    BHI LBB5_2036
    WORD $0x6b05027f // cmp    w19, w5
    BGT LBB5_2036
    WORD $0x1b061e6e // madd    w14, w19, w6, w7
    WORD $0xf10005ad // subs    x13, x13, #1
    WORD $0x5100c1d3 // sub    w19, w14, #48
    BNE LBB5_2032
LBB5_2036:
    WORD $0x1b140e63 // madd    w3, w19, w20, w3
LBB5_2037:
    WORD $0xaa1f03f5 // mov    x21, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    CMP $0, R25
    BEQ LBB5_2377
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0xd2effe15 // mov    x21, #9218868437227405312
    WORD $0x7104d87f // cmp    w3, #310
    BGT LBB5_2377
    WORD $0xaa1f03f5 // mov    x21, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x3105287f // cmn    w3, #330
    BLT LBB5_2377
    WORD $0x7100047f // cmp    w3, #1
    BLT LBB5_2101
    WORD $0xb201e7f3 // mov    x19, #-7378697629483820647
    WORD $0x2a1f03e5 // mov    w5, wzr
    WORD $0xf2933353 // movk    x19, #39322
    WORD $0x92800004 // mov    x4, #-1
    WORD $0x52800147 // mov    w7, #10
    WORD $0xf2e03333 // movk    x19, #409, lsl #48
    WORD $0x2a1903fb // mov    w27, w25
