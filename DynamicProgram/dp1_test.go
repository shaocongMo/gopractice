package leetcode

import "testing"

func TestDp1(t *testing.T) {
	arr := []int{3, 34, 4, 12, 5, 2}
	t.Logf("arr: %v S:%d Output:%v", arr, 9, Solu(arr, 9))
	t.Logf("arr: %v S:%d Output:%v", arr, 10, Solu(arr, 10))
	t.Logf("arr: %v S:%d Output:%v", arr, 11, Solu(arr, 11))
	t.Logf("arr: %v S:%d Output:%v", arr, 12, Solu(arr, 12))
	t.Logf("arr: %v S:%d Output:%v", arr, 13, Solu(arr, 13))
}
