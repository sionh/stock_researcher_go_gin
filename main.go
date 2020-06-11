package main

import (
	"github.com/gin-gonic/gin"
	"html/template"

	_ "github.com/mattn/go-sqlite3"

	common "./common"
	controller "./controllers/controller"
)

func main() {
	serve()
}

func serve() {

	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"GetQuarterFromNum":       common.GetQuarterFromNum,
		"Comma":                   common.Comma,
		"RoundFloatStrIfParsable": common.RoundFloatStrIfParsable,
	})
	router.Static("/static", "./views/static")
	router.LoadHTMLGlob("views/templates/*.html")

	router.GET("shikihos", controller.SearchShikihos)
	router.GET("shikihos/code/:code", controller.DispShikiho)

	router.Run()
}
