package book

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func CreateBook(name string, total int, author, createTime string) (book *Book) {
	book := &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
	return
}

func (b *Book) CanBorrow(c int) bool {
	return b.Total >= c
}

func (b *Book) Borrow(c int) (err error) {
	if b.CanBorrow(c) == false {
		err = error.New("Stock is not enough")
		return
	}

	b.Total -= c
	return
}

func (b *Book) Back(c int) (err error) {
	b.Total += c
	return

}
