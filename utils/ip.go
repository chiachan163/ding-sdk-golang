package utils

import (
	"net"
	"strconv"
	"strings"
)

func inet_aton(ipnr net.IP) (arr []int) {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	arr = append(arr, b0, b1, b2, b3)
	return
}

func GetIntranetIp() (ip []int, err error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return inet_aton(ipnet.IP), nil
			}
		}
	}
	return
}
