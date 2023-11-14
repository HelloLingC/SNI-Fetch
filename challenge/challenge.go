package challenge

import (
	"crypto/tls"
	"net/http"
	"fmt"
	"sync"
)
var mu sync.Mutex

func Check(host string, sniNum *int) bool {
	infoStr := ""
	result := test(host, &infoStr)
	mu.Lock()
	if(result) {
		fmt.Printf("[%d]\033[32m %v  Available\n\033[0m", *sniNum, host)
	} else {
		fmt.Printf("[%d] %s %s\n", *sniNum, host, infoStr)
	}
	mu.Unlock()
	return result 
}

func test(host string, info *string) bool {
	url := fmt.Sprintf("https://%v", host)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}
	// tp := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 	MinVersion:         tls.VersionTLS11,
	// 	MaxVersion:         tls.VersionTLS13,
	// 	},
	// }
	
	c := &http.Client{
		Transport: &http.Transport{},
	}
	res, _ := c.Do(req)
	// When http status code is not 200 OK
	if res == nil {
		*info = fmtRedColor(" Unreachable")
		return false
	} else {
		// *info = fmtGreenColor(" Accessible")
	}
	defer res.Body.Close()
	if res.TLS == nil {
		*info += fmtRedColor(" TLS CONN ")
		return fail("Failed to create a TLS conn")
	}
	if len(res.TLS.VerifiedChains) == 0 {
		*info += " Unknown TLS cert"
	}

	checkTLS := res.TLS.Version == tls.VersionTLS13
	checkH2 :=  res.ProtoMajor == 2 

	if checkTLS {
		*info += fmtGreenColor("  TLSv1.3  ")
	} else {
		*info += fmtRedColor("  TLSv1.3  ")
	}

	if checkH2 {
		*info += fmtGreenColor(" HTTP/2  ")
	} else {
		*info += fmtRedColor(" HTTP/2  ")
	}
	return checkTLS && checkH2
}

func fmtRedColor(s string) string {
	return fmt.Sprintf("\033[31m%v\033[0m", s)
}

func fmtGreenColor(s string) string {
	return fmt.Sprintf("\033[32m%v\033[0m", s)
}

func fail(i string) bool {
	return false
}
