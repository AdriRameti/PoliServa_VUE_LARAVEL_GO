package main

import (
	"fmt"
	"poliserva/Config"
	"poliserva/Routes"
	"time"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	count := 0
	if err != nil {
		fmt.Println("Status:", err);
		for{
			if err == nil{
				fmt.Println("Status is NIL")
				break
			}
			fmt.Println("Status is not NIL")
			time.Sleep(time.Second)
			count++
			if count > 180{
				fmt.Printf("***************")
				fmt.Println("DB Connection failure")
				panic(err)
			}
			Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
		}
	}
	fmt.Println("DB Connection successful")
	defer Config.DB.Close()

	r := Routes.SetupRouter()

	r.Run(":8003") // listen and serve on 0.0.0.0:8003 (for windows "localhost:8003")
}

