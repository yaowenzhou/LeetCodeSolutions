/**
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 12:55:26
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 13:07:17
 * @version      : 
 * @Description  : 
 */
public class App {
    public static void main(String[] args) throws Exception {
        ListNode l1 = new ListNode(2);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);
        ListNode l2 = new ListNode(5);
        l2.next = new ListNode(6);
        l2.next.next = new ListNode(4);

        Solution1 sol1 = new Solution1();
        ListNode ret = sol1.addTwoNumbers(l1, l2);

        while (ret != null) {
            System.out.print(ret.val);
            ret = ret.next;
        }
    }
}
