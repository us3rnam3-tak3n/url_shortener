package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"url_shortening/database_conn"
	"url_shortening/routers"
)

func init() {
	//open a db connection
	database_conn.Conn()
}

func main() {
	routers.Initroutes()
}