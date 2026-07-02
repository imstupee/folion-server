package system

import (
	"folion-server/src/network"
)

type FolionInstance struct {
	SafeBoot           bool
	Config             *Config
	ConnectionHahndler *network.ConnectionHandler
}

func NewFolionInstance(_config *Config, _connectionhandler *network.ConnectionHandler) *FolionInstance {
	return &FolionInstance{
		Config:             _config,
		ConnectionHahndler: _connectionhandler,
	}

}

func (folionInstance *FolionInstance) StartFolionServer() error {
	if err := folionInstance.ConnectionHahndler.StartConnectionHandler(); err != nil {
		return err
	}
	return nil
}
