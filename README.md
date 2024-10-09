# Goilerplate Quickstart - Jump Start Your Go & Templ UI Project

This repository provides a quickstart template for projects using Goilerplate, a modern UI component library for Go and Templ. Get a head start on your project with a pre-configured environment and example setup.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (latest version recommended)
- [Node.js and npm](https://nodejs.org/) (for running npx with Tailwind CSS)
- `make` (usually pre-installed on macOS and Linux; Windows users might need to install it separately)

## Quick Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/axzilla/goilerplate-quickstart.git
   cd goilerplate-quickstart
   ```

2. **Install Go Dependencies**

   ```bash
   go mod tidy
   ```

3. **Install Templ**

   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

4. **Start the Development Server**

   ```bash
   make dev
   ```

   This command will use npx to run Tailwind CSS, so no separate npm install is needed.

5. **View Your Project**
   Open your browser and navigate to [http://localhost:7331](http://localhost:7331)

## Development Workflow

- The `make dev` command starts a development server with hot-reloading for Go, Templ, and Tailwind CSS.
- Edit files in the `ui` directory to modify layouts and pages.
- Tailwind CSS styles can be adjusted in `assets/css/input.css`.
- The main application logic is in `main.go`.

## Customization

- Add new pages in `ui/pages/` and new layouts in `ui/layouts/`.
- Modify Tailwind configuration in `tailwind.config.js`.
- Extend the `Makefile` for additional build processes if needed.

## Using Goilerplate Components

This quickstart is pre-configured to use Goilerplate. To use Goilerplate components:

1. Import Goilerplate in your Templ files:

   ```go
   import "github.com/axzilla/goilerplate/pkg/components"
   ```

2. Use components in your Templ templates:
   ```go
   @components.Button(components.ButtonProps{Text: "Click me"})
   ```

Refer to the [Goilerplate documentation](https://github.com/axzilla/goilerplate) for a full list of available components and their usage.

## Troubleshooting

- If Templ files are not updating, ensure the Templ CLI is installed and run `templ generate` manually.
- For Tailwind CSS issues, check that the `input.css` file is correctly linked in your HTML.

## Contributing

Contributions to improve this quickstart template are welcome. Please feel free to submit issues or pull requests.

## License

This quickstart project is open-source and available under the [MIT License](LICENSE).

---

Happy coding with Goilerplate! We're excited to see what you build. If you create something cool, consider sharing it with the community!
