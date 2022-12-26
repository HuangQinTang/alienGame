package main

import (
	"log"
	"net/http"
	"testing"
)

func Test_Web(t *testing.T) {
	if err := http.ListenAndServe(":8889", http.FileServer(http.Dir("."))); err != nil {
		log.Fatal(err)
	}
}
