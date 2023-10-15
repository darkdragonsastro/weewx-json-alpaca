package alpaca

import (
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
)

type AlpacaDiscovery struct {
	log  *zap.Logger
	port int
	stop bool
}

func NewAlpacaDiscovery(log *zap.Logger, port int) *AlpacaDiscovery {
	return &AlpacaDiscovery{
		log:  log,
		port: port,
		stop: false,
	}
}

// Handle incoming UDP packets.
func (a *AlpacaDiscovery) handleDiscovery(conn *net.UDPConn) {
	for {
		if a.stop {
			break
		}

		buf := make([]byte, 16)

		conn.SetReadDeadline((time.Now().Add(time.Second * 5)))

		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			if err.(net.Error).Timeout() {
				continue
			}

			a.log.Error("Error reading UDP packet", zap.Error(err))
			continue
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

func (a *AlpacaDiscovery) StopDiscovery() {
	a.stop = true
}
