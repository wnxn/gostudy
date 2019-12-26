#include <iostream>

using namespace std;

class people{
public:
    int age;
    int getAge(){
        return age;
    };
    people(int a): age(a){};
};

class student: public people{
public:
    int id;
    int getId(){
        return id;
    };
    student(int i, int a):id(i), people(a){};
};

struct man {
    int years;
    int getYears(){
        return years;
    }
};

struct teacher:people{
    teacher(int a): people(a){};
};

class staff:public man{
    
};

int main(){
    student s(2,3);
    cout << s.getAge() << " " << s.getId()<< endl;
    teacher t(2);
    cout << t.getAge() << endl;
    staff st;
    cout << st.getYears() << endl;
}