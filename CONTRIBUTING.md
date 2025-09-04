### Contributing to gofred

Thanks for your interest in contributing! This document outlines how to propose changes.

---

#### Code of conduct

By participating, you agree to uphold our standards of respectful, inclusive collaboration. See `CODE_OF_CONDUCT.md`.

---

#### Getting started
- Fork the repo and clone your fork.
- Ensure you have Go 1.25+ installed.
- Install dependencies via `go mod download`.
- Build the wasm demo once to ensure your environment is set up: `make build`.

---

#### Development workflow
- Create a topic branch from `master`.
- Keep changes focused and atomic.
- Write clear commit messages.
- Run checks locally:
  - `go fmt ./...`
  - `go vet ./...`
  - `make build` (verifies WASM build)

---

#### Pull requests
- Open a PR against `master`.
- Describe the change: motivation, approach, tradeoffs.
- Add tests or sample usage where appropriate.
- Ensure CI is green.

---

#### Issue reports
- Use the issue templates.
- Include: expected vs actual behavior, reproduction steps, environment details (OS, browser, Go version).

---

#### Style and guidelines
- Prefer clear, explicit names over abbreviations.
- Handle errors explicitly and early.
- Avoid deep nesting; use guard clauses.
- Keep public APIs well-documented with Go doc comments.

---

#### Security

Please report security issues privately. See `SECURITY.md`.

---

#### License

By contributing, you agree that your contributions are licensed under the MIT License.


