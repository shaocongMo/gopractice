package leetcode

import(
	"fmt"
	"testing"
)

func Test800(t *testing.T){
	myHashMap := Constructor()
	myHashMap.Put(1, 1)
	myHashMap.Put(2, 2)
	fmt.Println(myHashMap.Get(1))
	fmt.Println(myHashMap.Get(2))
	fmt.Println(myHashMap.Get(3))
	myHashMap.Put(2, 1)
	fmt.Println(myHashMap.Get(2))
	myHashMap.Remove(2)
	fmt.Println(myHashMap.Get(2))
	fmt.Println(myHashMap)
}
