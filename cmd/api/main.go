package main
import (
	// "fmt"
	"net/http"
	"encoding/json"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request){
		writeJSON(w,http.StatusOK,"")
	})

	http.ListenAndServe("localhost:8080", mux)
}

// writeJSON marshals the provided data to JSON, sets the content type, and writes the response.
func writeJSON(w http.ResponseWriter, status int, data any) error {
	// Set the header before writing the status code or body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	// Encode the data directly into the response writer
	return json.NewEncoder(w).Encode(data)
}