package main

import (
	opensearch "test-api/es"
	_ "test-api/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
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
	logs.Register("opensearch", opensearch.NewOpensarch)
	logs.SetLogger("opensearch", `{"dsn":"https://localhost:9200","index":"go-test-index1"}`)

	beego.Run(":8080")
}
