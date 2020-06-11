package controller

import (
	strconv "strconv"

	db "../../models/db"
	//entity "../../models/entity"
	"github.com/gin-gonic/gin"
)

// SearchShikihos は 全ての商品情報を取得する
func SearchShikihos(c *gin.Context) {
	param := c.Request.URL.Query()

	shikihos := db.SearchShikihos(param)

	// URLへのアクセスに対してJSONを返す
	c.HTML(200, "index.html", gin.H{
		"shikihos": shikihos,
		"year":     []string{"2020", "2019", "2018", "2017"},
		"quarter":  param,
	})
}

// DispShikiho は 全ての商品情報を取得する
func DispShikiho(c *gin.Context) {
	param := c.Request.URL.Query()
	code, _ := strconv.Atoi(c.Param("code"))
	shikihos := db.GetShikiho(code)
	shikiho := shikihos[0]

	// URLへのアクセスに対してJSONを返す
	c.HTML(200, "code.html", gin.H{
		"shikiho":  shikiho,
		"shikihos": shikihos,
		"year":     []string{"2020", "2019", "2018", "2017"},
		"quarter":  param,
	})
}
