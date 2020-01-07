#include <iostream>

using namespace std;

int main(){
    int id[sizeof(unsigned long)];
    for(int i = 0; i<sizeof(unsigned long); i++){
        id[i] = i*i;
        printf("id[%d]=%d\n", i, id[i]);
    }
}