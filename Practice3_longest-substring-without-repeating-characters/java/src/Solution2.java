import java.util.HashMap;
import java.util.Map;

/**
 * @Package      : []
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 14:51:14 Tuesday
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 15:18:31
 * @Version      : 
 * @Description  : 
 */

public class Solution2 {
    public int lengthOfLongestSubstring(String s) {
        int maxLen = 0;
        int start = 0;
        int end = 0;
        Character curChar = null;
        Map<Character, Integer> map = new HashMap<>();
        for (end = 0; end < s.length(); end++) {
            curChar = s.charAt(end);
            // map中没有存储过当前字符，或者已经存储的索引值小于start
            if (!map.containsKey(curChar) || map.get(curChar) < start) {
                // 1. 更新当前字符的索引
                map.put(curChar, end);
            }
            // map中已经存储过当前字符的索引，且存储的索引值大于start
            else {
                // 1. 更新 maxLen
                maxLen = maxLen > (end - start) ? maxLen : (end - start);
                // 2. 更新start为 当前字符已经存储的索引值 + 1
                start = map.get(curChar) + 1;
                // 3. 更新该字符的索引
                map.put(curChar, end);
            }
        }
        // 为了执行效率，只有当遇到重复字符时，才会更新索引
        // 当一直没有重复字符时，实际的窗口长度一直在增加，但是maxLen却没有更新，因此需要最后再更新一遍索引
        return maxLen > (end - start) ? maxLen : (end - start);
    };
}
