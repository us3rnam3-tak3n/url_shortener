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

	//var todos []todo
	//db.Find(&todos)
	//if len(todos) <= 0 {
	//	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	//	return
	//}
	//
	//comp:=false
	//
	//for _, item := range todos {
	//	if item.Long == ss {
	//		comp = true
	//		break
	//	}
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
	var cou int=0
	db.Model(&todo{}).Where("`long` LIKE ?", todof.Long).Count(&cou)
	if cou >= 1{

		db.Where("`long` = ?", c.PostForm("long")).First(&todof)
		c.HTML(http.StatusOK, "short.html", gin.H{
			"long" : todof.Long,
			"sho": todof.Short,
		})
		return
	}
	db.Save(&todof)
	todof.Short = longtoshort(todof.ID)
	db.Save(&todof)
	c.HTML(http.StatusOK,"short.html",gin.H{
		"sho" : todof.Short,
		"long": todof.Long,
	})
}
func home(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",gin.H{})
}
func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", home)
	r.POST("/action", createTodo)
	r.Run(":8084")

}