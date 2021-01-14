import java.util.HashSet;
import java.util.Set;

/**
 * @Package      : []
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 14:34:23 Tuesday
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 14:47:39
 * @Version      : 
 * @Description  : 
 */

public class Solution1 {
    public int lengthOfLongestSubstring(String s) {
        int maxLen = 0;
        int start = 0;
        int end = 0;
        Set<Character> set = new HashSet<>();
        for (start = 0; start < s.length(); start++) {
            for (end = start; end < s.length(); end++) {
                // 如果该字符已经在集合中出现
                // 1. 清空set
                // 2. 终止本次循环
                if (set.contains(s.charAt(end))) {
                    set.clear();
                    break;
                }
                // 如果该字符在集合中未出现过
                // 1. 将该字符插入到集合中
                else {
                    set.add(s.charAt(end));
                }
            }
            // 内循环终止后，得到的是当前外循环能够得到的最长子串，此时和更新maxLen
            maxLen = maxLen > (end - start) ? maxLen : (end - start);
        }
        return maxLen;
    }
}
