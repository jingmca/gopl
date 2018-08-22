package gopl

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

//App will show
func App() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "%s %s %s\n", r.RemoteAddr, r.URL.Path, "200")
		fmt.Fprintf(w, "URL.Path = %s", r.URL.Path)
		mu.Lock()
		count++
		mu.Unlock()

	})

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintf(w, "Requests:%3d", count)
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "123")
	})

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
		}
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form [%q] = %q\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
