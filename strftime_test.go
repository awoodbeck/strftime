package strftime

import (
	"fmt"
	"testing"
	"time"
)

var (
	t1  = time.Date(2018, time.July, 9, 13, 14, 15, 0, time.UTC)
	t2  = time.Date(1950, time.December, 10, 4, 45, 59, 0, time.UTC)
	t3  = time.Date(2016, time.January, 1, 13, 14, 15, 0, time.UTC)
	t4  = time.Date(2015, time.January, 1, 13, 14, 15, 0, time.UTC)
	now = time.Now()

	tc = []struct {
		time     *time.Time
		format   string
		expected string
	}{
		{time: &t1, format: "%", expected: "%"},
		{time: &t1, format: "%%", expected: "%"},
		{time: &t1, format: "%%n", expected: "%n"},
		{time: &t1, format: "%%t", expected: "%t"},
		{time: &t1, format: "%n%t", expected: "\n\t"},
		{time: &t1, format: "%a", expected: "Mon"},
		{time: &t1, format: "%A", expected: "Monday"},
		{time: &t1, format: "%b", expected: "Jul"},
		{time: &t1, format: "%h", expected: "Jul"},
		{time: &t1, format: "%B", expected: "July"},
		{time: &t1, format: "%c", expected: "Mon Jul  9 13:14:15 2018"},
		{time: &t1, format: "%C", expected: "20"},
		{time: &t1, format: "%d", expected: "09"},
		{time: &t1, format: "%D", expected: "07/09/18"},
		{time: &t1, format: "%e", expected: " 9"},
		{time: &t1, format: "%F", expected: "2018-07-09"},
		{time: &t1, format: "%g", expected: "18"},
		{time: &t1, format: "%G", expected: "2018"},
		{time: &t1, format: "%H", expected: "13"},
		{time: &t1, format: "%I", expected: "01"},
		{time: &t1, format: "%j", expected: "190"},
		{time: &t1, format: "%k", expected: "13"},
		{time: &t2, format: "%k", expected: " 4"},
		{time: &t1, format: "%l", expected: " 1"},
		{time: &t2, format: "%l", expected: " 4"},
		{time: &t1, format: "%m", expected: "07"},
		{time: &t1, format: "%M", expected: "14"},
		{time: &t1, format: "%n", expected: "\n"},
		{time: &t1, format: "%p", expected: "PM"},
		{time: &t2, format: "%p", expected: "AM"},
		{time: &t1, format: "%P", expected: "pm"},
		{time: &t2, format: "%P", expected: "am"},
		{time: &t1, format: "%r", expected: "01:14:15 PM"},
		{time: &t2, format: "%r", expected: "04:45:59 AM"},
		{time: &t1, format: "%R", expected: "13:14"},
		{time: &t2, format: "%R", expected: "04:45"},
		{time: &t1, format: "%s", expected: "1531142055"},
		{time: &t1, format: "%S", expected: "15"},
		{time: &t1, format: "%t", expected: "\t"},
		{time: &t1, format: "%T", expected: "13:14:15"},
		{time: &t2, format: "%T", expected: "04:45:59"},
		{time: &t1, format: "%u", expected: "1"},
		{time: &t2, format: "%u", expected: "7"},
		{time: &t1, format: "%V", expected: "28"},
		{time: &t3, format: "%V", expected: "53"}, // 3 January days in this week.
		{time: &t4, format: "%V", expected: "01"}, // 4 January days in this week.
		{time: &t1, format: "%w", expected: "1"},
		{time: &t2, format: "%w", expected: "0"},
		{time: &t1, format: "%x", expected: "07/09/2018"},
		{time: &t1, format: "%X", expected: "13:14:15"},
		{time: &t2, format: "%X", expected: "04:45:59"},
		{time: &t1, format: "%y", expected: "18"},
		{time: &t1, format: "%Y", expected: "2018"},
		{time: &t1, format: "%z", expected: "+0000"},
		{time: &t1, format: "%Z", expected: "UTC"},
		{time: &t1, format: "foo", expected: "foo"},
		{time: &t1, format: "bar%", expected: "bar%"},
		{time: &t1, format: "%1", expected: "%1"},
		{time: nil, format: "%Y", expected: now.Format("2006")},
		{time: &t1, format: "%Y-%m-%dtest\n\t%Z", expected: "2018-07-09test\n\tUTC"},
	}
)

func TestFormat(t *testing.T) {
	for i, c := range tc {
		actual := Format(c.time, c.format)
		if actual != c.expected {
			t.Errorf("%d: expected: %q; actual: %q", i, c.expected, actual)
		}
	}
}

func ExampleFormat() {
	t := time.Date(2018, time.July, 9, 13, 14, 15, 0, time.UTC)

	fmt.Println(Format(&t, "%c"))
	fmt.Println(Format(&t, "%%Y-%%m-%%d == %Y-%m-%d"))
	fmt.Println(Format(&t, "%A is day number %w of the week."))
	fmt.Println(Format(&t, "Last century was%n the %Cth century."))
	fmt.Println(Format(&t, "Time zone: %Z"))

	// Output:
	// Mon Jul  9 13:14:15 2018
	// %Y-%m-%d == 2018-07-09
	// Monday is day number 1 of the week.
	// Last century was
	//  the 20th century.
	// Time zone: UTC
}
