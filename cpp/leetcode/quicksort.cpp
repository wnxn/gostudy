#include <iostream>
#include <vector>
using namespace std;

void help(vector<int> &arr, int low, int high);


void printVector(const vector<int>& vec){
    for(auto item: vec){
        cout << " " << item ;
    }
    cout  << endl;
}

vector<int> quickSort(vector<int> arr){
    if(arr.size() <= 1){
        return arr;
    }
    int low = 0;
    int high = arr.size()-1;
    cout << low << " " << high << endl;
    help(arr, low, high);
    return arr;
}

void help(vector<int> &arr, int low, int high){
    if(high <= low){
        return;
    }
    int l = low;
    int h = high;
    while(low < high){
            while(low < high && arr[low] <= arr[high]){
                high--;
            }
            swap(arr[low], arr[high]);
            while(low < high && arr[low] <= arr[high]){
                low++;
            }
            swap(arr[low], arr[high]);
    };
    help(arr, l, high-1);
    help(arr, high+1, h);
}

int main(){
    int arr[]={23,43,12,4,56,23,98};
    vector<int> vec(arr, arr+7);
    vector<int> v2=  quickSort(vec);
    for(auto item: vec){
        cout << " " << item ;
    }
    cout << endl;
    for(auto item: v2){
        cout << " " << item ;
    }
    cout << endl;
}