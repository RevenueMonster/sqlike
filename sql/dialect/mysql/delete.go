package mysql

import (
	sqlstmt "github.com/RevenueMonster/sqlike/sql/stmt"
	"github.com/RevenueMonster/sqlike/sqlike/actions"
)

// Delete :
func (ms *MySQL) Delete(stmt sqlstmt.Stmt, f *actions.DeleteActions) (err error) {
	err = buildStatement(stmt, ms.parser, f)
	if err != nil {
		return
	}
	return
}
