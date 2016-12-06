package utils

import (
	"time"

	"github.com/mcuadros/go-time-aggregator"

	"gopkg.in/mgo.v2/bson"
)

type UniqueStringSlice []string

func (u *UniqueStringSlice) Add(entries ...string) int {
	count := 0

	for _, entry := range entries {
		if u.add(entry) {
			count++
		}
	}

	return count
}

func (u *UniqueStringSlice) add(e string) bool {
	if e == "" {
		return false
	}
	for _, v := range *u {
		if v == e {
			return false
		}
	}

	*u = UniqueStringSlice(append(*u, e))
	return true
}

func (u *UniqueStringSlice) Delete(entries ...string) int {
	count := 0

	for _, entry := range entries {
		if u.delete(entry) {
			count++
		}
	}

	return count
}

func (u *UniqueStringSlice) delete(e string) bool {
	s := *u

	index := -1
	for i, v := range *u {
		if v == e {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	*u = UniqueStringSlice(append(s[:index], s[index+1:]...))
	return true
}

type UniqueObjectIdSlice []bson.ObjectId

func (u *UniqueObjectIdSlice) Add(ids ...bson.ObjectId) int {
	count := 0
	for _, id := range ids {
		if u.add(id) {
			count++
		}
	}

	return count
}

func (u *UniqueObjectIdSlice) add(e bson.ObjectId) bool {
	if e == "" {
		return false
	}
	for _, v := range *u {
		if v == e {
			return false
		}
	}

	*u = UniqueObjectIdSlice(append(*u, e))
	return true
}

func (u *UniqueObjectIdSlice) Delete(ids ...bson.ObjectId) int {
	count := 0

	for _, id := range ids {
		if u.delete(id) {
			count++
		}
	}

	return count
}

func (u *UniqueObjectIdSlice) delete(e bson.ObjectId) bool {
	s := *u

	index := -1
	for i, v := range *u {
		if v == e {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	*u = UniqueObjectIdSlice(append(s[:index], s[index+1:]...))
	return true
}

func CastStringSliceToIfaceSlice(ss []string) []interface{} {
	var in []interface{}
	for _, s := range ss {
		in = append(in, s)
	}
	return in
}

type TimeAggregatorMap struct {
	u      []aggregator.Unit
	Values map[string]*aggregator.TimeAggregator
}

func NewTimeAggregatorMap(units ...aggregator.Unit) *TimeAggregatorMap {
	return &TimeAggregatorMap{
		u:      units,
		Values: make(map[string]*aggregator.TimeAggregator, 0),
	}
}

func (m *TimeAggregatorMap) Sum(key string, a *aggregator.TimeAggregator) {
	m.Get(key).Sum(a)
}

func (m *TimeAggregatorMap) Get(key string) *aggregator.TimeAggregator {
	a, ok := m.Values[key]
	if !ok {
		a, _ = aggregator.NewTimeAggregator(m.u...)
		m.Values[key] = a
	}

	return a
}

func (m *TimeAggregatorMap) Add(key string, when time.Time, bytes int64) {
	m.Get(key).Add(when, bytes)
}
