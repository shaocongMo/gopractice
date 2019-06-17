package mysqldump

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config interface {
	DataSourceName() string
	TarDataSourceName() string
	GetTableNameRe() string
	GetTableNameExceptRe() string
	GetFilePath() string
}

type config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	TableNameRe string
	TableNameExceptRe string
	Charset     string
	FilePath    string
	TarHost     string
	TarPort     string
	TarUsername string
	TarPassword string
	TarDatabase string
}

func NewConfig(configFilePath string) (Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open config file error : %s", err))
	}
	defer file.Close()
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read config file error : %s", err))
	}
	var config config
	json.Unmarshal(fileData, &config)
	return &config, nil
}

// "root:root@tcp(127.0.0.1:port)/go_todo?charset=utf8"
func (c *config) DataSourceName() string {
	return c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?charset=" + c.Charset
}

func (c *config) TarDataSourceName() string {
	return c.TarUsername + ":" + c.TarPassword + "@tcp(" + c.TarHost + ":" + c.TarHost + ")/" + c.TarDatabase + "?charset=" + c.Charset
}

func (c *config) GetTableNameRe() string {
	return c.TableNameRe
}

func (c *config) GetTableNameExceptRe() string{
	return c.TableNameExceptRe
}

func (c *config) GetFilePath() string {
	return c.FilePath
}
