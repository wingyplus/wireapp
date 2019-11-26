package main

import "net/http"

type page struct {
}

func (p *page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world"))
}

func newPage() *page {
	return &page{}
}
