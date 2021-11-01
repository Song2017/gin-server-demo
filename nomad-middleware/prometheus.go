package apiserver

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var requestCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "http_server_requests_total",
	Help: "The total number of requests events",
})

var failCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "http_server_requests_error",
	Help: "The total number of requests with error",
})

var StatusCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_server_requests_status",
	Help: "Number of requests status in total",
},
	[]string{"code"},
)

var timingSummary = promauto.NewSummary(prometheus.SummaryOpts{
	Name: "http_server_requests_milliseconds",
	Help: "The total timing of requests",
})

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestCounter.Inc()
		now := time.Now()

		c.Next()

		timingSummary.Observe(float64(time.Now().UnixNano()-now.UTC().UnixNano()) / 1000000)
		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			failCounter.Inc()
		}
		StatusCounter.With(
			prometheus.Labels{
				"code": strconv.FormatInt(int64(statusCode), 10),
			},
		).Inc()
	}
}
