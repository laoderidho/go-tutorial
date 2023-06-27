package removeelemen

import "fmt"

func RemoveElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Println(nums)
	return len(nums)
}
