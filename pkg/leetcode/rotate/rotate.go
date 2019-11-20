package rotate

func Rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	nums = append(reverse(nums[:length-k]), reverse(nums[length-k:])...)
	reverse(nums)
}

func reverse(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
	return nums
}
