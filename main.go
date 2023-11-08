package main

import "fmt"

func printSegitiga(tinggi int) {
	k := 0

	for i := tinggi; i >= 1; i-- {
		for l := 0; l < k; l++ {
			fmt.Print(" ")
		}

		k++

		for j := (i*2 - 1); j >= 1; j-- {
			fmt.Print("*")
		}

		fmt.Print("\n")
	}
}

func main() {
	tinggi := 5
	printSegitiga(tinggi)
}
