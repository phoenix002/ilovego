package main

func ReadWrite() bool {
	file.Open("test1.txt")
	defer file.Close()
	if failureX {
		return false
	}
	if failureY {
		return false
	return true
}
func main(){

}

