package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hmap map[string]interface{}


type Context struct {
	// Origin objects
	Writer http.ResponseWriter
	Req *http.Request
	// Request info
	Path string
	Method string
	// Response info
	StatusCode int
}

//
//  newContext
//  @Description: create a Context object
//  @param w, a http.ResponseWriter object
//  @param req, a *http.Request pointer
//  @return *Context, a *Context pointer that points to the new Context object
//
func newContext(w http.ResponseWriter, req *http.Request) *Context{
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

//
//  PostForm
//  @Description: get a value from the post form by the key
//  @receiver c
//  @param key
//  @return string
//
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//
//  Query
//  @Description: query a value in the URL by the key. E.g., c.Query("name") => "Alice"
//  @receiver c, the context object
//  @param key
//  @return string, the value
//
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//
//  Status
//  @Description: set the status code for the response
//  @receiver c
//  @param code
//
func (c *Context) Status(code int){
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//
//  SetHeader
//  @Description: set a key-value field for the response header
//  @receiver c
//  @param key
//  @param value
//
func (c *Context) SetHeader(key string, value string){
	c.Writer.Header().Set(key, value)
}

//
//  String
//  @Description: write a string of data into the response
//  @receiver c
//  @param code, the status code
//  @param format, the formatted string
//  @param values, the formatted values that will be inserted into the string
//
func (c *Context) String(code int, format string, values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,values...)))
}

//
//  JSON
//  @Description: write a json object to the response body
//  @receiver c
//  @param code, the status code
//  @param obj, the json object
//
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//
//  Data
//  @Description: write raw data to the response
//  @receiver c
//  @param code, the status code
//  @param data, a byte slice
//
func (c *Context) Data(code int , data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

//
//  HTML
//  @Description: write html content to the response
//  @receiver c
//  @param code, the status code
//  @param html, a string that contains the html file
//
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

