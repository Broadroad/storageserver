package server

import (
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Server control all this app
type Server struct {
	address     string        // addrss to open connection
	listener    net.Listener  // listener listen xx port
	joinsniffer chan net.Conn //joinsniffer get all connection
}

// NewServer return a new server
func NewServer() *Server {
	return &Server{}
}

// Run start the server
func (this *Server) Run() {
	log.Info("Server is starting...")
	// init server
	interruptHandler()

	this.listen()
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
	b := make([]byte, 100)
	for {
		_, err := conn.Read(b)
		if err == io.EOF {
			break
		}
	}
	conn.Write([]byte("HTTP/1.1 200 OK\n"))
}

// interruptHandler handle signal when server close
func interruptHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("captured %v, stopping profiler and exiting..", sig)
		os.Exit(1)
	}()
}
