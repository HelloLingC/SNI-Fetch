package challenge

import (
	"crypto/tls"
	"github.com/gookit/color"
	"net"
	"net/http"
	"fmt"
	"time"
)

func Check(host string) bool {
	fmt.Printf("Start checking %v ...\n", host)
	result := checkTLSv3(host) && checkHTTP2(host)
	if(result) {
		color.Green.Printf("%v is available\n", host)
	}
	return result 
}

func checkTLSv3(host string)  bool {
	timeout := 5 * time.Second
	addr := net.JoinHostPort(host, "443")
	// Todo: handle www redirect
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: timeout}, "tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS11,
		MaxVersion:         tls.VersionTLS13,
	})
	if err != nil {
		return fail("Failed")
	}
	defer conn.Close()

	// Check the negotiated TLS version
	tlsVer := conn.ConnectionState().Version
	if (tlsVer != tls.VersionTLS13) {
		return fail("TLSv1.3 Failed")
	}
	return true
}

func checkHTTP2(host string) bool {
	url := "https://" + host
	client := &http.Client{
		Transport: &http.Transport{},
	}

	resp, err := client.Get(url)
	if err != nil {
		return fail("Failed")
	}
	defer resp.Body.Close()

	if !(resp.ProtoMajor == 2) {
		return fail("HTTP/2 Failed")
	}
	return true
}

func fail(info string) bool {
	color.Red.Printf(info + "\n")
	return false
}
