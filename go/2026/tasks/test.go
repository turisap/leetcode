package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	a := Product{ID: 1, Name: "Notebook", Price: 9.98}
	if a == (Product{ID: 1, Name: "Notebook", Price: 9.99}) {
		fmt.Println("match")
		return
	}
	fmt.Println("no match")
}
