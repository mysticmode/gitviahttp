/*

MIT License

Copyright (c) 2019 Nirmal Kumar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

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
