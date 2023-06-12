#include "../to_lower.c"
#include<assert.h>

void test_to_lower(const char* input, const char* expect) {
    int len = sizeof(input);
    char* dst = (char*)malloc(len);
    to_lower(dst, input, len);
    assert(strncmp(expect, dst) == 0);
    free(dst);
}

int main() {
    test_to_lower("Hello, World!", "hello, world!");
    test_to_lower("12345", "12345");
    test_to_lower("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz");
}