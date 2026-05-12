package core

import (
	"sort"
	"strings"

	"github.com/SennovE/qrafter/dialect"
)

type ColumnBinder interface {
	Bind(name string, table TableRef)
}

type TableRef struct {
	Name  string
	Alias string
}

var _ = (Renderer)(TableRef{})

type TablesSet = map[TableRef]struct{}

func (t TableRef) SQLName() string {
	if t.Alias == "" {
		return t.Name
	}
	return t.Alias
}

func (t TableRef) Render(w *strings.Builder, d dialect.DialectRenderer) {
	if t.Alias == "" {
		w.WriteString(d.QuoteIdent(t.Name))
	} else {
		w.WriteString(d.QuoteIdent(t.Name))
		w.WriteString(" AS ")
		w.WriteString(d.QuoteIdent(t.Alias))
	}
}

func GetSortedTables(tables TablesSet) []TableRef {
	sortedTables := make([]TableRef, 0, len(tables))
	for table := range tables {
		sortedTables = append(sortedTables, table)
	}
	sort.Slice(sortedTables, func(i, j int) bool {
		if sortedTables[i].Name == sortedTables[j].Name {
			return sortedTables[i].Alias < sortedTables[j].Alias
		}
		return sortedTables[i].Name < sortedTables[j].Name
	})
	return sortedTables
}
