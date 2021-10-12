package controller

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/nokazn/go-api-template/controller/dto"
	"github.com/nokazn/go-api-template/model/entity"
	"github.com/nokazn/go-api-template/model/repository"
)

type taskController struct {
	tr repository.TaskRepository
}

type TaskController interface {
	Get(w http.ResponseWriter, r http.Request)
	GetAll(w http.ResponseWriter, r http.Request)
	Update(w http.ResponseWriter, r http.Request)
	Delete(w http.ResponseWriter, r http.Request)
}

func NewTaskController(tr repository.TaskRepository) TaskController {
	return &taskController{tr}
}

/*
 * ----------------------------------------------------------------------------------------------------
 */
func (tc *taskController) Get(w http.ResponseWriter, r http.Request) {
	taskId, err := strconv.Atoi(r.URL.Path)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	task, err := tc.tr.Find(taskId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	output, _ := json.MarshalIndent(task, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (tc *taskController) GetAll(w http.ResponseWriter, r http.Request) {
	tasks, err := tc.tr.FindAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var tasksResponseList []dto.TaskResponse
	for _, task := range tasks {
		tasksResponseList = append(tasksResponseList, dto.TaskResponse{
			Id:      task.Id,
			Title:   task.Title,
			Content: task.Content,
		})
	}

	var tasksResponse dto.TasksResponse
	tasksResponse.Tasks = tasksResponseList

	output, _ := json.MarshalIndent(tasksResponse.Tasks, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *taskController) Post(w http.ResponseWriter, r http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var taskRequest dto.TaskRequest
	json.Unmarshal(body, &taskRequest)

	task := entity.TaskEntity{
		Title:   taskRequest.Title,
		Content: taskRequest.Content,
	}
	id, err := tc.tr.Create(task)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *taskController) Update(w http.ResponseWriter, r http.Request) {
	taskId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var taskRequest dto.TaskRequest
	json.Unmarshal(body, &taskRequest)

	task := entity.TaskEntity{
		Id:      taskId,
		Title:   taskRequest.Title,
		Content: taskRequest.Content,
	}

	err = tc.tr.Update(task)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (tc *taskController) Delete(w http.ResponseWriter, r http.Request) {
	taskId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = tc.tr.Delete(taskId)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
