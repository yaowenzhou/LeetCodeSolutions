/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 08:41:19
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 10:02:09
 * @version      :
 * @Description  :
 */

#include <vector>
#include <iostream>
#include "solution1.hpp"
#include "solution2.hpp"

using namespace std;

int main() {
    vector<int> src{ 3,2,4 };
    int target = 6;

    cout << "---->Solution1" << endl;
    Solution1 sol1 = Solution1();
    vector<int> ret1 = sol1.twoSum(src, target);
    if (ret1.size() != 0) {
        cout << ret1[0] << "--" << ret1[1] << endl;
    }

    cout << "---->Solution2" << endl;
    Solution2 sol2 = Solution2();
    vector<int> ret2 = sol2.twoSum(src, target);
    if (ret2.size() != 0) {
        cout << ret2[0] << "--" << ret2[1];
    }
    return 0;
}