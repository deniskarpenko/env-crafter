package entity

import "bitbucket.org/go-scrapper/docker-compose/internal/container/domain/valueobject"

type Container struct {
	ContainerName string
	Image         valueobject.Image
}
