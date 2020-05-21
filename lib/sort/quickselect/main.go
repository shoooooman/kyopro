package main

import (
	"fmt"
	"log"
	"math/rand"
)

// Lomuto partition scheme
// 左から順に見ていき、適宜前に出てきた要素と交換する方法
func partition(nums []int, left, right int) int {
	pivot := right

	j := left
	for i := left; i < right; i++ {
		if nums[i] < nums[pivot] {
			// pivot未満の値(ith)と以上の値(jth)を交換
			nums[i], nums[j] = nums[j], nums[i]
			// jはpivot以上の数の位置を表す
			j++
		}
	}
	nums[pivot], nums[j] = nums[j], nums[pivot]
	return j
}

// Hoare partition scheme
// 左と右を交換していく方法
func partition2(nums []int, left, right int) int {
	pivot := right
	i, j := left, right-1
	for {
		if i < 0 || i > len(nums)-1 || pivot < 0 || pivot > len(nums)-1 {
			log.Println(i, pivot, len(nums))
		}
		for nums[i] < nums[pivot] {
			i++
		}
		for i < j && nums[j] > nums[pivot] {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[pivot], nums[i] = nums[i], nums[pivot]
	return i
}

// ランダムにピボットを取る
func randomPartition(nums []int, left, right int) int {
	pivot := rand.Intn(right-left+1) + left
	nums[pivot], nums[right] = nums[right], nums[pivot]
	return partition(nums, left, right)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	// pivot := partition(nums, left, right)
	// pivot := partition2(nums, left, right)
	pivot := randomPartition(nums, left, right)
	if k == pivot {
		return nums[pivot]
	}
	if k < pivot {
		return quickSelect(nums, left, pivot-1, k)
	}
	return quickSelect(nums, pivot+1, right, k)
}

func main() {
	nums := []int{4, 2, 9, 1, 0, 3, 2, 5}
	k := 3
	ans := quickSelect(nums, 0, len(nums)-1, k)
	fmt.Println(ans)
}
