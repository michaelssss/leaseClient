package leaseClient

import (
	"fmt"
	"net"
	"os"
)

func MakeDiscover() net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv6 := addr.To4(); ipv6 == nil && addr.IsGlobalUnicast() && !addr.IsLinkLocalUnicast() {
			fmt.Println(addr.String())
			return addr
		}
	}
	//return ""
	return nil
}
