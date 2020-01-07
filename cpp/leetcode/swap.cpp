#include <iostream>
#include <vector>
using namespace std;

void swap(int& a, int& b){
    a ^=b;
    b ^= a;
    a ^= b;
}

int main(){
    int a = 234;
    int b = -23;
    swap(a,b);
    cout << a << " " << b << endl;
}