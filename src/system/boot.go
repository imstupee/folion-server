package system

import (
	"folion-server/src/infrastructure/logging"
	"folion-server/src/network"
	"folion-server/src/system/config"
	"folion-server/src/system/x"
	"log/slog"
)

func BootFolion() {
	x.Cautious(func() error {
		return logging.InitLogger()
	})

	slog.Debug("Started logging")

	x.Cautious(func() error {
		return config.LoadConfig()
	})

	connectionHandler := network.NewConnectionHandler(config.GetInstance().Address,
		config.GetInstance().Port)

	instance := NewFolionInstance(
		config.GetInstance(),
		connectionHandler,
	)

	x.Cautious(func() error {
		return instance.StartFolionServer()
	})

	slog.Debug("Started Folion!")

}
