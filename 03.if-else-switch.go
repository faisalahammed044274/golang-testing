package main

import "fmt"

func main() {
age := 19;

if age >= 18 {
	fmt.Println("You are an adult....");
}else if(age <= 10){
	fmt.Println("You are a child....");
} else {
	fmt.Println("You are a teenager....");
}
}