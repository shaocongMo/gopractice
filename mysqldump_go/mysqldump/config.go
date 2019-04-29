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
}

type config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	TableNameRe string
	Charset     string
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

// "root:root@tcp(127.0.0.1)/go_todo?charset=utf8"
func (c *config) DataSourceName() string {
	return c.Host + ":" + c.Username + "@tcp(" + c.Host + ")/" + c.Database + "?charset=" + c.Charset
}
