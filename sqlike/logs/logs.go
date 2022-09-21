package logs

import (
	sqlstmt "github.com/RevenueMonster/sqlike/sql/stmt"
)

// Logger :
type Logger interface {
	Debug(stmt *sqlstmt.Statement)
}
