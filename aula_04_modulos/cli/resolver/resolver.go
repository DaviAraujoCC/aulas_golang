package resolver

import (
	"fmt"
	"net"
)

func ResolverDNS(host string) {

	ip_addrs, err := net.LookupHost(host)
	if err != nil {
		panic(err)
	}

	for _, ip := range ip_addrs {
		fmt.Println(ip)
	}
}
