package bookRepository

import (
	"books-list/models"
	"database/sql"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select * from books")

	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	return book, err
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books where id = $1", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
