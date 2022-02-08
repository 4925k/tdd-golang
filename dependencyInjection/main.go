package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "beb")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(GreeterHandler)))
}
