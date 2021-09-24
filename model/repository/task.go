package repository

import (
	"log"

	"github.com/nokazn/go-api-template/model/entity"
)

type taskRepository struct{}

type TaskRepository interface {
	Find(id int) (task entity.TaskEntity, err error)
	FindAll() (tasks []entity.TaskEntity, err error)
	Create(task entity.TaskEntity) (id int, err error)
	Update(task entity.TaskEntity) error
	Delete(id int) error
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (t *taskRepository) Find(id int) (task entity.TaskEntity, err error) {
	task = entity.TaskEntity{}
	row, err := Db.Query("SELECT id, title, content FROM task WHERE id = ?", id)
	if err != nil {
		log.Print("Failed to get task.", err)
		return task, err
	}

	err = row.Scan(&task.Id, &task.Title, &task.Content)
	if err != nil {
		log.Print("Failed to scan row.", err)
		return task, err
	}
	return task, err
}

func (t *taskRepository) FindAll() (tasks []entity.TaskEntity, err error) {
	tasks = []entity.TaskEntity{}
	rows, err := Db.Query("SELECT id, title, content FROM task ORDER BY created_at DESC")
	if err != nil {
		log.Print("Failed to get tasks.", err)
		return tasks, err
	}

	for rows.Next() {
		task := entity.TaskEntity{}
		err = rows.Scan(&task.Id, &task.Title, &task.Content)
		if err != nil {
			log.Print("Failed to scan row.", err)
			return
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func (t *taskRepository) Create(task entity.TaskEntity) (id int, err error) {
	_, err = Db.Exec("INSERT INTO task (title, content) VALUES (?, ?)", task.Title, task.Content)
	if err != nil {
		log.Println("Failed to create task.", err)
		return id, err
	}
	err = Db.QueryRow("SELECT id FROM task ORDER BY created_at DESC LIMIT 1").Scan(&id)
	return id, err
}

func (t *taskRepository) Update(task entity.TaskEntity) error {
	_, err := Db.Exec("UPDATE task SET title = ?, content = ? WHERE id = ?", &task.Title, &task.Content, &task.Id)
	if err != nil {
		log.Print("Failed to update task.", err)
	}
	return err
}

func (t *taskRepository) Delete(id int) error {
	_, err := Db.Exec("DELETE FROM task WHERE id = ?", id)
	if err != nil {
		log.Print("Failed to delete task.", err)
	}
	return err
}
