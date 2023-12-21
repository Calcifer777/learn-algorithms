package greedy

import (
  "testing"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
  slots := []Interval{
    Interval{10, 13},
    Interval{2, 7},
    Interval{0, 1},
    Interval{3, 5},
    Interval{14, 19},
    Interval{8, 10},
    Interval{16,17},
  }
  target := Interval{0, 20}
  scheduled := Schedule(target, slots)
  expected := []Interval{
    Interval{0,1},
    Interval{3,5},
    Interval{8,10},
    Interval{10,13},
    Interval{16,17},
  }
  assert.ElementsMatch(t,
    expected,
    scheduled,
  )
}

func TestColor(t *testing.T) {
  slots := []Interval{
    Interval{0, 1},
    Interval{0, 1},
    Interval{0, 2},
    Interval{1, 2},
    Interval{1, 4},
    Interval{3, 5},
    Interval{3, 5},
    Interval{4,7},
    Interval{6,7},
    Interval{6,7},
  }
  scheduled := Color(slots)
  expected := []int{1, 2, 3, 1, 2, 1, 3, 2, 1, 3}
  assert.Equal(t,
    expected,
    scheduled,
  )
}
