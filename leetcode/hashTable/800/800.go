package leetcode

type MyHashMap struct {
    bucket [][]*HashMapElem
}

type HashMapElem struct{
	key int
	value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	bucket := make([][]*HashMapElem, 1000)
	return MyHashMap{bucket : bucket}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	bucketIndex := getHashKey(key)
	for _, elem := range this.bucket[bucketIndex]{
		if elem.key == key{
			elem.value = value
			return
		}
	}
	this.bucket[bucketIndex] = append(this.bucket[bucketIndex], &HashMapElem{key: key, value: value})
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	bucketIndex := getHashKey(key)
	for _, elem := range this.bucket[bucketIndex]{
		if elem.key == key{
			return elem.value
		}
	}
	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    bucketIndex := getHashKey(key)
	for i, elem := range this.bucket[bucketIndex]{
		if elem.key == key{
			this.bucket[bucketIndex] = append(this.bucket[bucketIndex][:i], this.bucket[bucketIndex][i+1:]...)
		}
	}
}

func getHashKey(key int) int{
	return key % 1000
}