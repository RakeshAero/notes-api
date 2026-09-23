package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	
	"github.com/RakeshAero/notes-api/internal/database"
)

func main(){

	//Open the DB Connection
	db, err := database.Open("data/notes.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	
	//Migrate the DB
	if err := database.Migrate(db); err != nil {
		panic(err)
	}


	//Start http server
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request){
		writeJSON(w,http.StatusOK,map[string]string{
			"status" : "ok",
		})
	})
	fmt.Println("Server started at localhost:8081")
	err = http.ListenAndServe("localhost:8081", mux)
	if err != nil {
		panic(err)
	}
}





// writeJSON helper
func writeJSON(w http.ResponseWriter, status int, data any) error {
	// Set the header before writing the status code or body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	// Encode the data directly into the response writer
	return json.NewEncoder(w).Encode(data)
}