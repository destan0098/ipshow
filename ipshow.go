package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	//domain := u.Hostname()
	//
	//ips, err := net.LookupIP(domain)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	ipshowj := flag.Bool("ip", false, "Show Just IP")
	ipdomainshowj := flag.Bool("ipd", false, "Show Just IP AND Domain")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()
	if *help {
		fmt.Println("-h : To Show Help")
		fmt.Println("-ip : To Show Just IP")
		fmt.Println("-ipd : To Show Domain with IP")
	}
	scanner := bufio.NewScanner(os.Stdin)
	if *ipshowj == true {

		for scanner.Scan() {
			line := scanner.Text()

			ips, err := net.LookupIP(line)
			if err != nil {
				fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
				continue
			}

			fmt.Println(ips[0])
		}
		// Read input data line by line

	} else if *ipdomainshowj == true {
		for scanner.Scan() {
			line := scanner.Text()
			ips, err := net.LookupIP(line)
			if err != nil {
				fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
				continue
			}
			ipurl := fmt.Sprintf("%s", ips[0])
			fmt.Println(line + " : " + ipurl)
		}

	} else {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Received:", line)
			ips, err := net.LookupIP(line)
			if err != nil {
				fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
				continue
			}
			ipurl := fmt.Sprintf("%s", ips[0])
			fmt.Println(line + " : " + ipurl)
		}
	}
}
