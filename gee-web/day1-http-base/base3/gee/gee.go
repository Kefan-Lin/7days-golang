package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)


//  Engine
//  @Description: Engine implements the interface of ServeHTTP
//
type Engine struct{
	router map[string]HandlerFunc
}

//  New
//  @Description: New is the constructor of gee.Engine
//  @return *Engine, a pointer to an Engine struct
//
func New() *Engine{
	return &Engine{router: make(map[string]HandlerFunc)}
}

//
//  addRoute
//  @Description: add a route in the Engine object
//  @receiver engine
//  @param method, is one of the HTTP methods like GET, POST, etc
//  @param pattern, is the URL pattern, such as "/hello/world"
//  @param handler, is the handler function for this route
//
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc){
	key := method + "-" + pattern
	engine.router[key] = handler
}

//
//  GET
//  @Description: resister a GET handler for a URL
//  @receiver engine
//  @param pattern, is the URL pattern, such as "/hello/world"
//  @param handler, is the handler function for this route
//
func (engine *Engine) GET(pattern string, handler HandlerFunc){
	engine.addRoute("GET", pattern, handler)
}

//
//  POST
//  @Description: resister a POST handler for a URL
//  @receiver engine
//  @param pattern, is the URL pattern, such as "/hello/world"
//  @param handler, is the handler function for this route
//
func (engine *Engine) POST(pattern string, handler HandlerFunc){
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request){
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w,r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}


