package builders

import (
	"fmt"
	"log"
	"os"
)

type SQLBuilder struct{}

func (s *SQLBuilder) Create(info *ServiceInfo) (Service, error) {
	rootDirectory, err := GetAppRoot()

	if err != nil {
		log.Print("Can't get app root directory for SQL builder")
		return Service{}, err
	}

	dirEnvPath := fmt.Sprintf("%s/docker_app/env/", rootDirectory)

	if _, err := os.Stat(dirEnvPath); os.IsNotExist(err) {

		err = os.MkdirAll(dirEnvPath, DirPerm)
		if err != nil {
			return Service{}, err
		}
	}

	envSqlName := fmt.Sprintf("%s.env", info.ImageName)

	envPath := dirEnvPath + envSqlName

	eb := EnvBuilder{}

	_, err = eb.Build(info.Envs, envPath)

	if err != nil {

		return Service{}, fmt.Errorf("Can't build environment variable for SQL: %w", err)
	}

	return Service{
		Image:         fmt.Sprintf("%s:%s", info.ImageName, info.Tag),
		Ports:         ConvertRelatedPortsToString(info.Ports),
		Volumes:       ConvertVolumeToString(info.Volumes),
		ContainerName: fmt.Sprintf("%s_%s", info.ImageName, info.Tag),
		EnvFile:       "env/" + envSqlName,
		DependsOn:     info.DependsOn,
		Networks:      info.Networks,
	}, nil
}
