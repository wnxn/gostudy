#include <iostream>

using namespace std;

class people{
public:
    virtual void print(){
        cout << "people" << endl;
    };
    virtual ~people(){
        cout << "delete people" << endl;
    }
    virtual int getAge(){
        return age;
    }
    int age;
};

class student: public people{
public:
    void print(){
        cout << "student" << endl;
    }
     ~student(){
        cout << "delete student" << endl;
    }
    virtual int getAge(){
        return 2;
    }
    student(){};
};

class kinder: public student{
public:
    void print(){
        cout << "kinder" << endl;
    }
    ~kinder(){
        cout << "delete kinder" << endl;
    }
};

void getVirtual(){
    people* p = new student();
    cout << p->getAge() << endl;
}

int main(){
//     people* p = new kinder();
//     p->print();
//     people& r = *p;
//     r.print();
//   //  delete p;
//     r.print();
//     people* p2 = p;
//     p2->print();
    getVirtual();
}

