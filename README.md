# TemplUI Quickstart

Get started with TemplUI, an enterprise-ready UI component library for Go and Templ. This template provides a pre-configured setup for building professional web applications with TemplUI components.

## Features

- ⚡️ Enterprise-ready components
- 🛡️ CSP compliant by default
- 🔄 HTMX optimized
- 🧩 Component driven architecture
- 🎨 Fully customizable design
- 🚀 Development server with hot reload
- 📱 Responsive by default
- 🌙 Light and dark mode support

## Security

This template is CSP-compliant out of the box through TemplUI's built-in security features:

- Uses Alpine.js CSP build
- All components use nonces via templ
- No inline scripts

If you plan to add external scripts (like HTMX or highlight.js), you'll need to:

1. Add the CSP middleware
2. Configure allowed script sources

See the [Documentation](https://templui.io/docs/how-to-use) for details.

## Prerequisites

Before starting, ensure you have these tools installed:

- Go (1.20 or later)
- Templ
- Tailwind CSS CLI (Standalone)
- Make (for development scripts)
- Docker (optional, for deployment)

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
   Since we're using TemplUI as a package, you need to configure Tailwind to process its components:

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

## Development

Start the development server with hot reload:

```bash
make dev
```

Your application will be running at [http://localhost:7331](http://localhost:7331)

## Deployment

This template includes a production-ready Dockerfile for easy deployment:

```bash
# Build the image
docker build -t templui-app .

# Run the container
docker run -p 8090:8090 templui-app
```

Your application will be available at `http://localhost:8090`

## Learn More

- [TemplUI Documentation](https://templui.io/docs/how-to-use)
- [Available Components](https://templui.io/docs/components)
- [Theme Customization](https://templui.io/docs/themes)

## Contributing

Issues and pull requests are welcome! Please read our [contributing guidelines](https://github.com/axzilla/templui/blob/main/CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
