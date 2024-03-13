#include "scanning.h"

long validate_one(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    fsm_init(m, FSM_VAL);
    return fsm_exec_1(m, src, p, flags);
}
