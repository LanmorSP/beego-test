package main

import (
	_ "test-api/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}
	beego.BConfig.CopyRequestBody = true
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=test password=test1234 dbname=test host=localhost port=5432 sslmode=disable",
	)
	beego.Run(":8080")
}
