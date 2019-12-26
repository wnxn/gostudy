#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int dominantIndex(vector<int>& nums) {
        int f1=0,f2=0;
        int i1=0,i2=0;
        for(int i=0;i<nums.size();i++){
            if (nums[i]>f1){
                f2 = f1;
                i2=i1;
                f1 = nums[i];
                i1 = i;
            }else if (nums[i] > f2){
                f2 = nums[i];
                i2 = i;
            }
        }
        if( f1 >= f2*2){
            return i1;
        }else{
            return -1;
        }
    }
};

int main(){
    Solution s;
    int a[4]={0,0,0,3};
    vector<int> v(a, a+4);
    cout << s.dominantIndex(v) << endl;
}