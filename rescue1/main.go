package main

//import "fmt"
func main(){

}
func throwsPanic(f func()) (b bool) {
defer func() {
if x := recover(); x != nil {
b = true
}
}()
f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
return
}

