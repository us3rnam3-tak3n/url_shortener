package controller

import (
	"github.com/gin-gonic/gin"
	"url_shortening/service"
)

//Calls function to load HTML basic homepage.
func Home(c *gin.Context){
	service.Home(c)
}

func CreateTodo(c *gin.Context) {
	service.CreateTodo(c)
}
