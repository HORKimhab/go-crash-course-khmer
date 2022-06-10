package main

import "fmt"

func main() {
  // signed Integer: can be negative or positive 
  // e.g -92, OR 92 | 12 
  // unsigned Integer: must be positive 
  // e.g 0 - 838 
  var kh int = 39949499
  
  // int8 = -128 to 127 
  var kh1 int8 = 125
  
  // unint32 = 4294967295
  var ds uint16 = 4294967295
  
  // uint, 
  
	fmt.Printf("Golang Integer\n")
	fmt.Printf("Type: %T, value: %v\n", kh, kh)
	fmt.Printf("Type: %T, value: %v\n", kh1, kh1)
	fmt.Printf("Type: %T, value: %v\n", ds, ds)
}