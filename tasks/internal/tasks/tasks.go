package tasks

import "time"

type Tags []string

type Task struct {
	Id   int       `json:"id"`
	Task string    `json:"task"`
	Tags Tags      `json:"tags"`
	Due  time.Time `json:"due"` // deadline date
}

type Tasks []Tasks
