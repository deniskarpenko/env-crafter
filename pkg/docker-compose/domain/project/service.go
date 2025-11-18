package project

import (
	valueobjects2 "bitbucket.org/go-scrapper/docker-compose/pkg/docker-compose/domain/project/valueobjects"
)

type Service struct {
	image         *valueobjects2.Image
	build         *valueobjects2.Build
	containerName string
	ports         []valueobjects2.Port
	volumes       []valueobjects2.Volume
	environment   []valueobjects2.Env
	envFiles      []valueobjects2.EnvFile
	dependsOn     []string
}
