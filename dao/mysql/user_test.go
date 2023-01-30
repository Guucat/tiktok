package mysql

import (
	"testing"
)

func init() {
	Init()
}
func TestGetFollowCount(t *testing.T) {
	tests := []struct {
		Input int64
		Want  int64
	}{
		{1, 1},
		{2, 0},
		{0, 0},
	}
	for _, test := range tests {
		got := GetFollowCount(test.Input)
		if got != test.Want {
			t.Errorf("input: %d, got: %d, want: %d", test.Input, got, test.Want)
		}
	}
}

func TestGetFollowerCount(t *testing.T) {
	tests := []struct {
		Input int64
		Want  int64
	}{
		{1, 0},
		{2, 1},
		{0, 0},
	}
	for _, test := range tests {
		got := GetFollowerCount(test.Input)
		if got != test.Want {
			t.Errorf("input: %d, got: %d, want: %d", test.Input, got, test.Want)
		}
	}
}

func TestGetIsFollower(t *testing.T) {
	tests := []struct {
		You   int64
		Other int64
		Want  bool
	}{
		{1, 1, false},
		{1, 2, true},
		{0, 1, false},
	}
	for _, test := range tests {
		got := IsFollower(test.You, test.Other)
		if got != test.Want {
			t.Errorf("%d followed %d, got: %v, want: %v", test.You, test.Other, got, test.Want)
		}
	}
}

func TestIsExistById(t *testing.T) {
	tests := []struct {
		Input int64
		Want  string
	}{
		{1, "test"},
		{2, "tt"},
		{0, ""},
	}
	for _, test := range tests {
		got := IsExistById(test.Input)
		if got != test.Want {
			t.Errorf("input: %v, got: %v, want: %v", test.Input, got, test.Want)
		}
	}
}
