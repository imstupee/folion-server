package system

import (
	"folion-server/src/network"
	"folion-server/src/system/config"
)

type FolionInstance struct {
	SafeBoot           bool
	Config             *config.Config
	ConnectionHahndler *network.ConnectionHandler
}

func NewFolionInstance(_config *config.Config, _connectionhandler *network.ConnectionHandler) *FolionInstance {
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

func (folionInstance *FolionInstance) StopFolionServer() error {
	return nil
}
