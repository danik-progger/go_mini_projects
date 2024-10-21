package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter domain:\n")

	for scanner.Scan() {
		domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord := checkDomain(scanner.Text())
		fmt.Printf("domain: %v\n", domain)
		fmt.Printf("hasMX: %v\n", hasMX)
		fmt.Printf("hasSPF: %v\n", hasSPF)
		fmt.Printf("spfRecord: %v\n", spfRecord)
		fmt.Printf("hasDMARC: %v\n", hasDMARC)
		fmt.Printf("dmarcRecord: %v\n", dmarcRecord)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func checkDomain(domain string) (string, bool, bool, string, bool, string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARK1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	return domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
}
