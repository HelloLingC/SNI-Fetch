package main

import (
	"flag"
	"fmt"
	"sni-fetch/fetch"
)

func main() {
	
	targetIP := flag.String("t", "", "The target IP")
	sniNum := flag.Int("n", 1, "The required number of sni")
	conNum := flag.Int("c", 5, "")
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
	`
	// 8 lines
	fmt.Printf(title + "\n")
}
