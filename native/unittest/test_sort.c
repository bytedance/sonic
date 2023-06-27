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
    char temp[6] = "abcdf";
    mpr[0].key.buf = "abcd";
    mpr[0].key.len = 5;
    
    mpr[1].key.buf = "abcdf";
    mpr[1].key.len = 6;

    mpr[2].key.buf = "测试字段";
    mpr[2].key.len = 13;
    
    mpr[3].key.buf = "";
    mpr[3].key.len = 1;

    mpr[4].key.buf = temp;
    mpr[4].key.len = 5;

    MapPair expect[5];

    expect[0].key.buf = "";
    expect[0].key.len = 0;
    
    expect[1].key.buf = "abcd";
    expect[1].key.len = 5;

    expect[2].key.buf = temp;
    expect[2].key.len = 5;
    
    expect[3].key.buf = "abcdf";
    expect[3].key.len = 6;

    expect[4].key.buf = "测试字段";
    expect[4].key.len = 13;

    MapPair mpr_2[5];
    memcpy(mpr_2, mpr, 5 * sizeof(MapPair));

    test_heap_sort(mpr, expect, 5);   
    test_insert_sort(mpr_2, expect, 5); 

    // test when n is a large number
    MapPair mpr_n_heap[4096];
    MapPair mpr_n_insert[4096];
    MapPair expect_n[4096];
    for(int i=0; i<4096; i++) {
        mpr_n_heap[i].key.buf = "abcda";
        mpr_n_heap[i].key.len = 6;

        expect_n[i].key.buf = "abcda";
        expect_n[i].key.len = 6;
    }
    mpr_n_heap[2000].key.buf = "";
    mpr_n_heap[2000].key.len = 1;
    mpr_n_heap[1000].key.buf = "aabb";
    mpr_n_heap[1000].key.len = 5;
    mpr_n_heap[3000].key.buf = "ss";
    mpr_n_heap[3000].key.len = 3;
    mpr_n_heap[4000].key.buf = "中文字符";
    mpr_n_heap[4000].key.len = 13;
    mpr_n_heap[4095].key.buf = "";
    mpr_n_heap[4095].key.len = 1;

    expect_n[0].key.buf = "";
    expect_n[0].key.len = 1;
    expect_n[1].key.buf = "";
    expect_n[1].key.len = 1;
    expect_n[2].key.buf = "aabb";
    expect_n[2].key.len = 5;
    expect_n[2000].key.buf = "abcda";
    expect_n[2000].key.len = 6;
    expect_n[1000].key.buf = "abcda";
    expect_n[1000].key.len = 6;
    expect_n[3000].key.buf = "abcda";
    expect_n[3000].key.len = 6;
    expect_n[4000].key.buf = "abcda";
    expect_n[4000].key.len = 6;
    expect_n[4094].key.buf = "ss";
    expect_n[4094].key.len = 3;
    expect_n[4095].key.buf = "中文字符";
    expect_n[4095].key.len = 13;

    memcpy(mpr_n_insert, mpr_n_heap, 4096 * sizeof(MapPair));
    test_heap_sort(mpr_n_heap, expect_n, 4096);   
    test_insert_sort(mpr_n_insert, expect_n, 4096); 
}