package main

import "os"

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}
//command-line-arguments
//runtime.main_main·f: function main is undeclared in the main package
