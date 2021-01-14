/*
 * @Author       : yaowenzhou
 * @Date         : 2021-01-12 10:44:37
 * @LastEditors  : yaowenzhou
 * @LastEditTime : 2021-01-12 10:46:24
 * @version      : 
 * @Description  : 
 */

#ifndef __LIST_NODE__
#define __LIST_NODE__
struct ListNode {
    int val;
    ListNode* next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode* next) : val(x), next(next) {}
};
#endif
