package main

import (
	"io/ioutil"
	"log"
	"net"
	"path"
	"strconv"
	"strings"
)

type ipTtl struct {
	IP  net.IP
	TTL uint32
}

type Hosts map[string]ipTtl

func ParseHost(filename string, hosts Hosts) {
	ttl, err := strconv.Atoi(strings.Replace(path.Ext(filename), ".", "", -1))
	if err != nil || ttl == 0 {
		ttl = 600
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("[WARN] open hosts fila failed", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		ip := net.ParseIP(fields[0])
		if ip == nil {
			continue
		}
		for _, domain := range fields[1:] {
			hosts[domain] = ipTtl{IP: ip, TTL: uint32(ttl)}
		}
	}
}

func ParseHostsFiles(filenames []string) (hosts Hosts) {
	hosts = make(Hosts)
	for _, fn := range filenames {
		ParseHost(fn, hosts)
	}
	return
}
