package dialect

import (
	"fmt"

	"github.com/SennovE/qrafter/internal/utils"
)

const mysqlOffsetOnlyLimit = "18446744073709551615"

// MySQL renders qrafter queries using MySQL identifier and LIMIT/OFFSET syntax.
//
// qrafter does not validate whether every query feature is supported by MySQL.
// In particular, PostgreSQL-style RETURNING, UPDATE ... FROM, DELETE ... USING,
// and NULLS FIRST/LAST ordering need dialect-specific handling before they can
// be considered fully portable.
type MySQL struct {
	BaseDialect
}

// QuoteIdent renders a MySQL backtick-quoted identifier.
func (MySQL) QuoteIdent(ident string) string {
	return utils.QuoteWith(ident, "`")
}

// LimitOffset renders MySQL LIMIT/OFFSET clauses.
func (MySQL) LimitOffset(limit, offset int) string {
	switch {
	case limit > 0 && offset > 0:
		return fmt.Sprintf("LIMIT %d, %d", offset, limit)
	case limit > 0:
		return fmt.Sprintf("LIMIT %d", limit)
	case offset > 0:
		return fmt.Sprintf("LIMIT %s OFFSET %d", mysqlOffsetOnlyLimit, offset)
	default:
		return ""
	}
}
