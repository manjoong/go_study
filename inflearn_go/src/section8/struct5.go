package main	

import "fmt"
import "reflect"

type car struct{
	name string "차량명"
	color string "색상"
	company string "제조사"
}

func main()  {

	tag:= reflect.TypeOf(car{})

	for i:=0; i< tag.NumField(); i++{
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
