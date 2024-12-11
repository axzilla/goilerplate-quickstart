# templUI Quickstart

Get started with templUI, a modern UI component library for Go and Templ. This template provides a pre-configured setup for building web applications with templUI components.

## Features

- ✨ Pre-configured templUI components
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

For installation instructions, visit our [documentation](https://templui.io/docs/how-to-use#requirements).

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/axzilla/templui-quickstart.git
   cd templui-quickstart
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Configure Tailwind**
   Since we're using templUI as a package, you need to configure Tailwind to process its components:

   a. Get your Go path:

   ```bash
   go env GOPATH
   ```

   b. Add the path to your `tailwind.config.js` content array:

   ```js
   content: [
     "./**/*.{html,templ,go}",
     "${GOPATH}/pkg/mod/github.com/axzilla/templui@*/**/*.{go,templ}", // Replace ${GOPATH} with your actual Go path
   ];
   ```

4. **Start Development Server**

   ```bash
   make dev
   ```

   Your application will be running at [http://localhost:7331](http://localhost:7331)

## Learn More

- [templUI Documentation](https://templui.io/docs/how-to-use)
- [Available Components](https://templui.io/docs/components)
- [Theme Customization](https://templui.io/docs/themes)

## Contributing

Issues and pull requests are welcome! Please read our [contributing guidelines](https://github.com/axzilla/templui/blob/main/CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
