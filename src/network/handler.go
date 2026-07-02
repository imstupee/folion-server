package network

type ConnectionHandler struct {
	address string
	port    int
}

func NewConnectionHandler(_address string, _port int) *ConnectionHandler {
	return &ConnectionHandler{
		address: _address,
		port:    _port,
	}
}

func (connectionHandler *ConnectionHandler) StartConnectionHandler() error {

	return nil
}
