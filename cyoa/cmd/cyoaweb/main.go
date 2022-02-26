// go build cmd/cyoaweb/main.go
// go run cmd/cyoaweb/main.go
// go run cmd/cyoaweb/main.go --help

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arnxv0/golang-projects/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "port to start web server on")
	file := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in '%s'.\n", *file)

	fileReader, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(fileReader)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", story) // %+v prints the struct as a string

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
