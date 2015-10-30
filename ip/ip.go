package ip

import (
	"fmt"
	"math"
	"net"
	"strconv"
)

type IPRange struct {
	StartIP   net.IP
	CIDRBlock *net.IPNet
}

func NewIPRange(cols []string) *IPRange {
	ipAddr, cidrBlock, _ := parseCIDR(cols[3], cols[4])

	return &IPRange{
		StartIP:   ipAddr,
		CIDRBlock: cidrBlock,
	}
}

func parseCIDR(rawIP string, value string) (net.IP, *net.IPNet, error) {
	return net.ParseCIDR(fmt.Sprintf("%s/%s", rawIP, calcCIDRBits(value)))
}

func calcCIDRBits(rawValue string) string {
	value, _ := strconv.ParseFloat(rawValue, 64)
	return fmt.Sprint(32 - math.Log2(value))
}

func (this IPRange) Contains(ipAddr net.IP) bool {
	return this.CIDRBlock.Contains(ipAddr)
}
