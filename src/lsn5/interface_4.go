package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read data")
}

func (f *File) Write() {
	fmt.Println("Write data to file")
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func Test1(r Reader) {
	r.Read()
}

func main() {
	var f *File
	var b interface{} = f
	//Test(&f)
	//Test1(&f)

	if v, ok := b.(ReadWriter); ok {
		fmt.Println("%v", ok, v)
	}
}
