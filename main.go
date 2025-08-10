package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "test-backend/docs"
	"test-backend/internal/user"
)

// @title           User API
// @version         1.0
// @description     Simple user API with Gin and Swagger

// @host      localhost:8080
// @BasePath  /
func main() {
	repo := user.NewInMemoryRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r := gin.Default()

	// Swagger docs endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routes
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
