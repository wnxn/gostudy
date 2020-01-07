#include <iostream>

using namespace std;

const char* copystr(string s){
    return s.c_str();
}

int main(){
    string s = "hello";
    cout << copystr(s) << endl;
}