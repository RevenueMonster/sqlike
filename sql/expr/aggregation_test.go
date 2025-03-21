package expr

import (
	"testing"

	"github.com/RevenueMonster/sqlike/sqlike/primitive"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	require.Equal(t, primitive.Aggregate{
		Field: wrapColumn("a"),
		By:    primitive.Sum,
	}, Sum("a"))
	return
}
