package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/gohouse/gorose"
	_ "github.com/mattn/go-sqlite3"
)

var dbc orm.Ormer
var dbx *gorose.Engin

// var config = &gorose.Config{
//     Driver: "sqlite3",
//     Dsn:    "data.db",
// }
var err error

func init() {
	// 注册驱动
	_ = orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 注册默认数据库
	_ = orm.RegisterDataBase("default", "sqlite3", "data.db")
	// 注册model
	orm.RegisterModel(
		new(Config),
		new(Article),
		new(Comment),
		new(Category),
		new(Link),
		new(Message),
		new(Notice),
		new(Attachment),
	)
	// 自动建表
	_ = orm.RunSyncdb("default", false, true)

	// gorose
	dbx, err = gorose.Open(&gorose.Config{Driver: "sqlite3", Dsn: "./data.db"})
	if err != nil {
		panic(err.Error())
	}
}

func DB() gorose.IOrm {
	return dbx.NewOrm()
}
