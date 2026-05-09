package vodka

// Error Handler Middleware
func ErrorHandler() HandlerFunc {
	return func(c *Context) {
		c.Next() // Pass request through all middlewares first

		if len(c.Errors) != 0 {
			err := c.Errors[len(c.Errors)-1] // Get the last error

			c.JSON(err.Status, M{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
