package main

import "fmt"

// Basic Function
func swap(p, y string) (string, string){
	return y, p
}

func main() {

	// IF Statements
	var result int;
	var number1, number2 = 4, 2
	result = number1 + number2;

	if(result != 6){
		fmt.Println("result tidak sama dengan ", result)
	}else{
		fmt.Println("result sama dengan ", result)
	}

	// Loop Statements
	var angka int = 2;
	for angka < 10{
		fmt.Printf("Data: %d\n", angka)
		angka++;
		if angka > 20{
			break;
		}
	}

	// Calling function swap()
	a, b := swap("Prima", "Yudantra")
	fmt.Println(a, b)

	// Basic Array
	var arrayData = []float32{1,2,3,4,5.2}

	// Basic For Each
	for index,value := range arrayData{
		fmt.Println("Data Array ke",index," : ",value)
	}

	// Pointer *
	var data_ptr int = 8;
	var ptr *int

	ptr = &data_ptr;

	fmt.Printf("Data pointer is : %d\n", *ptr)

}