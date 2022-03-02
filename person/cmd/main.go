package main


import (
	//"person/model"
   p "person/model"
	"fmt"
)
//在主函数中创建结构体的数值
func main(){
	fmt.Println("---") //---
	//包的名字，结构体的名字
	//p1 := model.Person{"坚持",22,"女"} //{坚持 22 女}
	//fmt.Println(p1)
	//fmt.Println(p1.Name,p1.Age,p1.Sex) //坚持 22 女
	//别名
	//p1 := p.Person{"梦想",18,"女"}
	//fmt.Println(p1)
	//Age改为私有age //.\main.go:20:26: implicit assignment of unexported field 'age' in model.Person literal
	p1 := p.Person{"梦想",18,"女"}
	fmt.Println(p1)
}
//CreateFile main.go: The system cannot find the file specified.
