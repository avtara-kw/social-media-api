package main

import (
	"github.com/avtara-kw/social-media-api/app/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	port := ":" + os.Getenv("PORT")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("server running at port ", port)
	router.Run(port)
}
