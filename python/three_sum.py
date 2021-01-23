from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        return self.three_sum1(nums=nums)

    def three_sum1(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return []

        nums.sort()
        res = set()
        for i, v in enumerate(nums[:-2]):
            if i >= 1 and v == nums[i - 1]:
                continue

            dup = {}

            for x in nums[i + 1 :]:
                if x not in dup:
                    dup[-v - x] = 1
                else:
                    res.add((v, -v - x, x))
                pass

        return [list(item) for item in res]

    def three_sum2(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return []
        nums.sort()

        res = []
        for i, v in enumerate(nums[:-2]):
            if i >= 1 and v == nums[i - 1]:
                continue

            left, right = i + 1, len(nums) - 1

            while left < right:
                two_sum = nums[left] + nums[right]
                if two_sum > -v:
                    right -= 1
                elif two_sum < -v:
                    left += 1
                else:
                    res.append([v, nums[left], nums[right]])

                    while left < right and nums[left] == nums[left + 1]:
                        left += 1

                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1

                    left += 1
                    right -= 1

        return res
