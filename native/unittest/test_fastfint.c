#include "native.c"
#include <assert.h>
#include <string.h>

void test_u64toa(uint64_t input, const char* expect) {
    char buf[64] = {0};
    u64toa(buf, input);
    assert(strcmp(expect, buf) == 0);
}

int main() {
    test_u64toa(0, "0");
    test_u64toa(1, "1");
    test_u64toa(1345, "1345");
    // test max uint64
    test_u64toa(18446744073709551615ULL, "18446744073709551615");
}