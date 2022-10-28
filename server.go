package main

import (
	"bufio"
	"log"
	"net"
	"sync"
)

var clients []*client
var wg sync.WaitGroup

const (
	PORT     string = ":8080"
	PROTOCOL string = "tcp"
)

type client struct {
	conn net.Conn
	// nickname string
	// room     string
}

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
	}
}

func receiveClientMsg(conn net.Conn) {
	for {
		// fmt.Print(len(clients))
		msg, _ := bufio.NewReader(conn).ReadString('\n')

		// checks if there are at least 2 clients for communication
		if len(clients) > 1 {
			for i := 0; i < len(clients); i++ {

				if clients[i].conn != conn {
					clients[i].conn.Write([]byte(msg))
				}
			}
		}

	}
}

func runServer() {
	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on: %s\n", PORT)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		newClient := newClient(conn)
		clients = append(clients, newClient)

		go receiveClientMsg(conn)
	}
}

func main() {
	runServer()
}
