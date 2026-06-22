package handlers

import (
	"books/helpers"
	"books/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Задаем заголовок ответа
	helpers.InitHeader(w)

	// Вытаскиваем идетификатор книги из пути запроса
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		msg := fmt.Sprintf("Ошибка в передаче идентификатора %s книги", mux.Vars(r)["id"])
		log.Println(msg) // лоигруемся для себя
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helpers.Message{
			Message: msg,
		})
		return
	}

	// Ищем книгу по идентификатору
	book, ok := models.FindBookById(id)

	// Негативный сценарий - не нашли книгу в БД
	if !ok {
		msg := fmt.Sprintf("Книга с идентификатором %d не найдена", id)
		log.Println(msg) // логируемся для себя
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(helpers.Message{
			Message: msg,
		})
		return
	}

	// Позитивный сценарий - нашли книгу в БД
	if ok {
		msg := fmt.Sprintf("Книга с идентификатором %d найдена", id)
		log.Println(msg) // логируемся
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(book)
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	helpers.InitHeader(w)

	// Вычитываем книгу из тела запроса
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		msg := fmt.Sprintln("Ошибка с JSON-файлом в запросе")
		log.Println(msg)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helpers.Message{
			Message: msg,
		})
		return
	}

	// Добавляем книгу в БД
	book.Id = len(models.Db) + 1
	models.AddBook(book)

	// Возвращаем клиенту книгу, которую он пытается добавить в БД
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	helpers.InitHeader(w)
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		msg := fmt.Sprintf("Ошибка в передаче идентификатора %s книги", mux.Vars(r)["id"])
		log.Println(msg) // лоигруемся для себя
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helpers.Message{
			Message: msg,
		})
		return
	}

	_, ok := models.FindBookById(id)

	// Негативный сценарий - не нашли книгу в БД
	if !ok {
		msg := fmt.Sprintf("Книга с идентификатором %d не найдена", id)
		log.Println(msg) // логируемся для себя
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(helpers.Message{
			Message: msg,
		})
		return
	}

	var newBook models.Book
	if ok {
		err := json.NewDecoder(r.Body).Decode(&newBook)

		if err != nil {
			msg := fmt.Sprintln("Ошибка с JSON-файлом в запросе")
			log.Println(msg)
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(helpers.Message{
				Message: msg,
			})
			return
		}
	}

	// Обновляем книгу в БД
	newBook.Id = id // обновляем идентификатор новой книги искомым идентификатором
	models.UpdateBook(id, newBook)

	// Возвращаем ответ клиенту
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(newBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//
}
