    CMP $0, R21
    BEQ LBB5_544
    WORD $0x110ff640 // add    w0, w18, #1021
    WORD $0x3110ea5f // cmn    w18, #1082
    BHI LBB5_531
    WORD $0xb201e7f2 // mov    x18, #-7378697629483820647
    WORD $0x52800141 // mov    w1, #10
    WORD $0xf2933352 // movk    x18, #39322
    WORD $0xf2e03332 // movk    x18, #409, lsl #48
    B LBB5_509
LBB5_506:
    WORD $0x2a1f03f1 // mov    w17, wzr
LBB5_507:
    WORD $0x2a1f03f4 // mov    w20, wzr
LBB5_508:
    WORD $0x1100f002 // add    w2, w0, #60
    WORD $0x3101e01f // cmn    w0, #120
    WORD $0x2a1403f5 // mov    w21, w20
    WORD $0x2a0203e0 // mov    w0, w2
    WORD $0x2a1403e4 // mov    w4, w20
    BGE LBB5_532
LBB5_509:
    WORD $0xaa1f03e5 // mov    x5, xzr
    WORD $0xaa1f03e4 // mov    x4, xzr
    WORD $0x0ab57ea2 // bic    w2, w21, w21, asr #31
LBB5_510:
    WORD $0xeb05005f // cmp    x2, x5
    BEQ LBB5_513
    WORD $0x386569a3 // ldrb    w3, [x13, x5]
    WORD $0x910004a5 // add    x5, x5, #1
    WORD $0x9b010c83 // madd    x3, x4, x1, x3
    WORD $0xd100c064 // sub    x4, x3, #48
    WORD $0xd37cfc83 // lsr    x3, x4, #60
    CMP $0, R3
    BEQ LBB5_510
    WORD $0xaa0403e3 // mov    x3, x4
    WORD $0x2a0503e2 // mov    w2, w5
    B LBB5_515
LBB5_513:
    CMP $0, R4
    BEQ LBB5_507
LBB5_514:
    WORD $0x8b040883 // add    x3, x4, x4, lsl #2
    WORD $0x11000442 // add    w2, w2, #1
    WORD $0xd37ff863 // lsl    x3, x3, #1
    WORD $0xeb12009f // cmp    x4, x18
    WORD $0xaa0303e4 // mov    x4, x3
    BLO LBB5_514
LBB5_515:
    WORD $0x6b15005f // cmp    w2, w21
    BGE LBB5_519
    WORD $0x2a0203e5 // mov    w5, w2
    WORD $0xaa1f03e4 // mov    x4, xzr
    WORD $0x93407ca5 // sxtw    x5, w5
    WORD $0x93407e86 // sxtw    x6, w20
    WORD $0x8b0501a7 // add    x7, x13, x5
LBB5_517:
    WORD $0xd37cfc73 // lsr    x19, x3, #60
    WORD $0x9240ec63 // and    x3, x3, #0xfffffffffffffff
    WORD $0x321c0673 // orr    w19, w19, #0x30
    WORD $0x382469b3 // strb    w19, [x13, x4]
    WORD $0x386468f3 // ldrb    w19, [x7, x4]
    WORD $0x91000484 // add    x4, x4, #1
    WORD $0x9b014c63 // madd    x3, x3, x1, x19
    WORD $0x8b0400b3 // add    x19, x5, x4
    WORD $0xeb06027f // cmp    x19, x6
    WORD $0xd100c063 // sub    x3, x3, #48
    BLT LBB5_517
    WORD $0x2a0403f4 // mov    w20, w4
    CMP $0, R3
    BNE LBB5_521
    B LBB5_523
LBB5_519:
    WORD $0x2a1f03f4 // mov    w20, wzr
    B LBB5_521
LBB5_520:
    WORD $0xd37cfc65 // lsr    x5, x3, #60
    WORD $0x11000694 // add    w20, w20, #1
    WORD $0x321c04a5 // orr    w5, w5, #0x30
    WORD $0x382469a5 // strb    w5, [x13, x4]
    WORD $0x9240ec64 // and    x4, x3, #0xfffffffffffffff
    WORD $0x8b040883 // add    x3, x4, x4, lsl #2
    WORD $0xd37ff863 // lsl    x3, x3, #1
    CMP $0, R4
    BEQ LBB5_523
LBB5_521:
    WORD $0x93407e84 // sxtw    x4, w20
    WORD $0xeb04021f // cmp    x16, x4
    BHI LBB5_520
    WORD $0xd37cfc64 // lsr    x4, x3, #60
    WORD $0xf100009f // cmp    x4, #0
    WORD $0x1a9f05ce // csinc    w14, w14, wzr, eq
    WORD $0x9240ec64 // and    x4, x3, #0xfffffffffffffff
    WORD $0x8b040883 // add    x3, x4, x4, lsl #2
    WORD $0xd37ff863 // lsl    x3, x3, #1
    CMP $0, R4
    BNE LBB5_521
LBB5_523:
    WORD $0x4b020231 // sub    w17, w17, w2
    WORD $0x7100069f // cmp    w20, #1
    WORD $0x11000631 // add    w17, w17, #1
    BLT LBB5_528
    WORD $0x2a1403e2 // mov    w2, w20
    WORD $0x8b0d0043 // add    x3, x2, x13
    WORD $0x385ff063 // ldurb    w3, [x3, #-1]
    WORD $0x7100c07f // cmp    w3, #48
    BNE LBB5_508
LBB5_525:
    WORD $0xf1000454 // subs    x20, x2, #1
    BLS LBB5_506
    WORD $0x51000842 // sub    w2, w2, #2
    WORD $0x386249a3 // ldrb    w3, [x13, w2, uxtw]
    WORD $0xaa1403e2 // mov    x2, x20
    WORD $0x7100c07f // cmp    w3, #48
    BEQ LBB5_525
    B LBB5_508
LBB5_528:
    CMP $0, R20
    BNE LBB5_508
    B LBB5_506
LBB5_529:
    WORD $0x7110025f // cmp    w18, #1024
    BGT LBB5_330
    WORD $0x51000652 // sub    w18, w18, #1
    B LBB5_578
LBB5_531:
    WORD $0x2a1503e4 // mov    w4, w21
    WORD $0x2a0003e2 // mov    w2, w0
LBB5_532:
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0xaa1f03e1 // mov    x1, xzr
    WORD $0x4b0203f2 // neg    w18, w2
    WORD $0x0aa47c83 // bic    w3, w4, w4, asr #31
    WORD $0x52800142 // mov    w2, #10
LBB5_533:
    WORD $0xeb00007f // cmp    x3, x0
    BEQ LBB5_540
    WORD $0x386069a5 // ldrb    w5, [x13, x0]
    WORD $0x91000400 // add    x0, x0, #1
    WORD $0x9b021421 // madd    x1, x1, x2, x5
    WORD $0xd100c021 // sub    x1, x1, #48
    WORD $0x9ad22425 // lsr    x5, x1, x18
    CMP $0, R5
    BEQ LBB5_533
    WORD $0x2a0003e3 // mov    w3, w0
LBB5_536:
    WORD $0x92800000 // mov    x0, #-1
    WORD $0x6b04007f // cmp    w3, w4
    WORD $0x9ad22000 // lsl    x0, x0, x18
    WORD $0xaa2003e2 // mvn    x2, x0
    BGE LBB5_546
    WORD $0x2a0303e0 // mov    w0, w3
    WORD $0xaa1f03e6 // mov    x6, xzr
    WORD $0x93407c00 // sxtw    x0, w0
    WORD $0x93407e84 // sxtw    x4, w20
    WORD $0x8b0001a5 // add    x5, x13, x0
    WORD $0x52800147 // mov    w7, #10
LBB5_538:
    WORD $0x9ad22433 // lsr    x19, x1, x18
    WORD $0x8a020021 // and    x1, x1, x2
    WORD $0x1100c273 // add    w19, w19, #48
    WORD $0x382669b3 // strb    w19, [x13, x6]
    WORD $0x386668b3 // ldrb    w19, [x5, x6]
    WORD $0x910004c6 // add    x6, x6, #1
    WORD $0x9b074c21 // madd    x1, x1, x7, x19
    WORD $0x8b060013 // add    x19, x0, x6
    WORD $0xeb04027f // cmp    x19, x4
    WORD $0xd100c021 // sub    x1, x1, #48
    BLT LBB5_538
    WORD $0x2a0603e0 // mov    w0, w6
    B LBB5_547
LBB5_540:
    CMP $0, R1
    BEQ LBB5_544
    WORD $0x9ad22420 // lsr    x0, x1, x18
    CMP $0, R0
    BEQ LBB5_543
    WORD $0x92800002 // mov    x2, #-1
    WORD $0x4b030231 // sub    w17, w17, w3
    WORD $0x9ad22042 // lsl    x2, x2, x18
    WORD $0x2a1f03e0 // mov    w0, wzr
    WORD $0xaa2203e2 // mvn    x2, x2
    WORD $0x11000631 // add    w17, w17, #1
    B LBB5_549
LBB5_543:
    WORD $0x8b010820 // add    x0, x1, x1, lsl #2
    WORD $0x11000463 // add    w3, w3, #1
    WORD $0xd37ff801 // lsl    x1, x0, #1
    WORD $0x9ad22420 // lsr    x0, x1, x18
    CMP $0, R0
    BEQ LBB5_543
    B LBB5_536
LBB5_544:
    WORD $0x2a1f03ef // mov    w15, wzr
    WORD $0x12807fb2 // mov    w18, #-1022
    B LBB5_679
LBB5_545:
    WORD $0x5280002a // mov    w10, #1
    WORD $0x9280002c // mov    x12, #-2
    B LBB5_802
LBB5_546:
    WORD $0x2a1f03e0 // mov    w0, wzr
LBB5_547:
    WORD $0x4b030231 // sub    w17, w17, w3
    WORD $0x11000631 // add    w17, w17, #1
    CMP $0, R1
    BNE LBB5_549
    B LBB5_551
LBB5_548:
    WORD $0xf100007f // cmp    x3, #0
    WORD $0x1a9f05ce // csinc    w14, w14, wzr, eq
    WORD $0x8a020023 // and    x3, x1, x2
    WORD $0x8b030861 // add    x1, x3, x3, lsl #2
    WORD $0xd37ff821 // lsl    x1, x1, #1
    CMP $0, R3
    BEQ LBB5_551
LBB5_549:
    WORD $0x9ad22423 // lsr    x3, x1, x18
    WORD $0x93407c04 // sxtw    x4, w0
    WORD $0xeb04021f // cmp    x16, x4
    BLS LBB5_548
    WORD $0x1100c063 // add    w3, w3, #48
    WORD $0x11000400 // add    w0, w0, #1
    WORD $0x382469a3 // strb    w3, [x13, x4]
    WORD $0x8a020023 // and    x3, x1, x2
    WORD $0x8b030861 // add    x1, x3, x3, lsl #2
    WORD $0xd37ff821 // lsl    x1, x1, #1
    CMP $0, R3
    BNE LBB5_549
LBB5_551:
    WORD $0x7100041f // cmp    w0, #1
    BLT LBB5_555
    WORD $0x2a0003e1 // mov    w1, w0
    WORD $0x8b0d0032 // add    x18, x1, x13
    WORD $0x385ff252 // ldurb    w18, [x18, #-1]
    WORD $0x7100c25f // cmp    w18, #48
    BNE LBB5_560
LBB5_553:
    WORD $0xaa0103f2 // mov    x18, x1
    WORD $0xf1000421 // subs    x1, x1, #1
    BLS LBB5_576
    WORD $0x51000a40 // sub    w0, w18, #2
    WORD $0x386049a0 // ldrb    w0, [x13, w0, uxtw]
    WORD $0x7100c01f // cmp    w0, #48
    BEQ LBB5_553
    B LBB5_577
LBB5_555:
    WORD $0x12807fb2 // mov    w18, #-1022
    CMP $0, R0
    BEQ LBB5_706
    WORD $0x2a0003f5 // mov    w21, w0
    B LBB5_580
LBB5_557:
    WORD $0xaa3103e9 // mvn    x9, x17
    WORD $0xcb264131 // sub    x17, x9, w6, uxtw
    B LBB5_233
LBB5_558:
    WORD $0xaa3103e9 // mvn    x9, x17
    WORD $0xcb254131 // sub    x17, x9, w5, uxtw
    B LBB5_233
LBB5_559:
    WORD $0xaa3103e9 // mvn    x9, x17
    WORD $0xcb244131 // sub    x17, x9, w4, uxtw
    B LBB5_233
LBB5_560:
    WORD $0x12807fb2 // mov    w18, #-1022
    WORD $0x2a0003f5 // mov    w21, w0
    B LBB5_580
LBB5_561:
    WORD $0xaa0803f2 // mov    x18, x8
    WORD $0xf84b8e4d // ldr    x13, [x18, #184]!
    WORD $0xf85e8240 // ldur    x0, [x18, #-24]
    WORD $0x8b0c11ab // add    x11, x13, x12, lsl #4
    WORD $0x385f000e // ldurb    w14, [x0, #-16]
    WORD $0x71001ddf // cmp    w14, #7
    BEQ LBB5_711
    WORD $0x710019df // cmp    w14, #6
    BNE LBB5_727
    WORD $0xaa1003f1 // mov    x17, x16
    WORD $0x38401622 // ldrb    w2, [x17], #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_574
    WORD $0xd284c00f // mov    x15, #9728
    WORD $0x5280002e // mov    w14, #1
    WORD $0xf2c0002f // movk    x15, #1, lsl #32
    WORD $0x9ac221ce // lsl    x14, x14, x2
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_574
    WORD $0x39400602 // ldrb    w2, [x16, #1]
    WORD $0x91000a11 // add    x17, x16, #2
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_773
    WORD $0x5280002e // mov    w14, #1
    WORD $0x9ac221ce // lsl    x14, x14, x2
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_773
    WORD $0xf940490e // ldr    x14, [x8, #144]
    WORD $0xcb0e022f // sub    x15, x17, x14
    WORD $0xf100fdff // cmp    x15, #63
    BHI LBB5_570
    WORD $0xf9404d10 // ldr    x16, [x8, #152]
    WORD $0x92800011 // mov    x17, #-1
    WORD $0x9acf222f // lsl    x15, x17, x15
    WORD $0xea0f020f // ands    x15, x16, x15
    BNE LBB5_573
    WORD $0x910101d1 // add    x17, x14, #64
LBB5_570:
    ADR LCPI5_0, R14
    ADR LCPI5_1, R15
    ADR LCPI5_2, R16
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001c0 // ldr    q0, [x14, :lo12:.LCPI5_0]
    WORD $0xd101022e // sub    x14, x17, #64
    WORD $0x3dc001e2 // ldr    q2, [x15, :lo12:.LCPI5_1]
    WORD $0x3dc00203 // ldr    q3, [x16, :lo12:.LCPI5_2]
LBB5_571:
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
    WORD $0x1e2600af // fmov    w15, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e260090 // fmov    w16, s4
    WORD $0x1e2600b1 // fmov    w17, s5
    WORD $0x1e2600c1 // fmov    w1, s6
    WORD $0x33103df0 // bfi    w16, w15, #16, #16
    WORD $0xaa11820f // orr    x15, x16, x17, lsl #32
    WORD $0xaa01c1ef // orr    x15, x15, x1, lsl #48
    WORD $0xb10005ff // cmn    x15, #1
    BEQ LBB5_571
    WORD $0xaa2f03ef // mvn    x15, x15
    WORD $0xa9093d0e // stp    x14, x15, [x8, #144]
LBB5_573:
    WORD $0xdac001ef // rbit    x15, x15
    WORD $0xdac011ef // clz    x15, x15
    WORD $0x8b0f01d1 // add    x17, x14, x15
    WORD $0x38401622 // ldrb    w2, [x17], #1
LBB5_574:
    WORD $0x7101f45f // cmp    w2, #125
    BNE LBB5_774
LBB5_575:
    WORD $0xb940cd0e // ldr    w14, [x8, #204]
    WORD $0x8b0c11ac // add    x12, x13, x12, lsl #4
    WORD $0x110005ce // add    w14, w14, #1
    WORD $0xb900cd0e // str    w14, [x8, #204]
    B LBB5_724
LBB5_576:
    WORD $0x2a1f03f1 // mov    w17, wzr
LBB5_577:
    WORD $0x51000654 // sub    w20, w18, #1
    WORD $0x12807fb2 // mov    w18, #-1022
    WORD $0x2a1403f5 // mov    w21, w20
LBB5_578:
    CMP $0, R21
    BEQ LBB5_584
    WORD $0x2a1403e0 // mov    w0, w20
LBB5_580:
    WORD $0x394001a2 // ldrb    w2, [x13]
    WORD $0x93407ea1 // sxtw    x1, w21
    WORD $0x7100c45f // cmp    w2, #49
    BNE LBB5_659
    WORD $0x7100043f // cmp    w1, #1
    BNE LBB5_585
LBB5_582:
    WORD $0x5282b182 // mov    w2, #5516
    WORD $0x8b0101ef // add    x15, x15, x1
    WORD $0x386269ef // ldrb    w15, [x15, x2]
    CMP $0, R15
    BNE LBB5_661
    WORD $0x52800202 // mov    w2, #16
    B LBB5_664
LBB5_584:
    WORD $0x2a1f03ef // mov    w15, wzr
    B LBB5_679
LBB5_585:
    WORD $0x394005a2 // ldrb    w2, [x13, #1]
    WORD $0x7100c45f // cmp    w2, #49
    BNE LBB5_659
    WORD $0x7100083f // cmp    w1, #2
    BEQ LBB5_582
    WORD $0x394009a2 // ldrb    w2, [x13, #2]
    WORD $0x7100c45f // cmp    w2, #49
    BNE LBB5_659
    WORD $0x71000c3f // cmp    w1, #3
    BEQ LBB5_582
    WORD $0x39400da2 // ldrb    w2, [x13, #3]
    WORD $0x7100c05f // cmp    w2, #48
    BNE LBB5_777
    WORD $0x7100103f // cmp    w1, #4
    BEQ LBB5_582
    WORD $0x394011a2 // ldrb    w2, [x13, #4]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100143f // cmp    w1, #5
    BEQ LBB5_582
    WORD $0x394015a2 // ldrb    w2, [x13, #5]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100183f // cmp    w1, #6
    BEQ LBB5_582
    WORD $0x394019a2 // ldrb    w2, [x13, #6]
    WORD $0x7100cc5f // cmp    w2, #51
    BNE LBB5_798
    WORD $0x71001c3f // cmp    w1, #7
    BEQ LBB5_582
    WORD $0x39401da2 // ldrb    w2, [x13, #7]
    WORD $0x7100c05f // cmp    w2, #48
    BNE LBB5_777
    WORD $0x7100203f // cmp    w1, #8
    BEQ LBB5_582
    WORD $0x394021a2 // ldrb    w2, [x13, #8]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100243f // cmp    w1, #9
    BEQ LBB5_582
    WORD $0x394025a2 // ldrb    w2, [x13, #9]
    WORD $0x7100d05f // cmp    w2, #52
    BNE LBB5_823
    WORD $0x7100283f // cmp    w1, #10
    BEQ LBB5_582
    WORD $0x394029a2 // ldrb    w2, [x13, #10]
    WORD $0x7100d85f // cmp    w2, #54
    BNE LBB5_829
    WORD $0x71002c3f // cmp    w1, #11
    BEQ LBB5_582
    WORD $0x39402da2 // ldrb    w2, [x13, #11]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100303f // cmp    w1, #12
    BEQ LBB5_582
    WORD $0x394031a2 // ldrb    w2, [x13, #12]
    WORD $0x7100d45f // cmp    w2, #53
    BNE LBB5_830
    WORD $0x7100343f // cmp    w1, #13
    BEQ LBB5_582
    WORD $0x394035a2 // ldrb    w2, [x13, #13]
    WORD $0x7100c45f // cmp    w2, #49
    BNE LBB5_659
    WORD $0x7100383f // cmp    w1, #14
    BEQ LBB5_582
    WORD $0x394039a2 // ldrb    w2, [x13, #14]
    WORD $0x7100d45f // cmp    w2, #53
    BNE LBB5_830
    WORD $0x71003c3f // cmp    w1, #15
    BEQ LBB5_582
    WORD $0x39403da2 // ldrb    w2, [x13, #15]
    WORD $0x7100d85f // cmp    w2, #54
    BNE LBB5_829
    WORD $0x7100403f // cmp    w1, #16
    BEQ LBB5_582
    WORD $0x394041a2 // ldrb    w2, [x13, #16]
    WORD $0x7100d45f // cmp    w2, #53
    BNE LBB5_830
    WORD $0x7100443f // cmp    w1, #17
    BEQ LBB5_582
    WORD $0x394045a2 // ldrb    w2, [x13, #17]
    WORD $0x7100d05f // cmp    w2, #52
    BNE LBB5_823
    WORD $0x7100483f // cmp    w1, #18
    BEQ LBB5_582
    WORD $0x394049a2 // ldrb    w2, [x13, #18]
    WORD $0x7100c05f // cmp    w2, #48
    BNE LBB5_777
    WORD $0x71004c3f // cmp    w1, #19
    BEQ LBB5_582
    WORD $0x39404da2 // ldrb    w2, [x13, #19]
    WORD $0x7100d05f // cmp    w2, #52
    BNE LBB5_823
    WORD $0x7100503f // cmp    w1, #20
    BEQ LBB5_582
    WORD $0x394051a2 // ldrb    w2, [x13, #20]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100543f // cmp    w1, #21
    BEQ LBB5_582
    WORD $0x394055a2 // ldrb    w2, [x13, #21]
    WORD $0x7100cc5f // cmp    w2, #51
    BNE LBB5_798
    WORD $0x7100583f // cmp    w1, #22
    BEQ LBB5_582
    WORD $0x394059a2 // ldrb    w2, [x13, #22]
    WORD $0x7100d85f // cmp    w2, #54
    BNE LBB5_829
    WORD $0x71005c3f // cmp    w1, #23
    BEQ LBB5_582
    WORD $0x39405da2 // ldrb    w2, [x13, #23]
    WORD $0x7100cc5f // cmp    w2, #51
    BNE LBB5_798
    WORD $0x7100603f // cmp    w1, #24
    BEQ LBB5_582
    WORD $0x394061a2 // ldrb    w2, [x13, #24]
    WORD $0x7100c45f // cmp    w2, #49
    BNE LBB5_659
    WORD $0x7100643f // cmp    w1, #25
    BEQ LBB5_582
    WORD $0x394065a2 // ldrb    w2, [x13, #25]
    WORD $0x7100d85f // cmp    w2, #54
    BNE LBB5_829
    WORD $0x7100683f // cmp    w1, #26
    BEQ LBB5_582
    WORD $0x394069a2 // ldrb    w2, [x13, #26]
    WORD $0x7100d85f // cmp    w2, #54
    BNE LBB5_829
    WORD $0x71006c3f // cmp    w1, #27
    BEQ LBB5_582
    WORD $0x39406da2 // ldrb    w2, [x13, #27]
    WORD $0x7100e05f // cmp    w2, #56
    BNE LBB5_832
    WORD $0x7100703f // cmp    w1, #28
    BEQ LBB5_582
    WORD $0x394071a2 // ldrb    w2, [x13, #28]
    WORD $0x7100c05f // cmp    w2, #48
    BNE LBB5_777
    WORD $0x7100743f // cmp    w1, #29
    BEQ LBB5_582
    WORD $0x394075a2 // ldrb    w2, [x13, #29]
    WORD $0x7100e45f // cmp    w2, #57
    BNE LBB5_833
    WORD $0x7100783f // cmp    w1, #30
    BEQ LBB5_582
    WORD $0x394079a2 // ldrb    w2, [x13, #30]
    WORD $0x7100c05f // cmp    w2, #48
    BNE LBB5_777
    WORD $0x71007c3f // cmp    w1, #31
    BEQ LBB5_582
    WORD $0x39407da2 // ldrb    w2, [x13, #31]
    WORD $0x7100e05f // cmp    w2, #56
    BNE LBB5_832
    WORD $0x7100803f // cmp    w1, #32
    BEQ LBB5_582
    WORD $0x394081a2 // ldrb    w2, [x13, #32]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100843f // cmp    w1, #33
    BEQ LBB5_582
    WORD $0x394085a2 // ldrb    w2, [x13, #33]
    WORD $0x7100c05f // cmp    w2, #48
    BNE LBB5_777
    WORD $0x7100883f // cmp    w1, #34
    BEQ LBB5_582
    WORD $0x394089a2 // ldrb    w2, [x13, #34]
    WORD $0x7100cc5f // cmp    w2, #51
    BNE LBB5_798
    WORD $0x71008c3f // cmp    w1, #35
    BEQ LBB5_582
    WORD $0x39408da2 // ldrb    w2, [x13, #35]
    WORD $0x7100c45f // cmp    w2, #49
    BNE LBB5_659
    WORD $0x7100903f // cmp    w1, #36
    BEQ LBB5_582
    WORD $0x394091a2 // ldrb    w2, [x13, #36]
    WORD $0x7100c85f // cmp    w2, #50
    BNE LBB5_789
    WORD $0x7100943f // cmp    w1, #37
    BEQ LBB5_582
    WORD $0x394095a2 // ldrb    w2, [x13, #37]
    WORD $0x7100d45f // cmp    w2, #53
    BNE LBB5_830
    WORD $0x7100983f // cmp    w1, #38
    BEQ LBB5_582
    B LBB5_662
LBB5_659:
    WORD $0x5280062f // mov    w15, #49
LBB5_660:
    WORD $0x6b0f005f // cmp    w2, w15
    BHS LBB5_662
LBB5_661:
    WORD $0x528001e2 // mov    w2, #15
    B LBB5_663
LBB5_662:
    WORD $0x52800202 // mov    w2, #16
LBB5_663:
    WORD $0x7100043f // cmp    w1, #1
    BLT LBB5_673
LBB5_664:
    WORD $0x0b010044 // add    w4, w2, w1
    WORD $0x92407c21 // and    x1, x1, #0xffffffff
    WORD $0x93407c83 // sxtw    x3, w4
    WORD $0xb202e7e6 // mov    x6, #-3689348814741910324
    WORD $0xaa1f03ef // mov    x15, xzr
    WORD $0x91000421 // add    x1, x1, #1
    WORD $0xd1000463 // sub    x3, x3, #1
    WORD $0xd2ff4005 // mov    x5, #-432345564227567616
    WORD $0xf29999a6 // movk    x6, #52429
    WORD $0x92800127 // mov    x7, #-10
    B LBB5_666
LBB5_665:
    WORD $0xf100029f // cmp    x20, #0
    WORD $0x1a9f05ce // csinc    w14, w14, wzr, eq
    WORD $0x51000484 // sub    w4, w4, #1
    WORD $0xd1000463 // sub    x3, x3, #1
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xf100043f // cmp    x1, #1
    BLS LBB5_668
LBB5_666:
    WORD $0x51000833 // sub    w19, w1, #2
    WORD $0xeb10007f // cmp    x3, x16
    WORD $0x387349b3 // ldrb    w19, [x13, w19, uxtw]
    WORD $0x8b13d5ef // add    x15, x15, x19, lsl #53
    WORD $0x8b0501f3 // add    x19, x15, x5
    WORD $0x9bc67e6f // umulh    x15, x19, x6
    WORD $0xd343fdef // lsr    x15, x15, #3
    WORD $0x9b074df4 // madd    x20, x15, x7, x19
    BHS LBB5_665
    WORD $0x1100c294 // add    w20, w20, #48
    WORD $0x382369b4 // strb    w20, [x13, x3]
    WORD $0x51000484 // sub    w4, w4, #1
    WORD $0xd1000463 // sub    x3, x3, #1
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xf100043f // cmp    x1, #1
    BHI LBB5_666
LBB5_668:
    WORD $0xf1002a7f // cmp    x19, #10
    BLO LBB5_673
    WORD $0x93407c81 // sxtw    x1, w4
    WORD $0xb202e7e3 // mov    x3, #-3689348814741910324
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xf29999a3 // movk    x3, #52429
    WORD $0x92800124 // mov    x4, #-10
    B LBB5_671
LBB5_670:
    WORD $0xf10000df // cmp    x6, #0
    WORD $0x1a9f05ce // csinc    w14, w14, wzr, eq
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xf10025ff // cmp    x15, #9
    WORD $0xaa0503ef // mov    x15, x5
    BLS LBB5_673
LBB5_671:
    WORD $0x9bc37de5 // umulh    x5, x15, x3
    WORD $0xeb10003f // cmp    x1, x16
    WORD $0xd343fca5 // lsr    x5, x5, #3
    WORD $0x9b043ca6 // madd    x6, x5, x4, x15
    BHS LBB5_670
    WORD $0x1100c0c6 // add    w6, w6, #48
    WORD $0x382169a6 // strb    w6, [x13, x1]
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xf10025ff // cmp    x15, #9
    WORD $0xaa0503ef // mov    x15, x5
    BHI LBB5_671
LBB5_673:
    WORD $0x0b00004f // add    w15, w2, w0
    WORD $0x0b110051 // add    w17, w2, w17
    WORD $0xeb2fc21f // cmp    x16, w15, sxtw
    WORD $0x1a9081ef // csel    w15, w15, w16, hi
    WORD $0x710005ff // cmp    w15, #1
    BLT LBB5_678
    WORD $0x8b0d01f0 // add    x16, x15, x13
    WORD $0x385ff210 // ldurb    w16, [x16, #-1]
    WORD $0x7100c21f // cmp    w16, #48
    BNE LBB5_679
LBB5_675:
    WORD $0xf10005f0 // subs    x16, x15, #1
    BLS LBB5_695
    WORD $0x510009ef // sub    w15, w15, #2
    WORD $0x386f49a0 // ldrb    w0, [x13, w15, uxtw]
    WORD $0xaa1003ef // mov    x15, x16
    WORD $0x7100c01f // cmp    w0, #48
    BEQ LBB5_675
    WORD $0x2a1003ef // mov    w15, w16
    B LBB5_679
LBB5_678:
    CMP $0, R15
    BEQ LBB5_702
LBB5_679:
    WORD $0x7100523f // cmp    w17, #20
    BLE LBB5_681
    WORD $0x9280000d // mov    x13, #-1
    B LBB5_707
LBB5_681:
    WORD $0x7100063f // cmp    w17, #1
    BLT LBB5_686
    WORD $0x2a1103e1 // mov    w1, w17
    WORD $0x0aaf7de2 // bic    w2, w15, w15, asr #31
    WORD $0xd1000420 // sub    x0, x1, #1
    WORD $0xaa1f03f0 // mov    x16, xzr
    WORD $0xeb02001f // cmp    x0, x2
    WORD $0x52800144 // mov    w4, #10
    WORD $0x9a823000 // csel    x0, x0, x2, lo
    WORD $0xaa0d03e5 // mov    x5, x13
    WORD $0x91000403 // add    x3, x0, #1
LBB5_683:
    CMP $0, R2
    BEQ LBB5_687
    WORD $0x384014a6 // ldrb    w6, [x5], #1
    WORD $0xd1000421 // sub    x1, x1, #1
    WORD $0xd1000442 // sub    x2, x2, #1
    WORD $0x9b041a10 // madd    x16, x16, x4, x6
    WORD $0xd100c210 // sub    x16, x16, #48
    CMP $0, R1
    BNE LBB5_683
    WORD $0xaa0303e0 // mov    x0, x3
    B LBB5_687
LBB5_686:
    WORD $0x2a1f03e0 // mov    w0, wzr
    WORD $0xaa1f03f0 // mov    x16, xzr
LBB5_687:
    WORD $0x6b000221 // subs    w1, w17, w0
    BLE LBB5_694
    WORD $0x7100103f // cmp    w1, #4
    BLO LBB5_692
    WORD $0x52800022 // mov    w2, #1
    WORD $0x52800143 // mov    w3, #10
    WORD $0x25d8e040 // ptrue    p0.d, vl2
    WORD $0x4e080c40 // dup    v0.2d, x2
    WORD $0x121e7422 // and    w2, w1, #0xfffffffc
    WORD $0x0b020000 // add    w0, w0, w2
    WORD $0x4e080c62 // dup    v2.2d, x3
    WORD $0x4ea01c01 // mov    v1.16b, v0.16b
    WORD $0x4e181e01 // mov    v1.d[1], x16
    WORD $0x2a0203f0 // mov    w16, w2
LBB5_690:
    WORD $0x71001210 // subs    w16, w16, #4
    WORD $0x04d00040 // mul    z0.d, p0/m, z0.d, z2.d
    WORD $0x04d00041 // mul    z1.d, p0/m, z1.d, z2.d
    BNE LBB5_690
    WORD $0x4ec07822 // zip2    v2.2d, v1.2d, v0.2d
    WORD $0x6b02003f // cmp    w1, w2
    WORD $0x4ec03820 // zip1    v0.2d, v1.2d, v0.2d
    WORD $0x04d00040 // mul    z0.d, p0/m, z0.d, z2.d
    WORD $0x25d8e020 // ptrue    p0.d, vl1
    WORD $0x6e004001 // ext    v1.16b, v0.16b, v0.16b, #8
    WORD $0x04d00020 // mul    z0.d, p0/m, z0.d, z1.d
    WORD $0x9e660010 // fmov    x16, d0
    BEQ LBB5_694
LBB5_692:
    WORD $0x4b000220 // sub    w0, w17, w0
LBB5_693:
    WORD $0x8b100a10 // add    x16, x16, x16, lsl #2
    WORD $0x71000400 // subs    w0, w0, #1
    WORD $0xd37ffa10 // lsl    x16, x16, #1
    BNE LBB5_693
LBB5_694:
    TST $(1<<31), R17
    BEQ LBB5_696
    B LBB5_703
LBB5_695:
    WORD $0xaa1f03f0 // mov    x16, xzr
    WORD $0x2a1f03f1 // mov    w17, wzr
    WORD $0x510005ef // sub    w15, w15, #1
LBB5_696:
    WORD $0x6b1101ff // cmp    w15, w17
    BLE LBB5_703
    WORD $0x387149a0 // ldrb    w0, [x13, w17, uxtw]
    WORD $0x7100d41f // cmp    w0, #53
    BNE LBB5_701
    WORD $0x11000621 // add    w1, w17, #1
    WORD $0x6b0f003f // cmp    w1, w15
    BNE LBB5_701
    CMP $0, R14
    BEQ LBB5_709
    WORD $0x5280002d // mov    w13, #1
    WORD $0x8b2d420d // add    x13, x16, w13, uxtw
    WORD $0xd2e0040e // mov    x14, #9007199254740992
    WORD $0xeb0e01bf // cmp    x13, x14
    BNE LBB5_707
    B LBB5_704
LBB5_701:
    WORD $0x7100d01f // cmp    w0, #52
    WORD $0x1a9f97ed // cset    w13, hi
    WORD $0x8b2d420d // add    x13, x16, w13, uxtw
    WORD $0xd2e0040e // mov    x14, #9007199254740992
    WORD $0xeb0e01bf // cmp    x13, x14
    BNE LBB5_707
    B LBB5_704
LBB5_702:
    WORD $0xaa1f03f0 // mov    x16, xzr
LBB5_703:
    WORD $0x2a1f03ed // mov    w13, wzr
    WORD $0x8b3f420d // add    x13, x16, wzr, uxtw
    WORD $0xd2e0040e // mov    x14, #9007199254740992
    WORD $0xeb0e01bf // cmp    x13, x14
    BNE LBB5_707
LBB5_704:
    WORD $0x710ffa5f // cmp    w18, #1022
    BGT LBB5_330
    WORD $0x11000652 // add    w18, w18, #1
    WORD $0xd2e0020d // mov    x13, #4503599627370496
    B LBB5_707
LBB5_706:
    WORD $0xaa1f03ed // mov    x13, xzr
LBB5_707:
    WORD $0x110ffe4e // add    w14, w18, #1023
    WORD $0x9374d1af // sbfx    x15, x13, #52, #1
    WORD $0x120029ce // and    w14, w14, #0x7ff
    WORD $0x8a0ed1ee // and    x14, x15, x14, lsl #52
    B LBB5_333
LBB5_708:
    WORD $0x2a1f03e7 // mov    w7, wzr
    B LBB5_405
LBB5_709:
    CMP $0, R17
    BEQ LBB5_703
    WORD $0x5100062e // sub    w14, w17, #1
    WORD $0x386e49ad // ldrb    w13, [x13, w14, uxtw]
    WORD $0x120001ad // and    w13, w13, #0x1
    WORD $0x8b2d420d // add    x13, x16, w13, uxtw
    WORD $0xd2e0040e // mov    x14, #9007199254740992
    WORD $0xeb0e01bf // cmp    x13, x14
    BNE LBB5_707
    B LBB5_704
LBB5_711:
    WORD $0xaa1003f1 // mov    x17, x16
    WORD $0x38401622 // ldrb    w2, [x17], #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_722
    WORD $0xd284c00f // mov    x15, #9728
    WORD $0x5280002e // mov    w14, #1
    WORD $0xf2c0002f // movk    x15, #1, lsl #32
    WORD $0x9ac221ce // lsl    x14, x14, x2
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_722
    WORD $0x39400602 // ldrb    w2, [x16, #1]
    WORD $0x91000a11 // add    x17, x16, #2
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_775
    WORD $0x5280002e // mov    w14, #1
    WORD $0x9ac221ce // lsl    x14, x14, x2
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_775
    WORD $0xf940490e // ldr    x14, [x8, #144]
    WORD $0xcb0e022f // sub    x15, x17, x14
    WORD $0xf100fdff // cmp    x15, #63
    BHI LBB5_718
    WORD $0xf9404d10 // ldr    x16, [x8, #152]
    WORD $0x92800011 // mov    x17, #-1
    WORD $0x9acf222f // lsl    x15, x17, x15
    WORD $0xea0f020f // ands    x15, x16, x15
    BNE LBB5_721
    WORD $0x910101d1 // add    x17, x14, #64
LBB5_718:
    ADR LCPI5_0, R14
    ADR LCPI5_1, R15
    ADR LCPI5_2, R16
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001c0 // ldr    q0, [x14, :lo12:.LCPI5_0]
    WORD $0xd101022e // sub    x14, x17, #64
    WORD $0x3dc001e2 // ldr    q2, [x15, :lo12:.LCPI5_1]
    WORD $0x3dc00203 // ldr    q3, [x16, :lo12:.LCPI5_2]
LBB5_719:
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
    WORD $0x1e2600af // fmov    w15, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e260090 // fmov    w16, s4
    WORD $0x1e2600b1 // fmov    w17, s5
    WORD $0x1e2600c1 // fmov    w1, s6
    WORD $0x33103df0 // bfi    w16, w15, #16, #16
    WORD $0xaa11820f // orr    x15, x16, x17, lsl #32
    WORD $0xaa01c1ef // orr    x15, x15, x1, lsl #48
    WORD $0xb10005ff // cmn    x15, #1
    BEQ LBB5_719
    WORD $0xaa2f03ef // mvn    x15, x15
    WORD $0xa9093d0e // stp    x14, x15, [x8, #144]
LBB5_721:
    WORD $0xdac001ef // rbit    x15, x15
    WORD $0xdac011ef // clz    x15, x15
    WORD $0x8b0f01d1 // add    x17, x14, x15
    WORD $0x38401622 // ldrb    w2, [x17], #1
LBB5_722:
    WORD $0x7101745f // cmp    w2, #93
    BNE LBB5_776
LBB5_723:
    WORD $0xb940d10e // ldr    w14, [x8, #208]
    WORD $0x8b0c11ac // add    x12, x13, x12, lsl #4
    WORD $0x110005ce // add    w14, w14, #1
    WORD $0xb900d10e // str    w14, [x8, #208]
LBB5_724:
    WORD $0xcb0b000d // sub    x13, x0, x11
    WORD $0xf940058e // ldr    x14, [x12, #8]
    WORD $0xd344fdad // lsr    x13, x13, #4
    WORD $0xf900550e // str    x14, [x8, #168]
    WORD $0xb9000d6d // str    w13, [x11, #12]
    WORD $0xf940016d // ldr    x13, [x11]
    WORD $0xb900099f // str    wzr, [x12, #8]
    WORD $0xf940590c // ldr    x12, [x8, #176]
    WORD $0xb940e50e // ldr    w14, [x8, #228]
    WORD $0x92609dad // and    x13, x13, #0xffffffff000000ff
    WORD $0xeb0e019f // cmp    x12, x14
    WORD $0xf900016d // str    x13, [x11]
    BLS LBB5_726
    WORD $0xf140059f // cmp    x12, #1, lsl #12
    WORD $0xb900e50c // str    w12, [x8, #228]
    BHI LBB5_96
LBB5_726:
    WORD $0xf9405510 // ldr    x16, [x8, #168]
    WORD $0xd100058b // sub    x11, x12, #1
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0xb100061f // cmn    x16, #1
    WORD $0xf900590b // str    x11, [x8, #176]
    BNE LBB5_1924
    B LBB5_834
LBB5_727:
    WORD $0x3940016c // ldrb    w12, [x11]
    WORD $0xf100199f // cmp    x12, #6
    BNE LBB5_731
    WORD $0x3943210c // ldrb    w12, [x8, #200]
    CMP $0, R12
    BEQ LBB5_747
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0xaa1003ef // mov    x15, x16
    B LBB5_935
LBB5_730:
    WORD $0x52800062 // mov    w2, #3
    B LBB5_337
LBB5_731:
    WORD $0xaa1003ef // mov    x15, x16
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0x384015e0 // ldrb    w0, [x15], #1
    WORD $0x7100801f // cmp    w0, #32
    BHI LBB5_779
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ac021ad // lsl    x13, x13, x0
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_779
    WORD $0x39400600 // ldrb    w0, [x16, #1]
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0x91000a0f // add    x15, x16, #2
    WORD $0x7100801f // cmp    w0, #32
    BHI LBB5_780
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ac021ad // lsl    x13, x13, x0
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_780
    WORD $0xf9404910 // ldr    x16, [x8, #144]
    WORD $0xcb1001ec // sub    x12, x15, x16
    WORD $0xf100fd9f // cmp    x12, #63
    BHI LBB5_782
    WORD $0xf9404d0d // ldr    x13, [x8, #152]
    WORD $0x9280000e // mov    x14, #-1
    WORD $0x9acc21cc // lsl    x12, x14, x12
    WORD $0xea0c01af // ands    x15, x13, x12
    BEQ LBB5_781
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_2399
LBB5_738:
    WORD $0x9280000d // mov    x13, #-1
    WORD $0xaa1203e9 // mov    x9, x18
    WORD $0x9280000e // mov    x14, #-1
    WORD $0x9280000a // mov    x10, #-1
    B LBB5_187
LBB5_739:
    WORD $0x2a1f03f0 // mov    w16, wzr
    WORD $0x5284e201 // mov    w1, #10000
    B LBB5_215
LBB5_740:
    WORD $0xcb0a0129 // sub    x9, x9, x10
    WORD $0x39400131 // ldrb    w17, [x9]
    WORD $0x5100c22a // sub    w10, w17, #48
    WORD $0x7100255f // cmp    w10, #9
    BHI LBB5_778
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03ea // mov    x10, xzr
    WORD $0x52800149 // mov    w9, #10
LBB5_742:
    WORD $0x8b0a024e // add    x14, x18, x10
    WORD $0x9b097dad // mul    x13, x13, x9
    WORD $0x8b3141ad // add    x13, x13, w17, uxtw
    WORD $0x394005d1 // ldrb    w17, [x14, #1]
    WORD $0xd100c1ad // sub    x13, x13, #48
    WORD $0x5100c22e // sub    w14, w17, #48
    WORD $0x710025df // cmp    w14, #9
    WORD $0xfa529942 // ccmp    x10, #18, #2, ls
    WORD $0x9100054a // add    x10, x10, #1
    BLO LBB5_742
    WORD $0x710025df // cmp    w14, #9
    BHI LBB5_785
    WORD $0xaa1f03f0 // mov    x16, xzr
LBB5_745:
    WORD $0x8b100249 // add    x9, x18, x16
    WORD $0x91000610 // add    x16, x16, #1
    WORD $0x8b0a0129 // add    x9, x9, x10
    WORD $0x39400531 // ldrb    w17, [x9, #1]
    WORD $0x5100c229 // sub    w9, w17, #48
    WORD $0x7100293f // cmp    w9, #10
    BLO LBB5_745
    WORD $0x8b0a0249 // add    x9, x18, x10
    WORD $0x5280002e // mov    w14, #1
    WORD $0x8b100129 // add    x9, x9, x16
    B LBB5_47
LBB5_747:
    WORD $0xaa1003f1 // mov    x17, x16
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0x38401632 // ldrb    w18, [x17], #1
    WORD $0x7100825f // cmp    w18, #32
    BHI LBB5_790
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ad221ad // lsl    x13, x13, x18
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_790
    WORD $0x39400612 // ldrb    w18, [x16, #1]
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0x91000a11 // add    x17, x16, #2
    WORD $0x7100825f // cmp    w18, #32
    BHI LBB5_791
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ad221ad // lsl    x13, x13, x18
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_791
    WORD $0xf940490f // ldr    x15, [x8, #144]
    WORD $0xcb0f022c // sub    x12, x17, x15
    WORD $0xf100fd9f // cmp    x12, #63
    BHI LBB5_793
    WORD $0xf9404d0d // ldr    x13, [x8, #152]
    WORD $0x9280000e // mov    x14, #-1
    WORD $0x9acc21cc // lsl    x12, x14, x12
    WORD $0xea0c01b0 // ands    x16, x13, x12
    BEQ LBB5_792
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_1641
LBB5_754:
    WORD $0x528001a2 // mov    w2, #13
    WORD $0xaa0903ef // mov    x15, x9
    B LBB5_73
LBB5_755:
    WORD $0x528001a2 // mov    w2, #13
    WORD $0xaa0903ef // mov    x15, x9
    B LBB5_77
LBB5_756:
    WORD $0x528001a2 // mov    w2, #13
    WORD $0xaa0903ef // mov    x15, x9
    B LBB5_67
LBB5_757:
    WORD $0x2538c3e2 // mov    z2.b, #31
LBB5_758:
    WORD $0xa400a203 // ld1b    { z3.b }, p0/z, [x16]
    WORD $0x2400a061 // cmpeq    p1.b, p0/z, z3.b, z0.b
    WORD $0x2401a063 // cmpeq    p3.b, p0/z, z3.b, z1.b
    WORD $0x2402a062 // cmpeq    p2.b, p0/z, z3.b, z2.b
    WORD $0x25834c64 // mov    p4.b, p3.b
    WORD $0x25c24025 // orrs    p5.b, p0/z, p1.b, p2.b
    BEQ LBB5_760
    WORD $0x259040a4 // brkb    p4.b, p0/z, p5.b
    WORD $0x25034084 // and    p4.b, p0/z, p4.b, p3.b
LBB5_760:
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_767
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_764
    WORD $0x25904063 // brkb    p3.b, p0/z, p3.b
    WORD $0x25414064 // ands    p4.b, p0/z, p3.b, p1.b
    BNE LBB5_106
    WORD $0x25024063 // and    p3.b, p0/z, p3.b, p2.b
    B LBB5_765
LBB5_764:
    WORD $0x25824843 // mov    p3.b, p2.b
    WORD $0x2550c020 // ptest    p0, p1.b
    BNE LBB5_106
LBB5_765:
    WORD $0x2550c060 // ptest    p0, p3.b
    BNE LBB5_796
    WORD $0x91008210 // add    x16, x16, #32
    B LBB5_758
LBB5_767:
    WORD $0x25904061 // brkb    p1.b, p0/z, p3.b
    WORD $0x2a1f03ea // mov    w10, wzr
    WORD $0x2520802c // cntp    x12, p0, p1.b
    WORD $0x8b10018c // add    x12, x12, x16
    WORD $0x9100058d // add    x13, x12, #1
    WORD $0xaa2f03ec // mvn    x12, x15
    WORD $0x8b0c01ac // add    x12, x13, x12
    WORD $0xaa0d03ef // mov    x15, x13
    TST $(1<<63), R12
    BNE LBB5_801
LBB5_768:
    WORD $0x2a1f03e2 // mov    w2, wzr
    B LBB5_803
LBB5_769:
    WORD $0x910009ef // add    x15, x15, #2
    B LBB5_73
LBB5_770:
    WORD $0x910009ef // add    x15, x15, #2
    B LBB5_77
LBB5_771:
    WORD $0x910009ef // add    x15, x15, #2
    WORD $0x528001a2 // mov    w2, #13
    B LBB5_67
LBB5_772:
    WORD $0x91000def // add    x15, x15, #3
    B LBB5_67
LBB5_773:
    WORD $0x7101f45f // cmp    w2, #125
    BEQ LBB5_575
LBB5_774:
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    B LBB5_837
LBB5_775:
    WORD $0x7101745f // cmp    w2, #93
    BEQ LBB5_723
LBB5_776:
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_1703
LBB5_777:
    WORD $0x5280060f // mov    w15, #48
    B LBB5_660
LBB5_778:
    WORD $0xaa1f03ea // mov    x10, xzr
    WORD $0x2a1f03ee // mov    w14, wzr
    WORD $0x2a1f03f0 // mov    w16, wzr
    WORD $0xaa1f03ed // mov    x13, xzr
    B LBB5_47
LBB5_779:
    WORD $0xaa0c03fd // mov    x29, x12
    WORD $0xf940016d // ldr    x13, [x11]
    WORD $0x7100b01f // cmp    w0, #44
    WORD $0x910401ad // add    x13, x13, #256
    WORD $0xf900016d // str    x13, [x11]
    BEQ LBB5_2401
    B LBB5_2413
LBB5_780:
    WORD $0xaa0c03fd // mov    x29, x12
    WORD $0xf940016d // ldr    x13, [x11]
    WORD $0x7100b01f // cmp    w0, #44
    WORD $0x910401ad // add    x13, x13, #256
    WORD $0xf900016d // str    x13, [x11]
    BEQ LBB5_2401
    B LBB5_2413
LBB5_781:
    WORD $0x9101020f // add    x15, x16, #64
LBB5_782:
    ADR LCPI5_0, R12
    ADR LCPI5_1, R13
    ADR LCPI5_2, R14
    WORD $0xd10101f0 // sub    x16, x15, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc00180 // ldr    q0, [x12, :lo12:.LCPI5_0]
    WORD $0x3dc001a2 // ldr    q2, [x13, :lo12:.LCPI5_1]
    WORD $0x3dc001c3 // ldr    q3, [x14, :lo12:.LCPI5_2]
LBB5_783:
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
    WORD $0x1e2600ac // fmov    w12, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e26008d // fmov    w13, s4
    WORD $0x1e2600ae // fmov    w14, s5
    WORD $0x1e2600cf // fmov    w15, s6
    WORD $0x33103d8d // bfi    w13, w12, #16, #16
    WORD $0xaa0e81ac // orr    x12, x13, x14, lsl #32
    WORD $0xaa0fc18d // orr    x13, x12, x15, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_783
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_2398
LBB5_785:
    WORD $0x2a1f03ee // mov    w14, wzr
    WORD $0x2a1f03f0 // mov    w16, wzr
    WORD $0x8b0a0249 // add    x9, x18, x10
    B LBB5_47
LBB5_786:
    WORD $0x7100061f // cmp    w16, #1
    BNE LBB5_242
    WORD $0x5280014a // mov    w10, #10
    WORD $0x9bca7daa // umulh    x10, x13, x10
    WORD $0xeb0a03ff // cmp    xzr, x10
    BEQ LBB5_824
    WORD $0x7100019f // cmp    w12, #0
    WORD $0x1280000a // mov    w10, #-1
    WORD $0x5a8a154a // cneg    w10, w10, eq
    WORD $0x52800030 // mov    w16, #1
    B LBB5_254
LBB5_789:
    WORD $0x5280064f // mov    w15, #50
    B LBB5_660
LBB5_790:
    WORD $0xaa0c03fd // mov    x29, x12
    B LBB5_1938
LBB5_791:
    WORD $0xaa0c03fd // mov    x29, x12
    B LBB5_1642
LBB5_792:
    WORD $0x910101f1 // add    x17, x15, #64
LBB5_793:
    ADR LCPI5_0, R12
    ADR LCPI5_1, R13
    ADR LCPI5_2, R14
    WORD $0xd101022f // sub    x15, x17, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc00180 // ldr    q0, [x12, :lo12:.LCPI5_0]
    WORD $0x3dc001a2 // ldr    q2, [x13, :lo12:.LCPI5_1]
    WORD $0x3dc001c3 // ldr    q3, [x14, :lo12:.LCPI5_2]
LBB5_794:
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
    WORD $0x1e2600ac // fmov    w12, s5
    WORD $0x4e71b8c5 // addv    h5, v6.8h
    WORD $0x4e71b8e6 // addv    h6, v7.8h
    WORD $0x1e26008d // fmov    w13, s4
    WORD $0x1e2600ae // fmov    w14, s5
    WORD $0x1e2600d0 // fmov    w16, s6
    WORD $0x33103d8d // bfi    w13, w12, #16, #16
    WORD $0xaa0e81ac // orr    x12, x13, x14, lsl #32
    WORD $0xaa10c18d // orr    x13, x12, x16, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_794
    WORD $0xaa1f03ec // mov    x12, xzr
    WORD $0xaa1f03fd // mov    x29, xzr
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_1640
LBB5_796:
    WORD $0x2a1f03ea // mov    w10, wzr
LBB5_797:
    WORD $0x25904041 // brkb    p1.b, p0/z, p2.b
    WORD $0x2520802c // cntp    x12, p0, p1.b
    WORD $0x8b0c0214 // add    x20, x16, x12
    WORD $0x9280000c // mov    x12, #-1
    B LBB5_802
LBB5_798:
    WORD $0x5280066f // mov    w15, #51
    B LBB5_660
LBB5_799:
    WORD $0x8b1300ca // add    x10, x6, x19
    WORD $0x3940014c // ldrb    w12, [x10]
    WORD $0x7100899f // cmp    w12, #34
    BNE LBB5_805
LBB5_800:
    WORD $0x8b1300ec // add    x12, x7, x19
    WORD $0x9100054d // add    x13, x10, #1
    WORD $0xcb0f018c // sub    x12, x12, x15
    WORD $0x5280002a // mov    w10, #1
    WORD $0xaa0d03ef // mov    x15, x13
    TST $(1<<63), R12
    BEQ LBB5_768
LBB5_801:
    WORD $0xaa0f03f4 // mov    x20, x15
LBB5_802:
    WORD $0x4b0c03e2 // neg    w2, w12
    WORD $0xaa1403ef // mov    x15, x20
LBB5_803:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x7100015f // cmp    w10, #0
    WORD $0x5280018a // mov    w10, #12
    WORD $0x5280008e // mov    w14, #4
    WORD $0x9a8a01ca // csel    x10, x14, x10, eq
    WORD $0xd2c0002e // mov    x14, #4294967296
    WORD $0xf90005ac // str    x12, [x13, #8]
    WORD $0xaa0b814a // orr    x10, x10, x11, lsl #32
    WORD $0xf940510c // ldr    x12, [x8, #160]
    WORD $0x8b0e014a // add    x10, x10, x14
    WORD $0xf9406110 // ldr    x16, [x8, #192]
    WORD $0xcb0f0129 // sub    x9, x9, x15
    WORD $0xb940d50b // ldr    w11, [x8, #212]
    WORD $0x9100818e // add    x14, x12, #32
    WORD $0x9100418c // add    x12, x12, #16
    WORD $0xeb1001df // cmp    x14, x16
    WORD $0xf90001aa // str    x10, [x13]
    WORD $0x1100056b // add    w11, w11, #1
    WORD $0x1a9f87ea // cset    w10, ls
    WORD $0xf900510c // str    x12, [x8, #160]
    WORD $0xb900d50b // str    w11, [x8, #212]
    TST $(1<<63), R9
    BEQ LBB5_302
    WORD $0x528000a2 // mov    w2, #5
    B LBB5_341
LBB5_805:
    WORD $0x8b1300ea // add    x10, x7, x19
    WORD $0x8b1300cd // add    x13, x6, x19
    WORD $0x3900014c // strb    w12, [x10]
    WORD $0x394005ac // ldrb    w12, [x13, #1]
    WORD $0x7100899f // cmp    w12, #34
    BEQ LBB5_815
    WORD $0x3900054c // strb    w12, [x10, #1]
    WORD $0x394009ad // ldrb    w13, [x13, #2]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_816
    WORD $0x8b1300cc // add    x12, x6, x19
    WORD $0x3900094d // strb    w13, [x10, #2]
    WORD $0x39400d8d // ldrb    w13, [x12, #3]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_817
    WORD $0x39000d4d // strb    w13, [x10, #3]
    WORD $0x3940118d // ldrb    w13, [x12, #4]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_818
    WORD $0x8b1300cc // add    x12, x6, x19
    WORD $0x3900114d // strb    w13, [x10, #4]
    WORD $0x3940158d // ldrb    w13, [x12, #5]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_819
    WORD $0x3900154d // strb    w13, [x10, #5]
    WORD $0x3940198d // ldrb    w13, [x12, #6]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_820
    WORD $0x8b1300cc // add    x12, x6, x19
    WORD $0x3900194d // strb    w13, [x10, #6]
    WORD $0x39401d8d // ldrb    w13, [x12, #7]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_821
    WORD $0x39001d4d // strb    w13, [x10, #7]
    WORD $0x91002273 // add    x19, x19, #8
    WORD $0x3940218c // ldrb    w12, [x12, #8]
    WORD $0x7100899f // cmp    w12, #34
    BNE LBB5_805
    WORD $0x8b1300ca // add    x10, x6, x19
    B LBB5_800
LBB5_814:
    WORD $0x5280002a // mov    w10, #1
    WORD $0x9280016c // mov    x12, #-12
    B LBB5_802
LBB5_815:
    WORD $0xcb0f00ea // sub    x10, x7, x15
    WORD $0x910009af // add    x15, x13, #2
    WORD $0x8b13014a // add    x10, x10, x19
    WORD $0x9100054c // add    x12, x10, #1
    B LBB5_822
LBB5_816:
    WORD $0x8b1300ca // add    x10, x6, x19
    WORD $0xcb0f00ec // sub    x12, x7, x15
    WORD $0x91000d4f // add    x15, x10, #3
    WORD $0x8b13018a // add    x10, x12, x19
    WORD $0x9100094c // add    x12, x10, #2
    B LBB5_822
LBB5_817:
    WORD $0xcb0f00ea // sub    x10, x7, x15
    WORD $0x9100118f // add    x15, x12, #4
    WORD $0x8b13014a // add    x10, x10, x19
    WORD $0x91000d4c // add    x12, x10, #3
    B LBB5_822
LBB5_818:
    WORD $0x8b1300ca // add    x10, x6, x19
    WORD $0xcb0f00ec // sub    x12, x7, x15
    WORD $0x9100154f // add    x15, x10, #5
    WORD $0x8b13018a // add    x10, x12, x19
    WORD $0x9100114c // add    x12, x10, #4
    B LBB5_822
LBB5_819:
    WORD $0xcb0f00ea // sub    x10, x7, x15
    WORD $0x9100198f // add    x15, x12, #6
    WORD $0x8b13014a // add    x10, x10, x19
    WORD $0x9100154c // add    x12, x10, #5
    B LBB5_822
LBB5_820:
    WORD $0x8b1300ca // add    x10, x6, x19
    WORD $0xcb0f00ec // sub    x12, x7, x15
    WORD $0x91001d4f // add    x15, x10, #7
    WORD $0x8b13018a // add    x10, x12, x19
    WORD $0x9100194c // add    x12, x10, #6
    B LBB5_822
LBB5_821:
    WORD $0xcb0f00ea // sub    x10, x7, x15
    WORD $0x9100218f // add    x15, x12, #8
    WORD $0x8b13014a // add    x10, x10, x19
    WORD $0x91001d4c // add    x12, x10, #7
LBB5_822:
    WORD $0x5280002a // mov    w10, #1
    TST $(1<<63), R12
    BEQ LBB5_768
    B LBB5_801
LBB5_823:
    WORD $0x5280068f // mov    w15, #52
    B LBB5_660
LBB5_824:
    WORD $0x385ff12a // ldurb    w10, [x9, #-1]
    WORD $0x8b0d09b0 // add    x16, x13, x13, lsl #2
    WORD $0xd37ffa10 // lsl    x16, x16, #1
    WORD $0x5100c14a // sub    w10, w10, #48
    WORD $0x93407d4a // sxtw    x10, w10
    WORD $0x937ffd51 // asr    x17, x10, #63
    WORD $0xab0a020a // adds    x10, x16, x10
    WORD $0x9a913630 // cinc    x16, x17, hs
    WORD $0x93400211 // sbfx    x17, x16, #0, #1
    WORD $0xca100221 // eor    x1, x17, x16
    WORD $0x52800030 // mov    w16, #1
    CMP $0, R1
    BNE LBB5_242
    TST $(1<<63), R17
    BNE LBB5_242
    CMP $0, R12
    BEQ LBB5_315
    WORD $0x9e630140 // ucvtf    d0, x10
    B LBB5_229
LBB5_828:
    WORD $0x8b1300d0 // add    x16, x6, x19
    WORD $0x5280002a // mov    w10, #1
    B LBB5_797
LBB5_829:
    WORD $0x528006cf // mov    w15, #54
    B LBB5_660
LBB5_830:
    WORD $0x528006af // mov    w15, #53
    B LBB5_660
LBB5_831:
    WORD $0x5280002a // mov    w10, #1
    WORD $0x9280016c // mov    x12, #-12
    WORD $0xaa0603f4 // mov    x20, x6
    B LBB5_802
LBB5_832:
    WORD $0x5280070f // mov    w15, #56
    B LBB5_660
LBB5_833:
    WORD $0x5280072f // mov    w15, #57
    B LBB5_660
LBB5_834:
    WORD $0xaa1f03eb // mov    x11, xzr
    CMP ZR, ZR
    BNE LBB5_1926
LBB5_835:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa1103ef // mov    x15, x17
    B LBB5_341
LBB5_836:
LBB5_837:
    WORD $0x7100885f // cmp    w2, #34
    BNE LBB5_909
    WORD $0xf940390e // ldr    x14, [x8, #112]
    WORD $0xcb0a0230 // sub    x16, x17, x10
    WORD $0xaa1103ef // mov    x15, x17
    WORD $0x2518e3e0 // ptrue    p0.b
    WORD $0x2538cb80 // mov    z0.b, #92
    WORD $0x2538c441 // mov    z1.b, #34
    TST $(1<<5), R14
    BNE LBB5_910
    WORD $0x2518e402 // pfalse    p2.b
    B LBB5_842
LBB5_840:
    WORD $0x25904063 // brkb    p3.b, p0/z, p3.b
    WORD $0x25414063 // ands    p3.b, p0/z, p3.b, p1.b
    BNE LBB5_847
LBB5_841:
    WORD $0x910081ef // add    x15, x15, #32
LBB5_842:
    WORD $0xa400a1e2 // ld1b    { z2.b }, p0/z, [x15]
    WORD $0x2401a043 // cmpeq    p3.b, p0/z, z2.b, z1.b
    WORD $0x2400a041 // cmpeq    p1.b, p0/z, z2.b, z0.b
    WORD $0x25834c64 // mov    p4.b, p3.b
    WORD $0x25c24025 // orrs    p5.b, p0/z, p1.b, p2.b
    BEQ LBB5_844
    WORD $0x259040a4 // brkb    p4.b, p0/z, p5.b
    WORD $0x25034084 // and    p4.b, p0/z, p4.b, p3.b
LBB5_844:
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_920
    WORD $0x2550c060 // ptest    p0, p3.b
    BNE LBB5_840
    WORD $0x2550c020 // ptest    p0, p1.b
    BEQ LBB5_841
LBB5_847:
    WORD $0x25904021 // brkb    p1.b, p0/z, p1.b
    WORD $0x5299fa12 // mov    w18, #53200
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0x52832321 // mov    w1, #6425
    WORD $0x8b0d01ef // add    x15, x15, x13
    WORD $0x5288c8c3 // mov    w3, #17990
    WORD $0x52872725 // mov    w5, #14649
    WORD $0x52848014 // mov    w20, #9216
    WORD $0x72b9f9f2 // movk    w18, #53199, lsl #16
    WORD $0x3201c3e0 // mov    w0, #-2139062144
    WORD $0x72a32321 // movk    w1, #6425, lsl #16
    WORD $0x3202c7e2 // mov    w2, #-1061109568
    WORD $0x72a8c8c3 // movk    w3, #17990, lsl #16
    WORD $0x3203cbe4 // mov    w4, #-522133280
    WORD $0x72a72725 // movk    w5, #14649, lsl #16
    WORD $0x3200c3e6 // mov    w6, #16843009
    WORD $0x5297fde7 // mov    w7, #49135
    WORD $0x528017b3 // mov    w19, #189
    WORD $0x72bf9414 // movk    w20, #64672, lsl #16
    WORD $0xaa0f03fa // mov    x26, x15
    WORD $0xaa0f03f9 // mov    x25, x15
    ADR ESCAPED_TAB, R21
    WORD $0x910002b5 // add    x21, x21, :lo12:ESCAPED_TAB
    WORD $0x2538cb80 // mov    z0.b, #92
    WORD $0x2538c441 // mov    z1.b, #34
    WORD $0x2518e401 // pfalse    p1.b
    WORD $0x2538c3e2 // mov    z2.b, #31
LBB5_848:
    WORD $0x3940074d // ldrb    w13, [x26, #1]
    WORD $0xf101d5bf // cmp    x13, #117
    BEQ LBB5_851
    WORD $0x386d6aad // ldrb    w13, [x21, x13]
    CMP $0, R13
    BEQ LBB5_908
    WORD $0x3800172d // strb    w13, [x25], #1
    WORD $0x91000b56 // add    x22, x26, #2
    WORD $0xaa1903f7 // mov    x23, x25
    B LBB5_871
LBB5_851:
    WORD $0xb840234d // ldur    w13, [x26, #2]
    WORD $0x0a2d0016 // bic    w22, w0, w13
    WORD $0x0b1201b7 // add    w23, w13, w18
    WORD $0x6a1702df // tst    w22, w23
    BNE LBB5_1660
    WORD $0x0b0101b7 // add    w23, w13, w1
    WORD $0x2a0d02f7 // orr    w23, w23, w13
    WORD $0x7201c2ff // tst    w23, #0x80808080
    BNE LBB5_1660
    WORD $0x1200d9b7 // and    w23, w13, #0x7f7f7f7f
    WORD $0x4b170058 // sub    w24, w2, w23
    WORD $0x0b0302fb // add    w27, w23, w3
    WORD $0x0a1b0318 // and    w24, w24, w27
    WORD $0x6a16031f // tst    w24, w22
    BNE LBB5_1660
    WORD $0x4b170098 // sub    w24, w4, w23
    WORD $0x0b0502f7 // add    w23, w23, w5
    WORD $0x0a170317 // and    w23, w24, w23
    WORD $0x6a1602ff // tst    w23, w22
    BNE LBB5_1660
    WORD $0x5ac009ad // rev    w13, w13
    WORD $0x91001b56 // add    x22, x26, #6
    WORD $0x1200cdaf // and    w15, w13, #0xf0f0f0f
    WORD $0x0a6d10cd // bic    w13, w6, w13, lsr #4
    WORD $0x2a0d0dad // orr    w13, w13, w13, lsl #3
    WORD $0x0b0f01ad // add    w13, w13, w15
    WORD $0x2a4d11ad // orr    w13, w13, w13, lsr #4
    WORD $0x53105daf // ubfx    w15, w13, #16, #8
    WORD $0x12001dad // and    w13, w13, #0xff
    WORD $0x2a0f21af // orr    w15, w13, w15, lsl #8
    WORD $0x710201ff // cmp    w15, #128
    BLO LBB5_903
    WORD $0x91001337 // add    x23, x25, #4
LBB5_857:
    WORD $0x711ffdff // cmp    w15, #2047
    BLS LBB5_905
    WORD $0x514039ed // sub    w13, w15, #14, lsl #12
    WORD $0x312005bf // cmn    w13, #2049
    BLS LBB5_869
    WORD $0x530a7ded // lsr    w13, w15, #10
    WORD $0x7100d9bf // cmp    w13, #54
    BHI LBB5_906
    WORD $0x394002cd // ldrb    w13, [x22]
    WORD $0x710171bf // cmp    w13, #92
    BNE LBB5_906
    WORD $0x394006cd // ldrb    w13, [x22, #1]
    WORD $0x7101d5bf // cmp    w13, #117
    BNE LBB5_906
    WORD $0xb84022cd // ldur    w13, [x22, #2]
    WORD $0x0a2d0018 // bic    w24, w0, w13
    WORD $0x0b1201b9 // add    w25, w13, w18
    WORD $0x6a19031f // tst    w24, w25
    BNE LBB5_1699
    WORD $0x0b0101b9 // add    w25, w13, w1
    WORD $0x2a0d0339 // orr    w25, w25, w13
    WORD $0x7201c33f // tst    w25, #0x80808080
    BNE LBB5_1699
    WORD $0x1200d9b9 // and    w25, w13, #0x7f7f7f7f
    WORD $0x4b19005a // sub    w26, w2, w25
    WORD $0x0b03033b // add    w27, w25, w3
    WORD $0x0a1b035a // and    w26, w26, w27
    WORD $0x6a18035f // tst    w26, w24
    BNE LBB5_1699
    WORD $0x4b19009a // sub    w26, w4, w25
    WORD $0x0b050339 // add    w25, w25, w5
    WORD $0x0a190359 // and    w25, w26, w25
    WORD $0x6a18033f // tst    w25, w24
    BNE LBB5_1699
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
    BHI LBB5_907
    WORD $0x781fc2e7 // sturh    w7, [x23, #-4]
    WORD $0x2a0d03ef // mov    w15, w13
    WORD $0x381fe2f3 // sturb    w19, [x23, #-2]
    WORD $0x91000ef7 // add    x23, x23, #3
    WORD $0x710201bf // cmp    w13, #128
    BHS LBB5_857
    WORD $0xd10012f9 // sub    x25, x23, #4
    B LBB5_904
LBB5_869:
    WORD $0x530c7ded // lsr    w13, w15, #12
    WORD $0x52801018 // mov    w24, #128
    WORD $0x52801019 // mov    w25, #128
    WORD $0x321b09ad // orr    w13, w13, #0xe0
    WORD $0x33062df8 // bfxil    w24, w15, #6, #6
    WORD $0x330015f9 // bfxil    w25, w15, #0, #6
    WORD $0xd10006ef // sub    x15, x23, #1
    WORD $0x381fc2ed // sturb    w13, [x23, #-4]
    WORD $0x381fd2f8 // sturb    w24, [x23, #-3]
    WORD $0x381fe2f9 // sturb    w25, [x23, #-2]
LBB5_870:
    WORD $0xaa0f03f7 // mov    x23, x15
LBB5_871:
    WORD $0x394002cd // ldrb    w13, [x22]
    WORD $0xaa1603ef // mov    x15, x22
    WORD $0xaa1603fa // mov    x26, x22
    WORD $0xaa1703f9 // mov    x25, x23
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_848
    WORD $0xaa1d03fb // mov    x27, x29
    WORD $0xaa1f03f8 // mov    x24, xzr
    WORD $0xa41842c3 // ld1b    { z3.b }, p0/z, [x22, x24]
    WORD $0x25814422 // mov    p2.b, p1.b
    TST $(1<<5), R14
    BEQ LBB5_876
    B LBB5_874
LBB5_873:
    WORD $0xe41842e3 // st1b    { z3.b }, p0, [x23, x24]
    WORD $0x91008318 // add    x24, x24, #32
    WORD $0xa41842c3 // ld1b    { z3.b }, p0/z, [x22, x24]
    WORD $0x25814422 // mov    p2.b, p1.b
    TST $(1<<5), R14
    BEQ LBB5_876
LBB5_874:
    WORD $0x2402a062 // cmpeq    p2.b, p0/z, z3.b, z2.b
    WORD $0x2401a064 // cmpeq    p4.b, p0/z, z3.b, z1.b
    WORD $0x2400a063 // cmpeq    p3.b, p0/z, z3.b, z0.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BNE LBB5_877
LBB5_875:
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BEQ LBB5_878
    B LBB5_1648
LBB5_876:
    WORD $0x2401a064 // cmpeq    p4.b, p0/z, z3.b, z1.b
    WORD $0x2400a063 // cmpeq    p3.b, p0/z, z3.b, z0.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BEQ LBB5_875
LBB5_877:
    WORD $0x259040c5 // brkb    p5.b, p0/z, p6.b
    WORD $0x250440a5 // and    p5.b, p0/z, p5.b, p4.b
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_1648
LBB5_878:
    WORD $0x2550c080 // ptest    p0, p4.b
    WORD $0x25824845 // mov    p5.b, p2.b
    WORD $0x1a9f07ed // cset    w13, ne
    BEQ LBB5_880
    WORD $0x25904085 // brkb    p5.b, p0/z, p4.b
    WORD $0x250240a5 // and    p5.b, p0/z, p5.b, p2.b
LBB5_880:
    TST $(1<<5), R14
    BEQ LBB5_882
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_1682
LBB5_882:
    CMP $0, R13
    BEQ LBB5_884
    WORD $0x25904082 // brkb    p2.b, p0/z, p4.b
    WORD $0x25034043 // and    p3.b, p0/z, p2.b, p3.b
LBB5_884:
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_873
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x8b1802f9 // add    x25, x23, x24
    WORD $0xaa0f03fa // mov    x26, x15
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x394001ed // ldrb    w13, [x15]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_848
LBB5_886:
    WORD $0x8b1802f9 // add    x25, x23, x24
    WORD $0x8b1802da // add    x26, x22, x24
    WORD $0x3900032d // strb    w13, [x25]
    WORD $0x3940074d // ldrb    w13, [x26, #1]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_895
    WORD $0x3900072d // strb    w13, [x25, #1]
    WORD $0x39400b4d // ldrb    w13, [x26, #2]
    WORD $0x710171bf // cmp    w13, #92
    BEQ LBB5_896
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x39000b2d // strb    w13, [x25, #2]
    WORD $0x8b1802ed // add    x13, x23, x24
    WORD $0x39400df9 // ldrb    w25, [x15, #3]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_897
    WORD $0x39000db9 // strb    w25, [x13, #3]
    WORD $0x394011f9 // ldrb    w25, [x15, #4]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_898
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x390011b9 // strb    w25, [x13, #4]
    WORD $0x8b1802ed // add    x13, x23, x24
    WORD $0x394015f9 // ldrb    w25, [x15, #5]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_899
    WORD $0x390015b9 // strb    w25, [x13, #5]
    WORD $0x394019f9 // ldrb    w25, [x15, #6]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_900
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x390019b9 // strb    w25, [x13, #6]
    WORD $0x8b1802ed // add    x13, x23, x24
    WORD $0x39401df9 // ldrb    w25, [x15, #7]
    WORD $0x7101733f // cmp    w25, #92
    BEQ LBB5_901
    WORD $0x39001db9 // strb    w25, [x13, #7]
    WORD $0x91002318 // add    x24, x24, #8
    WORD $0x394021ed // ldrb    w13, [x15, #8]
    WORD $0x710171bf // cmp    w13, #92
    BNE LBB5_886
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x8b1802f9 // add    x25, x23, x24
    WORD $0xd10005fa // sub    x26, x15, #1
    B LBB5_902
LBB5_895:
    WORD $0x9100074f // add    x15, x26, #1
    WORD $0x91000739 // add    x25, x25, #1
    B LBB5_902
LBB5_896:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x8b1802f6 // add    x22, x23, x24
    WORD $0x910005ba // add    x26, x13, #1
    WORD $0x910009af // add    x15, x13, #2
    WORD $0x91000ad9 // add    x25, x22, #2
    B LBB5_902
LBB5_897:
    WORD $0x910009fa // add    x26, x15, #2
    WORD $0x91000def // add    x15, x15, #3
    WORD $0x91000db9 // add    x25, x13, #3
    B LBB5_902
LBB5_898:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x8b1802f6 // add    x22, x23, x24
    WORD $0x91000dba // add    x26, x13, #3
    WORD $0x910011af // add    x15, x13, #4
    WORD $0x910012d9 // add    x25, x22, #4
    B LBB5_902
LBB5_899:
    WORD $0x910011fa // add    x26, x15, #4
    WORD $0x910015ef // add    x15, x15, #5
    WORD $0x910015b9 // add    x25, x13, #5
    B LBB5_902
LBB5_900:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x8b1802f6 // add    x22, x23, x24
    WORD $0x910015ba // add    x26, x13, #5
    WORD $0x910019af // add    x15, x13, #6
    WORD $0x91001ad9 // add    x25, x22, #6
    B LBB5_902
LBB5_901:
    WORD $0x910019fa // add    x26, x15, #6
    WORD $0x91001def // add    x15, x15, #7
    WORD $0x91001db9 // add    x25, x13, #7
LBB5_902:
    WORD $0x9100075a // add    x26, x26, #1
    WORD $0xaa1b03fd // mov    x29, x27
    B LBB5_848
LBB5_903:
    WORD $0x2a0f03ed // mov    w13, w15
LBB5_904:
    WORD $0x3800172d // strb    w13, [x25], #1
    WORD $0xaa1903f7 // mov    x23, x25
    B LBB5_871
LBB5_905:
    WORD $0x53067ded // lsr    w13, w15, #6
    WORD $0x52801018 // mov    w24, #128
    WORD $0x321a05ad // orr    w13, w13, #0xc0
    WORD $0x330015f8 // bfxil    w24, w15, #0, #6
    WORD $0xd1000aef // sub    x15, x23, #2
    WORD $0x381fc2ed // sturb    w13, [x23, #-4]
    WORD $0x381fd2f8 // sturb    w24, [x23, #-3]
    B LBB5_870
LBB5_906:
    WORD $0xd10006ed // sub    x13, x23, #1
    WORD $0x781fc2e7 // sturh    w7, [x23, #-4]
    WORD $0x381fe2f3 // sturb    w19, [x23, #-2]
    WORD $0xaa0d03f7 // mov    x23, x13
    B LBB5_871
LBB5_907:
    WORD $0x0b0f29ad // add    w13, w13, w15, lsl #10
    WORD $0x52801019 // mov    w25, #128
    WORD $0x0b1401ad // add    w13, w13, w20
    WORD $0x5280101a // mov    w26, #128
    WORD $0x53127daf // lsr    w15, w13, #18
    WORD $0x5280101b // mov    w27, #128
    WORD $0x321c0def // orr    w15, w15, #0xf0
    WORD $0x3300171b // bfxil    w27, w24, #0, #6
    WORD $0x330c45b9 // bfxil    w25, w13, #12, #6
    WORD $0x33062dba // bfxil    w26, w13, #6, #6
    WORD $0x381fc2ef // sturb    w15, [x23, #-4]
    WORD $0x381fd2f9 // sturb    w25, [x23, #-3]
    WORD $0x381fe2fa // sturb    w26, [x23, #-2]
    WORD $0x381ff2fb // sturb    w27, [x23, #-1]
    B LBB5_871
LBB5_908:
    WORD $0x92800031 // mov    x17, #-2
    WORD $0x4b1103e2 // neg    w2, w17
    B LBB5_341
LBB5_909:
    WORD $0x528000e2 // mov    w2, #7
    WORD $0xaa1103ef // mov    x15, x17
    B LBB5_341
LBB5_910:
    WORD $0x2538c3e2 // mov    z2.b, #31
LBB5_911:
    WORD $0xa400a1e3 // ld1b    { z3.b }, p0/z, [x15]
    WORD $0x2400a061 // cmpeq    p1.b, p0/z, z3.b, z0.b
    WORD $0x2401a063 // cmpeq    p3.b, p0/z, z3.b, z1.b
    WORD $0x2402a062 // cmpeq    p2.b, p0/z, z3.b, z2.b
    WORD $0x25834c64 // mov    p4.b, p3.b
    WORD $0x25c24025 // orrs    p5.b, p0/z, p1.b, p2.b
    BEQ LBB5_913
    WORD $0x259040a4 // brkb    p4.b, p0/z, p5.b
    WORD $0x25034084 // and    p4.b, p0/z, p4.b, p3.b
LBB5_913:
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_920
    WORD $0x2550c060 // ptest    p0, p3.b
    BEQ LBB5_917
    WORD $0x25904063 // brkb    p3.b, p0/z, p3.b
    WORD $0x25414064 // ands    p4.b, p0/z, p3.b, p1.b
    BNE LBB5_847
    WORD $0x25024063 // and    p3.b, p0/z, p3.b, p2.b
    B LBB5_918
LBB5_917:
    WORD $0x25824843 // mov    p3.b, p2.b
    WORD $0x2550c020 // ptest    p0, p1.b
    BNE LBB5_847
LBB5_918:
    WORD $0x2550c060 // ptest    p0, p3.b
    BNE LBB5_1683
    WORD $0x910081ef // add    x15, x15, #32
    B LBB5_911
LBB5_920:
    WORD $0x25904061 // brkb    p1.b, p0/z, p3.b
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0x52800092 // mov    w18, #4
    WORD $0x8b0f01ad // add    x13, x13, x15
    WORD $0x910005af // add    x15, x13, #1
    WORD $0xaa3103ed // mvn    x13, x17
    WORD $0x8b0d01f1 // add    x17, x15, x13
    WORD $0xaa0f03e0 // mov    x0, x15
    TST $(1<<63), R17
    BNE LBB5_1684
LBB5_921:
    WORD $0xaa0003ef // mov    x15, x0
    WORD $0x384015ed // ldrb    w13, [x15], #1
    WORD $0x710081bf // cmp    w13, #32
    BHI LBB5_932
    WORD $0x5280002e // mov    w14, #1
    WORD $0xd284c001 // mov    x1, #9728
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0xf2c00021 // movk    x1, #1, lsl #32
    WORD $0xea0101df // tst    x14, x1
    BEQ LBB5_932
    WORD $0x3940040d // ldrb    w13, [x0, #1]
    WORD $0x9100080f // add    x15, x0, #2
    WORD $0x710081bf // cmp    w13, #32
    BHI LBB5_987
    WORD $0x5280002e // mov    w14, #1
    WORD $0xd284c000 // mov    x0, #9728
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0xf2c00020 // movk    x0, #1, lsl #32
    WORD $0xea0001df // tst    x14, x0
    BEQ LBB5_987
    WORD $0xf9404900 // ldr    x0, [x8, #144]
    WORD $0xcb0001ed // sub    x13, x15, x0
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_928
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x9acd21ed // lsl    x13, x15, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_931
    WORD $0x9101000f // add    x15, x0, #64
LBB5_928:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R1
    WORD $0xd10101e0 // sub    x0, x15, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00023 // ldr    q3, [x1, :lo12:.LCPI5_2]
LBB5_929:
    WORD $0xadc21404 // ldp    q4, q5, [x0, #64]!
    WORD $0x4e211c90 // and    v16.16b, v4.16b, v1.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411c06 // ldp    q6, q7, [x0, #32]
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
    WORD $0x1e2600c1 // fmov    w1, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa0f81cd // orr    x13, x14, x15, lsl #32
    WORD $0xaa01c1ad // orr    x13, x13, x1, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_929
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa9093500 // stp    x0, x13, [x8, #144]
LBB5_931:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d000f // add    x15, x0, x13
    WORD $0x384015ed // ldrb    w13, [x15], #1
LBB5_932:
    WORD $0x7100e9bf // cmp    w13, #58
    BNE LBB5_988
LBB5_933:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0xaa108250 // orr    x16, x18, x16, lsl #32
    WORD $0xf90005b1 // str    x17, [x13, #8]
    WORD $0xf940510e // ldr    x14, [x8, #160]
    WORD $0xf90001b0 // str    x16, [x13]
    WORD $0xf9406111 // ldr    x17, [x8, #192]
    WORD $0x910041c0 // add    x0, x14, #16
    WORD $0x910081ce // add    x14, x14, #32
    WORD $0xeb1101df // cmp    x14, x17
    WORD $0xf9005100 // str    x0, [x8, #160]
    BHI LBB5_1469
    WORD $0xaa0f03f0 // mov    x16, x15
LBB5_935:
    WORD $0x384015f2 // ldrb    w18, [x15], #1
    WORD $0x7100825f // cmp    w18, #32
    BHI LBB5_946
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ad221ad // lsl    x13, x13, x18
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_946
    WORD $0x39400612 // ldrb    w18, [x16, #1]
    WORD $0x910005ef // add    x15, x15, #1
    WORD $0x7100825f // cmp    w18, #32
    BHI LBB5_968
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ad221ad // lsl    x13, x13, x18
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_968
    WORD $0xf9404910 // ldr    x16, [x8, #144]
    WORD $0xcb1001ed // sub    x13, x15, x16
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_942
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x9acd21ed // lsl    x13, x15, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_945
    WORD $0x9101020f // add    x15, x16, #64
LBB5_942:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R17
    WORD $0xd10101f0 // sub    x16, x15, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00223 // ldr    q3, [x17, :lo12:.LCPI5_2]
LBB5_943:
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
    BEQ LBB5_943
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa9093510 // stp    x16, x13, [x8, #144]
LBB5_945:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d020f // add    x15, x16, x13
    WORD $0x384015f2 // ldrb    w18, [x15], #1
LBB5_946:
    WORD $0xaa2a03ed // mvn    x13, x10
    WORD $0x528000c2 // mov    w2, #6
    WORD $0x8b0f01b0 // add    x16, x13, x15
    WORD $0x71016a5f // cmp    w18, #90
    BGT LBB5_969
LBB5_947:
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x5100c24d // sub    w13, w18, #48
    WORD $0x710029bf // cmp    w13, #10
    BHS LBB5_989
    WORD $0x5200002d // eor    w13, w1, #0x1
    WORD $0x3941c10e // ldrb    w14, [x8, #112]
    WORD $0x934001a6 // sbfx    x6, x13, #0, #1
    WORD $0xcb0d01e5 // sub    x5, x15, x13
    TST $(1<<1), R14
    BNE LBB5_992
LBB5_949:
    WORD $0x394000b3 // ldrb    w19, [x5]
    WORD $0x7100c26d // subs    w13, w19, #48
    BNE LBB5_962
    WORD $0xaa0503f1 // mov    x17, x5
    WORD $0x38401e2d // ldrb    w13, [x17, #1]!
    WORD $0x7100b9bf // cmp    w13, #46
    BEQ LBB5_1141
    WORD $0xaa1f03e3 // mov    x3, xzr
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x710115bf // cmp    w13, #69
    BEQ LBB5_953
    WORD $0x710195bf // cmp    w13, #101
    BNE LBB5_1159
LBB5_953:
    WORD $0x2a0403e7 // mov    w7, w4
LBB5_954:
    WORD $0xaa1103e0 // mov    x0, x17
    WORD $0x12800012 // mov    w18, #-1
    WORD $0x38401c13 // ldrb    w19, [x0, #1]!
    WORD $0x7100b67f // cmp    w19, #45
    BEQ LBB5_956
    WORD $0x52800032 // mov    w18, #1
    WORD $0x7100ae7f // cmp    w19, #43
    BNE LBB5_957
LBB5_956:
    WORD $0x38402e33 // ldrb    w19, [x17, #2]!
    WORD $0xaa1103e0 // mov    x0, x17
LBB5_957:
    WORD $0x52800062 // mov    w2, #3
    WORD $0x5100c26d // sub    w13, w19, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1607
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x52800151 // mov    w17, #10
LBB5_959:
    WORD $0x8b0d000e // add    x14, x0, x13
    WORD $0x1b114c42 // madd    w2, w2, w17, w19
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x5100c042 // sub    w2, w2, #48
    WORD $0x394005d3 // ldrb    w19, [x14, #1]
    WORD $0x5100c26e // sub    w14, w19, #48
    WORD $0x710029df // cmp    w14, #10
    BLO LBB5_959
    WORD $0x8b0d0011 // add    x17, x0, x13
    WORD $0xd10005ae // sub    x14, x13, #1
    WORD $0xf10025df // cmp    x14, #9
    BHS LBB5_1608
LBB5_961:
    WORD $0x1b121c47 // madd    w7, w2, w18, w7
    B LBB5_1171
LBB5_962:
    WORD $0x52800062 // mov    w2, #3
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_1035
    WORD $0xaa1f03e3 // mov    x3, xzr
    WORD $0xaa1f03f2 // mov    x18, xzr
    WORD $0x5280014d // mov    w13, #10
LBB5_964:
    WORD $0x8b1200ae // add    x14, x5, x18
    WORD $0x9b0d7c71 // mul    x17, x3, x13
    WORD $0x91000652 // add    x18, x18, #1
    WORD $0x8b334231 // add    x17, x17, w19, uxtw
    WORD $0x394005d3 // ldrb    w19, [x14, #1]
    WORD $0xd100c223 // sub    x3, x17, #48
    WORD $0x5100c26e // sub    w14, w19, #48
    WORD $0x710029df // cmp    w14, #10
    BLO LBB5_964
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x2a1f03e7 // mov    w7, wzr
    WORD $0x8b1200b1 // add    x17, x5, x18
    WORD $0xd100064d // sub    x13, x18, #1
    WORD $0xf1004dbf // cmp    x13, #19
    BHS LBB5_1609
LBB5_966:
    WORD $0x7100ba7f // cmp    w19, #46
    BNE LBB5_1152
    WORD $0x38401e20 // ldrb    w0, [x17, #1]!
    WORD $0x52800062 // mov    w2, #3
    WORD $0xaa1103f3 // mov    x19, x17
    WORD $0x5100c00d // sub    w13, w0, #48
    WORD $0x710029bf // cmp    w13, #10
    BLO LBB5_1147
    B LBB5_1598
LBB5_968:
    WORD $0xaa2a03ed // mvn    x13, x10
    WORD $0x528000c2 // mov    w2, #6
    WORD $0x8b0f01b0 // add    x16, x13, x15
    WORD $0x71016a5f // cmp    w18, #90
    BLE LBB5_947
LBB5_969:
    WORD $0x910005f1 // add    x17, x15, #1
    WORD $0x7101b65f // cmp    w18, #109
    BLE LBB5_1028
    WORD $0x7101ba5f // cmp    w18, #110
    BEQ LBB5_1051
    WORD $0x7101d25f // cmp    w18, #116
    BEQ LBB5_1047
    WORD $0x7101ee5f // cmp    w18, #123
    BNE LBB5_341
    WORD $0xf940550b // ldr    x11, [x8, #168]
    WORD $0xaa0803f2 // mov    x18, x8
    WORD $0xf900040b // str    x11, [x0, #8]
    WORD $0x528000cb // mov    w11, #6
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
    BHI LBB5_985
    WORD $0x5280002e // mov    w14, #1
    WORD $0xd284c000 // mov    x0, #9728
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0xf2c00020 // movk    x0, #1, lsl #32
    WORD $0xea0001df // tst    x14, x0
    BEQ LBB5_985
    WORD $0x394005ed // ldrb    w13, [x15, #1]
    WORD $0x91000631 // add    x17, x17, #1
    WORD $0x710081bf // cmp    w13, #32
    BHI LBB5_1241
    WORD $0x5280002e // mov    w14, #1
    WORD $0xd284c00f // mov    x15, #9728
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0xf2c0002f // movk    x15, #1, lsl #32
    WORD $0xea0f01df // tst    x14, x15
    BEQ LBB5_1241
    WORD $0xf940490f // ldr    x15, [x8, #144]
    WORD $0xcb0f022d // sub    x13, x17, x15
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_981
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x92800011 // mov    x17, #-1
    WORD $0x9acd222d // lsl    x13, x17, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_984
    WORD $0x910101f1 // add    x17, x15, #64
LBB5_981:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R0
    WORD $0xd101022f // sub    x15, x17, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00003 // ldr    q3, [x0, :lo12:.LCPI5_2]
LBB5_982:
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
    BEQ LBB5_982
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa909350f // stp    x15, x13, [x8, #144]
LBB5_984:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d01f1 // add    x17, x15, x13
    WORD $0x3840162d // ldrb    w13, [x17], #1
LBB5_985:
    WORD $0x7101f5bf // cmp    w13, #125
    BNE LBB5_1242
LBB5_986:
    WORD $0xb940cd0b // ldr    w11, [x8, #204]
    WORD $0x8b03104d // add    x13, x2, x3, lsl #4
    WORD $0x1100056b // add    w11, w11, #1
    WORD $0xb900cd0b // str    w11, [x8, #204]
    B LBB5_1069
LBB5_987:
    WORD $0x7100e9bf // cmp    w13, #58
    BEQ LBB5_933
LBB5_988:
    WORD $0x52800102 // mov    w2, #8
    B LBB5_341
LBB5_989:
    WORD $0x71008a5f // cmp    w18, #34
    BEQ LBB5_1071
    WORD $0x7100b65f // cmp    w18, #45
    BNE LBB5_341
    WORD $0x52800021 // mov    w1, #1
    WORD $0x5200002d // eor    w13, w1, #0x1
    WORD $0x3941c10e // ldrb    w14, [x8, #112]
    WORD $0x934001a6 // sbfx    x6, x13, #0, #1
    WORD $0xcb0d01e5 // sub    x5, x15, x13
    TST $(1<<1), R14
    BEQ LBB5_949
LBB5_992:
    WORD $0xcb05012d // sub    x13, x9, x5
    WORD $0x92800004 // mov    x4, #-1
    WORD $0xeb0601b1 // subs    x17, x13, x6
    BEQ LBB5_1226
    WORD $0x394000ad // ldrb    w13, [x5]
    WORD $0x924000c7 // and    x7, x6, #0x1
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_997
    WORD $0x2a1f03f2 // mov    w18, wzr
    WORD $0x52800024 // mov    w4, #1
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf100063f // cmp    x17, #1
    BEQ LBB5_1227
    WORD $0x394004ad // ldrb    w13, [x5, #1]
    WORD $0x2a1f03f2 // mov    w18, wzr
    WORD $0x52800024 // mov    w4, #1
    WORD $0x5100b9ad // sub    w13, w13, #46
    WORD $0x7100ddbf // cmp    w13, #55
    BHI LBB5_1246
    WORD $0x5280002e // mov    w14, #1
    WORD $0xb20903e2 // mov    x2, #36028797027352576
    WORD $0x9acd21ce // lsl    x14, x14, x13
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2800022 // movk    x2, #1
    WORD $0xea0201df // tst    x14, x2
    BEQ LBB5_1227
LBB5_997:
    WORD $0x92800012 // mov    x18, #-1
    WORD $0xf100423f // cmp    x17, #16
    BLO LBB5_1606
    ADR LCPI5_1, R13
    ADR LCPI5_2, R14
    WORD $0x4f01e5c0 // movi    v0.16b, #46
    WORD $0xaa1f03e4 // mov    x4, xzr
    WORD $0x4f01e561 // movi    v1.16b, #43
    WORD $0x92800002 // mov    x2, #-1
    WORD $0x4f01e5a2 // movi    v2.16b, #45
    WORD $0x3dc001a5 // ldr    q5, [x13, :lo12:.LCPI5_1]
    WORD $0x4f06e603 // movi    v3.16b, #208
    WORD $0x3dc001c7 // ldr    q7, [x14, :lo12:.LCPI5_2]
    WORD $0x4f00e544 // movi    v4.16b, #10
    WORD $0x12800013 // mov    w19, #-1
    WORD $0x4f06e7e6 // movi    v6.16b, #223
    WORD $0x92800003 // mov    x3, #-1
    WORD $0x4f02e4b0 // movi    v16.16b, #69
LBB5_999:
    WORD $0x3ce468b1 // ldr    q17, [x5, x4]
    WORD $0x6e208e32 // cmeq    v18.16b, v17.16b, v0.16b
    WORD $0x6e218e33 // cmeq    v19.16b, v17.16b, v1.16b
    WORD $0x6e228e34 // cmeq    v20.16b, v17.16b, v2.16b
    WORD $0x4e238635 // add    v21.16b, v17.16b, v3.16b
    WORD $0x4e261e31 // and    v17.16b, v17.16b, v6.16b
    WORD $0x6e308e31 // cmeq    v17.16b, v17.16b, v16.16b
    WORD $0x6e353495 // cmhi    v21.16b, v4.16b, v21.16b
    WORD $0x4eb31e93 // orr    v19.16b, v20.16b, v19.16b
    WORD $0x4eb11e54 // orr    v20.16b, v18.16b, v17.16b
    WORD $0x4eb41eb4 // orr    v20.16b, v21.16b, v20.16b
    WORD $0x4eb31e94 // orr    v20.16b, v20.16b, v19.16b
    WORD $0x4e251e94 // and    v20.16b, v20.16b, v5.16b
    WORD $0x4e070294 // tbl    v20.16b, { v20.16b }, v7.16b
    WORD $0x4e251e52 // and    v18.16b, v18.16b, v5.16b
    WORD $0x4e251e31 // and    v17.16b, v17.16b, v5.16b
    WORD $0x4e251e73 // and    v19.16b, v19.16b, v5.16b
    WORD $0x4e71ba94 // addv    h20, v20.8h
    WORD $0x4e070252 // tbl    v18.16b, { v18.16b }, v7.16b
    WORD $0x4e070231 // tbl    v17.16b, { v17.16b }, v7.16b
    WORD $0x4e070273 // tbl    v19.16b, { v19.16b }, v7.16b
    WORD $0x1e26028d // fmov    w13, s20
    WORD $0x4e71ba52 // addv    h18, v18.8h
    WORD $0x4e71ba31 // addv    h17, v17.8h
    WORD $0x2a2d03ed // mvn    w13, w13
    WORD $0x4e71ba73 // addv    h19, v19.8h
    WORD $0x32103dad // orr    w13, w13, #0xffff0000
    WORD $0x5ac001ad // rbit    w13, w13
    WORD $0x1e260257 // fmov    w23, s18
    WORD $0x1e260236 // fmov    w22, s17
    WORD $0x5ac011b4 // clz    w20, w13
    WORD $0x1e260275 // fmov    w21, s19
    WORD $0x7100429f // cmp    w20, #16
    BEQ LBB5_1001
    WORD $0x1ad4226d // lsl    w13, w19, w20
    WORD $0x0a2d02f7 // bic    w23, w23, w13
    WORD $0x0a2d02d6 // bic    w22, w22, w13
    WORD $0x0a2d02b5 // bic    w21, w21, w13
LBB5_1001:
    WORD $0x510006ed // sub    w13, w23, #1
    WORD $0x6a1701ad // ands    w13, w13, w23
    BNE LBB5_1318
    WORD $0x510006cd // sub    w13, w22, #1
    WORD $0x6a1601ad // ands    w13, w13, w22
    BNE LBB5_1318
    WORD $0x510006ad // sub    w13, w21, #1
    WORD $0x6a1501ad // ands    w13, w13, w21
    BNE LBB5_1318
    CMP $0, R23
    BEQ LBB5_1007
    WORD $0x5ac002ed // rbit    w13, w23
    WORD $0xb100065f // cmn    x18, #1
    WORD $0x5ac011ad // clz    w13, w13
    BNE LBB5_1442
    WORD $0x8b0d0092 // add    x18, x4, x13
LBB5_1007:
    CMP $0, R22
    BEQ LBB5_1010
    WORD $0x5ac002cd // rbit    w13, w22
    WORD $0xb100047f // cmn    x3, #1
    WORD $0x5ac011ad // clz    w13, w13
    BNE LBB5_1442
    WORD $0x8b0d0083 // add    x3, x4, x13
LBB5_1010:
    CMP $0, R21
    BEQ LBB5_1013
    WORD $0x5ac002ad // rbit    w13, w21
    WORD $0xb100045f // cmn    x2, #1
    WORD $0x5ac011ad // clz    w13, w13
    BNE LBB5_1442
    WORD $0x8b0d0082 // add    x2, x4, x13
LBB5_1013:
    WORD $0x7100429f // cmp    w20, #16
    BNE LBB5_1037
    WORD $0xd1004231 // sub    x17, x17, #16
    WORD $0x91004084 // add    x4, x4, #16
    WORD $0xf1003e3f // cmp    x17, #15
    BHI LBB5_999
    WORD $0x8b07012d // add    x13, x9, x7
    WORD $0x8b0400a7 // add    x7, x5, x4
    WORD $0xcb0f01ad // sub    x13, x13, x15
    WORD $0xcb0601ad // sub    x13, x13, x6
    WORD $0xeb0401bf // cmp    x13, x4
    BEQ LBB5_1038
LBB5_1016:
    WORD $0x8b0601ed // add    x13, x15, x6
    WORD $0xaa2703ee // mvn    x14, x7
    WORD $0x8b1100ef // add    x15, x7, x17
    WORD $0x8b0d01c4 // add    x4, x14, x13
    WORD $0xcb0500e6 // sub    x6, x7, x5
    WORD $0xaa0703f3 // mov    x19, x7
    B LBB5_1019
