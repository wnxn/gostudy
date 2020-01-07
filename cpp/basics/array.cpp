#include <iostream>
#include <vector>
#include <assert.h>
#include <iomanip>
using namespace std;

int main(){
    int a[5] = {1,2,3};
    for (int i = sizeof(a)/sizeof(*a)-1; i>=0;i--){
        cout << " " << a[i];
    }
    cout << endl;

    vector<int> v(a, a+5);
    for(vector<int>::iterator item = v.begin(); item != v.end(); item++){
        cout << " " << *item ;
    }
    cout << endl;
    for (int item:v){
        cout << " " << item++;
    }
    cout << endl;
        for (int item:v){
        cout << " " << item;
    }
    cout << endl;

    double balance[]={100.0, 2.0, 3.4};
    assert(sizeof(balance)/sizeof(*balance)==3);
    cout << "Element" << std::setw(13) << "Value" << endl;

    cout << sizeof(balance) << endl;
    double *p = balance;
    cout << *p << endl;
    printf("%0xu, %0xu\n", balance+2, balance+1);
    printf("%f, %f, %d\n", *(balance+2), *(balance+1), sizeof(*balance));
    strcpy
}