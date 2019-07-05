package leaseClient

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func MakeDiscover() net.IP {
	ipdiscoverlist := []string{"http://ifconfig.me/ip", "https://ifconfig.co/ip"}
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
	if isIPv6(theip) {
		start := strings.LastIndex(theip, "[")
		end := strings.LastIndex(theip, "]")
		return net.ParseIP(theip[start+1 : end])
	} else {
		return net.ParseIP(strings.Split(theip, ":")[0])
	}
}
func isIPv6(ip string) bool {
	return strings.Contains(ip, "[")
}
