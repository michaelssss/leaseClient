package main

import (
	"fmt"
	"os"
	"net"
	"leaseClient/alidns"
	"net/url"
	"time"
	"net/http"
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
	now:=time.Now().UTC().Format("2006-01-02T15:04:05Z")
	base := alidns.SignatureBase{"XML", "2015-01-09", "", "HMAC-SHA1", now, "1.0", "f59ed6a9-83fc-473b-9cc6-99c95df3856e"}
	signarture := alidns.GetAllDomains{&base, "DescribeDomainRecords", "liangyumingblog.com"}
	fmt.Println(url.QueryEscape(alidns.Sign(signarture.ToStringSignMap())))
	fmt.Println(url.QueryEscape(now))
	resp,err:=http.Get("http://alidns.aliyuncs.com/?Format=XML&Action=DescribeDomainRecords&AccessKeyId=LTAIvRMcO78UieFf&SignatureMethod=HMAC-SHA1&DomainName=liangyumingblog.com&SignatureNonce=f59ed6a9-83fc-473b-9cc6-99c95df3856e&Version=2015-01-09&SignatureVersion=1.0&Signature=ov4nFkxS02JkfgU9w%2FkJ95nHFuA%3D&Timestamp=2017-12-19T08%3A35%3A54Z")
	if(err!=nil){
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Status)
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
