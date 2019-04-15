// file: web/controllers/hello_controller.go

package controllers

import (
	"github.com/kataras/iris/mvc"
)

// HelloController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type HomeController struct{}

var homeView = mvc.View{
	Name: "hello.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
	},
}

// Get will return a predefined view with bind data.
//
// `mvc.Result` is just an interface with a `Dispatch` function.
// `mvc.Response` and `mvc.View` are the built'n result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func (c *HomeController) Get() mvc.Result {
	return helloView
}
