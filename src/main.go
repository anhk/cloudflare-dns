package main

import (
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"net"
)

// getSysIp 获取系统IP地址
func getSysIp(infName string) (string, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, v := range ifs {
		if v.Name != infName {
			continue
		}
		addrs, err := v.Addrs()
		if err != nil {
			return "", err
		}

		for _, vv := range addrs {
			ip, _, err := net.ParseCIDR(vv.String())
			if err != nil {
				return "", err
			}
			return ip.String(), nil
		}

	}
	return "", errors.New("no record.")
}

func main() {

	cfg, err := load("/usr/local/etc/cloudflare-dns.json")
	if err != nil {
		panic(err)
	}

	sysIp, err := getSysIp(cfg.InfName)
	if err != nil {
		panic(err)
	}

	api, err := cloudflare.New(cfg.CfKey, "anhk@ir0.cn")
	if err != nil {
		panic(err)
	}

	list, err := api.DNSRecords(cfg.ZoneId, cloudflare.DNSRecord{Name: cfg.KeyWord})
	if err != nil {
		panic(err)
	}

	if len(list) == 1 && list[0].Content == sysIp {
		return
	}

	if _, err := api.CreateDNSRecord(cfg.ZoneId, cloudflare.DNSRecord{
		Type:    "A",
		Name:    cfg.KeyWord,
		Content: sysIp,
	}); err != nil {
		panic(err)
	}
}
