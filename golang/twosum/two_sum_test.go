package twosum

import (
	"testing"

	"github.com/BeanNan/leetcode/golang/utils"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expect := []int{0, 1}
	actual := twoSum(nums, target)

	isEqual := utils.EqualSliceInt(expect, actual)

	if !isEqual {
		t.Errorf("expected result: %v, actual result: %v", expect, actual)
	}

}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expect := []int{1, 2}
	actual := twoSum(nums, target)

	isEqual := utils.EqualSliceInt(expect, actual)

	if !isEqual {
		t.Errorf("expected result: %v, actual result: %v", expect, actual)
	}

}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expect := []int{0, 1}
	actual := twoSum(nums, target)

	isEqual := utils.EqualSliceInt(expect, actual)

	if !isEqual {
		t.Errorf("expected result: %v, actual result: %v", expect, actual)
	}

}

func TestTwoSum4(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expect := []int{0, 1}
	actual := twoSum2(nums, target)

	isEqual := utils.EqualSliceInt(expect, actual)

	if !isEqual {
		t.Errorf("expected result: %v, actual result: %v", expect, actual)
	}

}

func TestTwoSum5(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expect := []int{1, 2}
	actual := twoSum2(nums, target)

	isEqual := utils.EqualSliceInt(expect, actual)

	if !isEqual {
		t.Errorf("expected result: %v, actual result: %v", expect, actual)
	}

}

func TestTwoSum6(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expect := []int{0, 1}
	actual := twoSum2(nums, target)

	isEqual := utils.EqualSliceInt(expect, actual)

	if !isEqual {
		t.Errorf("expected result: %v, actual result: %v", expect, actual)
	}

}
