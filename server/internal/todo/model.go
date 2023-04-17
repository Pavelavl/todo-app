package todo

type Todo struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title,omitempty"`
	Completed bool   `json:"completed" bson:"completed"`
}

type CreateTodoDTO struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
}
