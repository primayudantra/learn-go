package main

import "fmt"

type User struct{
	name string
	age int
	email string
}

func main(){
	var Prima User
	var Isna 	User

	Prima.name = "Prima Yudantra"
	Prima.age = 23
	Prima.email = "prima.yudantra@gmail.com"

	Isna.name = "Hey You"
	Isna.age = 25
	Isna.email = "hey.you@gmail.com"

	printUser(Prima)
	printUser(Isna)
}

func printUser(user User){
	fmt.Printf("User name: %s\n", user.name);
	fmt.Printf("User age: %d\n", user.age);
	fmt.Printf("User email: %s\n", user.email);
}