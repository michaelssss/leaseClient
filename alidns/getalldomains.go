package alidns

import (
	"net/url"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetAllDomains struct {
	Base       *SignatureBase
	Action     string
	DomainName string
}
type operation interface {
	Fire() string
}

func (getalldomain *GetAllDomains) ToStringSignMap() map[string]string {
	sMap := getalldomain.Base.ToStringSignMap()
	sMap["Action"] = getalldomain.Action
	sMap["DomainName"] = getalldomain.DomainName
	return sMap
}
func (getalldomains *GetAllDomains) Fire() string {
	map1 := getalldomains.ToStringSignMap()
	strs := make([]string, 0)
	map1["Signature"] = Sign(map1, getalldomains.Base.AccessKey)
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
