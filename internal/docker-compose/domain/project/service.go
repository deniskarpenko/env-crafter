package project

import "bitbucket.org/go-scrapper/docker-compose/internal/docker-compose/domain/project/valueobjects"

type Service struct {
	image         *valueobjects.Image
	build         *valueobjects.Build
	containerName string
}
