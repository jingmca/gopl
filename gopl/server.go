package gopl

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//App will show
func App() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "%s %s %s\n", r.RemoteAddr, r.URL.Path, "200")
		fmt.Fprintf(w, "URL.Path = %s", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
