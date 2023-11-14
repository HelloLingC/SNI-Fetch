package fetch
import(
    "sni-fetch/challenge"
	"fmt"
	"strings"
	"sync"
)

var mu sync.Mutex
// The number of sni met the conditions
var vaildSNIs []string
// The number of already checked sni
var sniNum = 1

func HandleRecords(rs []Record) {
	ch := make(chan struct{})
	if fetch.Con > len(domainList) {
		fetch.Con = len(domainList)
	}
	dIndex := 0
	for i:=0; i < fetch.Con; i++ {
		go processChallenge(domainList[dIndex], ch)
		dIndex++
	}

	for {
		<- ch
		if fetch.Num == 0 {
			if sniNum == len(domainList) {
				break
			}
		} else {
			if fetch.Num == len(vaildSNIs) {
				break
			}
		}
		go processChallenge(domainList[dIndex], ch)
		dIndex++
	}
	output()
}

func processChallenge(domain string, ch chan struct{}) {
	if(challenge.Check(domain, &sniNum)) {
		vaildSNIs = append(vaildSNIs, domain)
	}
	mu.Lock()
	sniNum++
	mu.Unlock()
	ch <- struct{}{}
}

func output() {
	fmt.Printf("\n\033[32m[Finished] Found %v SNIs available: \n %v\033[0m", len(vaildSNIs), strings.Join(vaildSNIs, "\n"))
}