package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	wg.Done()

	for {
		recvMsg, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Print(">" + recvMsg)
	}
}

func sendMessage(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		msgToSend := strings.Trim(scanner.Text(), "\r\n")
		conn.Write([]byte(msgToSend + "\n"))

	}
}

func runClient() {

	fmt.Println("Joined server")

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		fmt.Println("Unable to connect to server: ", err.Error())
		os.Exit(0)
	}

	wg.Add(2)
	go sendMessage(conn, &wg)
	go socketRecv(conn, &wg)
	wg.Wait()
}

func main() {
	runClient()
}
