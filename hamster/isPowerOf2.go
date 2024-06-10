package hamster

func isPowerOf2(n int) bool {
	return n&(n-1) == 0
}
