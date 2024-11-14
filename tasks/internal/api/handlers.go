package api

import (
	"encoding/json"
	"log"
	"net/http"
	"tesks-service/internal/storage"
	"tesks-service/internal/tasks"
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

func H_CreateTask(w http.ResponseWriter, r *http.Request) {
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
	storage.CreateTask(task)

}

func H_GetTask(w http.ResponseWriter, r *http.Request) {
}

func H_GetAllTasks(w http.ResponseWriter, r *http.Request) {
}

func H_DeleteTask(w http.ResponseWriter, r *http.Request) {
}

func H_DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
}

func H_FindByTag(w http.ResponseWriter, r *http.Request) {
}

func H_DueTasks(w http.ResponseWriter, r *http.Request) {
}
