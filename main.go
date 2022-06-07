package main

import (
	"github.com/avtara-kw/social-media-api/app/config"
	"github.com/avtara-kw/social-media-api/app/middleware"
	comments2 "github.com/avtara-kw/social-media-api/businesses/comments"
	photos2 "github.com/avtara-kw/social-media-api/businesses/photos"
	social_medias2 "github.com/avtara-kw/social-media-api/businesses/social_medias"
	users2 "github.com/avtara-kw/social-media-api/businesses/users"
	comments3 "github.com/avtara-kw/social-media-api/controllers/comments"
	photos3 "github.com/avtara-kw/social-media-api/controllers/photos"
	social_medias3 "github.com/avtara-kw/social-media-api/controllers/social_medias"
	users3 "github.com/avtara-kw/social-media-api/controllers/users"
	"github.com/avtara-kw/social-media-api/repositories/databases/comments"
	"github.com/avtara-kw/social-media-api/repositories/databases/photos"
	"github.com/avtara-kw/social-media-api/repositories/databases/social_medias"
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

	socialMediaRepo := social_medias.NewRepoPostgresql(db)
	socialMediaService := social_medias2.NewSocialMediasService(socialMediaRepo)
	socialMediaController := social_medias3.NewSocialMediasController(socialMediaService)

	commentRepo := comments.NewRepoPostgresql(db)
	commentService := comments2.NewCommentsService(commentRepo)
	commentController := comments3.NewCommentsController(commentService)

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

	socialMediaRoute := router.Group("/socialmedias")
	{
		socialMediaRoute.Use(middleware.Auth())
		socialMediaRoute.POST("/", socialMediaController.Post)
		socialMediaRoute.GET("/", socialMediaController.GetAll)
		socialMediaRoute.DELETE("/:id", socialMediaController.Delete)
		socialMediaRoute.PUT("/:id", socialMediaController.Update)
	}

	commentRoute := router.Group("/comments")
	{
		commentRoute.Use(middleware.Auth())
		commentRoute.POST("/", commentController.Post)
		commentRoute.GET("/", commentController.GetAll)
		commentRoute.DELETE("/:id", commentController.Delete)
		commentRoute.PUT("/:id", commentController.Update)
	}

	log.Println("server running at port ", port)
	router.Run(port)
}
