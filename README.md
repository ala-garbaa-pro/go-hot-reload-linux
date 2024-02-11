# Go Hot Reload Linux: A Go Web Application Template


## Introduction

Go Hot Reload is an open-source starter template, designed to streamline the development process with hot reloading for Go, Templ (HTML), htmx and Tailwind CSS. It's a perfect starting point for server side web development, particularly for those using `htmx`.

## Features

- 🔄 **Live Reloading**: Automatic Go rebuild using [`air`](https://github.com/cosmtrek/air). Configured in `.air.toml`.
- 📄 **Templ**: Templ is a component and typed html templating system - it plays nice with htmx.
- 💅 **Tailwind**: CSS framework for responsive, customizable UI components with [Tailwind](https://tailwindcss.com/).
- 🚀 **HTMX**: Enhance your HTML with AJAX, WebSockets, and more using [HTMX](https://htmx.org/), enabling rich interactions with minimal JavaScript.

## Getting Started

### Setting Up the Environment

1. **Install pnpm Modules**: Run `pnpm install` to set up necessary modules.
2. **Environment Configuration**: Create a `.env` file with the following content:
   ```
   PORT=3333
   ```
3. **Launch the Development Server**: Start the server using `pnpm run dev`.
4. **Hot Reload**: Change the tailwind classes in `templates/landing.html` to see hot reload in action.

### Running Tests

The server includes an integration test.

- Execute tests using `go test ./...`.
- Or using pnpm command `pnpm run test`

### Commands

Full list of commands for `pnpm run` are:

```json
"watch:tailwind": "npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --watch",
"watch:templ": "templ generate -path web/view  --watch --proxy='http://localhost:3333'",
"watch:go": "air",
"dev": "concurrently \"pnpm run watch:tailwind\" \"pnpm run watch:go\" \"pnpm run watch:templ\"",
"test": "go test ./..."
```