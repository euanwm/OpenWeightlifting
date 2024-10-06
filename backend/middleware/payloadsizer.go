package middleware

import (
	"backend/structs"
	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	bytes int
	bank  *structs.BeanCounter
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(data)
	rw.bytes += size
	rw.bank.AddBytes(uint64(size))
	return size, err
}

func PayloadSizer(theBank *structs.BeanCounter) gin.HandlerFunc {
	return func(c *gin.Context) {
		rw := &responseWriter{
			ResponseWriter: c.Writer,
			bytes:          0,
			bank:           theBank,
		}
		c.Writer = rw
		c.Next()
	}
}
