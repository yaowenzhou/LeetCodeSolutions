/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 15:48:47
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-13 22:37:32
 * @version      :
 * @Description  :
 */
#ifndef __SOLUTION1__
#define __SOLUTION1__
#include <vector>
#include <iostream>

using namespace std;

class Solution1 {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int len1 = nums1.size();
        int len2 = nums2.size();

        // 题目要求两个数组不会同时为空
        // 首先考虑一下其中一个数组为空的情况，此时直接返回长数组的中位数即可
        if (len1 == 0) {
            return (nums2[len2 / 2] + nums2[(len2 - 1) / 2]) / 2.0;
        }
        if (len2 == 0) {
            return (nums1[len1 / 2] + nums1[(len1 - 1) / 2]) / 2.0;
        }

        // totalSize 表示两个数组的元素总数
        int totalSize = len1 + len2;
        // 由于totalSize为奇数和totalSize为偶数的情况不同
        // leftSize 表示左边元素的个数
        // totalSize 为奇数，比如3，(totalSize + 1) / 2 == 2，leftSize == rightSize + 1
        // totalSize 为偶数，比如4，(totalSize + 1) / 2 == 2，leftSize == rightSize
        // 这样就可以暂时统一奇偶数的情况，并将之放到最后讨论
        int leftSize = (totalSize + 1) / 2;
        // 分别定义两个引用用于表示长数组和短数组
        const vector<int>& shortNums = len1 <= len2 ? nums1 : nums2;
        const vector<int>& longNums = len1 > len2 ? nums1 : nums2;

        // 定义一个变量，用于表示长数组中左半部分最大值的索引
        // 并且假定短数组全部处于左半部分
        int k = leftSize - shortNums.size() - 1;
        // 定义一个变量用于表示短数组需要调整的元素个数，正数表示将元素划分到右边，负数表示将元素划分到左边
        // 如果需要调整元素划分，每次调整shortNumsMoveSize / 2个元素
        // 由于初始状态下，短数组的元素全部在左边，因此第一次调整时，需要调整的元素个数为shortNums.size() / 2;
        int shortNumsMoveSize = shortNums.size() / 2;
        if (shortNumsMoveSize == 0) {
            shortNumsMoveSize = 1;
        }
        // 定义两个变量分别表示左边的最大值和右边的最小值
        // 由于假定的短数组全部处于左半部分
        // 因此默认左半部分的最大值取 --shortNums.end() 和 longNums[k] 之间较大者
        // 如果两个数组一样长，则shortNums[shortNums.length - 1]表示左边的最大值
        int leftMax = shortNums[shortNums.size() - 1];
        if(len1 !=len2) {
            leftMax = *(--shortNums.end()) >= longNums[k] ? *(--shortNums.end()) : longNums[k];
        }
        // 默认右半部分的最小值取 longNums[k+1]
        int rightMin = longNums[k + 1];

        // 所有元素划分成两部分的要求是leftMax <= rightMin，此即为循环的终止条件
        while (leftMax > rightMin) {
            // 进入循环后开始调整元素划分
            // 短数组需要向右边多划分shortNumsMoveSize，
            // 为了保证左右两边元数数量的平衡，长数组需要相应的向左边多划分shortNumsMoveSize个元素
            // k需要相应的增加 shortNumsMoveSize
            k = k + shortNumsMoveSize;
            if (k > leftSize - 1) {
                k = leftSize - 1;
            }
            
            shortNumsMoveSize = shortNumsMoveSize > 0 ? shortNumsMoveSize : -shortNumsMoveSize;
            // 调整好 k 的位置后，即完成了元素划分
            // 做出相应的比较并且重新设置相应的值
            /* 此处说明一下
             * 默认短数组全部划分到左边，因此一旦进入循环，说明短数组一定有部分要划分到右边，而长数组必定是左右两边都有
             * shortNumsMoveSize >= shortNumsMoveSize / 2;
             * 第一次调整时，调整的元素数量是 shortNums.size() / 2;
             * 以后每次调整时都调整shortNumsMoveSize / 2;个元素
             * 极端情况下可能会出现整个短数组都会调整到右边
             * 此时leftMax = longNums[k]
             * 在两个数组中
             * shortNums[leftSize - (k + 1) - 1](shortLeft) <= shortNums[leftSize - (k + 1)](shortRight)
             * longNums[k](longLeft) <= longNums[k + 1](longRight)
             * 是否需要调整由 leftMax和rightMin确定，当不需要调整时，shortNumsMoveSize的值是多少无所谓
             * 需要调整时，分两种情况，shortNums向右调整和shortNums向左调整
             * leftMax > rightMin
             * 此时如果 leftMax == shortLeft，则将短数组向右调整
             * 此时有: shortLeft >
             * 如果 leftMax == longLeft，则说明将长数组向右调整，相对应的，短数组则需要向左调整
            */
            // 循环开始时，会判断 k + 1 == leftSize
            // 如果条件成立，说明整个短数组都被划分到了右边
            // 因此在此之前，shortNums[leftSize - (k + 1) - 1] 不会出现数组越界的情况
            // 如果整个短数组都被移到了最右边，则leftMax == longNums[k];
            // 如果leftMax == longNums[k];
            // 则 长数组应该向右调整，短数组应该向左调整，shortNumsMoveSize应该为负
            if ((k + 1) >= leftSize || longNums[k] > shortNums[leftSize - (k + 1) - 1]) {
                leftMax = longNums[k];
                // 为了保证恰好shortNumsMoveSize / 2 == 0时依然可以调整元素
                shortNumsMoveSize = (shortNumsMoveSize / 2) == 0 ? -1 : (-shortNumsMoveSize / 2);
            }
            else {
                leftMax = shortNums[leftSize - (k + 1) - 1]; // shortLeft
                // 为了保证恰好shortNumsMoveSize / 2 == 0时依然可以调整元素
                shortNumsMoveSize = (shortNumsMoveSize / 2) == 0 ? 1 : (shortNumsMoveSize / 2);
            }
            // 最最极端的情况: 长数组和短数组一样长，且长数组的最大值小于短数组的最小值
            if ((k + 1) == longNums.size()) {
                rightMin = shortNums[0];
            }
            else {
                rightMin = shortNums[leftSize - (k + 1)] < longNums[k + 1] ? shortNums[leftSize - (k + 1)] : longNums[k + 1];
            }
        }
        // 退出循环后，返回中位数
        // 如果 totalSize 为奇数，则返回 leftMax
        if (totalSize % 2 == 1) {
            return leftMax * 1.0;
        }
        // 如果 totalSize 为偶数，则返回 (leftMax + rightMin) / 2
        return (leftMax + rightMin) / 2.0;
    }

};
#endif
