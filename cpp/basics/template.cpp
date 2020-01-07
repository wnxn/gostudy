#include <iostream>

using namespace std;

int add(int a, int b){
    cout << "add int" << endl;
    return a+b;
}

template <typename T> 
T add(T a, T b){
    cout << "add TA" << endl;
    return a+b;
}



int main(){
    int a=1,b=2;
    add(a,b);
}