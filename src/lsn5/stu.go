package stu

type Student struct {
	Name  string
	Grade string
	ID    string
	Sex   string
	Books []*Book
}

func CreateStudent(name, grade, id, sex string) *Student {
	stu := &Student{
		Name:  name,
		Grade: grade,
		ID:    id,
		Sex:   sex,
	}
	return stu
}

func (s *Student) AddBook(b *Book) {
	s.Books = append(s.Books, b)
}
