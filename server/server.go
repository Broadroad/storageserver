package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"bufio"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Server control all this app
type Server struct {
	address     string        // addrss to open connection
	listener    net.Listener  // listener listen xx port
	joinsniffer chan net.Conn //joinsniffer get all connection
}

// NewServer return a new server
func NewServer(address string) *Server {
	return &Server{address: address}
}

// Run start the server
func (s *Server) Run() {
	log.Info("Server is starting...")
	// init server
	interruptHandler()
	s.listen()
}

// listen accept connection
func (s *Server) listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

// handleConnection really handle connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Info("Connection come from " + conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}
		fmt.Println(message)
		log.WithFields(log.Fields{
			"message":    message,
		}).Info("get message")
		conn.Write([]byte("HTTP/1.1 200 OK\n"))	
	}
}

// interruptHandler handle signal when server close
func interruptHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.WithFields(log.Fields{
			"sig":    sig,
		}).Info("stopping profiler and exiting...")	
		os.Exit(1)
	}()
}
