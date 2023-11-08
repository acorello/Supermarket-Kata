package test_fixtures_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
)

func TestCatalogLoading_RowCountMatches(t *testing.T) {
	filename := "catalog.csv"
	lineCount := countLines(filename)
	lineCountMinusHeader := lineCount - 1
	assert.Equal(t, lineCountMinusHeader, test_fixtures.Catalog().Len())
}

func countLines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file: %v", err)
	}
	return lineCount
}
