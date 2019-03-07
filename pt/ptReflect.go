package pt

import (
	"fmt"
	"reflect"
	"runtime"
)

var RefPt = func(inter interface{}) {
	v := reflect.ValueOf(inter)
	k := v.Kind()
	t := v.Type()

	Ptf("[%v][%v][%v]\n", k, t, v)
}

var RefPtl = func(inter interface{}) {
	_, _, line, _ := runtime.Caller(1)

	v := reflect.ValueOf(inter)
	k := v.Kind()
	t := v.Type()

	Ptf("l(%d): [%v][%v][%v]\n", line, k, t, v)
}

var RefPtn = func(inter interface{}) {
	v := reflect.ValueOf(inter)
	k := v.Kind()
	t := v.Type()

	Ptn(k)
	Ptn(t)
	Ptn(v)
}

var RefPtln = func(inter interface{}) {
	_, _, line, _ := runtime.Caller(1)
	fmt.Printf("l(%d): \n", line)

	v := reflect.ValueOf(inter)
	k := v.Kind()
	t := v.Type()

	Ptn(k)
	Ptn(t)
	Ptn(v)
}
