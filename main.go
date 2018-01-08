package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//注册接口
	router.GET("/say/:name", SayHello)
	//监听端口
	http.ListenAndServe(":10012", router)
}

func SayHello(c *gin.Context){
	name := c.Param("name")
	if "" == name{
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "param is nil", "status": http.StatusExpectationFailed})
	}
	say := "hello " + name
	c.JSON(http.StatusOK, gin.H{"message":say , "status": http.StatusOK})
}
