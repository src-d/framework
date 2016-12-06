package utils

import (
	"time"

	. "gopkg.in/check.v1"
)

func (s *UtilsSuite) TestTimeWithTimezone_New(c *C) {
	loc, _ := time.LoadLocation("Europe/Berlin")

	d := time.Date(2015, time.July, 1, 0, 0, 0, 0, loc)
	t := NewTimeWithTimezone(d)

	c.Assert(t.Time.String(), Equals, d.String())
	c.Assert(t.Timezone, Equals, Timezone("+0200"))
}
