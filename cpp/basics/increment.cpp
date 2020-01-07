#include <iostream>

using namespace std;

int main(){
    int a = 4;
    a+=(a++);
    assert(a==9);

    a+=(++a);
    assert(a==20);

    (++a)+=a;
    cout << a <<  endl;
}