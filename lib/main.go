package main

import (
	"C"
)

//export ffiCheck
func ffiCheck(mes *C.char) *C.char {
	var db = NewBookDatabase()
	var id = "abcde12345"
	result := fetch(db, id)

	tex := C.GoString(mes)
	return C.CString(tex + result.Bookname)
}

func main() {
}

type BookDatabase struct{}
type Book struct {
	Bookname string "hogehoge"
}

func NewBookDatabase() *BookDatabase {
	var db BookDatabase
	return &db
}

func (d *BookDatabase) GetBook(id string) *Book {
	return &Book{Bookname: "okok"}
}

func (d *BookDatabase) CreateBook(data interface{}) {}

type BookReader interface {
	GetBook(string) *Book
}

func fetch(b BookReader, id string) *Book {
	return b.GetBook(id)
}
