package todo

import "context"

type Storage interface {
	Create(ctx context.Context, todo Todo) (string, error)
	Update(ctx context.Context, id string, todo Todo) error
	GetAll(ctx context.Context) ([]Todo, error)
	Delete(ctx context.Context, id string) error
}
