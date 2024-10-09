# Quickstart for Goilerplate - Modern UI Components for Go & Templ

Get started in just a few steps. Clone, run, and customize as you like!

## Steps to Run

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/axzilla/goilerplate-quickstart.git
   cd goilerplate-quickstart
   ```

2. **Install Dependencies**:
   Run the following command to install Go dependencies:

   ```bash
   go mod tidy
   ```

3. **Install Templ LSP**:
   To enable templating features, install the Templ Language Server Protocol (LSP) with:

   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

4. **Run Development Server**:
   Start the development server with:

   ```bash
   make dev
   ```

5. **Open in Browser**:
   Navigate to [http://localhost:7331](http://localhost:7331) to view the project.

6. **Customize & Experiment**:
   Feel free to modify the code, experiment with changes, and create something awesome. Break it, fix it, and make it your own!

## Requirements

- [Go](https://golang.org/dl/) for Go package management.
- [Templ](https://templ.guide/) for templating. Make sure the Templ LSP is installed with:
  ```bash
  go install github.com/a-h/templ/cmd/templ@latest
  ```
- `make` for running development commands.
- [Node.js and npm](https://nodejs.org/) for using `npx` with Tailwind CSS commands.

## License

This project is open-source. Feel free to use and modify the code as needed.

Happy hacking!
