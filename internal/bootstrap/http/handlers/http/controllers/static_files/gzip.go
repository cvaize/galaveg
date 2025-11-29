package static_files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Gzip(filepath, contentType string) gin.HandlerFunc {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	gzipContent, err := os.ReadFile(filepath + ".gz")
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Header("Content-Encoding", "gzip")
			c.Header("Vary", "Accept-Encoding")
			c.Data(http.StatusOK, contentType, gzipContent)
			return
		}

		c.Data(http.StatusOK, contentType, content)
	}
}
