package es

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/beego/beego/v2/core/logs"
	opensearch "github.com/opensearch-project/opensearch-go"
	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"
)

var levelNames = [...]string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}

// NewES returns a LoggerInterface
func NewOpensarch() logs.Logger {
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
	IndexName string `json:"index"`
	formatter logs.LogFormatter
	Formatter string `json:"formatter"`
}

func (el *esLogger) Format(lm *logs.LogMsg) *bytes.Buffer {
	msg := lm.Msg

	if len(lm.Args) > 0 {
		msg = fmt.Sprintf(lm.Msg, lm.Args...)
	}
	m, b := LogDocument{
		Timestamp: lm.When.Format(time.RFC3339),
		Msg:       msg,
		Level:     levelNames[lm.Level],
		Line:      lm.LineNumber,
		Filename:  lm.FilePath,
	}, new(bytes.Buffer)

	json.NewEncoder(b).Encode(m)

	return b
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

func (el *esLogger) getIndexName(lm *logs.LogMsg) string {
	if el.IndexName != "" {
		return el.IndexName
	} else {
		return fmt.Sprintf("backend_api_log_%d%02d%02d", lm.When.Year(), lm.When.Month(), lm.When.Day())
	}
}

// WriteMsg writes the msg and level into es
func (el *esLogger) WriteMsg(lm *logs.LogMsg) error {
	if lm.Level > el.Level {
		return nil
	}

	req := opensearchapi.IndexRequest{
		Index: el.getIndexName(lm),
		Body:  el.Format(lm),
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
	Level     string `json:"level"`
	Filename  string `json:"filename"`
	Line      int    `json:"line"`
}
