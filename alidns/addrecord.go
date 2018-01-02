package alidns

import (
	"net"
	"net/url"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
)

type addRecord struct {
	Base       *signatureBase
	Action     string
	DomainName string
	RR         string
	Type       string
	Value      string
}

func (addRecord *addRecord) ToStringSignMap() map[string]string {
	sMap := addRecord.Base.ToStringSignMap()
	sMap["Action"] = addRecord.Action
	sMap["DomainName"] = addRecord.DomainName
	sMap["RR"] = addRecord.RR
	sMap["Type"] = addRecord.Type
	sMap["Value"] = addRecord.Value
	return sMap
}
func (addRecord *addRecord) Fire() string {
	map1 := addRecord.ToStringSignMap()
	strs := make([]string, 0)
	map1["Signature"] = Sign(map1, addRecord.Base.AccessKey)
	strss := make([]string, 0)
	for index, _ := range map1 {
		strs = append(strs, index)
	}
	for index, value := range strs {
		strss = append(strss, url.QueryEscape(strs[index])+"="+url.QueryEscape(map1[value]))
	}
	strings.Join(strss, "&")
	resp, err := http.Get("http://alidns.aliyuncs.com/?" + strings.Join(strss, "&"))
	if (err != nil) {
		fmt.Println(err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	return string(b)
}
func AddRecord(domainName string, rr string, base *signatureBase, ip net.IP) addRecord {
	return addRecord{base, "AddDomainRecord", domainName, rr, "A", ip.String()}
}
