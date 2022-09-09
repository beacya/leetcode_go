package main

import "fmt"

func removeDuplicates(nums []int) int {
	arrLen := len(nums)
	startKey := 0
	for searchKey := 1; searchKey < arrLen; searchKey++ {
		if nums[startKey] != nums[searchKey] {
			if (searchKey - startKey) > 1 {
				nums[startKey+1] = nums[searchKey]
			}
			startKey++
		}
	}
	return startKey + 1
}

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 10, 10, 12}
	nums2 := []int{1, 1, 2}
	nums3 := []int{1, 2, 2}
	nums4 := []int{1, 2, 3}
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(removeDuplicates(nums3))
	fmt.Println(removeDuplicates(nums4))
}
