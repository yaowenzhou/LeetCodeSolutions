/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 15:43:00
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-13 22:37:07
 * @version      :
 * @Description  :
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cstdlib>
#include <ctime>
#include <iomanip>

#include "Solution1.hpp"

using namespace std;

int main()
{
    vector<int> nums1 = {};
    vector<int> nums2 = { 2, 3 };
    // vector<int> nums3;
    // srand((int)time(0));  // 产生随机种子  把0换成NULL也行
    // for (int i = 0; i < 20; i++)
    // {
    //     nums1.push_back(rand() % 90 + 10);
    // }

    // for (int i = 0; i < 5; i++)
    // {
    //     nums2.push_back(rand() % 90 + 10);
    // }

    // sort(nums1.begin(), nums1.end());
    // sort(nums2.begin(), nums2.end());

    // for_each(nums1.begin(), nums1.end(),
    //     [&nums3](int val)->void {
    //         cout << val << " ";
    //         nums3.push_back(val);
    //     });
    // cout << endl;
    // for_each(nums2.begin(), nums2.end(),
    //     [&nums3](int val)->void {
    //         cout << val << " ";
    //         nums3.push_back(val);
    //     });

    // cout << endl;
    // sort(nums3.begin(), nums3.end());
    // for_each(nums3.begin(), nums3.end(), [](int val)->void{
    //     cout << val << " ";
    // });
    // cout << endl;

    Solution1 sol1 = Solution1();
    double midNum = sol1.findMedianSortedArrays(nums1, nums2);
    cout << "---->" << midNum << endl;
    return 0;
}

/*
13 18 21 27 28 33 37 43 46 46 49 50 58 60 67 67 72 74 85 85
12 13 14 14 17 21 27 33 34 41 55 65 69 83 86 91 97 97
*/