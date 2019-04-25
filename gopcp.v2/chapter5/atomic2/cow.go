package cow

import (
	"errors"
	"sync/atomic"
)

type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}

type concurrentArray struct {
	length uint32
	val    atomic.Value
}

func NewConcurrentArray(length uint32) ConcurrentArray {
	array := concurrentArray{}
	array.length = length
	array.val.Store(make([]int, array.length))
	return &array
}

func (array *concurrentArray) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	newArray := make([]int, array.length)
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return
}

func (array *concurrentArray) Get(index uint32) (elem int, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	elem = array.val.Load().([]int)[index]
	return
}

func (array *concurrentArray) checkIndex(index uint32) (err error) {
	if index >= array.length || index < 0 {
		err = errors.New("index out of array")
	}
	return
}

func (array *concurrentArray) checkValue() (err error) {
	return
}

func (array *concurrentArray) Len() uint32 {
	return array.length
}
