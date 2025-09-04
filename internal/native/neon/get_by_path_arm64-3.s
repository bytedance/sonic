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
    BEQ LBB0_910
LBB0_908:
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
    BNE LBB0_908
LBB0_910:
    WORD $0x9e670078 // fmov    d24, x3
    WORD $0xd2407c42 // eor    x2, x2, #0xffffffff
    WORD $0x8a120052 // and    x18, x2, x18
    WORD $0xea110051 // ands    x17, x2, x17
    WORD $0x0e205b18 // cnt    v24.8b, v24.8b
    WORD $0x2e303b18 // uaddlv    h24, v24.8b
    WORD $0x1e260303 // fmov    w3, s24
    WORD $0x8b0e006e // add    x14, x3, x14
    BEQ LBB0_913
LBB0_911:
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
    BNE LBB0_911
LBB0_913:
    WORD $0x9e670258 // fmov    d24, x18
    WORD $0x937ffe10 // asr    x16, x16, #63
    WORD $0x9101014a // add    x10, x10, #64
    WORD $0x0e205b18 // cnt    v24.8b, v24.8b
    WORD $0x2e303b18 // uaddlv    h24, v24.8b
    WORD $0x1e26030c // fmov    w12, s24
    WORD $0x8b0e018e // add    x14, x12, x14
    WORD $0xaa0f03ec // mov    x12, x15
    WORD $0xf10101ef // subs    x15, x15, #64
    BGE LBB0_904
    B LBB0_888
LBB0_914:
    WORD $0xf9400408 // ldr    x8, [x0, #8]
    WORD $0xdac0008a // rbit    x10, x4
    B LBB0_918
LBB0_915:
    WORD $0x5ac001ea // rbit    w10, w15
    WORD $0x8b080128 // add    x8, x9, x8
    WORD $0x5ac0114a // clz    w10, w10
    WORD $0x8b0a0108 // add    x8, x8, x10
    WORD $0x9100090a // add    x10, x8, #2
LBB0_916:
    WORD $0xaa0903e8 // mov    x8, x9
    WORD $0xf900002a // str    x10, [x1]
    B LBB0_868
LBB0_917:
    WORD $0xf9400408 // ldr    x8, [x0, #8]
    WORD $0xdac0022a // rbit    x10, x17
LBB0_918:
    WORD $0xdac0114a // clz    x10, x10
    WORD $0xcb0c014a // sub    x10, x10, x12
    WORD $0x8b080148 // add    x8, x10, x8
    WORD $0x9100050a // add    x10, x8, #1
    WORD $0xf900002a // str    x10, [x1]
    WORD $0xf940040b // ldr    x11, [x0, #8]
    WORD $0xeb0b015f // cmp    x10, x11
    WORD $0x9a88256a // csinc    x10, x11, x8, hs
    WORD $0xda9f9128 // csinv    x8, x9, xzr, ls
    WORD $0xf900002a // str    x10, [x1]
    B LBB0_868
LBB0_919:
    WORD $0x92800028 // mov    x8, #-2
    WORD $0x5280004d // mov    w13, #2
    WORD $0x8b0d018c // add    x12, x12, x13
    WORD $0xab0b010b // adds    x11, x8, x11
    WORD $0x92800008 // mov    x8, #-1
    BLE LBB0_868
LBB0_920:
    WORD $0x39400188 // ldrb    w8, [x12]
    WORD $0x7101711f // cmp    w8, #92
    BEQ LBB0_919
    WORD $0x7100891f // cmp    w8, #34
    BEQ LBB0_953
    WORD $0x92800008 // mov    x8, #-1
    WORD $0x5280002d // mov    w13, #1
    WORD $0x8b0d018c // add    x12, x12, x13
    WORD $0xab0b010b // adds    x11, x8, x11
    WORD $0x92800008 // mov    x8, #-1
    BGT LBB0_920
    B LBB0_868
LBB0_923:
    WORD $0x928000c8 // mov    x8, #-7
    B LBB0_868
LBB0_924:
    WORD $0xb100051f // cmn    x8, #1
    BNE LBB0_926
LBB0_925:
    WORD $0x92800008 // mov    x8, #-1
    WORD $0xaa1603f8 // mov    x24, x22
LBB0_926:
    WORD $0xf9000038 // str    x24, [x1]
    B LBB0_868
LBB0_927:
    WORD $0x92800007 // mov    x7, #-1
    B LBB0_933
LBB0_928:
    WORD $0xcb0a010a // sub    x10, x8, x10
    B LBB0_916
LBB0_929:
    WORD $0xd1000508 // sub    x8, x8, #1
    B LBB0_868
LBB0_930:
    WORD $0xaa3a03e7 // mvn    x7, x26
    B LBB0_933
LBB0_931:
    WORD $0x92800008 // mov    x8, #-1
    WORD $0xf9000026 // str    x6, [x1]
    B LBB0_868
LBB0_932:
    WORD $0xaa3803e7 // mvn    x7, x24
LBB0_933:
    WORD $0xcb0702c8 // sub    x8, x22, x7
    WORD $0xd1000909 // sub    x9, x8, #2
    B LBB0_866
LBB0_934:
    WORD $0xaa1403e5 // mov    x5, x20
LBB0_935:
    WORD $0x92800008 // mov    x8, #-1
    WORD $0xf9000025 // str    x5, [x1]
    B LBB0_868
LBB0_936:
    WORD $0x9280001d // mov    x29, #-1
    B LBB0_956
LBB0_937:
    WORD $0x12001cca // and    w10, w6, #0xff
    WORD $0x7101855f // cmp    w10, #97
    BNE LBB0_951
    WORD $0x9100050a // add    x10, x8, #1
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a692a // ldrb    w10, [x9, x10]
    WORD $0x7101b15f // cmp    w10, #108
    BNE LBB0_951
    WORD $0x9100090a // add    x10, x8, #2
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a692a // ldrb    w10, [x9, x10]
    WORD $0x7101cd5f // cmp    w10, #115
    BNE LBB0_951
    WORD $0x91000d0a // add    x10, x8, #3
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a6929 // ldrb    w9, [x9, x10]
    WORD $0x7101953f // cmp    w9, #101
    BNE LBB0_951
    WORD $0x91001109 // add    x9, x8, #4
    B LBB0_866
LBB0_942:
    WORD $0xd100050a // sub    x10, x8, #1
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a692a // ldrb    w10, [x9, x10]
    WORD $0x7101b95f // cmp    w10, #110
    BNE LBB0_951
    WORD $0xf9000028 // str    x8, [x1]
    WORD $0x3868692a // ldrb    w10, [x9, x8]
    WORD $0x7101d55f // cmp    w10, #117
    BNE LBB0_951
    WORD $0x9100050a // add    x10, x8, #1
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a692a // ldrb    w10, [x9, x10]
    WORD $0x7101b15f // cmp    w10, #108
    BNE LBB0_951
    WORD $0x9100090a // add    x10, x8, #2
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a6929 // ldrb    w9, [x9, x10]
    WORD $0x7101b13f // cmp    w9, #108
    BNE LBB0_951
    B LBB0_950
LBB0_946:
    WORD $0xd100050a // sub    x10, x8, #1
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a692a // ldrb    w10, [x9, x10]
    WORD $0x7101d15f // cmp    w10, #116
    BNE LBB0_951
    WORD $0xf9000028 // str    x8, [x1]
    WORD $0x3868692a // ldrb    w10, [x9, x8]
    WORD $0x7101c95f // cmp    w10, #114
    BNE LBB0_951
    WORD $0x9100050a // add    x10, x8, #1
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a692a // ldrb    w10, [x9, x10]
    WORD $0x7101d55f // cmp    w10, #117
    BNE LBB0_951
    WORD $0x9100090a // add    x10, x8, #2
    WORD $0xf900002a // str    x10, [x1]
    WORD $0x386a6929 // ldrb    w9, [x9, x10]
    WORD $0x7101953f // cmp    w9, #101
    BNE LBB0_951
LBB0_950:
    WORD $0x91000d09 // add    x9, x8, #3
    B LBB0_866
LBB0_951:
    WORD $0x92800028 // mov    x8, #-2
    B LBB0_868
LBB0_952:
    WORD $0xd10006e8 // sub    x8, x23, #1
    B LBB0_868
LBB0_953:
    WORD $0xcb0a0188 // sub    x8, x12, x10
    WORD $0x9100050a // add    x10, x8, #1
    B LBB0_916
LBB0_954:
    WORD $0xaa3903fd // mvn    x29, x25
    B LBB0_956
LBB0_955:
    WORD $0xaa3703fd // mvn    x29, x23
LBB0_956:
    WORD $0xaa3d03e9 // mvn    x9, x29
    WORD $0x8b090109 // add    x9, x8, x9
    B LBB0_866
LBB0_957:
    WORD $0xaa1703f6 // mov    x22, x23
    B LBB0_925
LBB0_958:
    WORD $0x92800008 // mov    x8, #-1
    WORD $0xcb100229 // sub    x9, x17, x16
    B LBB0_867
LBB0_959:
    WORD $0x91000631 // add    x17, x17, #1
    WORD $0x92800048 // mov    x8, #-3
    WORD $0xcb100229 // sub    x9, x17, x16
    B LBB0_867
LBB0_960:
    WORD $0x8b0d014c // add    x12, x10, x13
    B LBB0_886
LBB0_961:
    WORD $0xf9400409 // ldr    x9, [x0, #8]
    WORD $0x92800008 // mov    x8, #-1
    B LBB0_867
LBB0_962:
    WORD $0xd100056c // sub    x12, x11, #1
    WORD $0xeb08019f // cmp    x12, x8
    BEQ LBB0_806
    WORD $0x8b09014c // add    x12, x10, x9
    WORD $0x8b08018c // add    x12, x12, x8
    WORD $0xcb080168 // sub    x8, x11, x8
    WORD $0x9100098c // add    x12, x12, #2
    WORD $0xd100090b // sub    x11, x8, #2
    B LBB0_886
MASK_USE_NUMBER:
    WORD $0x00000002 // .long    2
_UnquoteTab:
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00220000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x2F000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x0000005C
    WORD $0x00080000
    WORD $0x000C0000
    WORD $0x00000000
    WORD $0x000A0000
    WORD $0x000D0000
    WORD $0x0000FF09
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
    WORD $0x00000000
TEXT ·__get_by_path(SB), NOSPLIT, $0-40
	NO_LOCAL_POINTERS

_entry:
	MOVD 16(g), R16
	SUB $288, RSP, R17
	CMP  R16, R17
	BLS  _stack_grow

_get_by_path:
	MOVD s+0(FP), R0
	MOVD p+8(FP), R1
	MOVD path+16(FP), R2
	MOVD m+24(FP), R3
	MOVD ·_subr__get_by_path(SB), R11
	WORD $0x1000005e // adr x30, .+8
	JMP (R11)
	MOVD R0, ret+32(FP)
	RET

_stack_grow:
	MOVD R30, R3
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry

