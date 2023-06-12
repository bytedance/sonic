#include <stdio.h>
#include <malloc.h>
#include <stdbool.h>
#include <stdint.h>
#include <stddef.h>
#include <string.h>

#include <native.h>

#define ERR_INVAL       2
#define ERR_RECURSE_MAX 7

typedef struct {
    union { // u64 或 i64 array 的指针
        int64_t*  iptr;
        uint64_t* uptr; 
    };
    size_t len; // s 长度
    size_t cap; // slice 容量
} GoIntSlice;



bool isSpace(char a){
    if( a == 32){
        return true;
    }else{
        return false;
    }
}

bool isIntger(char a){
	if(a<48 ||a>57){
		return false;
	}else{
		return true;
	}
}

int charToNum(char c){
	return c-48;
}
// 输入参数：src 是 输入的json, p 是当前位置.
// 返回值：
// 如果是合法的json 数组，返回0 表示解析成功，
// 如果解析的元素是非法json，返回-ERR_INVAL
// 如果解析的元素是json 类型不匹配，但是 json 是合法的，返回 -ERR_MISMATCH
// 如果 slice cap 不够，返回 -ERR_RECURSE_MAX。
// 将解析结果存放到arr中，len +1，直到cap。

// 例子1
// 输入:
// src： 一个字符串，是json整数数组，例如"[1,2,3,4]" 
// p: 当前的位置，假如是0, 
// arr: 假如是len=0, cap=256
// 输出：
// p: 解析后指向的位置，这里应该是json结束位置，即9
// arr: uptr指针依次存入1,2,3，4. 然后len = 4, cap不变。
// 返回值是0

// 例子2
// 输入：
// src： 一个字符串，例如"{[]}"，和 "[1,2,3,4.5], 这个json片段是非法的或者不是数字
// p: 当前的位置，假如是0, 
// arr: 假如是len=0, cap=256
// 输出：
// p: 解析后指向的位置，这里应该是json出错位置
// arr: 解析后发现是非法的json，需要返回-ERR_INVALID，同时arr中长度需要重置为0
// 返回值是 ERR_INVAL

long decode_u64_array( const GoString* src, long* p, GoIntSlice* arr){ //无符号数字
    char* pos = src->buf;
    int i =0;
    
    while(isSpace(*(pos+i))){      //如果最开始前面有空格，先吃掉空格 
    	i++;                           
	}
    if(*(pos+i) != 91){  //第一个不是左中括号的话直接返回非法
        *p = i+1;    //p指向第一个出错后位置 
        arr->len = 0;
        return ERR_INVAL;
    }
    i++;  //是左括号
    int k =0;   //需要一个游标k,m
    int num = 0; //存储数字 
    while(*(pos+i) !=0){
        if(k==arr->cap){   //容量不够
        	*p = i+1;
            return ERR_RECURSE_MAX;
        }
        while(isSpace(*(pos+i))){   //是空格就往后跳 
        	i++;             
		}
		if(*(pos+i)<48 || *(pos+i)>57){
			*p = i+1;    //p指向第一个出错后位置 
        	arr->len = 0;
			return ERR_INVAL;                    //第一个不是数字肯定非法 
		}else{
			num = charToNum(*(pos+i));
		} 
		    
		i++; 
		while(!isSpace(*(pos+i))&& *(pos+i) !=44){   //后面不是空格也不是逗号, 那么说明是数字或者发生错误 
		 	if(isIntger(*(pos+i))){
		 		
		 		num = num*10 + charToNum(*(pos+i));
        		i++;		
			}else if(*(pos+i) ==93){             //右括号的话可以收起来 ，并返回值 
				(arr->uptr)[k] = num;
				arr->len = k+1;
				*p = i+1;
				return 0;
			}else{
				*p = i+1;               //不是数字不是逗号就出错了  p指向出错后位置 
        		arr->len = 0;
        		return ERR_INVAL;
			} 
	}
		while(isSpace(*(pos+i))){
			i++;
		}
		if(*(pos+i) ==44){          //是逗号的话收起来 
			(arr->uptr)[k] = num; 
			k++;
			i++;
		}
			 
			             
    }
    

  

}



long decode_i64_array(const GoString* src, long* p, GoIntSlice* arr){//有符号数字
	char* pos = src->buf;
	int i =0;
	    
	while(isSpace(*(pos+i))){      //如果最开始前面有空格，先吃掉空格 
	    i++;                           
	}
    if(*(pos+i) != 91){  //第一个不是左中括号的话直接返回非法
        *p = i+1;    //p指向第一个出错后位置 
        arr->len = 0;
        return ERR_INVAL;
    }
    i++;  //是左括号
    int k =0;     //需要一个游标k,m
    int num = 0;    //存储数字 
    char flag ='+';   //把flag初始化为正号 
    while(*(pos+i) !=0){
        if(k==arr->cap){   //容量不够
        	*p = i+1;
            return ERR_RECURSE_MAX;
        }
        while(isSpace(*(pos+i))){   //是空格就往后跳 
        	i++;             
		}
		if((*(pos+i)<48 || *(pos+i)>57) && (*(pos+i)!=43 && *(pos+i)!=45 )){
			
			*p = i+1;    //p指向第一个出错后位置 
        	arr->len = 0;
			return ERR_INVAL;                    //第一个既不是数字，又不是正负号，那么肯定非法 
		}else if(*(pos+i)==43 || *(pos+i)==45){
			flag =  *(pos+i);                     //是符号的话就定符号 ,并且后面紧挨着的第一个数字也存入num 并把i+1 
			i++;
			num = charToNum(*(pos+i));
			
		}else{
			num = charToNum(*(pos+i));            
		} 
		    
		i++; 
		while(!isSpace(*(pos+i))&& *(pos+i) !=44){   //后面不是空格也不是逗号, 那么说明是数字或者发生错误 
		 	if(isIntger(*(pos+i))){
		 		
		 		num = num*10 + charToNum(*(pos+i));
        		i++;		
			}else if(*(pos+i) ==93){             //右括号的话可以收起来 ，并返回值 
			    if(flag ==45){
			    	num = -(num);
				}
				(arr->iptr)[k] = num;
				arr->len = k+1;
				*p = i+1;
				return 0;
			}else{
				*p = i+1;               //不是数字不是逗号就出错了  p指向出错后位置 
        		arr->len = 0;
        		return ERR_INVAL;
			} 
	} 
		while(isSpace(*(pos+i))){
			i++;
		}
		if(*(pos+i) ==44){          //是逗号的话收起来 
			
			if(flag ==45){
			    num = -(num);
			}
			(arr->iptr)[k] = num; 
			k++;
			i++;
		}
		flag = '+';		          //重新把flag置为正号 
			             
    }
}



