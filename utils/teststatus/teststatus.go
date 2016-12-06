package teststatus

import (
	"reflect"
	"strings"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"srcd.works/framework/utils"
)

type Case struct {
	Code          interface{}
	Constructor   interface{}
	ValidPayloads []interface{}
	BadPayloads   []interface{}
}

func (cas *Case) Check(c *C) {
	if len(cas.ValidPayloads) == 0 {
		cas.CheckStatus(c, cas.Construct(), nil)
	}
	for _, pl := range cas.ValidPayloads {
		cas.CheckStatus(c, cas.Construct(pl), pl)
	}
}

func (cas *Case) Construct(payload ...interface{}) interface{} {
	args := []reflect.Value{}
	for _, p := range payload {
		args = append(args, reflect.ValueOf(p))
	}
	out := reflect.ValueOf(cas.Constructor).Call(args)
	pst := reflect.New(out[0].Type())
	reflect.Indirect(pst).Set(out[0])
	return pst.Interface()
}

func (cas *Case) CheckStatus(c *C, st interface{}, pl interface{}) {
	cas.CheckCode(c, st, cas.Code)

	cas.CheckPayload(c, st, pl)
	for _, nextPl := range cas.ValidPayloads {
		err := cas.SetPayload(st, nextPl)
		c.Assert(err, IsNil)
		cas.CheckPayload(c, st, nextPl)
		cas.CheckBSON(c, st)
	}

	cas.SetPayload(st, pl)
	for _, bad := range cas.BadPayloads {
		err := cas.SetPayload(st, bad)
		c.Assert(err, Not(IsNil))
		cas.CheckPayload(c, st, pl)
	}

	cas.CheckIfMethods(c, st)
}

func (cas *Case) CheckCode(c *C, st interface{}, code interface{}) {
	c.Assert(reflect.ValueOf(st).Elem().FieldByName("Code").Interface(), Equals, code)
}

func (cas *Case) CheckPayload(c *C, st interface{}, payload interface{}) {
	c.Assert(reflect.ValueOf(st).Elem().FieldByName("Payload").Interface(), DeepEquals, payload)
}

func (cas *Case) CheckBSON(c *C, st interface{}) {
	bs, err := bson.Marshal(st)
	c.Assert(err, IsNil)
	rst := reflect.ValueOf(st)
	out := reflect.New(rst.Type())
	err = bson.Unmarshal(bs, out.Interface())
	c.Assert(err, IsNil)
	c.Assert(reflect.Indirect(reflect.Indirect(out)).Interface(), DeepEquals, reflect.Indirect(rst).Interface())
}

func (cas *Case) SetPayload(st interface{}, payload ...interface{}) error {
	args := []reflect.Value{}
	for _, p := range payload {
		args = append(args, reflect.ValueOf(&p).Elem())
	}
	out := reflect.ValueOf(st).MethodByName("SetPayload").Call(args)
	ret := out[0].Interface()
	if ret == nil {
		return nil
	}
	return ret.(error)
}

func (cas *Case) CheckIfMethods(c *C, st interface{}) {
	rst := reflect.ValueOf(st)
	for i := 0; i < rst.NumMethod(); i++ {
		mt := rst.Type().Method(i)
		ifIdx := strings.Index(mt.Name, "If")
		if ifIdx < 0 || mt.Type.NumOut() != 2 || mt.Type.Out(1).Kind() != reflect.Bool {
			continue
		}
		out := rst.Method(i).Call(nil)
		stName := mt.Name[ifIdx+2:]
		if utils.CamelToSnake(stName) == rst.Elem().FieldByName("Code").String() {
			c.Assert(out[1].Bool(), Equals, true)
			c.Assert(out[0].Interface(), DeepEquals, rst.Elem().FieldByName("Payload").Interface())
		}
	}
}
