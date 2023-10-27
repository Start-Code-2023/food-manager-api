package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// HandlerDefault is a handler for the /default endpoint.
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to indicate that the response contains JSON data
	w.Header().Add("content-type", "application/json")

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodGet {
		http.Error(w, "The HTTP method used is not GET", http.StatusMethodNotAllowed)
		return
	}
}

// loadFile takes a filename as a string and returns the contents of the file as a string.
// Returns: a string, or an error and an empty string.
func loadFile(filename string) (string, error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
