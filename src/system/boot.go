package system

import (
	"folion-server/src/infrastructure/logging"
	"folion-server/src/network"
	"log/slog"
	"os"
)

func BootFolion() {
	err := logging.InitLogger()
	if err != nil {
		os.Stdout.WriteString("Failed to start logger. Exiting.")
		os.Exit(1)
	}
	slog.Debug("Started logging")

	config, err := LoadConfig()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	connectionHandler := network.NewConnectionHandler(config.Address, config.Port)

	instance := NewFolionInstance(
		config,
		connectionHandler,
	)

	if err := instance.StartFolionServer(); err != nil {
		slog.Error(err.Error())
	}

}
