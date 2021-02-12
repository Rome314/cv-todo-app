package logger


import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GinLogger(logger log.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		start := time.Now()
		c.Next()

		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		statusCode := c.Writer.Status()
		// clientIP := c.ClientIP()
		reqBts, _ := ioutil.ReadAll(c.Request.Body)

		if raw != "" {
			path = path + "?" + raw
		}

		entry := logger.WithFields(log.Fields{
			"timestamp":  timeStamp.Format("02-01-2006 15:04:05.000"),
			"method":     c.Request.Method,
			"path":       path,
			"statusCode": statusCode,
			// "clientIP":   clientIP,
			"latency": latency, // time to process
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode >= http.StatusInternalServerError {
				entry.Error("")
			} else if statusCode >= http.StatusBadRequest {

				entry = entry.WithField("request", string(reqBts))
				entry.Warn("")
			} else {
				entry.Info("")
			}
		}
	}
}
