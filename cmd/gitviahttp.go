package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mysticmode/gitviahttp"
)

func main() {
	var (
		isServerMode bool
		port         string
		repoDir      string
	)

	flag.BoolVar(&isServerMode, "server", false, "Specify true for the server mode else it will run in CLI mode")
	flag.StringVar(&port, "port", "8080", "Specifying the port where gitviahttp should run")
	flag.StringVar(&repoDir, "directory", ".", "Specify the directory where your repositories are located")

	flag.Parse()

	if isServerMode {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			gitviahttp.Context(w, r, ".")
		})
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else {
		fmt.Println("Hello, from CLI :)")
	}
}
