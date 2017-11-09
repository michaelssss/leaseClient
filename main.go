package main

import (
	"time"
	"fmt"
	"os"
	"net"
	"leaseClient/client"
)

func main() {
	ticket := time.NewTicker(time.Second * 20)
	func() {
		for _ = range ticket.C {
			client := leaseClient.Client{ClientName: "company", ClientAddr: getAddr().(*net.IPNet).IP.String()}
			fmt.Println(client)
			client.MakeDiscover("michaelssss.com:8888")
		}
	}()
}
func getAddr() net.Addr {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return address
			}

		}
	}
	return nil
}
