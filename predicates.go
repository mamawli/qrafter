package qrafter

import (
	"strings"

	"github.com/SennovE/qrafter/dialect"
	"github.com/SennovE/qrafter/internal/core"
	"github.com/SennovE/qrafter/internal/pred"
)

type Predicate struct {
	predicater core.Predicater
}

var _ = (core.Predicater)(Predicate{})

func newPredicate(p core.Predicater) Predicate {
	return Predicate{predicater: p}
}

func unwrapPredicates(ps []core.Predicater) []core.Predicater {
	res := make([]core.Predicater, len(ps))
	for i, p := range ps {
		if wrapped, ok := p.(Predicate); ok {
			res[i] = wrapped.predicater
			continue
		}
		res[i] = p
	}
	return res
}

func (p Predicate) Predicate() {}

func (p Predicate) Render(w *strings.Builder, d dialect.DialectRenderer) {
	p.predicater.Render(w, d)
}

func (p Predicate) Tables() core.TablesSet {
	return p.predicater.Tables()
}

func (p Predicate) Precedence() int {
	if prec, ok := p.predicater.(core.Precedencer); ok {
		return prec.Precedence()
	}
	return core.PrecedenceComparison
}

func And(ps ...core.Predicater) Predicate {
	return newPredicate(pred.Logical(pred.OpAnd, unwrapPredicates(ps)...))
}

func Or(ps ...core.Predicater) Predicate {
	return newPredicate(pred.Logical(pred.OpOr, unwrapPredicates(ps)...))
}

func (e Expression) Lt(v any) Predicate {
	return newPredicate(pred.Binary("<", e.selecter, asSelecter(v)))
}

func (e Expression) Gt(v any) Predicate {
	return newPredicate(pred.Binary(">", e.selecter, asSelecter(v)))
}

func (e Expression) Le(v any) Predicate {
	return newPredicate(pred.Binary("<=", e.selecter, asSelecter(v)))
}

func (e Expression) Ge(v any) Predicate {
	return newPredicate(pred.Binary(">=", e.selecter, asSelecter(v)))
}

func (e Expression) Eq(v any) Predicate {
	return newPredicate(pred.Binary("=", e.selecter, asSelecter(v)))
}
