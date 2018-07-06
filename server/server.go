package server

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const (
	// BUFFERSIZE define the read or write buffer size
	BUFFERSIZE = 2 * 1024 * 1024
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
	//buffer := new(bytes.Buffer)
	reader := bufio.NewReader(conn)
	data := make([]byte, 0, BUFFERSIZE)
	for {
		i := 1
		for i <= 5 {
			n, err := io.ReadFull(reader, data[:i])
			buf := bytes.NewBuffer(data)
			i = i + 1
			dataLen, _ := binary.ReadVarint(buf)
		}

		n, err := io.ReadFull(reader, data[:5])
		data = data[:n]
		if err != nil {
			if err != io.EOF {
				break
			}
		}

		dataLen := binary.BigEndian.Uint32(data)

		n, err = io.ReadFull(reader, data[:dataLen])

		if err != nil {
			conn.Close()
			return
		}
		message := "test"
		fmt.Println(message)
		log.WithFields(log.Fields{
			"message": message,
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
			"sig": sig,
		}).Info("stopping profiler and exiting...")
		os.Exit(1)
	}()
}
