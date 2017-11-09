package leaseClient

import (
	"net"
	"encoding/json"
)

type Client struct {
	ClientName string
	ClientAddr string
}
type ClientOperations interface {
	MakeDiscover(ip string)
	Renew(ip string)
}

func (client *Client) MakeDiscover(ip string) {
	conn, err := net.Dial("tcp", ip)
	defer conn.Close()
	result, _ := json.Marshal(client)
	if nil != err {
		panic(err)
	}
	content := Community{len(result), result}.ToByte()
	conn.Write(content)
}
func (client *Client) Renew(ip string) {
	client.MakeDiscover(ip)
}