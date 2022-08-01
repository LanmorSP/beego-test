package controllers

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
	opensearch "github.com/opensearch-project/opensearch-go"
	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"
)

// Operations about Users
type TestLogController struct {
	beego.Controller
}

// @Title Send Message Test
// @Description just test logs sender
// @Success 200 {object} models.User
// @router /send [get]
func (c *TestLogController) Logger() {
	logs.Warning("hello")
	logs.Error("asdasd")
	c.Data["json"] = "heelo"

	c.ServeJSON()
}

// @Title Just test send es log
// @Description logs
// @Success 200 {object} models.User
// @router /es [get]
func (c *TestLogController) Es() {

	conn, _ := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "admin",
	})

	// mapping := strings.NewReader(`{
	// 	'settings': {
	// 	  'index': {
	// 		   'number_of_shards': 4
	// 		   }
	// 		 }
	// 	}`)

	// // Create an index with non-default settings.
	// res := opensearchapi.IndicesCreateRequest{
	// 	Index: "go-test-index1",
	// 	Body:  mapping,
	// }
	// fmt.Println("creating index", res)

	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// document := strings.NewReader(fmt.Sprintf(`{
	// 	"time": "%s"
	//     "title": "Moneyball",
	//     "director": "Bennett Miller GG",
	//     "year": "2022"
	// }`, formatted))

	// document := strings.NewReader(fmt.Sprintf(`{
	// 	"time": "2022-08-01 14:50:33",
	//     "title": "Moneyball",
	//     "director": "Bennett Miller GG",
	//     "year": "2022"
	// }`, formatted))
	m, b := map[string]interface{}{
		"time":    formatted,
		"Message": "Hello",
		"user_id": 1,
	}, new(bytes.Buffer)
	log.Println(formatted)

	json.NewEncoder(b).Encode(m)
	req := opensearchapi.IndexRequest{
		Index: "go-test-index1",
		Body:  b,
	}
	fmt.Println(req.Body)
	insertResponse, err := req.Do(context.Background(), conn)
	if err != nil {
		c.Data["json"] = err
	} else {
		rs := make(map[string]interface{})
		fmt.Println(req.Body)
		rs["request"] = req
		rs["responese"] = insertResponse
		c.Data["json"] = rs
	}

	c.ServeJSON()
}
