package twosum

// TwoSum implemens a solution for "Two Sum" problem
// 	see: https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	complements := map[int]int{}
	for i, a := range nums {
		b := target - a
		if j, exists := complements[b]; exists {
			return []int{j, i}
		}

		complements[a] = i
	}

	return nil
}
