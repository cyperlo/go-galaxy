package main

import "net"

func isIPInCIDR(ipstr, cidrStr string) bool {
	ip := net.ParseIP(ipstr)
	_, ipNet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		return false
	}
	return ipNet.Contains(ip)
}

func main() {
}
