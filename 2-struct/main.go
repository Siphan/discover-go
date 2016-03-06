package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (u *user) userName() { // method to print user name
	fmt.Printf("Hello %s\n", u.Name)
}

func parseTime() time.Time {
	value := "March 7, 1917"    // date that we want to use
	layout := "January 2, 2006" // format that the date needs to be converted to using Go reference date
	t, _ := time.Parse(layout, value)
	return (t) // return time is converted format 1917-03-07 00:00:00 +0000 UTC
}

func age(DOB string) int { // method to calculate age based on input DOB
	birthDate := parseTime()                      // storing result of method parseTime into a variable that we can use
	years := time.Now().Year() - birthDate.Year() // gets the number of years since input birth
	return (years)
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("%s who was born in %s would be %d years old today\n", u.Name, u.City, age(u.DOB)) // calling variable, variable, method
}
