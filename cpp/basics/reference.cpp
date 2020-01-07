#include <iostream>

using namespace std;

int main(){
    int* i =new int(17);
    int r;
    //int& r = *i;
    printf("i=%d, r=%d\n", *i, r);
    delete i;
    r=19;
    printf("i=%d, r=%d\n", *i, r);
    int j = 20;
    r=j;
    printf("i=%d, j=%d, r=%d\n", *i, j, r);
}