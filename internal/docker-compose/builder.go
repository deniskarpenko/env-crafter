package docker

import "log/slog"

func NewWebBuilder(logger *slog.Logger) (*WebBuilder, error) {
	return &WebBuilder{
		Logger: logger,
	}, nil
}
