// +build !noasm !appengine !amd64

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·MoreStack(SB), NOSPLIT, $0 - 8
    NO_LOCAL_POINTERS
    RET
