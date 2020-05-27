package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// To build a server, we need know the functionality of it and how does it works
// There are three things a server neeed to deal with:
// 1. Listen to its Port (In this case I used port 27910 which is usually used for gaming)
// 2. Accept the request from the client and built a corresponding connection
// 3. Process the connection using goroutine and give feedback

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:27910") // I command the server to consistently listen to this port
	if err != nil {
		fmt.Println("The server failed to Listen, err:", err)
		return
	}

	for { // BY using this for loop, the server would co
		conn, err := listen.Accept() // Accept() would help the listener "listen" create a connection (type Conn) with its client
		if err != nil {
			fmt.Println("Connection built failed, err:", err)
			continue
		}
		go process(conn) // Use a goroutine to deal with the connection, between we want the server have the ability to deal with multiple client request at the sametime
	}
}

func process(conn net.Conn) {
	defer conn.Close() // In the end of processing, we need to make sure the connection is closed
	for {
		var buffer [1024]byte            // Create a byte slice
		reader := bufio.NewReader(conn)  // created a reader which is type *reader, which fetch the info from conn
		n, err := reader.Read(buffer[:]) // Reader request a byte slice as input and would return the length of the string
		if err != nil {                  // Error handling, if there is error in this part, server must fail to read input from the client
			fmt.Println("failed to read from client, err:", err)
			break
		}
		recvStr := string(buffer[:n])           //Store the inputv recieved in String form
		if strings.ToUpper(recvStr) == "EXIT" { // If the upper case form of string recieved is "EXIT", give a report
			fmt.Println("The client is over by command: ", recvStr)
			break
		}
		fmt.Println("Recieved messages from client: ", recvStr)                   // Print out what recieved from the server
		conn.Write([]byte("Response from the server: What to " + recvStr + " ?")) // Send the response of server back to client
	}
}
