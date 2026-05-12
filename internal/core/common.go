package core

import (
	"strings"

	"github.com/SennovE/qrafter/dialect"
)

const (
	PrecedenceOr = iota + 1
	PrecedenceAnd
	PrecedenceComparison
	PrecedenceAdditive
	PrecedenceMultiplicative
)

type Renderer interface {
	Render(w *strings.Builder, d dialect.DialectRenderer)
}

type Precedencer interface {
	Precedence() int
}

type Selecter interface {
	Renderer
	Tables() TablesSet
}

type Predicater interface {
	Selecter
	Predicate()
}

func RenderChild(r Renderer, parentPrecedence int, parenthesizeOnEqual bool, w *strings.Builder, d dialect.DialectRenderer) {
	child, ok := r.(Precedencer)
	if !ok {
		r.Render(w, d)
		return
	}

	childPrecedence := child.Precedence()
	if childPrecedence < parentPrecedence || childPrecedence == parentPrecedence && parenthesizeOnEqual {
		w.WriteString("(")
		r.Render(w, d)
		w.WriteString(")")
		return
	}

	r.Render(w, d)
}
