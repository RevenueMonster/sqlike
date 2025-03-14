package sqlstmt

import (
	"fmt"
	"sync"

	"reflect"

	"github.com/RevenueMonster/sqlike/reflext"
)

// BuildStatementFunc :
type BuildStatementFunc func(stmt Stmt, it interface{}) error

// StatementBuilder :
type StatementBuilder struct {
	mutex    *sync.Mutex
	builders map[interface{}]BuildStatementFunc
}

// NewStatementBuilder :
func NewStatementBuilder() *StatementBuilder {
	return &StatementBuilder{
		mutex:    new(sync.Mutex),
		builders: make(map[interface{}]BuildStatementFunc),
	}
}

// SetBuilder :
func (sb *StatementBuilder) SetBuilder(it interface{}, p BuildStatementFunc) {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()
	sb.builders[it] = p
}

// LookupBuilder :
func (sb *StatementBuilder) LookupBuilder(t reflect.Type) (blr BuildStatementFunc, ok bool) {
	blr, ok = sb.builders[t]
	return
}

// BuildStatement :
func (sb *StatementBuilder) BuildStatement(stmt Stmt, it interface{}) error {
	v := reflext.ValueOf(it)
	if x, ok := sb.builders[v.Type()]; ok {
		return x(stmt, it)
	}
	if x, ok := sb.builders[v.Kind()]; ok {
		return x(stmt, it)
	}

	return fmt.Errorf("sqlstmt: invalid data type support %v", v.Type())
}
