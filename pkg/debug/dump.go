package debug

import (
	"github.com/gin-gonic/gin"
	"github.com/goforj/godump"
)

func Dump(value any) {
	godump.Dump(value)
}

func DumpHTMLToResponse(c *gin.Context, value any) {
	c.Data(200, "text/html; charset=utf-8", []byte(godump.DumpHTML(value)))
}
