package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

type subscriber struct {
	name string
	rate float64
	active bool
}

func printInfo(s subscriber){
	fmt.Println("Name: ", s.name)
	fmt.Println("RATE: ", s.rate)
	fmt.Println("Active: ", s.active)
}



func main() {
	// var subscriber1 subscriber
	// subscriber1.name = "manjoong"
	// subscriber1.rate = 15.5
	// subscriber1.active = true

	// printInfo(subscriber1)

	// subscriber1.name = "manjoong"
	// fmt.Println("Name: ", subscriber1.name)

	// var subscriber2 subscriber
	// subscriber2.name ="joosun"
	// fmt.Println("Name: ", subscriber2.name)

	var s magazine.Subscriber
	s.rate = 4.99
	fmt.Println(s.rate)
}



