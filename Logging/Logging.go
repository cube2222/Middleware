package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.Handle("/", Log(HelloWorldHandler{}))
	http.ListenAndServe(":3000", nil)
}

type HelloWorldHandler struct {
}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!!!")
}

type LoggingMiddleware struct {
	wrappedHandler http.Handler
}

func (h LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling client with ip ", r.RemoteAddr)
	h.wrappedHandler.ServeHTTP(w, r)
	fmt.Println("Client with ip ", r.RemoteAddr, "handled.")
}

func Log(h http.Handler) LoggingMiddleware {
	return LoggingMiddleware{h}
}
