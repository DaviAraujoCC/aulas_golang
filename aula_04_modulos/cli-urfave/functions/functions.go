package functions

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func ResolverDNS(c *cli.Context) {

	ip_addrs, err := net.LookupHost(c.String("host"))
	if err != nil {
		panic(err)
	}

	for _, ip := range ip_addrs {
		fmt.Println(ip)
	}

}
