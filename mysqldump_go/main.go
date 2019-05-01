package main

import (
	"fmt"
	"gopractice/mysqldump_go/mysqldump"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		configPath := os.Args[1]
		config, err := mysqldump.NewConfig(configPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		db, err := mysqldump.NewDB(config)
		if err != nil {
			fmt.Println(err)
			return
		}
		db.Dump()
		fmt.Println("Finish")
	} else {
		fmt.Printf("please run with an args (config file path) ")
	}
}
