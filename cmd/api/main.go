package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/handlers"
	"github.com/imabg/authn/store"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	conn, err := store.NewPostgresStore(os.Getenv("POSTGRES_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Authn is up and running",
		})
	})
	// source store
	sourceStore := store.NewSourceStore(conn)
	sourceHandler := handlers.NewSourceHandler(sourceStore)
	//user store
	userStore := store.NewUserStore(conn)
	userHandler := handlers.NewUserHandler(userStore, sourceStore)
	//login store
	credStore := store.NewCredentialStore(conn)
	loginStore := store.NewLoginStore(conn)
	loginHandler := handlers.NewLoginHandler(loginStore, credStore)

	base := router.Group("/api/v1")
	//source routes
	sourceRoutes := base.Group("/sources")
	sourceRoutes.POST("/create", sourceHandler.Create)
	sourceRoutes.GET("/:id", sourceHandler.GetByID)

	//users routes
	userRoutes := base.Group("/users")
	userRoutes.POST("/email", userHandler.CreateViaEmail)
	userRoutes.POST("/phone", userHandler.CreateViaPhone)

	//auth routes
	loginRoutes := base.Group("/auth")
	loginRoutes.POST("/login/email", loginHandler.LoginViaEmail)

	//private routes
	//privateRoutes := base.Group("/").Use(middlewares.AuthMiddleware())

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
