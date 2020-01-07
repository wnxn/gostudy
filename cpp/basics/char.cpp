#include <iostream>
using namespace std;

int main(){
    char ch[5]={"1234"};
    cout << ch << endl;
    char ch2[4] = {'a','b','c','d'};
    cout << ch2 << endl;
    cout << &ch2 << " " << &(ch2[3]) << endl;
    cout << &ch << " " << &(ch[4]) << endl;
}