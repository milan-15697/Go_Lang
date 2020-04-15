package main

import "fmt"

func Factorial(num int){
	fact_calc:=1
	for i:=num; i>0; i--{
		fact = fact*i
	}
	fmt.Printf("Factorial of %d is %d = ",num,fact_calc)
}

func main(){
	var no int
	fmt.Scan(&no)
	Factorial(no)
}
