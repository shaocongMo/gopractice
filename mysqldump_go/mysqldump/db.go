package mysqldump

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Db interface {
	Connect() error
}

type db struct {
	config       Config
	db_connecter *xorm.Engine
}

func NewDB(config Config) (Db, error) {
	db := db{config: config}
	err := db.Connect()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("db connect error : %s", err))
	}
	return &db, nil
}

func (db *db) Connect() (err error) {
	db.db_connecter, err = xorm.NewEngine("mysql", db.config.DataSourceName())
	if err != nil {
		return errors.New(fmt.Sprintf("db connect error: %s", err))
	}
	return nil
}
