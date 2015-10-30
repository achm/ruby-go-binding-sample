package main

import (
	"./ip"
	"C"
	"encoding/gob"
	"net"
	"os"
)

func main() {
}

//export search
func search(rawIP *C.char) bool {
	ipAddr := net.ParseIP(C.GoString(rawIP))
	ipRanges := decodeIPRanges()

	for _, ipRange := range ipRanges {
		if ipRange.Contains(ipAddr) {
			return true
		}
	}

	return false
}

func decodeIPRanges() []ip.IPRange {
	var ipRanges []ip.IPRange

	ipRangesFile, _ := os.Open("./ips-go")
	decoder := gob.NewDecoder(ipRangesFile)
	decoder.Decode(&ipRanges)

	return ipRanges
}
