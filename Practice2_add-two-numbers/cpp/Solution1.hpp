/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 10:43:22
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 13:44:48
 * @version      : 
 * @Description  : 
 */

#include "ListNode.hpp"

#ifndef __SOLUTION1__
#define __SOLUTION1__
class Solution1 {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        // ret 保存需要返回的指针
        ListNode* ret = new ListNode(0, nullptr);
        // tmpNode 表示需要返回的链表的进度
        ListNode* tmpNode = ret;
        // extra 表示需要进位的数字
        int extra = 0;
        // tmpSum 表示当前操作的节点数字以及 extra 之和
        int tmpSum = 0;

        // 当 l1、 l2都遍历完毕，且没有需要进位的数字，则结束循环
        while(l1 != nullptr || l2 != nullptr || extra != 0) {
            tmpSum = (l1 != nullptr ? l1->val : 0) + (l2 != nullptr ? l2->val : 0) + extra;
            tmpNode->next = new ListNode(tmpSum % 10);
            tmpNode = tmpNode->next;
            extra = tmpSum / 10;

            if(l1 != nullptr) l1 = l1->next;
            if(l2 != nullptr) l2 = l2->next;
        }
        tmpNode = ret->next;
        delete ret;
        return tmpNode;
    }
};
#endif