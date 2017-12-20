package alidns

import (
	"net/url"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
)

type deleteRecord struct {
	Base     *signatureBase
	Action   string
	RecordId string
}

func (deleteRecord *deleteRecord) ToStringSignMap() map[string]string {
	sMap := deleteRecord.Base.ToStringSignMap()
	sMap["Action"] = deleteRecord.Action
	sMap["RecordId"] = deleteRecord.RecordId
	return sMap
}
func (deleteRecord *deleteRecord) Fire() string {
	map1 := deleteRecord.ToStringSignMap()
	strs := make([]string, 0)
	map1["Signature"] = Sign(map1, deleteRecord.Base.AccessKey)
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
func DeleteRecord(base *signatureBase, recordId string) deleteRecord {
	return deleteRecord{base, "DeleteDomainRecord", recordId}
}
