#include <iostream>

using namespace std;

enum myenum1 {
    white,
    yellow,
    orange = 5,
    red,
    black,
};

int main(){
    myenum1 t = red;
    printf("%d\n", t);
}