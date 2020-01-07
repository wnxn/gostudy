#include <iostream>

using namespace std;

static int si = 9;

int gi = 10;

void getAddr(int i){
    static int lsi = 14;
    int lvi = i;
    int* di = new int(i+9);
    int* di2 = new int(i+1);
    int* di3 = new int(i+2);
    printf("the addr of global static %0x\n", &si);
    printf("the addr of global variable %0x\n", &gi);
    printf("the addr of local static %0x\n", &lsi);
    printf("the addr of local variable %0x\n", &lvi);
    printf("the addr of formal variable %0x\n", &i);
    printf("the addr of dynamical variable %0x, %0x, %0x\n", &di, di, *(di));
    printf("the addr of dynamical variable %0x, %0x, %0x\n", &di2, di2, *di2);
    printf("the addr of dynamical variable %0x, %0x, %0x\n", &di3, di3, *di3);
    delete di;
    delete di2;
    delete di3;
}

int main(){
    getAddr(10);
}

// the addr of global static 60e9030
// the addr of global variable 60e9028
// the addr of local static 60e902c
// the addr of local variable e9b176dc
// the addr of formal variable e9b176d8
// the addr of dynamical variable e9b176c0, 7fc00350, 13
// the addr of dynamical variable e9b176c8, 7fc028c0, b
// the addr of dynamical variable e9b176d0, 7fc028d0, c
// 栈空间
// 静态、全局变量区
// 堆空间