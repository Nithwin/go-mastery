package main

import (
	"fmt"
)


func changeAge(age int) {
	age = 25
}

func main() {
	age := 21
	ptr := &age;
	changeAge(age)
	*ptr = 23
	fmt.Println(&age);
	fmt.Println(ptr);
	fmt.Println(*ptr);
	fmt.Println(&ptr);
	fmt.Println(age);

}