package main

import "fmt"

func main() {
  // Print(): function prints its arguments with their default format.
  var k, h string = "Hello", "Universe"
  var kh string = "Golang with HKimhab"
  var num, num2 = 22, 92
  fmt.Print(k, "\n")
  fmt.Print(h, "\n")
  fmt.Print(kh, "\n")
  fmt.Print("-------", "\n")
  fmt.Print(k, "\n", h, "\n", kh, "\n")
  fmt.Print(k,h)
  fmt.Print(num, num2)
	fmt.Printf("You are learning go output with HKimhab")
	
	/*  Println(): The Println() function is similar to Print() with 
    	the difference that a whitespace is added between the arguments, 
    	and a newline is added at the end: 
	*/
	fmt.Println() // fmt.Print("\n")
	fmt.Println("---Println()---")
	fmt.Println(k, h)
	fmt.Println(num, num2)
	fmt.Println(num)
	fmt.Println(num2)
	
	/*  The Printf() function first formats its argument based on the given 
	      formatting verb and then prints them.
        Here we will use two formatting verbs:
        + %v is used to print the value of the arguments
        + %T is used to print the type of the arguments
	*/
	
	fmt.Printf("num has value: %v and type: %T\n", num, num)
	fmt.Printf("kh has value: %v and type: %T\n", kh, kh)
	
}