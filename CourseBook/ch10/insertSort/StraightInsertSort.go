package insertSort

// 直接插入排序
func StraightInsertSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			temp := nums[i]
			var j int
			for j = i - 1; j >= 0 && temp < nums[j]; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = temp
		}
	}
	return nums
}
