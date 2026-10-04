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
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", input)
			os.Exit(1)
		}

		args := strings.Split(input, " ")
		cmd := strings.ToLower(args[0])
		switch cmd {
		case CD:
			var path string
			if len(args) == 1 {
				path = os.Getenv("HOME")
			} else {
				path = args[1]
			}
			if err := os.Chdir(path); err != nil {
				fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", args[1])
			}
		case EXIT:
			return
		case ECHO:
			fmt.Println(strings.TrimLeft(strings.TrimRight(fmt.Sprint(args[1:]), "]"), "["))
		case TYPE:
			command := args[1]
			switch args[1] {
			case ECHO, EXIT, PWD, TYPE:
				fmt.Println(command + " is a shell builtin")
			default:
				path, err := exec.LookPath(command)
				if err != nil {
					fmt.Println(command + ": not found")
				} else {
					fmt.Println(command, "is", path)
				}
			}
		case PWD:
			if path, err := os.Getwd(); err != nil {
				fmt.Errorf(err.Error())
			} else {
				fmt.Println(path)
			}
		case "":
			continue
		default:
			if _, err := exec.LookPath(args[0]); err != nil {
				fmt.Fprintln(os.Stderr, strings.TrimSuffix(cmd, "\n")+": command not found")
			} else {
				arguments := args[1:]
				command := exec.Command(args[0], arguments...)
				var out strings.Builder
				command.Stdout = &out
				if err := command.Run(); err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
				} else {
					fmt.Print(out.String())
				}
			}
		}
	}
}
