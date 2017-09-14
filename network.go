package metadata

import (
	"errors"
	"net"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/milosgajdos83/tenus"
)

func CreateDummyInterface() {
	link, err := tenus.NewLink("docker-metadata")
	if err != nil {
		logrus.Debugf(err.Error())
		return
	}
	ip, network, _ := net.ParseCIDR("169.254.169.254/32")
	link.SetLinkIp(ip, network)
	link.SetLinkUp()
}

func GetIpForInterface(iface string) (string, error) {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		if i.Name == iface || ((iface == "" && !strings.Contains(i.Name, "lo")) && (iface == "" && !strings.Contains(i.Name, "docker"))) {
			addrs, _ := i.Addrs()
			addrWithNet := addrs[0].String()
			addr := strings.Split(addrWithNet, "/")[0] + "\n"
			return addr, nil
		}
	}
	return "", errors.New("Interface address not found")
}
