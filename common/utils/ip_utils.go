package utils

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"regexp"
	"strings"
)

/**
 * 获取本地的ip <br/>
 * return [] string
 * 可能有双网卡，或者多个ip 。数组最后一个是默认的ip：127.0.0.1
 */
func GetLocalHostIps() (ipHosts []string) {
	address, err := net.InterfaceAddrs()
	localhost := "127.0.0.1"
	ips := make([]string, 0)
	if err != nil {
		log.Printf("Error: get lochost ip is wrong ..")
		ips = append(ips, localhost)
		return ips
	}

	for _, addr := range address {
		IpDr := addr.String()
		match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, IpDr)
		if !match {
			continue
		}
		ip := strings.Split(IpDr, "/")[0]
		if localhost != ip {
			ips = append(ips, ip)
		}
	}

	ips = append(ips, localhost)
	return ips
}

/**
  解析url
*/
func ParseUrl(baseUrl string) (string, error) {
	matched, _ := regexp.Match("(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]", []byte(baseUrl))
	if !matched {
		return baseUrl, errors.New(fmt.Sprintf("URL %s 协议头必须为 (http 或者 https)", baseUrl))
	}

	u, err := url.Parse(baseUrl)
	if err != nil {
		return baseUrl, err
	}

	baseUrl = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	return baseUrl, nil
}
