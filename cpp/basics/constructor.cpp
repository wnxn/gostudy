#include <iostream>

using namespace std;

struct CLS{
    int m_i;
    CLS(int i):m_i(i){}
    CLS(){
        cout << "CLS()" << endl;
        m_i = 2;
    }
};

int main(){
    CLS obj;
    cout << obj.m_i << endl;
}