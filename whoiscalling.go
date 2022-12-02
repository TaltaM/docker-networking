package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func RunServer() {
	fmt.Printf("Listening at port %s\n", os.Getenv("LISTEN_PORT"))
	listenPort, err := strconv.ParseInt(os.Getenv("LISTEN_PORT"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		incomingConnection, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		netData, err := bufio.NewReader(incomingConnection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Printf("-> from: %s while listening at %s\n",
			incomingConnection.RemoteAddr().String(),
			incomingConnection.LocalAddr().String())
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		incomingConnection.Write([]byte(myTime))
		incomingConnection.Close()
	}
}

func RunClient() {
	time.Sleep(5 * time.Second)

	fmt.Printf("Calling for %s\n", os.Getenv("PEER_ADDRESS"))
	peerAddress := os.Getenv("PEER_ADDRESS")
	connection, err := net.Dial("tcp", peerAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(connection, text+"\n")

	message, _ := bufio.NewReader(connection).ReadString('\n')
	fmt.Print("->: " + message)
	if strings.TrimSpace(string(text)) == "STOP" {
		fmt.Println("TCP client exiting...")
		return
	}
	connection.Close()
}

func main() {
	fmt.Println("Who is calling?")
	//	fmt.Println(os.Getenv("ROLE"))

	//	role := os.Getenv("ROLE")

	go RunClient()
	RunServer()
	// time.Sleep(5 * time.Second)
	// fmt.Println("Exiting")
}
