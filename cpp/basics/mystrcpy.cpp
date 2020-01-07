#include <iostream>

using namespace std;

char* mystrcpy(char * dst, const char* src){
    if(dst==NULL||src==NULL){
        return NULL;
    }
    int i =0;
    for(; src[i]!='\0'; i++){
        dst[i] = src[i];
    }
    dst[i] = '\0';
    return dst;
}