package main

import "fmt"

func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	newLen := m + n
	var num3 []int
	mStart := 0
	nStart := 0
	for i := 0; i < newLen; i++ {
		if mStart < m && nStart >= n {
			num3 = append(num3, nums1[mStart])
			mStart++
			break
		} else if nStart < n && mStart >= m {
			num3 = append(num3, nums2[nStart])
			nStart++
			break
		}
		if nums1[mStart] >= nums2[nStart] {
			num3 = append(num3, nums2[nStart])
			nStart++
		} else {
			num3 = append(num3, nums1[mStart])
			mStart++
		}
	}
	nums1 = num3
	fmt.Println(num3)
	fmt.Println(nums1)

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	mergeArray(nums1, m, nums2, n)
}
