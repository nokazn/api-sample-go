package dto

type TaskRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TaskResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TasksResponse struct {
	Tasks []TaskResponse `json:"task"`
}
