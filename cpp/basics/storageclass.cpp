#include <iostream>
using namespace std;

// main() is where program execution begins.
int main() {
    register int a=7;
    cout << a << endl;
    int& b = a;
    cout << b << endl;

   return 0;
}