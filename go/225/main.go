package main

import "fmt"

type MyStack struct {
	a []int
	b []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.a = append(this.a, x)
}

func (this *MyStack) Pop() int {
	for len(this.a) > 1 {
		this.b = append(this.b, this.a[0])
		this.a = this.a[1:]
	}
	res := this.a[0]
	this.a = this.a[:0]
	this.a, this.b = this.b, this.a
	return res
}

func (this *MyStack) Top() int {
	res := this.Pop()
	this.Push(res)
	return res
}

func (this *MyStack) Empty() bool {
	return len(this.a) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Top()
	param_2 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}
