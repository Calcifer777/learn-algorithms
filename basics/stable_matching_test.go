package basics

import (
	"fmt"
	"log/slog"
	"testing"
)

// const menPreferences = {
//   A: 'YXZ',
//   B: 'ZYX',
//   C: 'XZY'
// }

// const womenPreferences = {
//   X: 'BAC',
//   Y: 'CBA',
//   Z: 'ACB'
// }

func TestStableMatching(t *testing.T) {
	mPrefs := map[string][]string{
		"A": {"Y", "X", "Z"},
		"B": {"Y", "Z", "X"},
		"C": {"Z", "Y", "X"},
	}
	wPrefs := map[string][]string{
		"X": {"B", "A", "C"},
		"Y": {"C", "B", "A"},
		"Z": {"A", "C", "B"},
	}
	men, women := InitParties(mPrefs, wPrefs)
	StableMactch(men, women)
	for _, m := range men {
		slog.Info(fmt.Sprintf("m: %s, w: %s", m.name, m.matched.name))
	}

}
