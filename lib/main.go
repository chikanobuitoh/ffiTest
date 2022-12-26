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
	return C.CString(tex + result)
}

func main() {
}

type BookDatabase struct{}
type Book struct{}

func NewBookDatabase() *BookDatabase {
	var db BookDatabase
	return &db
}

func (d *BookDatabase) GetBook(id string) string {
	return "ok"
}

func (d *BookDatabase) CreateBook(data interface{}) {}

type BookReader interface {
	GetBook(string) string
}

func fetch(b BookReader, id string) string {
	return b.GetBook(id)
}
