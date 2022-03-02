//package  hello //go run: cannot run non-main package
//package adddemo //go run: cannot run non-main package


package main
import (
	"adddemo/modules"
	"fmt"
)
func main()  {
	//包名.方法名
	/*
	type Student struct {
		Name string
	}*/
	stu := modules.Student{
		Name :"坚持",
	}
	//fmt.Printf(stu)
	//.\hello.go:10:12: cannot use stu (type modules.Student) as type string in argument to fmt.Printf
	fmt.Println(stu)
}
