package service

import "bitbucket.org/go-scrapper/docker-compose/internal/docker-compose/domain/project"

type BuildService struct {
	projectRepo project.Repository
}
