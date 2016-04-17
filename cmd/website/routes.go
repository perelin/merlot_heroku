package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

var isRunningOnHeroku = false
var localFolderPrefix string

func setContext() {
	if os.Getenv("PORT") == "" {
		isRunningOnHeroku = false
		localFolderPrefix = "../../"
		log.Printf("Running outside of Heroku context")
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Printf("$PORT is not set, using default port: %s", port)
	}
	return port
}

func initRouter() {

	setContext()

	router.Use(gin.Logger())

	router.LoadHTMLGlob(localFolderPrefix + "templates/*")

	/*router.LoadHTMLFiles("templates/twitterlogin.tmpl.html",
	"templates/twittercallback.tmpl.html",
	"templates/h5bp/twissr.html")*/
	//router.LoadHTMLGlob("templates/h5bp/*.html")
	router.Static("/static", localFolderPrefix+"static")
	router.Static("/js", localFolderPrefix+"static/js")
	router.Static("/css", localFolderPrefix+"static/css")
	router.Static("/img", localFolderPrefix+"static/img")
	router.Static("/fonts", localFolderPrefix+"static/fonts")
	router.StaticFile("/favicon.ico", localFolderPrefix+"static/root/favicon.ico")
	router.StaticFile("/robots.txt", localFolderPrefix+"static/root/robots.txt")
	router.StaticFile("/apple-touch-icon.png", localFolderPrefix+"static/root/apple-touch-icon.png")

	/*
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		})*/

	router.GET("/", startPageHandler)

	router.GET("/impressum", impressumPageHandler)

	router.GET("/datenschutzerklaerung", datenschutzerklaerungPageHandler)
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
