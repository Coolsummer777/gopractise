package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopractice/biz/dal"
	"gopractice/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func ListAll(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Find(c *gin.Context) {

}

func CreateTable(c *gin.Context) {

	if !dal.DB_Db1.HasTable(&models.UserDbModel{}) {
		//dal.DB_Db1.Set("gorm:table_options","ENGINE=InnoDB").CreateTable(&models.UserDbModel{})
		dal.DB_Db1.AutoMigrate(&models.UserDbModel{})
		c.JSON(http.StatusOK, gin.H{
			"message": "create new table",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "table exists",
		})
	}

}

func CreateData(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	num, err := strconv.Atoi(c.Query("num"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i := 0; i < num; i++ {
			user := models.UserDbModel{}
			user.Age = rand.Intn(100)
			if user.Age%2 == 1 {
				user.Gender = "male"
			} else {
				user.Gender = "female"
			}
			for j := 0; j < 7; j++ {
				user.Email = user.Email + string(rand.Intn(100))
			}
			user.Email += "@qq.com"
			dal.DB_Db1.Table("user_list_test").Debug().Create(&user)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "add data into db",
			"num":     num,
		})
	}

}
