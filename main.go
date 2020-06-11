package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"

	controller "./controllers/controller"
)

func main() {
	serve()
}

func serve() {

	router := gin.Default()

	router.Static("/static", "./views/static")
	router.LoadHTMLGlob("views/templates/*.html")

	router.GET("shikihos", controller.SearchShikihos)

	router.Run()
}
