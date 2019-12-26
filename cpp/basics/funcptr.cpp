#include <iostream>
using namespace std;

typedef int (*p)(int,int);

int add(int a, int b){
    return a+b;
}

int main(){
    p f = add;
    cout << f(2,3) << endl;
}