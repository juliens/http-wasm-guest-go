package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	_ "github.com/stealthrocket/net/http"
)

func main() {
	handler.HandleRequestFn = handleRequest
	handler.HandleInMain()
}

// handleRequest implements a simple HTTP router.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	fmt.Println(http.Get("http://127.0.0.1:8085"))

	// // If the URI starts with /host, trim it and dispatch to the next handler.
	if uri := req.GetURI(); strings.HasPrefix(uri, "/host") {
		req.SetURI(uri[5:])
		next = true // proceed to the next handler on the host.
		return
	}

	// Serve a static response
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().WriteString("hello")
	return // skip the next handler, as we wrote a response.
}
