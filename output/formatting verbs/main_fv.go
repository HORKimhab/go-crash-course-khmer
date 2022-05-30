package main

import "fmt"

func main() {
  var h = 22.2 
  var kh = 92
  var str = "Hello Universe"
  var txt = "HKimhab"
  var ds = true
  var dsk = false
  var i = 12.925
  
  fmt.Printf("General Formating Verbs\n")
  fmt.Printf("%v\n", h)
  fmt.Printf("%#v\n", str)
  fmt.Printf("%T\n", str)
  fmt.Printf("%v%%\n", h)
  
  fmt.Println("-----------------")
  fmt.Printf("Integer Formating Verbs\n")
  fmt.Printf("%b\n", kh)
  fmt.Printf("%d\n", kh)
  fmt.Printf("%+d\n", kh)
  fmt.Printf("%o\n", kh)
  fmt.Printf("%O\n", kh)
  fmt.Printf("%x\n", kh)
  fmt.Printf("%X\n", kh)
  fmt.Printf("%#x\n", kh)
  fmt.Printf("%4d\n", kh)
  fmt.Printf("%-4d\n", kh)
  fmt.Printf("%04d\n", kh)
  
  fmt.Println("-----------------")
  fmt.Printf("String Formating Verbs\n")
  fmt.Printf("%s\n", txt)
  fmt.Printf("%q\n", txt)
  fmt.Printf("%8s\n", txt)
  fmt.Printf("%-8s\n", txt)
  fmt.Printf("%x\n", txt)
  fmt.Printf("% x\n", txt)
  
  fmt.Println("-----------------")
  fmt.Printf("Booleang Formating Verbs\n")
  fmt.Printf("%t\n", ds)
  fmt.Printf("%t\n", dsk)
  
  fmt.Println("\n-----------------")
  // i = 12.925
  fmt.Printf("Float Formating Verbs\n")
  fmt.Printf("%e\n", i)
  fmt.Printf("%f\n", i)
  fmt.Printf("%.2f\n", i)
  fmt.Printf("%6.2f\n", i)
  fmt.Printf("%7.2f\n", i)
  fmt.Printf("%g\n", i)
}