package leaseClient

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func MakeDiscover() net.IP {
	ipdiscoverlist := []string{"http://v6.ip.zxinc.org/getip"}
	http.DefaultClient.Timeout = time.Second
	var rsp *http.Response
	var err error
	for index := 0; index < len(ipdiscoverlist); index++ {
		rsp, err = http.Get(ipdiscoverlist[index])
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
	return net.ParseIP(theip)
}
