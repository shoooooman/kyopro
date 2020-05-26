package main

import "fmt"

// targetが存在する場合は最も左側の要素の添字を返す
// 存在しない場合は、targetより小さい要素が何個あるかを返す
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 2, 2, 2, 5, 7}
	target := 2
	ans := binarySearch(nums, target)
	fmt.Println(ans)
}
