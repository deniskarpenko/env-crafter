package project

import "context"

type Repository interface {
	Save(ctx context.Context, project *Project) error
	FindByID(ctx context.Context, id string) (*Project, error)
	FindAll(ctx context.Context) ([]*Project, error)
	Delete(ctx context.Context, id string) error
}
