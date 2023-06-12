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
    union { // u64 �� i64 array ��ָ��
        int64_t*  iptr;
        uint64_t* uptr; 
    };
    size_t len; // s ����
    size_t cap; // slice ����
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
// ���������src �� �����json, p �ǵ�ǰλ��.
// ����ֵ��
// ����ǺϷ���json ���飬����0 ��ʾ�����ɹ���
// ���������Ԫ���ǷǷ�json������-ERR_INVAL
// ���������Ԫ����json ���Ͳ�ƥ�䣬���� json �ǺϷ��ģ����� -ERR_MISMATCH
// ��� slice cap ���������� -ERR_RECURSE_MAX��
// �����������ŵ�arr�У�len +1��ֱ��cap��

// ����1
// ����:
// src�� һ���ַ�������json�������飬����"[1,2,3,4]" 
// p: ��ǰ��λ�ã�������0, 
// arr: ������len=0, cap=256
// �����
// p: ������ָ���λ�ã�����Ӧ����json����λ�ã���9
// arr: uptrָ�����δ���1,2,3��4. Ȼ��len = 4, cap���䡣
// ����ֵ��0

// ����2
// ���룺
// src�� һ���ַ���������"{[]}"���� "[1,2,3,4.5], ���jsonƬ���ǷǷ��Ļ��߲�������
// p: ��ǰ��λ�ã�������0, 
// arr: ������len=0, cap=256
// �����
// p: ������ָ���λ�ã�����Ӧ����json����λ��
// arr: ���������ǷǷ���json����Ҫ����-ERR_INVALID��ͬʱarr�г�����Ҫ����Ϊ0
// ����ֵ�� ERR_INVAL

long decode_u64_array( const GoString* src, long* p, GoIntSlice* arr){ //�޷�������
    char* pos = src->buf;
    int i =0;
    
    while(isSpace(*(pos+i))){      //����ʼǰ���пո��ȳԵ��ո� 
    	i++;                           
	}
    if(*(pos+i) != 91){  //��һ�������������ŵĻ�ֱ�ӷ��طǷ�
        *p = i+1;    //pָ���һ�������λ�� 
        arr->len = 0;
        return ERR_INVAL;
    }
    i++;  //��������
    int k =0;   //��Ҫһ���α�k,m
    int num = 0; //�洢���� 
    while(*(pos+i) !=0){
        if(k==arr->cap){   //��������
        	*p = i+1;
            return ERR_RECURSE_MAX;
        }
        while(isSpace(*(pos+i))){   //�ǿո�������� 
        	i++;             
		}
		if(*(pos+i)<48 || *(pos+i)>57){
			*p = i+1;    //pָ���һ�������λ�� 
        	arr->len = 0;
			return ERR_INVAL;                    //��һ���������ֿ϶��Ƿ� 
		}else{
			num = charToNum(*(pos+i));
		} 
		    
		i++; 
		while(!isSpace(*(pos+i))&& *(pos+i) !=44){   //���治�ǿո�Ҳ���Ƕ���, ��ô˵�������ֻ��߷������� 
		 	if(isIntger(*(pos+i))){
		 		
		 		num = num*10 + charToNum(*(pos+i));
        		i++;		
			}else if(*(pos+i) ==93){             //�����ŵĻ����������� ��������ֵ 
				(arr->uptr)[k] = num;
				arr->len = k+1;
				*p = i+1;
				return 0;
			}else{
				*p = i+1;               //�������ֲ��Ƕ��žͳ�����  pָ������λ�� 
        		arr->len = 0;
        		return ERR_INVAL;
			} 
	}
		while(isSpace(*(pos+i))){
			i++;
		}
		if(*(pos+i) ==44){          //�Ƕ��ŵĻ������� 
			(arr->uptr)[k] = num; 
			k++;
			i++;
		}
			 
			             
    }
    

  

}



long decode_i64_array(const GoString* src, long* p, GoIntSlice* arr){//�з�������
	char* pos = src->buf;
	int i =0;
	    
	while(isSpace(*(pos+i))){      //����ʼǰ���пո��ȳԵ��ո� 
	    i++;                           
	}
    if(*(pos+i) != 91){  //��һ�������������ŵĻ�ֱ�ӷ��طǷ�
        *p = i+1;    //pָ���һ�������λ�� 
        arr->len = 0;
        return ERR_INVAL;
    }
    i++;  //��������
    int k =0;     //��Ҫһ���α�k,m
    int num = 0;    //�洢���� 
    char flag ='+';   //��flag��ʼ��Ϊ���� 
    while(*(pos+i) !=0){
        if(k==arr->cap){   //��������
        	*p = i+1;
            return ERR_RECURSE_MAX;
        }
        while(isSpace(*(pos+i))){   //�ǿո�������� 
        	i++;             
		}
		if((*(pos+i)<48 || *(pos+i)>57) && (*(pos+i)!=43 && *(pos+i)!=45 )){
			
			*p = i+1;    //pָ���һ�������λ�� 
        	arr->len = 0;
			return ERR_INVAL;                    //��һ���Ȳ������֣��ֲ��������ţ���ô�϶��Ƿ� 
		}else if(*(pos+i)==43 || *(pos+i)==45){
			flag =  *(pos+i);                     //�Ƿ��ŵĻ��Ͷ����� ,���Һ�������ŵĵ�һ������Ҳ����num ����i+1 
			i++;
			num = charToNum(*(pos+i));
			
		}else{
			num = charToNum(*(pos+i));            
		} 
		    
		i++; 
		while(!isSpace(*(pos+i))&& *(pos+i) !=44){   //���治�ǿո�Ҳ���Ƕ���, ��ô˵�������ֻ��߷������� 
		 	if(isIntger(*(pos+i))){
		 		
		 		num = num*10 + charToNum(*(pos+i));
        		i++;		
			}else if(*(pos+i) ==93){             //�����ŵĻ����������� ��������ֵ 
			    if(flag ==45){
			    	num = -(num);
				}
				(arr->iptr)[k] = num;
				arr->len = k+1;
				*p = i+1;
				return 0;
			}else{
				*p = i+1;               //�������ֲ��Ƕ��žͳ�����  pָ������λ�� 
        		arr->len = 0;
        		return ERR_INVAL;
			} 
	} 
		while(isSpace(*(pos+i))){
			i++;
		}
		if(*(pos+i) ==44){          //�Ƕ��ŵĻ������� 
			
			if(flag ==45){
			    num = -(num);
			}
			(arr->iptr)[k] = num; 
			k++;
			i++;
		}
		flag = '+';		          //���°�flag��Ϊ���� 
			             
    }
}



