package v3

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	woffset int64
	roffset int64
	rcond   *sync.Cond
	dataLen uint32
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.roffset)
		if atomic.CompareAndSwapInt64(&df.roffset, offset, (offset + int64(df.dataLen))) {
			break
		}

	}

	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				fmt.Println("wait data")
				df.rcond.Wait()
				fmt.Println("signal")
				continue
			}
			return
		}
		d = bytes
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.woffset)
		if atomic.CompareAndSwapInt64(&df.woffset, offset, (offset + int64(df.dataLen))) {
			break
		}
	}

	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer func() {
		df.fmutex.Unlock()
	}()
	_, err = df.f.Write(bytes)
	df.rcond.Signal()
	return
}

func (df *myDataFile) RSN() int64 {
	return atomic.LoadInt64(&df.roffset) / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	return atomic.LoadInt64(&df.woffset) / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return uint32(df.dataLen)
}

func (df *myDataFile) Close() error {
	if df.f != nil {
		return df.f.Close()
	}
	return nil
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}
