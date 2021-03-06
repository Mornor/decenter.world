// A simple redirect server that forwards any http:// request to https://decenter.world.
package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://decenter.world/", http.StatusFound)
	})
	addr := ":80"
	fmt.Printf("Serving redirect server on %s..\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
