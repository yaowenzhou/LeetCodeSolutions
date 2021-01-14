/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 13:43:32
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 15:15:47
 * @version      :
 * @Description  : 滑动窗口
 */
#ifndef __SOLUTION2__
#define __SOLUTION2__

#include <string>
#include <unordered_map>

using namespace std;
class Solution2 {
public:
    int lengthOfLongestSubstring(string s) {
        // 使用map存储该字符及其索引
        unordered_map<char, int> map;
        int maxLen = 0; // 记录最长子串长度
        int n = s.size();
        int start = 0, end = 0;
        // 外层循环，子串起始字符
        for (end = 0; end < n; end++) {
            // 查找该字符的索引
            auto iter = map.find(s[end]);
            // 如果该元素未在map中出现过或者其索引小于start，则说明，在当前的窗口中，并未出现该字符
            if (iter == map.end() || iter->second < start) {
                // 插入或者修改该字符对应的索引
                map[s[end]] = end;
            }
            // 该元素在map中出现过，并且其索引大于start
            else {
                // 1. 将当前的窗口长度和maxLen对比，取大者赋值给maxLen
                maxLen = maxLen > (end - start) ? maxLen : (end - start);
                // 2. 需要将窗口的起始位置修改为 该字符的索引 + 1
                start = iter->second + 1;
                // 更新字符s[i] 的索引值
                map[s[end]] = end;
            }
        }
        // 为了效率，在每次循环的时候，如果字符不重复，则不更新maxLen
        // 因此需要在最后出来的时候，再更新一次maxLen
        return maxLen > (end - start) ? maxLen : (end - start);
    };
};

#endif
