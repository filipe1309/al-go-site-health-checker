package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const MONITORING = 3
const DELAY = 5

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()

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

	sites := readFileSites()

	fmt.Println(sites)
	fmt.Println(reflect.TypeOf(sites))

	for i := 0; i < MONITORING; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(DELAY * time.Second)
	}

	fmt.Println()
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "loaded with success")
		registerLog(site, true)
	} else {
		fmt.Println("Site", site, "not loaded. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readFileSites() []string {
	var sites []string

	file, err := os.Open("hello/sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("hello/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
