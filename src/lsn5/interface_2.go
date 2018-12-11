package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type Test interface {
	Hello()
}

type BMW struct {
	Name string
}

func (b BMW) GetName() string {
	return b.Name
}

func (b BMW) Run() {
	fmt.Println(b.Name + "is Runing")
}

func (b BMW) DiDi() {
	fmt.Println("DiDi")
}

func (b *BMW) Hello() {
	fmt.Println("Hello")
}

func main() {
	bmw := &BMW{
		Name: "Q7",
	}
	var bi Carer
	bi = bmw
	bi.Run()
	var t1 Test
	t1 = bmw
	t1.Hello()
}
