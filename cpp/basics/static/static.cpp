#include "static.h"

int a = 1;
static int b = 2;

void func(void){
    static int i = 5;
    i++;
    std::cout << "i is" << i << std::endl;
}

static int staticAdd(int a, int b){
    cout << "staticAdd" << endl;
    return a+b;
}

extern int normalAdd(int a, int b){
    cout << "normalAdd" << endl;
    return staticAdd(a,b);
}

void staticPlus(){
    static int i =4;
    i++;
    cout << "staticPlus " <<  i << endl;
}