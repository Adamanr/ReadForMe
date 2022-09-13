package v1

import "github.com/gin-gonic/gin"

func MainPage(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"title": "Главная",
	})
}
