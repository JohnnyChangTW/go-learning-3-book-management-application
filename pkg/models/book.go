package models

import (
	"github.com/JohnnyChangTW/go-learning-3-book-management-application/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model         // todo: What's this ? Why ?
	Name        string `gorm:""json:"name"` // todo: What's this ? Why ?
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // todo: What's this ? Why ?
}

func (b *Book) CreateBook() *Book { // todo: What is func (b *Book) CreateBook()
	db.NewRecord(b) // todo: Why this
	db.Create(&b)   // todo: And this
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
