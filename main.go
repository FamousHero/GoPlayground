package main

import "fmt" 
// import "class"

func main(){
	fmt.Println("Hello world")
	p1 := Person{name: "Johnny", age: 29, alive: false}
	fmt.Println(p1.getName())

	s1 := "Hi "
	s2 := "there"
	s_final := concat(s1, s2)
	fmt.Println(s_final)

	x := 4
	fmt.Println(square(x))
}