package main

import (
	 "github.com/gin-gonic/gin"
	 "github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	"gopractice/consts"
)


func main()  {
	r := gin.Default()
	db ,err := gorm.Open(consts.MYSQL,consts.DSNDB1)
	if err != nil {
		panic(err.Error())
	}
	r.Run()
	defer db.Close()

}
