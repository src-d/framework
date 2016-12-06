package tests

import (
	"fmt"
	"reflect"

	"gopkg.in/check.v1"
)

type elementsEqual struct {
	*check.CheckerInfo
}

var ElementsEqual = &elementsEqual{
	&check.CheckerInfo{"ElementsEqual", []string{"list", "expected"}},
}

func (e *elementsEqual) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "Need exactly two parameters"
	}

	p1 := reflect.ValueOf(params[0])
	p2 := reflect.ValueOf(params[1])

	if !e.checkKind(p1) || !e.checkKind(p2) {
		return false, "Parameters must be lists or arrays"
	}

	len1 := p1.Len()
	len2 := p2.Len()

	if len1 != len2 {
		msg := fmt.Sprintf(
			"Lengths not equal for actual %v and expected %v",
			params[0], params[1],
		)
		return false, msg
	}

	m := map[interface{}]bool{}
	for i := 0; i < len1; i += 1 {
		x := p1.Index(i).Interface()
		m[x] = true
	}

	for i := 0; i < len2; i += 1 {
		x := p2.Index(i).Interface()
		if !m[x] {
			msg := fmt.Sprintf("Element `%v` expected and not present in %v", x, params[0])
			return false, msg
		}
	}

	return true, ""
}

func (e *elementsEqual) checkKind(v reflect.Value) bool {
	k := v.Kind()

	return (k == reflect.Array) || (k == reflect.Slice)
}
