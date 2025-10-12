package docker

import "bitbucket.org/go-scrapper/docker-compose/internal/builders/builders"

type MlBuilder struct{}

func (mb *MlBuilder) Build(images []builders.ServiceInfo) error {
	return nil
}
