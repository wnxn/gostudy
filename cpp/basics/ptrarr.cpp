#include <stdio.h>

using namespace std;

class people{
public:
    people(){};
    people(const char* s):name(s){
        printf("create people with name %s\n", s);
    }
    virtual ~people(){
        printf("delete people %s\n", name);
    }
    const char* name;
};

class student: public people{
public:
    student(){};
    student(const char* s):people(s){
        printf("create student %s\n", s);
    }
    ~student(){
        printf("delete student %s\n", name);
    }
};

int main(){
 //  student* s = new student[3]{student("wang"), student("song"), student("zhu")};
    student* s=new student[3];
    s[0] = student("wang");
    s[1] = student("song");
    s[2] = student("zhu");
    printf("%d, %d\n", sizeof(s), sizeof(*s));
    printf("%s\n", (s+2)->name);
    delete[] s;
 }