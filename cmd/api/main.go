package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/handlers"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/utils"
	"github.com/jmoiron/sqlx"
	"log"
)

type ServerType struct {
	Config      utils.Config
	TokenMaster utils.Maker
	Router      *gin.Engine
	SourceStore *store.SourceStore
	UserStore   *store.UserStore
	LoginStore  *store.LoginStore
}

func main() {
	server, err := initServer()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := store.NewPostgresStore(server.Config.PostgresDBURL)
	if err != nil {
		log.Fatal(err)
	}

	server.Router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Authn is up and running",
		})
	})

	server.initialiseStore(conn)
	server.setupRoutes()

	err = server.Router.Run(":8080")
	if err != nil {
		return
	}
}

func initServer() (*ServerType, error) {
	config, err := utils.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("while loading config %s", err.Error())
	}
	tokenMaster, err := utils.NewPasetoMaker(config.TokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("while creating token %s", err.Error())
	}
	s := gin.Default()
	return &ServerType{
		Config:      config,
		TokenMaster: tokenMaster,
		Router:      s,
	}, nil
}

func (s *ServerType) initialiseStore(conn *sqlx.DB) {
	s.LoginStore = store.NewLoginStore(conn)
	s.UserStore = store.NewUserStore(conn)
	s.SourceStore = store.NewSourceStore(conn)
}

func (s *ServerType) setupRoutes() {
	const base = "/api/v1"
	baseRoutes := s.Router.Group(base)

	//	source routes
	sourceRoutes := baseRoutes.Group("/sources")
	sourceHandler := handlers.NewSourceHandler(s.SourceStore)
	sourceRoutes.POST("/create", sourceHandler.Create)
	sourceRoutes.GET("/:id", sourceHandler.GetByID)

	//	user routes
	userRoutes := baseRoutes.Group("/users")
	userHandler := handlers.NewUserHandler(s.UserStore, s.SourceStore)
	userRoutes.POST("/email", userHandler.CreateViaEmail)
	userRoutes.POST("/phone", userHandler.CreateViaPhone)

	//	auth  routes
	authRoutes := baseRoutes.Group("/auth")
	authHandler := handlers.NewLoginHandler(s.LoginStore, s.TokenMaster, s.Config)
	authRoutes.POST("/login/email", authHandler.LoginViaEmail)

}
