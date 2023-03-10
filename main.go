package main

import "net/http"

func main() {
	// start a server on port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}
