package main

// fmt to running go project
import "fmt"

// Outside function of constant
// Untyped == inferred: type of vari refer to value
const PI = 3.14
// Type is string 
const HTTP = "http"

func main(){
	// constants = unchangeble and read-only 
	// syntax of constant: const NAMEOFCONS type = value
	/* --- Rule of Constants ---
		+ Constant names follow the same naming rules as variables
		+ Constant names are usually written in uppercase letters (for 		easy identification and differentiation from variables)
		+ Constants can be declared both inside and outside of a function
	*/
	/* Type of Constant: 
		1. Typed const
		2. UnTyped const
	*/

	// Typed const 
	const NAME string = "Golang"
	fmt.Println("---Learning Constant Go---")
	fmt.Println(NAME)
	fmt.Println(PI)

	// Declare Const as multi in block 
	const (
		CAPITAL = "Phnom Penh"
		CURRENCY = "KHR"
		DATEFORMAT = "dd//mm/yyyy"
	)
	fmt.Println("Const in block as multi var")
	fmt.Println(CAPITAL)
	fmt.Println(CURRENCY)
	fmt.Println(DATEFORMAT)

}