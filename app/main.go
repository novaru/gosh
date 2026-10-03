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
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", cmd)
			os.Exit(1)
		}
		fmt.Printf("%s: command not found\n", strings.TrimSuffix(cmd, "\n"))
	}
}
