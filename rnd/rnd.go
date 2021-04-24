package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("start")

	// Code below highlights what works so far in the reflect prototype
	str := "hitchhiker"
	// showType(42)
	// showType(uint8(42))
	// showType(str)
	// showType(&str)
	// spt := reflect.PtrTo(reflect.TypeOf(str))
	// fmt.Println(spt)
	// showType(struct {
	// 	A int32
	// 	b string
	// }{42, str})

	m := MyType{
		X: 42,
		y: str,
	}
	showType(m)
	mt := reflect.TypeOf(m)
	// fmt.Println("MyType fields:")
	// for i := 0; i < mt.NumField(); i++ {
	// 	fmt.Println(mt.Field(i))
	// }
	// fmt.Println("MyType methods:")
	// for i := 0; i < mt.NumMethod(); i++ {
	// 	fmt.Println(mt.Method(i))
	// }
	showType(mt)

	//nsi := NonStructT(42)
	//showType(nsi)

	// line below does not work in Go, it would neet to be generated by go2hx
	// var ms = makeNamedGoType("interface", "github_com.pxshadow.go4hx.rnd.MyStringer")

	// fmt.Println(ms.String(), " methods:")
	// for i := 0; i < ms.NumMethod(); i++ {
	// 	fmt.Println(ms.Method(i))
	// }
	// fmt.Println(mt.Implements(ms)) // this is the key requirement for a type switch

	var k reflect.Kind
	k = 5
	showType(k)

	//testHarness()
}

func showType(i interface{}) {
	Ti := reflect.TypeOf(i)
	Ki := Ti.Kind()
	Kui := uint(Ki)
	Kn := Ki.String()
	fmt.Println(i, Ti.String(), Kui, Kn)
}

type MyType struct {
	X uint8
	y string
}

func (mt MyType) String() string {
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
//func makeNamedGoType(interface_or_namedType, haxePathToFunction string) reflect.Type

//#go2hx stdgo.reflect.Reflect.testHarness
//func testHarness()
