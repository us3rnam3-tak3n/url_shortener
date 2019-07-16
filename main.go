package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"url_shortening/database_conn"
	"url_shortening/routers"
)

//To initialise database connection
func init() {
	//open a db connection
	database_conn.Conn()
}

//Calls routes to do required functions
func main() {
	routers.Initroutes()
}