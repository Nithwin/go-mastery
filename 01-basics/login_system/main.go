package main

import "fmt"

func main(){
	var username, password string;

	fmt.Print("Enter username: ");
	fmt.Scan(&username);
	fmt.Print("Enter password: ");
	fmt.Scan(&password);

	if username == "nithwin" && password == "password" {
		fmt.Println("Access granted!");
	} else {
		fmt.Println("Access denied!");
	}
}
