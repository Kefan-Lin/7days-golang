package main

import(
	"net/http"
	"gee"
)

func main(){
	engine := gee.New()
	engine.GET("/", func(c *gee.Context){
		c.HTML(http.StatusOK, "<h1>Hello Gee<h1>")
	})
	engine.GET("/hello", func(c *gee.Context){
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})
	engine.POST("/login", func(c *gee.Context){
		c.JSON(http.StatusOK, gee.Hmap{
			"username" : c.PostForm("username"),
			"password" : c.PostForm("password"),
		})
	})

	engine.Run(":9999")
}
