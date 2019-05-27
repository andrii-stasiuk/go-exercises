package main

import (
	"bufio"
	"flag"
	"fmt"
	printer "go-exercises/client-server/client/inc"
	requests "go-exercises/client-server/client/inc"
	"os"
	"strings"
)

func main() {
	var addrPtr = flag.String("addr", "127.0.0.1:8000", "API server IPv4 address and port")
	flag.Parse()
	fmt.Println("Client is using IP-address (and port): " + *addrPtr)
	for {
		fmt.Print("Choose an operation adduser|changeuser|deluser|getuser or exit: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		command := strings.Fields(input)
		if len(command) > 0 {
			switch command[0] {
			case "adduser":
				var name, surname, email string
				fmt.Print("Enter Name: ")
				fmt.Scanln(&name)
				fmt.Print("Enter Surname: ")
				fmt.Scanln(&surname)
				fmt.Print("Enter Email: ")
				fmt.Scanln(&email)
				printer.MapPrinter(requests.PostRequest(*addrPtr, "", name, surname, email))
			case "changeuser":
				var id, name, surname, email string
				fmt.Print("Enter ID: ")
				fmt.Scanln(&id)
				fmt.Print("Enter Name: ")
				fmt.Scanln(&name)
				fmt.Print("Enter Surname: ")
				fmt.Scanln(&surname)
				fmt.Print("Enter Email: ")
				fmt.Scanln(&email)
				printer.MapPrinter(requests.PostRequest(*addrPtr, id, name, surname, email))
			case "deluser", "deleteuser":
				if len(command) > 1 {
					printer.MapPrinter(requests.DeleteRequest(*addrPtr, command[1]))
				} else {
					fmt.Println("Error: user ID is not specified")
				}
			case "getuser":
				if len(command) > 1 {
					printer.MapPrinter(requests.GetRequest(*addrPtr, command[1]))
				} else {
					fmt.Println("Error: user ID is not specified")
				}
			case "e", "exit":
				fmt.Println("Client has been successfully shut down")
				return
			default:
				fmt.Println("Unknown task: " + input)
			}
		}
	}
}
