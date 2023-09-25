package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/destan0098/ipshow/IPINfo"
	"net"
	"os"
)

func domainwithpip(ipshowj, ipdomainshowj, ipdomaininfshowj bool) {
	scanner := bufio.NewScanner(os.Stdin)
	if ipshowj == true {
		// Show just the IP address
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

	} else if ipdomainshowj == true {
		// Show both the domain and the IP address
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

	} else if ipdomaininfshowj == true {
		for scanner.Scan() {
			line := scanner.Text()
			ips, err := net.LookupIP(line)
			if err != nil {
				fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
				continue
			}
			ipurl := fmt.Sprintf("%s", ips[0])
			fmt.Println(line + " : " + ipurl)
			IPINfo.IpInfo(ipurl)
		}
	} else {
		// Default behavior: Show received input and IP address
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
func domainwithname(ipshowj, ipdomainshowj, ipdomaininfshowj bool, domains string) {

	if ipshowj == true {
		// Show just the IP address

		ips, err := net.LookupIP(domains)
		if err != nil {
			fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
			recover()
		}

		fmt.Println(ips[0])

		// Read input data line by line

	} else if ipdomainshowj == true {
		// Show both the domain and the IP address

		ips, err := net.LookupIP(domains)
		if err != nil {
			fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
			recover()
		}
		ipurl := fmt.Sprintf("%s", ips[0])
		fmt.Println(domains + " : " + ipurl)

	} else if ipdomaininfshowj == true {

		ips, err := net.LookupIP(domains)
		if err != nil {
			fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
			recover()
		}
		ipurl := fmt.Sprintf("%s", ips[0])
		fmt.Println(domains + " : " + ipurl)
		IPINfo.IpInfo(ipurl)

	} else {
		// Default behavior: Show received input and IP address

		fmt.Println("Received:", domains)
		ips, err := net.LookupIP(domains)
		if err != nil {
			fmt.Println("Error: Not RESOLVED , Maybe it doesn't have an IP. ")
			recover()
		}
		ipurl := fmt.Sprintf("%s", ips[0])
		fmt.Println(domains + " : " + ipurl)

	}
}
func main() {

	//Receive input from the user.
	// Define command-line flags
	ipshowj := flag.Bool("ip", false, "Show Just IP")
	ipdomainshowj := flag.Bool("ipd", false, "Show Just IP AND Domain")
	ipdomaininfshowj := flag.Bool("ipin", false, "Show Just IP AND Domain")
	domains := flag.String("d", "google.com", "domain to check")
	pipeline := flag.Bool("pipe", false, "domain to check from pipe line")
	help := flag.Bool("h", false, "Show help")
	// Parse command-line arguments
	flag.Parse()
	// Display help message if requested
	if *help {
		helpme()

	}
	//give from Pipe Line
	// Create a scanner to read input data
	if *pipeline == true {
		domainwithpip(*ipshowj, *ipdomainshowj, *ipdomaininfshowj)
	} else if *domains != "" {
		domainwithname(*ipshowj, *ipdomainshowj, *ipdomaininfshowj, *domains)
	} else {
		fmt.Println("Please Add domain ")
		helpme()
	}

}
func helpme() {
	fmt.Println("-h : To Show Help")
	fmt.Println("-ip : To Show Just IP")
	fmt.Println("-ipd : To Show Domain with IP")
	fmt.Println("-ipin : To Show Domain with IP info")
	fmt.Println("-d : Domain to Check")
	fmt.Println("-pipe : Domain to Check From Pipe line")
}
