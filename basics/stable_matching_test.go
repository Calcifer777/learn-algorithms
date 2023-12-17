package basics

import (
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

func TestStableMatching1(t *testing.T) {
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
		slog.Info(m.String())
	}

}

func TestStableMatching2(t *testing.T) {
	mPrefs := map[string][]string{
		"0": {"7", "5", "6", "4"},
		"1": {"5", "4", "6", "7"},
		"2": {"4", "5", "6", "7"},
		"3": {"4", "5", "6", "7"},
	}
	wPrefs := map[string][]string{
		"4": {"0", "1", "2", "3"},
		"5": {"0", "1", "2", "3"},
		"6": {"0", "1", "2", "3"},
		"7": {"0", "1", "2", "3"},
	}
	men, women := InitParties(mPrefs, wPrefs)
	StableMactch(men, women)
	for _, m := range men {
		slog.Info(m.String())
	}

}
