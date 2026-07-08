package main

import (
	"fmt"
	"os"
	"bufio"
)

func add(todos []string) []string {
	fmt.Print("Enter your todo: ")
	reader := bufio.NewReader(os.Stdin);
	todo, _ := reader.ReadString('\n');
	todos = append(todos, todo);
	fmt.Println("Todo added!...");

	return todos;
};

func remove(todos []string) []string {
	var index int;
	fmt.Print("Enter your todo no: ")
	fmt.Scan(&index);
	todos = append(todos[:index-1], todos[index:]...)
	fmt.Println("Todo removed!...");
	return todos;
};

func view(todos []string) {
	for i, value := range todos {
		fmt.Println(i+1,".", value);
	}
}

func main(){
	var todos []string;

	for {
		var option int;
		fmt.Print("1.add \n2.remove \n3.view \n4.clear \n5.exit \n Your option (in number): ");
		fmt.Scan(&option);

		if option == 1 {
			todos = add(todos);
		} else if option == 2 {
			todos = remove(todos);
		} else if option == 3 {
			view(todos);
		} else if option == 4 {
			todos = nil;
			fmt.Println("Your Todo list cleared...");
		} else {
			break;
		}
		fmt.Println();
	}
}
