
#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

/*
 This function should only used in user goroutine, it will spin the P until finish.
 TODO: support signal
 func __c_systemstack_1arg(
        arg unsafe.Pointer
        fn  unsafe.Pointer
    ) uintptr
*/
TEXT ·__c_systemstack_1arg(SB), NOSPLIT, $16-24
	MOVD    arg0+0(FP), R0
    MOVD    fn+8(FP),   R1
    MOVD	48(g),      R2	// R2 = g.m
	MOVD	0(R2),      R3	// R3 = m.g0

    // store caller's fp and lr
    STP     (R29, R30), 16(RSP)

	// switch to g0
	MOVD	R3, g
	MOVD	0x38(g), R4 // g_sched+gobuf_sp

    // call cb in system stack
    MOVD    RSP, R29    // mov fp, sp
    MOVD	R4, RSP
    WORD    $0xd63f0020 // blr x2

    // restore sp
    MOVD    R29, RSP    // mov sp, fp

    // restore caller's fp and lr
    LDP     16(RSP), (R29, R30)

    // back to usrg
    MOVD	0x30(g),     R2 // R2 = g_m
    MOVD	0xc0(R2),    g  // g = m_curg
 
    // restore ret
    MOVD    R0, ret+16(FP)
	RET


/*
 This function should only used in user goroutine, it will spin the P until finish.
 func __c_systemstack_1arg(
        arg unsafe.Pointer
        fn  unsafe.Pointer
    ) uintptr
*/
TEXT ·__c_systemstack_2arg(SB), NOSPLIT, $16-24
	MOVD    arg0+0(FP), R0
    MOVD    fn+8(FP),   R1
    MOVD	48(g),      R2	// R2 = g.m
	MOVD	0(R2),      R3	// R3 = m.g0

    // store caller's fp and lr
    STP     (R29, R30), 16(RSP)

	// switch to g0
	MOVD	R3, g
	MOVD	0x38(g), R4 // g_sched+gobuf_sp

    // call cb in system stack
    MOVD    RSP, R29    // mov fp, sp
    MOVD	R4, RSP
    WORD    $0xd63f0020 // blr x2

    // restore sp
    MOVD    R29, RSP    // mov sp, fp

    // restore caller's fp and lr
    LDP     16(RSP), (R29, R30)

    // back to usrg
    MOVD	0x30(g),     R2 // R2 = g_m
    MOVD	0xc0(R2),    g  // g = m_curg
 
    // restore ret
    MOVD    R0, ret+16(FP)
	RET