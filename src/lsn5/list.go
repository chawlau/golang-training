package main

import "fmt"

type LinkNode struct {
	Next *LinkNode
	Data interface{}
}

type Link struct {
	Head *LinkNode
	Tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	newNode := &LinkNode{
		Data: data,
		Next: nil,
	}

	if p.Tail == nil && p.Head == nil {
		p.Tail = newNode
		p.Head = newNode
		return
	}

	newNode.Next = p.Head
	p.Head = newNode
}

func (p *Link) InsertTail(data interface{}) {
	newNode := &LinkNode{
		Data: data,
		Next: nil,
	}

	if p.Tail == nil {
		p.Tail = newNode
		p.Head = newNode
		return
	}

	p.Tail.Next = newNode
	p.Tail = newNode
}

func (p *Link) Scan() {
	head := p.Head
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

func main() {
	var intLink Link
	for i := 0; i < 10; i++ {
		intLink.InsertTail(i)
	}
	intLink.Scan()
}
