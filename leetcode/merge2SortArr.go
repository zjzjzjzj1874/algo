package leetcode

func merge2SortArr(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, m+n)

	p1 := 0
	p2 := 0
	for p1 < m && p2 < n {
		if nums1[p1] == 0 {
			p1++
			continue
		}
		if nums2[p2] == 0 {
			p2++
			continue
		}
		if nums1[p1] < nums2[p2] {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}
	}

	for p1 < m {
		if nums1[p1] == 0 {
			p1++
			continue
		}
		res = append(res, nums1[p1])
		p1++
	}

	for p2 < n {
		if nums2[p2] == 0 {
			p2++
			continue
		}
		res = append(res, nums2[p2])
		p2++
	}

	copy(nums1, res)
}
