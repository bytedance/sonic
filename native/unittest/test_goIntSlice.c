#include "../goIntSlice.c"

#include <stdio.h>
#include <malloc.h>
#include <stdbool.h>
#include <stdint.h>
#include <stddef.h>
#include <string.h>



int main(){
    // ��ʼ��GoIntSlice
    GoIntSlice demo;
    size_t cap = 4096;
    //demo.uptr = (uint64_t*)malloc(cap * sizeof(uint64_t));
    demo.iptr = (int64_t*)malloc(cap * sizeof(int64_t));
    demo.len = 0;
    demo.cap = cap;

    // long* pp;  // ������һ��Ұָ�룬����û�г�ʼ����
    long p = 0;
    long* pp = &p; 
    GoString test;
    /*
    Ŀǰ��������������У�
	 "{[1,2,3,4]"  �������� 
	 "[1,2,3.5,4]"  ��������
     "[1,  2,3, 4]"  ��ȷ���������в�����ո�
     "[  1,2,3,4]"   ��ȷ���������в�����ո�
	 "[1,2,3,4]"  ��ȷ���� 
     "[1,2,3]"    ��ȷ���� 
     "[1  ,2,3,4]"   ��ȷ���������в�����ո�
    */
    
    
    const char* str = "[1,-2,3,-4,5]"; // ʹ��C�����ַ���д�����Ӽ�࣬���ҽ�β�Զ�������'\0'
    test.buf = str;
    test.len = strlen(str);
	
	
	long res = decode_i64_array(&test,pp,&demo);
    //long res = decode_u64_array(&test,pp,&demo);
    
	printf("%lld ,%lld\n",*pp,res);
    for(int z=0;z<demo.len;z++){
    	printf("%d   ",demo.iptr[z]);
	}

    // �������ͷ��ڴ�
    free(demo.uptr);
}
