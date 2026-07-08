package main

import "fmt"

func main(){
	// var arr = [5]int{1,2,3,4,5};
	// for _, n := range arr {
	// 	fmt.Println(n)
	// }

	var arr []int

	arr = append(arr, 5);
	arr = append(arr, 5);
	arr = append(arr, 5);
	arr = append(arr, 5);

	for _, n := range arr {
		fmt.Println(n);
	}

}
