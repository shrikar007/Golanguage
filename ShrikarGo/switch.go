package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	switch name {
	case "Ajitem":
		fmt.Println("Hello, Ajitem!")
	case "Babulal":
		fmt.Println("Hello, Babulal!")
		  break
              
	case "Rohit":
		fmt.Println("Hello, Rohit!")
                fallthrough
   
	case "Ankit":
		fmt.Println("Hello, Ankit!")
	case "Soni":
		fmt.Println("Hello, Soni!")
		fallthrough
	case "Kalpashree":
		fmt.Println("Hello, Kalpashree!")
	default:
		fmt.Println("Hello, " + name + " you are not a part of GoLang Training. Would you like to join us?")
	}
}
