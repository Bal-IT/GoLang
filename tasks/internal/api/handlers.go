package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tesks-service/internal/storage"
	"tesks-service/internal/tasks"

	"github.com/gorilla/mux"
)

/*
Реализовать REST API сервис для задач и использовать в качестве хранилища - sqlite db:
Можно использовать стандартный Mux, можно Gorilla/mux

// POST   /tasks/              :  создаёт задачу и возвращает её ID
// GET    /tasks/<taskid>      :  возвращает одну задачу по её ID
// GET    /tasks/              :  возвращает все задачи
// DELETE /tasks/<taskid>      :  удаляет задачу по ID
// DELETE /tasks/              :  удаляет все задачи
// GET    /tags/<tagname>      :  возвращает список задач с заданным тегом
// GET    /due/<yy>/<mm>/<dd>  :  возвращает список задач, запланированных на указанную дату


type task struct {
	id int
	Text string
	Tags []string
	Due time.Time // deadline date
}
*/

type Message struct {
	Message string
}

type LastInsertId struct {
	Id int `json:"id"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func H_CreateTask(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Creating new task ...")
	var task tasks.Task

	// Check json file
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400 error
		message := Message{Message: "provided json file is invalid."}
		json.NewEncoder(w).Encode(message)
		return
	}
	id := storage.CreateTask(task)

	if id < 1 {
		w.WriteHeader(http.StatusBadRequest) // 400 error
		message := Message{Message: "Что-то пошо не так"}
		json.NewEncoder(w).Encode(message)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	lastInsertId := LastInsertId{Id: id}
	json.NewEncoder(w).Encode(lastInsertId)
}

func H_GetTask(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		w.WriteHeader(http.StatusBadRequest) // 400 error
		message := Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(w).Encode(message)
		return
	}

	task, ok := storage.GetTask(id)
	log.Println("Get task with id:", id)
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404 error
		message := Message{Message: "task with that ID does not exist in database."}
		json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusOK) // 200
		json.NewEncoder(w).Encode(task)
	}

}

func H_GetAllTasks(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	tasks, ok := storage.GetAllTask()
	log.Println("Get all task")
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404 error
		message := Message{Message: "task with that ID does not exist in database."}
		json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusOK) // 200
		json.NewEncoder(w).Encode(tasks)
	}
}

func H_DeleteTask(w http.ResponseWriter, r *http.Request) {
}

func H_DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
}

func H_FindByTag(w http.ResponseWriter, r *http.Request) {
}

func H_DueTasks(w http.ResponseWriter, r *http.Request) {
}
