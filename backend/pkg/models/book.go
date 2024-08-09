package models

import (
	"time"

	"github.com/Sattuu21/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Author    string         `json:"author"`
	Type      string         `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func init() {
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	DB.Create(&b)
	return b
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	result := DB.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := DB.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var b Book
	DB.Where("ID=?", ID).Delete(b)
	return b
}
