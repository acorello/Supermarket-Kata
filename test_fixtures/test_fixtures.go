package test_fixtures

import (
	"bufio"
	_ "embed"
	"log"
	"strconv"
	"strings"

	"dev.acorello.it/go/supermarket-kata/item"
	"dev.acorello.it/go/supermarket-kata/money"
)

//go:embed catalog.csv
var csv string

// this path is
var catalog = mustFromCSV(csv)

func Catalog() item.Catalog {
	return catalog
}

func mustFromCSV(csvString string) item.Catalog {
	scanner := bufio.NewScanner(strings.NewReader(csvString))
	checkAndSkipHeader(scanner, csvString)
	items := []item.Item{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		id := item.Id(strings.TrimSpace(fields[0]))
		price, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			log.Fatalf("failed to parse price: %v", err)
		}
		items = append(items, item.Item{
			Id:    id,
			Price: money.Cents(uint(price)),
			Unit:  strings.TrimSpace(fields[2]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error occurred during scanning: %v", err)
	}
	return item.NewCatalog(items...)
}

func checkAndSkipHeader(scanner *bufio.Scanner, path string) {
	const HEADER = "ITEM_ID,PRICE_CENTS,UNIT"
	read := scanner.Scan()
	if !read {
		log.Fatalf("empty file %q", path)
	}
	if scanner.Text() != HEADER {
		log.Fatalf("expected header %q but got %q", HEADER, scanner.Text())
	}
}
