package server

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sugandasu/ruru/jongi"
	"github.com/sugandasu/ruru/nibirudb"

	"github.com/sugandasu/go-boilerplate/config"
	_ "github.com/sugandasu/go-boilerplate/docs"
	"github.com/sugandasu/go-boilerplate/internal/app/handler"
	"github.com/sugandasu/go-boilerplate/internal/app/repository"
	"github.com/sugandasu/go-boilerplate/internal/app/service/auth"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunRestServer(cfg *config.Config) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     cfg.Rest.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     cfg.Rest.AllowedHeaders,
		AllowCredentials: true,
		MaxAge:           300,
	}))

	db := nibirudb.NewDatabaseConnection(&cfg.DB)

	userRepo := repository.NewUserRepository(db)

	authService := auth.NewAuthService(cfg, userRepo)

	authHandler := handler.NewAuthHandler(authService)
	v1 := e.Group("/v1", jongi.EchoErrorMiddleware)
	authHandler.RegisterRoutes(v1)
	v1.Use(jongi.EchoAuthMiddleware(cfg.Jwt.SecretKey))

	// Swagger
	e.GET("/sewage/*", echoSwagger.WrapHandler)

	// Serve static file
	e.Static("/static", "static")

	addr := fmt.Sprintf("%s:%d", cfg.Rest.Host, cfg.Rest.Port)
	log.Printf("Running REST API on http://%s", addr)
	if err := e.Start(addr); err != nil {
		log.Println(err)
	}
}
