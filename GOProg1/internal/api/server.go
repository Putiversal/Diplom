package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func printArrayRand(arr []string) {
	for {
		i := len(arr)
		for i > 0 {
			fmt.Printf(arr[i])
			i = i - 1
		}
	}
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/myTemplate", func(c *gin.Context) {
		c.HTML(http.StatusOK, "myTemplate.tmpl", gin.H{
			"title": "MyTemplate website", "array": []string{"Колбаса", "Сыр", "Хлеб"},
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
