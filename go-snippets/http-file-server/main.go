package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// serve files from the ./tmp directory, but alias it with /fs/, however,
	// strip the /fs/ from the path is it is only the alias
	http.Handle("/fs/", http.StripPrefix("/fs/", http.FileServer(http.Dir("./tmp"))))
	// serve files directly from the root of this server.
	http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/pdf", servePDF)
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	_, _ = io.WriteString(w, `
alias rewrite
<img src="/fs/dottics-logo.png"/>
directly from root
<img src="/dottics-logo.png"/>
from root in a nested folder structure
<img src="/tmp/dottics-logo.png"/>
`)
}
