package client

import (
	"github.com/hekmon/transmissionrpc"
	"log"
	"omniarr/internal/config"
	"strconv"
)

var Transmission *transmissionrpc.Client

func init() {
	port, err := strconv.ParseUint(config.AppConfig.TransmissionPort, 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	client, err := transmissionrpc.New(
		config.AppConfig.TransmissionHost,
		config.AppConfig.TransmissionUser,
		config.AppConfig.TransmissionPass,
		&transmissionrpc.AdvancedConfig{
			Port: uint16(port),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	Transmission = client

	log.Println("🔌 Transmission client initialized")
}
