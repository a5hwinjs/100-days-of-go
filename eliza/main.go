package main

import ("fmt"
	"myapp/doctor")

func main() { 

	var whattosay string
	whattosay = doctor.Intro()
	fmt.Println(whattosay)

}