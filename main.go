package main

import (
	"github.com/avtara-kw/social-media-api/app/config"
	users2 "github.com/avtara-kw/social-media-api/businesses/users"
	users3 "github.com/avtara-kw/social-media-api/controllers/users"
	"github.com/avtara-kw/social-media-api/repositories/databases/users"
	"github.com/avtara-kw/social-media-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", utils.PasswordRule)
	}

	usersRepo := users.NewRepoMySQL(db)
	usersService := users2.NewUserService(usersRepo)
	usersController := users3.NewUserController(usersService)

	port := ":" + os.Getenv("PORT")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/users", usersController.Registration)

	log.Println("server running at port ", port)
	router.Run(port)
}
