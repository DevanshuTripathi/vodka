package vodka

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context) // Handler Function with Context wrapping

type M map[string]any // Shortcut map

const abortIndex int8 = 63 // High Abort Number

var validate = validator.New() // validator for struct binding

// Custom error type
type VodkaError struct {
	Err    error
	Status int
}

// return string error
func (v *VodkaError) Error() string {
	return v.Err.Error()
}

type Context struct {
	Writer     http.ResponseWriter // net/http response writer
	Request    *http.Request       // net/http request
	Params     httprouter.Params   // URL Parameters for dynamic routing
	Keys       map[string]any      // Key-Value Store
	Errors     []*VodkaError       // Stores errors along the middleware chain
	handlers   []HandlerFunc       // stores middleware funcs and also main handler func
	index      int8                // tracks current step
	queryCache url.Values          // Caches query parameter values for fast access
	isAborted  bool
}

// Abort http request
func (c *Context) Abort() {
	c.index = abortIndex
	c.isAborted = true // Aborted
}

// Step By Step execution of middlewares doesnt handle abort
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

// Set Values into Keys into the context to be used later by another middleware
func (c *Context) Set(key string, value any) {
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}
	c.Keys[key] = value
}

// Get value from a key after a previous middleware assigned it
func (c *Context) Get(key string) (value any, exists bool) {
	value, exists = c.Keys[key]
	return
}

// Retrieves text value from a miltipart form
func (c *Context) FormValue(key string) string {
	return c.Request.FormValue(key)
}

// Retrieves a single file from multipart form
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			return nil, err
		}
	}

	file, header, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	file.Close()

	return header, nil
}

// Saves Uploaded File to destination on disk
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
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
	if err := decoder.Decode(obj); err != nil {
		return err // Failed to parse JSON
	}

	if err := validate.Struct(obj); err != nil {
		return err // Failed validation
	}

	return nil
}

// Encode Response to JSON
func (c *Context) JSON(statusCode int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)
	json.NewEncoder(c.Writer).Encode(obj)
}

// Stores Errors in the context to be handled by ErrorHandler Middleware
func (c *Context) Error(statusCode int, err error) {
	if err == nil {
		return
	}
	c.Errors = append(c.Errors, &VodkaError{
		Err:    err,
		Status: statusCode,
	})
}

// Basic String Response
func (c *Context) String(statusCode int, text string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write([]byte(text))
}
