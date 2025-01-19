package mac

import (
	"crypto/rand"
	"net"
)

func GetNode() ([]byte, error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if len(iface.HardwareAddr) == 6 {
			return iface.HardwareAddr, nil
		}

	}
	// if no mac was found
	randomNode := make([]byte, 6)
	_, err = rand.Read(randomNode)

	if err != nil {
		return nil, err
	}

	return randomNode, nil
}
