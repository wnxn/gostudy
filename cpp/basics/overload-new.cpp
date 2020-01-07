#include <iostream>

class student{
public:
    void* operator new(size_t t);
};

void* student::operator new(size_t t){
    printf("new student %u\n", t);
    void *temp = ::operator new(t);
    if(temp != NULL){
        printf("succeed new\n");
    }
    return temp;
}

int main(){
    student* s = new student();
    printf("%d\n", sizeof(s)/sizeof(*s));
    delete[] s;
}