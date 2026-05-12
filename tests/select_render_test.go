package tests

import (
	"testing"

	"github.com/SennovE/qrafter"
	"github.com/SennovE/qrafter/dialect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectRender_ParenthesizesLowerPrecedencePredicate(t *testing.T) {
	u := User{}
	require.NoError(t, qrafter.Bind(&u))

	query := qrafter.Select(u.UserName).Where(
		qrafter.And(
			u.UserName.Eq("ABC"),
			qrafter.Or(
				u.Age.Ge("1"),
				qrafter.Const("Test").Eq(u.UserName),
			),
		),
	)

	assert.Equal(
		t,
		`SELECT "table"."user_name" FROM "table" WHERE "table"."user_name" = 'ABC' AND ("table"."userAge" >= '1' OR 'Test' = "table"."user_name")`,
		query.Render(dialect.PostgreSQL{}),
	)
}

func TestSelectRender_ParenthesizesLowerPrecedenceExpression(t *testing.T) {
	u := User{}
	require.NoError(t, qrafter.Bind(&u))

	query := qrafter.Select(
		u.Age.Add(1).Mul(2),
	)

	assert.Equal(
		t,
		`SELECT ("table"."userAge" + 1) * 2 FROM "table"`,
		query.Render(dialect.PostgreSQL{}),
	)
}

func TestSelectRender_ParenthesizesRightPeerForNonAssociativeExpression(t *testing.T) {
	query := qrafter.Select(
		qrafter.Const(10).Sub(qrafter.Const(7).Sub(3)),
	)

	assert.Equal(t, `SELECT 10 - (7 - 3)`, query.Render(dialect.PostgreSQL{}))
}
