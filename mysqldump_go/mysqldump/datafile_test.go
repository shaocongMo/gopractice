package mysqldump

import (
	"testing"
	"os"
)

func TestDataFile(t *testing.T) {
	t.Run("testNewDataFile", testNewDataFile)
}

func testNewDataFile(t *testing.T) {
	path := "./db.sql"
	df, err := NewDataFile(path)
	defer func() {
			if df != nil{
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