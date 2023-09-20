package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {

	//Receive input from the user.
	// Define command-line flags
	ipshowj := flag.Bool("ip", false, "Show Just IP")
	ipdomainshowj := flag.Bool("ipd", false, "Show Just IP AND Domain")
	help := flag.Bool("h", false, "Show help")
	// Parse command-line arguments
	flag.Parse()
	// Display help message if requested
	if *help {
		fmt.Println("-h : To Show Help")
		fmt.Println("-ip : To Show Just IP")
		fmt.Println("-ipd : To Show Domain with IP")
	}
	//give from Pipe Line
	// Create a scanner to read input data
	scanner := bufio.NewScanner(os.Stdin)
	if *ipshowj == true {
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

	} else if *ipdomainshowj == true {
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
