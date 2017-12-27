package discovery

import "github.com/grandcat/zeroconf"

type ZeroConfServer struct{
	server* zeroconf.Server
}

func (this* ZeroConfServer)RegisterService(port int) {
	var err error
	this.server, err = zeroconf.Register("GoSpot", "_spotify-connect._tcp", "local.", port, []string{"VERSION=1.0", "CPath=/"}, nil)
	if err != nil {
		panic(err)
	}
}

func (this* ZeroConfServer)RemoveService() {
	this.server.Shutdown()
}