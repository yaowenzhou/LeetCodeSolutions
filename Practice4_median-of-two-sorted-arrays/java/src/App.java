/**
 * @Author       : yaowenzhou
 * @Date         : 2021-01-13 20:58:11
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-13 22:27:58
 * @version      : 
 * @Description  : 
 */
public class App {
    public static void main(String[] args) throws Exception {
        int[] nums1 = new int[] { 4 };
        int[] nums2 = new int[] { 1, 2, 3 };
        Solution1 sol1 = new Solution1();
        System.out.println(sol1.findMedianSortedArrays(nums1, nums2));
    }
}
