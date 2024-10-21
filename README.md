# Simple Go Web Server Tutorial

This tutorial demonstrates how to build a simple Go web server that serves static files and renders HTML templates for different routes, such as Home, Projects, Skills, About, and Contact pages.

## Prerequisites

- Install [Go](https://golang.org/doc/install)
- Basic understanding of Go
- Familiarity with HTML


## Steps

1. **Setup Project Structure:** Create the directories `static` and `templates`.
2. **Static Files:** Add your static files (e.g., CSS, images) inside the `static` folder.
3. **HTML Templates:** Add your HTML templates inside the `templates` folder (e.g., home.html, projects.html, etc.).
4. **Running the Server:**
   - Navigate to the project directory.
   - Run the server using:
     ```bash
     go run main.go
     ```
   - The server will start at `http://localhost:8080`.

This simple Go web server allows you to serve static files and HTML templates. It is a great starting point for building a personal portfolio or project website using Go.
