package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"fastapi-digital-library/internal/infrastructure/memory"
	httpInt "fastapi-digital-library/internal/interface/http"
	"fastapi-digital-library/internal/usecase"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(httpInt.LoggingMiddleware())

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(cfg))

	repo := memory.NewBookRepository()
	service := usecase.NewBookService(repo)
	queue := httpInt.NewTaskQueue(100)

	h := httpInt.NewBookHandler(service, queue)
	h.Register(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
