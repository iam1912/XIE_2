package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iam1912/XIE_2/control"
)

func main() {
	r := gin.Default()
	r.Static("/view", "./view")
	r.LoadHTMLGlob("./view/*")
	r.GET("/login", LoginHandler)
	r.POST("/login", LoginerHandler)

	r.GET("/index", ViewHandler)

	r.GET("/edit", EditHandler)
	r.POST("/edit", PostHandler)

	r.Run(":8080")
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginerHandler(c *gin.Context) {
	control.LoginerHandler(c)
}

func ViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func EditHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.html", nil)
}

func PostHandler(c *gin.Context) {
	control.PostHandler(c)
}
