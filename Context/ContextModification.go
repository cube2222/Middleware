package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"
)

func main() {
	http.Handle("/", Name(HelloWorldHandler{}))
	http.ListenAndServe(":3000", nil)
}

type HelloWorldHandler struct {
}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if name, ok := r.Context().
		Value("name").(string); ok == true && len(name) > 0 {
		fmt.Fprintln(w ,"Hello ", name, "!!!")
	} else {
		fmt.Fprintln(w, "Hello nameless client.")
	}
}

type NamingMiddleware struct {
	wrappedHandler http.Handler
}

func (h NamingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.wrappedHandler.ServeHTTP(
		w,
		r.WithContext(context.WithValue(r.Context(), "name", "Ben")),
	)
}

func Name(h http.Handler) NamingMiddleware {
	return NamingMiddleware{h}
}