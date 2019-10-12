package leaseClient

import (
	"fmt"
	"net"
)

func MakeDiscover() net.IP {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			if ipv6, ok := addr.(*net.IPNet); ok && nil == ipv6.IP.To4() && ipv6.IP.IsGlobalUnicast() && !ipv6.IP.IsLinkLocalUnicast() {
				fmt.Println(ipv6.IP.String())
				return ipv6.IP
			}
		}
	}

	//return ""
	return nil
}
