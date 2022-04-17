package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "工具",
	})
}

func CheckName(c *gin.Context) {
	allStudentList := map[string]int{"张三": 29, "李四": 30}
	var notInList []string

	commitList := c.DefaultPostForm("name_list", "张三")
	for k, v := range allStudentList {
		if strings.Contains(commitList, k) {
			fmt.Println(k, "在提交的名单中")
		} else {
			notInList = append(notInList, k+strconv.Itoa(v))
			fmt.Println(k, "不在提交的名单中")
		}
	}

	c.HTML(http.StatusOK, "check_name.html", gin.H{
		"NameList": notInList,
	})
}
