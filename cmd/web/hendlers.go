package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" { // Обработчик 404 страница не найдена
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(404)
		w.Write([]byte("Страница не найдена"))
		return
	}
	w.Write([]byte("Hello"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Отображение заметки..."))
	fmt.Println(id)

}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод запрещен", 405)
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}
