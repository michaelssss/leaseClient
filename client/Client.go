package leaseClient

import (
	"net"
	"fmt"
	"io"
	"strings"
	"time"
)

func MakeDiscover(ip string, key string) net.IP {
	conn, err := net.Dial("tcp", ip)
	defer func() {
		if nil != conn {
			err := conn.Close()
			if nil != err {
				fmt.Println(err.Error())
			}
		}
	}()
	if nil == conn {
		fmt.Println("connect to ", ip, " failed ,time is ", time.Now().String())
	}
	if nil != err {
		fmt.Println(err.Error())
		return net.ParseIP("")
	}
	_, err = conn.Write([]byte(key))
	if nil != err {
		fmt.Println(err.Error())
	}
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
