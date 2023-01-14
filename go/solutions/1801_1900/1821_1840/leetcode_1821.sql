-- https://leetcode.cn/problems/find-customers-with-positive-revenue-this-year/
-- 这是一个sql题
-- Write your MySQL query statement below
SELECT customer_id FROM Customers WHERE year = 2021 AND revenue > 0;