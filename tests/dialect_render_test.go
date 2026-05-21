package tests

import (
	"testing"

	q "github.com/SennovE/qrafter"
	"github.com/SennovE/qrafter/dialect"
	"github.com/stretchr/testify/assert"
)

func TestMySQLDialectRender(t *testing.T) {
	users := q.MustNewTable[User]()

	sql, args := q.Select(users.UserName).
		Where(users.Age.Ge("18")).
		Limit(10).
		Offset(20).
		Render(dialect.MySQL{})

	assert.Equal(t, "SELECT `table`.`user_name`\nFROM `table`\nWHERE `table`.`userAge` >= ?\nLIMIT 20, 10", sql)
	assert.Equal(t, []any{"18"}, args)
	assert.Equal(t, "`weird``name`", dialect.MySQL{}.QuoteIdent("weird`name"))
}

func TestMySQLDialectRender_OffsetWithoutLimit(t *testing.T) {
	users := q.MustNewTable[User]()

	sql, args := q.Select(users.UserName).
		Offset(20).
		Render(dialect.MySQL{})

	assert.Equal(
		t,
		"SELECT `table`.`user_name`\nFROM `table`\nLIMIT 18446744073709551615 OFFSET 20",
		sql,
	)
	assert.Empty(t, args)
}

func TestSQLiteDialectRender(t *testing.T) {
	users := q.MustNewTable[User]()

	sql, args := q.Select(users.UserName).
		Offset(20).
		Render(dialect.SQLite{})

	assert.Equal(t, `SELECT "table"."user_name"
FROM "table"
LIMIT -1 OFFSET 20`, sql)
	assert.Empty(t, args)
}

func TestSQLiteDialectRender_BoolLiteral(t *testing.T) {
	sql, args := q.Select(q.Literal(true), q.Literal(false)).
		Render(dialect.SQLite{})

	assert.Equal(t, "SELECT 1, 0", sql)
	assert.Empty(t, args)
}
