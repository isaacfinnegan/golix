package main 
	import (
		"fmt"
		"html"
		"log"
		"net/http"
		"github.com/gorrilla/mux"
	)

	func main() {
		router := mux.NewRouter().StrictSlash(true)
		router.HanldeFunc("/",Index)
		log.Fatal(http.ListenAndServe(":8080",router))

	}

	func Index(w http.ResponseWriter, r
		*http.Request) {
			fmt.Fprintf(w, "Hello, %q",html.EscapeString(r.URL.Path))
		}
	
