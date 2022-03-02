package model        //申明person包
type Person struct { //首字母大写代表公共的可以被其他包访问
	Name string //首字母大写代表外部能访问
	//Age int
	age int
	Sex string

}
