package main

import "fmt"

func main(){
	var arrayData = []float32{1,2,3,4,5.2}
	for index,value := range arrayData{
		fmt.Println("Data Array ke",index," : ",value)
	}

	fmt.Println("---------------")

	dataString := map[string] string{"name":"Prima Yudantra","email":"prima.yudantra@gmail.com","city":"Jakarta","country":"Indonesia"}

	for key,value:= range dataString{
		fmt.Println("Your",key,"is",value);
	}
}
