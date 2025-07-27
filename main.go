package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided. Type `help` for available commands.")
		return
	}

	command := os.Args[1]

	switch command {
	case "help":
		fmt.Println("Available commands:")
		fmt.Println("  help    - Show available commands")
		fmt.Println("  whoami  - Show the current system user")
	case "whoami":
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("Error fetching user:", err)
			return
		}
		fmt.Println("Current user:", currentUser.Username)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Type `help` for available commands.")
	}
}
