package main

import "github.com/gin-gonic/gin"

func customizeRegister(r *gin.Engine){
	//注册全局前缀
	dbtest := r.Group("/dbtest")//操作数据库，增删改查
	tmpRouter := r.Group("/temp")//用于异常测试，空测试等


	registerDbtest(dbtest)
	registerTempRouter(tmpRouter)
}

func registerDbtest(g *gin.RouterGroup){
	g.GET("/list")
	g.POST("")
	g.DELETE("delete")
	g.PUT("")
}

func registerTempRouter(g *gin.RouterGroup){

	g.GET("")
}