package entities

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:root@/golang?charset=utf8")
	if err != nil {
		panic(err)
	}
	Engine = engine
	if os.Getenv("DEBUG") == "TRUE" {
		Engine.ShowSQL(true)
		Engine.Logger().SetLevel(core.LOG_DEBUG)
	}
}
