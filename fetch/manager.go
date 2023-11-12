package fetch
import(
    "sni-fetch/challenge"
	"fmt"
	"strings"
	"sync"
)

var mu sync.Mutex
var vaildSNIs []string
var sniNum = 1

func HandleRecords(rs []Record) {
	// completedCh := make(chan struct{}, requiredSNINum)
	var wg sync.WaitGroup
	if fetch.Con > len(domainList) {
		fetch.Con = len(domainList)
	}

	index := 0
	for i:= 0; i < len(domainList) / fetch.Con; i++ {
		// Create goruntimes with the number of fetch.Con
		for j := 0; j < fetch.Con; j++ {
			wg.Add(1)
			go processChallenge(domainList[index], &wg)
			index++
		}
		wg.Wait()
		if(fetch.Num == 0) {
			// Require to fetch check all the rcds
			if(sniNum == len(domainList)) {
				output()
				return
			}
		} else {
			if(len(vaildSNIs) >= fetch.Num) {
				output()
				return
			}
		}
	}

	output()
}

func processChallenge(domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	if(challenge.Check(domain, &sniNum)) {
		vaildSNIs = append(vaildSNIs, domain)
	}
	mu.Lock()
	sniNum++
	mu.Unlock()
}

func output() {
	fmt.Printf("\n\033[32m[Finished] Found %v SNIs available: \n %v\033[0m]", len(vaildSNIs), strings.Join(vaildSNIs, "\n"))
}