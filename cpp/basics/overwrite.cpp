#include <iostream>

using namespace std;

class people{
public:
    int getage(){
        return age;
    }
private:
    int age;
};


class student: public people{
public:
    // int getage(int a){
    //     return a;
    // }
};

int main(){
    student* s=new student();
    cout << s->getage() << endl;
}