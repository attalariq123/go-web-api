package repository

import (
	"web-api/pkg"

	"gorm.io/gorm"
)

type Repository interface {
	Read() ([]pkg.Book, error)
	FindById(ID int) (pkg.Book, error)
	Create(book pkg.Book) (pkg.Book, error)
	Update(book pkg.Book) (pkg.Book, error)
	Delete(book pkg.Book) (pkg.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Read() ([]pkg.Book, error) {
	var books []pkg.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(ID int) (pkg.Book, error) {
	var book pkg.Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Create(book pkg.Book) (pkg.Book, error) {
	err := r.db.Create(&book).Error

	return book, err

}

func (r *repository) Update(book pkg.Book) (pkg.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book pkg.Book) (pkg.Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
