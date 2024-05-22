package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s4(f)
}

func getMeasurementsFile() (*os.File, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}

	return f, nil
}
