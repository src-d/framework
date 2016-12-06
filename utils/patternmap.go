package utils

import (
	"regexp"
	"strings"
)

type PatternGetter func(string) string
type PatternMap struct {
	Default PatternGetter
	cases   map[string]string
}

func NewPatternMap(def PatternGetter) *PatternMap {
	return &PatternMap{
		Default: def,
		cases:   map[string]string{},
	}
}

func (m *PatternMap) Get(s string) string {
	p, ok := m.cases[s]
	if ok {
		return p
	}
	return m.Default(s)
}

func (m *PatternMap) MustGetRegexp(s string) *regexp.Regexp {
	pat := m.Get(s)

	return regexp.MustCompile(pat)
}

func (m *PatternMap) Set(s string, p string) *PatternMap {
	m.cases[s] = p

	return m
}

type PatternBuilder struct{}

func (b *PatternBuilder) Default() *PatternMap {
	return NewPatternMap(b.DefaultPatternGetter())
}

func (b *PatternBuilder) DefaultPatternGetter() PatternGetter {
	return b.DefaultPattern
}

func (b *PatternBuilder) DefaultPattern(s string) string {
	return `\b` + s + `\b`
}

func (b *PatternBuilder) JoinPatterns(ps ...string) string {
	withParens := make([]string, 0, len(ps))
	for _, p := range ps {
		withParens = append(withParens, "("+p+")")
	}

	joined := `(?i)(` + strings.Join(withParens, "|") + `)`

	return joined
}

func (b *PatternBuilder) JoinDefaults(ns ...string) string {
	defs := make([]string, 0, len(ns))
	getter := b.DefaultPatternGetter()
	for _, n := range ns {
		defs = append(defs, getter(n))
	}
	joined := b.JoinPatterns(defs...)

	return joined
}
