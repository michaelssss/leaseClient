package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"leaseClient/alidns"
	"leaseClient/client"
	"os"
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
var config leaseClient.Config

func main() {
	ss := os.Args
	configPath := ss[1]
	config = leaseClient.NewConfig()
	config.Parse(readFile(configPath))
	accessKey := config.Get("accessKey")
	accessId := config.Get("accessId")
	ttype := config.Get("type")
	communityKey := config.Get("communityKey")
	ipport := config.Get("ipport")
	rr := config.Get("rr")
	domain := config.Get("domain")
	ticket := time.NewTicker(time.Second * 20)
	func() {
		for _ = range ticket.C {
			ip := leaseClient.MakeDiscover(ipport, communityKey)
			if nil != ip && ip.String() != nowIp {

				base := alidns.SignatureBase(accessKey, accessId)
				getAllDomains := alidns.GetAllDomains(domain, &base)

				json1 := getAllDomains.Fire()
				resp := Resp{}
				err := json.Unmarshal([]byte(json1), &resp)
				if nil != err {
					fmt.Println(err.Error())
				}
				for _, value := range resp.DomainRecords.Record {
					if value.RR == rr {
						deleteDomain := alidns.DeleteRecord(&base, value.RecordId)
						deleteDomain.Fire()
					}
				}
				addRecord := alidns.AddRecord(domain, rr, ttype, &base, ip)
				addRecord.Fire()
				nowIp = ip.String()
			}
		}
	}()
}

func readFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if nil != err {
		fmt.Println(err.Error())
		panic(err)
	}
	return file
}
