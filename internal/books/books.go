package books

import (
	"errors"
	"strings"
	"time"
)

type book struct {
	id     int
	Name   string
	Author string
	Year   int
}

var Books []book
var idNow = 0

func ErrorForNewBook(name string, author string, year int) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if author == "" {
		return errors.New("name is empty")
	}
	if year < 1500 {
		return errors.New("year is too small")
	}
	t := time.Now()
	yearNow := t.Year()
	if year > yearNow {
		return errors.New("year is over current")
	}
	return nil
}

func NewBook(name string, author string, year int) error {
	if ErrorForNewBook(name, author, year) != nil {
		return ErrorForNewBook(name, author, year)
	}
	Books = append(Books, book{
		id:     idNow,
		Name:   name,
		Author: author,
		Year:   year,
	})
	idNow++
	return nil
}

func GetBookByName(name string) ([]book, error) {
	var FindBooks []book
	if name == "" {
		return Books, errors.New("name is empty")
	}
	for _, n := range Books {
		if name == n.Name {
			FindBooks = append(FindBooks, book{
				id:     n.id,
				Name:   n.Name,
				Author: n.Author,
				Year:   n.Year,
			})
		} else if strings.Contains(n.Name, name) {
			FindBooks = append(FindBooks, book{
				id:     n.id,
				Name:   n.Name,
				Author: n.Author,
				Year:   n.Year,
			})
		} else {
			return FindBooks, errors.New("book was not found")
		}
	}
	return FindBooks, nil
}

func UpdateBookById(id int, name string, author string, year int) (book, error) {
	if ErrorForNewBook(name, author, year) != nil {
		return book{}, ErrorForNewBook(name, author, year)
	}
	if id < 0 || id >= idNow {
		return book{}, errors.New("bad id")
	}
	if Books[id].Name == name && Books[id].Author == author && Books[id].Year == year {
		return Books[id], errors.New("nothing to update")
	}
	if Books[id].Name != name {
		Books[id].Name = name
	}
	if Books[id].Author != author {
		Books[id].Author = author
	}
	if Books[id].Year != year {
		Books[id].Year = year
	}
	return Books[id], nil
}

func GetAllBooks() []book {
	return Books
}

func DeleteBook(id int) error {
	if id < 0 || id > idNow {
		return errors.New("bad id")
	}
	Books = append(Books[:id], Books[id+1:]...)
	return nil
}
