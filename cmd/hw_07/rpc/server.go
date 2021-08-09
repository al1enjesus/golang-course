package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"unicode/utf8"
)

type User struct {
	id int
	name string
}

func (u User) String() string {
	return fmt.Sprintf("[User with id: %d and name: %s] ", u.id, u.name)
}

type FakeDb struct {
	db []User
}

func (fdb FakeDb) String() string {
	var output string
	for _, user := range fdb.db {
		output = fmt.Sprintf("%s %s", user.String(), output)
	}
	return output
}

func (fdb FakeDb) IsExist(name string) bool {
	for _, user := range fdb.db {
		if user.name == name {
			return true
		}
	}
	return false
}

func (fdb *FakeDb) AddUser(name string) {
	currentId := len(fdb.db)
	fdb.db = append(fdb.db, User{currentId, name})
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()

	var fdb FakeDb

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		length := utf8.RuneCountInString(string(message))
		message = message[:(length-2)]

		fmt.Println("Message from client: ", message)

		command := strings.Split(message, " ")[0]

		switch command {
		case "register":
			if len(strings.Split(message, " ")) == 1 {
				conn.Write([]byte("Incorrect input. Use: [register <name>]\n"))
				break
			}
			userName := strings.Split(message, " ")[1]
			if fdb.IsExist(userName) {
				conn.Write([]byte("This user already exists.\n"))
			} else {
				fdb.AddUser(userName)
				conn.Write([]byte(fmt.Sprintf("User successfully created with id: %d\n", len(fdb.db) - 1)))
			}
		case "list":
			conn.Write([]byte(fdb.String() + "\n"))
		default:
			conn.Write([]byte("Unknown command.\n"))
		}
	}
}