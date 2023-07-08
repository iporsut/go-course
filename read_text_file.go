package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	raw, err := os.ReadFile("note.txt")
	if err != nil {
		log.Fatal(err)
	}

	bWord := bytes.Fields(raw) // bWord is a slice of []byte
	fmt.Println(bWord)
}
