#include <iostream>

using namespace std;

class student{
public:
    student(int p){ };
};

class teacher{
public:
    explicit teacher(int p){};
};

int main(){
    student s=3;
    teacher t=3;
}