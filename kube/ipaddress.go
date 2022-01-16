package kube

type IpAddress struct {
	Ip      string `json:"ip"`
	Netmask int    `json:"netmask"`
	Count   int    `json:"count"`
}
