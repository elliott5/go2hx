package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("start")

	// Code below highlights what works so far in the reflect prototype
	str := "hitchhiker"
	showType("42", 42, false, false)
	showType("uint8(42)", uint8(42), false, false)
	showType("str", str, false, false)
	showType("&str", &str, false, false)
	spt := reflect.PtrTo(reflect.TypeOf(str))
	fmt.Println(spt)
	showType("anonymous", struct {
		A int32
		b string
		K reflect.Kind
	}{42, str, 4}, true, false)

	m := MyType{
		X: 42,
		y: str,
	}
	showType("MyType", m, true, true)
	mt := reflect.TypeOf(m)
	showType("reflect.Type", mt, false, true)

	var k reflect.Kind = 5
	showType("reflect.Kind", k, false, true)

	nst := NonStructT(42)                     // needs different code to be generated by go2hx to find type
	showType("NonStructT", nst, false, false) // TODO can't see the methods yet

	// line below does not work in Go, it would neet to be generated by go2hx
	var ms = makeNamedGoType("interface", "github_com.pxshadow.go4hx.rnd.MyStringer")

	showType("MyStringer", ms, false, true)

	fmt.Println("MyType implements MyStringer", mt.Implements(ms))                       // ...or not!                    // this is the key requirement for a type switch
	fmt.Println("reflect.Type implements MyStringer", reflect.TypeOf(mt).Implements(ms)) // this is the key requirement for a type switch
	fmt.Println("reflect.Kind implements MyStringer", reflect.TypeOf(k).Implements(ms))  // this is the key requirement for a type switch

	A := [2]uint64{0, 1}
	showType("array", A, false, false)

	S := []bool{true, false}
	showType("slice", S, false, false)

	Map := map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	showType("map", Map, false, false)

	c := make(chan int8)
	showType("chan", c, false, false)

	showType("func", showType, false, false)

	//testHarness()
}

func showType(name string, i interface{}, isStruct, isNamed bool) {
	Ti := reflect.TypeOf(i)
	Ki := Ti.Kind()
	Kui := uint(Ki)
	Kn := Ki.String()
	fmt.Println(name, Ti.String(), Kui, Kn)
	if isStruct {
		fmt.Println("fields:")
		for i := 0; i < Ti.NumField(); i++ {
			fmt.Println("=", Ti.Field(i))
		}
	}
	if isNamed {
		fmt.Println("methods:")
		for i := 0; i < Ti.NumMethod(); i++ {
			fmt.Println("-", Ti.Method(i))
		}
	}
}

type Vertex struct {
	Lat, Long float64
}

type MyType struct {
	X uint8
	y string
}

// MyType does not implement the MyStringer interface
func (mt MyType) StringX() string {
	return mt.y
}

type MyStringer interface {
	String() string
}

type NonStructT uint8

func (nst NonStructT) String() string {
	return "NonStructT-int"
}

//#go2hx stdgo.reflect.Reflect.makeNamedGoType
func makeNamedGoType(interface_or_namedType, haxePathToFunction string) reflect.Type

//#go2hx stdgo.reflect.Reflect.testHarness
//func testHarness()
