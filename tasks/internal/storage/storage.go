package storage

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"tesks-service/internal/tasks"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

/*
type Storage struct {
	db *sql.DB
}

func New() *Storage {
	conn, err := sql.Open("sqlite", "./todo.db")
	if err != nil {
		log.Fatalf("Connect to database failed: %v\n", err)
	}

	return &Storage{db: conn}
}


func (s *Storage) CreateDatabase() {
	// Create table block
	// SQL statement to create table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, due TEXT, checked INTEGER);
	`
	// Execute SQL statement
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt = `
	CREATE TABLE IF NOT EXISTS tags (id INTEGER PRIMARY KEY AUTOINCREMENT, taskid INTEGER, tag TEXT);
	`
	// Execute SQL statement
	_, err = s.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
*/

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Init() {
	// os.Remove("./todo.db")
	// Open()
	// createDatabase()
}

func Open() {
	var err error
	db, err = sql.Open("sqlite", "./todo.db")

	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

func CreateTask(task tasks.Task) int {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Prepare SQL statement than can be reused. Char "?" for SQLite, char "%" for MySQL, PostgreSQL
	stmt, err := tx.Prepare("INSERT INTO tasks(task, due, checked) VALUES(?, ?, 0)")
	if err != nil {
		log.Fatal(err)
	}
	// close prepared statement before exiting program
	defer stmt.Close()

	_, err = stmt.Exec(task.Task, task.Due.Format(time.DateTime))
	if err != nil {
		log.Fatal(err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// Получаем ID
	rows, err := db.Query("SELECT last_insert_rowid()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lastId int
	for rows.Next() {
		rows.Scan(&lastId)
	}

	if len(task.Tags) == 0 {
		return lastId
	}

	// Теги
	tx2, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = tx2.Prepare("INSERT INTO tags(taskid, tag) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	//
	for i := range task.Tags {
		_, err = stmt.Exec(lastId, task.Tags[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := tx2.Commit(); err != nil {
		log.Fatal(err)
	}

	return lastId
}

func GetTask(id int) (tasks.Task, bool) {
	stmt, err := db.Prepare(`select t.id, t.task, t.due, group_concat(tg.tag) as tag
from tasks t
left join tags tg on t.id = tg.taskid
where t.id = ?
group by t.id, t.task, t.due 
		`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	// Rows must be closed
	defer rows.Close()

	var ts tasks.Task

	ok := false
	for rows.Next() {
		var id int
		var task string
		var sdue string
		var tag string
		// use pointers to get data
		err = rows.Scan(&id, &task, &sdue, &tag)
		if err != nil {
			log.Fatal(err)
		}

		ts.Id = id
		ts.Task = task
		ts.Tags = strings.Split(tag, ",")

		date, err := time.Parse(time.DateTime, sdue)
		if err == nil {
			ts.Due = date
		}

		ok = true
	}
	if !ok {
		return ts, false
	}

	return ts, true
}

func GetAllTask() (tasks.Tasks, bool) {
	stmt, err := db.Prepare(`select t.id, t.task, t.due, group_concat(tg.tag) as tag
from tasks t
left join tags tg on t.id = tg.taskid
group by t.id, t.task, t.due 
		`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	// Rows must be closed
	defer rows.Close()

	var ts tasks.Task
	var tasks tasks.Tasks

	for rows.Next() {
		var id int
		var task string
		var sdue string
		var tag *string
		// use pointers to get data
		err = rows.Scan(&id, &task, &sdue, &tag)
		if err != nil {
			log.Fatal(err)
		}

		ts.Id = id
		ts.Task = task

		if tag != nil {
			ts.Tags = strings.Split(*tag, ",")
		}

		date, err := time.Parse(time.DateTime, sdue)
		if err == nil {
			ts.Due = date
		}

		tasks = append(tasks, ts)
	}
	if len(tasks) == 0 {
		return tasks, false
	}

	return tasks, true
}

func DeleteTask(id int) error {

	tx, err := db.Begin()

	stmt, err := db.Prepare("delete from tasks where id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("delete from tags where taskid = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return nil
}
