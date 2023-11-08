package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Task 1
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

// Task 2
func genPass(password string, level string) {
	rand.Seed(time.Now().UnixNano())

	if level == "Low" {
		randomInt := rand.Intn(1000) + 10000
		fmt.Print(password, randomInt)
	} else if level == "Med" {
		randomInt := rand.Intn(1000) + 1000
		fmt.Print(password, randomInt, "@")
	} else if level == "Strong" {
		hurufPertama := password[0:1]
		hurufKedua := strings.ToUpper(password[1:2])
		hurufKetiga := password[2:3]
		hurufKeempat := strings.ToUpper(password[3:4])
		kata := hurufPertama + hurufKedua + hurufKetiga + hurufKeempat
		randomInt := rand.Intn(1000) + 1000
		fmt.Print(kata, "@", randomInt)
	}
}

func main() {
	//Task 1
	tinggi := 5
	printSegitiga(tinggi)

	// Task 2
	genPass("abcd", "Strong")
}
