package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func startPageHandler(c *gin.Context) {
	//c.String(http.StatusOK, "Hello !")

	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"products": Products,
	})
}

func impressumPageHandler(c *gin.Context) {
	//c.String(http.StatusOK, "Hello !")

	c.HTML(http.StatusOK, "impressum.tmpl.html", gin.H{
		"products": Products,
	})
}

func main() {
	//log.Printf("HW!")
	//log.Printf("The type is: %T \n", Products)
	//log.Printf("values: %+v", Products)
	initRouter()
}
