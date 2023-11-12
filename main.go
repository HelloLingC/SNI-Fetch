package main

import (
	"flag"
	"fmt"
	"sni-fetch/fetch"
)

func main() {
	
	targetIP := flag.String("t", "", "The target IP")
	sniNum := flag.Int("n", 1, "The required number of sni")
	conNum := flag.Int("c", 10, "The number of concurrent checks in a single round")
	flag.Parse()
	printProcInfo()
	fetch.Start(fetch.Fetch{
		Addr : *targetIP,
		Num : *sniNum,
		Con : *conNum})
}


func printProcInfo() {
	title := `
 ███████╗███╗   ██╗██╗ ███████╗███████╗████████╗ ██████╗██╗  ██╗
 ██╔════╝████╗  ██║██║ ██╔════╝██╔════╝╚══██╔══╝██╔════╝██║  ██║
 ███████╗██╔██╗ ██║██║ █████╗  █████╗     ██║   ██║     ███████║
 ╚════██║██║╚██╗██║██║ ██╔══╝  ██╔══╝     ██║   ██║     ██╔══██║
 ███████║██║ ╚████║██║ ██║     ███████╗   ██║   ╚██████╗██║  ██║
 ╚══════╝╚═╝  ╚═══╝╚═╝ ╚═╝     ╚══════╝   ╚═╝    ╚═════╝╚═╝  ╚═╝                
	
  Github: https://github.com/HelloLingC/SNI-Fetch
 `
	fmt.Printf(title + "\n")
}
