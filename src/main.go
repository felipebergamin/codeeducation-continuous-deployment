package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(message string) string {
	return fmt.Sprintf("<b>%s<b>", message)
}

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
