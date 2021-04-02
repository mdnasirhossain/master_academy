package main

import "fmt"

type Doctor struct{
   Name string
   Education string
   Age int
   Salary float32

}

//method
func (d Doctor)Speak(){
fmt.Println(d.Name,"can speak")
}

//method
func (d Doctor)getName() string{
return d.Name
}

//method
func (d Doctor)getSalaryInfo() float32{
return d.Salary
}

func main(){
//var d= Doctor { "Abir", "MBBS", 30, 50000.00, }
  //var d Doctor {Name:"Abir", Education:"MBBS", 
                 //Age: 30, Salary: 50000.00,}

//var d=Doctor{Name:"Abir", Education:"MBBS", 
  //         Age: 30, Salary: 50000.00,}


  var d Doctor
  d.Name="Abir"
  d.Education= "BDS"
  d.Age=30
  d.Salary=50000.00

 fmt.Println (d.getName())
 fmt.Println (d.getSalaryInfo())
  
}