package leaseClient

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// 自定义Dialer
var myDial = &net.Dialer{
	Timeout:   10 * time.Second,
	KeepAlive: 30 * time.Second,
	DualStack: false,
}

// 自定义DialContext
var myDialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
	//network = "tcp4" //仅使用ipv4
	network = "tcp6" //仅使用ipv6
	return myDial.DialContext(ctx, network, addr)
}

func MakeDiscover() net.IP {
	ipdiscoverlist := []string{"http://ip.zxinc.org/getip"}
	var client = &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           myDialContext,
			MaxIdleConns:          20,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	var rsp *http.Response
	var err error
	for index := 0; index < len(ipdiscoverlist); index++ {
		rsp, err = client.Get(ipdiscoverlist[index])
		if nil == err {
			break
		}
	}
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	err = nil
	ipbyte, err := ioutil.ReadAll(rsp.Body)
	if nil != err {
		fmt.Println(err.Error())
	}
	theip := string(ipbyte)
	fmt.Println(theip)
	return net.ParseIP(theip)
}
