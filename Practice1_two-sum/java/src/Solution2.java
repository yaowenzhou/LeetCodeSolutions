import java.util.HashMap;
import java.util.Map;

/**
 * @Package      : []
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 10:19:19 Tuesday
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 10:35:45
 * @Version      : 
 * @Description  : 使用HashMap
 */

public class Solution2 {
    public int[] twoSum(int[] nums, int target) throws Exception {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (map.containsKey(target - nums[i])) {
                return new int[] { map.get(target - nums[i]), i };
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("No two sum solution");
    }
}
