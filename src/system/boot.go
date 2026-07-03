package system

import (
	"folion-server/src/infrastructure/logging"
	"folion-server/src/network"
	"folion-server/src/network/requests"
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

	requestQueue := requests.NewRequestQueue(config.GetInstance().MaxWorkers)
	requestQueue.Run()

	connectionHandler := network.NewConnectionHandler(config.GetInstance().Address,
		config.GetInstance().Port, requestQueue)

	instance := NewFolionInstance(
		config.GetInstance(),
		connectionHandler,
	)

	x.Cautious(func() error {
		return instance.StartFolionServer()
	})

	slog.Debug("Started Folion!")

}
