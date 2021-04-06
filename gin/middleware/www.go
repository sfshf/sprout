package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func WWW(root string) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		fpath := filepath.Join(root, filepath.FromSlash(p))
		_, err := os.Stat(fpath)
		if err != nil && os.IsNotExist(err) {
			fpath = filepath.Join(root, "index.html")
		}
		c.File(fpath)
		c.Abort()
	}
}
