package dto

type TaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TaskResponse struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TasksResponse struct {
	Tasks []TaskRequest `json:"task"`
}
