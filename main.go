package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	dir  = flag.String("d", ".", "directory to serve")
	port = flag.Int("p", 8080, "port to serve on")
)

func main() {
	flag.Parse()
	fmt.Printf("Serving '%s' on port %d\n", *dir, *port)
	perror(http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(*dir))))
}

func perror(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
