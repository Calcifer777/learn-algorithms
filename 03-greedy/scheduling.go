package greedy

import (
  "fmt"
  "log/slog"
)

type Interval struct {
  s, f int
}

func StartBefore(i1, i2 Interval) bool {
  return i1.s < i2.s || (i1.s == i2.s && i1.f <= i2.f)
}

func  EndsBefore(i1, i2 Interval) bool {
  return i1.f < i2.f
}

func Overlap(i1, i2 Interval) bool {
  var prev, next Interval
  if i1.s < i2.s {
    prev = i1
    next = i2
  } else {
    prev = i2
    next = i1
  }
  return prev.f > next.s
}

func Schedule(target Interval, slots []Interval) []Interval {
  out := make([]Interval, 0)
  if len(slots) == 0 {
    return out
  }
  slotsSorted := SortIntervals(slots, EndsBefore)
  out = append(out, slotsSorted[0])
  for _, i := range slotsSorted {
    if i.s < out[len(out)-1].f {
      continue
    } else {
      out = append(out, i)
    }
  }
  return out
}

func Color(slots []Interval) []int {
  slotsSorted := SortIntervals(slots, StartBefore)
  colors := make([]int, len(slotsSorted))
  var excludedLabels map[int]bool
  for j, slot := range slotsSorted {
    slog.Info(fmt.Sprintf("assigning label to %v\n", slot))
    excludedLabels = make(map[int]bool)
    for i, cmp := range slotsSorted[:j] {
      slog.Info(fmt.Sprintf("\tComparing with %v\n", cmp))
      if Overlap(slot, cmp) {
        slog.Info(fmt.Sprintf("\t\tOverlaps! Excluding: %d\n", colors[i]))
        excludedLabels[colors[i]] = true
      }
    }
    slog.Info(fmt.Sprintf("\tExcluded labels for %v: %v\n", slot, excludedLabels))
    labelLoop:
    for label := 1; label <= len(slotsSorted); label++ {
      if !excludedLabels[label] {
        slog.Info(fmt.Sprintf("\tassigned label %d\n", label))
        colors[j] = label
        break labelLoop
      }
    }
  }
  return colors
}

func SortIntervals(is []Interval, cmp func(i1, i2 Interval) bool) []Interval {
  if len(is) <= 1 {
    return is
  } else {
    ref := is[0]
    larger, smaller := make([]Interval, 0), make([]Interval, 0)
    for _, i := range is[1:] {
      if cmp(ref, i) {
        larger = append(larger, i)
      } else {
        smaller = append(smaller, i)
      }
    }
    return append(
      append(SortIntervals(smaller, cmp), ref),
      SortIntervals(larger, cmp)...,
    )

  }
}

