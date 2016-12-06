package utils

import (
	"time"

	"github.com/mcuadros/go-time-aggregator"

	. "gopkg.in/check.v1"
)

func (s *UtilsSuite) TestUniqueStringSlice_Append(c *C) {
	u := make(UniqueStringSlice, 0)

	c.Assert(u.Add("foo"), Equals, 1)
	c.Assert(u, DeepEquals, UniqueStringSlice{"foo"})

	c.Assert(u.Add("foo"), Equals, 0)
	c.Assert(u, DeepEquals, UniqueStringSlice{"foo"})

	c.Assert(u.Add("qux"), Equals, 1)
	c.Assert(u, DeepEquals, UniqueStringSlice{"foo", "qux"})
}

func (s *UtilsSuite) TestUniqueStringSlice_Delete(c *C) {
	u := make(UniqueStringSlice, 0)
	c.Assert(u.Add("foo", "qux", "bar"), Equals, 3)

	c.Assert(u.Delete("qux"), Equals, 1)
	c.Assert(u, DeepEquals, UniqueStringSlice{"foo", "bar"})
}

func (s *UtilsSuite) TestUniqueStringSlice_DeleteNotFound(c *C) {
	u := make(UniqueStringSlice, 0)
	u.Add("foo")
	u.Add("bar")

	c.Assert(u.Delete("qux"), Equals, 0)
	c.Assert(u, DeepEquals, UniqueStringSlice{"foo", "bar"})
}

func (s *UtilsSuite) TestUniqueObjectIdSlice_Append(c *C) {
	u := make(UniqueObjectIdSlice, 0)

	c.Assert(u.Add("foo"), Equals, 1)
	c.Assert(u, DeepEquals, UniqueObjectIdSlice{"foo"})

	c.Assert(u.Add("foo"), Equals, 0)
	c.Assert(u, DeepEquals, UniqueObjectIdSlice{"foo"})

	c.Assert(u.Add("qux"), Equals, 1)
	c.Assert(u, DeepEquals, UniqueObjectIdSlice{"foo", "qux"})
}

func (s *UtilsSuite) TestUniqueObjectIdSlice_Delete(c *C) {
	u := make(UniqueObjectIdSlice, 0)
	u.Add("foo")
	u.Add("qux")
	u.Add("bar")

	c.Assert(u.Delete("qux"), Equals, 1)
	c.Assert(u, DeepEquals, UniqueObjectIdSlice{"foo", "bar"})
}

func (s *UtilsSuite) TestUniqueObjectIdSlice_DeleteNotFound(c *C) {
	u := make(UniqueObjectIdSlice, 0)
	u.Add("foo")
	u.Add("bar")

	c.Assert(u.Delete("qux"), Equals, 0)
	c.Assert(u, DeepEquals, UniqueObjectIdSlice{"foo", "bar"})
}

func (s *UtilsSuite) TestUniqueObjectIdSlice_IgnoreEmpty(c *C) {
	u := make(UniqueStringSlice, 0)
	c.Assert(u.Add(""), Equals, 0)
	c.Assert(len(u), Equals, 0)
}

func (p *UtilsSuite) TestLanguagesTimeAggregator(c *C) {
	var lta1 = NewTimeAggregatorMap(aggregator.Year, aggregator.YearDay)
	var lta2 = NewTimeAggregatorMap(aggregator.Year, aggregator.YearDay)

	jan1_15 := time.Date(2015, time.January, 1, 12, 0, 0, 0, time.UTC)

	lta1.Add("Python", jan1_15, 10)
	lta2.Add("Python", jan1_15, 20)
	lta1.Add("Go", jan1_15, 100)
	lta2.Add("Go", jan1_15, 200)

	taPython1 := lta1.Values["Python"]
	taPython2 := lta2.Values["Python"]
	taGo1 := lta1.Values["Go"]
	taGo2 := lta2.Values["Go"]

	c.Assert(taPython1.Get(jan1_15), Equals, int64(10))
	c.Assert(taPython2.Get(jan1_15), Equals, int64(20))
	c.Assert(taGo1.Get(jan1_15), Equals, int64(100))
	c.Assert(taGo2.Get(jan1_15), Equals, int64(200))

	var err error
	err = taPython1.Sum(taPython2)
	c.Assert(err, IsNil)
	c.Assert(taPython1.Get(jan1_15), Equals, int64(30))

	err = taGo1.Sum(taGo2)
	c.Assert(err, IsNil)
	c.Assert(taGo1.Get(jan1_15), Equals, int64(300))
}
