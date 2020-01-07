#include <iostream>
#include <string>
using namespace std;

int* add(int* a, int * b){
    int c = *a+*b;
    return &c;
}

int main(){
    // string s1 = "hello";
    // string s2 = "world";
    // char c1[10];
    // char c2[10];
    // strcpy(c1, s1.c_str());
    // strcpy(c2, s2.c_str());
    // char* c3 = strcat(c1,c2);
    // cout << c1 << endl;
    // cout << c2 << endl;
    // cout << c3 << endl;
    cout << "hello" << endl;
    int a = 1, b= 2;
    int* c= add(&a,&b);;
    cout << *c << endl;
    
}