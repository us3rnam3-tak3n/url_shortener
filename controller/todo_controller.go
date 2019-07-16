package controller

import (
	"github.com/gin-gonic/gin"
	"url_shortening/service"
)

func Home(c *gin.Context){
	service.Home(c)
}

func CreateTodo(c *gin.Context) {
	service.CreateTodo(c)
}
