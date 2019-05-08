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
func ShellSort(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
			}
		}
		key = key / 2
	}
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
	result := []int{}
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] > right[rightIndex] {
			result = append(result, right[rightIndex])
			rightIndex++
		} else {
			result = append(result, left[leftIndex])
			leftIndex++
		}
	}
	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)
	return result
}

// 快速排序
func QuickSort(nums []int) {
	qsort(nums)
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

// from https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F#Go
func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

// 堆排序
// 用一维数组代表树
// 父节点 i
// 父节点的左子节点 2*i + 1
// 父节点的右子节点 2*i + 2
func HeapSort(nums []int) {
	heapSize := len(nums)
	heapNum := heapSize / 2
	// 构建最大堆
	for i := heapNum; i > -1; i-- {
		maxHeap(nums, i, heapSize-1)
	}
	for i := heapSize - 1; i > -1; i-- {
		// 把最大堆的堆顶（最大值）放到数组最后
		nums[i], nums[0] = nums[0], nums[i]
		// 重新构建最大堆 新堆大小 = 旧堆大小 - 1 （即不对已排好的数组尾巴进行堆处理）
		maxHeap(nums, 0, i-1)
	}
}

func maxHeap(nums []int, i int, end int) {
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && nums[r] > nums[l] {
		n = r
	}
	if nums[i] > nums[n] {
		return
	}
	nums[i], nums[n] = nums[n], nums[i]
	maxHeap(nums, n, end)
}

// 计数排序
// 找出待排序的数组中最大和最小的元素
// 统计数组中每个值为 i的元素出现的次数，存入数组C 的第  i项
// 对所有的计数累加（从 C 中的第一个元素开始，每一项和前一项相加）
// 反向填充目标数组：将每个元素 i放在新数组的第 C[i] 项，每放一个元素就将 C[i] 减去1
func CountingSort(nums []int) {
	size := len(nums)
	if size <= 0 {
		return
	}
	min, max := nums[0], nums[0]
	for i := 1; i < size; i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	countSize := max - min + 1
	count := make([]int, countSize)
	for i := 0; i < size; i++ {
		count[nums[i]-min]++
	}
	for j, i := 0, 0; j < countSize; {
		if count[j] != 0 {
			count[j]--
			nums[i] = j + min
			i++
		} else {
			j++
		}

	}
}

// 桶排序
func BucketSort(nums []int) {
	size := len(nums)
	if size <= 0 {
		return
	}
	min, max := nums[0], nums[0]
	for i := 1; i < size; i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	// 依据实际情况设定桶的容量
	bucketSize := 5
	bucketnum := (max-min)/bucketSize + 1
	bucket := make([][]int, bucketnum)
	for i := 0; i < bucketnum; i++ {
		bucket[i] = []int{}
	}
	for i := 0; i < size; i++ {
		index := (nums[i] - min) / bucketSize
		bucket[index] = append(bucket[index], nums[i])
	}
	index := 0
	for i := 0; i < bucketnum; i++ {
		InsertSort(bucket[i])
		for j := 0; j < len(bucket[i]); j++ {
			nums[index] = bucket[i][j]
			index++
		}
	}
}

// 基数排序
func RedixSort(nums []int) {
	// index [0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]
	// redix [_, -9, -8, -7, -6, -5, -4, -3, -2, -1,  0,  1,  2,  3,  4,  5,  6,  7,  8,  9]
	size := len(nums)
	if size <= 0 {
		return
	}
	max := nums[0]
	for i := 1; i < size; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	exp := 1
	for max/exp > 0 {
		redix := make([][]int, 20)
		for i := 0; i < 20; i++ {
			redix[i] = []int{}
		}
		pushToRedix(nums, exp, redix)
		exp = exp * 10
	}
}

func pushToRedix(nums []int, remNum int, redix [][]int) {
	index := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]*nums[i] >= remNum%10-1 {
			if index == -1 {
				index = i
			}
			redixIndex := nums[i]/remNum%10 + 10
			redix[redixIndex] = append(redix[redixIndex], nums[i])
		}
	}
	for i := index; i < len(redix); i++ {
		for j := 0; j < len(redix[i]); j++ {
			nums[index] = redix[i][j]
			index++
		}
	}
}
