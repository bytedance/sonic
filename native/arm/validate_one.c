#include "scanning.h"

long validate_one(const GoString *src, long *p, StateMachine *m) {
    fsm_init(m, FSM_VAL);
    return fsm_exec_1(m, src, p, MASK_VALIDATE_STRING);
}
