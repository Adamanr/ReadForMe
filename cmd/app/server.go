package main

import (
	v1 "MyPay/internal/controller/http/v1"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	router := gin.New()
	router.LoadHTMLGlob("./pkg/assets/templates/*/*.*")
	router.Static("/static", "./pkg/assets/static")

	router.GET("/", v1.MainPage)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Сайт не смог запустить на :8080 по причине: %s", err)
	}
}
