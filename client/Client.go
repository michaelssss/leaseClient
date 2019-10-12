package leaseClient

import (
	"net"
	"os"
)

func MakeDiscover() net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv6 := addr.To16(); ipv6 != nil && !ipv6.IsLinkLocalUnicast() {
			return ipv6
		}
	}
	//return ""
	return nil
}
