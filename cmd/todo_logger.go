package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	loggs "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"runtime"
	"time"
)

func cpuUsageMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			loggs.SetLevel(loggs.DebugLevel)

			var beforeCPU runtime.MemStats
			runtime.ReadMemStats(&beforeCPU)

			err := next(c)

			var afterCPU runtime.MemStats
			runtime.ReadMemStats(&afterCPU)

			duration := time.Since(start)
			cpuUsage := afterCPU.TotalAlloc - beforeCPU.TotalAlloc

			logger.Info("Request details",
				zap.String("URI", c.Request().RequestURI),
				zap.Duration("Duration", duration),
				zap.Uint64("CPU usage (bytes)", cpuUsage),
			)

			fmt.Printf("logger:%s Request %s took %v, CPU usage: %v bytes\n", "printF", c.Request().RequestURI, duration, cpuUsage)
			log.Errorf("logger:%s Request %s took %v, CPU usage: %v bytes\n", "loggerErrof", c.Request().RequestURI, duration, cpuUsage)
			loggs.Debugf("logger:%s Request %s took %v, CPU usage: %v bytes\n", "debug", c.Request().RequestURI, duration, cpuUsage)

			return err
		}
	}
}

//func main() {
//	logger, _ := zap.NewProduction()
//	defer logger.Sync()
//
//	e := echo.New()
//
//	e.Use(cpuUsageMiddleware(logger))
//
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	e.GET("/fet", func(c echo.Context) error {
//		return c.String(http.StatusOK, "Hello, World!")
//	})
//
//	e.Start(":8080")
//}
