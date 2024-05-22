package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"
)

func main() {
	f, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	b := true
	t := time.Now()
	var fl float64 = 1.4
	str := "dwE4xfGuLfccUy3RCM7DBFkQtfub7pgD"
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Sizeof(fl))
	fmt.Println(unsafe.Sizeof(str))
}

func getMeasurementsFile() (*os.File, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}

	return f, nil
}
