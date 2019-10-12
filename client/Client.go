package leaseClient

import (
	"net"
)

func MakeDiscover() net.IP {
	conn, err := net.Dial("tcp6", "ipv6.michaelssss.cc:80")
	if nil != conn {
		defer conn.Close()
	}
	if nil == err {
		return conn.LocalAddr().(*net.IPNet).IP
	}
	return nil
}
