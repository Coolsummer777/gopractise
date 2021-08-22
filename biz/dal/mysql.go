package dal

import (
	"github.com/jinzhu/gorm"
	"gopractice/consts"
)

var (
	DB_GormTest     *gorm.DB
	DB_Db1          *gorm.DB
	DB_TestDatabase *gorm.DB
)

func InitMysql() {
	var err error
	DB_GormTest, err = gorm.Open(consts.MYSQL, consts.DSNPREFIX+consts.DBNAME_GORM_TEST+consts.DB_CONF)
	if err != nil {
		panic(err.Error())
	}

	DB_TestDatabase, err = gorm.Open(consts.MYSQL, consts.DSNPREFIX+consts.DBNAME_TEST_DATABASE+consts.DB_CONF)
	if err != nil {
		panic(err.Error())
	}

	DB_Db1, err = gorm.Open(consts.MYSQL, consts.DSNPREFIX+consts.DBNAME_DB1+consts.DB_CONF)
	if err != nil {
		panic(err.Error())
	}

}
