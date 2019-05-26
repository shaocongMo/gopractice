package leetcode

import(
	"fmt"
	"testing"
)

func Test799(t *testing.T){
	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(1))
	fmt.Println(myHashSet.Contains(3))
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(2))
	myHashSet.Remove(2)
	fmt.Println(myHashSet.Contains(2))
}
