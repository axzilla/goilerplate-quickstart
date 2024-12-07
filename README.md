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

- Go (1.20 or later)
- Templ
- Tailwind CSS CLI (Standalone)
- Make (for development scripts)

For installation instructions, visit our [documentation](https://goilerplate.com/docs/how-to-use#requirements).

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

## Learn More

- [Goilerplate Documentation](https://goilerplate.com/docs/how-to-use)
- [Available Components](https://goilerplate.com/docs/components)
- [Theme Customization](https://goilerplate.com/docs/themes)

## Contributing

Issues and pull requests are welcome! Please read our [contributing guidelines](https://github.com/axzilla/goilerplate/blob/main/CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
