package mysqldump

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type Db interface {
	Connect() error
	InitDumpTables() error
	Dump() error
}

type mysqlDb struct {
	config      Config
	engine      *xorm.Engine
	dump_tables []*core.Table
}

func NewDB(config Config) (Db, error) {
	db := mysqlDb{config: config}
	err := db.Connect()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("db connect error : %s", err))
	}
	db.InitDumpTables()
	return &db, nil
}

func (db *mysqlDb) Connect() (err error) {
	db.engine, err = xorm.NewEngine("mysql", db.config.DataSourceName())
	if err != nil {
		return errors.New(fmt.Sprintf("db connect error: %s", err))
	}
	return nil
}

func (db *mysqlDb) InitDumpTables() error {
	tables, err := db.engine.DBMetas()
	if err != nil {
		return err
	}
	var table_index int = 0
	tableNameReStr := db.config.GetTableNameRe()
	tableNameRe := regexp.MustCompile(tableNameReStr)
	for _, table := range tables {
		if strings.EqualFold(tableNameReStr, "") || tableNameRe.MatchString(table.Name) {
			tables[table_index] = table
			table_index += 1
		}
	}
	db.dump_tables = tables[table_index:]
	return nil
}

func (db *mysqlDb) Dump() error {
	return nil
	// return db.engine.DumpTablesToFile(db.dump_tables, db.config.GetFilePath())
}
