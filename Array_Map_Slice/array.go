package main

import "fmt"

func main() {

	// var students [3]string
	// students[0] = "Nasir"
	// students[1] = "Hossain"
	// students[2] = "Shuvo"

	// fmt.Println(students)
	// fmt.Println(len(students))
	// fmt.Println(students[0])
	// fmt.Println(students[1])
	// fmt.Println(students[2])

	students := [...]string{"Nasir", "Hossain", "Shuvo", "Josim"}
	fmt.Println(students)
	fmt.Println(len(students))
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
	fmt.Println(students[3])

}
