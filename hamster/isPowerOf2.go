package hamster

func isPowerOf2(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	fn := float64(n)
	for {
		if fn > 1 {
			fn = fn / 2
			continue
		} else {
			break
		}
	}

	return fn == 1
}
