package main

import (
	"fmt"
)


type P115 struct {
}


type MinStack struct {
	Data []int
	MinData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack {
		Data: []int{},
		MinData: []int{},
	}
}

func (this *MinStack) Push(val int)  {
	this.Data = append(this.Data, val)
	// <=: to avoid duplicate min num insert
	if len(this.MinData) == 0 || val <= this.GetMin() {
		this.MinData = append(this.MinData, val)
	}
}


func (this *MinStack) Pop()  {
	if len(this.Data) <= 0 {
		return
	}

	if this.Top() == this.GetMin() {
		this.MinData = this.MinData[: len(this.MinData) - 1]
	}

	this.Data = this.Data[: len(this.Data) - 1]
}


func (this *MinStack) Top() int {
	return this.Data[len(this.Data) - 1]
}


func (this *MinStack) GetMin() int {
	return this.MinData[len(this.MinData) - 1]
}


func execute1() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}


func (p P115)execute1() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}


func (p P115)execute2() {
	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}


func main() {
	var p P115
	p.execute1()
	p.execute2()
}
