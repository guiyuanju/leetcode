package main

import (
	"fmt"
	"reflect"
)

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	assertEq(-3, ms.GetMin())
	ms.Pop()
	assertEq(0, ms.Top())
	assertEq(-2, ms.GetMin())
}

func assertEq(a, b any) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("Ok: %v\n", a)
	} else {
		fmt.Printf("Failed: %v != %v\n", a, b)
	}
}

type MinStack struct {
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	if len(this.mins) > 0 {
		this.mins = append(this.mins, min(val, this.mins[len(this.mins)-1]))
	} else {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	this.vals = this.vals[:len(this.vals)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
