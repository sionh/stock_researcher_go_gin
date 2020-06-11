package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"

	entity "../entity"
)

// DB接続する
func open() *gorm.DB {
	DBMS := "sqlite3"
	DBNAME := "db.sqlite3"
	db, err := gorm.Open(DBMS, DBNAME)

	if err != nil {
		panic(err.Error())
	}

	// 詳細なログを表示
	db.LogMode(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Shikiho{})

	fmt.Println("db connected: ", &db)
	return db
}

// SearchShikihos は全shikihoを取得します
func SearchShikihos(param url.Values) []entity.Shikiho {
	shikihos := []entity.Shikiho{}

	db := open()
	// select
	if len(param.Get("year")) != 0 {
		db = db.Where("year = ?", param.Get("year"))
	}
	if len(param.Get("quarter")) != 0 {
		db = db.Where("quarter = ?", param.Get("quarter"))
	}
	if len(param.Get("code")) != 0 {
		db = db.Where("code like ?", "%"+param.Get("code")+"%")
	}
	if len(param.Get("compName")) != 0 {
		db = db.Where("compName like ?", "%"+param.Get("compName")+"%")
	}
	if len(param.Get("feature")) != 0 {
		db = db.Where("feature like ?", "%"+param.Get("feature")+"%")
	}
	if len(param.Get("topic")) != 0 {
		db = db.Where("topic like ?", "%"+param.Get("topic")+"%")
	}
	if len(param.Get("outlook")) != 0 {
		db = db.Where("outlook like ?", "%"+param.Get("outlook")+"%")
	}
	db = db.Order("code asc").Find(&shikihos)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return shikihos
}

// GetShikiho は shikihoの履歴をすべて取得する
func GetShikiho(code int) []entity.Shikiho {
	shikihos := []entity.Shikiho{}

	db := open()
	db = db.Where("code = ?", code).Order("year desc").Order("quarter desc")
	// select
	db.Find(&shikihos)

	defer db.Close()

	return shikihos
}

// InsertShikiho は shikihoにレコードを追加する
func InsertShikiho(registerShikiho *entity.Shikiho) {
	db := open()
	// insert
	db.Create(&registerShikiho)
	defer db.Close()
}

// UpdateStateShikiho は shikihoの指定したレコードの状態を変更する
func UpdateStateShikiho(shikihoID int, shikihoState int) {
	shikiho := []entity.Shikiho{}

	db := open()
	// update
	db.Model(&shikiho).Where("ID = ?", shikihoID).Update("State", shikihoState)
	defer db.Close()
}

// DeleteShikiho は shikihoの指定したレコードを削除する
func DeleteShikiho(shikihoID int) {
	shikiho := []entity.Shikiho{}

	db := open()
	// delete
	db.Delete(&shikiho, shikihoID)
	defer db.Close()
}
