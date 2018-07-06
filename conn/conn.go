package conn

import (
	"errors"
	"net"
	"time"
)

var (
	// ErrConnUnavialible is returned when Connection Unavaliable
	ErrConnUnavialible = errors.New("Connection Unavaliable")
	// ErrConnBroken is returned when Conenction broken
	ErrConnBroken = errors.New("Connection Broken")
	// ErrConnSigInter is returned when Connection sig interrupt
	ErrConnSigInter = errors.New("Connection sig interrupt")
	// ErrConnTimeout is returned when Connection timeout
	ErrConnTimeout = errors.New("Connection timeout")
)

// Conn hold on connection from client
type Conn struct {
	Conn         net.Conn
	ReadTimeout  time.Duration // sets the read timeout in the connection.
	WriteTimeout time.Duration // sets the write timeout in the connection.
}

// NewConn return new Conn with timeout
func NewConn(c net.Conn, readTimeout, writeTimeout time.Duration) *Conn {
	return &Conn{
		Conn:         c,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}
