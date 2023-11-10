package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%d] - %s - [%s] %s | %s | %s \n",
			param.StatusCode,
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.Latency,
		)
	})

}

func SetUpLogOutput() {

	f, err := os.Create("gin.log")
	if err != nil {
		fmt.Println("File creating error", err)
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
