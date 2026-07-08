package main

import "fmt"


func add(x, y int) int {
	return x + y;
}

func subtract(x, y int) int {
	return x - y;
}

func multiply(x, y int) int {
	return x * y;
}

func divide(x, y int) int {
	return x + y;
}

func sum(numbers ...int) int {
	res := 0;
	for _, n := range numbers {
		res += n;
	}

	return res;
}

func main(){

	fmt.Println(add(4,5));
	fmt.Println(subtract(7,5));
	fmt.Println(multiply(5,5));
	fmt.Println(divide(20,4));
	fmt.Println(sum(5,6,7,8));
	fmt.Println(sum(1,2,3,4,5,6,7,8));
}
