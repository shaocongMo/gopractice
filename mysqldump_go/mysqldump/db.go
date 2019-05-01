package mysqldump

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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
	tarEngine   *xorm.Engine
	dump_tables []*core.Table
	data        string
	dataAddLock sync.Mutex
	datafile    DataFile
}

func NewDB(config Config) (Db, error) {
	db := mysqlDb{config: config}
	err := db.Connect()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("db connect error : %s", err))
	}
	db.InitDumpTables()
	if db.config.GetFilePath() != "" {
		datafile, err := NewDataFile(db.config.GetFilePath())
		db.datafile = datafile
		if err != nil {
			return nil, errors.New(fmt.Sprintf("db datafile create error : %s", err))
		}
	}
	return &db, nil
}

func (db *mysqlDb) Connect() (err error) {
	db.engine, err = xorm.NewEngine("mysql", db.config.DataSourceName())
	if err != nil {
		return errors.New(fmt.Sprintf("db connect error: %s", err))
	}
	db.tarEngine, err = xorm.NewEngine("mysql", db.config.TarDataSourceName())
	if err != nil {
		return errors.New(fmt.Sprintf("db connect tar error: %s", err))
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
	db.dump_tables = tables[:table_index]
	return nil
}

func (db *mysqlDb) Dump() error {
	table_num := len(db.dump_tables)
	var wg sync.WaitGroup
	wg.Add(table_num)
	var dialect core.Dialect = db.engine.Dialect()
	for _, table := range db.dump_tables {
		time.Sleep(time.Millisecond * 100)
		go func(t *core.Table) {
			defer wg.Done()
			start := time.Now()
			fmt.Printf("[%v] %s Start.\n", start.Format(time.RFC822), t.Name)
			dropTableSql := "IF EXISTS DROP TABLE `" + t.Name + "`;"
			createTableSql := dialect.CreateTableSql(t, "", t.StoreEngine, "") + ";"

			cols := t.ColumnsSeq()
			colNames := dialect.Quote(strings.Join(cols, dialect.Quote(", ")))

			rows, err := db.engine.DB().Query("SELECT " + colNames + " FROM " + t.Name)

			insertSql := ""
			if err != nil {
				fmt.Println("table get rows error: ", err)
			} else {
				insertSqlHead := "INSERT INTO " + t.Name + " (" + colNames + ") VALUES "
				insertSqlTail := ""
				for rows.Next() {
					dest := make([]interface{}, len(cols))
					err := rows.ScanSlice(&dest)
					if err != nil {
						fmt.Println("table ScanSlice rows error: ", err)
						break
					}
					for i, d := range dest {
						col := t.GetColumn(cols[i])
						if col == nil {
							break
						}
						if i == 0 {
							insertSqlTail += "( " + colValToSting(d, col)
						} else {
							insertSqlTail += colValToSting(d, col)
						}
						if i < len(cols)-1 {
							insertSqlTail += ", "
						}
						if i == len(cols)-1 {
							insertSqlTail += "),"
						}

					}
				}
				rows.Close()
				if strings.EqualFold(insertSqlTail, "") {
					insertSql = ""
				} else {
					insertSql = insertSqlHead + insertSqlTail[:len(insertSqlTail)-1] + ";"
				}
			}
			session := db.tarEngine.Prepare()
			session.Exec(dropTableSql)
			session.Exec(createTableSql)
			session.Exec(insertSql)
			session.Commit()
			session.Close()

			// writeSql := dropTableSql + "\n" + createTableSql + "\n" + insertSql + "\n\n"
			// db.datafile.Write([]byte(writeSql))

			end := time.Now()
			fmt.Printf("[%s] %s Finsh speed %d seconds.\n", end.Format(time.RFC822), t.Name, end.Unix()-start.Unix())

			// db.dataAddLock.Lock()
			// fmt.Println(t.Name, end.Unix() -start.Unix())
			// db.data = db.data + writeSql
			// db.dataAddLock.Unlock()
		}(table)
	}
	wg.Wait()
	// db.dataAddLock.Lock()
	// db.datafile.Write([]byte(db.data))
	// db.dataAddLock.Unlock()
	return nil
}

func colValToSting(val interface{}, col *core.Column) string {
	if val == nil {
		return "NULL"
	} else if col.SQLType.IsText() || col.SQLType.IsTime() {
		var v = fmt.Sprintf("%s", val)
		if strings.HasSuffix(v, " +0000 UTC") {
			return fmt.Sprintf("'%s'", v[0:len(v)-len(" +0000 UTC")])
		} else {
			return "'" + strings.Replace(v, "'", "''", -1) + "'"
		}
	} else if col.SQLType.IsBlob() {
		if reflect.TypeOf(val).Kind() == reflect.Slice {
			return fmt.Sprintf("%s", fmt.Sprintf("0x%x", val.([]byte)))
		} else if reflect.TypeOf(val).Kind() == reflect.String {
			return fmt.Sprintf("'%s'", val.(string))
		}
	} else if col.SQLType.IsNumeric() {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Slice:
			if col.SQLType.Name == core.Bool {
				return fmt.Sprintf("%v", strconv.FormatBool(val.([]byte)[0] != byte('0')))
			} else {
				return fmt.Sprintf("%s", string(val.([]byte)))
			}
		case reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int:
			if col.SQLType.Name == core.Bool {
				return fmt.Sprintf("%v", strconv.FormatBool(reflect.ValueOf(val).Int() > 0))
			} else {
				return fmt.Sprintf("%v", val)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if col.SQLType.Name == core.Bool {
				return fmt.Sprintf("%v", strconv.FormatBool(reflect.ValueOf(val).Uint() > 0))
			} else {
				return fmt.Sprintf("%v", val)
			}
		default:
			return fmt.Sprintf(", %v", val)
		}
	} else {
		s := fmt.Sprintf("%v", val)
		if strings.Contains(s, ":") || strings.Contains(s, "-") {
			if strings.HasSuffix(s, " +0000 UTC") {
				return fmt.Sprintf("'%s'", s[0:len(s)-len(" +0000 UTC")])
			} else {
				return fmt.Sprintf("'%s'", s)
			}
		} else {
			return fmt.Sprintf("%s", s)
		}
	}
	return "NULL"
}
