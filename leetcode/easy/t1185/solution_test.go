package t1185

import "testing"

func TestDayOfTheWeek(t *testing.T) {
	tests := []struct {
		day   int
		month int
		year  int
		week  string
	}{
		{31, 8, 2019, "Saturday"},
		{18, 7, 1999, "Sunday"},
		{15, 8, 1993, "Sunday"},
	}

	for _, test := range tests {
		got := dayOfTheWeek(test.day, test.month, test.year)
		if got != test.week {
			t.Errorf("dayOfTheWeek(%d, %d, %d) = %s, want %s", test.day, test.month, test.year, got, test.week)
		}
	}
}
