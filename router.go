package main

import (
	"github.com/gin-gonic/gin"
	"gopractice/biz/handler"
)

func customizeRegister(r *gin.Engine) {
	//注册数据库操作前缀
	db := r.Group("/db") //操作数据库，增删改查

	registerDB(db)
}

func registerDB(db *gin.RouterGroup) {
	operateRouter := db.Group("/operate")
	operateWithConfRouter := db.Group("/operate_with_condition")
	exceptionRouter := db.Group("/exception")
	testRouter := db.Group("/test")

	registerOperateRouter(operateRouter)
	registerOperateWithConfRouter(operateWithConfRouter)
	registerExceptionRouter(exceptionRouter)
	registerTestRouter(testRouter)
}

func registerOperateRouter(g *gin.RouterGroup) {

}

func registerOperateWithConfRouter(g *gin.RouterGroup) {

}

func registerExceptionRouter(g *gin.RouterGroup) {

}

func registerTestRouter(g *gin.RouterGroup) {
	g.GET("create_table", handler.CreateTable)
	g.GET("create_data", handler.CreateData)
}
