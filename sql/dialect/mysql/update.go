package mysql

import (
	sqlstmt "github.com/RevenueMonster/sqlike/sql/stmt"
	"github.com/RevenueMonster/sqlike/sqlike/actions"
)

// Update :
func (ms *MySQL) Update(stmt sqlstmt.Stmt, f *actions.UpdateActions) (err error) {
	err = buildStatement(stmt, ms.parser, f)
	if err != nil {
		return
	}
	return
}
