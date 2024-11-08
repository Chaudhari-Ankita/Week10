package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./templates")) // created a variable 'fs' that points to the "./templates" directory.
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":80", nil))
}
