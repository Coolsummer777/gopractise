package main

import (
	"github.com/gin-gonic/gin"
	"gopractice/biz/handler"
)

func InitWeb(r *gin.Engine) {

	customizeRegister(r)

	//测试框架与网络基础情况
	r.GET("/ping",handler.Ping)
}
