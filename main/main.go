package main

import metadata "github.com/hjdr4/docker-metadata"
import "os"

func main() {
	var iface string
	if len(os.Args) > 1 {
		iface = os.Args[1]
	} else {
		iface = ""
	}
	metadata.SetIface(iface)
	metadata.Serve()
}
