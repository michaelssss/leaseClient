package leaseClient

import (
	 "strings"
	"fmt"
)

type Config interface {
	Parse(bytes []byte) config
	Get(key string) string
	Set(key string, value string) bool
	ToString()
}
type config struct {
	configMap map[string]string
}

func (cfg *config) Parse(bytes []byte) config {
	stringss := strings.Split(string(bytes), "\r\n")
	for index := range stringss {
		sss := strings.Split(stringss[index], "=")
		key := sss[0]
		value := sss[1]
		cfg.configMap[key] = value
	}
	return *cfg
}
func (cfg *config) Get(key string) string {
	return cfg.configMap[key]
}
func (cfg *config) Set(key string, value string) bool {
	return false
}
func (cfg *config) ToString() {
	for key, value := range cfg.configMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}
}
func NewConfig() Config {
	return &config{configMap: make(map[string]string)}
}
func GetConfig(maps map[string]string) Config {
	return &config{configMap: maps}
}
