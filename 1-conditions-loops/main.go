package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Printf("I am a student of the %s\n", school)
	} else {
		fmt.Printf("I am not a student of the %s\n", school)
	}
	beautifulWeather := true
	if beautifulWeather { // boolean variable is already true, so this is enough
		fmt.Printf("It's a beautiful weather!\n")
	}

	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"} // slice with 3 strings
	for i := 0; i < 3; i++ {
		fmt.Printf("%v is a founder\n", holbertonFounders[i]) // slice depends on i, each string will be printed
	}
}
