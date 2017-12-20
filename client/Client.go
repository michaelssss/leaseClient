package leaseClient

import (
	"net"
	"fmt"
	"io"
)

func MakeDiscover(ip string, key string) {
	conn, err := net.Dial("tcp", ip)
	defer conn.Close()
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
	ipp := net.ParseIP(string(publicIpByte)[0:9])
	fmt.Println(ipp)
}
