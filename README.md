# Goilerplate Quickstart

Get started with Goilerplate, a modern UI component library for Go and Templ. This template provides a pre-configured setup for building web applications with Goilerplate components.

## Features

- ✨ Pre-configured Goilerplate components
- 🎨 Ready-to-use theme and styling
- 🚀 Development server with hot reload
- 🛠️ Example components and layouts
- 📱 Responsive by default
- 🌙 Light and dark mode support

## Prerequisites

Before starting, ensure you have these tools installed:

1. **Go** (1.20 or later)

   ```bash
   # Verify installation
   go version

   # Install from https://golang.org/dl/
   ```

2. **Templ**

   ```bash
   # Install
   go install github.com/a-h/templ/cmd/templ@latest

   # Verify installation
   templ version
   ```

3. **Tailwind CSS CLI** (Standalone)

   ```bash
   # Choose one installation method:

   # macOS (Homebrew)
   brew install tailwindcss

   # macOS (Apple Silicon)
   curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
   chmod +x tailwindcss-macos-arm64
   sudo mv tailwindcss-macos-arm64 /usr/local/bin/tailwindcss

   # macOS (Intel)
   curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
   chmod +x tailwindcss-macos-x64
   sudo mv tailwindcss-macos-x64 /usr/local/bin/tailwindcss

   # Linux (x64)
   curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
   chmod +x tailwindcss-linux-x64
   sudo mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

   # Windows (x64)
   # Download from: https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-windows-x64.exe
   # Add to your PATH as 'tailwindcss'

   # Verify installation
   tailwindcss --help
   ```

4. **Make** (for development scripts)
   - Pre-installed on macOS and Linux
   - Windows users might need to install it separately

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/axzilla/goilerplate-quickstart.git
   cd goilerplate-quickstart
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Configure Tailwind**
   Since we're using Goilerplate as a package, you need to configure Tailwind to process its components:

   a. Get your Go path:

   ```bash
   go env GOPATH
   ```

   b. Add the path to your `tailwind.config.js` content array:

   ```js
   content: [
     "./**/*.{html,templ,go}",
     "${GOPATH}/pkg/mod/github.com/axzilla/goilerplate@*/**/*.{go,templ}", // Replace ${GOPATH} with your actual Go path
   ];
   ```

4. **Start Development Server**

   ```bash
   make dev
   ```

   Your application will be running at [http://localhost:7331](http://localhost:7331)

## Project Structure

```
.
├── assets/
│   └── css/
│       ├── input.css    # Tailwind imports and custom styles
│       └── output.css   # Generated CSS (don't edit)
├── ui/
│   ├── components/      # Your custom components
│   ├── layouts/         # Page layouts
│   └── pages/          # Page templates
├── main.go             # Application entry point
├── Makefile           # Development commands
└── tailwind.config.js # Tailwind configuration
```

## Development Commands

- `make dev` - Start development server with hot reload (Templ, Tailwind, and Go)
- `make build` - Build for production
- `make clean` - Clean build artifacts

## Customization

1. **Add New Pages**

   - Create new templates in `ui/pages/`
   - Add routes in `main.go`

2. **Create Layouts**

   - Add layout templates in `ui/layouts/`
   - Use them in your pages with `@layouts.YourLayout()`

3. **Modify Styles**

   - Edit `assets/css/input.css` for custom styles
   - Configure theme in `tailwind.config.js`

4. **Add Components**

   - Create new components in `ui/components/`
   - Import existing Goilerplate components:

     ```go
     import "github.com/axzilla/goilerplate/pkg/components"

     // Use in templates
     @components.Button(components.ButtonProps{Text: "Click me"})
     ```

## Troubleshooting

- **Styles not updating?** Make sure Tailwind is watching for changes and your GOPATH is correctly configured
- **Components not found?** Run `go mod tidy` to ensure all dependencies are installed
- **Build errors?** Run `make clean` and try again
- **Other issues?** Check our [documentation](https://goilerplate.com/docs/how-to-use) or [open an issue](https://github.com/axzilla/goilerplate-quickstart/issues)

## Learn More

- [Goilerplate Documentation](https://goilerplate.com/docs/how-to-use)
- [Available Components](https://goilerplate.com/docs/components)
- [Theme Customization](https://goilerplate.com/docs/themes)

## Contributing

Issues and pull requests are welcome! Please read our [contributing guidelines](https://github.com/axzilla/goilerplate/blob/main/CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
