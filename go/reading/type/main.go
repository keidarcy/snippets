package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	for i := reflect.Invalid; i <= reflect.UnsafePointer; i++ {
		fmt.Printf("%d: %s\n", i, i)
	}

	var (
		// invalidA invalid
		boolA       bool
		intA        int
		int8A       int8
		int16A      int16
		int32A      int32
		int64A      int64
		uintA       uint
		uint8A      uint8
		uint16A     uint16
		uint32A     uint32
		uint64A     uint64
		uintptrA    uintptr
		float32A    float32
		float64A    float64
		complex64A  complex64
		complex128A complex128
		arrayA      [3]int
		chanA       chan int
		funcA       func()
		// interfaceA interface
		mapA           map[int]int
		ptrA           *int
		sliceA         []int
		stringA        string
		structA        struct{}
		unsafePointerA unsafe.Pointer
	)

	var list []interface{}
	list = append(list, boolA, intA, int8A, int16A, int32A, int64A, uintA, uint8A, uint16A, uint32A, uint64A, uintptrA, float32A, float64A, complex64A, complex128A, arrayA, chanA, funcA, mapA, ptrA, sliceA, stringA, structA, unsafePointerA)

	for _, v := range list {
		t := reflect.TypeOf(v)
		fmt.Printf("Type: %v, Size: %v, Align: %v\n", t.Kind(), t.Size(), t.Align())
		// chan, slice, func, map is address

		fmt.Printf("%s, Comparable: %t\n", t.Kind(), t.Comparable())
		// slice, func, map -> false
	}

	var (
		arrayB [3]int8
		chanB  chan int8
		mapB   map[int8]int8
		ptrB   *int8
		sliceB []int8
	)
	list = nil
	list = append(list, arrayB, chanB, mapB, ptrB, sliceB)
	for _, v := range list {
		t := reflect.TypeOf(v)
		fmt.Printf("%s:  %d\n", t.Kind(), t.Size())

	}

	type A1 struct {
		// byte 0~7
		a int
		// byte 8
		b int8
		// byte 9
		c int8
	}

	type A2 struct {
		// byte 0
		a int8
		// byte 1~7 not used
		// byte 8~15
		b int
		// byte 16
		c int8
	}

	var v1 A1
	var v2 A2
	fmt.Printf("struct size: %d, %d", reflect.TypeOf(v1).Size(), reflect.TypeOf(v2).Size())
}
