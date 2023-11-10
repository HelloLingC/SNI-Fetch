package main

import (
	"flag"
	"regexp"
	"log"
	//"fmt"
	"sni-fetch/fetch"
)

func main() {
	ipPattern := `\b(?:\d{1,3}\.){3}\d{1,3}\b`
	ipRegex := regexp.MustCompile(ipPattern)
	targetIP := flag.String("t", "", "The target IP")
	sniNum := flag.Int("n", 1, "The required number of sni")
	flag.Parse()
	isVaildIP := ipRegex.MatchString(*targetIP)
	if !isVaildIP {
		log.Fatal("Please give a correct IP address!")
	}
	
	fetch.Start(*targetIP, *sniNum)
}
