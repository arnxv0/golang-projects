// go build cmd/cyoaweb/main.go
// go run cmd/cyoaweb/main.go
// go run cmd/cyoaweb/main.go --help

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/arnxv0/golang-projects/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in '%s'.\n", *file)

	fileReader, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	var story cyoa.Story
	fileJSON := json.NewDecoder(fileReader)
	if err := fileJSON.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
