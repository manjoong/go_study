package main


import(
	"fmt"
	"os"
	"strconv"
	"log"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0	
	for _, number := range numbers{
		sum += number
	}
	return sum / float64(len(numbers))
}

func main(){
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments{
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers =append(numbers, number)
	}
	//가변 인자를 받는 함수에 대해서 슬라이드를 넣을때, ...을 넣어줘야 한다.
	fmt.Printf("평균: %0.2f\n", average(numbers...))
}