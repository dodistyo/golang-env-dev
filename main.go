package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)
func main() {
    fmt.Println("halo dunia") //change this string value to test auto reload
    r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
