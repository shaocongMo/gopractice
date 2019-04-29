package mysqldump

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	t.Run("newDb", testNewDb)
	t.Run("dbConnect", testDbConnect)
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
