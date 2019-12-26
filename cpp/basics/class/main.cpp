#include <iostream>
#include "box.h"

using namespace std;

int main(){
    Box b;
    b.setBreadth(2);
    b.setHeigth(3);
    b.setLength(4);
    cout << b.getVolume() << endl;
}