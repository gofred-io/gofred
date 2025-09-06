### gofred

[![CI](https://img.shields.io/github/actions/workflow/status/gofred-io/gofred/ci.yml?label=CI)](https://github.com/gofred-io/gofred/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/gofred-io/gofred.svg)](https://pkg.go.dev/github.com/gofred-io/gofred)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Build responsive web applications in Go with WebAssembly.

gofred lets you write UI in pure Go and run it in the browser via WebAssembly, with a small set of primitives for layout, components, routing, and events.

> **🚀 Getting Started?** Use the [gofred-cli](https://github.com/gofred-io/gofred-cli) to quickly create and run gofred applications with hot reload development server.

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
- CLI Tool
- Website
- Contributing
- License

#### Features
- **WASM-first**: Compile Go to WebAssembly and run directly in the browser.
- **Composable widgets**: `container`, `row`, `column`, `text`, `image`, `button`, and more.
- **Routing and navigation**: Lightweight router and navigation helpers.
- **Responsive layout**: Breakpoints and spacing utilities.
- **Type-safe**: Pure Go code with compile-time type checking.

---

#### Requirements
- Go 1.25 or newer (as defined in `go.mod`)
- A modern browser with WebAssembly support (all evergreen browsers)

---

#### Install

```bash
go get github.com/gofred-io/gofred
```

This repository is intended to be used as a library. For a complete development experience with hot reload and development server, use the [gofred-cli](https://github.com/gofred-io/gofred-cli) tool.

---

#### Quick start

The easiest way to get started with gofred is using the [gofred-cli](https://github.com/gofred-io/gofred-cli):

```bash
# Install gofred-cli
curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/main/install.sh | bash

# Create a new application
gofred app create my-app

# Navigate to your app and run it
cd my-app
gofred app run
```

This will create a complete gofred application with hot reload development server.

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

---

#### Development

For library development:

- Build wasm: `make build`
- Run tests: `go test ./...`
- Run linter: `golangci-lint run`

The CI workflow runs basic build and vet steps.

For application development, use the [gofred-cli](https://github.com/gofred-io/gofred-cli) which provides hot reload and development server.

---

#### Browser support

Any modern browser with WebAssembly support. If you need older browser support, consider a fallback or a progressive enhancement approach.

---

#### CLI Tool

For the best development experience, use the [gofred-cli](https://github.com/gofred-io/gofred-cli) which provides:

- **Quick project setup**: Create new gofred applications with a single command
- **Hot reload**: Automatic recompilation and browser refresh on file changes
- **Development server**: Built-in server with WebSocket support
- **Cross-platform**: Support for Windows, macOS, and Linux
- **VS Code integration**: Automatic workspace configuration

Install and get started:

```bash
# Install gofred-cli
curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/main/install.sh | bash

# Create and run your first app
gofred app create my-app
cd my-app
gofred app run
```

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


