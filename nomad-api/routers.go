package apiserver

import (
	middleware "apiserver/v1/nomad-middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// prometheus statics
	prometheus.MustRegister(middleware.StatusCounter)

	// middlewares
	router.Use(
		middleware.LoggerMiddleware(),
		middleware.PrometheusMiddleware(),

		gin.Recovery(),
	)

	// endpoints
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server page")
	})

	root_group := router.Group("/")
	{
		root_group.GET("/health", GetHealth)
		root_group.GET("/metrics", gin.WrapH(promhttp.Handler()))

		// http://localhost:8000/cbec/openapi.yaml
		root_group.GET("/openapi.yaml",
			func(c *gin.Context) {
				if gin.Mode() == gin.ReleaseMode {
					c.String(404, "openapi.yaml not found")
					return
				}
			},
			func(c *gin.Context) { c.File("./swagger_server/api/openapi.yaml") },
		)
		// swagger UI: url pointing to API definition
		// http://localhost:8000/swagger/index.html
		router.GET("/swagger/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler, ginSwagger.URL("/openapi.yaml")))
	}

	platform_group := router.Group("/platform")
	{
		platform_group.POST(
			"/cypher",
			middleware.AuthorizationMiddleware(),
			BatchOperatecypher,
		)
	}

	return router
}
