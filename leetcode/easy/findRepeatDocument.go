package easy

// LCR 120. 寻找文件副本
// 设备中存有 n 个文件，文件 id 记于数组 documents。若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。
//
// 示例 1：
//
// 输入：documents = [2, 5, 3, 0, 5, 0]
// 输出：0 或 5
//
// 提示：
//
// 0 ≤ documents[i] ≤ n-1
// 2 <= n <= 100000

// 解题：使用数组实现Hash表，数组的大小是0，n-1，则nums[i] = i;
// Hash表实现比较简单，这里略过
func findRepeatDocument(documents []int) int {
	n := len(documents)
	for i := 0; i < n; i++ {
		for documents[i] != documents[documents[i]] {
			documents[i], documents[documents[i]] = documents[documents[i]], documents[i]
		}
	}

	for i := 0; i < n; i++ {
		if documents[i] != i {
			return documents[i]
		}
	}

	return -1
}
