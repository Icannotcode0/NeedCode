package _1_two_sum

func twoSum(nums []int, target int) []int {
	recorder := make(map[int]int)
	for i, num := range nums {
		recorder[num] = i
	}
	ans := []int{}

	for i := range nums {
		complement := target - nums[i]
		val, ok := recorder[complement]
		if !ok || val == i {
			continue
		}
		ans = append(ans, []int{i, val}...)
		break
	}
	return ans
}
