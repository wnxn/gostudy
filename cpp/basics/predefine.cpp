#include <iostream>

using namespace std;

class A{};

int main(){
    cout << __FILE__ << endl;
    cout << __LINE__ << endl;
    cout << sizeof(A) << endl;
}