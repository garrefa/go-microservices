package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello world!")
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Ooooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", data)
}
