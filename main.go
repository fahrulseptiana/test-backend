package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "test-backend/docs"
	"test-backend/internal/auth"
	"test-backend/internal/user"
)

// @title           User API
// @version         1.0
// @description     Simple user API with Gin and Swagger

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)
	jwtKey := []byte("secret")
	authHandler := auth.NewHandler(service, jwtKey)

	r := gin.Default()

	// Swagger docs endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	authorized := r.Group("/")
	authorized.Use(auth.JWTMiddleware(jwtKey))
	{
		authorized.GET("/users", handler.GetUsers)
		authorized.GET("/users/:id", handler.GetUser)
		authorized.POST("/users", handler.CreateUser)
		authorized.PUT("/users/:id", handler.UpdateUser)
		authorized.DELETE("/users/:id", handler.DeleteUser)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
