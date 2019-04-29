package mysqldump

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("newConfig", testNewConfig)
	t.Run("dataSourceName", testConfigDataSourceName)
}

func testNewConfig(t *testing.T) {
	NewConfig("./test_config.json")
}

func testConfigDataSourceName(t *testing.T) {
	config, err := NewConfig("./test_config.json")
	if err != nil {
		t.Fatalf("new config err %s", err)
	}
	dataSourceName := config.DataSourceName()
	fmt.Println(dataSourceName)
}
