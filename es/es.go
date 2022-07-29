package es

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	opensearch "github.com/opensearch-project/opensearch-go"
	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"
)

const IndexName = "go-test-index1"

// NewES returns a LoggerInterface
func NewES() logs.Logger {
	cw := &esLogger{
		Level: logs.LevelDebug,
	}
	return cw
}

// esLogger will log msg into ES
// before you using this implementation,
// please import this package
// usually means that you can import this package in your main package
// for example, anonymous:
// import _ "github.com/beego/beego/v2/core/logs/es"
type esLogger struct {
	*opensearch.Client
	DSN       string `json:"dsn"`
	Level     int    `json:"level"`
	formatter logs.LogFormatter
	Formatter string `json:"formatter"`
}

func (el *esLogger) Format(lm *logs.LogMsg) string {
	msg := lm.OldStyleFormat()
	idx := LogDocument{
		Timestamp: lm.When.Format(time.RFC3339),
		Msg:       msg,
	}
	body, err := json.Marshal(idx)
	if err != nil {
		return msg
	}
	return string(body)
}

func (el *esLogger) SetFormatter(f logs.LogFormatter) {
	el.formatter = f
}

// {"dsn":"http://localhost:9200/","level":1}
func (el *esLogger) Init(config string) error {
	err := json.Unmarshal([]byte(config), el)
	if err != nil {
		return err
	}
	if el.DSN == "" {
		return errors.New("empty dsn")
	} else if u, err := url.Parse(el.DSN); err != nil {
		return err
	} else if u.Path == "" {
		return errors.New("missing prefix")
	} else {
		conn, err := opensearch.NewClient(opensearch.Config{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Addresses: []string{"https://localhost:9200"},
			Username:  "admin", // For testing only. Don't store credentials in code.
			Password:  "admin",
		})
		if err != nil {
			return err
		}
		el.Client = conn
	}
	if len(el.Formatter) > 0 {
		fmtr, ok := logs.GetFormatter(el.Formatter)
		if !ok {
			return errors.New(fmt.Sprintf("the formatter with name: %s not found", el.Formatter))
		}
		el.formatter = fmtr
	}
	return nil
}

// WriteMsg writes the msg and level into es
func (el *esLogger) WriteMsg(lm *logs.LogMsg) error {
	if lm.Level > el.Level {
		return nil
	}
	log.Println("msssg")
	msg := el.formatter.Format(lm)

	req := opensearchapi.IndexRequest{
		Index: IndexName,
		Body:  strings.NewReader(msg),
	}
	_, err := req.Do(context.Background(), el.Client)
	return err
}

// Destroy is a empty method
func (el *esLogger) Destroy() {
}

// Flush is a empty method
func (el *esLogger) Flush() {
}

type LogDocument struct {
	Timestamp string `json:"timestamp"`
	Msg       string `json:"msg"`
}

func init() {
	logs.Register(logs.AdapterEs, NewES)
}
