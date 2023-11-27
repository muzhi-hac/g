package gee

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Method     string
	Path       string
	StatusCode int
	Params     map[string]string
	//StatusCode int
	index    int
	engine   *Engine
	handlers []HandleFunc
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}

}
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
	//c.Writer.Write("")
}

func (c *Context) String(code int, obj interface{}) {
	// Set the Content-Type header to indicate JSON response
	c.Writer.Header().Set("Content-Type", "application/json")

	// Set the status code for the response
	c.Writer.WriteHeader(code)

	// Encode the object to JSON and write it to the response
	err := json.NewEncoder(c.Writer).Encode(obj)
	if err != nil {
		// Handle error (e.g., log it or return an internal server error)
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
