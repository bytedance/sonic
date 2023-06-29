#include <stdio.h>
#include <malloc.h>
#include <stdbool.h>
#include <stdint.h>
#include <stddef.h>
#include <string.h>

#include "native.h"
#include "types.h"

#define ERR_INVAL       2
#define ERR_RECURSE_MAX 7

typedef struct {
    union {                    // the pointer of u64 or i64 array 
        int64_t*  iptr;
        uint64_t* uptr; 
    };
    size_t len;                // the length of slice
    size_t cap;                // the capacity of slice 
} GoIntSlice;

bool is_space(char a){
    if( a == ' '||a == '\t'||a == '\r'||a == '\n'){
        return true;
    }else{
        return false;
    }
}

bool is_integer(char a){
    return a>='0'&& a<='9';
}

int char_to_num(char c){
    return c-'0';
}

long decode_u64_array( const GoString* src, long* p, GoIntSlice* arr){  
    const char* pos = src->buf;
    int i =*p;
    int len = src->len;
    //check space
    while(i<len && is_space(pos[i])){                                  
        i++;                           
    }
    //check left bracket and eof
    if(i >= len || pos[i] != '['){                                   
        *p = i;        
        arr->len = 0;
        return ERR_INVAL;
    }
    //when program runs here, it's a left parenthesis
    i++;     
    //K+1 represents the number of digits in the string src
    int k =0;
    //Num is used to store the current number                                                
    int num = 0;                                              
    while(i < len){                   
        //Jump back if it's a space
        while(i < len && is_space(pos[i])){     
            i++;             
	}
	//eof or the first one is not a number, it must be illegal
	if(i >= len || pos[i]<'0' || pos[i]>'9'){
	    //"[   ]" is true example
	    if(pos[i]==']' && k==0){
	        arr->len = k;
	        *p = i + 1;
	        return 0;
	    }
	    *p = i;                                     
            arr->len = 0;
	    return ERR_INVAL;                             
	}else{
	    num = char_to_num(pos[i]); 
	} 	     
	i++;
	//"0123"is false
	if(num == 0 && is_integer(pos[i])){
	    *p = i;                                     
            arr->len = 0;
	    return ERR_INVAL;
	}
	//parse the digital
	while(i < len && is_integer(pos[i])){		 		
	    num = num*10 + char_to_num(pos[i]);
            i++;		
	}
	while(i < len && pos[i] != ',' && pos[i] != ']'){     
            if(isspace(pos[i])){
                i++;
	    }else{
	        *p = i;                                     
                arr->len = 0;
                return ERR_INVAL;
	    }      
	}	
	//If the capacity is insufficient, return ERR_ RECURSE_ MAX
	if(k >= arr->cap){
            *p = i;
            return ERR_RECURSE_MAX;
        }
        (arr->uptr)[k++] = num;
	if (pos[i] == ']') {
            arr->len = k;
	    *p = i + 1;
	    return 0;
        }else{
            i++;
        }						 			             
    }    
    *p = i;
    arr->len = 0;
    return ERR_INVAL;   
}

long decode_i64_array(const GoString* src, long* p, GoIntSlice* arr){   
    const char* pos = src->buf;
    int i =0;
    int len = src->len;   
    while(i < len && is_space(pos[i])){                                             
        i++;                           
    }
    if(i >= len || pos[i] != '['){                                                  
        *p = i;                                                       
        arr->len = 0;
        return ERR_INVAL;
    }
    i++;                                                                
    int k =0;                                                          
    int num = 0;
    //Define a flag to represent the symbol of a signed number                                                        
    char flag ='+';       
    while(i < len){
        while(i < len && is_space(pos[i])){                                         
            i++;             
	}
	if(i >= len || (pos[i]<'0' || pos[i]>'9') && (pos[i]!='-' )){
	    if(pos[i]==']' && k==0){
	        arr->len = k;
	        *p = i + 1;
	        return 0;
	    }		
	    *p = i;                                                   
            arr->len = 0;
	    //The first one is neither a number nor a Plus or minus sign, so it must be illegal 
	    return ERR_INVAL;                        
	}else if(pos[i]=='+' || pos[i]=='-'){		
	    //Determine symbols , and the first digit immediately after it is also stored in num and i+1 
	    flag =  pos[i];                           
	    i++;
	    num = char_to_num(pos[i]);			
	}else{
	    num = char_to_num(pos[i]);            
	} 	    
	i++;	
	if(num == 0 && is_integer(pos[i])){
            *p = i;                                     
            arr->len = 0;
	    return ERR_INVAL;
	}
	while(i < len && is_integer(pos[i])){		 		
	    num = num*10 + char_to_num(pos[i]);
            i++;		
	}
	while(i < len && pos[i] != ',' && pos[i] != ']'){     
            if(isspace(pos[i])){
                i++;
	    }else{
		*p = i;                                     
                arr->len = 0;
                return ERR_INVAL;
	    }      
	}	
	//If the capacity is insufficient, return ERR_ RECURSE_ MAX
	if(k >= arr->cap){
            *p = i;
            return ERR_RECURSE_MAX;
        }
        if(flag =='-'){
	    num = -(num);
	}
        (arr->iptr)[k++] = num;
	if (pos[i] == ']') {
            arr->len = k;
	    *p = i + 1;
	    return 0;
        }else{
            i++;
        }
	flag = '+';	      //Reset flag to positive sign	 			             
    }
    *p = i;
    arr->len = 0;
    return ERR_INVAL;
}
