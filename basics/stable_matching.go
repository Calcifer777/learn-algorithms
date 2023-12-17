package basics

import (
	"fmt"
	"log/slog"
	"strings"
)

func match(mPref []*Party, wPref []*Party) {
}

type Party struct {
	name     string
	prefs    []*Party
	matched  *Party
	proposed []string
}

func NewParty(name string, prefs []*Party) Party {
	return Party{name, prefs, nil, make([]string, 0)}
}

func (p *Party) String() string {
	matched := "n.d."
	if p.matched != nil {
		matched = p.matched.name
	}
	prefs := make([]string, 0)
	for _, p := range p.prefs {
		prefs = append(prefs, p.name)
	}
	return fmt.Sprintf(
		"P{n: %s, prefs: %s, m: %s}",
		p.name,
		strings.Join(prefs, ", "),
		matched,
	)
}

func (p *Party) AddPref(pref *Party) {
	(*p).prefs = append(p.prefs, pref)
}

func Contains(ps []*Party, name string) (*Party, bool) {
	for _, p := range ps {
		if p.name == name {
			return p, true
		}
	}
	return nil, false

}
func InitParties(
	mPref map[string][]string,
	wPref map[string][]string,
) ([]*Party, []*Party) {

	men := make([]*Party, 0)
	women := make([]*Party, 0)

	// Init mens
	for mName, p := range mPref {
		ws := make([]*Party, 0)
		for _, wName := range p {
			w, ok := Contains(women, wName)
			if !ok {
				w_ := NewParty(wName, nil)
				women = append(women, &w_)
				w = &w_
			}
			ws = append(ws, w)
		}
		m := NewParty(mName, ws)
		men = append(men, &m)
	}

	// Init womens
	for wName, p := range wPref {
		w, ok := Contains(women, wName)
		if !ok {
			panic(fmt.Errorf("%s not found", wName))
		}
		for _, mName := range p {
			m, ok := Contains(men, mName)
			if !ok {
				panic(fmt.Errorf("%s not found", mName))
			}
			w.AddPref(m)
		}
		fmt.Println("")
	}

	for _, m := range men {
		slog.Info("init", slog.String("m", m.String()))
	}
	for _, w := range women {
		slog.Info("init", slog.String("w", w.String()))
	}
	return men, women
}

func IsIn[T comparable](arr []T, needle T) bool {
	for _, t := range arr {
		if needle == t {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](arr []T, needle T) int {
	for i, t := range arr {
		if t == needle {
			return i
		}
	}
	return -1
}

func StableMactch(men []*Party, women []*Party) {
	unmatched := true
	for unmatched {
		unmatched = false
		for _, m := range men {
			if m.matched != nil {
				continue
			}
			unmatched = true
		inner:
			for _, w := range m.prefs {
				if IsIn(m.proposed, w.name) {
					continue
				}
				m.proposed = append(m.proposed, w.name)
				if w.matched == nil {
					w.matched = m
					m.matched = w
					break inner
				} else if IndexOf(w.prefs, m) < IndexOf(w.prefs, w.matched) {
					w.matched.matched = nil
					m.matched = w
					w.matched = m
					break inner
				}
			}
		}
	}
}
