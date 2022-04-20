# library_handler

Учебный проект на работу с http.Handler

Покрытие тестами ~ 80%

Проект работает на порту 8000


- Создание книги

/create принимает json файл
```json
{
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
```
- Поиск книги по названию

/find

- Изменение книги

/update принимает json файл
```json
{
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
```

- Список всех книг

/books

- Удаление книги по id

/delete

### Запуск 

1. go run cmd/book/main.go
2. запросы можно посылать через postman (например :8000/create)


### Полезные ссылки для изучения основ 

http.Handler - https://golangify.com/http-handler-interface

postman - https://habr.com/ru/company/maxilect/blog/596789/
