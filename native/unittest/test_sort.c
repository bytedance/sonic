#include <assert.h>
#include <string.h>
#include "../sort.c"

void test_heap_sort(MapPair* input, MapPair* expect, int n) {
    heap_sort(input, n);
    for(int i=0; i<n; i++) {
        assert(strcmp(input[i].key.buf, expect[i].key.buf) == 0);
    }
}

void test_insert_sort(MapPair* input, MapPair* expect, int n) {
    insert_sort(input, n);
    for(int i=0; i<n; i++) {
        assert(strcmp(input[i].key.buf, expect[i].key.buf) == 0);
    }
}

int main() {
    MapPair mpr[5];

    mpr[0].key.buf = "abcd";
    mpr[0].key.len = 4;
    
    mpr[1].key.buf = "abcdf";
    mpr[1].key.len = 5;

    mpr[2].key.buf = "bbcf";
    mpr[2].key.len = 4;
    
    mpr[3].key.buf = "ma";
    mpr[3].key.len = 2;

    mpr[4].key.buf = "abcd";
    mpr[4].key.len = 4;

    MapPair expect[5];
    expect[0].key.buf = "abcd";
    expect[0].key.len = 4;
    
    expect[1].key.buf = "abcd";
    expect[1].key.len = 4;

    expect[2].key.buf = "abcdf";
    expect[2].key.len = 5;
    
    expect[3].key.buf = "bbcf";
    expect[3].key.len = 4;

    expect[4].key.buf = "ma";
    expect[4].key.len = 2;

    MapPair mpr_2[5];
    memcpy(mpr_2, mpr, 5 * sizeof(MapPair));

    test_heap_sort(mpr, expect, 5);   
    test_insert_sort(mpr_2, expect, 5); 
}