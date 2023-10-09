#include "scanning.h"

long skip_array(const GoString *src, long *p, StateMachine *m, uint64_t flags) {
    fsm_init(m, FSM_ARR_0);
    return fsm_exec_1(m, src, p, flags);
}
