package network

import (
	"fmt"
	"log/slog"
	"net"
)

type ConnectionHandler struct {
	address string
	port    string
}

func NewConnectionHandler(_address string, _port string) *ConnectionHandler {
	return &ConnectionHandler{
		address: _address,
		port:    _port,
	}
}

func (connectionHandler *ConnectionHandler) StartConnectionHandler() error {
	addr, err := net.ResolveTCPAddr("tcp", connectionHandler.port)
	if err != nil {
		return err
	}
	socket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	defer socket.Close()
	slog.Debug(fmt.Sprintf("Started Connection handler on port %s", connectionHandler.port))

	for {
		connection, err := socket.Accept()
		if err != nil {
			return err
		}

		go connectionHandler.HandlerConnection(&connection)
	}

	return nil
}

func (connectionHandler *ConnectionHandler) HandlerConnection(connection *net.Conn) error {
	return nil
}
