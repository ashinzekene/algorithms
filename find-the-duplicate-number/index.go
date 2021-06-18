package algorithms

import "sort"

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	previous := nums[len(nums)-1]
	for _, num := range nums {
		if num == previous {
			return num
		}
		previous = num
	}
	return -1
}

func findDuplicate1(nums []int) int {
	numsMap := make(map[int]bool)

	for _, num := range nums {
		if numsMap[num] {
			return num
		}
		numsMap[num] = true
	}

	return -1
}

func findDuplicate2(nums []int) int {
	fast := nums[0]
	slow := nums[0]
	started := false

	for fast != slow || !started {
		started = true
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = nums[0]
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
