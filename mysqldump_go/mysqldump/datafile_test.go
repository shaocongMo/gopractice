package mysqldump

import (
	"os"
	"sync"
	"testing"
)

func TestDataFile(t *testing.T) {
	// t.Run("testNewDataFile", testNewDataFile)
	t.Run("testWriteData", testWriteData)
}

func testNewDataFile(t *testing.T) {
	path := "./db.sql"
	df, err := NewDataFile(path)
	defer func() {
		if df != nil {
			df.Close()
		}
		if err := removeFile(path); err != nil {
			t.Errorf("Open file error: %s\n", err)
		}
	}()
	if err != nil {
		t.Fatalf("new datafile error : %s", err)
	}
	t.Logf("datafile: %v", df)
}

func testWriteData(t *testing.T) {
	path := "./db.sql"
	df, err := NewDataFile(path)
	defer func() {
		if df != nil {
			df.Close()
		}
		// if err := removeFile(path); err != nil {
		// 	t.Errorf("Open file error: %s\n", err)
		// }
	}()
	if err != nil {
		t.Fatalf("new datafile error : %s", err)
	}
	max := 100
	var wg sync.WaitGroup
	wg.Add(max)
	for i := 1; i <= max; i++ {
		go func(index int) {
			data := Data{byte(index), 'a', '\n'}
			df.Write(data)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func removeFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return nil
	}
	file.Close()
	return os.Remove(path)
}
