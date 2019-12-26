#include "box.h"

void Box::setLength(double v){
    length = v;
}

void Box::setBreadth(double v){
    breadth = v;
}

void Box::setHeigth(double v){
    height = v;
}

double Box::getVolume(){
    return length * breadth * height;
}