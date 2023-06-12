#include "../goIntSlice.c"

#include <stdio.h>
#include <malloc.h>
#include <stdbool.h>
#include <stdint.h>
#include <stddef.h>
#include <string.h>



int main(){
    // 初始化GoIntSlice
    GoIntSlice demo;
    size_t cap = 4096;
    //demo.uptr = (uint64_t*)malloc(cap * sizeof(uint64_t));
    demo.iptr = (int64_t*)malloc(cap * sizeof(int64_t));
    demo.len = 0;
    demo.cap = cap;

    // long* pp;  // 这种是一个野指针，而且没有初始化，
    long p = 0;
    long* pp = &p; 
    GoString test;
    /*
    目前测试用例满足的有：
	 "{[1,2,3,4]"  错误用例 
	 "[1,2,3.5,4]"  错误用例
     "[1,  2,3, 4]"  正确用例但是有不规则空格
     "[  1,2,3,4]"   正确用例但是有不规则空格
	 "[1,2,3,4]"  正确用例 
     "[1,2,3]"    正确用例 
     "[1  ,2,3,4]"   正确用例但是有不规则空格
    */
    
    
    const char* str = "[1,-2,3,-4,5]"; // 使用C语言字符串写法更加简洁，而且结尾自动加上了'\0'
    test.buf = str;
    test.len = strlen(str);
	
	
	long res = decode_i64_array(&test,pp,&demo);
    //long res = decode_u64_array(&test,pp,&demo);
    
	printf("%lld ,%lld\n",*pp,res);
    for(int z=0;z<demo.len;z++){
    	printf("%d   ",demo.iptr[z]);
	}

    // 测试完释放内存
    free(demo.uptr);
}
