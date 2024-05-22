package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

type TestCase struct {
	name   string
	input  *os.File
	result Result
}

/*
 1. Test Recursive
    go test -run TestOPTRecursiveSubTest

1a)Test Iterative

		go test -run TestOPTIterativeSubTest
	 2. check coverage
	    go test -coverprofile=coverage.out
	 3. check coverage by function:
	    go tool cover -func=coverage.out
	 4. check coverage:
	    go tool cover -html=coverage.out
*/
var r = Result{
	avg:     "-17.87",
	minTemp: "-50.60",
	maxTemp: "15.40",
	fK:      "-146.16",
}

func TestOPTSolution1(t *testing.T) {
	f, err := getMeasurementsFile()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	testCase := TestCase{
		name:   "basic test",
		input:  f,
		result: r,
	}

	t.Run(testCase.name, func(t *testing.T) {
		result := s1(testCase.input)
		assert.Equal(t, testCase.result, result)
	})
}

func TestOPTSolution2(t *testing.T) {
	f, err := getMeasurementsFile()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	testCase := TestCase{
		name:   "basic test",
		input:  f,
		result: r,
	}

	t.Run(testCase.name, func(t *testing.T) {
		result := s2(testCase.input)
		assert.Equal(t, testCase.result, result)
	})
}

func TestOPTSolution3(t *testing.T) {
	f, err := getMeasurementsFile()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	testCase := TestCase{
		name:   "basic test",
		input:  f,
		result: r,
	}

	t.Run(testCase.name, func(t *testing.T) {
		result := s3(testCase.input)
		assert.Equal(t, testCase.result, result)
	})
}
