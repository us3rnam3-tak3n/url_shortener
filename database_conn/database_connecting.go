package database_conn

import(
	"github.com/jinzhu/gorm"
	"url_shortening/Model"
)

var Db *gorm.DB

func Conn(){
	//open a db connection
	var err error
	Db, err = gorm.Open("mysql", "root:gaurav@/url_shortener?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&Model.Todo{})
}
