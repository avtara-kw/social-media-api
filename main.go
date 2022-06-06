package main

import (
	"github.com/avtara-kw/social-media-api/app/config"
	"github.com/avtara-kw/social-media-api/app/middleware"
	photos2 "github.com/avtara-kw/social-media-api/businesses/photos"
	users2 "github.com/avtara-kw/social-media-api/businesses/users"
	photos3 "github.com/avtara-kw/social-media-api/controllers/photos"
	users3 "github.com/avtara-kw/social-media-api/controllers/users"
	"github.com/avtara-kw/social-media-api/repositories/databases/photos"
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

	usersRepo := users.NewRepoPostgresql(db)
	usersService := users2.NewUserService(usersRepo)
	usersController := users3.NewUserController(usersService)

	photoRepo := photos.NewRepoPostgresql(db)
	photoService := photos2.NewPhotosService(photoRepo)
	photoController := photos3.NewPhotosController(photoService)

	port := ":" + os.Getenv("PORT")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRoute := router.Group("/users")
	{
		userRoute.POST("/", usersController.Registration)
		userRoute.POST("/login", usersController.Login)
		userRoute.Use(middleware.Auth())
		userRoute.DELETE("/", usersController.Delete)
		userRoute.PUT("/", usersController.Update)
	}

	photoRoute := router.Group("/photos")
	{
		photoRoute.Use(middleware.Auth())
		photoRoute.POST("/", photoController.Post)
		photoRoute.GET("/", photoController.GetAll)
		photoRoute.DELETE("/:id", photoController.Delete)
		photoRoute.PUT("/:id", photoController.Update)
	}

	log.Println("server running at port ", port)
	router.Run(port)
}
