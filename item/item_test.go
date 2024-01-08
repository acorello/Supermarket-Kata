package item_test

import (
	"testing"

	"dev.acorello.it/go/supermarket-kata/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var inventory = item.FixedInventory()

func TestFixedCatalog_RandomItems(t *testing.T) {
	wanted := 1
	require.GreaterOrEqual(t, inventory.Len(), wanted)

	items := inventory.RandomItems(wanted)
	got := len(items)

	assert.Equal(t, wanted, got)
}
