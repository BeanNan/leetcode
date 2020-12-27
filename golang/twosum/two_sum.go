package twosum

func twoSum(nums []int, target int) []int {

	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x]+nums[y] == target {
				return []int{x, y}
			}
		}
	}

	return []int{0, 0}
}

func twoSum2(nums []int, target int) []int {
	mapping := map[int]int{}

	for index, value := range nums {
		surplus := target - value

		if preIndex, ok := mapping[surplus]; ok {
			return []int{preIndex, index}
		}
		mapping[value] = index
	}

	return []int{0, 0}
}
