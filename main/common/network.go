package common

import (
	"XrayHelper/main/builds"
	e "XrayHelper/main/errors"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	tagNetwork = "network"
	timeout    = 3000
	dns        = "223.5.5.5:53"
)

// getHttpClient get a http client with custom dns
func getHttpClient(dns string, timeout time.Duration) *http.Client {
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{
			Resolver: &net.Resolver{
				PreferGo: true,
				Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
					d := net.Dialer{Timeout: timeout}
					return d.DialContext(ctx, "udp", dns)
				},
			},
		}
		return dialer.DialContext(ctx, network, addr)
	}
	return &http.Client{}
}

// Ping simple ping use target host&port(result max: 3000)
func Ping(protocol string, host string, port string) (result int) {
	start := time.Now()
	addr := net.JoinHostPort(host, port)
	result = -1
	switch strings.ToLower(protocol) {
	case "tcp", "http", "h2", "httpupgrade", "ws", "grpc":
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err != nil {
			return
		}
		_ = conn.Close()
	case "udp", "kcp", "mkcp":
		conn, err := net.DialTimeout("udp", addr, 2*time.Second)
		if err != nil {
			return
		}
		_ = conn.Close()
	case "quic":
		conn, err := net.DialTimeout("udp", addr, 2*time.Second)
		if err != nil {
			return
		}
		_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
		if _, err := conn.Write([]byte("\r12345678Q999\x00")); err != nil {
			_ = conn.Close()
			return
		}
		if _, err := conn.Read(make([]byte, 1024)); err != nil {
			_ = conn.Close()
			return
		}
	case "dns":
		conn, err := net.DialTimeout("udp", addr, 2*time.Second)
		if err != nil {
			return
		}
		_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
		if _, err := conn.Write([]byte("\x00\x00\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00")); err != nil {
			_ = conn.Close()
			return
		}
		if _, err := conn.Read(make([]byte, 1024)); err != nil {
			_ = conn.Close()
			return
		}
	default:
		return
	}
	result = int(time.Since(start).Milliseconds())
	return
}

// CheckLocalPort check whether the local port is listening
func CheckLocalPort(pid string, port string, ipv6 bool) bool {
	knetPath := "/proc/" + pid + "/net/tcp"
	if ipv6 {
		knetPath = "/proc/" + pid + "/net/tcp6"
	}
	i, _ := strconv.Atoi(port)
	port = fmt.Sprintf(":%X ", i)
	if knet, err := os.ReadFile(knetPath); err == nil {
		return strings.Contains(string(knet), port)
	}
	return false
}

func IsIPv6(cidr string) bool {
	ip, _, _ := net.ParseCIDR(cidr)
	if ip != nil && ip.To4() == nil {
		return true
	}
	return false
}

func CheckLocalDevice(dev string) bool {
	devices, err := net.Interfaces()
	if err != nil {
		return false
	}
	for _, device := range devices {
		if dev == device.Name {
			return true
		}
	}
	return false
}

// DownloadFile download file from url, and save to filepath
func DownloadFile(filepath string, url string) error {
	// get file from url
	client := getHttpClient(dns, timeout*time.Millisecond)
	request, _ := http.NewRequest("GET", url, nil)
	if len(builds.Config.XrayHelper.UserAgent) > 0 {
		request.Header.Set("User-Agent", builds.Config.XrayHelper.UserAgent)
	}
	response, err := client.Do(request)
	if err != nil {
		return e.New("cannot get file "+url+", ", err).WithPrefix(tagNetwork)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return e.New("bad http status "+response.Status+", ", err).WithPrefix(tagNetwork)
	}
	// open saveFile
	saveFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0755)
	if err != nil {
		return e.New("cannot open file "+filepath+", ", err).WithPrefix(tagNetwork)
	}
	defer func(saveFile *os.File) {
		_ = saveFile.Close()
	}(saveFile)
	_, err = io.Copy(saveFile, response.Body)
	if err != nil {
		return e.New("save file "+filepath+" failed, ", err).WithPrefix(tagNetwork)
	}
	return nil
}

// GetRawData get raw data from a url
func GetRawData(url string) ([]byte, error) {
	client := getHttpClient(dns, timeout*time.Millisecond)
	request, _ := http.NewRequest("GET", url, nil)
	if len(builds.Config.XrayHelper.UserAgent) > 0 {
		request.Header.Set("User-Agent", builds.Config.XrayHelper.UserAgent)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, e.New("cannot get url "+url+", ", err).WithPrefix(tagNetwork)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, e.New("bad http status "+response.Status+", ", err).WithPrefix(tagNetwork)
	}
	raw, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, e.New("read data failed, ", err).WithPrefix(tagNetwork)
	}
	return raw, nil
}
