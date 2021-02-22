//고루틴 기초1
package main
import "fmt"
import "time"

func main(){
	

	exe("t1") //가장 먼저 실행

	fmt.Println("Main routine start :", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second) // time.Second, Minute, Hour...
	fmt.Println("Main routine end :", time.Now())


}

func exe(name string){
	fmt.Println(name, " start: ", time.Now())
	for i:=0; i<1000; i++{
		fmt.Println(name, ">>>>>", i)
	}
	fmt.Println(name, " end : ", time.Now())
}




