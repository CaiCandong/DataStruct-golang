package insertSort

func BInsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			temp := nums[i]
			left, right := 0, i-1
			for left <= right {
				mid := (left + right) / 2
				if temp < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			for j := i; j > left; j-- {
				nums[j] = nums[j-1]
			}
			nums[left] = temp
		}
	}
}
