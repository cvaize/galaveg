package static_files

import (
	"galaveg/internal/config"
	"github.com/gin-gonic/gin"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func GzipStaticMiddleware(C *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			filePath := C.GetFolder(filepath.Join("public", c.Request.URL.Path))

			gzipPath := filePath + ".gz"
			if _, err := os.Stat(gzipPath); err == nil {
				file, err := os.Open(gzipPath)
				if err == nil {
					defer file.Close()

					// Устанавливаем заголовки
					c.Header("Content-Encoding", "gzip")
					c.Header("Vary", "Accept-Encoding")
					c.Header("Cache-Control", "public, max-age=31536000")

					ext := filepath.Ext(filePath)
					contentType := mime.TypeByExtension(ext)
					if contentType != "" {
						c.Header("Content-Type", contentType)
					}

					c.File(gzipPath)
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
