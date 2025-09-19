package main

import (
	"fmt"
	"net/http"
	"strings"
)

func ShowJob(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "showjob endpointgo  ")
	parts := strings.Split(r.URL.Path, "/")
	length := len(parts)
	id := parts[length-1]
	fmt.Fprintf(w, id)
	fmt.Println(length)
	fmt.Println(parts)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.URL.Host)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RequestURI())

}

/*

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Handler that extracts an integer ID from the URL path
func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Example URL: /item/123
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "You requested item with ID: %d\n", id)
}

func main() {
	http.HandleFunc("/item/", getByIDHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


*/
