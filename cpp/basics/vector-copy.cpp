#include <vector>
#include <iostream>

using namespace std;

class Student{
public:
    Student(){
        cout << "Student" << endl;
    };
    Student(const Student& s){
        cout << "copy studnet" << endl;
    };
    ~Student(){
        cout << "~Student" << endl;
    }
};

int main(){
    vector<Student> s;
    Student s1;
    cout << "start" << endl;
    // copy one time
    s.push_back(s1);
}

// Result
// Student
// start
// copy studnet
// ~Student
// ~Student