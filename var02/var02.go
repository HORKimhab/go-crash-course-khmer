package main

// fmt to running go project
import "fmt"

func main(){
	// Multi variable 
	var k1, kh, d int = 12, 22, 55

	// Multi without type 
	var k2, kh2 = 34, "Golang with HKimhab"
	
	// Multi with inferred 
	k3, d3 := "DSreykhuoch", 2022

	// c
	var (
		k4 int
		d4 int = 95
		ds string = "Hey Crush"
	)

	// Kh & KH 
	var Kh int  = 74
	var KH int 

	var k5 string = "Error"
	fmt.Println("Declare multi variable")
	fmt.Println("-------")
	fmt.Println(k1)
	fmt.Println(kh)
	fmt.Println(d)

	fmt.Println("Multi without type ")
	fmt.Println(k2)
	fmt.Println(kh2)

	fmt.Println("Multi with inferred")
	fmt.Println(k3)
	fmt.Println(d3)
	fmt.Println()

	fmt.Println("Multi with block")
	fmt.Println(k4)
	fmt.Println(d4)
	fmt.Println(ds)

	fmt.Println()
	fmt.Println(Kh)
	fmt.Println(KH)

	fmt.Println(k5)

	/* 
		Camel Case
		Each word, except the first, starts with a capital letter:
		myVariableName = "John"	

		Pascal Case
		Each word starts with a capital letter:
		MyVariableName = "John"

		Snake Case
		Each word is separated by an underscore character:
		my_variable_name = "John"
	*/
}