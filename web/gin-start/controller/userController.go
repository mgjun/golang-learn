package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

var Users = []User{{Id: "123", Name: "张三"}, {Id: "456", Name: "李四"}}

func GetUsers(r *gin.Engine, path string) {
	r.GET(path, func(c *gin.Context) {
		c.JSON(http.StatusOK, Users)
	})
}

func CreateUser(r *gin.Engine, path string) {
	r.POST(path, func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Users := append(Users, user)
		c.JSON(http.StatusOK, Users)
		return
	})
}

func ModifyUserId(r *gin.Engine, path string) {
	r.PUT(path, func(c *gin.Context) {
		var modifyUser User
		if err := c.ShouldBindJSON(&modifyUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := c.Param("id")
		for i := 0; i < len(Users); i++ {
			if Users[i].Id == id {
				Users[i].Name = modifyUser.Name
			}
		}
		c.JSON(http.StatusOK, Users)
		return
	})
}

func DeleteUserById(r *gin.Engine, path string) {
	r.DELETE(path, func(c *gin.Context) {
		id := c.Param("id")
		var newUsers []User
		for i := 0; i < len(Users); i++ {
			if Users[i].Id == id {
				continue
			}
			newUsers = append(newUsers, Users[i])
		}
		c.JSON(http.StatusOK, newUsers)
		return
	})
}
