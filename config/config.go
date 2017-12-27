package config

import "gotify/dh"

type Config struct {
	DeviceId string
	PrivateKey *dh.PrivateKey
}
