package main

import (
	"gotify/discovery"
	"time"
)

func main() {
	//log.SetLevel(log.DebugLevel)
	//channel := new(channel.Channel)
	//key, _ := dh.MakeKey(rand.Reader, dh.SpotifyConnect)
	//channel.Config = &config.Config{PrivateKey: key}
	//channel.Connect()
	//channel.Disconnect()
	discovery.StartServer(42424)
	dis := new(discovery.ZeroConfServer)
	print("Started Zero")
	dis.RegisterService(42424)
	time.Sleep(30 * time.Minute)
	dis.RemoveService()
}
