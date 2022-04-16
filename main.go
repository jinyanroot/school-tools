package main

import (
	"github.com/gin-gonic/gin"
	"school-tools/controller"
)

func main() {
	router := gin.Default()
	// 导入所有模板，多级目录结构需要这样写
	router.LoadHTMLGlob("web/templates/*")

	// website分组
	v := router.Group("/")
	{
		v.GET("/index.html", controller.IndexController)
		v.POST("/check_name.html", controller.CheckName)
	}

	err := router.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
