package models

import "github.com/jinzhu/gorm"

type UserDbModel struct {
	gorm.Model
	Name     string
	Gender   string
	Age      int
	Email    string
	location string
}

//TableName方法为结构体绑定表名，gorm会通过反射机制获取对应的表名，从而知道在数据库中操作哪张表
//在使用CreateTable或者AutoMigrate方法创建表时，会采用这里所绑定的表名
//如果缺乏TableName方法，gorm在根据struct创建表的时候会按照结构体名称自动生成表名，会使用复数形式，将大写开头的驼峰命名转化为下划线分割的全小写表名，并使用复数形式

//func (UserDbModel) TableName() string {
//	return "user_list_test"
//}
