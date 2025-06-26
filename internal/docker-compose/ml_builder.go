package docker

import "bitbucket.org/go-scrapper/docker-compose/internal/docker-compose/builders"

type MlBuilder struct{}

func (mb *MlBuilder) Build(images []builders.ServiceInfo) error {
	return nil
}
