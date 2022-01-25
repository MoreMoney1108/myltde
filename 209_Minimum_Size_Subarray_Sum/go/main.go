package minimumsizesubarraysum

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	l := int(math.MaxInt64)
	var start, end = 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		if sum >= target {
			curLen := end - start + 1
			if curLen == 1 {
				return curLen
			}
			if curLen < l {
				l = curLen
			}

			sum -= nums[start] + nums[end]
			start++
		} else {
			end++
		}
	}

	if l == int(math.MaxInt64) {
		return 0
	}
	return l
}
