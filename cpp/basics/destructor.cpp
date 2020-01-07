#include <iostream>
using namespace std;
class CDemo {
public:
    virtual ~CDemo() { cout << "CDemo destructor" << endl; }
};

class DDemo: public CDemo{
public:
     ~DDemo() override{cout << "DDemo destructor" << endl;}
};

void Func(CDemo obj) {
    cout << "func" << endl;
}
//CDemo d1;
// CDemo Test() {
//     cout << "test" << endl;
//     return d1;
// }

void TestInherOverwriteDes(){
    cout << "before ddemo" << endl;
    CDemo* d= new DDemo();
    delete d;
    cout << "after ddemo" << endl;
}

int main() {
    // CDemo d2;
    // Func(d2);
    // Test();
    // cout << "after test" << endl;
    TestInherOverwriteDes();
    return 0;
}