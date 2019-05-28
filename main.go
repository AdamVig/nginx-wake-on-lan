package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/mdlayher/wol"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleMissingParam(w http.ResponseWriter, name string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "missing query parameter '%s'", name)
}

func wake(w http.ResponseWriter, rawMac string, ip string) {
	log.Printf("waking computer with MAC '%s' using broadcast IP '%s'\n", rawMac, ip)

	c, err := wol.NewClient()
	if err != nil {
		log.Println("failed to create wake on LAN client", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mac, err := net.ParseMAC(rawMac)
	if err != nil {
		log.Println("failed to parse MAC address", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.Wake(ip+":9", mac)
	if err != nil {
		log.Println("failed to wake computer", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func wakeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if _, ok := r.Form["mac"]; !ok {
		handleMissingParam(w, "mac")
	} else if _, ok := r.Form["ip"]; !ok {
		handleMissingParam(w, "ip")
	} else {
		mac := r.Form.Get("mac")
		ip := r.Form.Get("ip")
		wake(w, mac, ip)
	}
}

func main() {
	http.HandleFunc("/wake", wakeHandler)
	http.HandleFunc("/", pageHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
