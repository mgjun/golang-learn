package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	users := []User{{ID: "123", Name: "张三"}, {ID: "456", Name: "李四"}}
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users := append(users, user)
		c.JSON(http.StatusOK, users)
		return
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var modifyUser User
		if err := c.ShouldBindJSON(&modifyUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i := 0; i < len(users); i++ {
			if users[i].ID == id {
				users[i].Name = modifyUser.Name
			}
		}
		c.JSON(http.StatusOK, users)
		return
	})

	r.DELETE("users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var newUsers []User
		for i := 0; i < len(users); i++ {
			if users[i].ID == id {
				continue
			}
			newUsers = append(newUsers, users[i])
		}
		c.JSON(http.StatusOK, newUsers)
		return
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
