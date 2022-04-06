package service

import (
	"web-api/pkg"
	"web-api/repository"
)

type Service interface {
	Read() ([]pkg.Book, error)
	FindById(ID int) (pkg.Book, error)
	Create(book pkg.Book) (pkg.Book, error)
	Update(ID int, book pkg.Book) (pkg.Book, error)
	Delete(ID int) (pkg.Book, error)
}

type service struct {
	bookRepository repository.Repository
}

func NewService(bookRepository repository.Repository) *service {
	return &service{bookRepository}
}

func (s *service) Read() ([]pkg.Book, error) {
	return s.bookRepository.Read()
}

func (s *service) FindById(ID int) (pkg.Book, error) {
	return s.bookRepository.FindById(ID)
}

func (s *service) Create(book pkg.Book) (pkg.Book, error) {
	return s.bookRepository.Create(book)
}

func (s *service) Update(ID int, book pkg.Book) (pkg.Book, error) {
	bookRes, err := s.bookRepository.FindById(ID)

	if bookRes.ID == 0 {
		return bookRes, err
	}

	bookRes.Title = book.Title
	bookRes.Description = book.Description
	bookRes.Price = book.Price
	bookRes.Rating = book.Rating
	bookRes.Discount = book.Discount

	return s.bookRepository.Update(bookRes)
}

func (s *service) Delete(ID int) (pkg.Book, error) {
	bookRes, _ := s.bookRepository.FindById(ID)
	return s.bookRepository.Delete(bookRes)
}
