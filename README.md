[![Build](https://github.com/http-wasm/http-wasm-guest-tinygo/workflows/build/badge.svg)](https://github.com/http-wasm/http-wasm-guest-tinygo)
[![Go Report Card](https://goreportcard.com/badge/github.com/http-wasm/http-wasm-guest-tinygo)](https://goreportcard.com/report/github.com/http-wasm/http-wasm-guest-tinygo)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

# http-wasm Guest Library for Go

[http-wasm][1] is HTTP client middleware implemented in [WebAssembly][2].
This is a Go WASI library that implements the [Guest ABI][4].

## Example
The following is an [example](examples/router) of routing middleware:

```go
package main

import (
	"strings"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = handleRequest
	handler.HandleInMain()
}

// handle implements a simple HTTP router.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	// If the URI starts with /host, trim it and dispatch to the next handler.
	if uri := req.GetURI(); strings.HasPrefix(uri, "/host") {
		req.SetURI(uri[5:])
		next = true // proceed to the next handler on the host.
		return
	}

	// Serve a static response
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().WriteString("hello")
	return // skip any handlers as the response is written.
}
```

If you make changes, you can rebuild this with Go v1.21 or higher like so:
```sh
GOOS=wasm GOARCH=wasi go build -o examples/router/main.wasm examples/router/main.go
```

There are also more [examples](examples) you may wish to try out!

# WARNING: This is an early draft

The current maturity phase is early draft. Once this is integrated with
[coraza][5] and [dapr][6], we can begin discussions about compatability.

[1]: https://github.com/http-wasm
[2]: https://webassembly.org/
[3]: https://wazero.io/languages/tinygo/
[4]: https://github.com/http-wasm/http-wasm-abi
[5]: https://github.com/corazawaf/coraza-proxy-wasm
[6]: https://github.com/http-wasm/components-contrib/
