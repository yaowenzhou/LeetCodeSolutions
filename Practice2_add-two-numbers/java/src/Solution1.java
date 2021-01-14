/**
 * @Package      : []
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 12:56:30 Tuesday
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 13:04:19
 * @Version      : 
 * @Description  : 
 */

public class Solution1 {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode ret = new ListNode();
        ListNode curNode = ret;
        int extra = 0;
        int tmpSum = 0;
        while (l1 != null || l2 != null || extra != 0) {
            tmpSum = (l1 != null ? l1.val : 0) + (l2 != null ? l2.val : 0) + extra;
            extra = tmpSum / 10;
            curNode.next = new ListNode(tmpSum % 10);
            curNode = curNode.next;

            if (l1 != null)
                l1 = l1.next;
            if (l2 != null)
                l2 = l2.next;
        }
        return ret.next;
    }
}
