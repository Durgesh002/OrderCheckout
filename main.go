package main

import "fmt"

func main() {
	fmt.Printf("%50v\n", "Welcome to the Gopher store")
	fmt.Println()
	mybill := createbill()
	prompts(mybill)

}
