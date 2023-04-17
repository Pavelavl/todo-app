package todo

import (
	"context"
	"todo-app/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateTodoDTO) (t Todo, err error) {
	return
}
