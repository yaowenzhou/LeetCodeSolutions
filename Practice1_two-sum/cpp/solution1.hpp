/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 08:26:19
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 09:48:48
 * @version      :
 * @Description  : 解法1: 双层循环暴力破解
 */
#ifndef SOLUTION1
#define SOLUTION1

#include <iostream>
#include <vector>

using namespace std;

// 暴力破解，双层循环
class Solution1 {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        for (int i = 0; i < nums.size(); i++) {
            for (int j = i + 1; j < nums.size(); j++) {
                if (nums[i] + nums[j] == target) return vector<int>{i, j};
            }
        }
        return vector<int>();
    }
};
#endif
