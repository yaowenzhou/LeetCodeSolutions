# [LeetCode第48题](https://leetcode.cn/problems/rotate-image/)
给定一个 $n×n$ 的二维矩阵`matrix`表示一个图像。请你将图像顺时针旋转`90`度。
你必须在`原地`旋转图像，这意味着你需要直接修改输入的二维矩阵。`请不要`使用另一个矩阵来旋转图像。
```
示例1:
```
$$
 \left|\begin{matrix}
    1 & 2 & 3 \\
    4 & 5 & 6 \\
    7 & 8 & 9
   \end{matrix} \right| => 
\left|\begin{matrix}
    7 & 4 & 1 \\
    8 & 5 & 2 \\
    9 & 6 & 3
   \end{matrix} \right|
$$
```
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]
```

<blockquote class="warn">示例2</blockquote>

$$
 \left|\begin{matrix}
    5 & 1 & 9 & 11 \\
    2 & 4 & 8 & 10 \\
    13 & 3 & 6 & 7 \\
    15 & 14 & 12 & 16
   \end{matrix} \right| => 
\left|\begin{matrix}
    15 & 13 & 2 & 5 \\
    14 & 3 & 4 & 1 \\
    12 & 6 & 8 & 9 \\
    16 & 7 & 10 & 11
   \end{matrix} \right|
$$
<blockquote class="warn">
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
</blockquote>

# 解题方案: 
1. 上下翻转

$$
 \left|\begin{matrix}
    5 & 1 & 9 & 11 \\
    2 & 4 & 8 & 10 \\
    13 & 3 & 6 & 7 \\
    15 & 14 & 12 & 16
   \end{matrix} \right| \frac{水平翻转}{=>}
\left|\begin{matrix}
    15 & 14 & 12 & 16 \\
    13 & 3 & 6 & 7 \\
    2 & 4 & 8 & 10 \\
    5 & 1 & 9 & 11 
   \end{matrix} \right|
$$

2. 左上右下对角线为轴翻转

$$
 \left|\begin{matrix}
    15 & 14 & 12 & 16 \\
    13 & 3  & 6  & 7  \\
    2  & 4  & 8  & 10 \\
    5  & 1  & 9  & 11 
   \end{matrix} \right| \frac{对角线翻转}{=>}
\left|\begin{matrix}
    15 & 13 & 2  & 5 \\
    14 & 3  & 4  & 1 \\
    12 & 6  & 8  & 9 \\
    16 & 7  & 10 & 11
   \end{matrix} \right|
$$

3. go代码实现
```go
func rotate(matrix [][]int) {
    n := len(matrix)
    // 水平翻转
    for i := 0; i < n/2; i++ {
        matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
    }
    // 主对角线翻转
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}
```