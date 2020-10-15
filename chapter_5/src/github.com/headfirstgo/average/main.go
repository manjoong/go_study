package main

import(
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main(){
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers{
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("평균: %0.2f\n", sum/sampleCount)
}