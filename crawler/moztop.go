package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strings"
)

func NewMozReader(head int) *strings.Reader {
	// TODO: Find a better solution for this.
	resp, err := http.Get("https://moz.com/top500/domains/csv")

	if err != nil {
		log.Fatalf("error performing request %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error getting moz top 500: %s\n", resp.Status)
	}
	if resp.Header.Get("Content-Type") != "text/csv" {
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
	// Skipping csv header.
	return strings.NewReader(strings.Join(urls[1:], "\n"))

}
