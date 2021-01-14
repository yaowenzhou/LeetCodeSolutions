/**
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 14:33:44
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 15:09:17
 * @version      : 
 * @Description  : 
 */
public class App {
    public static void main(String[] args) throws Exception {
        String str = "abcabcbb";

        Solution1 sol1 = new Solution1();
        System.out.println(sol1.lengthOfLongestSubstring(str));

        Solution2 sol2 = new Solution2();
        System.out.println(sol2.lengthOfLongestSubstring(str));
    }
}
