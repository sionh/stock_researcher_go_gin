package controller

import (
	//strconv "strconv"

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
