package dynprog

import (
	"slices"
)

type Interval struct {
	start, finish, value int
}

type NamedInterval struct {
	id int
	*Interval
}

func NewInterval(id, start, finish, value int) NamedInterval {
	return NamedInterval{id, &Interval{start, finish, value}}
}

type Schedule struct {
	intervals []NamedInterval
}

func NewSchedule(intervals []Interval) (schedule *Schedule) {
	s := Schedule{make([]NamedInterval, 0)}
	for _, i := range intervals {
		s.Add(i)
	}
	schedule = &s
	return
}

// Add interval to schedule, sorting by finishing time
func (s *Schedule) Add(i Interval) {
	idx := len(s.intervals)
	s.intervals = append(s.intervals, NamedInterval{idx, &i})
	slices.SortFunc(s.intervals, func(i1, i2 NamedInterval) int {
		return i1.finish - i2.finish
	})
}

func (s *Schedule) Opt() int {
	prevDisjoints := GetPrevDisjoints(s.intervals)
	memo := make([]int, 0)
	var inner func(idx int) (v int)
	inner = func(idx int) (v int) {
		if idx == 0 {
			return 0
		} else if idx < len(memo) {
			return memo[idx]
		} else {
			out := max(
				inner(idx-1),
				s.intervals[idx].value+prevDisjoints[idx],
			)
			memo = append(memo, out)
			return out
		}
	}
	return inner(len(s.intervals) - 1)
}

func GetPrevDisjoints(ints []NamedInterval) []int {
	out := make([]int, len(ints))
	for i := range ints {
		pIdx := 0
		for j := 0; j < i; j++ {
			if ints[j].finish <= ints[i].start {
				pIdx = j
			}
		}
		out[i] = pIdx
	}
	return out
}
