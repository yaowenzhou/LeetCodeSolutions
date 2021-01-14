/**
 * @Package      : []
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 10:25:59 Tuesday
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 10:38:15
 * @Version      : 
 * @Description  : 方案1，暴力破解
 */

public class Solution1 {
    public int[] twoSum(int[] nums, int target) throws Exception {
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[i] + nums[j] == target) {
                    return new int[] { i, j };
                }
            }
        }
        throw new IllegalArgumentException("No two sum solution");
    }
}
