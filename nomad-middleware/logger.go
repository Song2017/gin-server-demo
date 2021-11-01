package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	model "apiserver/v1/nomad-model"
)

type ResponseBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (r ResponseBodyWriter) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return r.ResponseWriter.Write(b)
}

func MaskURL(query *url.URL) string {
	paras := query.Query()
	if paras["Authorization"] != nil {
		return "user_id"
	}
	return query.RawQuery
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		body, _ := ioutil.ReadAll(c.Request.Body)
		nomadLog := model.NomadLog{
			Platform:    "NOMAD_ENVOY_UTILS",
			TimingStart: start.UTC().String(),
			RemoteAddr:  c.Request.RemoteAddr,
			Label:       strings.Join([]string{c.Request.Method, c.Request.URL.Path}, "-"),
			Input: fmt.Sprintf("query: %s, body: %s",
				MaskURL(c.Request.URL), string(body)),
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		w := &ResponseBodyWriter{
			Body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		// Process request
		c.Next()

		nomadLog.TimingEnd = time.Now().UTC().String()
		nomadLog.Latency = time.Since(start).Seconds()
		nomadLog.Response = fmt.Sprintf("status: %s, body: %s",
			strconv.Itoa(c.Writer.Status()), w.Body.String())
		log_data, _ := json.Marshal(nomadLog)
		log.Println(string(log_data))
	}
}
