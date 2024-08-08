# algo
算法



* [综合](./leetcode/README.md)
* [回溯算法](./leetcode/backtrack/README.md)
* [二分法](./leetcode/dichotomy/README.md)
* [动态规划](./leetcode/dp/README.md)
* [单调栈](./leetcode/monitonicStack/README.md)
* [前缀和]()

二维前缀和
-[ ] 1277. 统计全为 1 的正方形子矩阵（双倍经验 221. 最大正方形）
-[ ] 1504. 统计全 1 子矩形
-[ ] 304. 二维区域和检索 - 矩阵不可变

前缀和
- [x] 560.和为 K 的子数组
- [ ] 930.和相同的二元子数组 1592（同 560 题）
- [ ] 1524.和为奇数的子数组数目 
- [ ] 974.和可被 K 整除的子数组
- [ ] 523.连续的子数组和
- [ ] 3026.最大好子数组和
- [ ] 525.连续数组
- [ ] 面试题 17.05. 字母与数字
- [ ] 1124.表现良好的最长时间段 1908
- [ ] 2488.统计中位数为 K 的子数组
- [ ] 1590.使数组和能被 P 整除 2039
- [ ] 2949.统计美丽子字符串 II 2445
- [ ] 1983.范围和相等的最宽索引对（会员题）
- [ ] 2489.固定比率的子字符串数（会员题）
- [ ] 2955.同端子串的数量（会员题）

异或前缀和
- [ ] 1310.子数组异或查询 1460
- [ ] 1177.构建回文串检测 1848
- [ ] 1371.每个元音包含偶数次的最长子字符串 2041
- [ ] 1542.找出最长的超赞子字符串 2222
- [ ] 1915.最美子字符串的数目 2235
- [ ] 2791.树中可以形成回文的路径数 2677

* [图](./leetcode/graph/README.md)


### 计算机指标

L1 cache reference ......................... 0.5 ns
Branch mispredict ............................ 5 ns
L2 cache reference ........................... 7 ns
Mutex lock/unlock ........................... 25 ns
Main memory reference ...................... 100 ns
Compress 1K bytes with Zippy ............. 3,000 ns  =   3 µs
Send 2K bytes over 1 Gbps network ....... 20,000 ns  =  20 µs
SSD random read ........................ 150,000 ns  = 150 µs
Read 1 MB sequentially from memory ..... 250,000 ns  = 250 µs
Round trip within same datacenter ...... 500,000 ns  = 0.5 ms
Read 1 MB sequentially from SSD* ..... 1,000,000 ns  =   1 ms
Disk seek ........................... 10,000,000 ns  =  10 ms
Read 1 MB sequentially from disk .... 20,000,000 ns  =  20 ms
Send packet CA->Netherlands->CA .... 150,000,000 ns  = 150 ms

* 参考
  * [细说Cache-L1/L2/L3/TLB](https://zhuanlan.zhihu.com/p/31875174)
  * [计算机常见指标](https://gist.github.com/hellerbarde/2843375)
  * [让 CPU 告诉你硬盘和网络到底有多慢](https://cizixs.com/2017/01/03/how-slow-is-disk-and-network/)
  * [关于CPU、内存、硬盘及网络传输速度的整理](https://superxlcr.github.io/2020/01/11/%E5%85%B3%E4%BA%8ECPU%E3%80%81%E5%86%85%E5%AD%98%E3%80%81%E7%A1%AC%E7%9B%98%E5%8F%8A%E7%BD%91%E7%BB%9C%E4%BC%A0%E8%BE%93%E9%80%9F%E5%BA%A6%E7%9A%84%E6%95%B4%E7%90%86/)
  * [磁盘比内存慢几万倍](https://xiaolincoding.com/os/1_hardware/storage.html#%E5%AD%98%E5%82%A8%E5%99%A8%E7%9A%84%E5%B1%82%E6%AC%A1%E7%BB%93%E6%9E%84)