package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		length := utf8.RuneCountInString(string(message))
		message = message[:(length-2)]
		fmt.Println("Message Received:", message)

		var newMessage string
		value, err := strconv.ParseInt(message, 0, 64)
		if err != nil {
			newMessage = strings.ToUpper(message)
		} else {
			newMessage = strconv.Itoa(int(value * 2))
		}

		conn.Write([]byte(newMessage + "\n"))
	}
}