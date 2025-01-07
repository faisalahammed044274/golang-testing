package main

import "fmt"

func main() {
age := 20;
sex :="male"
sex = "female"

if age == 20 && sex == "male" {
	fmt.Println("You are a male of 20 years");
}elseif(age == 60 && sex == "male"){
	fmt.Println("You are a male of 60 years");
}elseif(age==20 && sex=="female"){
	fmt.Println("You are a female of 20 years");
	}else{
		fmt.Println("You are not a male of 20 years");
		}
// if age >= 18 {
// 	fmt.Println("You are an adult....");
// }else if(age <= 10){
// 	fmt.Println("You are a child....");
// } else {
// 	fmt.Println("You are a teenager....");
// }
}

//>, >=, <, <=, ==, !=, !=0, &&, ||, !, switch, for
// and => &&
// or => ||
// not => !