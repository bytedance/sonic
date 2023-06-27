#include <stdio.h>
#include <string.h>

#include "../goIntSlice.c"

void test_decode_i64_array(const char* test_str){
    GoIntSlice demo;
    size_t cap = 4096;
    demo.iptr = (int64_t*)malloc(cap * sizeof(int64_t));
    demo.len = 0;
    demo.cap = cap;
    
    long p = 0;
    long* pp = &p; 
    GoString test;
    
    const char* str = test_str; 
    test.buf = str;
    test.len = strlen(str);
	
    long res = decode_i64_array(&test,pp,&demo);
    
    printf("%lld ,%lld\n",*pp,res);
	
    for(int z=0;z<demo.len;z++){
    	printf(" %d ",demo.iptr[z]);    	
	}	
    free(demo.iptr);	
}

void test_decode_u64_array(const char* test_str){
    GoIntSlice demo;
    size_t cap = 4096;
    demo.uptr = (uint64_t*)malloc(cap * sizeof(uint64_t));
    demo.len = 0;
    demo.cap = cap;
    
    long p = 0;
    long* pp = &p; 
    GoString test;
    
    const char* str = test_str; 
    test.buf = str;
    test.len = strlen(str);
	
    long res = decode_u64_array(&test,pp,&demo);
    
    printf("%lld ,%lld\n",*pp,res);
	
    for(int z=0;z<demo.len;z++){
    	printf(" %d ",demo.uptr[z]);    	
	}	
    free(demo.uptr);
}

int main(){     
    char teststr[10][100] = {"   ","[1,20,3,4],","[1,2,3.5,4]","[1,20,3,4]","[  1,2,3,4]","[1,2,3,4]","[1,2,3]","[1  ,2,3,4]","[1,-2,-3,4]","[1,-2,  -3, 4]","[1, -2.3, 4]"};
    
    for(int i=0;i<11;i++){
    	test_decode_u64_array(teststr[i]);
    	printf("*");
    	test_decode_i64_array(teststr[i]);
    	printf("*");
    }    
}
