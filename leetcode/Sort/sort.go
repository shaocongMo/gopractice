package leetcode

//冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := true
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				flag = false
			}
		}
		// 已是有序状态， 无须继续排序
		if flag {
			break
		}
	}
}

// 选择排序
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}
		if min_index != i {
			nums[i], nums[min_index] = nums[min_index], nums[i]
		}
	}
}

// 插入排序
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i
		for ; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			} else {
				break
			}
		}
	}
}

// 希尔排序
func ShellSort(nums []int) {
	gap := len(nums) / 2
	for gap > 0 {
		for i, j := 0, gap; i < len(nums)-gap && j < len(nums); i, j = i+1, j+1 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		gap = gap / 2
	}
	InsertSort(nums)
}

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	nums = merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
	return nums
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	index, leftIndex, rightIndex := 0, 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] > right[rightIndex] {
			result[index] = right[rightIndex]
			rightIndex++
		} else {
			result[index] = left[leftIndex]
			leftIndex++
		}
		index++
	}
	for leftIndex < len(left) {
		result[index] = left[leftIndex]
		index++
		leftIndex++
	}
	for rightIndex < len(right) {
		result[index] = right[rightIndex]
		index++
		rightIndex++
	}
	return result
}

// 快速排序
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left int, right int) {
	if left < right {
		partitionIndex := partition(nums, left, right)
		quickSort(nums, left, partitionIndex-1)
		quickSort(nums, partitionIndex+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	// 基准值
	pivot := right
	right = pivot - 1
	for {
		for nums[right] >= nums[pivot] && right > left {
			right--
		}
		for nums[left] <= nums[pivot] && left <= right {
			left++
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right++
	}
	if nums[left] > nums[pivot] {
		nums[left], nums[pivot] = nums[pivot], nums[left]
		pivot = left
	}
	return pivot
}
