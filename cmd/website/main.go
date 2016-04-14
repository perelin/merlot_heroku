package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func startPageHandler(c *gin.Context) {
    //c.String(http.StatusOK, "Hello !")
    c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
        "title": "Main website",
    })
}

func main() {
	initRouter()
}
