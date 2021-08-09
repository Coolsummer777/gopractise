package main

import "github.com/gin-gonic/gin"

func CustomizeRegister(r *gin.Engine){


	_ = r.Group("/")
}