package main

import (
	"os"
	"leaseClient/alidns"
	"fmt"
	"encoding/json"
	"leaseClient/client"
	"time"
)

type Record struct {
	RR         string
	Status     string
	Value      string
	Weight     int
	RecordId   string
	Type       string
	DomainName string
	Locked     bool
	Line       string
	TTL        int
}
type DomainRecords struct {
	Record []Record
}
type Resp struct {
	PageNumber    int
	TotalCount    int
	PageSize      int
	RequestId     string
	DomainRecords DomainRecords
}

var nowIp string

func main() {
	ss := os.Args
	accessKey := ss[1]
	accessId := ss[2]
	communityKey := ss[3]
	ipport := ss[4]
	ticket := time.NewTicker(time.Second * 20)
	func() {
		for _ = range ticket.C {
			ip := leaseClient.MakeDiscover(ipport, communityKey)
			if nil != ip && ip.String() != nowIp {

				base := alidns.SignatureBase(accessKey, accessId)
				getAllDomains := alidns.GetAllDomains("liangyumingblog.com", &base)

				json1 := getAllDomains.Fire()
				resp := Resp{}
				err := json.Unmarshal([]byte(json1), &resp)
				if nil != err {
					fmt.Println(err.Error())
				}
				for _, value := range resp.DomainRecords.Record {
					if value.RR == "home" {
						deleteDomain := alidns.DeleteRecord(&base, value.RecordId)
						deleteDomain.Fire()
					}
				}
				addRecord := alidns.AddRecord("liangyumingblog.com", &base, ip)
				addRecord.Fire()
				nowIp = ip.String()
			}
		}
	}()

}
