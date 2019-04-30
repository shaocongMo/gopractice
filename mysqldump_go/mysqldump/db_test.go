package mysqldump

import (
	"fmt"
	"testing"
)

//  go test gopractice/mysqldump_go/mysqldump -run TestDb -v
func TestDb(t *testing.T) {
	// t.Run("newDb", testNewDb)
	// t.Run("dbConnect", testDbConnect)
	// t.Run("initDumpTables", testDbInitDumpTables)
	t.Run("testDbDumpTablesToFile", testDbDumpTablesToFile)
}

func testNewDb(t *testing.T) {
	config, err := NewConfig("./test_config.json")
	if err != nil {
		t.Fatalf("new config err %s", err)
	}
	db, err := NewDB(config)
	if err != nil {
		t.Fatalf("new db error: %s", err)
	}
	fmt.Println(db)
}

func testDbConnect(t *testing.T) {
	config, err := NewConfig("./test_config.json")
	if err != nil {
		t.Fatalf("new config err %s", err)
	}
	db, err := NewDB(config)
	if err != nil {
		t.Fatalf("new db error: %s", err)
	}
	err1 := db.Connect()
	if err1 != nil {
		t.Fatalf("db connect error: %s", err)
	}
}

func testDbInitDumpTables(t *testing.T) {
	config, _ := NewConfig("./test_config.json")
	db, _ := NewDB(config)
	err := db.InitDumpTables()
	if err != nil {
		t.Fatalf("db init dump tables error: %s", err)
	}
}

func testDbDumpTablesToFile(t *testing.T) {
	config, _ := NewConfig("./test_config.json")
	db, _ := NewDB(config)
	err := db.Dump()
	if err != nil {
		t.Fatalf("db dump table error: %s", err)
	}
}
