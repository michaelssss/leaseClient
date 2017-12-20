package alidns

import (
	"sort"
	"net/url"
	"strings"
	"crypto/sha1"
	"crypto/hmac"
	"encoding/base64"
	"time"
)

type SignatureBase struct {
	Format           string
	AccessKey        string
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
func Sign(sMap map[string]string, accessKey string) string {
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
	StringToSign := "GET&" + url.QueryEscape("/") + "&" + url.QueryEscape(strings.Join(strss, "&"))
	key := []byte( accessKey + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(StringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
func New(AccessKey string, AccessId string) SignatureBase {
	now := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	base := SignatureBase{"JSON", AccessKey, "2015-01-09", AccessId, "HMAC-SHA1", now, "1.0", now}
	return base
}
