package main // Declare the main package

import (
	"fmt"           // Import the fmt package for formatted I/O
	"html/template" // Import the html/template package for rendering HTML templates
	"log"           // Import the log package for logging errors
	"net/http"      // Import the net/http package for HTTP server and client
)

func main() {
	// Serve static files from the static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	// Handle requests to the root URL by calling the home function
	http.HandleFunc("/", home)
	// Handle requests to the /projects URL by calling the projects function
	http.HandleFunc("/projects", projects)
	// Handle requests to the /skills URL by calling the skills function
	http.HandleFunc("/skills", skills)
	// Handle requests to the /about URL by calling the about function
	http.HandleFunc("/about", about)
	// Handle requests to the /contact URL by calling the contact function
	http.HandleFunc("/contact", contact)

	// Print a message to the console indicating the server is running
	fmt.Println("🚀 Server is running on http://localhost:8080")
	// Start the server on port 8080 and listen for incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil { // Check for errors when starting the server
		log.Fatal(err) // Log the error and exit if the server fails to start
	}
}

// home function handles requests to the home page
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html") // Render the home.html template
}

// projects function handles requests to the projects page
func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html") // Render the projects.html template
}

// skills function handles requests to the skills page
func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html") // Render the skills.html template
}

// about function handles requests to the about page
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html") // Render the about.html template
}

// contact function handles requests to the contact page
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html") // Render the contact.html template
}

// renderTemplate function renders an HTML template based on the provided name
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl) // Parse the specified template file
	if err != nil {                                    // Check for errors during template parsing
		http.Error(w, err.Error(), http.StatusInternalServerError) // Send an error response
	}
	t.Execute(w, nil) // Execute the template and write it to the ResponseWriter
}
