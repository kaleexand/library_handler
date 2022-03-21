package books

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBook(t *testing.T) {
	err := NewBook("Мастер и Маргарита", "Булгаков", 1400)
	assert.Error(t, err)
	err = NewBook("", "Булгаков", 2022)
	assert.Error(t, err)
	err = NewBook("Мастер и Маргарита", "", 2022)
	assert.Error(t, err)
	err = NewBook("Мастер и Маргарита", "Булгаков", 2024)
	assert.Error(t, err)
	err = NewBook("Мастер и Маргарита", "Булгаков", 2022)
	assert.NoError(t, err)
}

func TestGetBookByName(t *testing.T) {
	NewBook("Мастер и Маргарита", "Булгаков", 2022)
	_, err := GetBookByName("Мастер и Маргарита")
	assert.NoError(t, err)
	_, err = GetBookByName("")
	assert.Error(t, err)
	_, err = GetBookByName("Лиса")
	assert.Error(t, err)
}

func TestUpdateBookById(t *testing.T) {
	NewBook("Мастер и Маргарита", "Булгаков", 2022)
	NewBook("Идиот", "Достоевский", 2020)

	_, err := UpdateBookById(5, "Мастер и Маргарита", "Булгаков", 2022)
	assert.Error(t, err)
	_, err = UpdateBookById(0, "Мастер и Маргарита", "Булгаков", 2022)
	assert.Error(t, err)
	_, err = UpdateBookById(1, "", "Булгаков", 2022)
	assert.Error(t, err)
	_, err = UpdateBookById(1, "Мастер и Маргарита", "", 2022)
	assert.Error(t, err)
	_, err = UpdateBookById(1, "Мастер и Маргарита", "Булгаков", 2024)
	assert.Error(t, err)
	_, err = UpdateBookById(-5, "Мастер и Маргарита", "Булгаков", 2020)
	assert.Error(t, err)
	_, err = UpdateBookById(0, "Мастер и Маргарита", "Булгаков", 2022)
	assert.Error(t, err)
	_, err = UpdateBookById(0, "Мастер и Маргарита", "Булгаков", 1000)
	assert.Error(t, err)
	_, err = UpdateBookById(0, "Бег", "Булгаков", 2022)
	assert.NoError(t, err)
}
func TestDeleteBook(t *testing.T) {
	err := DeleteBook(1)
	assert.NoError(t, err)
	err = DeleteBook(10)
	assert.Error(t, err)
	err = DeleteBook(-9)
	assert.Error(t, err)
}
