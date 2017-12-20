package leaseClient

import (
	"net"
	"fmt"
	"io"
	"strings"
)

func MakeDiscover(ip string, key string) net.IP {
	conn, err := net.Dial("tcp", ip)
	if nil != err {
		fmt.Println(err.Error())
	}
	conn.Write([]byte(key))
	buf := make([]byte, 64)
	publicIpByte := make([]byte, 0)
	for {
		i, err := conn.Read(buf)
		if io.EOF == err {
			break
		}
		publicIpByte = append(publicIpByte, buf[0:i]...)
	}
	return net.ParseIP(strings.Split(string(publicIpByte), ":")[0])
}
