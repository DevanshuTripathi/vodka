package vodka

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context) // Handler Function with Context wrapping

type M map[string]any // Shortcut map

type Context struct {
	Writer     http.ResponseWriter // net/http response writer
	Request    *http.Request       // net/http request
	Params     httprouter.Params   // URL Parameters for dynamic routing
	handlers   []HandlerFunc       // stores middleware funcs and also main handler func
	index      int8                // tracks current step
	queryCache url.Values          // Caches query parameter values for fast access
}

// Step By Step execution of middlewares
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Find Query Values
func (c *Context) Query(key string) string {
	if c.queryCache == nil {
		c.queryCache = c.Request.URL.Query()
	}
	return c.queryCache.Get(key)
}

// Default Query returns value if exists otherwise a default value
func (c *Context) DefaultQuery(key string, defautlValue string) string {
	if c.queryCache == nil {
		c.queryCache = c.Request.URL.Query()
	}

	if values, ok := c.queryCache[key]; ok && len(values) > 0 {
		return values[0]
	}

	return defautlValue
}

// Get Param Value
func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

// BindJSON parses the Request Body
func (c *Context) BindJSON(obj any) error {
	if c.Request.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	defer c.Request.Body.Close()

	decoder := json.NewDecoder(c.Request.Body)
	return decoder.Decode(obj)
}

func (c *Context) JSON(statusCode int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)
	json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) String(statusCode int, text string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write([]byte(text))
}
