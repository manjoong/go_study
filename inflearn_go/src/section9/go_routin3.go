//고루틴 기초1
package main
import "fmt"
import "time"
import "math/rand"
import "runtime"

func exe(name int){
	r := rand.Intn(100)
	fmt.Println(name, " Start", time.Now())
	for i:=0; i<100; i++{
		fmt.Println(name, " >>>>> ", r, i)
	}
	fmt.Println(name, " END", time.Now())

}


func main(){
	//고루틴
	//멀티 코어(다중 cpu) 최대한 활용
	runtime.GOMAXPROCS(runtime.NumCPU()) //현 시스템의 cpu 코어 개수 반환 후 설정
	
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0))

	fmt.Println("Main ROutin", time.Now())

	for i :=0; i<100; i++{
		go exe(i) //고루틴 100개 생성

	}

	time.Sleep(5 * time.Second)
	fmt.Println(" END", time.Now())


}
