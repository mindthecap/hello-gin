package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	metricsRouter := gin.Default()

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.Use(metricsRouter)

	router.GET("/api/v1/gin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello " + os.Getenv("USERNAME") + " from " + os.Getenv("MY_POD_NAME"))
	})

	// Run the main router on port 8000
	go func() {
		router.Run(":8000")
	}()

	// Run the metrics router on a different port, e.g., 8080
	metricsRouter.Run(":8080")
}
