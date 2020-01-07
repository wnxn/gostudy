#include <iostream>
#include <assert.h>
using namespace std;

int add(int n){
    static int i = 100;
    i+=n;
    return i;
}

int& addRef(int n){
    static int i = 100;
    i+=n;
    return i;
}

int* addPtr(int n){
    static int i = 100;
    i+=n;
    return &i;
}

int main(){
    int a1 = 0;
    a1 = add(2);
    assert(a1==102);
    a1 = add(3);
    assert(a1==105);

    int a2 = 0;
    a2 = addRef(4);
    assert(a2==104);
    a2 = 23;
    assert(a2==23);
    a2 = addRef(5);
    assert(a2==109);

    int* a3;
    a3 = addPtr(6);
    assert(*a3==106);
    *a3 = 23;
    assert(*a3==23);
    a3 = addPtr(7);
    assert(*a3==30);
}