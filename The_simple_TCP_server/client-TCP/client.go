package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:27910") // If success, this will return the pointer of the connection made
	if err != nil {
		fmt.Println("This is a problem in Dailing, err :", err)
		return
	}
	defer conn.Close() // Use defer to fianlly close the connection

	inputReader := bufio.NewReader(os.Stdin) // Return the detailed address for a new reader with default size

	for {
		input, _ := inputReader.ReadString('\n')  // Read the input of the user
		inputInfo := strings.Trim(input, "\r\n")  // Clean the input up
		if strings.ToUpper(inputInfo) == "EXIT" { // if any string, with a upper case form of "EXIT", is inputed, then terminate the client
			_, err = conn.Write([]byte("exit"))
			if err != nil {
				return
			}
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // Send what is inputed to the server
		if err != nil {
			return
		}

		buffer := [1024]byte{}
		n, err := conn.Read(buffer[:]) // read from the conn to get the response from the server and store the response into the buffer
		if err != nil {
			fmt.Println("failed to recieve from the server, err:", err)
			return
		}
		fmt.Println(string(buffer[:n])) // Print out the string version of the buffer
	}
}
