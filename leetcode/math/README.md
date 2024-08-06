
## 数学问题
数学问题的算法题，一般是先找规律，然后使用归纳法总结数学规律；能找到规律，答案呼之欲出；

再进一步，当一个题目的入参和出参都是一个int类型的时候，先从小到大枚举入参和出参的数学规律；四成的题目基本可以总结出规律。


### Easy
- [x] [292.Nim游戏](https://leetcode.cn/problems/nim-game/description/)

### medium
- [ ] [877. 石子游戏](https://leetcode.cn/problems/stone-game/description/)

### hard
- [ ] [810. 黑板异或游戏](https://leetcode.cn/problems/chalkboard-xor-game/description/)


### [概率问题](https://leetcode.cn/problems/implement-rand10-using-rand7/solutions/427572/cong-pao-ying-bi-kai-shi-xun-xu-jian-jin-ba-zhe-da/)

#### 均匀硬币，产生等概率
现有一枚不均匀的硬币 coin()，能够返回 0、1 两个值，其概率分别为 0.6、0.4。要求使用这枚硬币，产生均匀的概率分布。即编写一个函数 coin_new() 使得它返回 0、1 的概率均为 0.5。
```shell
# 不均匀硬币，返回 0、1 的概率分别为 0.6、0.4
int coin() {
  return (rand() % 10) > 5;
}
```
统计抛两次硬币的结果的概率分布：

| 结果 | 0            | 1            |
|----|--------------|--------------|
| 0  | 0.6*0.6=0.36 | 0.6*0.4=0.24 |
| 1  | 0.4*0.6=0.24 | 0.4*0.4=0.16 |
不难发现，连续抛两枚硬币得到 0 1 和 1 0 的概率分布是相同的。因此这道题的解法就是连续抛两次硬币，如果得到 0 1，返回 0；如果得到 1 0，返回 1；如果两次结果相同，则重新抛。
以此类推，无论这枚不均匀硬币的概率是多少，都可以用这种方法得到等概率的结果。

代码如下：
```shell
int coin_new() {
    while (true) {
        int a = coin();
        if (coin() != a) return a;
    }
}
```

#### 均匀硬币，产生不等概率
  现有一枚均匀的硬币 coin()，能够返回 0、1 两个值，其概率均为 0.5。要求编写一个函数 coin_new()，使得它返回指定的 0、1 概率分布。
```shell
# 均匀硬币
int coin() {
  return rand() % 2;
}
```

* P(0) = 1/4，P(1) = 3/4
对于均匀硬币而言，连续抛两次，得到 0 0、0 1、1 0、1 1 的概率均为 1/4。显然，只需要连续抛两次硬币，如果得到 0 0，返回 0，其他情况返回 1。
```shell
int coin_new() {
  return coin() || coin();
}
```
测试输出：
```
0: 0.249249
1: 0.750751
```

* P(0) = 1/3，P(1) = 2/3
连续抛两次硬币。如果得到 1 1，返回 0；如果得到 1 0 或 0 1，返回 1；如果得到 0 0，继续抛硬币。
```shell
int coin_new() {
    while (true) {
        int a = coin(), b = coin();
        if (a && b) return 0;
        if (a || b) return 1;
    }
}
```