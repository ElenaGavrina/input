package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your message: ")
	m, _ := reader.ReadString('\n')
	fmt.Printf("%s\n", m)
}
