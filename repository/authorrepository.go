package repository

import (
	"github.com/alexgunkel/library/models"
	"errors"
	"fmt"
)

func GetAuthorById(id int64) (models.Author, error) {
	var author models.Author
	db.First(&author, id)

	if 0 == author.ID {
		return models.Author{}, errors.New("Could not find author")
	}

	return author, nil
}

func GetAuthors() ([]models.Author, error) {
	var authors []models.Author
	db.Find(&authors)

	if len(authors) <= 0 {
		return []models.Author{}, errors.New("Could not find any authors")
	}

	return authors, nil
}

func GetAuthorBooks(id int64) ([]models.Book, error) {
	var books []models.Book
	author, err := GetAuthorById(id)
	if nil != err {
		return nil, err
	}


	db.Model(&author).Related(&books, "Books")

	if len(books) <= 0 {
		return nil, errors.New("No books found for author with " + fmt.Sprintf("%d", author.ID))
	}

	return books, nil
}

func AddAuthorBook(id int64, book *models.Book) error {
	author, err := GetAuthorById(id)

	if nil != err {
		return err
	}

	db.Model(&author).Association("books").Append(book)

	return nil
}

func AddAuthor(author *models.Author) {
	db.Save(author)
}

func UpdateAuthor(id int64, author *models.Author) {
	author.ID = uint(id)
	db.Save(author)
}

func DeleteAuthor(id int64) error {
	author, err := GetAuthorById(id)

	if nil != err {
		return err
	}

	db.Delete(&author)
	return nil
}
