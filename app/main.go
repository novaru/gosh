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
		input = strings.TrimSpace(input)

		args := strings.Split(input, " ")
		cmd := strings.ToLower(args[0])
		switch cmd {
		case EXIT:
			return
		case ECHO:
			fmt.Println(strings.TrimPrefix(input, "echo "))
		case TYPE:
			command := strings.TrimPrefix(input, "type ")
			switch command {
			case ECHO, EXIT, TYPE:
				fmt.Println(command + " is a shell builtin")
			default:
				path, err := exec.LookPath(command)
				if err != nil {
					fmt.Println(command + ": not found")
				} else {
					fmt.Println(command, "is", path)
				}
			}
		case "":
			continue
		default:
			if _, err := exec.LookPath(args[0]); err != nil {
				fmt.Println(strings.TrimSuffix(cmd, "\n") + ": command not found")
			} else {
				arguments := args[1:]
				command := exec.Command(args[0], arguments...)
				var out strings.Builder
				command.Stdout = &out
				if err := command.Run(); err != nil {
					fmt.Errorf(err.Error())
				} else {
					fmt.Print(out.String())
				}
			}
		}
	}
}
