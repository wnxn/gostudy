#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        if(nums.size() == 0){
            return -1;
        }
        int total = 0;
        for (int item:nums){
            total += item;
        }
        int left = 0;
        int right = total;
        for(int i=0;i<nums.size();i++){
            if(i==0){
                left = 0;
            }else{
                left += nums[i-1];
            }
            right = total - left - nums[i];
            if (left == right){
                return i;
            }
        }
        return -1;
    }
};

int main(){
    Solution s;
    int a[6]={-1,-1,-1,-1,-1,0};
    vector<int> v(a, a+6);
    cout << s.pivotIndex(v) << endl;
}