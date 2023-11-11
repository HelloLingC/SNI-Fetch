package fetch

import(
	"github.com/gocolly/colly/v2"
	"log"
	"fmt"
	"net"
	"strings"
)

type Fetch struct {
	Addr string
	Num int
	Con int
}

type Record struct {
	IP string
	Domains []string
}

var fetch *Fetch
var domainList []string

func errorHandle(r *colly.Response, err error) {
	log.Fatal("Error:", err)
}

func Start(f Fetch) {
	fmt.Printf("Start fetching...\n")
	fetch = &f
	ip := net.ParseIP(fetch.Addr)
	if ip == nil {
		log.Fatal("Inputed a invaild IP address")
	}
	apiHost := "https://bgp.he.net"
	api := "https://bgp.he.net/ip/"
	col := colly.NewCollector()
	col.OnHTML("td.nowrap > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Get %v\n", link)
		getDNSList(apiHost + link)
	})

	col.OnError(errorHandle)

	err := col.Visit(api + fetch.Addr)
	if err != nil {
		log.Fatal("Cannot get ASN info", err)
	}

}

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