#include "../native.c"
#include <assert.h>
#include <string.h>

void test_skipnumber(const char* number, size_t nb, size_t num_len) {
    assert(do_skip_number(number, nb) == num_len);
}

int main() {
    const char* buf = "2[";
    test_skipnumber(buf, 1, 1);
}