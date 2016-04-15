package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("$PORT is not set, using default port: %s", port)
	}
	return port
}

func initRouter() {
	router.Use(gin.Logger())

	router.LoadHTMLGlob("templates/*")

	/*router.LoadHTMLFiles("templates/twitterlogin.tmpl.html",
	"templates/twittercallback.tmpl.html",
	"templates/h5bp/twissr.html")*/
	//router.LoadHTMLGlob("templates/h5bp/*.html")
	router.Static("/static", "static")
	router.Static("/js", "static/js")
	router.Static("/css", "static/css")
	router.Static("/img", "static/img")
	router.Static("/fonts", "static/fonts")
	router.StaticFile("/favicon.ico", "static/root/favicon.ico")
	router.StaticFile("/robots.txt", "static/root/robots.txt")
	router.StaticFile("/apple-touch-icon.png", "static/root/apple-touch-icon.png")

	/*
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		})*/

	router.GET("/", startPageHandler)
	//router.GET("/tl", twitterLoginHandler)
	//router.GET("/tc", twitterCallbackHandler)
	//router.GET("/tc", twitterCallbackHandler2)
	//router.GET("/hometimeline/:userID", twitterHomeTimelineRSSHandler)

	/*
		router.GET("/tl", func(c *gin.Context) {
			c.HTML(http.StatusOK, "twitterlogin.tmpl.html", nil)
		})*/
	router.Run(":" + getPort())
}
