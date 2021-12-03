package files

import (
	"log"
	"os"
	"strings"
)

// ReadLines takes a file path and returns the lines of text
// in that file.
func ReadLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
