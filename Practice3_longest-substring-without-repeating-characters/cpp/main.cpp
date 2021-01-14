/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 13:09:31
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 14:20:55
 * @version      :
 * @Description  :
 */
#include <string>
#include <iostream>

#include "Solution1.hpp"
#include "Solution2.hpp"

using namespace std;

int main() {
    string str = "abcabcbb";

    Solution1 sol1 = Solution1();
    cout << sol1.lengthOfLongestSubstring(str) << endl;

    Solution2 sol2 = Solution2();
    cout << sol2.lengthOfLongestSubstring(str) << endl;
    return 0;
}