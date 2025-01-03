package shareurls

import (
	e "XrayHelper/main/errors"
	"XrayHelper/main/serial"
	"XrayHelper/main/shareurls/addon"
	"strings"
)

const (
	tagShareurl     = "shareurl"
	socksPrefix     = "socks://"
	ssPrefix        = "ss://"
	vmessPrefix     = "vmess://"
	vlessPrefix     = "vless://"
	trojanPrefix    = "trojan://"
	hysteriaPrefix  = "hysteria://"
	hysteria2Prefix = "hysteria2://"
	hy2Prefix       = "hy2://"
	wireguardPrefix = "wireguard://"
)

// ShareUrl implement this interface, that node can be converted to core OutboundObject
type ShareUrl interface {
	GetNodeInfoStr() string
	GetNodeInfo() *addon.NodeInfo
	ToOutboundWithTag(coreType string, tag string) (*serial.OrderedMap, error)
}

// Parse return a ShareUrl
func Parse(link string) (ShareUrl, error) {
	if strings.HasPrefix(link, socksPrefix) {
		return parseSocks(link)
	}
	if strings.HasPrefix(link, ssPrefix) {
		return parseShadowsocks(link)
	}
	if strings.HasPrefix(link, vmessPrefix) {
		return parseVmess(link)
	}
	if strings.HasPrefix(link, vlessPrefix) {
		return parseVLESS(link)
	}
	if strings.HasPrefix(link, trojanPrefix) {
		return parseTrojan(link)
	}
	if strings.HasPrefix(link, hysteriaPrefix) {
		return parseHysteria(link)
	}
	if strings.HasPrefix(link, hysteria2Prefix) || strings.HasPrefix(link, hy2Prefix) {
		return parseHysteria2(link)
	}
	if strings.HasPrefix(link, wireguardPrefix) {
		return parseWireguard(link)
	}
	return nil, e.New("not a supported share link").WithPrefix(tagShareurl)
}
