package main
import "fmt"

func calmDown(){
	recover()
}

func freakOut() {
	defer calmDown()
	panic("OH no")
}

func main(){
	freakOut()
	fmt.Println("exit nommally")
}