#include <iostream>

using namespace std;

class human{
public:
    virtual ~human(){
        cout << "human over ..." << endl;
    }
    void Disp(){
        cout << "human disp ..." << endl;
    }
    human(){
        cout << "human start ..." << endl;
    }

};

class man: public human{
public:
    ~man(){
        cout << "man over ..." << endl;
    }
    void Disp(){
        cout << "man disp ..." << endl;
    }
    man(){
        cout << "man start ..." << endl;
    }
};

int main(){
    human *p = new man();
    p->Disp();
    delete p;
}