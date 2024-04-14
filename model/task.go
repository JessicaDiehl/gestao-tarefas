package model

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	DueDate     time.Time
	AssignedTo  string // ID do usuário responsável pela tarefa
	CreatedAt   time.Time
}
