package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"os"
)

// hosts: add these lines: (use tabs to separate them, if that doesn't works for you)
// 127.0.0.1 mydomain.com
// 127.0.0.1 api.mydomain.com
func main() {
	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	//app.Use(logger.New())
	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		//Columns: true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKey: "logger_message",
	})

	app.Use(customLogger)
	//
	// REGISTER YOUR REST API
	//

	// http://api.mydomain.com/ ...
	// brackers are optional, it's just a visual declaration.
	api := app.Party("/api")

	{
		// http://api.mydomain.com/users/42
		api.Get("/users/:userid", func(ctx context.Context) {
			//	ctx.Writef("user with id: %s", ctx.Param("userid"))
		})

	}

	//
	// REGISTER THE PAGE AND ALL OTHER STATIC FILES
	// INCLUDING A FAVICON, CSS, JS and so on
	//

	// http://mydomain.com , here should be your index.html
	// which is the SPA frontend page
	app.StaticWeb("/", "./app/dist")

	// or catch all http errors:
	app.OnAnyErrorCode(customLogger, func(ctx context.Context) {
		// this should be added to the logs, at the end because of the `logger.Config#MessageContextKey`
		ctx.Values().Set("logger_message","a dynamic message passed to the logs")
		ctx.Writef("My Custom error page")
	})


	var port string = os.Getenv("portnumber");
	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}