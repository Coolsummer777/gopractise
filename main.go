package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopractice/biz/dal"
)

func main() {

	r := gin.Default()
	initService(r)
	defer r.Run()

}

func initService(r *gin.Engine) {
	initWeb(r)
	dal.InitMysql()
}
