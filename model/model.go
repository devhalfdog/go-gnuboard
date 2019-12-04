package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    // mssql 드라이버 로드
	_ "github.com/jinzhu/gorm/dialects/mysql"    // mysql 드라이버 로드
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgresql 드라이버 로드
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite 드라이버 로드
	"go-gnuboard/conf"
	"go-gnuboard/controller"
)

var db *gorm.DB

func DBConnect(DBData map[string]string) {
	config := conf.NewConfig()
	var err error

	if DBData["Database"] == "MySQL" {
		// MySQL
		db, err = gorm.Open("mysql", DBData["User"]+":"+DBData["Password"]+"@tcp("+
			DBData["Host"]+":"+DBData["Port"]+")/"+
			DBData["DB"]+"?charset="+config.DBCharset+"&parseTime=True&loc=Local")
	} else if DBData["Database"] == "PostgreSQL" {
		db, err = gorm.Open("postgresql", "host="+DBData["Host"]+" port="+
			DBData["Port"]+" user="+DBData["User"]+" dbname="+DBData["DB"]+
			" password"+DBData["Password"])
	}

	controller.ErrorController(err)

	defer db.Close()
}
