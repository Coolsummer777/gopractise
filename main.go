package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopractice/init"
)

func main() {

	r := gin.Default()
	init.InitWeb(r)
	defer r.Run()

}
