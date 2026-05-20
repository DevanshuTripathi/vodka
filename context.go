package vodka

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context) // Handler Function with Context wrapping

type M map[string]any // Shortcut map

const abortIndex int8 = 63 // High Index for Abort Number

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

var contextPool = sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

// helper function to reset Context
func (c *Context) Reset() {
	// Delete Old keys and still be able to resue the Keys map without re-allocation
	for key := range c.Keys {
		delete(c.Keys, key)
	}
	c.Writer = nil
	c.Request = nil
	c.Params = nil
	c.handlers = nil
	c.index = -1
	c.queryCache = nil
	c.Errors = nil
	c.isAborted = false
}

// helper to initialize Context
func (c *Context) Initialize(w http.ResponseWriter, r *http.Request, p httprouter.Params, handlers []HandlerFunc) {
	c.Writer = w
	c.Request = r
	c.Params = p
	c.handlers = handlers
	c.index = -1
}

// helper function to copy context for using it outisde the request lifecycle
func (c *Context) Copy() *Context {
	cp := &Context{
		Request:   c.Request,
		index:     c.index,
		isAborted: c.isAborted,
	}

	if c.Params != nil {
		cp.Params = make(httprouter.Params, len(c.Params))
		copy(cp.Params, c.Params)
	}

	if c.handlers != nil {
		cp.handlers = make([]HandlerFunc, len(c.handlers))
		copy(cp.handlers, c.handlers)
	}

	if c.Keys != nil {
		cp.Keys = make(map[string]any, len(c.Keys))
		for k, v := range c.Keys {
			cp.Keys[k] = v
		}
	}

	if c.Errors != nil {
		cp.Errors = make([]*VodkaError, len(c.Errors))
		copy(cp.Errors, c.Errors)
	}

	return cp
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
		if c.isAborted {
			return
		}
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

	if values, ok := c.queryCache[key]; ok && len(values) > 0 && values[0] != "" {
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
	c.Abort()
}

// Basic String Response
func (c *Context) String(statusCode int, text string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write([]byte(text))
}

// Helper to get request IP
func (c *Context) IP() string {
	return c.Request.RemoteAddr
}

// QueryInt returns query param as int or an error
func (c *Context) QueryInt(key string) (int, error) {
	val, err := strconv.Atoi(c.Query(key))
	if err != nil {
		return 0, fmt.Errorf(Red+"query param %q is not a valid int"+Reset, key)
	}
	return val, nil
}

// QueryBool returns query param as bool or an error
func (c *Context) QueryBool(key string) (bool, error) {
	val, err := strconv.ParseBool(c.Query(key))
	if err != nil {
		return false, fmt.Errorf(Red+"query param %q is not a valid bool"+Reset, key)
	}
	return val, nil
}

// ParamInt returns URL param as int or an error
func (c *Context) ParamInt(key string) (int, error) {
	val, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return 0, fmt.Errorf(Red+"param %q is not a valid int"+Reset, key)
	}
	return val, nil
}

// ParamBool returns URL param as bool or an error
func (c *Context) ParamBool(key string) (bool, error) {
	val, err := strconv.ParseBool(c.Param(key))
	if err != nil {
		return false, fmt.Errorf(Red+"param %q is not a valid bool"+Reset, key)
	}
	return val, nil
}

// SetCookie writes a cookie to the response.
// maxAge is in seconds. Use 0 for session cookie, negative to delete.
func (c *Context) SetCookie(name, value string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Path:     "/",
		HttpOnly: true,
	})
}

// Cookie returns the value of the named cookie from the request.
// Returns an error if the cookie is not found.
func (c *Context) Cookie(name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// ClearCookie deletes a cookie by setting its MaxAge to -1.
func (c *Context) ClearCookie(name string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
}
