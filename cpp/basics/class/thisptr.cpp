#include <iostream>

using namespace std;

class A{
public:
    A(){p=this;};
    ~A(){
        if(p!= NULL){
            // delete p;
            p=NULL;
        }
    }
    A* p;
};

int main(){
    A a;
}