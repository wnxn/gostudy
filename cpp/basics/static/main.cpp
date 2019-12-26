#include <iostream>
#include "static.h"

using namespace std;

static int c = 10;
int student::a = 2;
int main(){
   cout << normalAdd(3,4) << endl;
   staticPlus();staticPlus();staticPlus();
   //student::a = 1;
   cout << student::a << endl;
}