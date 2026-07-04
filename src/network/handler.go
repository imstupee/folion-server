package network

import (
	"encoding/json"
	"fmt"
	"folion-server/src/network/requests"
	"folion-server/src/system/x"
	"log/slog"
	"net"
)

type ConnectionHandler struct {
	address      string
	port         string
	requestQueue *requests.RequestQueue
}

func NewConnectionHandler(_address string, _port string, _requestQueue *requests.RequestQueue) *ConnectionHandler {
	return &ConnectionHandler{
		address:      _address,
		port:         _port,
		requestQueue: _requestQueue,
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
	slog.Debug(fmt.Sprintf("Started Connection handler on %s:%s", socket.Addr(), connectionHandler.port))

	for {
		connection, err := socket.Accept()
		if err != nil {
			return err
		}
		slog.Debug(fmt.Sprintf("Got connections from %s", connection.RemoteAddr()))
		go connectionHandler.HandleConnection(connection)
	}
}

func (connectionHandler *ConnectionHandler) registerSession(connection net.Conn) {

}

func (connectionHandler *ConnectionHandler) HandleConnection(connection net.Conn) error {
	defer connection.Close()
	decoder := json.NewDecoder(connection)

	var request requests.Request
	if err := decoder.Decode(&request); err != nil {
		return x.ErrRequestDecodeFailed.Message(
			fmt.Sprintf("Failed to decode request from %s", connection.RemoteAddr())).Wrap(err)
	}

	slog.Debug(request.Type)

	go func(_request requests.Request) {
		requests.RequestQueueChan <- _request
	}(request)

	return nil
}
