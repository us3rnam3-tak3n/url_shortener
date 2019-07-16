package routers

import(
	"github.com/gin-gonic/gin"
	"url_shortening/controller"
)

//Starts server using Gin, and to call required functions
func Initroutes(){
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", controller.Home)
	r.POST("/action", controller.CreateTodo)
	r.Run(":8084")
}
