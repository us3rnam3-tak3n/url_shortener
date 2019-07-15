package main

import (
	//	"fmt"
//	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

type(
	todo struct {
		gorm.Model
		Long	string `json:"long"`
		Short	string `json:"short"`
	}
)
func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:gaurav@/url_shortener?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo{})
}

func fetchAllTodo(c *gin.Context) {
	//c.String(http.StatusOK,"dfdfgd")
	c.HTML(http.StatusOK,"index.html",gin.H{})
	//var todos []todo
	//db.Find(&todos)
	//if len(todos) <= 0 {
	//	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	//	return
	//}
	//
	//for _, item := range todos {
	//	completed := false
	//	if item.Completed == 1 {
	//		completed = true
	//	} else {
	//		completed = false
	//	}
	//	//_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	//}
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

func longtoshort(n uint)  string{
	enc := [62]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","0","1","2","3","4","5","6","7","8","9"}
	var ans string
	for n!=0{
		ans+=enc[n%62]
		n/=62
	}
	return ans
}
func createTodo(c *gin.Context) {
	var todof todo
	todof.Long=c.PostForm("long")
	db.Save(&todof)
	todof.Short = longtoshort(todof.ID)
	db.Save(&todof)
	c.Redirect(http.StatusMovedPermanently,"http:localhost:8084/done")
}
func doneTodo(c *gin.Context){
	c.HTML(http.StatusOK,"short.html",gin.H{})
}
func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", fetchAllTodo)
	r.GET("/done", doneTodo)
	r.POST("/action", createTodo)
	r.Run(":8084")

}