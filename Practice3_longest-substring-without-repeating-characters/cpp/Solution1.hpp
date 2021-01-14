/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 13:09:47
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 14:25:22
 * @version      :
 * @Description  : Solution1 暴力破解，枚举所有子串
 */
#ifndef __SOLUTION1__
#define __SOLUTION1__

#include <string>
#include <unordered_set>

using namespace std;
class Solution1 {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_set<char> set;
        int maxLen = 0; // 记录最长子串长度
        int n = s.size();
        int left = 0, right = 0;
        // 外层循环，子串起始字符
        for (left = 0; left < n; left++) {
            // 内层循环，子串终止位置
            for (right = left; right < n; right++) {
                // 如果找不到某个字符，说明该字符未出现，可以插入
                if (set.find(s[right]) == set.end()) {
                    set.insert(s[right]);
                }
                // 如果找到该字符，则说明有重复字符，清空set，本次内循环终止
                else {
                    set.clear();
                    break;
                }
            }
            maxLen = maxLen > (right - left) ? maxLen : (right - left);
        }
        return maxLen;
    }
};

#endif
