import java.util.Arrays;

/**
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 10:18:56
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 10:35:29
 * @version      : 
 * @Description  : 
 */
public class App {
    public static void main(String[] args) throws Exception {
        int[] src = { 3, 2, 4 };
        int target = 6;

        System.out.println("----> Solution1");
        Solution1 sol1 = new Solution1();
        int[] ret1 = sol1.twoSum(src, target);
        System.out.println(Arrays.toString(ret1));

        System.out.println("----> Solution2");
        Solution2 sol2 = new Solution2();
        int ret2[] = sol2.twoSum(src, target);
        System.out.println(Arrays.toString(ret2));
    }
}
