package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// MozReader creates new reader which represents list of sites addresses from
// https://moz.com/top500 delimeted by new line charaster. Request timeout is 60s.
func MozReader(head int) *strings.Reader {
	// TODO: Create a better solution for providing top sites.

	client := http.Client{Timeout: time.Second * 60}
	fmt.Printf("Getting csv with top 500 from moz.com... ")
	start := time.Now()
	resp, err := client.Get("https://moz.com/top500/domains/csv")

	if err != nil {
		log.Fatalf("error performing request %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error getting moz top 500: %s\n", resp.Status)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/csv") {
		log.Fatalln("error getting csv data, wrong conten type")
	}
	defer resp.Body.Close()

	var urls []string

	reader := csv.NewReader(resp.Body)
	// haed + 1 because of csv header.
	for i := 0; i <= head; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading csv %s\n", err)
		}
		url := "http://" + strings.TrimRight(record[1], "/")

		urls = append(urls, url)

	}

	// Do not count strings join and new reader creation.
	fmt.Printf("done in %.2fs\n", time.Since(start).Seconds())
	// Skipping csv header.
	return strings.NewReader(strings.Join(urls[1:], "\n"))
}
