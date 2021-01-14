/**
 * @Package      : []
 * @Author       : yaowenzhou
 * @Date         : 2021-01-13 20:59:57 Wednesday
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-13 22:32:57
 * @Version      : 
 * @Description  : 
 */

public class Solution1 {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int len1 = nums1.length;
        int len2 = nums2.length;
        // 当有一个数组为空数组时，直接可以求得中位数
        if (len1 == 0) {
            return (nums2[len2 / 2] + nums2[(len2 - 1) / 2]) / 2.0;
        }
        if (len2 == 0) {
            return (nums1[len1 / 2] + nums1[(len1 - 1) / 2]) / 2.0;
        }
        // 定义两个数组分别表示长数组和短数组
        int[] shortNums = null;
        int[] longNums = null;
        if (len1 <= len2) {
            shortNums = nums1;
            longNums = nums2;
        } else {
            shortNums = nums2;
            longNums = nums1;
        }
        // totalSize表示数组总元素个数
        int totalSize = nums1.length + nums2.length;
        // leftSize 表示左边元素的个数，+1 保证了totalSize为奇数时，中位数处于左边
        int leftSize = (totalSize + 1) / 2;
        // 中位数的作用是将数组划分为两部分
        // 此时必须满足下面两个要求
        // 1. leftPart.Max <= rightPart.min
        // 2. leftPart.length == rightPart.length || leftPart.length == rightPart.length + 1;
        // 我们将总数量为奇数时，中位数归位到左边
        // 因此我们可以这么操作:
        // 1. 将数组划分为两部分，默认将短数组全部划分到右边
        //      定义一个变量用于表示长数组中，最右边的元素的索引
        int k = leftSize - shortNums.length - 1;
        // 由于当前的划分保证了元素数量的要求，因此调整时需要重新划分元素
        // 通过改变k可以实现重新划分元素
        // 定义一个变量moveSize用于表示 k 变动的大小
        // moveSize为正时，表示k增加|moveSize|，
        // moveSize为负时，表示k减小|moveSize|
        // moveSize的初始值为 shortNums.length / 2
        // 以后每次调整时 moveSize = (|moveSize|) / 2
        int moveSize = shortNums.length / 2;
        if (moveSize == 0) {
            moveSize = 1;
        }
        // 定义两个变量分别用于表示左边最大值和右边最小值
        // 如果两个数组一样长，则shortNums[shortNums.length - 1]表示左边的最大值
        // moveSize决定了调整的大小
        // leftMax 的值决定了调整方向
        int leftMax = shortNums[shortNums.length - 1];
        if (len1 != len2) {
            if (shortNums[shortNums.length - 1] > longNums[k]) {
                leftMax = shortNums[shortNums.length - 1];
            } else {
                leftMax = longNums[k];
                // shortNums左调整，k需要减小，moveSize为负
                moveSize = -moveSize;
            }
        }
        int rightMin = longNums[k + 1];

        // 中位数的要求是leftPart.Max <= rightPart.min，因此不满足该条件时，我们需要进入循环进行调整中位数的位置
        while (leftMax > rightMin) {
            // 1. 改变k值
            // 如果k > leftSize - 1，表示整个短数组全部调整到了右边，考虑到实际情况，需要将k设置为leftSize - 1
            // 第一次循环后，k的值为 k + shortNums.length/2，并且以后每次调整时，都会将调整的大小减半
            // 每次每个数组最少调整一个元素，这个过程在k向左靠拢必定会执行到，中必定可以找到中位数
            // 因此不会造成longNums的访问越界，故而只需要考虑shortNums的访问越界
            k = k + moveSize > leftSize - 1 ? leftSize - 1 : k + moveSize;

            // 2. 修改moveSize，其正负将根据实际情况另行设置
            // |moveSize| 最小值为1，保证了一直可以调整元素
            if (moveSize / 2 == 0) {
                moveSize = 1;
            } else if (moveSize < 0) {
                moveSize = -moveSize / 2;
            } else {
                moveSize = moveSize / 2;
            }
            // longNums[k](longLeft) 和 shortNums[leftSize - (k + 1) - 1](shortLeft)表示两个数组左边的最大值
            // 二者的最大值即为leftMax
            if (k + 1 == leftSize/*整个短数组全部调整到了右边*/ || longNums[k] > shortNums[leftSize - (k + 1) - 1]) {
                leftMax = longNums[k];
                // leftMax == longNums[k] 说明下次调整时，k需要减小
                moveSize = -moveSize;
            } else {
                leftMax = shortNums[leftSize - (k + 1) - 1];
            }
            // longNums[k + 1](longRight) 和 shortNums[leftSize - (k + 1)](shortRight)表示两个数组右边的最小值
            // 二者的最小值即为rightMin
            // 最最极端的情况，长数组和短数组长度一样，并且长数组的最大值小于短数组的最小值，即整个长数组都调整到了左边
            if (k + 1 == longNums.length) {
                rightMin = shortNums[0];
            } else {
                rightMin = longNums[k + 1] < shortNums[leftSize - (k + 1)] ? longNums[k + 1]
                        : shortNums[leftSize - (k + 1)];
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
}
