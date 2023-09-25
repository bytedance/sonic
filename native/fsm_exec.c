#include "scanning.h"
long fsm_exec(StateMachine *self, const GoString *src, long *p, uint64_t flags) {
    return fsm_exec_1(self, src, p, flags);
}