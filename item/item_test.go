package item_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/test_fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var catalog = test_fixtures.Catalog()

func TestCatalog_RandomItems(t *testing.T) {
	wanted := 1
	require.GreaterOrEqual(t, catalog.Len(), wanted)

	items := catalog.RandomItems(wanted)
	got := len(items)

	assert.Equal(t, wanted, got)
}
