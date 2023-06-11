package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Print("THis is some text")
	fmt.Print("this is some more text")
	sayHelloworld("Hello world from function")
	// var whattosay string
	// whattosay = "Hello world from variable"
	whattosay := "Hello world from variable"
	sayHelloworld(whattosay)
}

func sayHelloworld(whattosay string){
	fmt.Println(whattosay)
}