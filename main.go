package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	about = "about.gohtml"
	last  = "last.gohtml"
)

func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	engine.GET(about, func(context *gin.Context) {
		context.HTML(200, about, nil)
	})

	engine.GET("", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, about)
	})

	engine.GET(last, func(context *gin.Context) {
		rewards := []struct {
			From        string
			To          string
			Description string
		}{
			{From: "Alex", Description: "Thanks for helping me out the other day.", To: "Felix"},
			{From: "Felix", Description: "It is fun working with you.", To: "Tim"},
			{From: "Timo", Description: "Thanks for finding my bug.", To: "Alex"},
		}
		context.HTML(200, last, gin.H{
			"Rewards": rewards,
		})
	})
	_ = engine.Run(":8080")
}
