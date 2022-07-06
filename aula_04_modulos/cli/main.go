package main

import (
	"cli-exemplo/resolver"
	"flag"
	"fmt"
)

func main() {

	fmt.Println("Minha CLI")

	var host string

	flag.StringVar(&host, "host", "", "Informa o host para resolver o dns")
	flag.Parse()

	fmt.Println("Resolvendo dns do host: " + host)

	resolver.ResolverDNS(host)

}
