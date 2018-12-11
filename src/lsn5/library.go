package main

import "fmt"

type Book struct {
	Name        string
	Repliset    int
	Author      string
	PublishDate string
	BorrowFrom  string
}

type Student struct {
	Name           string
	Level          string
	ID             string
	Sex            string
	BorrowBookList []Book
}

type Library struct {
	Book
	Student
}

func (b *Book) SearchBook(condition ...string) {
	for i := 0; i < len(condition); i++ {
		fmt.Println("search ", condition[i])
	}
}

func (b *Book) BorrowBook(student string, bookList ...string) {
	for i := 0; i < len(bookList); i++ {
		bookInfo := &Book{
			BorrowFrom: student,
			Name:       bookList[i],
		}
		fmt.Println("borrow ", bookInfo)
	}
}

func (b *Book) ManageBook(bookList ...string) (bookInfo map[string]string) {
	bookInfo = make(map[string]string)
	for i := 0; i < len(bookList); i++ {
		fmt.Println("book ", bookList[i])
		bookInfo[bookList[i]] = "shuizhihan"
	}
	return
}

func main() {
	var lib Library
	lib.SearchBook("ming", "zong", "yue")
	lib.BorrowBook("ming", "zong", "yue")
	lib.ManageBook("ming", "zong", "yue")
}
