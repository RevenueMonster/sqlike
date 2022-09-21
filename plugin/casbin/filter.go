package casbin

import (
	"github.com/RevenueMonster/sqlike/sql/expr"
	"github.com/RevenueMonster/sqlike/sqlike/primitive"
)

// Filter :
func Filter(fields ...interface{}) primitive.Group {
	return expr.And(fields...)
}
