package fetch
import(
    "sni-fetch/challenge"
	"fmt"
	_ "log"
	"strings"
	"sync"
)

var vaildSNIs []string
var sniNum = 0

func HandleRecords(rs []Record) {
	// completedCh := make(chan struct{}, requiredSNINum)
	var wg sync.WaitGroup

	con := 5
	index := 0
	for i:= 0; i < len(domainList) / con; i++ {
		for j := 0; j < con; j++ {
			wg.Add(1)
			go processChallenge(domainList[index], &wg)
			index++
		}
		wg.Wait()
		if(sniNum >= requiredSNINum) {
			output()
			return
		}
	}

	output()
}

func processChallenge(domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	if(challenge.Check(domain)) {
		vaildSNIs = append(vaildSNIs, domain)
		sniNum++
	}
}

func output() {
	fmt.Printf("[Finished] Found %v SNIs available: \n %v", len(vaildSNIs), strings.Join(vaildSNIs, "\n"))
}