package mysqldump

import (
	// "errors"
	// "fmt"
	// "io"
	"os"
	"sync"
	"sync/atomic"
)

type Data []byte

type DataFile interface {
	Write(d Data) (err error)
	Close() error
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	woffset int64
}

func (df *myDataFile) Write(d Data) (err error) {
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.woffset)
		if atomic.CompareAndSwapInt64(&df.woffset, offset, (offset + int64(len(d)))) {
			break
		}
	}

	df.fmutex.Lock()
	defer func() {
		df.fmutex.Unlock()
	}()
	_, err = df.f.Write(d)
	return
}

func (df *myDataFile) Close() error {
	if df.f != nil {
		return df.f.Close()
	}
	return nil
}

func NewDataFile(path string) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	df := &myDataFile{f: f}
	return df, nil
}
