package main

import (
	"fmt"
	"github.com/DmitriyVTitov/size"
	"log"
	"os"
	"time"
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
	fmt.Println(size.Of(b))
	fmt.Println(size.Of(fl))
	fmt.Println(size.Of(str))
	fmt.Println(size.Of(t))

	s1 := Celsius{
		temp:      1.23284,
		time:      time.Time{},
		active:    false,
		source:    "dwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgDdwE4xfGuLfccUy3RCM7DBFkQtfub7pgD",
		footprint: 0.33,
	}

	s2 := CelsiusReorder{
		temp:      1.23284,
		time:      time.Time{},
		active:    false,
		source:    "d",
		footprint: 0.33,
	}
	fmt.Println("===========")
	fmt.Println(size.Of(s1))
	fmt.Println(size.Of(s2))
}

func getMeasurementsFile() (*os.File, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}

	return f, nil
}
