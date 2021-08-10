package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopractice/biz/handler"
	"gopractice/consts"
	_ "gopractice/consts"
)

var (
	DB_GormTest *gorm.DB
	DB_Db1 *gorm.DB
	DB_TestDatabase *gorm.DB
)

func initWeb(r *gin.Engine) {

	//注册路由
	customizeRegister(r)

	//测试框架与网络基础情况
	r.GET("/ping",handler.Ping)
}


func initDB(){
	var err error
	DB_GormTest,err = gorm.Open(consts.MYSQL,consts.DSNPREFIX+consts.DBNAME_GORM_TEST+consts.DB_CONF)
	if err != nil {
		panic(err.Error())
	}

	DB_TestDatabase,err = gorm.Open(consts.MYSQL,consts.DSNPREFIX+consts.DBNAME_TEST_DATABASE+consts.DB_CONF)
	if err != nil {
		panic(err.Error())
	}

	DB_Db1,err = gorm.Open(consts.MYSQL,consts.DSNPREFIX+consts.DBNAME_DB1+consts.DB_CONF)
	if err != nil {
		panic(err.Error())
	}

}