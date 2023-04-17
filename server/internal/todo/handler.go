package todo

import (
	"net/http"
	"todo-app/internal/handlers"
	"todo-app/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	todosURL = "/todos"
	todoURL  = "/todo/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(todosURL, h.GetTodos)
	router.POST(todosURL, h.CreateTodo)
	router.DELETE(todoURL, h.DeleteTodoById)
	router.PATCH(todoURL, h.PatchTodoById)
}

func (h *handler) GetTodos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("list todos"))
}

func (h *handler) CreateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("todo created"))
}

func (h *handler) DeleteTodoById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("todo deleted"))
}

func (h *handler) PatchTodoById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("todo updated"))
}
