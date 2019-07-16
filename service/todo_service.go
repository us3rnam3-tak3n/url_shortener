package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url_shortening/Model"
	"url_shortening/database_conn"
)

func longtoshort(n uint)  string{
	enc := [62]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","0","1","2","3","4","5","6","7","8","9"}
	var ans string
	for n!=0{
		ans+=enc[n%62]
		n/=62
	}
	return ans
}

func Home(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",gin.H{})
}

func CreateTodo(c *gin.Context) {
	var todof Model.Todo
	todof.Long=c.PostForm("long")
	var cou int=0
	database_conn.Db.Model(&Model.Todo{}).Where("`long` LIKE ?", todof.Long).Count(&cou)
	if cou >= 1{

		database_conn.Db.Where("`long` = ?", c.PostForm("long")).First(&todof)
		c.HTML(http.StatusOK, "short.html", gin.H{
			"long" : todof.Long,
			"sho": todof.Short,
		})
		return
	}
	database_conn.Db.Save(&todof)
	todof.Short = longtoshort(todof.ID)
	database_conn.Db.Save(&todof)
	c.HTML(http.StatusOK,"short.html",gin.H{
		"sho" : todof.Short,
		"long": todof.Long,
	})
}

