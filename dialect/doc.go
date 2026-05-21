// Package dialect contains SQL rendering dialects for qrafter queries.
//
// The included dialects cover syntax that qrafter can currently vary safely:
// identifier quoting, literals, placeholders, and LIMIT/OFFSET rendering.
// Feature-level support is still query-dependent; for example MySQL needs
// dialect-specific rendering before PostgreSQL-style RETURNING, UPDATE ... FROM,
// DELETE ... USING, or NULLS FIRST/LAST can be considered portable.
package dialect
