package repository

import (
	"github.com/alexgunkel/library/models"
	"errors"
	"fmt"
)

func GetBookById(id int64) (models.Book, error) {
	var book models.Book
	db.First(&book, id)

	if 0 == book.ID {
		return models.Book{}, errors.New("Could not find book")
	}

	return book, nil
}

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	db.Find(&books)

	if len(books) <= 0 {
		return []models.Book{}, errors.New("Could not find any books")
	}

	return books, nil
}

func GetBookAuthors(id int64) ([]models.Author, error) {
	var authors []models.Author
	book, err := GetBookById(id)
	if nil != err {
		return nil, err
	}


	db.Model(&book).Related(&authors, "Authors")

	if len(authors) <= 0 {
		return nil, errors.New("No authors found for book with " + fmt.Sprintf("%d", book.ID))
	}

	return authors, nil
}

func AddBookAuthor(bookId int64, author *models.Author) error {
	book, err := GetBookById(bookId)

	if nil != err {
		return err
	}

	db.Model(&book).Association("Authors").Append(author)

	return nil
}

func AddBook(book *models.Book) {
	db.Save(book)
}

func UpdateBook(id int64, book *models.Book) {
	book.ID = uint(id)
	db.Save(book)
}

func DeleteBook(id int64) error {
	book, err := GetBookById(id)

	if nil != err {
		return err
	}

	db.Delete(&book)
	return nil
}
