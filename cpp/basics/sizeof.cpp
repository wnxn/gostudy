#include <iostream>
using namespace std;

void getSize();

int main() {
   cout << "Size of char : " << sizeof(char) << endl;
   cout << "Size of short : " << sizeof(short) << endl;
    cout << "Size of short int : " << sizeof(short int) << endl;
   cout << "Size of int : " << sizeof(int) << endl;

   cout << "Size of long int : " << sizeof(long int) << endl;
   cout << "Size of float : " << sizeof(float) << endl;
   cout << "Size of double : " << sizeof(double) << endl;
   cout << "Size of wchar_t : " << sizeof(wchar_t) << endl;
   getSize();
   return 0;
}

void getSize(){
   cout << "getSize" << endl;
   char str[] = "world";
   cout << sizeof(str) << ' ';
   char *p= str; cout << sizeof(p) << ' ';
   char i = 10; cout << sizeof(i) << ' ';
   cout << str << " " << p << endl;
   void *pp = malloc(10); cout << sizeof(p) << endl;
}