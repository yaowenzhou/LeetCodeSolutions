/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 08:40:54
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 11:04:41
 * @version      :
 * @Description  : 方案2，使用 unordered_map
 */
#ifndef SOLUTION2
#define SOLUTION2

#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// 暴力破解，双层循环
class Solution2 {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> umap;
        for (int i = 0; i < nums.size(); i++) {
            auto iter = umap.find(target - nums[i]);
            if (iter != umap.end()) {
                return vector<int>{iter->second, i};
            }

            umap.emplace(nums[i], i); // 向umap中插入数据，键值对为 num[i] 以及其索引
        }

        return vector<int>();
    }
};
#endif
