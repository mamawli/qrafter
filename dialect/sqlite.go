package dialect

import "fmt"

// SQLite renders qrafter queries using SQLite placeholder and LIMIT/OFFSET
// syntax.
type SQLite struct {
	BaseDialect
}

// Literal renders SQLite-friendly inline SQL literals.
func (SQLite) Literal(value any) string {
	switch v := value.(type) {
	case bool:
		if v {
			return "1"
		}
		return "0"
	default:
		return BaseDialect{}.Literal(v)
	}
}

// LimitOffset renders SQLite LIMIT/OFFSET clauses.
func (SQLite) LimitOffset(limit, offset int) string {
	switch {
	case limit > 0 && offset > 0:
		return fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	case limit > 0:
		return fmt.Sprintf("LIMIT %d", limit)
	case offset > 0:
		return fmt.Sprintf("LIMIT -1 OFFSET %d", offset)
	default:
		return ""
	}
}
