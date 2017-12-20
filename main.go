package main

import (
	"os"
	"leaseClient/client"
	"leaseClient/alidns"
	"fmt"
)

func main() {
	//ticket := time.NewTicker(time.Second * 20)
	//func() {
	//	for _ = range ticket.C {
	//		client := leaseClient.Client{ClientName: "company", ClientAddr: getAddr().(*net.IPNet).IP.String()}
	//		fmt.Println(client)
	//		client.MakeDiscover("127.0.0.1:8888")
	//	}
	//}()
	ss := os.Args
	accessKey := ss[1]
	accessId := ss[2]
	communityKey := ss[3]
	ipport := ss[4]
	leaseClient.MakeDiscover(ipport, communityKey)
	base := alidns.SignatureBase(accessKey, accessId)
	GetAllDomains := alidns.GetAllDomains("liangyumingblog.com", &base)
	//fmt.Println(url.QueryEscape(alidns.Sign(GetAllDomains.ToStringSignMap())))
	json1 := GetAllDomains.Fire()
	fmt.Println(json1)
}
