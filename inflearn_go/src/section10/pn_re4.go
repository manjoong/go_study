//panic & recover (4)

package main

import "fmt"
import "os"

func fileOpen(filename string){
	defer func(){
		if r := recover(); r !=nil {
			fmt.Println("Error message", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else{
		fmt.Println(f.Name())
	}

	defer f.Close()

	
}
func main(){
	//go panic 함수


	fileOpen("underfined.txt")
	

	fmt.Println("end main")

}