package sqlike

import (
	"github.com/RevenueMonster/sqlike/reflext"
	"github.com/RevenueMonster/sqlike/sql/util"
	"github.com/RevenueMonster/sqlike/sqlike/logs"
)

func getLogger(logger logs.Logger, debug bool) logs.Logger {
	if debug {
		return logger
	}
	return nil
}

// we should skip column generated by virtual & stored columns on insertion and migration
func skipColumns(sfs []reflext.StructFielder, omits util.StringSlice) (fields []reflext.StructFielder) {
	fields = make([]reflext.StructFielder, 0, len(sfs))
	length := len(omits)
	for _, sf := range sfs {
		// omit all the struct field with `generated_column` tag, it shouldn't include when inserting to the db
		if _, ok := sf.Tag().LookUp("generated_column"); ok {
			continue
		}
		// omit all the field provided by user
		if length > 0 && omits.IndexOf(sf.Name()) > -1 {
			continue
		}
		fields = append(fields, sf)
	}
	return
}
