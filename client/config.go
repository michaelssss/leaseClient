package leaseClient

import (
	"fmt"
	"strings"
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
	deliters := []string{"\r\n", "\r", "\n"}
	for deliter := range deliters {
		stringss := strings.Split(string(bytes), deliters[deliter])
		if len(stringss) > 1 {
			for index := range stringss {
				if strings.Contains(stringss[index], "=") {
					sss := strings.Split(stringss[index], "=")
					key := strings.Replace(sss[0], " ", "", -1)
					key = strings.Replace(key, "\r", "", -1)
					key = strings.Replace(key, "\n", "", -1)
					value := strings.Replace(sss[1], " ", "", -1)
					value = strings.Replace(value, "\r", "", -1)
					value = strings.Replace(value, "\n", "", -1)
					cfg.configMap[key] = value
				}
			}
		}
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
