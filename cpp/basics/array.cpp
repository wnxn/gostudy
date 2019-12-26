#include <iostream>
#include <vector>
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
}