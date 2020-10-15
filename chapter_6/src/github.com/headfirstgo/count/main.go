package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"os"
)

func main(){
	lines, err := datafile.GetStrings("vites.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := make(map[string]int)
	for _, line := range lines{
		count[line]++
	}
	for name, count := range counts {
		fmt.Printf("%s : %d\n", name, count)
	}
}