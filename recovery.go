package vodka

import (
	"log"
	"net/http"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf(Red+"[Vodka Panic Recovery] %v\n"+Reset, err)
				c.Abort()
				if rw, ok := c.Writer.(*responseWriter); ok && rw.wroteHeader {
					return
				}
				c.JSON(http.StatusInternalServerError, M{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}

