package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	showIntro()
	showMenu()
	command := readCommand()

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Showing Logs...")
	// } else if command == 0 {
	// 	fmt.Println("Exiting...")
	// } else {
	// 	fmt.Println("Wrong command!")
	// }

	switch command {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Wrong command!")
		os.Exit(-1)
	}

}

func showIntro() {
	name := "Filipe" // = var name string = "Filipe"
	age := 30
	version := 1.1
	fmt.Println("Hi,", name, ", age:", age)
	fmt.Println("This app is in version,", version)

	fmt.Println("The type of var name is:", reflect.TypeOf(name))
	fmt.Println("The type of var age is:", reflect.TypeOf(age))
	fmt.Println("The type of var version is:", reflect.TypeOf(version))
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {

	var command int
	// fmt.Scanf("%d", &command)
	fmt.Scan(&command)
	fmt.Println("Command address: ", &command)
	fmt.Println("Command: ", command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://www.alura.com.br"
	resp, err := http.Get(site)
	fmt.Println(resp, err)
}
