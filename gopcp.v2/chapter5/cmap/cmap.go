package cmap

import (
	"math"
	"sync/atomic"
)

const (
	MAX_CONCURRENCY            int     = 65536
	DEFAULT_BUCKET_NUMBER      int     = 16
	DEFAULT_BUCKET_LOAD_FACTOR float64 = 0.75
	DEFAULT_BUCKET_MAX_SIZE    uint64  = 1000
)

type ConcurrentMap interface {
	Concurrency() int
	Put(key string, element interface{}) (bool, error)
	Get(key string) interface{}
	Delete(key string) bool
	Len() uint64
}

type myConcurrentMap struct {
	concurrency int
	segments    []Segment
	total       uint64
}

func NewConcurrentMap(concurrency int, pairRedistributor PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, newIllegalParameterError("concurrency is too small")
	}
	if concurrency > MAX_CONCURRENCY {
		return nil, newIllegalParameterError("concurrency is too large")
	}
	cmap := &myConcurrentMap{}
	cmap.concurrency = concurrency
	cmap.segments = make([]Segment, concurrency)
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
	}
	return cmap, nil
}

func (cmap *myConcurrentMap) Concurrency() int {
	return cmap.concurrency
}

func (cmap *myConcurrentMap) Put(key string, element interface{}) (bool, error) {
	p, err := newPair(key, element)
	if err != nil {
		return false, err
	}
	s := cmap.findSegment(p.Hash())
	ok, err := s.Put(p)
	if ok {
		atomic.AddUint64(&cmap.total, 1)
	}
	return ok, err
}

func (cmap *myConcurrentMap) Get(key string) interface{} {
	keyHash := hash(key)
	s := cmap.findSegment(keyHash)
	pair := s.GetWithHash(key, keyHash)
	if pair == nil {
		return nil
	}
	return pair.Element()
}

func (cmap *myConcurrentMap) Delete(key string) bool {
	s := cmap.findSegment(hash(key))
	if s.Delete(key) {
		atomic.AddUint64(&cmap.total, ^uint64(0))
		return true
	}
	return false
}

func (cmap *myConcurrentMap) Len() uint64 {
	return atomic.LoadUint64(&cmap.total)
}

func (cmap *myConcurrentMap) findSegment(keyhash uint64) Segment {
	if cmap.concurrency == 1 {
		return cmap.segments[0]
	}
	var keyHash32 uint32
	if keyhash > math.MaxUint32 {
		keyHash32 = uint32(keyhash >> 32)
	} else {
		keyHash32 = uint32(keyhash)
	}
	return cmap.segments[int(keyHash32>>16)%(cmap.concurrency-1)]
}
