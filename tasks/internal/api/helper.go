package api

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Пытаемся отконфигурировать наш API инстанс (а конкретнее - поле logger)
func (a *API) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

/*
// POST   /tasks/              :  создаёт задачу и возвращает её ID
// GET    /tasks/<taskid>      :  возвращает одну задачу по её ID
// GET    /tasks/              :  возвращает все задачи
// DELETE /tasks/<taskid>      :  удаляет задачу по ID
// DELETE /tasks/              :  удаляет все задачи
// GET    /tags/<tagname>      :  возвращает список задач с заданным тегом
// GET    /due/<yy>/<mm>/<dd>  :  возвращает список задач, запланированных на указанную дату

*/

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configreRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})

	a.router.HandleFunc("/tasks/", H_CreateTask).Methods("POST")
	a.router.HandleFunc("/tasks/{id}", H_GetTask).Methods("GET")
	a.router.HandleFunc("/tasks/", H_GetAllTasks).Methods("GET")
	a.router.HandleFunc("/tasks/{id}", H_DeleteTask).Methods("DELETE")
	a.router.HandleFunc("/tasks/", H_DeleteAllTasks).Methods("DELETE")
	a.router.HandleFunc("/tags/{id}", H_FindByTag).Methods("GET")
	a.router.HandleFunc("/due/{yy}/{mm}/{dd}", H_DueTasks).Methods("GET")

}
