package main

import (
	"fmt"
	"log"
	"net/http"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleMissingParam(w http.ResponseWriter, name string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "missing query parameter '%s'", name)
}

func wakeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if _, ok := r.Form["mac"]; !ok {
		handleMissingParam(w, "mac")
	} else if _, ok := r.Form["ip"]; !ok {
		handleMissingParam(w, "ip")
	} else {
		fmt.Fprintf(w, "waking computer with MAC '%s' using broadcast IP '%s'", r.Form.Get("mac"), r.Form.Get("ip"))
	}
}

func main() {
	http.HandleFunc("/wake", wakeHandler)
	http.HandleFunc("/", pageHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
