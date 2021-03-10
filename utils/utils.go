package utils

import (
	"net"
)

func GetInterfaceIP(name string) (net.IP) {
	ifi, err := net.InterfaceByName(name)
	if err != nil {
		return nil
	}
	addrs, err := ifi.Addrs()
	if err != nil {
		return nil
	}
	for _, addr := range addrs {
		ip, inet, err := net.ParseCIDR(addr.String())
		if err != nil {
			return nil
		}
		if len(inet.Mask) == 4 {
			return ip
		}
	}
	return nil
}
