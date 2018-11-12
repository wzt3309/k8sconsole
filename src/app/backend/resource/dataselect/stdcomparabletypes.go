package dataselect

import "strings"

type StdComparableInt int

func (self StdComparableInt) Compare(otherV ComparableValue) int {
	other := otherV.(StdComparableInt)
	return intsCompare(int(self), int(other))
}

func (self StdComparableInt) Contains(otherV ComparableValue) bool {
	return self.Compare(otherV) == 0
}

type StdComparableString string

func (self StdComparableString) Compare(otherV ComparableValue) int {
	other := otherV.(StdComparableString)
	return strings.Compare(string(self), string(other))
}

func (self StdComparableString) Contains(otherV ComparableValue) bool {
	other := otherV.(StdComparableString)
	return strings.Contains(string(self), string(other))
}

func intsCompare(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}

func ints64Compare(a, b int64) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}