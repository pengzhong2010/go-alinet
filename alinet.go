package alinet

import (
	"fmt"
	"net"
	"os"
)

//自动获取ip
func GetIntranetIp() string {
	ipLocal := ""

	//以太网网卡名称为eth0
	inter, err := net.InterfaceByName("eth0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	addrs, err := inter.Addrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//ip地址一个ip4一个ip6
	//for _, addr := range addrs {
	//        fmt.Println(addr.String())
	//}
	//addrs, err := net.InterfaceAddrs()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("ip:", ipnet.IP.String())
				if ipnet.IP.String() != "172.17.0.1" {
					if !IsPublicIP(ipnet.IP) {
						ipLocal = ipnet.IP.String()
					}
				}
			}
		}
	}
	return ipLocal
}

//判断是否为外网ip
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
