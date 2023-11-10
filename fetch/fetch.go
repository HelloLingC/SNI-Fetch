package fetch

import(
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

type Record struct {
	IP string
	Domains []string
}

var requiredSNINum int

func errorHandle(r *colly.Response, err error) {
	log.Fatal("Error:", err)
}

func Start(addr string, num int) {
	requiredSNINum = num
	apiHost := "https://bgp.he.net"
	api := "https://bgp.he.net/ip/"
	col := colly.NewCollector()
	col.OnHTML("td.nowrap > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		getDNSList(apiHost + link)
	})

	col.OnError(errorHandle)

	err := col.Visit(api + addr)
	if err != nil {
		log.Fatal("Cannot get ASN info", err)
	}

}

var domainList []string
func getDNSList(url string) {
	var waitList []Record
	url += "#_dnsrecords"
	c := colly.NewCollector()
	c.OnHTML("tr", func(row *colly.HTMLElement) {
		var IP string
		var domains []string
		i := 0
		row.ForEach("td", func(_ int, e *colly.HTMLElement) {
			// Skip PTR records
			if i == 0 {
				IP = e.Text
			} else if i == 2 {
				domains = strings.Split(e.Text, ", ")
				domainList = append(domainList, domains...)
			}
			i++
		})
		record := Record{IP, domains}
		waitList = append(waitList, record)
	})
	c.OnScraped(func(r *colly.Response) {
		HandleRecords(waitList)
	})
	c.OnError(errorHandle)

	err := c.Visit(url)
	if err != nil {
		log.Fatal("Cannot get DNS list", err)
	}
	
}