package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	name string
}

func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello %q", h.name)
}
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi There! %q", ctx.RequestURI())
}

func main() {

	myhandler := &MyHandler{
		name: "Yahiya",
	}

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/home":
			fastHTTPHandler(ctx)
		case "/index":
			myhandler.HandleFastHTTP(ctx)
		default:
			ctx.Error("not found!", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(":8082", m)
}
