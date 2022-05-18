package main

import (

	"poliserva/Config"
	"poliserva/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
	
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err);
	}

	defer Config.DB.Close()

	r := Routes.SetupRouter()

	r.Run(":8003") // listen and serve on 0.0.0.0:8003 (for windows "localhost:8003")
}

