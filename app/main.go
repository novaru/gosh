package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", input)
			os.Exit(1)
		}

		args := strings.Split(strings.TrimSpace(input), " ")
		cmd := strings.ToLower(args[0])
		switch cmd {
		case EXIT:
			return
		case ECHO:
			fmt.Println(strings.TrimPrefix(strings.TrimSpace(input), "echo "))
		case TYPE:
			command := strings.TrimPrefix(strings.TrimSpace(input), "type ")
			switch command {
			case ECHO, EXIT, TYPE:
				fmt.Println(command + " is a shell builtin")
			default:
				path, err := exec.LookPath(command)
				if err != nil {
					fmt.Println(command + ": not found")
				}
				fmt.Println(command, "is", path)
			}
		case "":
			continue
		default:
			fmt.Println(strings.TrimSuffix(cmd, "\n") + ": command not found")
		}
	}
}
