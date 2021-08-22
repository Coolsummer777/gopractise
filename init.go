package main

import (
	"github.com/gin-gonic/gin"
	"gopractice/biz/handler"
	_ "gopractice/consts"
)

func initWeb(r *gin.Engine) {

	//注册路由
	customizeRegister(r)

	//测试框架与网络基础情况
	r.GET("/ping", handler.Ping)
}
