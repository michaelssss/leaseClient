package alidns

import (
	"sort"
	"net/url"
	"strings"
	"fmt"
	"crypto/sha1"
	"crypto/hmac"
	"encoding/base64"
)

type SignatureBase struct {
	Format           string
	Version          string
	AccessKeyId      string
	SignatureMethod  string
	Timestamp        string
	SignatureVersion string
	SignatureNonce   string
}
type sign interface {
	ToStringSignMap() map[string]string
}

func (signture *SignatureBase) ToStringSignMap() map[string]string {
	sMap := make(map[string]string)
	sMap["Format"] = signture.Format
	sMap["Version"] = signture.Version
	sMap["AccessKeyId"] = signture.AccessKeyId
	sMap["SignatureMethod"] = signture.SignatureMethod
	sMap["Timestamp"] = signture.Timestamp
	sMap["SignatureVersion"] = signture.SignatureVersion
	sMap["SignatureNonce"] = signture.SignatureNonce
	return sMap
}
func Sign(sMap map[string]string) string {
	strs := make([]string, 0)
	for index, _ := range sMap {
		strs = append(strs, index)
	}
	strs = strs[0:]
	sort.Strings(strs)
	strss := make([]string, 0)
	for index, value := range strs {
		strss = append(strss, url.QueryEscape(strs[index])+"="+url.QueryEscape(sMap[value]))
	}
	fmt.Println(strings.Join(strss, "&"))
	StringToSign := "GET&" + url.QueryEscape("/") + "&" + url.QueryEscape(strings.Join(strss, "&"))
	fmt.Println(StringToSign)
	key := []byte("" + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(StringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
