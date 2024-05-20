package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("./measurements.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s1(f)
}
