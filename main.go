package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	dir           = flag.String("d", ".", "directory to serve")
	port          = flag.Int("p", 8080, "port to serve on")
	enableLogging = flag.Bool("enablelogging", true, "used to enable / disable logging")
	logFavicon    = flag.Bool("logfavicon", false, "whether to log favicon requests")
)

func main() {
	flag.Parse()
	log.Printf("- Serving '%s' on port %d\n", *dir, *port)
	handler := loggingHandler{http.FileServer(http.Dir(*dir))}
	log.Fatalln("-", http.ListenAndServe(":"+strconv.Itoa(*port), handler))
}

type loggingHandler struct {
	h http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusWriter := &statusStoringResponseWriter{ResponseWriter: w}
	h.h.ServeHTTP(statusWriter, r)
	switch {
	case !*enableLogging:
	case !*logFavicon && r.RequestURI == "/favicon.ico":
	default:
		log.Printf("- %d %s %s%s", statusWriter.status, r.Method, r.Host, r.RequestURI)
	}
}

type statusStoringResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusStoringResponseWriter) WriteHeader(n int) {
	w.ResponseWriter.WriteHeader(n)
	// Attempt to handle multiple WriteHeader() calls
	if w.status == 0 {
		w.status = n
	}
}
