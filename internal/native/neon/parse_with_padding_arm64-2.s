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
