/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 10:40:29
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 12:52:08
 * @version      :
 * @Description  :
 */
#include <iostream>
#include "ListNode.hpp"
#include "Solution1.hpp"

using namespace std;
int main() {
    ListNode* l1 = new ListNode(2);
    l1->next = new ListNode(4);
    l1->next->next = new ListNode(7);
    ListNode* l2 = new ListNode(5);
    l2->next = new ListNode(6);
    l2->next->next = new ListNode(4);
    
    Solution1 sol1 = Solution1();
    ListNode* ret = sol1.addTwoNumbers(l1, l2);
    while(ret != nullptr) {
        cout << ret->val;
        ret = ret->next;
    }
    return 0;
}
