package challenge

import (
	"crypto/tls"
	_ "golang.org/x/term"
	_ "github.com/gookit/color"
	"net"
	"net/http"
	"fmt"
	"time"
	"sync"
)
var mu sync.Mutex
func Check(host string, sniNum *int) bool {
	result := checkTLSv3(host) && checkHTTP2(host)
	mu.Lock()
	if(result) {
		fmt.Printf("[%d] %v \033[32mavailable\n\033[0m", *sniNum, host)
	} else {
		fmt.Printf("[%d] %v \033[31mfailed\n\033[0m", *sniNum, host)
	}
	mu.Unlock()
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
		return fail("Failed when checking TLS")
	}
	defer conn.Close()

	// Check the negotiated TLS version
	tlsVer := conn.ConnectionState().Version
	if (tlsVer != tls.VersionTLS13) {
		return fail("not supports TLSv1.3")
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
		return fail("Failed when checking h2")
	}
	defer resp.Body.Close()

	if !(resp.ProtoMajor == 2) {
		return fail("not supports HTTP/2")
	}
	return true
}

func fail(info string) bool {
	// color.Red.Printf(info + "\n")
	return false
}
