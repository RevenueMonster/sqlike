package debug

import (
	"github.com/RevenueMonster/sqlike/sql/dialect"
	"github.com/RevenueMonster/sqlike/sql/dialect/mysql"
	sqlstmt "github.com/RevenueMonster/sqlike/sql/stmt"
)

// ToSQL :
func ToSQL(src interface{}) error {
	ms := dialect.GetDialectByDriver("mysql").(*mysql.MySQL)
	sqlstmt.NewStatement(ms)
	return nil
}
