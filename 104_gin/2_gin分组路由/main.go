package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	rg := r.Group("/user")
	rg.GET("/login", one)
	rg.GET("two", two)
	r.Run()
}
func one(c *gin.Context) {
	c.String(200, "login")
}
func two(c *gin.Context) {
	c.String(200, "two")
}
