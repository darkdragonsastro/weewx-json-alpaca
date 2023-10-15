package alpaca

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

type AlpacaDiscovery struct {
	log  *zap.Logger
	port int
}

func NewAlpacaDiscovery(log *zap.Logger, port int) *AlpacaDiscovery {
	return &AlpacaDiscovery{
		log:  log,
		port: port,
	}
}

// Handle incoming UDP packets.
func (a *AlpacaDiscovery) handleDiscovery(conn *net.UDPConn) {
	for {
		buf := make([]byte, 16)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		a.log.Info("Received UDP packet", zap.String("packet", string(buf[0:n])))

		if string(buf) == "alpacadiscovery1" {
			conn.WriteToUDP([]byte(fmt.Sprintf("{\"AlpacaPort\":%d}", a.port)), addr)
		}
	}
}

// Start a background listener for Alpaca Discovery UDP packet.
func (a *AlpacaDiscovery) StartDiscovery() error {
	addr, err := net.ResolveUDPAddr("udp", ":32227")
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	go a.handleDiscovery(conn)

	return nil
}
