package main

import "fmt"

func update(a *int, name *string) {

	*a = *a + 5              
	*name = *name + " Nasir" 
	fmt.Println(a, name)
}

func main() {

	// var p *int //0xc000006028=0xc0000100a0
	// var x int  //0xc0000100a0=100 
	// //d := &doctor{"Shuvo", "BBA"}

	// fmt.Println(&p, &x)
	// fmt.Println(p, x) //zero values pointer=nil, integar= 0

	// x = 10
	// p = &x
	// *p = 100 

	// var p2 **int
	// p2 = &p
	// //*p2=p
	// //**p2==*p2->p->x=100

	// fmt.Println(x, p, *p, p2, *p2, **p2)

	x := 5
	name := "mostain"
	update(&x, &name)
	fmt.Println(&x, &name, x, name)

}
