package session

import (
	"bytes"
	"encoding/json"
	"folion-server/src/system/x"
	"net"

	"github.com/google/uuid"
)

type Session struct {
	ID   uuid.UUID
	conn net.Conn
}

func NewSession(_conn net.Conn) *Session {
	return &Session{
		ID:   uuid.New(),
		conn: _conn,
	}
}

func (session *Session) Write(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return x.ErrSocketWriteFailed.Wrap(err)
	}
	bytesData := bytes.NewBuffer(jsonData)
	session.conn.Write(bytesData.Bytes())

	return nil
}
