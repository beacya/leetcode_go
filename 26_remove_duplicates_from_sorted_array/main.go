package main

import "fmt"

func removeDuplicates(nums []int) int {
	arrLen := len(nums)
	startKey := 0
	for searchKey := 0; searchKey < arrLen; searchKey++ {
		if nums[startKey] != nums[searchKey] && startKey+1 != searchKey {
			nums[startKey+1] = nums[searchKey]
		}
		startKey++
	}
	fmt.Println(nums)
	return startKey + 1
}

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 10, 10, 12}
	nums2 := []int{1, 1, 2}
	nums3 := []int{1, 2}
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(removeDuplicates(nums3))
}
