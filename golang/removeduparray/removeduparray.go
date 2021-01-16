package removeduparray

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		next := nums[i]
		if prev == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			//nums = remove(nums, i)
			fmt.Println(nums)
		}
		prev = next
	}

	return len(nums)
}

func remove(nums []int, index int) []int {

	//for i := index; i < len(nums)-1; i++ {
	//	nums[i] = nums[i+1]
	//}
	//copy(nums[index:], nums[index+1:])
	nums = append(nums[:index], nums[index+1:]...)
	return nums
}
