package main

import "fmt"

func main() {

	x := make(map[string]string)
	x["name"] = "Nasir"
	x["age"] = "24"
	x["address"] = "Pabna"

	delete(x, "name")

	fmt.Println(x)
	fmt.Println(x["name"])
	fmt.Println(x["age"])
	fmt.Println(x["address"])

}