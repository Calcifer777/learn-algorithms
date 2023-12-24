package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	slots := []Interval{
		{10, 13},
		{2, 7},
		{0, 1},
		{3, 5},
		{14, 19},
		{8, 10},
		{16, 17},
	}
	target := Interval{0, 20}
	scheduled := Schedule(target, slots)
	expected := []Interval{
		{0, 1},
		{3, 5},
		{8, 10},
		{10, 13},
		{16, 17},
	}
	assert.ElementsMatch(t,
		expected,
		scheduled,
	)
}

func TestColor(t *testing.T) {
	slots := []Interval{
		{0, 1},
		{0, 1},
		{0, 2},
		{1, 2},
		{1, 4},
		{3, 5},
		{3, 5},
		{4, 7},
		{6, 7},
		{6, 7},
	}
	scheduled := Color(slots)
	expected := []int{1, 2, 3, 1, 2, 1, 3, 2, 1, 3}
	assert.Equal(t,
		expected,
		scheduled,
	)
}

func TestJobSchedule(t *testing.T) {
	slots := []Job{
		{2, 4},
		{3, 6},
		{1, 2},
	}
	scheduled := EarliestDeadlineFirst(slots)
	expected := []JobSchedule{
		{Job{1, 2}, 0},
		{Job{2, 4}, 1},
		{Job{3, 6}, 3},
	}
	assert.ElementsMatch(t,
		expected,
		scheduled,
	)
}
