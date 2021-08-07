package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()

	fmt.Println("Allowed commands: list, register <name>")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(fmt.Sprintf("Message from server: %s", message))
	}
}