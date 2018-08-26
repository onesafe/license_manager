package log

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger(out io.Writer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		comment := ctx.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Fprintf(out, "%v\t%d\t%v\t%s\t%s\t%s\t%s\n", end.Format(time.RFC3339), statusCode, latency, clientIP, method, path, comment)

	}
}