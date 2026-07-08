package main

import "fmt"

func main(){
	var n int;
	fmt.Print("Enter the Number: ");
	fmt.Scan(&n);

	for i := 1; i <= 10; i++ {
		fmt.Println(i, " x ", n, " = ", n*i);
	}
}