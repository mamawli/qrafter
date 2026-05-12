package qrafter

import (
	"github.com/SennovE/qrafter/internal/core"
	"github.com/SennovE/qrafter/internal/expr"
)

type Aggregate struct {
	Expression
}

var _ = (core.Aggregater)(Aggregate{})

func newAggregate(s core.Selecter) Aggregate {
	return Aggregate{Expression: newExpression(s)}
}

func (a Aggregate) Aggregate() {}

func (a Aggregate) As(alias string) Aggregate {
	return newAggregate(expr.Alias(a.selecter, alias))
}

func AggregateFunc(name string, args ...any) Aggregate {
	return newAggregate(expr.Function(name, asSelecters(args)...))
}

func Count(args ...any) Aggregate {
	if len(args) == 0 {
		return AggregateFunc("COUNT", Star())
	}
	return AggregateFunc("COUNT", args...)
}

func Sum(v any) Aggregate {
	return AggregateFunc("SUM", v)
}

func Avg(v any) Aggregate {
	return AggregateFunc("AVG", v)
}

func Min(v any) Aggregate {
	return AggregateFunc("MIN", v)
}

func Max(v any) Aggregate {
	return AggregateFunc("MAX", v)
}
