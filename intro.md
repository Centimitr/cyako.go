# Intro to Cyako.go
This is a quick intro about how to write a backend using Cyako.go quickly.

The project Cyako is to use an JSON wrapped data structure ( which contains id, method, param, data and error) to simplify the data exchange programming in simple C/S application. It avoid to use a router to allocate your clients' requests to a specific controller, maybe a function,  because it directly use the method name. 

In this project, Cyako.go ,realized the request/response type data exchange apis. The most amazing place is, it load middlewares, processor modules ( groups of controllers) and many things appear in a nearly future, as _ anonymous packages and register them to the Cyako global object. In this way, we can access modules, middlewares using its type name, and also, using the method name to access a module's processor ( controller). So we don't need to consider the router, the rest apis, which the route a controller function should be bind to , and how to name the controllers. All these are waste of time when we want to have a trying on a simple idea. All we want to do is to make it work, make it runnable, in a moment.

In this structure, it is easy to use a number of well-tested community middlewares written for http frameworks too. Few steps like write some essential hooks and register it is the most you need to do.

## Import packages

Here is an example of a Cyako.go server's main.go. We import middlewares and processor logic modules, and then setup a server. That's all.

```go
package main

import (

	// framework
	cyako "github.com/Cyako/Cyako.go"

	// middlewares
	_ "github.com/Cyako/Cyako.go/cache"
	_ "github.com/Cyako/Cyako.go/jsonbase"
	_ "github.com/Cyako/Cyako.go/statistics"

	// processor codules
	_ "github.com/Cyako/module"

	// systec library
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	c := cyako.Ins()
	c.PrintLoadInfo()

	// SERVER
	fmt.Println(" Running...")
	http.Handle("/echo", websocket.Handler(c.Server))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

```



## a simple processor module

## a simple middleware

