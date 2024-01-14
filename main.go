package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	proxy2 "github.com/onuragtas/reverse-proxy/proxy"
)

type Settings struct {
	Listen   string `json:"listen"`
	Mappings []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	} `json:"mappings"`
}

var settings Settings

func main() {
	path, _ := os.Getwd()
	bytes, _ := os.ReadFile(path + "/settings.json")
	json.Unmarshal(bytes, &settings)

	listener, err := net.Listen("tcp", settings.Listen)
	if err != nil {
		panic("connection error:" + err.Error())
	}
	log.Println("Proxy listening on", settings.Listen, "...")
	for {
		conn, err := listener.Accept()
		proxy := proxy2.Proxy{Src: conn, OnResponse: onResponse, OnRequest: onRequest, RequestHost: setDestination}
		if err != nil {
			fmt.Println("Accept Error:", err)
			continue
		}
		go proxy.Handle()
	}
}

func onRequest(srcLocal, srcRemote, dstLocal, dstRemote string, request []byte) {
	log.Println(srcLocal, "->", srcRemote, "->", dstLocal, "->", dstRemote, string(request))
}

func onResponse(dstRemote, dstLocal, srcRemote, srcLocal string, response []byte) {
	log.Println(dstRemote, "->", dstLocal, "->", srcRemote, "->", srcLocal)
}

func setDestination(req []byte, host string, src net.Conn) string {
	for _, mapping := range settings.Mappings {
		if strings.Contains(string(req), mapping.Path) {
			return mapping.URL
		}
	}
	return ""
}
