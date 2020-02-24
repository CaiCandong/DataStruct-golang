package insertSort

func BPathsInsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	helper := make([]int, len(nums))
	helper[0] = nums[0]
	left, right := 1, len(nums)-1
	for i := 1; i < len(nums); i++ {
		if nums[i] > helper[0] {
			//左插入
			j := left
			for j > 0 && helper[j-1] > nums[i] {
				helper[j] = helper[j-1]
				j--
			}
			helper[j] = nums[i]
			left++
		} else {
			//右插入
			j := right
			for j < len(nums)-1 && helper[j+1] < nums[i] {
				helper[j] = helper[j+1]
				j++
			}
			helper[j] = nums[i]
			right--
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = helper[(right+i+1)%len(nums)]
	}
}
