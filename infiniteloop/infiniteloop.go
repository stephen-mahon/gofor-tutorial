package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	findRandomeNumber(rand.Intn(100))

}

func findRandomeNumber(randomNumber int) {
	count := 1
	numberFound := false

	for {
		number := rand.Intn(1000)
		if number == randomNumber {
			numberFound = true
			break
		}
		count++
	}

	if numberFound {
		fmt.Printf("Number #%v found afer %v attempt(s)", randomNumber, count)
	}
}
