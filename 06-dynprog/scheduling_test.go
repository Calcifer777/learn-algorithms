package dynprog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var schedule *Schedule = NewSchedule([]Interval{
	{0, 1, 1},
	{0, 2, 2},
	{2, 3, 1},
	{3, 4, 1},
	{2, 5, 1},
	{5, 8, 5},
})

func TestPrevDisjoints(t *testing.T) {
	out := GetPrevDisjoints(schedule.intervals)
	var expected []int = []int{
		0,
		0,
		1,
		2,
		1,
		4,
	}
	assert.ElementsMatch(t, expected, out)
}

func TestSchedule(t *testing.T) {
	out := schedule.Opt()
	expected := 9
	assert.Equal(t, expected, out)
}
