package main

import (
	"bufio"
	"fmt"
	"os"
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
		case "exit":
			return
		case "echo":
			fmt.Println(strings.TrimPrefix(strings.TrimSpace(input), "echo "))
		case "":
			continue
		default:
			fmt.Printf("%s: command not found\n", strings.TrimSuffix(cmd, "\n"))
		}
	}
}
