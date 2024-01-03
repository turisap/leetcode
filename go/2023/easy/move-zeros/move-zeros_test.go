package main
import "testing"

        /*
        	1) Test Recursive
        	    go test -run TestMVZRecursiveSubTest
        	1a)Test Iterative
        	    go test -run TestMVZIterativeSubTest
        	2) check coverage
        	    go test -coverprofile=coverage.out
        	3) check coverage by function:
        	    go tool cover -func=coverage.out
        	4) check coverage:
        	    go tool cover -html=coverage.out
        */


        func TestMVZRecursiveSubTest(t *testing.T) {
        	for _, testCase := range testCases {
        		t.Run(testCase.name, func(t *testing.T) {
        			result := mvzRecursive()
        			if result != testCase.result {
        				t.Errorf("MVZ %s produces %d, got %d", testCase.s, testCase.result, result)
        			}
        		})
        	}
        }

        func TestMVZIterativeSubTest(t *testing.T) {
        	for _, testCase := range testCases {
        		t.Run(testCase.name, func(t *testing.T) {
        			result := mvzIterative()
        			if result != testCase.result {
        				t.Errorf("MVZ %s produces %d, got %d", testCase.s, testCase.result, result)
        			}
        		})
        	}
        }
        
