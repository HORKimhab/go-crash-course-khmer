// You can edit this code!
// Click here and start typing.
package main

// fmt to running go project
import "fmt"

/*
	main() is from package main
	learning Golang with HKimhab
	----
	int- stores integers (whole numbers), such as 123 or -123
	float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
	string - stores text, such as "Hello World". String values are surrounded by double quotes
	bool- stores values with two states: true or false

	Declare variable
	1. var variablename type = value
	2. variablename := value // type is inferred
	3. var variablename = value // type is inferred
*/
// name11 :=  12
var name string = "HKimhab" // type is string 

func main() {
	/* var name string = "HKimhab" // type is string  */
	name1 := 22 // type is inferred 
	var name2 = true // type is inferred 

	// Declare without value 
	var a1 string 
	var a2 bool
	var a3 int 

	// Declare datatype and declare  value 
	var user_name string 
	user_name = "HKimhab Coding"
	

	fmt.Println("You are learning declare variable in golang with HKimhab")

	// use variable name 
	fmt.Println(name)
	// Output name 1
	fmt.Println(name1)
	// Output name2 
	fmt.Println(name2)

	fmt.Println("Output variable without declare value")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(user_name)

	// fmt.Println(name11)
}