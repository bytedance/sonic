    WORD $0x2a1903f5 // mov    w21, w25
    B LBB5_2045
LBB5_2042:
    CMP $0, R27
    BEQ LBB5_2419
LBB5_2043:
    WORD $0x2a1b03f9 // mov    w25, w27
    WORD $0x2a1b03f5 // mov    w21, w27
LBB5_2044:
    WORD $0x0b050285 // add    w5, w20, w5
    WORD $0x7100007f // cmp    w3, #0
    BLE LBB5_2103
LBB5_2045:
    WORD $0x7100207f // cmp    w3, #8
    BLS LBB5_2048
    WORD $0x52800374 // mov    w20, #27
    CMP $0, R21
    BEQ LBB5_2044
    WORD $0x12800358 // mov    w24, #-27
    B LBB5_2074
LBB5_2048:
    ADR POW_TAB, R13
    WORD $0x910001ad // add    x13, x13, :lo12:POW_TAB
    WORD $0xb86359b4 // ldr    w20, [x13, w3, uxtw #2]
    CMP $0, R21
    BEQ LBB5_2044
    WORD $0x4b1403f8 // neg    w24, w20
    WORD $0x3100f71f // cmn    w24, #61
    BLS LBB5_2053
    B LBB5_2074
LBB5_2050:
    WORD $0x2a1f03e3 // mov    w3, wzr
LBB5_2051:
    WORD $0x2a1f03f9 // mov    w25, wzr
LBB5_2052:
    WORD $0x1100f2d8 // add    w24, w22, #60
    WORD $0x2a1903f5 // mov    w21, w25
    WORD $0x3101e2df // cmn    w22, #120
    BGE LBB5_2073
LBB5_2053:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03fa // mov    x26, xzr
    WORD $0x2a1803f6 // mov    w22, w24
    WORD $0x0ab57eb7 // bic    w23, w21, w21, asr #31
LBB5_2054:
    WORD $0xeb0d02ff // cmp    x23, x13
    BEQ LBB5_2057
    WORD $0x386d6a4e // ldrb    w14, [x18, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b073b4e // madd    x14, x26, x7, x14
    WORD $0xd100c1da // sub    x26, x14, #48
    WORD $0xd37cff4e // lsr    x14, x26, #60
    CMP $0, R14
    BEQ LBB5_2054
    WORD $0xaa1a03f8 // mov    x24, x26
    WORD $0x2a0d03f7 // mov    w23, w13
    B LBB5_2059
LBB5_2057:
    CMP $0, R26
    BEQ LBB5_2051
LBB5_2058:
    WORD $0x8b1a0b4d // add    x13, x26, x26, lsl #2
    WORD $0x110006f7 // add    w23, w23, #1
    WORD $0xd37ff9b8 // lsl    x24, x13, #1
    WORD $0xeb13035f // cmp    x26, x19
    WORD $0xaa1803fa // mov    x26, x24
    BLO LBB5_2058
LBB5_2059:
    WORD $0x6b1502ff // cmp    w23, w21
    BGE LBB5_2063
    WORD $0x2a1703ed // mov    w13, w23
    WORD $0xaa1f03e6 // mov    x6, xzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x93407f35 // sxtw    x21, w25
    WORD $0x8b0d0259 // add    x25, x18, x13
LBB5_2061:
    WORD $0xd37cff0e // lsr    x14, x24, #60
    WORD $0x9240ef18 // and    x24, x24, #0xfffffffffffffff
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x38266a4e // strb    w14, [x18, x6]
    WORD $0x38666b2e // ldrb    w14, [x25, x6]
    WORD $0x910004c6 // add    x6, x6, #1
    WORD $0x9b073b0e // madd    x14, x24, x7, x14
    WORD $0xd100c1d8 // sub    x24, x14, #48
    WORD $0x8b0601ae // add    x14, x13, x6
    WORD $0xeb1501df // cmp    x14, x21
    BLT LBB5_2061
    WORD $0x2a0603f9 // mov    w25, w6
    CMP $0, R24
    BNE LBB5_2065
    B LBB5_2067
LBB5_2063:
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_2065
LBB5_2064:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x9240ef0d // and    x13, x24, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d8 // lsl    x24, x14, #1
    CMP $0, R13
    BEQ LBB5_2067
LBB5_2065:
    WORD $0x93407f2d // sxtw    x13, w25
    WORD $0xd37cff0e // lsr    x14, x24, #60
    WORD $0xeb0d005f // cmp    x2, x13
    BLS LBB5_2064
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x11000739 // add    w25, w25, #1
    WORD $0x382d6a4e // strb    w14, [x18, x13]
    WORD $0x9240ef0d // and    x13, x24, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d8 // lsl    x24, x14, #1
    CMP $0, R13
    BNE LBB5_2065
LBB5_2067:
    WORD $0x4b17006d // sub    w13, w3, w23
    WORD $0x7100073f // cmp    w25, #1
    WORD $0x110005a3 // add    w3, w13, #1
    BLT LBB5_2072
    WORD $0x2a1903ed // mov    w13, w25
    WORD $0x8b1201ae // add    x14, x13, x18
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_2052
LBB5_2069:
    WORD $0xf10005b9 // subs    x25, x13, #1
    BLS LBB5_2050
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d4a4e // ldrb    w14, [x18, w13, uxtw]
    WORD $0xaa1903ed // mov    x13, x25
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2069
    B LBB5_2052
LBB5_2072:
    CMP $0, R25
    BNE LBB5_2052
    B LBB5_2050
LBB5_2073:
    WORD $0x2a1903f5 // mov    w21, w25
LBB5_2074:
    WORD $0xaa1f03fb // mov    x27, xzr
    WORD $0xaa1f03f7 // mov    x23, xzr
    WORD $0x4b1803f6 // neg    w22, w24
    WORD $0x0ab57eba // bic    w26, w21, w21, asr #31
LBB5_2075:
    WORD $0xeb1b035f // cmp    x26, x27
    BEQ LBB5_2082
    WORD $0x387b6a4d // ldrb    w13, [x18, x27]
    WORD $0x9100077b // add    x27, x27, #1
    WORD $0x9b0736ed // madd    x13, x23, x7, x13
    WORD $0xd100c1b7 // sub    x23, x13, #48
    WORD $0x9ad626ed // lsr    x13, x23, x22
    CMP $0, R13
    BEQ LBB5_2075
    WORD $0x2a1b03fa // mov    w26, w27
LBB5_2078:
    WORD $0x9ad6208d // lsl    x13, x4, x22
    WORD $0x6b15035f // cmp    w26, w21
    WORD $0xaa2d03f8 // mvn    x24, x13
    BGE LBB5_2086
    WORD $0x2a1a03ed // mov    w13, w26
    WORD $0xaa1f03fb // mov    x27, xzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x93407f26 // sxtw    x6, w25
    WORD $0x8b0d0255 // add    x21, x18, x13
LBB5_2080:
    WORD $0x9ad626ee // lsr    x14, x23, x22
    WORD $0x8a1802f7 // and    x23, x23, x24
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x383b6a4e // strb    w14, [x18, x27]
    WORD $0x387b6aae // ldrb    w14, [x21, x27]
    WORD $0x9100077b // add    x27, x27, #1
    WORD $0x9b073aee // madd    x14, x23, x7, x14
    WORD $0xd100c1d7 // sub    x23, x14, #48
    WORD $0x8b1b01ae // add    x14, x13, x27
    WORD $0xeb0601df // cmp    x14, x6
    BLT LBB5_2080
    B LBB5_2087
LBB5_2082:
    CMP $0, R23
    BEQ LBB5_2096
    WORD $0x9ad626ed // lsr    x13, x23, x22
    CMP $0, R13
    BEQ LBB5_2085
    WORD $0x9ad6208d // lsl    x13, x4, x22
    WORD $0x4b1a006e // sub    w14, w3, w26
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0xaa2d03f8 // mvn    x24, x13
    WORD $0x110005c3 // add    w3, w14, #1
    B LBB5_2089
LBB5_2085:
    WORD $0x8b170aed // add    x13, x23, x23, lsl #2
    WORD $0x1100075a // add    w26, w26, #1
    WORD $0xd37ff9b7 // lsl    x23, x13, #1
    WORD $0x9ad626ed // lsr    x13, x23, x22
    CMP $0, R13
    BEQ LBB5_2085
    B LBB5_2078
LBB5_2086:
    WORD $0x2a1f03fb // mov    w27, wzr
LBB5_2087:
    WORD $0x4b1a006d // sub    w13, w3, w26
    WORD $0x110005a3 // add    w3, w13, #1
    CMP $0, R23
    BNE LBB5_2089
    B LBB5_2091
LBB5_2088:
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x8a1802ed // and    x13, x23, x24
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d7 // lsl    x23, x14, #1
    CMP $0, R13
    BEQ LBB5_2091
LBB5_2089:
    WORD $0x9ad626ed // lsr    x13, x23, x22
    WORD $0x93407f66 // sxtw    x6, w27
    WORD $0xeb06005f // cmp    x2, x6
    BLS LBB5_2088
    WORD $0x1100c1ad // add    w13, w13, #48
    WORD $0x1100077b // add    w27, w27, #1
    WORD $0x38266a4d // strb    w13, [x18, x6]
    WORD $0x8a1802ed // and    x13, x23, x24
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d7 // lsl    x23, x14, #1
    CMP $0, R13
    BNE LBB5_2089
LBB5_2091:
    WORD $0x7100077f // cmp    w27, #1
    BLT LBB5_2042
    WORD $0x2a1b03ed // mov    w13, w27
    WORD $0x8b1201ae // add    x14, x13, x18
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_2043
LBB5_2093:
    WORD $0xf10005b5 // subs    x21, x13, #1
    BLS LBB5_2098
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d4a4e // ldrb    w14, [x18, w13, uxtw]
    WORD $0xaa1503ed // mov    x13, x21
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2093
    WORD $0x2a1503f9 // mov    w25, w21
    WORD $0x2a1503fb // mov    w27, w21
    B LBB5_2044
LBB5_2096:
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0x2a1f03f5 // mov    w21, wzr
    B LBB5_2044
LBB5_2097:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xcb0003fd // neg    x29, x0
    B LBB5_2385
LBB5_2098:
    WORD $0x510005b9 // sub    w25, w13, #1
LBB5_2099:
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0x0b050285 // add    w5, w20, w5
    B LBB5_2102
LBB5_2100:
    WORD $0x5ac001ad // rbit    w13, w13
    WORD $0xaa2203ee // mvn    x14, x2
    WORD $0x5ac011ad // clz    w13, w13
    WORD $0xcb0d01c2 // sub    x2, x14, x13
    B LBB5_1918
LBB5_2101:
    WORD $0x2a1f03e5 // mov    w5, wzr
LBB5_2102:
    WORD $0x2a1903fb // mov    w27, w25
LBB5_2103:
    WORD $0xb201e7f6 // mov    x22, #-7378697629483820647
    WORD $0xb202e7f3 // mov    x19, #-3689348814741910324
    WORD $0xf2933356 // movk    x22, #39322
    WORD $0xf90007fd // str    x29, [sp, #8]
    WORD $0xf29999b3 // movk    x19, #52429
    WORD $0x92800134 // mov    x20, #-10
    WORD $0xb2607ff5 // mov    x21, #-4294967296
    WORD $0xf2e03336 // movk    x22, #409, lsl #48
    WORD $0x52800157 // mov    w23, #10
    WORD $0x2a1903fa // mov    w26, w25
    WORD $0x2a1b03fd // mov    w29, w27
    B LBB5_2109
LBB5_2104:
    WORD $0x5100075a // sub    w26, w26, #1
LBB5_2105:
    WORD $0x2a1f03e3 // mov    w3, wzr
LBB5_2106:
    TST $(1<<31), R25
    BNE LBB5_2137
LBB5_2107:
    WORD $0x2a1a03fd // mov    w29, w26
    WORD $0x2a1a03fb // mov    w27, w26
LBB5_2108:
    WORD $0x4b1900a5 // sub    w5, w5, w25
LBB5_2109:
    TST $(1<<31), R3
    BNE LBB5_2112
    CMP $0, R3
    BNE LBB5_2196
    WORD $0x3940024d // ldrb    w13, [x18]
    WORD $0x7100d5bf // cmp    w13, #53
    BLO LBB5_2115
    B LBB5_2196
LBB5_2112:
    WORD $0x3100207f // cmn    w3, #8
    BHS LBB5_2115
    WORD $0x52800379 // mov    w25, #27
    CMP $0, R29
    BEQ LBB5_2179
    WORD $0x2a1d03fb // mov    w27, w29
    B LBB5_2116
LBB5_2115:
    WORD $0x4b0303ed // neg    w13, w3
    ADR POW_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:POW_TAB
    WORD $0xb86d59d9 // ldr    w25, [x14, w13, uxtw #2]
    CMP $0, R27
    BEQ LBB5_2108
LBB5_2116:
    WORD $0x52800d0e // mov    w14, #104
    ADR LSHIFT_TAB, R13
    WORD $0x910001ad // add    x13, x13, :lo12:LSHIFT_TAB
    WORD $0x93407f7d // sxtw    x29, w27
    WORD $0x9bae372d // umaddl    x13, w25, w14, x13
    WORD $0x2a1903f8 // mov    w24, w25
    WORD $0xaa1203e6 // mov    x6, x18
    WORD $0xaa1d03e7 // mov    x7, x29
    WORD $0xb84045bb // ldr    w27, [x13], #4
    WORD $0xaa0d03e4 // mov    x4, x13
LBB5_2117:
    WORD $0x3840148e // ldrb    w14, [x4], #1
    CMP $0, R14
    BEQ LBB5_2122
    WORD $0x394000de // ldrb    w30, [x6]
    WORD $0x6b0e03df // cmp    w30, w14
    BNE LBB5_2165
    WORD $0xf10004e7 // subs    x7, x7, #1
    WORD $0x910004c6 // add    x6, x6, #1
    BNE LBB5_2117
    WORD $0x387d69ad // ldrb    w13, [x13, x29]
    CMP $0, R13
    BEQ LBB5_2122
LBB5_2121:
    WORD $0x5100077b // sub    w27, w27, #1
LBB5_2122:
    WORD $0x710007bf // cmp    w29, #1
    BLT LBB5_2132
    WORD $0x0b1d036d // add    w13, w27, w29
    WORD $0x92407fae // and    x14, x29, #0xffffffff
    WORD $0x93407da6 // sxtw    x6, w13
    WORD $0xaa1f03e4 // mov    x4, xzr
    WORD $0x93607dbd // sbfiz    x29, x13, #32, #32
    WORD $0xd10004cd // sub    x13, x6, #1
    WORD $0x910005de // add    x30, x14, #1
    B LBB5_2125
LBB5_2124:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x8b1503bd // add    x29, x29, x21
    WORD $0xd10004ed // sub    x13, x7, #1
    WORD $0xd10007de // sub    x30, x30, #1
    WORD $0xf10007df // cmp    x30, #1
    BLS LBB5_2127
LBB5_2125:
    WORD $0x51000bce // sub    w14, w30, #2
    WORD $0xaa0d03e7 // mov    x7, x13
    WORD $0xeb0201bf // cmp    x13, x2
    WORD $0x386e4a4e // ldrb    w14, [x18, w14, uxtw]
    WORD $0xd100c1ce // sub    x14, x14, #48
    WORD $0x9ad821ce // lsl    x14, x14, x24
    WORD $0x8b0401c6 // add    x6, x14, x4
    WORD $0x9bd37cce // umulh    x14, x6, x19
    WORD $0xd343fdc4 // lsr    x4, x14, #3
    WORD $0x9b14188e // madd    x14, x4, x20, x6
    BHS LBB5_2124
    WORD $0x1100c1cd // add    w13, w14, #48
    WORD $0x38276a4d // strb    w13, [x18, x7]
    WORD $0x8b1503bd // add    x29, x29, x21
    WORD $0xd10004ed // sub    x13, x7, #1
    WORD $0xd10007de // sub    x30, x30, #1
    WORD $0xf10007df // cmp    x30, #1
    BHI LBB5_2125
LBB5_2127:
    WORD $0xf10028df // cmp    x6, #10
    BLO LBB5_2132
    WORD $0x93407ced // sxtw    x13, w7
    WORD $0xd10005b8 // sub    x24, x13, #1
    B LBB5_2130
LBB5_2129:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0xd1000718 // sub    x24, x24, #1
    WORD $0xf100249f // cmp    x4, #9
    WORD $0xaa0d03e4 // mov    x4, x13
    BLS LBB5_2132
LBB5_2130:
    WORD $0x9bd37c8d // umulh    x13, x4, x19
    WORD $0xeb02031f // cmp    x24, x2
    WORD $0xd343fdad // lsr    x13, x13, #3
    WORD $0x9b1411ae // madd    x14, x13, x20, x4
    BHS LBB5_2129
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x38386a4e // strb    w14, [x18, x24]
    WORD $0xd1000718 // sub    x24, x24, #1
    WORD $0xf100249f // cmp    x4, #9
    WORD $0xaa0d03e4 // mov    x4, x13
    BHI LBB5_2130
LBB5_2132:
    WORD $0x0b1a036d // add    w13, w27, w26
    WORD $0x0b030363 // add    w3, w27, w3
    WORD $0xeb2dc05f // cmp    x2, w13, sxtw
    WORD $0x1a8281ba // csel    w26, w13, w2, hi
    WORD $0x7100075f // cmp    w26, #1
    BLT LBB5_2164
    WORD $0x8b12034d // add    x13, x26, x18
    WORD $0x385ff1ad // ldurb    w13, [x13, #-1]
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_2106
LBB5_2134:
    WORD $0xf100074d // subs    x13, x26, #1
    BLS LBB5_2104
    WORD $0x51000b4e // sub    w14, w26, #2
    WORD $0xaa0d03fa // mov    x26, x13
    WORD $0x386e4a4e // ldrb    w14, [x18, w14, uxtw]
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2134
    WORD $0x2a0d03fa // mov    w26, w13
    B LBB5_2106
LBB5_2137:
    WORD $0x3100f73f // cmn    w25, #61
    BHI LBB5_2166
    WORD $0x2a1903fb // mov    w27, w25
    B LBB5_2142
LBB5_2139:
    WORD $0x2a1f03e3 // mov    w3, wzr
LBB5_2140:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_2141:
    WORD $0x1100f364 // add    w4, w27, #60
    WORD $0x3101e37f // cmn    w27, #120
    WORD $0x2a0403fb // mov    w27, w4
    BGE LBB5_2167
LBB5_2142:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03f8 // mov    x24, xzr
    WORD $0x0aba7f5d // bic    w29, w26, w26, asr #31
LBB5_2143:
    WORD $0xeb0d03bf // cmp    x29, x13
    BEQ LBB5_2146
    WORD $0x386d6a4e // ldrb    w14, [x18, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b173b0e // madd    x14, x24, x23, x14
    WORD $0xd100c1d8 // sub    x24, x14, #48
    WORD $0xd37cff0e // lsr    x14, x24, #60
    CMP $0, R14
    BEQ LBB5_2143
    WORD $0xaa1803e4 // mov    x4, x24
    WORD $0x2a0d03fd // mov    w29, w13
    B LBB5_2148
LBB5_2146:
    CMP $0, R24
    BEQ LBB5_2140
LBB5_2147:
    WORD $0x8b180b0d // add    x13, x24, x24, lsl #2
    WORD $0x110007bd // add    w29, w29, #1
    WORD $0xd37ff9a4 // lsl    x4, x13, #1
    WORD $0xeb16031f // cmp    x24, x22
    WORD $0xaa0403f8 // mov    x24, x4
    BLO LBB5_2147
LBB5_2148:
    WORD $0x6b1a03bf // cmp    w29, w26
    BGE LBB5_2153
    WORD $0x2a1d03ed // mov    w13, w29
    WORD $0x93407f46 // sxtw    x6, w26
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0xaa1203e7 // mov    x7, x18
    WORD $0xcb0d00da // sub    x26, x6, x13
LBB5_2150:
    WORD $0xd37cfc8e // lsr    x14, x4, #60
    WORD $0x9240ec84 // and    x4, x4, #0xfffffffffffffff
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0xd10004c6 // sub    x6, x6, #1
    WORD $0xeb0601bf // cmp    x13, x6
    WORD $0x390000ee // strb    w14, [x7]
    WORD $0x386d68ee // ldrb    w14, [x7, x13]
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0x9b17388e // madd    x14, x4, x23, x14
    WORD $0xd100c1c4 // sub    x4, x14, #48
    BNE LBB5_2150
    CMP $0, R4
    BNE LBB5_2154
    B LBB5_2158
LBB5_2153:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_2154:
    B LBB5_2156
LBB5_2155:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x9240ec8d // and    x13, x4, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9c4 // lsl    x4, x14, #1
    CMP $0, R13
    BEQ LBB5_2158
LBB5_2156:
    WORD $0x93407f4d // sxtw    x13, w26
    WORD $0xd37cfc8e // lsr    x14, x4, #60
    WORD $0xeb0d005f // cmp    x2, x13
    BLS LBB5_2155
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x1100075a // add    w26, w26, #1
    WORD $0x382d6a4e // strb    w14, [x18, x13]
    WORD $0x9240ec8d // and    x13, x4, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9c4 // lsl    x4, x14, #1
    CMP $0, R13
    BNE LBB5_2156
LBB5_2158:
    WORD $0x4b1d006d // sub    w13, w3, w29
    WORD $0x7100075f // cmp    w26, #1
    WORD $0x110005a3 // add    w3, w13, #1
    BLT LBB5_2163
    WORD $0x2a1a03ed // mov    w13, w26
    WORD $0x8b1201ae // add    x14, x13, x18
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_2141
LBB5_2160:
    WORD $0xf10005ba // subs    x26, x13, #1
    BLS LBB5_2139
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d4a4e // ldrb    w14, [x18, w13, uxtw]
    WORD $0xaa1a03ed // mov    x13, x26
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2160
    B LBB5_2141
LBB5_2163:
    CMP $0, R26
    BNE LBB5_2141
    B LBB5_2139
LBB5_2164:
    CMP $0, R26
    BNE LBB5_2106
    B LBB5_2105
LBB5_2165:
    BLO LBB5_2121
    B LBB5_2122
LBB5_2166:
    WORD $0x2a1903e4 // mov    w4, w25
LBB5_2167:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03fe // mov    x30, xzr
    WORD $0x4b0403fb // neg    w27, w4
    WORD $0x0aba7f44 // bic    w4, w26, w26, asr #31
LBB5_2168:
    WORD $0xeb0d009f // cmp    x4, x13
    BEQ LBB5_2175
    WORD $0x386d6a4e // ldrb    w14, [x18, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b173bce // madd    x14, x30, x23, x14
    WORD $0xd100c1de // sub    x30, x14, #48
    WORD $0x9adb27ce // lsr    x14, x30, x27
    CMP $0, R14
    BEQ LBB5_2168
    WORD $0x2a0d03e4 // mov    w4, w13
LBB5_2171:
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x6b1a009f // cmp    w4, w26
    WORD $0x9adb21ad // lsl    x13, x13, x27
    WORD $0xaa2d03f8 // mvn    x24, x13
    BGE LBB5_2180
    WORD $0x2a0403ed // mov    w13, w4
    WORD $0x93407f46 // sxtw    x6, w26
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0xaa1203e7 // mov    x7, x18
    WORD $0xcb0d00dd // sub    x29, x6, x13
LBB5_2173:
    WORD $0x9adb27ce // lsr    x14, x30, x27
    WORD $0x8a1803da // and    x26, x30, x24
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0xd10004c6 // sub    x6, x6, #1
    WORD $0xeb0601bf // cmp    x13, x6
    WORD $0x390000ee // strb    w14, [x7]
    WORD $0x386d68ee // ldrb    w14, [x7, x13]
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0x9b173b4e // madd    x14, x26, x23, x14
    WORD $0xd100c1de // sub    x30, x14, #48
    BNE LBB5_2173
    B LBB5_2181
LBB5_2175:
    CMP $0, R30
    BEQ LBB5_2192
    WORD $0x9adb27cd // lsr    x13, x30, x27
    CMP $0, R13
    BEQ LBB5_2178
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x4b04006e // sub    w14, w3, w4
    WORD $0x9adb21ad // lsl    x13, x13, x27
    WORD $0x2a1f03fd // mov    w29, wzr
    WORD $0xaa2d03f8 // mvn    x24, x13
    WORD $0x110005c3 // add    w3, w14, #1
    B LBB5_2183
LBB5_2178:
    WORD $0x8b1e0bcd // add    x13, x30, x30, lsl #2
    WORD $0x11000484 // add    w4, w4, #1
    WORD $0xd37ff9be // lsl    x30, x13, #1
    WORD $0x9adb27cd // lsr    x13, x30, x27
    CMP $0, R13
    BEQ LBB5_2178
    B LBB5_2171
LBB5_2179:
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_2109
LBB5_2180:
    WORD $0x2a1f03fd // mov    w29, wzr
LBB5_2181:
    WORD $0x4b04006d // sub    w13, w3, w4
    WORD $0x110005a3 // add    w3, w13, #1
    CMP $0, R30
    BNE LBB5_2183
    B LBB5_2185
LBB5_2182:
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x8a1803cd // and    x13, x30, x24
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9de // lsl    x30, x14, #1
    CMP $0, R13
    BEQ LBB5_2185
LBB5_2183:
    WORD $0x9adb27cd // lsr    x13, x30, x27
    WORD $0x93407fa4 // sxtw    x4, w29
    WORD $0xeb04005f // cmp    x2, x4
    BLS LBB5_2182
    WORD $0x1100c1ad // add    w13, w13, #48
    WORD $0x110007bd // add    w29, w29, #1
    WORD $0x38246a4d // strb    w13, [x18, x4]
    WORD $0x8a1803cd // and    x13, x30, x24
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9de // lsl    x30, x14, #1
    CMP $0, R13
    BNE LBB5_2183
LBB5_2185:
    WORD $0x710007bf // cmp    w29, #1
    BLT LBB5_2190
    WORD $0x2a1d03ed // mov    w13, w29
    WORD $0x8b1201ae // add    x14, x13, x18
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_2191
LBB5_2187:
    WORD $0xf10005bb // subs    x27, x13, #1
    BLS LBB5_2193
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d4a4e // ldrb    w14, [x18, w13, uxtw]
    WORD $0xaa1b03ed // mov    x13, x27
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2187
    WORD $0x2a1b03fa // mov    w26, w27
    WORD $0x2a1b03fd // mov    w29, w27
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_2109
LBB5_2190:
    CMP $0, R29
    BEQ LBB5_2194
LBB5_2191:
    WORD $0x2a1d03fa // mov    w26, w29
    WORD $0x2a1d03fb // mov    w27, w29
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_2109
LBB5_2192:
    WORD $0x2a1f03fa // mov    w26, wzr
    WORD $0x2a1f03fd // mov    w29, wzr
    WORD $0x2a1f03fb // mov    w27, wzr
    WORD $0x4b1900a5 // sub    w5, w5, w25
    B LBB5_2109
LBB5_2193:
    WORD $0x510005ba // sub    w26, w13, #1
    B LBB5_2195
LBB5_2194:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_2195:
    WORD $0x2a1f03e3 // mov    w3, wzr
    B LBB5_2107
LBB5_2196:
    WORD $0x310ff8bf // cmn    w5, #1022
    BGT LBB5_2223
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0xf94007fd // ldr    x29, [sp, #8]
    CMP $0, R27
    BEQ LBB5_2238
    WORD $0x110ff4a6 // add    w6, w5, #1021
    WORD $0x3110e8bf // cmn    w5, #1082
    BHI LBB5_2225
    WORD $0xb201e7e5 // mov    x5, #-7378697629483820647
    WORD $0x52800147 // mov    w7, #10
    WORD $0xf2933345 // movk    x5, #39322
    WORD $0xf2e03325 // movk    x5, #409, lsl #48
    B LBB5_2203
LBB5_2200:
    WORD $0x2a1f03e3 // mov    w3, wzr
LBB5_2201:
    WORD $0x2a1f03fa // mov    w26, wzr
LBB5_2202:
    WORD $0x1100f0cd // add    w13, w6, #60
    WORD $0x3101e0df // cmn    w6, #120
    WORD $0x2a1a03fb // mov    w27, w26
    WORD $0x2a0d03e6 // mov    w6, w13
    WORD $0x2a1a03e4 // mov    w4, w26
    BGE LBB5_2226
LBB5_2203:
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x0abb7f73 // bic    w19, w27, w27, asr #31
LBB5_2204:
    WORD $0xeb0d027f // cmp    x19, x13
    BEQ LBB5_2207
    WORD $0x386d6a4e // ldrb    w14, [x18, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b073a8e // madd    x14, x20, x7, x14
    WORD $0xd100c1d4 // sub    x20, x14, #48
    WORD $0xd37cfe8e // lsr    x14, x20, #60
    CMP $0, R14
    BEQ LBB5_2204
    WORD $0xaa1403e4 // mov    x4, x20
    WORD $0x2a0d03f3 // mov    w19, w13
    B LBB5_2209
LBB5_2207:
    CMP $0, R20
    BEQ LBB5_2201
LBB5_2208:
    WORD $0x8b140a8d // add    x13, x20, x20, lsl #2
    WORD $0x11000673 // add    w19, w19, #1
    WORD $0xd37ff9a4 // lsl    x4, x13, #1
    WORD $0xeb05029f // cmp    x20, x5
    WORD $0xaa0403f4 // mov    x20, x4
    BLO LBB5_2208
LBB5_2209:
    WORD $0x6b1b027f // cmp    w19, w27
    BGE LBB5_2213
    WORD $0x2a1303ee // mov    w14, w19
    WORD $0xaa1f03ed // mov    x13, xzr
    WORD $0x93407dd4 // sxtw    x20, w14
    WORD $0x93407f55 // sxtw    x21, w26
    WORD $0x8b140256 // add    x22, x18, x20
LBB5_2211:
    WORD $0xd37cfc8e // lsr    x14, x4, #60
    WORD $0x9240ec84 // and    x4, x4, #0xfffffffffffffff
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x382d6a4e // strb    w14, [x18, x13]
    WORD $0x386d6ace // ldrb    w14, [x22, x13]
    WORD $0x910005ad // add    x13, x13, #1
    WORD $0x9b07388e // madd    x14, x4, x7, x14
    WORD $0xd100c1c4 // sub    x4, x14, #48
    WORD $0x8b0d028e // add    x14, x20, x13
    WORD $0xeb1501df // cmp    x14, x21
    BLT LBB5_2211
    WORD $0x2a0d03fa // mov    w26, w13
    CMP $0, R4
    BNE LBB5_2215
    B LBB5_2217
LBB5_2213:
    WORD $0x2a1f03fa // mov    w26, wzr
    B LBB5_2215
LBB5_2214:
    WORD $0xd37cfc8d // lsr    x13, x4, #60
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x9240ec8d // and    x13, x4, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9c4 // lsl    x4, x14, #1
    CMP $0, R13
    BEQ LBB5_2217
LBB5_2215:
    WORD $0x93407f4d // sxtw    x13, w26
    WORD $0xeb0d005f // cmp    x2, x13
    BLS LBB5_2214
    WORD $0xd37cfc8e // lsr    x14, x4, #60
    WORD $0x1100075a // add    w26, w26, #1
    WORD $0x321c05ce // orr    w14, w14, #0x30
    WORD $0x382d6a4e // strb    w14, [x18, x13]
    WORD $0x9240ec8d // and    x13, x4, #0xfffffffffffffff
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9c4 // lsl    x4, x14, #1
    CMP $0, R13
    BNE LBB5_2215
LBB5_2217:
    WORD $0x4b13006d // sub    w13, w3, w19
    WORD $0x7100075f // cmp    w26, #1
    WORD $0x110005a3 // add    w3, w13, #1
    BLT LBB5_2222
    WORD $0x2a1a03ed // mov    w13, w26
    WORD $0x8b1201ae // add    x14, x13, x18
    WORD $0x385ff1ce // ldurb    w14, [x14, #-1]
    WORD $0x7100c1df // cmp    w14, #48
    BNE LBB5_2202
LBB5_2219:
    WORD $0xf10005ba // subs    x26, x13, #1
    BLS LBB5_2200
    WORD $0x510009ad // sub    w13, w13, #2
    WORD $0x386d4a4e // ldrb    w14, [x18, w13, uxtw]
    WORD $0xaa1a03ed // mov    x13, x26
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2219
    B LBB5_2202
LBB5_2222:
    CMP $0, R26
    BNE LBB5_2202
    B LBB5_2200
LBB5_2223:
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x510004a6 // sub    w6, w5, #1
    WORD $0xd2effe15 // mov    x21, #9218868437227405312
    WORD $0xf94007fd // ldr    x29, [sp, #8]
    WORD $0x711000bf // cmp    w5, #1024
    BLE LBB5_2255
    B LBB5_2377
LBB5_2224:
    WORD $0xaa2203ee // mvn    x14, x2
    WORD $0xcb2d41c2 // sub    x2, x14, w13, uxtw
    B LBB5_1918
LBB5_2225:
    WORD $0x2a1b03e4 // mov    w4, w27
    WORD $0x2a0603ed // mov    w13, w6
LBB5_2226:
    WORD $0xaa1f03e5 // mov    x5, xzr
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x4b0d03f3 // neg    w19, w13
    WORD $0x0aa47c95 // bic    w21, w4, w4, asr #31
    WORD $0x5280014d // mov    w13, #10
LBB5_2227:
    WORD $0xeb0502bf // cmp    x21, x5
    BEQ LBB5_2234
    WORD $0x38656a4e // ldrb    w14, [x18, x5]
    WORD $0x910004a5 // add    x5, x5, #1
    WORD $0x9b0d3a8e // madd    x14, x20, x13, x14
    WORD $0xd100c1d4 // sub    x20, x14, #48
    WORD $0x9ad3268e // lsr    x14, x20, x19
    CMP $0, R14
    BEQ LBB5_2227
    WORD $0x2a0503f5 // mov    w21, w5
LBB5_2230:
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x6b0402bf // cmp    w21, w4
    WORD $0x9ad321ad // lsl    x13, x13, x19
    WORD $0xaa2d03e6 // mvn    x6, x13
    BGE LBB5_2240
    WORD $0x2a1503ed // mov    w13, w21
    WORD $0xaa1f03e7 // mov    x7, xzr
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x93407f44 // sxtw    x4, w26
    WORD $0x8b0d0245 // add    x5, x18, x13
    WORD $0x52800156 // mov    w22, #10
LBB5_2232:
    WORD $0x9ad3268e // lsr    x14, x20, x19
    WORD $0x8a060294 // and    x20, x20, x6
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x38276a4e // strb    w14, [x18, x7]
    WORD $0x386768ae // ldrb    w14, [x5, x7]
    WORD $0x910004e7 // add    x7, x7, #1
    WORD $0x9b163a8e // madd    x14, x20, x22, x14
    WORD $0xd100c1d4 // sub    x20, x14, #48
    WORD $0x8b0701ae // add    x14, x13, x7
    WORD $0xeb0401df // cmp    x14, x4
    BLT LBB5_2232
    WORD $0x2a0703e5 // mov    w5, w7
    B LBB5_2241
LBB5_2234:
    WORD $0x12807fa6 // mov    w6, #-1022
    CMP $0, R20
    BEQ LBB5_2238
    WORD $0x9ad3268d // lsr    x13, x20, x19
    CMP $0, R13
    BEQ LBB5_2237
    WORD $0x9280000d // mov    x13, #-1
    WORD $0x4b15006e // sub    w14, w3, w21
    WORD $0x9ad321ad // lsl    x13, x13, x19
    WORD $0x2a1f03e5 // mov    w5, wzr
    WORD $0xaa2d03e6 // mvn    x6, x13
    WORD $0x110005c3 // add    w3, w14, #1
    B LBB5_2243
LBB5_2237:
    WORD $0x8b140a8d // add    x13, x20, x20, lsl #2
    WORD $0x110006b5 // add    w21, w21, #1
    WORD $0xd37ff9b4 // lsl    x20, x13, #1
    WORD $0x9ad3268d // lsr    x13, x20, x19
    CMP $0, R13
    BEQ LBB5_2237
    B LBB5_2230
LBB5_2238:
    WORD $0x2a1f03e7 // mov    w7, wzr
    B LBB5_2351
LBB5_2239:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x9280002f // mov    x15, #-2
    B LBB5_2442
LBB5_2240:
    WORD $0x2a1f03e5 // mov    w5, wzr
LBB5_2241:
    WORD $0x4b15006d // sub    w13, w3, w21
    WORD $0x110005a3 // add    w3, w13, #1
    CMP $0, R20
    BNE LBB5_2243
    B LBB5_2245
LBB5_2242:
    WORD $0xf10001bf // cmp    x13, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x8a06028d // and    x13, x20, x6
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d4 // lsl    x20, x14, #1
    CMP $0, R13
    BEQ LBB5_2245
LBB5_2243:
    WORD $0x9ad3268d // lsr    x13, x20, x19
    WORD $0x93407ca4 // sxtw    x4, w5
    WORD $0xeb04005f // cmp    x2, x4
    BLS LBB5_2242
    WORD $0x1100c1ad // add    w13, w13, #48
    WORD $0x110004a5 // add    w5, w5, #1
    WORD $0x38246a4d // strb    w13, [x18, x4]
    WORD $0x8a06028d // and    x13, x20, x6
    WORD $0x8b0d09ae // add    x14, x13, x13, lsl #2
    WORD $0xd37ff9d4 // lsl    x20, x14, #1
    CMP $0, R13
    BNE LBB5_2243
LBB5_2245:
    WORD $0x710004bf // cmp    w5, #1
    BLT LBB5_2249
    WORD $0x2a0503e4 // mov    w4, w5
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0x8b12008d // add    x13, x4, x18
    WORD $0x385ff1ad // ldurb    w13, [x13, #-1]
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_2250
LBB5_2247:
    WORD $0xaa0403ed // mov    x13, x4
    WORD $0xf1000484 // subs    x4, x4, #1
    BLS LBB5_2253
    WORD $0x510009ae // sub    w14, w13, #2
    WORD $0x386e4a4e // ldrb    w14, [x18, w14, uxtw]
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2247
    B LBB5_2254
LBB5_2249:
    WORD $0xaa1f03e4 // mov    x4, xzr
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0x2a0503ed // mov    w13, w5
    CMP $0, R5
    BNE LBB5_2256
    B LBB5_2376
LBB5_2250:
    WORD $0x2a0503ed // mov    w13, w5
    B LBB5_2256
LBB5_2251:
    WORD $0x7101f5df // cmp    w14, #125
    BNE LBB5_2452
    WORD $0xaa0803f2 // mov    x18, x8
    WORD $0xd3487dad // ubfx    x13, x13, #8, #24
    WORD $0xf84b8e4b // ldr    x11, [x18, #184]!
    WORD $0xa97ec240 // ldp    x0, x16, [x18, #-24]
    WORD $0xb940164e // ldr    w14, [x18, #20]
    WORD $0xb9402a4f // ldr    w15, [x18, #40]
    WORD $0x110005ce // add    w14, w14, #1
    WORD $0x8b10116b // add    x11, x11, x16, lsl #4
    WORD $0x0b0d01ef // add    w15, w15, w13
    WORD $0xb900164e // str    w14, [x18, #20]
    WORD $0xcb0b000e // sub    x14, x0, x11
    WORD $0xb9002a4f // str    w15, [x18, #40]
    WORD $0xd344fdce // lsr    x14, x14, #4
    WORD $0xa9403d70 // ldp    x16, x15, [x11]
    WORD $0xf81f024f // stur    x15, [x18, #-16]
    WORD $0x92609e0f // and    x15, x16, #0xffffffff000000ff
    WORD $0x2901396d // stp    w13, w14, [x11, #8]
    WORD $0xf85f824d // ldur    x13, [x18, #-8]
    WORD $0xb9402e4e // ldr    w14, [x18, #44]
    WORD $0xf900016f // str    x15, [x11]
    WORD $0xeb0e01bf // cmp    x13, x14
    BHI LBB5_1701
    B LBB5_1923
LBB5_2253:
    WORD $0x2a1f03e3 // mov    w3, wzr
LBB5_2254:
    WORD $0x510005ba // sub    w26, w13, #1
    WORD $0x12807fa6 // mov    w6, #-1022
    WORD $0x2a1a03fb // mov    w27, w26
LBB5_2255:
    WORD $0x2a1f03e7 // mov    w7, wzr
    WORD $0x2a1a03e5 // mov    w5, w26
    WORD $0x2a1b03ed // mov    w13, w27
    CMP $0, R27
    BEQ LBB5_2351
LBB5_2256:
    WORD $0x39400253 // ldrb    w19, [x18]
    WORD $0x93407da7 // sxtw    x7, w13
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_2333
    WORD $0x710004ff // cmp    w7, #1
    BNE LBB5_2259
LBB5_2258:
    WORD $0x5282b18d // mov    w13, #5516
    ADR LSHIFT_TAB, R14
    WORD $0x910001ce // add    x14, x14, :lo12:LSHIFT_TAB
    WORD $0x52800215 // mov    w21, #16
    WORD $0x8b0701ce // add    x14, x14, x7
    WORD $0x386d69cd // ldrb    w13, [x14, x13]
    CMP $0, R13
    BNE LBB5_2334
    B LBB5_2336
LBB5_2259:
    WORD $0x39400653 // ldrb    w19, [x18, #1]
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_2333
    WORD $0x710008ff // cmp    w7, #2
    BEQ LBB5_2258
    WORD $0x39400a53 // ldrb    w19, [x18, #2]
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_2333
    WORD $0x71000cff // cmp    w7, #3
    BEQ LBB5_2258
    WORD $0x39400e53 // ldrb    w19, [x18, #3]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_2333
    WORD $0x710010ff // cmp    w7, #4
    BEQ LBB5_2258
    WORD $0x39401253 // ldrb    w19, [x18, #4]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710014ff // cmp    w7, #5
    BEQ LBB5_2258
    WORD $0x39401653 // ldrb    w19, [x18, #5]
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710018ff // cmp    w7, #6
    BEQ LBB5_2258
    WORD $0x39401a53 // ldrb    w19, [x18, #6]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_2333
    WORD $0x71001cff // cmp    w7, #7
    BEQ LBB5_2258
    WORD $0x39401e53 // ldrb    w19, [x18, #7]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_2333
    WORD $0x710020ff // cmp    w7, #8
    BEQ LBB5_2258
    WORD $0x39402253 // ldrb    w19, [x18, #8]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710024ff // cmp    w7, #9
    BEQ LBB5_2258
    WORD $0x39402653 // ldrb    w19, [x18, #9]
    WORD $0x52800694 // mov    w20, #52
    WORD $0x7100d27f // cmp    w19, #52
    BNE LBB5_2333
    WORD $0x710028ff // cmp    w7, #10
    BEQ LBB5_2258
    WORD $0x39402a53 // ldrb    w19, [x18, #10]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_2333
    WORD $0x71002cff // cmp    w7, #11
    BEQ LBB5_2258
    WORD $0x39402e53 // ldrb    w19, [x18, #11]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710030ff // cmp    w7, #12
    BEQ LBB5_2258
    WORD $0x39403253 // ldrb    w19, [x18, #12]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_2333
    WORD $0x710034ff // cmp    w7, #13
    BEQ LBB5_2258
    WORD $0x39403653 // ldrb    w19, [x18, #13]
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_2333
    WORD $0x710038ff // cmp    w7, #14
    BEQ LBB5_2258
    WORD $0x39403a53 // ldrb    w19, [x18, #14]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_2333
    WORD $0x71003cff // cmp    w7, #15
    BEQ LBB5_2258
    WORD $0x39403e53 // ldrb    w19, [x18, #15]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_2333
    WORD $0x710040ff // cmp    w7, #16
    BEQ LBB5_2258
    WORD $0x39404253 // ldrb    w19, [x18, #16]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_2333
    WORD $0x710044ff // cmp    w7, #17
    BEQ LBB5_2258
    WORD $0x39404653 // ldrb    w19, [x18, #17]
    WORD $0x52800694 // mov    w20, #52
    WORD $0x7100d27f // cmp    w19, #52
    BNE LBB5_2333
    WORD $0x710048ff // cmp    w7, #18
    BEQ LBB5_2258
    WORD $0x39404a53 // ldrb    w19, [x18, #18]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_2333
    WORD $0x71004cff // cmp    w7, #19
    BEQ LBB5_2258
    WORD $0x39404e53 // ldrb    w19, [x18, #19]
    WORD $0x52800694 // mov    w20, #52
    WORD $0x7100d27f // cmp    w19, #52
    BNE LBB5_2333
    WORD $0x710050ff // cmp    w7, #20
    BEQ LBB5_2258
    WORD $0x39405253 // ldrb    w19, [x18, #20]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710054ff // cmp    w7, #21
    BEQ LBB5_2258
    WORD $0x39405653 // ldrb    w19, [x18, #21]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_2333
    WORD $0x710058ff // cmp    w7, #22
    BEQ LBB5_2258
    WORD $0x39405a53 // ldrb    w19, [x18, #22]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_2333
    WORD $0x71005cff // cmp    w7, #23
    BEQ LBB5_2258
    WORD $0x39405e53 // ldrb    w19, [x18, #23]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_2333
    WORD $0x710060ff // cmp    w7, #24
    BEQ LBB5_2258
    WORD $0x39406253 // ldrb    w19, [x18, #24]
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_2333
    WORD $0x710064ff // cmp    w7, #25
    BEQ LBB5_2258
    WORD $0x39406653 // ldrb    w19, [x18, #25]
    WORD $0x528006d4 // mov    w20, #54
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_2333
    WORD $0x710068ff // cmp    w7, #26
    BEQ LBB5_2258
    WORD $0x39406a53 // ldrb    w19, [x18, #26]
    WORD $0x7100da7f // cmp    w19, #54
    BNE LBB5_2333
    WORD $0x71006cff // cmp    w7, #27
    BEQ LBB5_2258
    WORD $0x39406e53 // ldrb    w19, [x18, #27]
    WORD $0x52800714 // mov    w20, #56
    WORD $0x7100e27f // cmp    w19, #56
    BNE LBB5_2333
    WORD $0x710070ff // cmp    w7, #28
    BEQ LBB5_2258
    WORD $0x39407253 // ldrb    w19, [x18, #28]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_2333
    WORD $0x710074ff // cmp    w7, #29
    BEQ LBB5_2258
    WORD $0x39407653 // ldrb    w19, [x18, #29]
    WORD $0x52800734 // mov    w20, #57
    WORD $0x7100e67f // cmp    w19, #57
    BNE LBB5_2333
    WORD $0x710078ff // cmp    w7, #30
    BEQ LBB5_2258
    WORD $0x39407a53 // ldrb    w19, [x18, #30]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_2333
    WORD $0x71007cff // cmp    w7, #31
    BEQ LBB5_2258
    WORD $0x39407e53 // ldrb    w19, [x18, #31]
    WORD $0x52800714 // mov    w20, #56
    WORD $0x7100e27f // cmp    w19, #56
    BNE LBB5_2333
    WORD $0x710080ff // cmp    w7, #32
    BEQ LBB5_2258
    WORD $0x39408253 // ldrb    w19, [x18, #32]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710084ff // cmp    w7, #33
    BEQ LBB5_2258
    WORD $0x39408653 // ldrb    w19, [x18, #33]
    WORD $0x52800614 // mov    w20, #48
    WORD $0x7100c27f // cmp    w19, #48
    BNE LBB5_2333
    WORD $0x710088ff // cmp    w7, #34
    BEQ LBB5_2258
    WORD $0x39408a53 // ldrb    w19, [x18, #34]
    WORD $0x52800674 // mov    w20, #51
    WORD $0x7100ce7f // cmp    w19, #51
    BNE LBB5_2333
    WORD $0x71008cff // cmp    w7, #35
    BEQ LBB5_2258
    WORD $0x39408e53 // ldrb    w19, [x18, #35]
    WORD $0x52800634 // mov    w20, #49
    WORD $0x7100c67f // cmp    w19, #49
    BNE LBB5_2333
    WORD $0x710090ff // cmp    w7, #36
    BEQ LBB5_2258
    WORD $0x39409253 // ldrb    w19, [x18, #36]
    WORD $0x52800654 // mov    w20, #50
    WORD $0x7100ca7f // cmp    w19, #50
    BNE LBB5_2333
    WORD $0x710094ff // cmp    w7, #37
    BEQ LBB5_2258
    WORD $0x39409653 // ldrb    w19, [x18, #37]
    WORD $0x528006b4 // mov    w20, #53
    WORD $0x7100d67f // cmp    w19, #53
    BNE LBB5_2333
    WORD $0x52800215 // mov    w21, #16
    WORD $0x710098ff // cmp    w7, #38
    BEQ LBB5_2258
    B LBB5_2335
LBB5_2333:
    WORD $0x52800215 // mov    w21, #16
    WORD $0x6b14027f // cmp    w19, w20
    BHS LBB5_2335
LBB5_2334:
    WORD $0x528001f5 // mov    w21, #15
LBB5_2335:
    WORD $0x710004ff // cmp    w7, #1
    BLT LBB5_2345
LBB5_2336:
    WORD $0x0b0702b4 // add    w20, w21, w7
    WORD $0x92407ced // and    x13, x7, #0xffffffff
    WORD $0x93407e8e // sxtw    x14, w20
    WORD $0xb202e7f7 // mov    x23, #-3689348814741910324
    WORD $0xaa1f03e4 // mov    x4, xzr
    WORD $0x910005a7 // add    x7, x13, #1
    WORD $0xd10005d3 // sub    x19, x14, #1
    WORD $0xd2ff4016 // mov    x22, #-432345564227567616
    WORD $0xf29999b7 // movk    x23, #52429
    WORD $0x92800138 // mov    x24, #-10
    B LBB5_2338
LBB5_2337:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0x51000694 // sub    w20, w20, #1
    WORD $0xd1000673 // sub    x19, x19, #1
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf10004ff // cmp    x7, #1
    BLS LBB5_2340
LBB5_2338:
    WORD $0x510008ed // sub    w13, w7, #2
    WORD $0xeb02027f // cmp    x19, x2
    WORD $0x386d4a4d // ldrb    w13, [x18, w13, uxtw]
    WORD $0x8b0dd48d // add    x13, x4, x13, lsl #53
    WORD $0x8b1601ad // add    x13, x13, x22
    WORD $0x9bd77dae // umulh    x14, x13, x23
    WORD $0xd343fdc4 // lsr    x4, x14, #3
    WORD $0x9b18348e // madd    x14, x4, x24, x13
    BHS LBB5_2337
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x38336a4e // strb    w14, [x18, x19]
    WORD $0x51000694 // sub    w20, w20, #1
    WORD $0xd1000673 // sub    x19, x19, #1
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf10004ff // cmp    x7, #1
    BHI LBB5_2338
LBB5_2340:
    WORD $0xf10029bf // cmp    x13, #10
    BLO LBB5_2345
    WORD $0x93407e8d // sxtw    x13, w20
    WORD $0x92800133 // mov    x19, #-10
    WORD $0xd10005a7 // sub    x7, x13, #1
    WORD $0xb202e7ed // mov    x13, #-3689348814741910324
    WORD $0xf29999ad // movk    x13, #52429
    B LBB5_2343
LBB5_2342:
    WORD $0xf10001df // cmp    x14, #0
    WORD $0x1a9f0421 // csinc    w1, w1, wzr, eq
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf100249f // cmp    x4, #9
    WORD $0xaa1403e4 // mov    x4, x20
    BLS LBB5_2345
LBB5_2343:
    WORD $0x9bcd7c8e // umulh    x14, x4, x13
    WORD $0xeb0200ff // cmp    x7, x2
    WORD $0xd343fdd4 // lsr    x20, x14, #3
    WORD $0x9b13128e // madd    x14, x20, x19, x4
    BHS LBB5_2342
    WORD $0x1100c1ce // add    w14, w14, #48
    WORD $0x38276a4e // strb    w14, [x18, x7]
    WORD $0xd10004e7 // sub    x7, x7, #1
    WORD $0xf100249f // cmp    x4, #9
    WORD $0xaa1403e4 // mov    x4, x20
    BHI LBB5_2343
LBB5_2345:
    WORD $0x0b0502ad // add    w13, w21, w5
    WORD $0x0b0302a3 // add    w3, w21, w3
    WORD $0xeb2dc05f // cmp    x2, w13, sxtw
    WORD $0x1a8281a7 // csel    w7, w13, w2, hi
    WORD $0x710004ff // cmp    w7, #1
    BLT LBB5_2350
    WORD $0x8b1200ed // add    x13, x7, x18
    WORD $0x385ff1ad // ldurb    w13, [x13, #-1]
    WORD $0x7100c1bf // cmp    w13, #48
    BNE LBB5_2351
LBB5_2347:
    WORD $0xf10004ed // subs    x13, x7, #1
    BLS LBB5_2365
    WORD $0x510008ee // sub    w14, w7, #2
    WORD $0xaa0d03e7 // mov    x7, x13
    WORD $0x386e4a4e // ldrb    w14, [x18, w14, uxtw]
    WORD $0x7100c1df // cmp    w14, #48
    BEQ LBB5_2347
    WORD $0x2a0d03e7 // mov    w7, w13
    B LBB5_2351
LBB5_2350:
    WORD $0xaa1f03e2 // mov    x2, xzr
    WORD $0x2a1f03e4 // mov    w4, wzr
    CMP $0, R7
    BEQ LBB5_2374
LBB5_2351:
    WORD $0x92800004 // mov    x4, #-1
    WORD $0x7100507f // cmp    w3, #20
    BGT LBB5_2376
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0xaa1f03e2 // mov    x2, xzr
    WORD $0x7100047f // cmp    w3, #1
    BLT LBB5_2357
    WORD $0x2a0303ed // mov    w13, w3
    WORD $0x0aa77ce5 // bic    w5, w7, w7, asr #31
    WORD $0xd10005ae // sub    x14, x13, #1
    WORD $0xaa1f03e2 // mov    x2, xzr
    WORD $0xeb0501df // cmp    x14, x5
    WORD $0x52800154 // mov    w20, #10
    WORD $0x9a8531c4 // csel    x4, x14, x5, lo
    WORD $0xaa1203f5 // mov    x21, x18
    WORD $0x91000493 // add    x19, x4, #1
LBB5_2354:
    CMP $0, R5
    BEQ LBB5_2357
    WORD $0x384016ae // ldrb    w14, [x21], #1
    WORD $0xd10005ad // sub    x13, x13, #1
    WORD $0xd10004a5 // sub    x5, x5, #1
    WORD $0x9b14384e // madd    x14, x2, x20, x14
    WORD $0xd100c1c2 // sub    x2, x14, #48
    CMP $0, R13
    BNE LBB5_2354
    WORD $0xaa1303e4 // mov    x4, x19
LBB5_2357:
    WORD $0x6b040065 // subs    w5, w3, w4
    BLE LBB5_2364
    WORD $0x710010bf // cmp    w5, #4
    BLO LBB5_2362
    WORD $0x5280002d // mov    w13, #1
    WORD $0x5280014e // mov    w14, #10
    WORD $0x121e74b3 // and    w19, w5, #0xfffffffc
    WORD $0x25d8e040 // ptrue    p0.d, vl2
    WORD $0x0b130084 // add    w4, w4, w19
    WORD $0x4e080da3 // dup    v3.2d, x13
    WORD $0x2a1303ed // mov    w13, w19
    WORD $0x4e080dc5 // dup    v5.2d, x14
    WORD $0x4ea31c64 // mov    v4.16b, v3.16b
    WORD $0x4e081c44 // mov    v4.d[0], x2
LBB5_2360:
    WORD $0x710011ad // subs    w13, w13, #4
    WORD $0x04d000a3 // mul    z3.d, p0/m, z3.d, z5.d
    WORD $0x04d000a4 // mul    z4.d, p0/m, z4.d, z5.d
    BNE LBB5_2360
    WORD $0x4ec33885 // zip1    v5.2d, v4.2d, v3.2d
    WORD $0x6b1300bf // cmp    w5, w19
    WORD $0x4ec37883 // zip2    v3.2d, v4.2d, v3.2d
    WORD $0x04d000a3 // mul    z3.d, p0/m, z3.d, z5.d
    WORD $0x25d8e020 // ptrue    p0.d, vl1
    WORD $0x6e034064 // ext    v4.16b, v3.16b, v3.16b, #8
    WORD $0x04d00083 // mul    z3.d, p0/m, z3.d, z4.d
    WORD $0x9e660062 // fmov    x2, d3
    BEQ LBB5_2364
LBB5_2362:
    WORD $0x4b04006d // sub    w13, w3, w4
LBB5_2363:
    WORD $0x8b02084e // add    x14, x2, x2, lsl #2
    WORD $0x710005ad // subs    w13, w13, #1
    WORD $0xd37ff9c2 // lsl    x2, x14, #1
    BNE LBB5_2363
LBB5_2364:
    WORD $0x2a1f03e4 // mov    w4, wzr
    TST $(1<<31), R3
    BEQ LBB5_2366
    B LBB5_2374
LBB5_2365:
    WORD $0xaa1f03e2 // mov    x2, xzr
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0x510004e7 // sub    w7, w7, #1
LBB5_2366:
    WORD $0x6b0300ff // cmp    w7, w3
    BLE LBB5_2372
    WORD $0x38634a4d // ldrb    w13, [x18, w3, uxtw]
    WORD $0x7100d5bf // cmp    w13, #53
    BNE LBB5_2373
    WORD $0x1100046e // add    w14, w3, #1
    WORD $0x6b0701df // cmp    w14, w7
    BNE LBB5_2373
    WORD $0x52800024 // mov    w4, #1
    CMP $0, R1
    BNE LBB5_2374
    CMP $0, R3
    BEQ LBB5_2372
    WORD $0x5100046d // sub    w13, w3, #1
    WORD $0x386d4a4d // ldrb    w13, [x18, w13, uxtw]
    WORD $0x120001a4 // and    w4, w13, #0x1
    WORD $0x8b244044 // add    x4, x2, w4, uxtw
    WORD $0xd2e0040d // mov    x13, #9007199254740992
    WORD $0xeb0d009f // cmp    x4, x13
    BEQ LBB5_2375
    B LBB5_2376
LBB5_2372:
    WORD $0x2a1f03e4 // mov    w4, wzr
    WORD $0x8b3f4044 // add    x4, x2, wzr, uxtw
    WORD $0xd2e0040d // mov    x13, #9007199254740992
    WORD $0xeb0d009f // cmp    x4, x13
    BEQ LBB5_2375
    B LBB5_2376
LBB5_2373:
    WORD $0x7100d1bf // cmp    w13, #52
    WORD $0x1a9f97e4 // cset    w4, hi
LBB5_2374:
    WORD $0x8b244044 // add    x4, x2, w4, uxtw
    WORD $0xd2e0040d // mov    x13, #9007199254740992
    WORD $0xeb0d009f // cmp    x4, x13
    BNE LBB5_2376
LBB5_2375:
    WORD $0x110004cd // add    w13, w6, #1
    WORD $0xaa1f03f4 // mov    x20, xzr
    WORD $0x710ff8df // cmp    w6, #1022
    WORD $0xd2e00204 // mov    x4, #4503599627370496
    WORD $0xd2effe15 // mov    x21, #9218868437227405312
    WORD $0x2a0d03e6 // mov    w6, w13
    BGT LBB5_2377
LBB5_2376:
    WORD $0x110ffccd // add    w13, w6, #1023
    WORD $0x9374d08e // sbfx    x14, x4, #52, #1
    WORD $0x120029ad // and    w13, w13, #0x7ff
    WORD $0xaa0403f4 // mov    x20, x4
    WORD $0x8a0dd1d5 // and    x21, x14, x13, lsl #52
LBB5_2377:
    WORD $0x9240ce8d // and    x13, x20, #0xfffffffffffff
    WORD $0x7100b41f // cmp    w0, #45
    WORD $0xaa1501ad // orr    x13, x13, x21
    WORD $0x1e620223 // scvtf    d3, w17
    WORD $0xb24101ae // orr    x14, x13, #0x8000000000000000
    WORD $0xb9401ff9 // ldr    w25, [sp, #28]
    WORD $0x9a8d01cd // csel    x13, x14, x13, eq
    WORD $0x9e6701a4 // fmov    d4, x13
LBB5_2378:
    WORD $0x1e640863 // fmul    d3, d3, d4
LBB5_2379:
    WORD $0x9e660060 // fmov    x0, d3
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x52800082 // mov    w2, #4
    WORD $0xd2effe0e // mov    x14, #9218868437227405312
    WORD $0x9240f80d // and    x13, x0, #0x7fffffffffffffff
    WORD $0xeb0e01bf // cmp    x13, x14
    BNE LBB5_2384
LBB5_2380:
    WORD $0xf1004d9f // cmp    x12, #19
    BEQ LBB5_1804
LBB5_2381:
    WORD $0xf1002d9f // cmp    x12, #11
    BEQ LBB5_2385
    WORD $0xf1000d9f // cmp    x12, #3
    BNE LBB5_341
LBB5_2383:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x5280006c // mov    w12, #3
    B LBB5_2386
LBB5_2384:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x5280026c // mov    w12, #19
    WORD $0xaa108190 // orr    x16, x12, x16, lsl #32
    WORD $0xaa0f03f1 // mov    x17, x15
    WORD $0xaa0003fd // mov    x29, x0
    WORD $0xf90005a0 // str    x0, [x13, #8]
    WORD $0xb940d90e // ldr    w14, [x8, #216]
    WORD $0xf90001b0 // str    x16, [x13]
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0x110005ce // add    w14, w14, #1
    B LBB5_2387
LBB5_2385:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x5280016c // mov    w12, #11
LBB5_2386:
    WORD $0xf90005bd // str    x29, [x13, #8]
    WORD $0xb940d90e // ldr    w14, [x8, #216]
    WORD $0xaa108190 // orr    x16, x12, x16, lsl #32
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0xaa0f03f1 // mov    x17, x15
    WORD $0x2a0203e1 // mov    w1, w2
    WORD $0x110005ce // add    w14, w14, #1
    WORD $0xf90001b0 // str    x16, [x13]
LBB5_2387:
    WORD $0x9100424d // add    x13, x18, #16
    WORD $0xb900d90e // str    w14, [x8, #216]
    WORD $0xf900510d // str    x13, [x8, #160]
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a8101a2 // csel    w2, w13, w1, eq
    CMP $0, R1
    BNE LBB5_2444
LBB5_2388:
    WORD $0xf940610d // ldr    x13, [x8, #192]
    WORD $0x9100824e // add    x14, x18, #32
    WORD $0xeb0d01df // cmp    x14, x13
    BHI LBB5_2444
    WORD $0xaa1103ef // mov    x15, x17
    WORD $0x384015e0 // ldrb    w0, [x15], #1
    WORD $0x7100801f // cmp    w0, #32
    BHI LBB5_2412
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ac021ad // lsl    x13, x13, x0
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_2412
    WORD $0x39400620 // ldrb    w0, [x17, #1]
    WORD $0x91000a2f // add    x15, x17, #2
    WORD $0x7100801f // cmp    w0, #32
    BHI LBB5_2412
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ac021ad // lsl    x13, x13, x0
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_2412
    WORD $0xf9404910 // ldr    x16, [x8, #144]
    WORD $0xcb1001ed // sub    x13, x15, x16
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_2396
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x9acd21ed // lsl    x13, x15, x13
    WORD $0xea0d01cf // ands    x15, x14, x13
    BNE LBB5_2399
    WORD $0x9101020f // add    x15, x16, #64
LBB5_2396:
    WORD $0x4f04e5e3 // movi    v3.16b, #143
    WORD $0xd10101f0 // sub    x16, x15, #64
LBB5_2397:
    WORD $0xadc21604 // ldp    q4, q5, [x16, #64]!
    WORD $0x4e231c90 // and    v16.16b, v4.16b, v3.16b
    WORD $0x4e100010 // tbl    v16.16b, { v0.16b }, v16.16b
    WORD $0xad411e06 // ldp    q6, q7, [x16, #32]
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
    WORD $0x1e2600af // fmov    w15, s5
    WORD $0x1e2600d1 // fmov    w17, s6
    WORD $0x33103dae // bfi    w14, w13, #16, #16
    WORD $0xaa0f81cd // orr    x13, x14, x15, lsl #32
    WORD $0xaa11c1ad // orr    x13, x13, x17, lsl #48
    WORD $0xb10005bf // cmn    x13, #1
    BEQ LBB5_2397
LBB5_2398:
    WORD $0xaa2d03ef // mvn    x15, x13
    WORD $0xa9093d10 // stp    x16, x15, [x8, #144]
LBB5_2399:
    WORD $0xdac001ed // rbit    x13, x15
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d020f // add    x15, x16, x13
    WORD $0x384015e0 // ldrb    w0, [x15], #1
LBB5_2400:
    WORD $0xf940016d // ldr    x13, [x11]
    WORD $0x7100b01f // cmp    w0, #44
    WORD $0x910401ad // add    x13, x13, #256
    WORD $0xf900016d // str    x13, [x11]
    BNE LBB5_2413
LBB5_2401:
    WORD $0x394001e2 // ldrb    w2, [x15]
    WORD $0x910005f1 // add    x17, x15, #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_1702
    WORD $0xd284c00e // mov    x14, #9728
    WORD $0x5280002d // mov    w13, #1
    WORD $0xf2c0002e // movk    x14, #1, lsl #32
    WORD $0x9ac221ad // lsl    x13, x13, x2
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1702
    WORD $0x394005e2 // ldrb    w2, [x15, #1]
    WORD $0x91000631 // add    x17, x17, #1
    WORD $0x7100805f // cmp    w2, #32
    BHI LBB5_1702
    WORD $0x5280002d // mov    w13, #1
    WORD $0x9ac221ad // lsl    x13, x13, x2
    WORD $0xea0e01bf // tst    x13, x14
    BEQ LBB5_1702
    WORD $0xf940490f // ldr    x15, [x8, #144]
    WORD $0xcb0f022d // sub    x13, x17, x15
    WORD $0xf100fdbf // cmp    x13, #63
    BHI LBB5_2408
    WORD $0xf9404d0e // ldr    x14, [x8, #152]
    WORD $0x92800010 // mov    x16, #-1
    WORD $0x9acd220d // lsl    x13, x16, x13
    WORD $0xea0d01cd // ands    x13, x14, x13
    BNE LBB5_2411
    WORD $0x910101f1 // add    x17, x15, #64
LBB5_2408:
    ADR LCPI5_0, R13
    ADR LCPI5_1, R14
    ADR LCPI5_2, R16
    WORD $0xd101022f // sub    x15, x17, #64
    WORD $0x4f04e5e1 // movi    v1.16b, #143
    WORD $0x3dc001a0 // ldr    q0, [x13, :lo12:.LCPI5_0]
    WORD $0x3dc001c2 // ldr    q2, [x14, :lo12:.LCPI5_1]
    WORD $0x3dc00203 // ldr    q3, [x16, :lo12:.LCPI5_2]
LBB5_2409:
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
    BEQ LBB5_2409
    WORD $0xaa2d03ed // mvn    x13, x13
    WORD $0xa909350f // stp    x15, x13, [x8, #144]
LBB5_2411:
    WORD $0xdac001ad // rbit    x13, x13
    WORD $0xdac011ad // clz    x13, x13
    WORD $0x8b0d01f1 // add    x17, x15, x13
    WORD $0x38401622 // ldrb    w2, [x17], #1
    B LBB5_1703
LBB5_2412:
    WORD $0xf940016d // ldr    x13, [x11]
    WORD $0x7100b01f // cmp    w0, #44
    WORD $0x910401ad // add    x13, x13, #256
    WORD $0xf900016d // str    x13, [x11]
    BEQ LBB5_2401
LBB5_2413:
    WORD $0x7101741f // cmp    w0, #93
    BNE LBB5_2451
    WORD $0xaa0803f2 // mov    x18, x8
    WORD $0xd3487dad // ubfx    x13, x13, #8, #24
    WORD $0xf84b8e4b // ldr    x11, [x18, #184]!
    WORD $0xa97ec640 // ldp    x0, x17, [x18, #-24]
    WORD $0xb9401a4e // ldr    w14, [x18, #24]
    WORD $0xb9402650 // ldr    w16, [x18, #36]
    WORD $0x110005ce // add    w14, w14, #1
    WORD $0x8b111171 // add    x17, x11, x17, lsl #4
    WORD $0x0b0d0210 // add    w16, w16, w13
    WORD $0xcb11000b // sub    x11, x0, x17
    WORD $0xd344fd6b // lsr    x11, x11, #4
    WORD $0xb9001a4e // str    w14, [x18, #24]
    WORD $0xb9002650 // str    w16, [x18, #36]
    WORD $0xa9403a30 // ldp    x16, x14, [x17]
    WORD $0xf81f024e // stur    x14, [x18, #-16]
    WORD $0x92609e0e // and    x14, x16, #0xffffffff000000ff
    WORD $0x29012e2d // stp    w13, w11, [x17, #8]
    WORD $0xf85f824b // ldur    x11, [x18, #-8]
    WORD $0xb9402e4d // ldr    w13, [x18, #44]
    WORD $0xf900022e // str    x14, [x17]
    WORD $0xeb0d017f // cmp    x11, x13
    BLS LBB5_2416
    WORD $0xf140057f // cmp    x11, #1, lsl #12
    WORD $0xb900e50b // str    w11, [x8, #228]
    BHI LBB5_2418
LBB5_2416:
    WORD $0xf9405510 // ldr    x16, [x8, #168]
    WORD $0xd100056b // sub    x11, x11, #1
    WORD $0xb100061f // cmn    x16, #1
    WORD $0xf900590b // str    x11, [x8, #176]
    BEQ LBB5_2418
    WORD $0xaa0f03f1 // mov    x17, x15
    B LBB5_1924
LBB5_2418:
    WORD $0xaa1f03eb // mov    x11, xzr
    WORD $0xaa0f03f1 // mov    x17, x15
    CMP ZR, ZR
    BNE LBB5_1926
    B LBB5_835
LBB5_2419:
    WORD $0x2a1f03f9 // mov    w25, wzr
    B LBB5_2099
LBB5_2420:
    WORD $0xaa0503e3 // mov    x3, x5
    WORD $0x92800000 // mov    x0, #-1
    WORD $0x92800001 // mov    x1, #-1
    B LBB5_1777
LBB5_2421:
    WORD $0xaa0703ef // mov    x15, x7
    WORD $0xf1004d9f // cmp    x12, #19
    BNE LBB5_2381
    B LBB5_1804
LBB5_2422:
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0x5284e202 // mov    w2, #10000
    B LBB5_1736
LBB5_2423:
    WORD $0xcb1101ef // sub    x15, x15, x17
    WORD $0x394001e7 // ldrb    w7, [x15]
    WORD $0x5100c0ed // sub    w13, w7, #48
    WORD $0x710025bf // cmp    w13, #9
    BHI LBB5_2445
    WORD $0xaa1f03e0 // mov    x0, xzr
    WORD $0xaa1f03f1 // mov    x17, xzr
    WORD $0x5280014d // mov    w13, #10
LBB5_2425:
    WORD $0x8b1100ae // add    x14, x5, x17
    WORD $0x9b0d7c0f // mul    x15, x0, x13
    WORD $0x8b2741ef // add    x15, x15, w7, uxtw
    WORD $0x394005c7 // ldrb    w7, [x14, #1]
    WORD $0xd100c1e0 // sub    x0, x15, #48
    WORD $0x5100c0e2 // sub    w2, w7, #48
    WORD $0x7100245f // cmp    w2, #9
    WORD $0xfa529a22 // ccmp    x17, #18, #2, ls
    WORD $0x91000631 // add    x17, x17, #1
    BLO LBB5_2425
    WORD $0x8b1100af // add    x15, x5, x17
    WORD $0x7100245f // cmp    w2, #9
    BHI LBB5_2446
    WORD $0xaa1f03e3 // mov    x3, xzr
LBB5_2428:
    WORD $0x8b0300ad // add    x13, x5, x3
    WORD $0x91000463 // add    x3, x3, #1
    WORD $0x8b1101ad // add    x13, x13, x17
    WORD $0x394005a7 // ldrb    w7, [x13, #1]
    WORD $0x5100c0ed // sub    w13, w7, #48
    WORD $0x710029bf // cmp    w13, #10
    BLO LBB5_2428
    WORD $0x8b1100ad // add    x13, x5, x17
    WORD $0x52800021 // mov    w1, #1
    WORD $0x8b0301af // add    x15, x13, x3
    B LBB5_1748
LBB5_2430:
    WORD $0xaa1d03fb // mov    x27, x29
    WORD $0xaa0403f1 // mov    x17, x4
    WORD $0x2538cb83 // mov    z3.b, #92
    WORD $0x2538c444 // mov    z4.b, #34
    WORD $0x2538c3e5 // mov    z5.b, #31
LBB5_2431:
    WORD $0xa400a226 // ld1b    { z6.b }, p0/z, [x17]
    WORD $0x2403a0c3 // cmpeq    p3.b, p0/z, z6.b, z3.b
    WORD $0x2404a0c4 // cmpeq    p4.b, p0/z, z6.b, z4.b
    WORD $0x2405a0c2 // cmpeq    p2.b, p0/z, z6.b, z5.b
    WORD $0x25845085 // mov    p5.b, p4.b
    WORD $0x25c24066 // orrs    p6.b, p0/z, p3.b, p2.b
    BEQ LBB5_2433
    WORD $0x259040c5 // brkb    p5.b, p0/z, p6.b
    WORD $0x250440a5 // and    p5.b, p0/z, p5.b, p4.b
LBB5_2433:
    WORD $0x2550c0a0 // ptest    p0, p5.b
    BNE LBB5_2440
    WORD $0x2550c080 // ptest    p0, p4.b
    BEQ LBB5_2437
    WORD $0x25904084 // brkb    p4.b, p0/z, p4.b
    WORD $0x25434085 // ands    p5.b, p0/z, p4.b, p3.b
    BNE LBB5_2450
    WORD $0x25024084 // and    p4.b, p0/z, p4.b, p2.b
    B LBB5_2438
LBB5_2437:
    WORD $0x25824844 // mov    p4.b, p2.b
    WORD $0x2550c060 // ptest    p0, p3.b
    BNE LBB5_2450
LBB5_2438:
    WORD $0x2550c080 // ptest    p0, p4.b
    BNE LBB5_2478
    WORD $0x91008231 // add    x17, x17, #32
    B LBB5_2431
LBB5_2440:
    WORD $0xaa1b03fd // mov    x29, x27
LBB5_2441:
    WORD $0x25904081 // brkb    p1.b, p0/z, p4.b
    WORD $0x2a1f03f9 // mov    w25, wzr
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0x8b1101ad // add    x13, x13, x17
    WORD $0x910005b1 // add    x17, x13, #1
    WORD $0xaa2403ed // mvn    x13, x4
    WORD $0x8b0d022f // add    x15, x17, x13
    WORD $0x2a1f03e1 // mov    w1, wzr
    TST $(1<<63), R15
    BEQ LBB5_2443
LBB5_2442:
    WORD $0x4b0f03e1 // neg    w1, w15
LBB5_2443:
    WORD $0xf940510d // ldr    x13, [x8, #160]
    WORD $0x7100033f // cmp    w25, #0
    WORD $0x5280018e // mov    w14, #12
    WORD $0x52800092 // mov    w18, #4
    WORD $0x9a8e024e // csel    x14, x18, x14, eq
    WORD $0xf90005af // str    x15, [x13, #8]
    WORD $0xaa1081ce // orr    x14, x14, x16, lsl #32
    WORD $0xf9405112 // ldr    x18, [x8, #160]
    WORD $0xd2c0002f // mov    x15, #4294967296
    WORD $0xb940d510 // ldr    w16, [x8, #212]
    WORD $0x8b0f01ce // add    x14, x14, x15
    WORD $0x9100424f // add    x15, x18, #16
    WORD $0x11000610 // add    w16, w16, #1
    WORD $0xf90001ae // str    x14, [x13]
    WORD $0xf900510f // str    x15, [x8, #160]
    WORD $0xb900d510 // str    w16, [x8, #212]
    WORD $0x7100003f // cmp    w1, #0
    WORD $0x5280016d // mov    w13, #11
    WORD $0x1a8101a2 // csel    w2, w13, w1, eq
    CMP $0, R1
    BEQ LBB5_2388
LBB5_2444:
    WORD $0xaa1103ef // mov    x15, x17
    B LBB5_341
LBB5_2445:
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0xaa1f03f1 // mov    x17, xzr
    WORD $0x2a1f03e3 // mov    w3, wzr
    WORD $0xaa1f03e0 // mov    x0, xzr
    B LBB5_1748
LBB5_2446:
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x2a1f03e3 // mov    w3, wzr
    B LBB5_1748
LBB5_2447:
    WORD $0x7100047f // cmp    w3, #1
    BNE LBB5_1958
    WORD $0x5280014d // mov    w13, #10
    WORD $0x9bcd7c0d // umulh    x13, x0, x13
    WORD $0xeb0d03ff // cmp    xzr, x13
    BEQ LBB5_2473
    WORD $0x7100025f // cmp    w18, #0
    WORD $0x1280000d // mov    w13, #-1
    WORD $0x5a8d15b1 // cneg    w17, w13, eq
    WORD $0x52800023 // mov    w3, #1
    B LBB5_1970
LBB5_2450:
    WORD $0xaa1b03fd // mov    x29, x27
    B LBB5_1836
LBB5_2451:
    WORD $0x52800142 // mov    w2, #10
    B LBB5_341
LBB5_2452:
    WORD $0x52800122 // mov    w2, #9
    WORD $0xaa1103ef // mov    x15, x17
    B LBB5_341
LBB5_2453:
    WORD $0x8b1802ce // add    x14, x22, x24
    WORD $0x8b1802ef // add    x15, x23, x24
    WORD $0x394001cd // ldrb    w13, [x14]
    WORD $0x710089bf // cmp    w13, #34
    BNE LBB5_2456
LBB5_2454:
    WORD $0x910005d1 // add    x17, x14, #1
    WORD $0xcb0401ef // sub    x15, x15, x4
LBB5_2455:
    WORD $0x52800039 // mov    w25, #1
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x2a1f03e1 // mov    w1, wzr
    TST $(1<<63), R15
    BEQ LBB5_2443
    B LBB5_2442
LBB5_2456:
    WORD $0x8b1802ee // add    x14, x23, x24
    WORD $0x8b1802cf // add    x15, x22, x24
    WORD $0x390001cd // strb    w13, [x14]
    WORD $0x394005ed // ldrb    w13, [x15, #1]
    WORD $0x710089bf // cmp    w13, #34
    BEQ LBB5_2466
    WORD $0x390005cd // strb    w13, [x14, #1]
    WORD $0x394009ef // ldrb    w15, [x15, #2]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_2467
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x390009cf // strb    w15, [x14, #2]
    WORD $0x39400daf // ldrb    w15, [x13, #3]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_2468
    WORD $0x39000dcf // strb    w15, [x14, #3]
    WORD $0x394011af // ldrb    w15, [x13, #4]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_2469
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x390011cf // strb    w15, [x14, #4]
    WORD $0x394015af // ldrb    w15, [x13, #5]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_2470
    WORD $0x390015cf // strb    w15, [x14, #5]
    WORD $0x394019af // ldrb    w15, [x13, #6]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_2471
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0x390019cf // strb    w15, [x14, #6]
    WORD $0x39401daf // ldrb    w15, [x13, #7]
    WORD $0x710089ff // cmp    w15, #34
    BEQ LBB5_2472
    WORD $0x39001dcf // strb    w15, [x14, #7]
    WORD $0x91002318 // add    x24, x24, #8
    WORD $0x394021ad // ldrb    w13, [x13, #8]
    WORD $0x710089bf // cmp    w13, #34
    BNE LBB5_2456
    WORD $0x8b1802ce // add    x14, x22, x24
    WORD $0x8b1802ef // add    x15, x23, x24
    B LBB5_2454
LBB5_2465:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x9280016f // mov    x15, #-12
    B LBB5_2442
LBB5_2466:
    WORD $0xcb0402ed // sub    x13, x23, x4
    WORD $0x910009f1 // add    x17, x15, #2
    WORD $0x8b1801ad // add    x13, x13, x24
    WORD $0x910005af // add    x15, x13, #1
    B LBB5_2455
LBB5_2467:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0xcb0402ee // sub    x14, x23, x4
    WORD $0x91000db1 // add    x17, x13, #3
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910009af // add    x15, x13, #2
    B LBB5_2455
LBB5_2468:
    WORD $0xcb0402ee // sub    x14, x23, x4
    WORD $0x910011b1 // add    x17, x13, #4
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x91000daf // add    x15, x13, #3
    B LBB5_2455
LBB5_2469:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0xcb0402ee // sub    x14, x23, x4
    WORD $0x910015b1 // add    x17, x13, #5
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910011af // add    x15, x13, #4
    B LBB5_2455
LBB5_2470:
    WORD $0xcb0402ee // sub    x14, x23, x4
    WORD $0x910019b1 // add    x17, x13, #6
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910015af // add    x15, x13, #5
    B LBB5_2455
LBB5_2471:
    WORD $0x8b1802cd // add    x13, x22, x24
    WORD $0xcb0402ee // sub    x14, x23, x4
    WORD $0x91001db1 // add    x17, x13, #7
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x910019af // add    x15, x13, #6
    B LBB5_2455
LBB5_2472:
    WORD $0xcb0402ee // sub    x14, x23, x4
    WORD $0x910021b1 // add    x17, x13, #8
    WORD $0x8b1801cd // add    x13, x14, x24
    WORD $0x91001daf // add    x15, x13, #7
    B LBB5_2455
LBB5_2473:
    WORD $0x385ff1ed // ldurb    w13, [x15, #-1]
    WORD $0x8b00080e // add    x14, x0, x0, lsl #2
    WORD $0xd37ff9ce // lsl    x14, x14, #1
    WORD $0x52800023 // mov    w3, #1
    WORD $0x5100c1ad // sub    w13, w13, #48
    WORD $0x93407dad // sxtw    x13, w13
    WORD $0x937ffdb1 // asr    x17, x13, #63
    WORD $0xab0d01cd // adds    x13, x14, x13
    WORD $0x9a913631 // cinc    x17, x17, hs
    WORD $0x9340022e // sbfx    x14, x17, #0, #1
    WORD $0xca1101d1 // eor    x17, x14, x17
    CMP $0, R17
    BNE LBB5_1958
    TST $(1<<63), R14
    BNE LBB5_1958
    CMP $0, R18_PLATFORM
    BEQ LBB5_2479
    WORD $0x2a1f03e1 // mov    w1, wzr
    WORD $0x9e6301a3 // ucvtf    d3, x13
    B LBB5_1915
LBB5_2477:
    WORD $0x8b1802d1 // add    x17, x22, x24
    WORD $0x52800039 // mov    w25, #1
LBB5_2478:
    WORD $0x25904041 // brkb    p1.b, p0/z, p2.b
    WORD $0x9280000f // mov    x15, #-1
    WORD $0x2520802d // cntp    x13, p0, p1.b
    WORD $0xaa1b03fd // mov    x29, x27
    WORD $0x8b0d0231 // add    x17, x17, x13
    B LBB5_2442
LBB5_2479:
    WORD $0x2a1f03e2 // mov    w2, wzr
    WORD $0xaa0d03fd // mov    x29, x13
    B LBB5_2383
LBB5_2480:
    WORD $0x52800039 // mov    w25, #1
    WORD $0x9280016f // mov    x15, #-12
    WORD $0xaa1603f1 // mov    x17, x22
    B LBB5_2442
MASK_USE_NUMBER:
    WORD $0x00000002 // .long    2
P10_TAB:
    WORD $0x00000000; WORD $0x3ff00000
    WORD $0x00000000; WORD $0x40240000
    WORD $0x00000000; WORD $0x40590000
    WORD $0x00000000; WORD $0x408f4000
    WORD $0x00000000; WORD $0x40c38800
    WORD $0x00000000; WORD $0x40f86a00
    WORD $0x00000000; WORD $0x412e8480
    WORD $0x00000000; WORD $0x416312d0
    WORD $0x00000000; WORD $0x4197d784
    WORD $0x00000000; WORD $0x41cdcd65
    WORD $0x20000000; WORD $0x4202a05f
    WORD $0xe8000000; WORD $0x42374876
    WORD $0xa2000000; WORD $0x426d1a94
    WORD $0xe5400000; WORD $0x42a2309c
    WORD $0x1e900000; WORD $0x42d6bcc4
    WORD $0x26340000; WORD $0x430c6bf5
    WORD $0x37e08000; WORD $0x4341c379
    WORD $0x85d8a000; WORD $0x43763457
    WORD $0x674ec800; WORD $0x43abc16d
    WORD $0x60913d00; WORD $0x43e158e4
    WORD $0x78b58c40; WORD $0x4415af1d
    WORD $0xd6e2ef50; WORD $0x444b1ae4
    WORD $0x064dd592; WORD $0x4480f0cf
POW10_M128_TAB:
    WORD $0xcd60e453; WORD $0x1732c869
    WORD $0x081c0288; WORD $0xfa8fd5a0
    WORD $0x205c8eb4; WORD $0x0e7fbd42
    WORD $0x05118195; WORD $0x9c99e584
    WORD $0xa873b261; WORD $0x521fac92
    WORD $0x0655e1fa; WORD $0xc3c05ee5
    WORD $0x52909ef9; WORD $0xe6a797b7
    WORD $0x47eb5a78; WORD $0xf4b0769e
    WORD $0x939a635c; WORD $0x9028bed2
    WORD $0xecf3188b; WORD $0x98ee4a22
    WORD $0x3880fc33; WORD $0x7432ee87
    WORD $0xa82fdeae; WORD $0xbf29dcab
    WORD $0x06a13b3f; WORD $0x113faa29
    WORD $0x923bd65a; WORD $0xeef453d6
    WORD $0xa424c507; WORD $0x4ac7ca59
    WORD $0x1b6565f8; WORD $0x9558b466
    WORD $0x0d2df649; WORD $0x5d79bcf0
    WORD $0xa23ebf76; WORD $0xbaaee17f
    WORD $0x107973dc; WORD $0xf4d82c2c
    WORD $0x8ace6f53; WORD $0xe95a99df
    WORD $0x8a4be869; WORD $0x79071b9b
    WORD $0xb6c10594; WORD $0x91d8a02b
    WORD $0x6cdee284; WORD $0x9748e282
    WORD $0xa47146f9; WORD $0xb64ec836
    WORD $0x08169b25; WORD $0xfd1b1b23
    WORD $0x4d8d98b7; WORD $0xe3e27a44
    WORD $0xe50e20f7; WORD $0xfe30f0f5
    WORD $0xb0787f72; WORD $0x8e6d8c6a
    WORD $0x5e51a935; WORD $0xbdbd2d33
    WORD $0x5c969f4f; WORD $0xb208ef85
    WORD $0x35e61382; WORD $0xad2c7880
    WORD $0xb3bc4723; WORD $0xde8b2b66
    WORD $0x21afcc31; WORD $0x4c3bcb50
    WORD $0x3055ac76; WORD $0x8b16fb20
    WORD $0x2a1bbf3d; WORD $0xdf4abe24
    WORD $0x3c6b1793; WORD $0xaddcb9e8
    WORD $0x34a2af0d; WORD $0xd71d6dad
    WORD $0x4b85dd78; WORD $0xd953e862
    WORD $0x40e5ad68; WORD $0x8672648c
    WORD $0x6f33aa6b; WORD $0x87d4713d
    WORD $0x511f18c2; WORD $0x680efdaf
    WORD $0xcb009506; WORD $0xa9c98d8c
    WORD $0x2566def2; WORD $0x0212bd1b
    WORD $0xfdc0ba48; WORD $0xd43bf0ef
    WORD $0xf7604b57; WORD $0x014bb630
    WORD $0xfe98746d; WORD $0x84a57695
    WORD $0x35385e2d; WORD $0x419ea3bd
    WORD $0x7e3e9188; WORD $0xa5ced43b
    WORD $0x828675b9; WORD $0x52064cac
    WORD $0x5dce35ea; WORD $0xcf42894a
    WORD $0xd1940993; WORD $0x7343efeb
    WORD $0x7aa0e1b2; WORD $0x818995ce
    WORD $0xc5f90bf8; WORD $0x1014ebe6
    WORD $0x19491a1f; WORD $0xa1ebfb42
    WORD $0x77774ef6; WORD $0xd41a26e0
    WORD $0x9f9b60a6; WORD $0xca66fa12
    WORD $0x955522b4; WORD $0x8920b098
    WORD $0x478238d0; WORD $0xfd00b897
    WORD $0x5d5535b0; WORD $0x55b46e5f
    WORD $0x8cb16382; WORD $0x9e20735e
    WORD $0x34aa831d; WORD $0xeb2189f7
    WORD $0x2fddbc62; WORD $0xc5a89036
    WORD $0x01d523e4; WORD $0xa5e9ec75
    WORD $0xbbd52b7b; WORD $0xf712b443
    WORD $0x2125366e; WORD $0x47b233c9
    WORD $0x55653b2d; WORD $0x9a6bb0aa
    WORD $0x696e840a; WORD $0x999ec0bb
    WORD $0xeabe89f8; WORD $0xc1069cd4
    WORD $0x43ca250d; WORD $0xc00670ea
    WORD $0x256e2c76; WORD $0xf148440a
    WORD $0x6a5e5728; WORD $0x38040692
    WORD $0x5764dbca; WORD $0x96cd2a86
    WORD $0x04f5ecf2; WORD $0xc6050837
    WORD $0xed3e12bc; WORD $0xbc807527
    WORD $0xc633682e; WORD $0xf7864a44
    WORD $0xe88d976b; WORD $0xeba09271
    WORD $0xfbe0211d; WORD $0x7ab3ee6a
    WORD $0x31587ea3; WORD $0x93445b87
    WORD $0xbad82964; WORD $0x5960ea05
    WORD $0xfdae9e4c; WORD $0xb8157268
    WORD $0x298e33bd; WORD $0x6fb92487
    WORD $0x3d1a45df; WORD $0xe61acf03
    WORD $0x79f8e056; WORD $0xa5d3b6d4
    WORD $0x06306bab; WORD $0x8fd0c162
    WORD $0x9877186c; WORD $0x8f48a489
    WORD $0x87bc8696; WORD $0xb3c4f1ba
    WORD $0xfe94de87; WORD $0x331acdab
    WORD $0x29aba83c; WORD $0xe0b62e29
    WORD $0x7f1d0b14; WORD $0x9ff0c08b
    WORD $0xba0b4925; WORD $0x8c71dcd9
    WORD $0x5ee44dd9; WORD $0x07ecf0ae
    WORD $0x288e1b6f; WORD $0xaf8e5410
    WORD $0xf69d6150; WORD $0xc9e82cd9
    WORD $0x32b1a24a; WORD $0xdb71e914
    WORD $0x3a225cd2; WORD $0xbe311c08
    WORD $0x9faf056e; WORD $0x892731ac
    WORD $0x48aaf406; WORD $0x6dbd630a
    WORD $0xc79ac6ca; WORD $0xab70fe17
    WORD $0xdad5b108; WORD $0x092cbbcc
    WORD $0xb981787d; WORD $0xd64d3d9d
    WORD $0x08c58ea5; WORD $0x25bbf560
    WORD $0x93f0eb4e; WORD $0x85f04682
    WORD $0x0af6f24e; WORD $0xaf2af2b8
    WORD $0x38ed2621; WORD $0xa76c5823
    WORD $0x0db4aee1; WORD $0x1af5af66
    WORD $0x07286faa; WORD $0xd1476e2c
    WORD $0xc890ed4d; WORD $0x50d98d9f
    WORD $0x847945ca; WORD $0x82cca4db
    WORD $0xbab528a0; WORD $0xe50ff107
    WORD $0x6597973c; WORD $0xa37fce12
    WORD $0xa96272c8; WORD $0x1e53ed49
    WORD $0xfefd7d0c; WORD $0xcc5fc196
    WORD $0x13bb0f7a; WORD $0x25e8e89c
    WORD $0xbebcdc4f; WORD $0xff77b1fc
    WORD $0x8c54e9ac; WORD $0x77b19161
    WORD $0xf73609b1; WORD $0x9faacf3d
    WORD $0xef6a2417; WORD $0xd59df5b9
    WORD $0x75038c1d; WORD $0xc795830d
    WORD $0x6b44ad1d; WORD $0x4b057328
    WORD $0xd2446f25; WORD $0xf97ae3d0
    WORD $0x430aec32; WORD $0x4ee367f9
    WORD $0x836ac577; WORD $0x9becce62
    WORD $0x93cda73f; WORD $0x229c41f7
    WORD $0x244576d5; WORD $0xc2e801fb
    WORD $0x78c1110f; WORD $0x6b435275
    WORD $0xed56d48a; WORD $0xf3a20279
    WORD $0x6b78aaa9; WORD $0x830a1389
    WORD $0x345644d6; WORD $0x9845418c
    WORD $0xc656d553; WORD $0x23cc986b
    WORD $0x416bd60c; WORD $0xbe5691ef
    WORD $0xb7ec8aa8; WORD $0x2cbfbe86
    WORD $0x11c6cb8f; WORD $0xedec366b
    WORD $0x32f3d6a9; WORD $0x7bf7d714
    WORD $0xeb1c3f39; WORD $0x94b3a202
    WORD $0x3fb0cc53; WORD $0xdaf5ccd9
    WORD $0xa5e34f07; WORD $0xb9e08a83
    WORD $0x8f9cff68; WORD $0xd1b3400f
    WORD $0x8f5c22c9; WORD $0xe858ad24
    WORD $0xb9c21fa1; WORD $0x23100809
    WORD $0xd99995be; WORD $0x91376c36
    WORD $0x2832a78a; WORD $0xabd40a0c
    WORD $0x8ffffb2d; WORD $0xb5854744
    WORD $0x323f516c; WORD $0x16c90c8f
    WORD $0xb3fff9f9; WORD $0xe2e69915
    WORD $0x7f6792e3; WORD $0xae3da7d9
    WORD $0x907ffc3b; WORD $0x8dd01fad
    WORD $0xdf41779c; WORD $0x99cd11cf
    WORD $0xf49ffb4a; WORD $0xb1442798
    WORD $0xd711d583; WORD $0x40405643
    WORD $0x31c7fa1d; WORD $0xdd95317f
    WORD $0x666b2572; WORD $0x482835ea
    WORD $0x7f1cfc52; WORD $0x8a7d3eef
    WORD $0x0005eecf; WORD $0xda324365
    WORD $0x5ee43b66; WORD $0xad1c8eab
    WORD $0x40076a82; WORD $0x90bed43e
    WORD $0x369d4a40; WORD $0xd863b256
    WORD $0xe804a291; WORD $0x5a7744a6
    WORD $0xe2224e68; WORD $0x873e4f75
    WORD $0xa205cb36; WORD $0x711515d0
    WORD $0x5aaae202; WORD $0xa90de353
    WORD $0xca873e03; WORD $0x0d5a5b44
    WORD $0x31559a83; WORD $0xd3515c28
    WORD $0xfe9486c2; WORD $0xe858790a
    WORD $0x1ed58091; WORD $0x8412d999
    WORD $0xbe39a872; WORD $0x626e974d
    WORD $0x668ae0b6; WORD $0xa5178fff
    WORD $0x2dc8128f; WORD $0xfb0a3d21
    WORD $0x402d98e3; WORD $0xce5d73ff
    WORD $0xbc9d0b99; WORD $0x7ce66634
    WORD $0x881c7f8e; WORD $0x80fa687f
    WORD $0xebc44e80; WORD $0x1c1fffc1
    WORD $0x6a239f72; WORD $0xa139029f
    WORD $0x66b56220; WORD $0xa327ffb2
    WORD $0x44ac874e; WORD $0xc9874347
    WORD $0x0062baa8; WORD $0x4bf1ff9f
    WORD $0x15d7a922; WORD $0xfbe91419
    WORD $0x603db4a9; WORD $0x6f773fc3
    WORD $0xada6c9b5; WORD $0x9d71ac8f
    WORD $0x384d21d3; WORD $0xcb550fb4
    WORD $0x99107c22; WORD $0xc4ce17b3
    WORD $0x46606a48; WORD $0x7e2a53a1
    WORD $0x7f549b2b; WORD $0xf6019da0
    WORD $0xcbfc426d; WORD $0x2eda7444
    WORD $0x4f94e0fb; WORD $0x99c10284
    WORD $0xfefb5308; WORD $0xfa911155
    WORD $0x637a1939; WORD $0xc0314325
    WORD $0x7eba27ca; WORD $0x793555ab
    WORD $0xbc589f88; WORD $0xf03d93ee
    WORD $0x2f3458de; WORD $0x4bc1558b
    WORD $0x35b763b5; WORD $0x96267c75
    WORD $0xfb016f16; WORD $0x9eb1aaed
    WORD $0x83253ca2; WORD $0xbbb01b92
    WORD $0x79c1cadc; WORD $0x465e15a9
    WORD $0x23ee8bcb; WORD $0xea9c2277
    WORD $0xec191ec9; WORD $0x0bfacd89
    WORD $0x7675175f; WORD $0x92a1958a
    WORD $0x671f667b; WORD $0xcef980ec
    WORD $0x14125d36; WORD $0xb749faed
    WORD $0x80e7401a; WORD $0x82b7e127
    WORD $0x5916f484; WORD $0xe51c79a8
    WORD $0xb0908810; WORD $0xd1b2ecb8
    WORD $0x37ae58d2; WORD $0x8f31cc09
    WORD $0xdcb4aa15; WORD $0x861fa7e6
    WORD $0x8599ef07; WORD $0xb2fe3f0b
    WORD $0x93e1d49a; WORD $0x67a791e0
    WORD $0x67006ac9; WORD $0xdfbdcece
    WORD $0x5c6d24e0; WORD $0xe0c8bb2c
    WORD $0x006042bd; WORD $0x8bd6a141
    WORD $0x73886e18; WORD $0x58fae9f7
    WORD $0x4078536d; WORD $0xaecc4991
    WORD $0x506a899e; WORD $0xaf39a475
    WORD $0x90966848; WORD $0xda7f5bf5
    WORD $0x52429603; WORD $0x6d8406c9
    WORD $0x7a5e012d; WORD $0x888f9979
    WORD $0xa6d33b83; WORD $0xc8e5087b
    WORD $0xd8f58178; WORD $0xaab37fd7
    WORD $0x90880a64; WORD $0xfb1e4a9a
    WORD $0xcf32e1d6; WORD $0xd5605fcd
    WORD $0x9a55067f; WORD $0x5cf2eea0
    WORD $0xa17fcd26; WORD $0x855c3be0
    WORD $0xc0ea481e; WORD $0xf42faa48
    WORD $0xc9dfc06f; WORD $0xa6b34ad8
    WORD $0xf124da26; WORD $0xf13b94da
    WORD $0xfc57b08b; WORD $0xd0601d8e
    WORD $0xd6b70858; WORD $0x76c53d08
    WORD $0x5db6ce57; WORD $0x823c1279
    WORD $0x0c64ca6e; WORD $0x54768c4b
    WORD $0xb52481ed; WORD $0xa2cb1717
    WORD $0xcf7dfd09; WORD $0xa9942f5d
    WORD $0xa26da268; WORD $0xcb7ddcdd
    WORD $0x435d7c4c; WORD $0xd3f93b35
    WORD $0x0b090b02; WORD $0xfe5d5415
    WORD $0x4a1a6daf; WORD $0xc47bc501
    WORD $0x26e5a6e1; WORD $0x9efa548d
    WORD $0x9ca1091b; WORD $0x359ab641
    WORD $0x709f109a; WORD $0xc6b8e9b0
    WORD $0x03c94b62; WORD $0xc30163d2
    WORD $0x8cc6d4c0; WORD $0xf867241c
    WORD $0x425dcf1d; WORD $0x79e0de63
    WORD $0xd7fc44f8; WORD $0x9b407691
    WORD $0x12f542e4; WORD $0x985915fc
    WORD $0x4dfb5636; WORD $0xc2109436
    WORD $0x17b2939d; WORD $0x3e6f5b7b
    WORD $0xe17a2bc4; WORD $0xf294b943
    WORD $0xeecf9c42; WORD $0xa705992c
    WORD $0x6cec5b5a; WORD $0x979cf3ca
    WORD $0x2a838353; WORD $0x50c6ff78
    WORD $0x08277231; WORD $0xbd8430bd
    WORD $0x35246428; WORD $0xa4f8bf56
    WORD $0x4a314ebd; WORD $0xece53cec
    WORD $0xe136be99; WORD $0x871b7795
    WORD $0xae5ed136; WORD $0x940f4613
    WORD $0x59846e3f; WORD $0x28e2557b
    WORD $0x99f68584; WORD $0xb9131798
    WORD $0x2fe589cf; WORD $0x331aeada
    WORD $0xc07426e5; WORD $0xe757dd7e
    WORD $0x5def7621; WORD $0x3ff0d2c8
    WORD $0x3848984f; WORD $0x9096ea6f
    WORD $0x756b53a9; WORD $0x0fed077a
    WORD $0x065abe63; WORD $0xb4bca50b
    WORD $0x12c62894; WORD $0xd3e84959
    WORD $0xc7f16dfb; WORD $0xe1ebce4d
    WORD $0xabbbd95c; WORD $0x64712dd7
    WORD $0x9cf6e4bd; WORD $0x8d3360f0
    WORD $0x96aacfb3; WORD $0xbd8d794d
    WORD $0xc4349dec; WORD $0xb080392c
    WORD $0xfc5583a0; WORD $0xecf0d7a0
    WORD $0xf541c567; WORD $0xdca04777
    WORD $0x9db57244; WORD $0xf41686c4
    WORD $0xf9491b60; WORD $0x89e42caa
    WORD $0xc522ced5; WORD $0x311c2875
    WORD $0xb79b6239; WORD $0xac5d37d5
    WORD $0x366b828b; WORD $0x7d633293
    WORD $0x25823ac7; WORD $0xd77485cb
    WORD $0x02033197; WORD $0xae5dff9c
    WORD $0xf77164bc; WORD $0x86a8d39e
    WORD $0x0283fdfc; WORD $0xd9f57f83
    WORD $0xb54dbdeb; WORD $0xa8530886
    WORD $0xc324fd7b; WORD $0xd072df63
    WORD $0x62a12d66; WORD $0xd267caa8
    WORD $0x59f71e6d; WORD $0x4247cb9e
    WORD $0x3da4bc60; WORD $0x8380dea9
    WORD $0xf074e608; WORD $0x52d9be85
    WORD $0x8d0deb78; WORD $0xa4611653
    WORD $0x6c921f8b; WORD $0x67902e27
    WORD $0x70516656; WORD $0xcd795be8
    WORD $0xa3db53b6; WORD $0x00ba1cd8
    WORD $0x4632dff6; WORD $0x806bd971
    WORD $0xccd228a4; WORD $0x80e8a40e
    WORD $0x97bf97f3; WORD $0xa086cfcd
    WORD $0x8006b2cd; WORD $0x6122cd12
    WORD $0xfdaf7df0; WORD $0xc8a883c0
    WORD $0x20085f81; WORD $0x796b8057
    WORD $0x3d1b5d6c; WORD $0xfad2a4b1
    WORD $0x74053bb0; WORD $0xcbe33036
    WORD $0xc6311a63; WORD $0x9cc3a6ee
    WORD $0x11068a9c; WORD $0xbedbfc44
    WORD $0x77bd60fc; WORD $0xc3f490aa
    WORD $0x15482d44; WORD $0xee92fb55
    WORD $0x15acb93b; WORD $0xf4f1b4d5
    WORD $0x2d4d1c4a; WORD $0x751bdd15
    WORD $0x2d8bf3c5; WORD $0x99171105
    WORD $0x78a0635d; WORD $0xd262d45a
    WORD $0x78eef0b6; WORD $0xbf5cd546
    WORD $0x16c87c34; WORD $0x86fb8971
    WORD $0x172aace4; WORD $0xef340a98
    WORD $0xae3d4da0; WORD $0xd45d35e6
    WORD $0x0e7aac0e; WORD $0x9580869f
    WORD $0x59cca109; WORD $0x89748360
    WORD $0xd2195712; WORD $0xbae0a846
    WORD $0x703fc94b; WORD $0x2bd1a438
    WORD $0x869facd7; WORD $0xe998d258
    WORD $0x4627ddcf; WORD $0x7b6306a3
    WORD $0x5423cc06; WORD $0x91ff8377
    WORD $0x17b1d542; WORD $0x1a3bc84c
    WORD $0x292cbf08; WORD $0xb67f6455
    WORD $0x1d9e4a93; WORD $0x20caba5f
    WORD $0x7377eeca; WORD $0xe41f3d6a
    WORD $0x7282ee9c; WORD $0x547eb47b
    WORD $0x882af53e; WORD $0x8e938662
    WORD $0x4f23aa43; WORD $0xe99e619a
    WORD $0x2a35b28d; WORD $0xb23867fb
    WORD $0xe2ec94d4; WORD $0x6405fa00
    WORD $0xf4c31f31; WORD $0xdec681f9
    WORD $0x8dd3dd04; WORD $0xde83bc40
    WORD $0x38f9f37e; WORD $0x8b3c113c
    WORD $0xb148d445; WORD $0x9624ab50
    WORD $0x4738705e; WORD $0xae0b158b
    WORD $0xdd9b0957; WORD $0x3badd624
    WORD $0x19068c76; WORD $0xd98ddaee
    WORD $0x0a80e5d6; WORD $0xe54ca5d7
    WORD $0xcfa417c9; WORD $0x87f8a8d4
    WORD $0xcd211f4c; WORD $0x5e9fcf4c
    WORD $0x038d1dbc; WORD $0xa9f6d30a
    WORD $0x0069671f; WORD $0x7647c320
    WORD $0x8470652b; WORD $0xd47487cc
    WORD $0x0041e073; WORD $0x29ecd9f4
    WORD $0xd2c63f3b; WORD $0x84c8d4df
    WORD $0x00525890; WORD $0xf4681071
    WORD $0xc777cf09; WORD $0xa5fb0a17
    WORD $0x4066eeb4; WORD $0x7182148d
    WORD $0xb955c2cc; WORD $0xcf79cc9d
    WORD $0x48405530; WORD $0xc6f14cd8
    WORD $0x93d599bf; WORD $0x81ac1fe2
    WORD $0x5a506a7c; WORD $0xb8ada00e
    WORD $0x38cb002f; WORD $0xa21727db
    WORD $0xf0e4851c; WORD $0xa6d90811
    WORD $0x06fdc03b; WORD $0xca9cf1d2
    WORD $0x6d1da663; WORD $0x908f4a16
    WORD $0x88bd304a; WORD $0xfd442e46
    WORD $0x043287fe; WORD $0x9a598e4e
    WORD $0x15763e2e; WORD $0x9e4a9cec
    WORD $0x853f29fd; WORD $0x40eff1e1
    WORD $0x1ad3cdba; WORD $0xc5dd4427
    WORD $0xe68ef47c; WORD $0xd12bee59
    WORD $0xe188c128; WORD $0xf7549530
    WORD $0x301958ce; WORD $0x82bb74f8
    WORD $0x8cf578b9; WORD $0x9a94dd3e
    WORD $0x3c1faf01; WORD $0xe36a5236
    WORD $0x3032d6e7; WORD $0xc13a148e
    WORD $0xcb279ac1; WORD $0xdc44e6c3
    WORD $0xbc3f8ca1; WORD $0xf18899b1
    WORD $0x5ef8c0b9; WORD $0x29ab103a
    WORD $0x15a7b7e5; WORD $0x96f5600f
    WORD $0xf6b6f0e7; WORD $0x7415d448
    WORD $0xdb11a5de; WORD $0xbcb2b812
    WORD $0x3464ad21; WORD $0x111b495b
    WORD $0x91d60f56; WORD $0xebdf6617
    WORD $0x00beec34; WORD $0xcab10dd9
    WORD $0xbb25c995; WORD $0x936b9fce
    WORD $0x40eea742; WORD $0x3d5d514f
    WORD $0x69ef3bfb; WORD $0xb84687c2
    WORD $0x112a5112; WORD $0x0cb4a5a3
    WORD $0x046b0afa; WORD $0xe65829b3
    WORD $0xeaba72ab; WORD $0x47f0e785
    WORD $0xe2c2e6dc; WORD $0x8ff71a0f
    WORD $0x65690f56; WORD $0x59ed2167
    WORD $0xdb73a093; WORD $0xb3f4e093
    WORD $0x3ec3532c; WORD $0x306869c1
    WORD $0xd25088b8; WORD $0xe0f218b8
    WORD $0xc73a13fb; WORD $0x1e414218
    WORD $0x83725573; WORD $0x8c974f73
    WORD $0xf90898fa; WORD $0xe5d1929e
    WORD $0x644eeacf; WORD $0xafbd2350
    WORD $0xb74abf39; WORD $0xdf45f746
    WORD $0x7d62a583; WORD $0xdbac6c24
    WORD $0x328eb783; WORD $0x6b8bba8c
    WORD $0xce5da772; WORD $0x894bc396
    WORD $0x3f326564; WORD $0x066ea92f
    WORD $0x81f5114f; WORD $0xab9eb47c
    WORD $0x0efefebd; WORD $0xc80a537b
    WORD $0xa27255a2; WORD $0xd686619b
    WORD $0xe95f5f36; WORD $0xbd06742c
    WORD $0x45877585; WORD $0x8613fd01
    WORD $0x23b73704; WORD $0x2c481138
    WORD $0x96e952e7; WORD $0xa798fc41
    WORD $0x2ca504c5; WORD $0xf75a1586
    WORD $0xfca3a7a0; WORD $0xd17f3b51
    WORD $0xdbe722fb; WORD $0x9a984d73
    WORD $0x3de648c4; WORD $0x82ef8513
    WORD $0xd2e0ebba; WORD $0xc13e60d0
    WORD $0x0d5fdaf5; WORD $0xa3ab6658
    WORD $0x079926a8; WORD $0x318df905
    WORD $0x10b7d1b3; WORD $0xcc963fee
    WORD $0x497f7052; WORD $0xfdf17746
    WORD $0x94e5c61f; WORD $0xffbbcfe9
    WORD $0xedefa633; WORD $0xfeb6ea8b
    WORD $0xfd0f9bd3; WORD $0x9fd561f1
    WORD $0xe96b8fc0; WORD $0xfe64a52e
    WORD $0x7c5382c8; WORD $0xc7caba6e
    WORD $0xa3c673b0; WORD $0x3dfdce7a
    WORD $0x1b68637b; WORD $0xf9bd690a
    WORD $0xa65c084e; WORD $0x06bea10c
    WORD $0x51213e2d; WORD $0x9c1661a6
    WORD $0xcff30a62; WORD $0x486e494f
    WORD $0xe5698db8; WORD $0xc31bfa0f
    WORD $0xc3efccfa; WORD $0x5a89dba3
    WORD $0xdec3f126; WORD $0xf3e2f893
    WORD $0x5a75e01c; WORD $0xf8962946
    WORD $0x6b3a76b7; WORD $0x986ddb5c
    WORD $0xf1135823; WORD $0xf6bbb397
    WORD $0x86091465; WORD $0xbe895233
    WORD $0xed582e2c; WORD $0x746aa07d
    WORD $0x678b597f; WORD $0xee2ba6c0
    WORD $0xb4571cdc; WORD $0xa8c2a44e
    WORD $0x40b717ef; WORD $0x94db4838
    WORD $0x616ce413; WORD $0x92f34d62
    WORD $0x50e4ddeb; WORD $0xba121a46
