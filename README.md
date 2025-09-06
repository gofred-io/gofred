### gofred

[![CI](https://img.shields.io/github/actions/workflow/status/gofred-io/gofred/ci.yml?label=CI)](https://github.com/gofred-io/gofred/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/gofred-io/gofred.svg)](https://pkg.go.dev/github.com/gofred-io/gofred)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Build responsive web applications in Go with WebAssembly.

gofred lets you write UI in pure Go and run it in the browser via WebAssembly, with a small set of primitives for layout, components, routing, and events.

---

#### Table of contents
- Features
- Requirements
- Install
- Quick start
- Minimal example
- Project structure
- Development
- Browser support
- Website
- Contributing
- License

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

This repository is intended to be used as a library. For development and running applications, use the [gofred-cli](https://github.com/gofred-io/gofred-cli) tool.

---

#### Quick start

Use the [gofred-cli](https://github.com/gofred-io/gofred-cli) tool to create and run applications:

1) Install gofred-cli:

```bash
curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/refs/heads/master/install.sh | bash
```

2) Create a new application:

```bash
gofred app create my-app --package my-app
```

3) Run the application:

```bash
cd my-app
gofred app run
```

This will start a development server with hot reload capabilities.

---

#### Minimal example

When you create a new application with gofred-cli, it generates a `main.go` file like this:

```go
package main

import (
	"github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/foundation/text"
)

func main() {
    app := text.New("Hello, world")
    application.Run(app)
}
```

Explore the `foundation/` packages for available widgets and options.

---

#### Project structure (high level)

- `basic/` – Low-level HTML-like elements (e.g. `div`, `span`, `img`, `svg`).
- `foundation/` – Higher-level components (layout, text, buttons, router, etc.).
- `hooks/` – React-like hooks for navigation, breakpoint, and state management.
- `widget/` – Core widget abstractions and lifecycle.

For development tooling and server setup, see [gofred-cli](https://github.com/gofred-io/gofred-cli).

---

#### Development

For development, use the [gofred-cli](https://github.com/gofred-io/gofred-cli) tool:

- Create app: `gofred app create my-app --package my-app`
- Run dev server: `gofred app run` (with hot reload)

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


