package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/wonderivan/logger"
	"lucky-draw/internal/lucky-drawn/apihandlers"
	"lucky-draw/web/controllers"
)

func init() {
	// https://github.com/wonderivan/logger
	//DEBG, INFO
	logConfig := `{
  	"TimeFormat": "2019-04-15 20:25:05",
  	"Console": {
   	 "level": "INFO",
   	 "color": true
  	},
  	"File": {
    	"filename": "./lucky-draw/lucky-draw.log",
    	"level": "INFO",
    	"daily": true,
    	"maxlines": 1000000,
    	"maxsize": 3,
    	"maxdays": -1,
    	"append": true,
    	"permit": "0660"
 	 }
	}`
	logger.SetLogger(logConfig)
}

func main() {
	app := iris.New()

	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.StaticWeb("/public", "./web/public")

	// Serve our controllers.
	//mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))

	mvc.New(app.Party("/home")).Handle(new(controllers.HomeController))

	app.Post("/reward",apihandlers.StartLuckyDraw)

	//app.Get("/home", apihandlers.GetIDMRequestTokenValidateIDMAccessToken, apihandlers.RouteRequestToNodejs)
	//// start the lucky draw
	//app.Post("/reward", apihandlers.ValidateIDMAccessToken, apihandlers.RefreshIDMAccessToken)

	app.Run(iris.Addr(":8080"))
}
