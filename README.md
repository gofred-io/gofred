### gofred

Build responsive web applications in Go with WebAssembly.

gofred lets you write UI in pure Go and run it in the browser via WebAssembly, with a small set of primitives for layout, components, routing, and events.

---

#### Features
- **WASM-first**: Compile Go to WebAssembly and run directly in the browser.
- **Composable widgets**: `container`, `row`, `column`, `text`, `image`, `button`, and more.
- **Routing and navigation**: Lightweight router and navigation helpers.
- **Responsive layout**: Breakpoints and spacing utilities.
- **Batteries included**: Dev server for local development.

---

#### Requirements
- Go 1.25 or newer (as defined in `go.mod`)
- A modern browser with WebAssembly support (all evergreen browsers)

---

#### Install

```bash
go get github.com/gofred-io/gofred
```

This repository is intended to be used as a library. The included `server/` and `main.go` demonstrate a minimal runnable example.

---

#### Quick start (example app in this repo)

1) Build the WebAssembly binary:

```bash
make build
```

This produces `server/main.wasm`.

2) Start the dev server:

```bash
make serve
```

Then open the printed local URL (defaults to `http://localhost:8080`). The `server/index.html` bootstraps `wasm_exec.js` and loads `main.wasm`.

---

#### Minimal example

Create your own `main.go` and use gofred widgets to build UI. For example:

```go
package main

import (
	"github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/foundation/text"
)

func main() {
    // See this repository's `main.go` for a complete runnable example
    application.Run(text.New("Hello, gofred!"))
}
```

Explore the `foundation/` packages for available widgets and options.

---

#### Project structure (high level)

- `basic/` – Low-level HTML-like elements (e.g. `div`, `span`, `img`, `svg`).
- `foundation/` – Higher-level components (layout, text, buttons, router, etc.).
- `hooks/` – React-like hooks for navigation, breakpoint, and state management.
- `widget/` – Core widget abstractions and lifecycle.
- `server/` – Minimal static server and HTML bootstrap for wasm.

---

#### Development

- Build wasm: `make build`
- Run dev server: `make serve`

The CI workflow runs basic build and vet steps.

---

#### Browser support

Any modern browser with WebAssembly support. If you need older browser support, consider a fallback or a progressive enhancement approach.

---

#### Contributing

Issues and pull requests are welcome. Please read `CONTRIBUTING.md` first.

---

#### Website

For the official website source and additional docs/material, see the gofred website repository: `https://github.com/gofred-io/gofred-website`.

The website's README is available at: `https://github.com/gofred-io/gofred-website/blob/master/README.md`.

---

#### License

MIT — see `LICENSE` for details.


