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
	//
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//
}
