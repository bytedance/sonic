#include "../native.c"
#include <assert.h>
#include <string.h>

void test_to_lower(const char* input, const char* expect) {
    char* dst;
    int len = sizeof(input);
    to_lower(dst, input, len);
    assert(strcmp(expect, buf) == 0);
}

int main() {
    test_u64toa("Hello, World!", "hello, world!");
    test_u64toa("12345", "12345");
    test_u64toa("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZ");
}